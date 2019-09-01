package mikrotik

import (
	"errors"
	"fmt"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	routeros "gopkg.in/routeros.v2"
)

//Dial to mikrotik router
func Dial(addr, user, pass string) (*Mikrotik, error) {
	c, err := routeros.Dial(addr, user, pass)
	if err != nil {
		return nil, err
	}

	mik := &Mikrotik{Conn: c}
	mik.setMikrotikCommands()

	return mik, nil
}

//DialTimeout dial to mikrotik router with timeout
func DialTimeout(addr, user, pass string, timeout time.Duration) (*Mikrotik, error) {
	c, err := routeros.DialTimeout(addr, user, pass, timeout)
	if err != nil {
		return nil, err
	}

	mik := &Mikrotik{Conn: c}
	mik.setMikrotikCommands()

	return mik, nil
}

//Mikrotik is common struct contains connection to device and API-tree
type Mikrotik struct {
	Conn      *routeros.Client
	connMutex sync.Mutex

	IP        ip
	System    system
	Interface netinterface
	PPP       ppp

	debug bool
}

func (mik *Mikrotik) Debug(debug bool) {
	mik.debug = debug
}

func (mik *Mikrotik) Close() {
	mik.Conn.Close()
}

func (mik *Mikrotik) setMikrotikCommands() {
	mik.IP = ip{
		Address: cmd{mikrotik: mik, path: "/ip/address"},
		Route:   cmd{mikrotik: mik, path: "/ip/route"},
		Firewall: firewall{
			NAT:    cmd{mikrotik: mik, path: "/ip/firewall/nat"},
			Mangle: cmd{mikrotik: mik, path: "/ip/firewall/mangle"},
		},
	}

	mik.System = system{
		mikrotik: mik,
		path:     "/system",
		Identity: identity{mikrotik: mik, path: "/system/identity"},
		NTP: ntp{
			Client: cfg{mikrotik: mik, path: "/system/ntp/client"},
		},
	}

	mik.Interface = netinterface{
		mikrotik:   mik,
		path:       "/interface",
		SSTPServer: cmd{mikrotik: mik, path: "/interface/sstp-server"},
		SSTPClient: cmd{mikrotik: mik, path: "/interface/sstp-client"},
		Wireless: wireless{
			cmd: cmd{
				mikrotik: mik,
				path:     "/interface/wireless",
			},
			SecurityProfiles: cmd{mikrotik: mik, path: "/interface/wireless/security-profiles"},
		},
		Lte: lte{
			mikrotik: mik,
			path:     "/interface/lte",
		},
	}

	mik.PPP = ppp{
		AAA:        cfg{mikrotik: mik, path: "/ppp/aaa"},
		Secret:     cmd{mikrotik: mik, path: "/ppp/secret"},
		L2tpSecret: cmd{mikrotik: mik, path: "/ppp/l2tp-secret"},
		Profile:    cmd{mikrotik: mik, path: "/ppp/profile"},
	}
}

func (mik *Mikrotik) Ping(addr string, count int) *Ping {
	return &Ping{mikrotik: mik, Address: addr, Count: count}
}

func (ping *Ping) Start() ([]*PingResponse, error) {
	re, err := ping.mikrotik.RunArgs("/ping", ToArgs(ping)...)
	if err != nil {
		return nil, err
	}

	var pingResp []*PingResponse
	for _, resp := range re.Re {
		ValuesFrom(resp.Map).To(&pingResp)
	}

	if len(pingResp) > 0 && pingResp[0].PacketLoss == 100 {
		return pingResp, errors.New("timeout")
	}

	return pingResp, nil
}

//Run one line command on mikrotik by api
func (mik *Mikrotik) Run(cmd string) (*routeros.Reply, error) {
	mik.connMutex.Lock()
	defer mik.connMutex.Unlock()

	log.Tracef("[Run] %v", cmd)
	re, err := mik.Conn.Run(cmd)
	log.Tracef("[Run](reply) %+v", re)
	if err != nil {
		mik.Conn.Run("")
	}

	return re, err
}

