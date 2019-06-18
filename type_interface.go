package mikrotik

type Interface struct {
	ID               string `mikrotik:".id"`
	Name             string `mikrotik:"name"`
	DefaultName      string `mikrotik:"default-name"`
	Type             string
	MTU              string `mikrotik:"mtu"`
	ActualMTU        string `mikrotik:"actual-mtu"`
	L2MTU            string `mikrotik:"l2mtu"`
	MaxL2MTU         string `mikrotik:"max-l2mtu"`
	MACAddress       string `mikrotik:"mac-address"`
	LastLinkDownTime string `mikrotik:"last-link-down-time"`
	LastLinkUpTime   string `mikrotik:"last-link-up-time"`
	FastPath         bool
	LinkDowns        int
	RxByte           int
	TxByte           int
	RxPacket         int
	TxPacket         int
	RxDrop           int
	TxDrop           int
	RxError          int
	TxError          int
	FpRxByte         int
	FpTxByte         int
	FpRxPacket       int
	FpTxPacket       int

	Running  bool
	Slave    bool
	Disabled bool
	Comment  string
}

// ====================================
//
// Bridge
//
// ====================================

// Bridge is the struct for Interface/Bridge/*
type Bridge struct {
	ID                string `mikrotik:".id"`
	Name              string `mikrotik:"name"`
	ActualMTU         string `mikrotik:"actual-mtu"`
	L2MTU             int    `mikrotik:"l2mtu"`
	MTU               int    `mikrotik:"mtu"`
	AdminMac          string `mikrotik:"admin-mac"`
	AgeingTime        string `mikrotik:"ageing-time"`
	Arp               string `mikrotik:"arp"`
	ArpTimeout        string `mikrotik:"arp-timeout"`
	AutoMac           string `mikrotik:"auto-mac"`
	FastForward       string `mikrotik:"fast-forward"`
	ForwardDelay      string `mikrotik:"forward-delay"`
	IgmpSnooping      string `mikrotik:"igmp-snooping"`
	MaxMessageAge     string `mikrotik:"max-message-age"`
	Priority          string `mikrotik:"priority"`
	ProtocolMode      string `mikrotik:"protocol-mode"`
	TransmitHoldCount string `mikrotik:"transmit-hold-count"`
	VlanFiltering     string `mikrotik:"vlan-filtering"`
}

// BridgeSettings is the struct for Interface/Bridge/Settings/*
type BridgeSettings struct {
	AllowFastPath            bool `mikrotik:"allow-fast-path"`
	BridgeFastForwardBytes   int  `mikrotik:"bridge-fast-forward-bytes"`
	BridgeFastForwardPackets int  `mikrotik:"bridge-fast-forward-packets"`
	BridgeFastPathActive     bool `mikrotik:"bridge-fast-path-active"`
	BridgeFastPathBytes      int  `mikrotik:"bridge-fast-path-bytes"`
	BridgeFastPathPackets    int  `mikrotik:"bridge-fast-path-packets"`
	UseIPFirewall            bool `mikrotik:"use-ip-firewall"`
	UseIPFirewallForPppoe    bool `mikrotik:"use-ip-firewall-for-pppoe"`
	UseIPFirewallForVlan     bool `mikrotik:"use-ip-firewall-for-vlan"`
}

