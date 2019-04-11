package mikrotik

import (
	"net"
	"time"
)

// ====================================
//
// Entities
//
// ====================================

type Ping struct {
	mikrotik *Mikrotik
	Address  string

	Count        int
	Interface    string
	Interval     int
	RoutingTable string
	Size         int
	SrcAddress   string
	TTL          int `mikrotik:"ttl"`
}

type PingResponse struct {
	Host string

	Status string

	Sent       int
	Received   int
	PacketLoss int

	TTL  int `mikrotik:"ttl"`
	Seq  int
	Size int

	Time   time.Duration
	MinRTT time.Duration `mikrotik:"min-rtt"`
	AvgRTT time.Duration `mikrotik:"avg-rtt"`
	MaxRTT time.Duration `mikrotik:"max-rtt"`
}

type SystemNTPClient struct {
	Enabled        string
	ServerDNSNames string `mikrotik:"server-dns-names"`
	Mode           string
}

type PPPprofile struct {
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

type Secret struct {
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

const (
	PPPServiceAny   = "any"
	PPPServiceAsync = "async"
	PPPServiceL2TP  = "l2tp"
	PPPServiceOVPN  = "ovpn"
	PPPServicePPPoE = "pppoe"
	PPPServicePPTP  = "pptp"
	PPPServiceSSTP  = "sstp"
)
