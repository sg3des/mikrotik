package mikrotik

// ====================================
//
// File
//
// ====================================

type SystemFile struct {
	ID                  string `mikrotik:".id"`
	Contents            string `mikrotik:"contents"`
	CreationTime        string `mikrotik:"creation-time"`
	PackageArchitecture string `mikrotik:"package-architecture"`
	PackageBuiltTime    string `mikrotik:"package-built-time"`
	PackageName         string `mikrotik:"package-name"`
	PackageVersion      string `mikrotik:"package-version"`
	Name                string `mikrotik:"name"`
	Size                int    `mikrotik:"size"`
	Type                string `mikrotik:"type"`
}

// ====================================
//
// SNTP Client
//
// ====================================

type SystemNTPClient struct {
	Enabled        string
	ServerDNSNames string `mikrotik:"server-dns-names"`
	Mode           string
}

// ====================================
//
// LEDs
//
// ===================================

type SystemLEDS struct {
	ID                  string `mikrotik :".id"`
	Default             bool   `mikrotik :"default"`
	Disabled            bool   `mikrotik :"disabled"`
	Interface           string `mikrotik :"interface"`
	ModemSignalTreshold int    `mikrotik :"modem-signal-treshold"`
	Leds                string `mikrotik :"leds"`
	Type                string `mikrotik :"type"`
}

const (
	SystemLedsApCap                  = "ap-cap"
	SystemLedsFlashAccess            = "flash-access"
	SystemLedsInterfaceActivity      = "interface-activity"
	SystemLedsInterfaceReceive       = "interface-receive"
	SystemLedsInterfaceSpeed         = "interface-speed"
	SystemLedsInterfaceStatus        = "interface-status"
	SystemLedsInterfaceTransmit      = "interface-transmit"
	SystemLedsModemSignal            = "modem-signal"
	SystemLedsOff                    = "off"
	SystemLedsOn                     = "on"
	SystemLedsPOEOut                 = "poe-out"
	SystemLedsWirelessSignalStrenght = "wireless-signal-strength"
	SystemLedsWirelessStruts         = "wireless-status"
)

type SystemLEDSSettings struct {
	AllLedsOff string `mikrotik :"all-leds-off"`
}

const (
	LedsOFFAfter1h   = "after-1h"
	LedsOFFAfter1min = "after-1min"
	LedsOFFImmediate = "immediate"
	LedsOFFNever     = "never"
)

// ====================================
//
// User
//
// ====================================

type SystemUserSSHKeys struct {
	ID       string `mikrotik :".id"`
	DSA      bool   `mikrotik :"DSA"`
	RSA      bool   `mikrotik :"RSA"`
	Bits     int    `mikrotik :"bits"`
	KeyOwner string `mikrotik :"key-owner"`
	User     string `mikrotik :"user"`
}

// ====================================
//
// Logging
//
// ====================================

type SystemLogging struct {
	ID       string `mikrotik :".id"`
	Action   string `mikrotik :"action"`
	Default  bool   `mikrotik :"default"`
	Disabled bool   `mikrotik :"disabled"`
	Invalid  bool   `mikrotik :"invalid"`
	Prefix   string `mikrotik :"prefix"`
	Topics   string `mikrotik :"topics"`
}

type SystemLoggingAction struct {
	ID               string `mikrotik:".id"`
	BsdSyslog        bool   `mikrotik:"bsd-syslog"`
	Default          string `mikrotik:"default"`
	DiskFileCount    int    `mikrotik:"disk-file-count"`
	DiskFileName     string `mikrotik:"disk-file-name"`
	DiskLinesPerFile int    `mikrotik:"disk-lines-per-file"`
	DiskStopOnFull   bool   `mikrotik:"disk-stop-on-full"`
	EmailStartTLS    bool   `mikrotik:"email-start-tls"`
	EmailTo          string `mikrotik:"email-to"`
	MemoryLines      int    `mikrotik:"memory-lines"`
	MemoryStopOnFull bool   `mikrotik:"memory-stop-on-full"`
	Name             string `mikrotik:"name"`
	Remember         bool   `mikrotik:"remember"`
	Remote           string `mikrotik:"remote"`
	RemotePort       string `mikrotik:"remote-port"`
	SrcAddress       string `mikrotik:"src-address"`
	SyslogFacility   string `mikrotik:"syslog-facility"`
	SyslogSeverity   string `mikrotik:"syslog-severity"`
	SyslogTimeFormat string `mikrotik:"syslog-time-format"`
	Target           string `mikrotik:"target"`
}

// ====================================
//
// Resource
//
// ====================================

type SystemResourceProperties struct {
	ArchitectureName     string `mikrotik:"architecture-name"`
	BadBlocks            string `mikrotik:"bad-blocks"`
	BoardName            string `mikrotik:"board-name"`
	BuildTime            string `mikrotik:"build-time"`
	CPU                  string `mikrotik:"cpu"`
	CPUCount             string `mikrotik:"cpu-count"`
	CPUFrequency         string `mikrotik:"cpu-frequency"`
	CPULoad              string `mikrotik:"cpu-load"`
	FactorySoftware      string `mikrotik:"factory-software"`
	FreeHddSpace         string `mikrotik:"free-hdd-space"`
	FreeMemory           string `mikrotik:"free-memory"`
	Platform             string `mikrotik:"platform"`
	TotalHddSpace        string `mikrotik:"total-hdd-space"`
	TotalMemory          string `mikrotik:"total-memory"`
	Uptime               string `mikrotik:"uptime"`
	Version              string `mikrotik:"version"`
	WriteSectSinceReboot string `mikrotik:"write-sect-since-reboot"`
	WriteSectTotal       string `mikrotik:"write-sect-total"`
}