//RunArgs run many line command on mikrotik by api
func (mik *Mikrotik) RunArgs(cmd string, args ...string) (*routeros.Reply, error) {
	mik.connMutex.Lock()
	defer mik.connMutex.Unlock()
	toRun := append([]string{cmd}, args...)
	log.Tracef("[RunArgs] %v", toRun)
	re, err := mik.Conn.RunArgs(toRun)

	log.Tracef("[RunArgs](reply) %+v", re)
	if err != nil {
		mik.Conn.Run("")
	}
	return re, err
}

//RunMarshal - run command and marhsal response to interface struct
func (mik *Mikrotik) RunMarshal(cmd string, v interface{}) error {
	re, err := mik.Run(cmd)
	if err != nil {
		return err
	}

	return mik.ParseResponce(re, v)
}

func (mik *Mikrotik) ParseResponce(re *routeros.Reply, v interface{}) error {
	for _, resp := range re.Re {
		if mik.debug {
			log.Debug(resp)
		}

		if err := ValuesFrom(resp.Map).To(v); err != nil {
			return err
		}
	}

	return nil
}

//Print returns all items by apipath and marshal it to passed structure
func (mik *Mikrotik) Print(apipath string, v interface{}) error {
	return mik.RunMarshal(apipath, v)
}

//Add item from passed struct to apipath
func (mik *Mikrotik) Add(apipath string, v interface{}) error {
	re, err := mik.RunArgs(apipath, ToArgs(v)...)
	if err != nil {
		return err
	}

	if id, ok := re.Done.Map["ret"]; ok {
		SetID(v, id)
	}

	return nil
}

//Set value of item by id
func (mik *Mikrotik) Set(apipath, id string, v interface{}) error {
	args := append([]string{"=.id=" + id}, ToArgs(v)...)
	_, err := mik.RunArgs(apipath, args...)
	return err
}

//SetOne set value to one field
func (mik *Mikrotik) SetOne(apipath, name, value string) error {
	_, err := mik.RunArgs(apipath, fmt.Sprintf("=%s=%s", name, value))
	return err
}

//Remove item by id
func (mik *Mikrotik) Remove(apipath, id string) error {
	_, err := mik.RunArgs(apipath, "=.id="+id)
	return err
}

//Enable item by id
func (mik *Mikrotik) Enable(apipath, id string) error {
	_, err := mik.RunArgs(apipath, "=.id="+id)
	return err
}

//Disable item by id
func (mik *Mikrotik) Disable(apipath, id string) error {
	_, err := mik.RunArgs(apipath, "=.id="+id)
	return err
}

//Comment - add comment to item by id
func (mik *Mikrotik) Comment(apipath, id, comment string) error {
	_, err := mik.RunArgs(apipath, "=.id="+id, "=comment="+comment)
	return err
}

// ====================================
//
// API tree
//
// ====================================

type ip struct {
	Address  cmd
	Route    cmd
	Firewall firewall
}
type firewall struct {
	NAT    cmd
	Mangle cmd
}

type cmd struct {
	mikrotik *Mikrotik
	path     string
}

func (c *cmd) List(v interface{}) error {
	return c.mikrotik.Print(c.path+"/print", v)
}

func (c *cmd) Find(where string, v interface{}) error {
	re, err := c.mikrotik.RunArgs(c.path+"/print", "?"+where)
	if err != nil {
		return err
	}

	return c.mikrotik.ParseResponce(re, v)
}

// func (c *cmd) Get(id string, v interface{}) error {
// 	re, err := c.mikrotik.RunArgs(c.path+"/print", "where", "?.id="+id)
// 	if err != nil {
// 		return err
// 	}

// 	return c.mikrotik.ParseResponce(re, v)
// }

func (c *cmd) Add(v interface{}) error {
	return c.mikrotik.Add(c.path+"/add", v)
}

func (c *cmd) Set(id string, v interface{}) error {
	return c.mikrotik.Set(c.path+"/set", id, v)
}

func (c *cmd) Remove(id string) error {
	return c.mikrotik.Remove(c.path+"/remove", id)
}

func (c *cmd) Enable(id string) error {
	return c.mikrotik.Enable(c.path+"/enable", id)
}

func (c *cmd) Disable(id string) error {
	return c.mikrotik.Disable(c.path+"/disable", id)
}