// BridgePort is the struct for Interface/Bridge/Port/*
type BridgePort struct {
	ID                    string `mikrotik:".id"`
	Nextid                string `mikrotik:".nextid"`
	AutoIsolate           bool   `mikrotik:"auto-isolate"`
	Bridge                string `mikrotik:"bridge"`
	BroadcastFlood        bool   `mikrotik:"broadcast-flood"`
	Comment               string `mikrotik:"comment"`
	DebugInfo             string `mikrotik:"debug-info"`
	Disabled              bool   `mikrotik:"disabled"`
	Dynamic               bool   `mikrotik:"dynamic"`
	Edge                  string `mikrotik:"edge"`
	EdgePort              bool   `mikrotik:"edge-port"`
	EdgePortDiscovery     bool   `mikrotik:"edge-port-discovery"`
	ExternalFdbStatus     bool   `mikrotik:"external-fdb-status"`
	Forwarding            bool   `mikrotik:"forwarding"`
	FrameTypes            string `mikrotik:"frame-types"`
	Horizon               string `mikrotik:"horizon"`
	Hw                    bool   `mikrotik:"hw"`
	HwOffload             bool   `mikrotik:"hw-offload"`
	HwOffloadGroup        string `mikrotik:"hw-offload-group"`
	Inactive              bool   `mikrotik:"inactive"`
	IngressFiltering      bool   `mikrotik:"ingress-filtering"`
	Interface             string `mikrotik:"interface"`
	InternalPathCost      int    `mikrotik:"internal-path-cost"`
	Learn                 string `mikrotik:"learn"`
	Learning              bool   `mikrotik:"learning"`
	PathCost              int    `mikrotik:"path-cost"`
	PointToPoint          string `mikrotik:"point-to-point"`
	PointToPointPort      bool   `mikrotik:"point-to-point-port"`
	PortNumber            int    `mikrotik:"port-number"`
	Priority              string `mikrotik:"priority"`
	Pvid                  int    `mikrotik:"pvid"`
	RestrictedRole        bool   `mikrotik:"restricted-role"`
	RestrictedTcn         bool   `mikrotik:"restricted-tcn"`
	Role                  string `mikrotik:"role"`
	SendingRstp           bool   `mikrotik:"sending-rstp"`
	Status                string `mikrotik:"status"`
	UnknownMulticastFlood bool   `mikrotik:"unknown-multicast-flood"`
	UnknownUnicastFlood   bool   `mikrotik:"unknown-unicast-flood"`
}

// ====================================
//
// Ethernet
//
// ====================================

// Ethernet is the struct for Interface/Ethernet/Print
type Ethernet struct {
	ID                      string `mikrotik:".id"`
	Advertise               string `mikrotik:"advertise"`
	Arp                     string `mikrotik:"arp"`
	ArpTimeout              string `mikrotik:"arp-timeout"`
	AutoNegotiation         bool   `mikrotik:"auto-negotiation"`
	Bandwidth               string `mikrotik:"bandwidth"`
	DefaultName             string `mikrotik:"default-name"`
	DriverRxByte            int    `mikrotik:"driver-rx-byte"`
	DriverRxPacket          int    `mikrotik:"driver-rx-packet"`
	DriverTxByte            int    `mikrotik:"driver-tx-byte"`
	DriverTxPacket          int    `mikrotik:"driver-tx-packet"`
	FullDuplex              bool   `mikrotik:"full-duplex"`
	L2MTU                   int    `mikrotik:"l2mtu"`
	LoopProtect             string `mikrotik:"loop-protect"`
	LoopProtectDisableTime  string `mikrotik:"loop-protect-disable-time"`
	LoopProtectSendInterval string `mikrotik:"loop-protect-send-interval"`
	LoopProtectStatus       string `mikrotik:"loop-protect-status"`
	MACAddress              string `mikrotik:"mac-address"`
	MTU                     int    `mikrotik:"mtu"`
	Name                    string `mikrotik:"name"`
	OrigMacAddress          string `mikrotik:"orig-mac-address"`
	Rx10241518              int    `mikrotik:"rx-1024-1518"`
	Rx128255                int    `mikrotik:"rx-128-255"`
	Rx1519Max               int    `mikrotik:"rx-1519-max"`
	Rx256511                int    `mikrotik:"rx-256-511"`
	Rx5121023               int    `mikrotik:"rx-512-1023"`
	Rx64                    int    `mikrotik:"rx-64"`
	Rx65127                 int    `mikrotik:"rx-65-127"`
	RxAlignError            int    `mikrotik:"rx-align-error"`
	RxBroadcast             int    `mikrotik:"rx-broadcast"`
	RxBytes                 int    `mikrotik:"rx-bytes"`
	RxFcsError              int    `mikrotik:"rx-fcs-error"`
	RxFlowControl           string `mikrotik:"rx-flow-control"`
	RxFragment              int    `mikrotik:"rx-fragment"`
	RxMulticast             int    `mikrotik:"rx-multicast"`
	RxOverflow              int    `mikrotik:"rx-overflow"`
	RxPause                 int    `mikrotik:"rx-pause"`
	RxTooLong               int    `mikrotik:"rx-too-long"`
	RxTooShort              int    `mikrotik:"rx-too-short"`
	Speed                   string `mikrotik:"speed"`
	Switch                  string `mikrotik:"switch"`
	Tx10241518              int    `mikrotik:"tx-1024-1518"`
	Tx128255                int    `mikrotik:"tx-128-255"`
	Tx1519Max               int    `mikrotik:"tx-1519-max"`
	Tx256511                int    `mikrotik:"tx-256-511"`
	Tx5121023               int    `mikrotik:"tx-512-1023"`
	Tx64                    int    `mikrotik:"tx-64"`
	Tx65127                 int    `mikrotik:"tx-65-127"`
	TxBroadcast             int    `mikrotik:"tx-broadcast"`
	TxBytes                 int    `mikrotik:"tx-bytes"`
	TxCollision             int    `mikrotik:"tx-collision"`
	TxDeferred              int    `mikrotik:"tx-deferred"`
	TxExcessiveCollision    int    `mikrotik:"tx-excessive-collision"`
	TxExcessiveDeferred     int    `mikrotik:"tx-excessive-deferred"`
	TxFlowControl           string `mikrotik:"tx-flow-control"`
	TxLateCollision         int    `mikrotik:"tx-late-collision"`
	TxMulticast             int    `mikrotik:"tx-multicast"`
	TxMultipleCollision     int    `mikrotik:"tx-multiple-collision"`
	TxPause                 int    `mikrotik:"tx-pause"`
	TxSingleCollision       int    `mikrotik:"tx-single-collision"`
	TxTooLong               int    `mikrotik:"tx-too-long"`
	TxUnderrun              int    `mikrotik:"tx-underrun"`

	Running  bool
	Slave    bool
	Disabled bool
}