type SystemResourceCPU struct {
	ID   string `mikrotik:".id"`
	CPU  string `mikrotik:"cpu"`
	Disk string `mikrotik:"disk"`
	Irq  string `mikrotik:"irq"`
	Load string `mikrotik:"load"`
}

type SystemResourceUSB struct {
	ID           string `mikrotik:".id"`
	Device       string `mikrotik:"device"`
	DeviceID     string `mikrotik:"device-id"`
	Name         string `mikrotik:"name"`
	Ports        int    `mikrotik:"ports"`
	SerialNumber string `mikrotik:"serial-number"`
	Speed        string `mikrotik:"speed"`
	UsbVersion   string `mikrotik:"usb-version"`
	Vendor       string `mikrotik:"vendor"`
	VendorID     string `mikrotik:"vendor-id"`
}

type SystemResourceIRQ struct {
	ID          string `mikrotik:".id"`
	ActiveCPU   string `mikrotik:"active-cpu"`
	Count       string `mikrotik:"count"`
	CPU         string `mikrotik:"cpu"`
	Irq         string `mikrotik:"irq"`
	PerCPUCount string `mikrotik:"per-cpu-count"`
	ReadOnly    string `mikrotik:"read-only"`
	Users       string `mikrotik:"users"`
}

type SystemResourcePCI struct {
	ID       string `mikrotik:".id"`
	Category string `mikrotik:"category"`
	Device   string `mikrotik:"device"`
	DeviceID string `mikrotik:"device-id"`
	Io       string `mikrotik:"io"`
	Irq      int    `mikrotik:"irq"`
	Memory   string `mikrotik:"memory"`
	Name     string `mikrotik:"name"`
	Vendor   string `mikrotik:"vendor"`
	VendorID string `mikrotik:"vendor-id"`
}

// ====================================
//
// Routerboard
//
// ====================================

type SystemRouterboard struct {
	BoardName       string `mikrotik:"board-name"`
	CurrentFirmware string `mikrotik:"current-firmware"`
	FactoryFirmware string `mikrotik:"factory-firmware"`
	FirmwareType    string `mikrotik:"firmware-type"`
	Model           string `mikrotik:"model"`
	Routerboard     string `mikrotik:"routerboard"`
	SerialNumber    string `mikrotik:"serial-number"`
	UpgradeFirmware string `mikrotik:"upgrade-firmware"`
}

type SystemRouterboardSettings struct {
	AutoUpgrade           bool   `mikrotik:"auto-upgrade"`
	BaudRate              int    `mikrotik:"baud-rate"`
	BootDelay             string `mikrotik:"boot-delay"`
	BootDevice            string `mikrotik:"boot-device"`
	BootProtocol          string `mikrotik:"boot-protocol"`
	CPUFrequency          string `mikrotik:"cpu-frequency"`
	CPUMode               string `mikrotik:"cpu-mode"`
	EnableJumperReset     bool   `mikrotik:"enable-jumper-reset"`
	EnterSetupOn          string `mikrotik:"enter-setup-on"`
	ForceBackupBooter     bool   `mikrotik:"force-backup-booter"`
	MemoryFrequency       string `mikrotik:"memory-frequency"`
	MemoryDataRate        string `mikrotik:"memory-data-rate"`
	InitDelay             string `mikrotik:"init-delay"`
	ProtectedRouterboot   string `mikrotik:"protected-routerboot"`
	ReformatHoldButton    string `mikrotik:"reformat-hold-button"`
	ReformatHoldButtonMax string `mikrotik:"reformat-hold-button-max"`
	RegulatoryDomainCe    bool   `mikrotik:"regulatory-domain-ce"`
	SilentBoot            bool   `mikrotik:"silent-boot"`
}

type SystemRouterboardModeButton struct {
	Enabled bool   `mikrotik:"enabled"`
	OnEvent string `mikrotik:"on-event"`
}

// ====================================
//
// Package
//
// ====================================

type SystemCheckForUpdates struct {
	Section          string `mikrotik:".section"`
	Channel          string `mikrotik:"channel"`
	InstalledVersion string `mikrotik:"installed-version"`
	LatestVersion    string `mikrotik:"latest-version"`
	Status           string `mikrotik:"status"`
}

type SystemPackage struct {
	ID        string `mikrotik:".id"`
	BuildTime string `mikrotik:"build-time"`
	Bundle    string `mikrotik:"bundle"`
	Disabled  bool   `mikrotik:"disabled"`
	Name      string `mikrotik:"name"`
	Scheduled string `mikrotik:"scheduled"`
	Version   string `mikrotik:"version"`
}

// ====================================
//
// Script
//
// ====================================

type SystemScript struct {
	ID                     string   `mikrotik:".id"`
	Comment                string   `mikrotik:"comment"`
	DontRequirePermissions bool     `mikrotik:"dont-require-permissions"`
	Invalid                bool     `mikrotik:"invalid"`
	LastStarted            string   `mikrotik:"last-started"`
	Name                   string   `mikrotik:"name"`
	Owner                  string   `mikrotik:"owner"`
	Policy                 []string `mikrotik:"policy"`
	RunCount               int      `mikrotik:"run-count"`
	Source                 string   `mikrotik:"source"`
}

type SystemScriptJob struct {
	ID      string   `mikrotik:".id"`
	Owner   string   `mikrotik:"owner"`
	Parent  string   `mikrotik:"parent"`
	Policy  []string `mikrotik:"policy"`
	Started string   `mikrotik:"started"`
	Type    string   `mikrotik:"type"`
}

type SystemScriptEnvironment struct {
	Owner   string   `mikrotik:"owner"`
	Policy  []string `mikrotik:"policy"`
	Started string   `mikrotik:"started"`
}
