package mikrotik

import "net"

// ====================================
//
// Profile
//
// ====================================

type PPPProfile struct {
	ID             string `mikrotik:".id"`
	Name           string
	UseMPLS        string `mikrotik:"use-mpls"`
	UseCompression string `mikrotik:"use-compression"`
	UseEncryption  string `mikrotik:"use-encryption"`
	OnlyOne        string `mikrotik:"only-one"`
	ChangeTCPMSS   string `mikrotik:"change-tcp-mss"`
	UseUPNP        string `mikrotik:"use-upnp"`
	AddressList    string `mikrotik:"address-list"`
	OnUp           string `mikrotik:"on-up"`
	OnDown         string `mikrotik:"on-down"`
	Default        bool   `mikrotik:"default"`
}

// ====================================
//
// Secret
//
// ====================================

type PPPSecret struct {
	ID            string `mikrotik:".id"`
	Name          string
	Password      string
	Service       string
	CallerID      string `mikrotik:"caller-id"`
	Profile       string
	LocalAddress  net.IP `mikrotik:"local-address"`
	RemoteAddress net.IP `mikrotik:"remote-address"`
	Routes        string
	LimitBytesIn  int    `mikrotik:"limit-bytes-in"`
	LimitBytesOut int    `mikrotik:"limit-bytes-out"`
	LastLoggedOut string `mikrotik:"last-logged-out,ro"`
	Disabled      bool
	Comment       string
}

// ====================================
//
// Active
//
// ====================================

type PPPActive struct {
	ID            string `mikrotik:"id"`
	Address       string `mikrotik:"address"`
	Bytes         int    `mikrotik:"bytes"`
	CallerID      string `mikrotik:"caller-id"`
	Encoding      string `mikrotik:"encoding"`
	LimitBytesIn  int    `mikrotik:"limit-bytes-in"`
	LimitBytesOut int    `mikrotik:"limit-bytes-out"`
	Name          string `mikrotik:"name"`
	Packets       int    `mikrotik:"packets"`
	Service       string `mikrotik:"service"`
	SessionID     string `mikrotik:"session-id"`
	Uptime        string `mikrotik:"uptime"`
}

// ====================================
//
// AAA
//
// ====================================

type PPPaaa struct {
	Accounting              bool   `mikrotik:"accounting"`
	InterimUpdate           string `mikrotik:"interim-update"`
	UseCircuitIDInNasPortID bool   `mikrotik:"use-circuit-id-in-nas-port-id"`
	UseRadius               bool   `mikrotik:"use-radius"`
}

const (
	PPPServiceAny   = "any"
	PPPServiceAsync = "async"
	PPPServiceL2TP  = "l2tp"
	PPPServiceOVPN  = "ovpn"
	PPPServicePPPoE = "pppoe"
	PPPServicePPTP  = "pptp"
	PPPServiceSSTP  = "sstp"
)