// ====================================
//
// SSTP
//
// ====================================

// SSTPClient is the struct for Interface/sstp-client/Print
type SSTPClient struct {
	ID string `mikrotik:".id"`

	AddDefaultRoute                    bool   `mikrotik:"add-default-route"`
	Authentication                     string `mikrotik:"authentication"`
	Certificate                        string `mikrotik:"certificate"`
	Comment                            string `mikrotik:"comment"`
	ConnectTo                          string `mikrotik:"connect-to"`
	DefaultRouteDistance               int    `mikrotik:"default-route-distance"`
	DialOnDemand                       bool   `mikrotik:"dial-on-demand"`
	Disabled                           bool   `mikrotik:"disabled"`
	HTTPProxy                          string `mikrotik:"http-proxy"`
	KeepAliveTimeout                   int    `mikrotik:"keepalive-timeout"`
	MaxMRU                             int    `mikrotik:"max-mru"`
	MaxMTU                             int    `mikrotik:"max-mtu"`
	MRRU                               string `mikrotik:"mrru"`
	Name                               string `mikrotik:"name"`
	Password                           string `mikrotik:"password"`
	PFS                                bool   `mikrotik:"pfs"`
	Profile                            string `mikrotik:"profile"`
	User                               string `mikrotik:"user"`
	TLSVersion                         string `mikrotik:"tls-version"`
	VerifyServerCertificate            bool   `mikrotik:"verify-server-certificate"`
	VerifyServerAddressFromCertificate bool   `mikrotik:"verify-server-address-from-certificate"`
}

// SSTPServer is the struct for Interface/sstp-server/server/*
type SSTPServer struct {
	ID   string `mikrotik:".id"`
	Name string `mikrotik:"name"`
	User string `mikrotik:"user"`

	Authentication          string `mikrotik:"authentication"`
	Certificate             string `mikrotik:"certificate"`
	DefaultProfile          string `mikrotik:"default-profile"`
	Enabled                 bool   `mikrotik:"enabled"`
	ForceAES                bool   `mikrotik:"force-aes"`
	KeepAliveTimeout        int    `mikrotik:"keepalive-timeout"`
	MaxMRU                  int    `mikrotik:"max-mru"`
	MaxMTU                  int    `mikrotik:"max-mtu"`
	MRRU                    string `mikrotik:"mrru"`
	PFS                     bool   `mikrotik:"pfs"`
	Port                    int    `mikrotik:"port"`
	TLSVersion              string `mikrotik:"tls-version"`
	VerifyClientCertificate bool   `mikrotik:"verify-client-certificate"`

	Running  bool
	Disabled bool
}

