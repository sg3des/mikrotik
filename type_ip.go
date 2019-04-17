package mikrotik

import (
	"fmt"
	"net"
)

// ====================================
//
// Address
//
// ====================================

// IPAddress is the struct for IP/Address*
type IPAddress struct {
	ID string `mikrotik:".id"`

	Address   string `mikrotik:"address"`
	Network   string `mikrotik:"network"`
	Broadcast string `mikrotik:"broadcast"`
	Netmask   string `mikrotik:"netmask"`

	Interface       string `mikrotik:"interface"`
	ActualInterface string `mikrotik:"actual-interface"`

	Invalid  bool   `mikrotik:"invalid"`
	Dynamic  bool   `mikrotik:"dynamic"`
	Disabled bool   `mikrotik:"disabled"`
	Comment  string `mikrotik:"comment"`
}

// ====================================
//
// Cloud
//
// ====================================

// IPCloud is the struct for IP/Cloud*
type IPCloud struct {
	DDNSEnabled        bool   `mikrotik:"ddns-enabled"`
	DDNSUpdateInterval string `mikrotik:"ddns-update-interval"`
	DNSName            string `mikrotik:"dns-name"`
	PublicAddress      string `mikrotik:"public-address"`
	PublicAddressV6    string `mikrotik:"public-address-ivp6"`
	Status             string `mikrotik:"status"`
	UpdateTime         bool   `mikrotik:"update-time"`
	UseLocalAddress    bool   `mikrotik:"use-local-address"` // /ip/cloud/advanced
	Warning            string `mikrotik:"warning"`
}

// ====================================
//
// SSH
//
// ====================================

// IPSSH is the struct for IP/SSH*
type IPSSH struct {
	AllowNoneCrypto          bool `mikrotik:"allow-none-crypto"`
	AlwaysAllowPasswordLogin bool `mikrotik:"always-allow-password-login"`
	ForwardingEnabled        bool `mikrotik:"forwarding-enabled"`
	HostKeySize              int  `mikrotik:"host-key-size"`
	StrongCrypto             bool `mikrotik:"strong-crypto"`
}

// ====================================
//
// Service
//
// ====================================

// IPService is the struct for IP/Service*
type IPService struct {
	ID          string `mikrotik:".id"`
	Address     string `mikrotik:"address"`
	Certificate string `mikrotik:"certificate"`
	Disabled    bool   `mikrotik:"disabled"`
	Invalid     bool   `mikrotik:"invalid"`
	Name        string `mikrotik:"name"`
	Port        int    `mikrotik:"port"`
}

// ====================================
//
// DHCP
//
// ====================================

// Works both with Client and Server

// IPDHCPOption is the struct for IP/DHCP-*/Option*
type IPDHCPOption struct {
	ID       string `mikrotik:".id"`
	Code     string `mikrotik:"code"`
	Default  string `mikrotik:"default"`
	Name     string `mikrotik:"name"`
	RawValue string `mikrotik:"raw-value"`
	Value    string `mikrotik:"value"`
}

// IPDHCPOptionSets is the struct for IP/DHCP-*/Option/Sets*
type IPDHCPOptionSets struct {
	ID      string `mikrotik:".id"`
	Name    string `mikrotik:"name"`
	Options string `mikrotik:"options"`
}

// ====================================
//
// DHCP Client
//
// ====================================

// IPDHCPClient is the struct for IP/DHCP-Client/*
type IPDHCPClient struct {
	ID string `mikrotik:".id"`

	AddDefaultRoute      string `mikrotik:"add-default-route"`
	Address              string `mikrotik:"address"` // read-only
	ClientID             string `mikrotik:"client-id"`
	Comment              string `mikrotik:"comment"`
	DefaultRouteDistance int    `mikrotik:"default-route-distance"`
	DHCPOptions          string `mikrotik:"dhcp-options"`
	DHCPServer           string `mikrotik:"dhcp-server"` // read-only
	Disabled             bool   `mikrotik:"disabled"`
	Dynamic              bool   `mikrotik:"dynamic"`
	ExpiresAfter         string `mikrotik:"expires-after"` // read-only
	Gateway              string `mikrotik:"gateway"`       // read-only
	HostName             string `mikrotik:"host-name"`
	Interface            string `mikrotik:"interface"`
	Invalid              bool   `mikrotik:"invalid"`       // read-only
	PrimaryDNS           string `mikrotik:"primary-dns"`   // read-only
	PrimaryNTP           string `mikrotik:"primary-ntp"`   // read-only
	SecondaryDNS         string `mikrotik:"secondary-dns"` // read-only
	SecondaryNTP         string `mikrotik:"secondary-ntp"` // read-only
	Status               string `mikrotik:"status"`        // read-only
	UsePeerDNS           bool   `mikrotik:"use-peer-dns"`
	UsePeerNTP           bool   `mikrotik:"use-peer-ntp"`
}