func (c *cmd) Comment(id, comment string) error {
	return c.mikrotik.Comment(c.path+"/comment", id, comment)
}

type system struct {
	mikrotik *Mikrotik
	path     string

	Identity identity
	NTP      ntp
}

func (s *system) Routerboard(v interface{}) error {
	return s.mikrotik.Print(s.path+"/routerboard/print", v)
}

type ntp struct {
	Client cfg
}

type cfg struct {
	mikrotik *Mikrotik
	path     string
}

func (c *cfg) Get(v interface{}) error {
	return c.mikrotik.Print(c.path+"/print", v)
}

func (c *cfg) Set(name, value string) error {
	return c.mikrotik.SetOne(c.path+"/set", name, value)
}

type identity struct {
	mikrotik *Mikrotik
	path     string
}

func (c *identity) Name() (string, error) {
	var resp struct {
		Name string
	}
	err := c.mikrotik.Print(c.path+"/print", &resp)
	return resp.Name, err
}

func (c *identity) SetName(name string) error {
	return c.mikrotik.SetOne(c.path+"/set", "name", name)
}

//netinterface not have add method
type netinterface struct {
	mikrotik *Mikrotik
	path     string

	SSTPClient cmd
	SSTPServer cmd
	Wireless   wireless
	Lte        lte
}

func (c *netinterface) List(v interface{}) error {
	return c.mikrotik.Print(c.path+"/print", v)
}

func (c *netinterface) Find(where string, v interface{}) error {
	re, err := c.mikrotik.RunArgs(c.path+"/print", "?"+where)
	if err != nil {
		return err
	}

	return c.mikrotik.ParseResponce(re, v)
}

// func (c *netinterface) GetByName(name string, v interface{}) error {
// 	re, err := c.mikrotik.RunArgs(c.path+"/print", "?name="+name)
// 	if err != nil {
// 		return err
// 	}

// 	return c.mikrotik.ParseResponce(re, v)
// }

func (c *netinterface) Set(id string, v interface{}) error {
	return c.mikrotik.Set(c.path+"/set", id, v)
}

func (c *netinterface) Remove(id string) error {
	return c.mikrotik.Remove(c.path+"/remove", id)
}

func (c *netinterface) Enable(id string) error {
	return c.mikrotik.Enable(c.path+"/enable", id)
}

func (c *netinterface) Disable(id string) error {
	return c.mikrotik.Disable(c.path+"/disable", id)
}

func (c *netinterface) Comment(id, comment string) error {
	return c.mikrotik.Comment(c.path+"/comment", id, comment)
}

type wireless struct {
	// mikrotik *Mikrotik
	// path     string
	cmd

	SecurityProfiles cmd
}

func (c *wireless) Scan(name, duration string) (APlist []*WirelessAP, err error) {
	re, err := c.mikrotik.RunArgs(c.path+"/scan", "=.id="+name, "=duration="+duration)
	if err != nil {
		return nil, err
	}

	var list []*WirelessAP
	err = c.mikrotik.ParseResponce(re, &list)
	if err != nil {
		return list, err
	}

	//remove duplicates
	for i := range list {
		if c.contains(APlist, list[i]) {
			continue
		}
		APlist = append(APlist, list[i])
	}

	return
}

func (c *wireless) contains(list []*WirelessAP, ap *WirelessAP) bool {
	for i := range list {
		if list[i].Address == ap.Address && list[i].SSID == ap.SSID {
			return true
		}
	}
	return false
}

type lte struct {
	mikrotik *Mikrotik
	path     string
}

func (l *lte) Set(id string, v interface{}) error {
	path := l.path
	return l.mikrotik.Set(path+"/set", id, v)
}

func (l *lte) InfoOnce(id string, lteInfo *LteInfo) error {
	res, err := l.mikrotik.RunArgs(l.path+"/info", "=.id="+id, "=once=")
	if err != nil {
		return err
	}

	err = l.mikrotik.ParseResponce(res, lteInfo)
	if err != nil {
		return err
	}

	return nil
}

func (l *lte) List(v interface{}) error {
	return l.mikrotik.Print(l.path+"/print", v)
}

type ppp struct {
	AAA        cfg
	L2tpSecret cmd
	Profile    cmd
	Secret     cmd
}
