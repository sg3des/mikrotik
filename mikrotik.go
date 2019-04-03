package mikrotik

import (
	"errors"
	"sync"

	routeros "gopkg.in/routeros.v2"
)

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

type ip struct {
	Address  cmd
	Route    cmd
	Firewall firewall
}
type firewall struct {
	NAT    cmd
	Mangle cmd
}

type ppp struct {
	AAA        cfg
	L2tpSecret cmd
	Profile    cmd
	Secret     cmd
}

// ====================================
//
// Mikrotik Utilities
//
// ====================================

// Debug activates the debug mode on the library
func (mik *Mikrotik) Debug(debug bool) {
	mik.debug = debug
}

// Close closes the connection with the Mikrotik
func (mik *Mikrotik) Close() {
	mik.Conn.Close()
}

// setMikrotikCommands sets the relative paths of the commands.
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
