package mikrotik

import (
	"log"
	"time"

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

//Run one line command on mikrotik by api
func (mik *Mikrotik) Run(cmd string) (*routeros.Reply, error) {
	mik.connMutex.Lock()
	defer mik.connMutex.Unlock()

	re, err := mik.Conn.Run(cmd)
	if err != nil {
		mik.Conn.Run("")
	}
	return re, err
}

//RunArgs run many line command on mikrotik by api
func (mik *Mikrotik) RunArgs(cmd string, args ...string) (*routeros.Reply, error) {
	mik.connMutex.Lock()
	defer mik.connMutex.Unlock()

	re, err := mik.Conn.RunArgs(append([]string{cmd}, args...))

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
			log.Println(resp)
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