// ====================================
//
// List
//
// ====================================

// List is the struct for Interface/List/*
type List struct {
	Name    string `mikrotik:"name"`
	Include string `mikrotik:"include"`
	Exclude string `mikrotik:"exclude"`
}

// ListMember is the struct for Interface/List/Member*
type ListMember struct {
	Interface string `mikrotik:"interface"`
	List      string `mikrotik:"list"`
}

// ====================================
//
// PPPOE Client
//
// ====================================

// PPPOEClient is the struct for Interface/pppoe-client/*
type PPPOEClient struct {
	ACMAC                string `mikrotik:"ac-mac"`
	ACName               string `mikrotik:"ac-name"`
	ActiveLinks          int    `mikrotik:"active-links"`
	AddDefaultRoute      bool   `mikrotik:"add-default-route"`
	Allow                string `mikrotik:"allow"`
	DefaultRouteDistance int    `mikrotik:"default-route-distance"`
	DialOnDemand         bool   `mikrotik:"dial-on-demand"`
	Encoding             string `mikrotik:"encoding"`
	Interface            string `mikrotik:"interface"`
	LocalAddress         string `mikrotik:"local-address"`
	KeepAliveTimeout     int    `mikrotik:"keepalive-timeout"`
	MRU                  int    `mikrotik:"mru"`
	MaxMRU               int    `mikrotik:"max-mru"`
	MTU                  int    `mikrotik:"mtu"`
	MaxMTU               int    `mikrotik:"max-mtu"`
	MRRU                 string `mikrotik:"mrru"`
	Name                 string `mikrotik:"name"`
	Password             string `mikrotik:"password"`
	Profile              string `mikrotik:"profile"`
	RemoteAddress        string `mikrotik:"remote-address"`
	ServiceName          string `mikrotik:"service-name"`
	Status               string `mikrotik:"status"`
	UsePeerDNS           bool   `mikrotik:"use-peer-dns"`
	User                 string `mikrotik:"user"`
	Uptime               string `mikrotik:"uptime"`
}

// ====================================
//
// Wireless
//
// ====================================

// WirelessInterface is the struct for Interface/Wireless/*
type WirelessInterface struct {
	ID                       string `mikrotik:".id"`
	DefaultName              string `mikrotik:"default-name"`
	Name                     string `mikrotik:"name"`
	MTU                      int    `mikrotik:"mtu"`
	L2MTU                    int    `mikrotik:"l2mtu"`
	MACAddress               string `mikrotik:"mac-address"`
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

// WirelessAP is the struct for Interface/Wireless/AP/*
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

// WirelessSecurityProfile is the struct for Interface/Wireless/Security/Profile*
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

// ====================================
//
// VLANs
//
// ====================================

// Structure for /interface/vlan/print

type VLANPrint struct {
	ID                      string `json:".id"`
	Arp                     string `json:"arp"`
	ArpTimeout              string `json:"arp-timeout"`
	Disabled                bool   `json:"disabled"`
	Interface               string `json:"interface"`
	L2Mtu                   int    `json:"l2mtu"`
	LoopProtect             string `json:"loop-protect"`
	LoopProtectDisableTime  string `json:"loop-protect-disable-time"`
	LoopProtectSendInterval string `json:"loop-protect-send-interval"`
	LoopProtectStatus       string `json:"loop-protect-status"`
	MacAddress              string `json:"mac-address"`
	Mtu                     int    `json:"mtu"`
	Name                    string `json:"name"`
	Running                 bool   `json:"running"`
	UseServiceTag           bool   `json:"use-service-tag"`
	VlanID                  int    `json:"vlan-id"`
}
