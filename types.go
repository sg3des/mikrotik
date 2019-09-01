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

//IPAddress /ip/address
type IPAddress struct {
	ID string `mikrotik:".id"`

	Address string
	Network string

	Interface       string
	ActualInterface string

	Invalid  bool
	Dynamic  bool
	Disabled bool
}

func (ipa IPAddress) String() string {
	return fmt.Sprintf("%s dev %s", ipa.Address, ipa.Interface)
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

type Interface struct {
	ID          string `mikrotik:".id"`
	Name        string
	DefaultName string
	Type        string
	MTU         string `mikrotik:"mtu"`
	ActualMTU   string `mikrotik:"actual-mtu"`
	L2MTU       string `mikrotik:"l2mtu"`
	MaxL2MTU    string `mikrotik:"max-l2mtu"`
	MACAddress  string `mikrotik:"mac-address"`
	FastPath    bool
	LinkDowns   int
	RxByte      int
	TxByte      int
	RxPacket    int
	TxPacket    int
	RxDrop      int
	TxDrop      int
	RxError     int
	TxError     int
	FpRxByte    int
	FpTxByte    int
	FpRxPacket  int
	FpTxPacket  int

	Running  bool
	Slave    bool
	Disabled bool
	Comment  string
}

type WirelessInterface struct {
	ID                       string `mikrotik:".id"`
	DefaultName              string `mikrotik:"default-name"`
	Name                     string `mikrotik:"name"`
	Mtu                      int    `mikrotik:"mtu"`
	L2mtu                    int    `mikrotik:"l2mtu"`
	MacCddress               string `mikrotik:"mac-address"`
	Arp                      string `mikrotik:"arp"`
	ArpTimeout               string `mikrotik:"arp-timeout"`
	DisableRunningCheck      string `mikrotik:"disable-running-check"`
	InterfaceType            string `mikrotik:"interface-type"`
	RadioName                string `mikrotik:"radio-name"`
	Mode                     string `mikrotik:"mode"`
	SSID                     string `mikrotik:"ssid"`
	Area                     string `mikrotik:"area"`
	FrequencyMode            string `mikrotik:"frequency-mode"`
	Country                  string `mikrotik:"country"`
	AntennaGain              int    `mikrotik:"antenna-gain"`
	Frequency                int    `mikrotik:"frequency"`
	Band                     string `mikrotik:"band"`
	ChannelWidth             string `mikrotik:"channel-width"`
	ScanList                 string `mikrotik:"scan-list"`
	WirelessProtocol         string `mikrotik:"wireless-protocol"`
	RateSet                  string `mikrotik:"rate-set"`
	SupportedRatesB          string `mikrotik:"supported-rates-b"`
	SupportedRatesAG         string `mikrotik:"supported-rates-a/g"`
	BasicRatesB              string `mikrotik:"basic-rates-b"`
	BasicRatesAG             string `mikrotik:"basic-rates-a/g"`
	MaxStationCount          int    `mikrotik:"max-station-count"`
	Distance                 string `mikrotik:"distance"`
	TxPowerMode              string `mikrotik:"tx-power-mode"`
	NoiseFloorThreshold      string `mikrotik:"noise-floor-threshold"`
	Nv2NoiseFloorOffset      string `mikrotik:"nv2-noise-floor-offset"`
	VlanMode                 string `mikrotik:"vlan-mode"`
	VlanID                   string `mikrotik:"vlan-id"`
	WDSmode                  string `mikrotik:"wds-mode"`
	WDSdefaultBridge         string `mikrotik:"wds-default-bridge"`
	WDSdefaultCost           int    `mikrotik:"wds-default-cost"`
	WDScostRange             string `mikrotik:"wds-cost-range"`
	WDSignoreSSID            string `mikrotik:"wds-ignore-ssid"`
	UpdateStatsInterval      string `mikrotik:"update-stats-interval"`
	BridgeMode               string `mikrotik:"bridge-mode"`
	DefaultAuthentication    bool   `mikrotik:"default-authentication"`
	DefaultForwarding        bool   `mikrotik:"default-forwarding"`
	DefaultAPTxLimit         int    `mikrotik:"default-ap-tx-limit"`
	DefaultClientTxLimit     int    `mikrotik:"default-client-tx-limit"`
	WMMsupport               string `mikrotik:"wmm-support"`
	HideSSID                 bool   `mikrotik:"hide-ssid"`
	SecurityProfile          string `mikrotik:"security-profile"`
	InterworkingProfile      string `mikrotik:"interworking-profile"`
	WPSmode                  string `mikrotik:"wps-mode"`
	StationRoaming           string `mikrotik:"station-roaming"`
	DisconnectTimeout        string `mikrotik:"disconnect-timeout"`
	OnFailRetryTime          string `mikrotik:"on-fail-retry-time"`
	PreambleMode             string `mikrotik:"preamble-mode"`
	Compression              string `mikrotik:"compression"`
	AllowSharedkey           string `mikrotik:"allow-sharedkey"`
	StationBridgeCloneMac    string `mikrotik:"station-bridge-clone-mac"`
	AmpduPriorities          string `mikrotik:"ampdu-priorities"`
	GuardInterval            string `mikrotik:"guard-interval"`
	HtSupportedMCS           string `mikrotik:"ht-supported-mcs"`
	HtBasicMCS               string `mikrotik:"ht-basic-mcs"`
	TxChains                 string `mikrotik:"tx-chains"`
	RxChains                 string `mikrotik:"rx-chains"`
	AmsduLimit               int    `mikrotik:"amsdu-limit"`
	AmsduThreshold           int    `mikrotik:"amsdu-threshold"`
	TdmaPeriodSize           int    `mikrotik:"tdma-period-size"`
	Nv2QueueCount            int    `mikrotik:"nv2-queue-count"`
	Nv2QOS                   string `mikrotik:"nv2-qos"`
	Nv2CellRadius            int    `mikrotik:"nv2-cell-radius"`
	Nv2Security              string `mikrotik:"nv2-security"`
	Nv2PresharedKey          string `mikrotik:"nv2-preshared-key"`
	HwRetries                string `mikrotik:"hw-retries"`
	FrameLifetime            string `mikrotik:"frame-lifetime"`
	AdaptiveNoiseImmunity    string `mikrotik:"adaptive-noise-immunity"`
	HwFragmentationThreshold string `mikrotik:"hw-fragmentation-threshold"`
	HwProtectionMode         string `mikrotik:"hw-protection-mode"`
	HwProtectionThreshold    string `mikrotik:"hw-protection-threshold"`
	FrequencyOffset          string `mikrotik:"frequency-offset"`
	RateSelection            string `mikrotik:"rate-selection"`
	MulticastHelper          string `mikrotik:"multicast-helper"`
	MulticastBuffering       string `mikrotik:"multicast-buffering"`
	KeepaliveFrames          string `mikrotik:"keepalive-frames"`
	Running                  bool   `mikrotik:"running"`
	Disabled                 bool   `mikrotik:"disabled"`
}

type WirelessAP struct {
	Address         string
	SSID            string `mikrotik:"ssid"`
	Channel         string
	SIG             int    `mikrotik:"sig"`
	NF              int    `mikrotik:"nf"`
	SNR             int    `mikrotik:"snr"`
	RadioName       string `mikrotik:"radio-name"`
	RouterOSVersion string `mikrotik:"routeros-version"`
	Section         int
}

type WirelessSecurityProfile struct {
	ID      string `mikrotik:".id"`
	Name    string
	Mode    string
	Default bool

	AuthenticationTypes string `mikrotik:"authentication-types"`
	UnicastCiphers      string `mikrotik:"unicast-ciphers"`
	GroupCiphers        string `mikrotik:"group-ciphers"`

	WpaPreSharedKey    string `mikrotik:"wpa-pre-shared-key"`
	Wpa2PreSharedKey   string `mikrotik:"wpa2-pre-shared-key"`
	SupplicantIdentity string `mikrotik:"supplicant-identity"`
	EAPmethods         string `mikrotik:"eap-methods"`

	TLSmode        string `mikrotik:"tls-mode"`
	TLScertificate string `mikrotik:"tls-certificate"`

	Mschapv2username string `mikrotik:"mschapv2-username"`
	Mschapv2password string `mikrotik:"mschapv2-password"`

	RadiusMACauthentication string `mikrotik:"radius-mac-authentication"`
	RadiusMACaccounting     string `mikrotik:"radius-mac-accounting"`
	RadiusEAPaccounting     string `mikrotik:"radius-eap-accounting"`
	RadiusMACformat         string `mikrotik:"radius-mac-format"`
	RadiusMACmode           string `mikrotik:"radius-mac-mode"`
	RadiusMACcaching        string `mikrotik:"radius-mac-caching"`
	InterimUpdate           string `mikrotik:"interim-update"`
	GroupKeyUpdate          string `mikrotik:"group-key-update"`

	ManagementProtection    string `mikrotik:"management-protection"`
	ManagementProtectionKey string `mikrotik:"management-protection-key"`
}

const (
	WPA_PSK  = "wpa-psk"
	WPA2_PSK = "wpa2-psk"
	WPA_EAP  = "wpa-eap"
	WPA2_EAP = "wpa2-eap"
)

const (
	WirelessSecurityMode_None               = "none"
	WirelessSecurityMode_DynamicKeys        = "dynamic-keys"
	WirelessSecurityMode_StaticKeysRequired = "static-keys-required"
	WirelessSecurityMode_StaticKeysOptional = "static-keys-optional"
)

type SSTPserver struct {
	ID   string `mikrotik:".id"`
	Name string
	User string

	Running  bool
	Disabled bool
}

type SSTPclient struct {
	ID                                 string `mikrotik:".id"`
	Name                               string
	ConnectTo                          string `mikrotik:"connect-to"`
	Certificate                        string
	VerifyServerCertificate            string `mikrotik:"verify-server-certificate"`
	VerifyServerAddressFromCertificate string `mikrotik:"verify-server-address-from-certificate"`
	User                               string
	Password                           string
	Profile                            string
	AddDefaultRoute                    string `mikrotik:"add-default-route"`
	DefaultRouteDistance               int    `mikrotik:"default-route-distance"`
	DialOnDemand                       bool
	Running                            bool
	Disabled                           bool
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

// Routerboard - information from /system/routerboard/print
type Routerboard struct {
	Routerboard     bool
	BoardName       string `mikrotik:"board-name"`
	Model           string
	SerialNumber    string `mikrotik:"serial-number"`
	FirmwareType    string `mikrotik:"firmware-type"`
	FactoryFirmware string `mikrotik:"factory-firmware"`
	CurrentFirmware string `mikrotik:"current-firmware"`
	UpgradeFirmware string `mikrotik:"upgrade-firmware"`
}

type LteInfo struct {
	// this data was sent without pin
	PinStatus     string `mikrotik:"pin-status"`
	Functionality string
	Manufacturer  string
	Model         string
	Revision      string
	IMEI          string `mikrotik:"imei"`

	// this data was empty withut pin and was sent with correct pin set.
	RegistrationStatus string `mikrotik:"registration-status"`
	CurrentOperator    string `mikrotik:"current-operator"`
	Lac                string
	CurrentCellID      string `mikrotik:"current-cellid"`
	EnbID              string `mikrotik:"enb-id"`
	SectorID           string `mikrotik:"sector-id"`
	PhyCellID          string `mikrotik:"phy-cellid"`
	AccessTechnology   string `mikrotik:"access-technology"`
	SessionUptime      string `mikrotik:"session-uptime"`
	IMSI               string `mikrotik:"imsi"`
	UICC               string `mikrotik:"uicc"`
	SubscriberNumber   string `mikrotik:"subscriber-number"`
	Earfcn             string
	Rsrp               int
	Rsrq               string
	Sinr               int
}

type LtePrint struct {
	ID           string `mikrotik:".id"`
	Name         string
	Mtu          int
	MacAddress   string `mikrotik:"mac-address"`
	Pin          string
	ApnProfiles  string `mikrotik:"apn-profiles"`
	AllowRoaming bool   `mikrotik:"allow-roaming"`
	NetworkMode  string `mikrotik:"network-mode"`
	Running      bool
	Disalbed     bool
}
