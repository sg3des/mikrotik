package mikrotik

// ====================================
//
// Address
//
// ====================================

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

type IPDHCPOption struct {
	ID       string `mikrotik:".id"`
	Code     string `mikrotik:"code"`
	Default  string `mikrotik:"default"`
	Name     string `mikrotik:"name"`
	RawValue string `mikrotik:"raw-value"`
	Value    string `mikrotik:"value"`
}

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
// DHCP Server `mikrotik:""`
//
// ====================================

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

type IPDHCPServerConfig struct {
	StoreLeasesDisk string `mikrotik:"store-leases-disk"`
}

// Da completare
