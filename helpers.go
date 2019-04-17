package mikrotik

import (
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