// ====================================
//
// DHCP Server
//
// ====================================

// IPDHCPServer is the struct for IP/DHCP-Server/*
type IPDHCPServer struct {
	ID                   string `mikrotik:".id"`
	AddArp               bool   `mikrotik:"add-arp"`
	AddressPool          string `mikrotik:"address-pool"`
	AllowDualStackQueue  bool   `mikrotik:"allow-dual-stack-queue"`
	AlwaysBroadcast      bool   `mikrotik:"always-broadcast"`
	Authoritative        string `mikrotik:"authoritative"`
	BootpLeaseTime       string `mikrotik:"bootp-lease-time"`
	BootpSupport         string `mikrotik:"bootp-support"`
	ConflictDetection    bool   `mikrotik:"conflict-detection"`
	DelayThreshold       string `mikrotik:"delay-threshold"`
	DHCPOptionSet        string `mikrotik:"dhcp-option-set"`
	Disabled             bool   `mikrotik:"disabled"`
	Dynamic              bool   `mikrotik:"dynamic"`
	InsertQueueBefore    string `mikrotik:"insert-queue-before"`
	Interface            string `mikrotik:"interface"`
	Invalid              bool   `mikrotik:"invalid"`
	LeaseScript          string `mikrotik:"lease-script"`
	LeaseTime            string `mikrotik:"lease-time"`
	Name                 string `mikrotik:"name"`
	ParentQueue          string `mikrotik:"parent-queue"`
	Relay                string `mikrotik:"relay"`
	SCRAddress           string `mikrotik:"src-address"`
	UseFramedAsClassless bool   `mikrotik:"use-framed-as-classles"`
	UseRadius            string `mikrotik:"use-radius"`
}

// IPDHCPServerLease is the struct for IP/DHCP-Server/Lease*
type IPDHCPServerLease struct {
	ID               string `mikrotik:".id"`
	ActiveAddress    string `mikrotik:"active-address"`
	ActiveClientId   string `mikrotik:"active-client-id"`
	ActiveMacAddress string `mikrotik:"active-mac-address"`
	ActiveServer     string `mikrotik:"active-server"`
	Address          string `mikrotik:"address"`
	AddressLists     string `mikrotik:"address-lists"`
	AlwaysBroadcast  bool   `mikrotik:"always-broadcast"`
	AgentCircuitId   string `mikrotik:"agent-circuit-id"`
	AgentRemoteId    string `mikrotik:"agent-remote-id"`
	BlockAccess      bool   `mikrotik:"block-access"`
	Blocked          string `mikrotik:"blocked"`
	ClientID         string `mikrotik:"client-id"`
	Comment          string `mikrotik:"comment"`
	DHCPOption       string `mikrotik:"dhcp-option"`
	Disabled         string `mikrotik:"disabled"`
	Dynamic          string `mikrotik:"dynamic"`
	ExpiresAfter     string `mikrotik:"expires-after"`
	HostName         string `mikrotik:"host-name"`
	LastSeen         string `mikrotik:"last-seen"`
	LeaseTime        string `mikrotik:"lease-time"`
	MacAddress       string `mikrotik:"mac-address"`
	Radius           string `mikrotik:"radius"`
	RateLimit        string `mikrotik:"rate-limit"`
	Server           string `mikrotik:"server"`
	SrcMacAddress    string `mikrotik:"src-mac-address"`
	Status           string `mikrotik:"status"`
	UseSrcMac        string `mikrotik:"use-src-mac"`
}

// IPDHCPServerNetwork is the struct for IP/DHCP-Server/Network*
type IPDHCPServerNetwork struct {
	ID            string `mikrotik:".id"`
	BootFileName  string `mikrotik:"boot-file-name"`
	Address       string `mikrotik:"address"`
	CapsManager   string `mikrotik:"caps-manager"`
	Comment       string `mikrotik:"comment"`
	DhcpOption    string `mikrotik:"dhcp-option"`
	DhcpOptionSet string `mikrotik:"dhcp-option-set"`
	DNSNone       bool   `mikrotik:"dns-none"`
	DNSServer     string `mikrotik:"dns-server"`
	Dynamic       bool   `mikrotik:"dynamic"`
	Domain        string `mikrotik:"domain"`
	Gateway       string `mikrotik:"gateway"`
	Netmask       int    `mikrotik:"netmask"`
	NextServer    string `mikrotik:"next-server"`
	NtpServer     string `mikrotik:"ntp-server"`
	WinsServer    string `mikrotik:"wins-server"`
}

