package mikrotik

import (
	"fmt"
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

//Route /ip/route
type Route struct {
	ID string `mikrotik:".id"`

	DstAddress    string
	PrefSrc       string
	Gateway       string
	GatewayStatus string

	Distance    int
	Scope       int
	TargetScope int

	Active   bool
	Dynamic  bool
	Static   bool
	Disabled bool

	Comment string
}

func (r Route) String() string {
	return fmt.Sprintf("%s %s %s %d %s", r.DstAddress, r.PrefSrc, r.Gateway, r.Distance, r.Comment)
}

//NATRule /ip/firewall/nat
type NATRule struct {
	ID string `mikrotik:".id"`

	Chain  string
	Action string

	Protocol    string
	SrcAddress  net.IP
	DstAddress  net.IP
	ToAddresses string
	ToPorts     string

	InInterface  string
	OutInterface string

	SrcPort int
	DstPort int
	Port    int

	PacketMark     string
	ConnectionMark string
	RoutingMark    string
	RoutingTable   string

	ConnectionType string

	Log       bool
	LogPrefix string

	Bytes    int
	Packets  int
	Invalid  bool
	Dynamic  bool
	Disabled bool

	Comment string
}

const (
	ChainSrcNAT = "srcnat"
	ChainDstNAT = "dstnat"
)

const (
	FirewallActionAccept              = "accept"
	FirewallActionAddDstToAddressList = "add-dst-to-address-list"
	FirewallActionDstNAT              = "dst-nat"
	FirewallActionJump                = "jump"
	FirewallActionLog                 = "log"
	FirewallActionMasquerade          = "masquerade"
	FirewallActionNetmap              = "netmap"
	FirewallActionPassthrough         = "passthrough"
	FirewallActionRedirect            = "redirect"
	FirewallActionReturn              = "return"
	FirewallActionSame                = "same"
	FirewallActionSrcNAT              = "src-nat"
)

//MangleRule /ip/firewall/mangle
type MangleRule struct {
	ID string `mikrotik:".id"`

	Action string
	Chain  string

	NewRoutingMark    string
	NewConnectionMark string
	NewPacketMark     string

	ConnectionMark string
	PacketMark     string

	Passthrough     bool
	ConnectionState string
	DstAddressType  string

	InInterface string
	Nth         string

	Log       bool
	LogPrefix string

	Bytes   int
	Packets int

	Invalid  bool
	Dynamic  bool
	Disabled bool

	Comment string
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