// IPDHCPServerConfig is the struct for IP/DHCP-Server/Config*
type IPDHCPServerConfig struct {
	StoreLeasesDisk string `mikrotik:"store-leases-disk"`
}

// IPDHCPServerAlert is the struct for IP/DHCP-Server/Alert*
type IPDHCPServerAlert struct {
	ID            string `mikrotik:".id"`
	AlertTimeout  string `mikrotik:"alert-timeout"`
	Disabled      string `mikrotik:"disabled"`
	Interface     string `mikrotik:"interface"`
	Invalid       string `mikrotik:"invalid"`
	OnAlert       string `mikrotik:"on-alert"`
	UnknownServer string `mikrotik:"unknown-server"`
	ValidServer   string `mikrotik:"valid-server"`
}

// ====================================
//
// DNS
//
// ====================================

// IPDNS is the struct for IP/DNS/*
type IPDNS struct {
	AllowRemoteRequests      bool   `mikrotik:"allow-remote-requests"`
	CacheMaxTTL              string `mikrotik:"cache-max-ttl"`
	CacheSize                int    `mikrotik:"cache-size"`
	CacheUsed                int    `mikrotik:"cache-used"`
	DynamicServers           string `mikrotik:"dynamic-servers"`
	MaxConcurrentQueries     int    `mikrotik:"max-concurrent-queries"`
	MaxConcurrentTCPSessions int    `mikrotik:"max-concurrent-tcp-sessions"`
	MaxUDPPacketSize         int    `mikrotik:"max-udp-packet-size"`
	QueryServerTimeout       string `mikrotik:"query-server-timeout"`
	QueryTotalTimeout        string `mikrotik:"query-total-timeout"`
	Servers                  string `mikrotik:"servers"`
}

// IPDNSStatic is the struct for IP/DNS/Static*
type IPDNSStatic struct {
	ID       string `mikrotik:".id"`
	Address  string `mikrotik:"address"`
	Disabled bool   `mikrotik:"disabled"`
	Dynamic  bool   `mikrotik:"dynamic"`
	Name     string `mikrotik:"name"`
	Regexp   string `mikrotik:"regexp"`
	TTL      string `mikrotik:"ttl"`
}

// IPDNSCache is the struct for IP/DNS/Cache*
type IPDNSCache struct {
	ID      string `mikrotik:".id"`
	Address string `mikrotik:"address"`
	Name    string `mikrotik:"name"`
	Static  bool   `mikrotik:"static"`
	TTL     string `mikrotik:"ttl"`
}

// IPDNSCacheAll is the struct for IP/DNS/Cache/All*
type IPDNSCacheAll struct {
	ID     string `mikrotik:".id"`
	Data   string `mikrotik:"data"`
	Name   string `mikrotik:"name"`
	Static string `mikrotik:"static"`
	TTL    string `mikrotik:"ttl"`
	Type   string `mikrotik:"type"`
}

// ====================================
//
// Firewall
//
// ====================================

type IPFirewallNATRule struct {
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

type IPFirewallMangleRule struct {
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

type IPFirewallFilter struct {
	ID                      string `mikrotik:".id"`
	Action                  string `mikrotik:"action"`
	AddressList             string `mikrotik:"address-list"`
	AddressListTimeout      string `mikrotik:"address-list-timeout"`
	Bytes                   string `mikrotik:"bytes"`
	Chain                   string `mikrotik:"chain"`
	Comment                 string `mikrotik:"comment"`
	ConnectionBytes         int    `mikrotik:"connection-bytes"`
	ConnectionLimit         int    `mikrotik:"connection-limit"`
	ConnectionMark          string `mikrotik:"connection-mark"`
	ConnectionNatState      string `mikrotik:"connection-nat-state"`
	ConnectionRate          int    `mikrotik:"connection-rate"`
	ConnectionState         string `mikrotik:"connection-state"`
	ConnectionType          string `mikrotik:"connection-type"`
	Content                 string `mikrotik:"content"`
	Dscp                    int    `mikrotik:"dscp"`
	DstAddress              string `mikrotik:"dst-address"`
	DstAddressList          string `mikrotik:"dst-address-list"`
	DstAddressType          string `mikrotik:"dst-address-type"`
	DstLimit                string `mikrotik:"dst-limit"`
	DstPort                 int    `mikrotik:"dst-port"`
	Fragment                bool   `mikrotik:"fragment"`
	Hotspot                 string `mikrotik:"hotspot"`
	IcmpOptions             int    `mikrotik:"icmp-options"`
	InBridgePort            string `mikrotik:"in-bridge-port"`
	InBridgePortList        string `mikrotik:"in-bridge-port-list"`
	InInterface             string `mikrotik:"in-interface"`
	InInterfaceList         string `mikrotik:"in-interface-list"`
	IngressPriority         int    `mikrotik:"ingress-priority"`
	IpsecPolicy             string `mikrotik:"ipsec-policy"`
	Ipv4Options             string `mikrotik:"ipv4-options"`
	JumpTarget              string `mikrotik:"jump-target"`
	Layer7Protocol          string `mikrotik:"layer7-protocol"`
	Limit                   int    `mikrotik:"limit"`
	LogPrefix               string `mikrotik:"log-prefix"`
	Nth                     int    `mikrotik:"nth"`
	OutBridgePort           string `mikrotik:"out-bridge-port"`
	OutBridgePortList       string `mikrotik:"out-bridge-port-list"`
	OutInterface            string `mikrotik:"out-interface"`
	OutInterfaceList        string `mikrotik:"out-interface-list"`
	PacketMark              string `mikrotik:"packet-mark"`
	PacketSize              int    `mikrotik:"packet-size"`
	PerConnectionClassifier string `mikrotik:"per-connection-classifier"`
	Port                    int    `mikrotik:"port"`
	Priority                int    `mikrotik:"priority"`
	Protocol                string `mikrotik:"protocol"`
	Psd                     int    `mikrotik:"psd"`
	Random                  int    `mikrotik:"random"`
	RejectWith              string `mikrotik:"reject-with"`
	RoutingTable            string `mikrotik:"routing-table"`
	RoutingMark             string `mikrotik:"routing-mark"`
	SrcAddress              string `mikrotik:"src-address"`
	SrcAddressList          string `mikrotik:"src-address-list"`
	SrcAddressType          string `mikrotik:"src-address-type"`
	SrcPort                 int    `mikrotik:"src-port"`
	SrcMacAddress           string `mikrotik:"src-mac-address"`
	TCPFlags                string `mikrotik:"tcp-flags"`
	TCPMss                  int    `mikrotik:"tcp-mss"`
	Time                    string `mikrotik:"time"`
	TLSHost                 string `mikrotik:"tls-host"`
	TTL                     int    `mikrotik:"ttl"`
	Dynamic                 string `mikrotik:"dynamic"`
	Invalid                 string `mikrotik:"invalid"`
	Packets                 string `mikrotik:"packets"`
}

// ====================================
//
// Neighbor
//
// ====================================

type IPNeighbor struct {
	ID                string `mikrotik:".id"`
	Address           string `mikrotik:"address"`
	Address6          string `mikrotik:"address6"`
	Age               string `mikrotik:"age"`
	Board             string `mikrotik:"board"`
	Identity          string `mikrotik:"identity"`
	Interface         string `mikrotik:"interface"`
	InterfaceName     string `mikrotik:"interface-name"`
	Ipv6              bool   `mikrotik:"ipv6"`
	MacAddress        string `mikrotik:"mac-address"`
	Platform          string `mikrotik:"platform"`
	SoftwareID        string `mikrotik:"software-id"`
	SystemCaps        string `mikrotik:"system-caps"`
	SystemCapsEnabled string `mikrotik:"system-caps-enabled"`
	Unpack            string `mikrotik:"unpack"`
	Uptime            string `mikrotik:"uptime"`
	Version           string `mikrotik:"version"`
}

type IPNeighbordDiscoverySettings struct {
	DiscoverInterfaceList string `mikrotik:"discover-interface-list"`
}

// ====================================
//
// Pools
//
// ====================================

type IPPool struct {
	ID       string `mikrotik:".id"`
	Name     string `mikrotik:"name"`
	NextPool string `mikrotik:"next-pool"`
	Ranges   string `mikrotik:"ranges"`
}

type IPPoolUsed struct {
	ID      string `mikrotik:".id"`
	Address string `mikrotik:"address"`
	Info    string `mikrotik:"info"`
	Owner   string `mikrotik:"owner"`
	Pool    string `mikrotik:"pool"`
}

// ====================================
//
// Route
//
// ====================================

type IPRoute struct {
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

func (r IPRoute) String() string {
	return fmt.Sprintf("%s %s %s %d %s", r.DstAddress, r.PrefSrc, r.Gateway, r.Distance, r.Comment)
}

// ====================================
//
// Arp
//
// ====================================

type IPArp struct {
	ID         string `mikrotik:".id"`
	DHCP       bool   `mikrotik:"DHCP"`
	Address    string `mikrotik:"address"`
	Complete   bool   `mikrotik:"complete"`
	Disabled   bool   `mikrotik:"disabled"`
	Dynamic    bool   `mikrotik:"dynamic"`
	Interface  string `mikrotik:"interface"`
	Invalid    bool   `mikrotik:"invalid"`
	MacAddress string `mikrotik:"mac-address"`
	Published  bool   `mikrotik:"published"`
}

