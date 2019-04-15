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
	ID                  string `json:".id"`
	Default             bool   `json:"default"`
	Disabled            bool   `json:"disabled"`
	Interface           string `json:"interface"`
	ModemSignalTreshold int    `json:"modem-signal-treshold"`
	Leds                string `json:"leds"`
	Type                string `json:"type"`
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
	AllLedsOff string `json:"all-leds-off"`
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
	ID       string `json:".id"`
	DSA      bool   `json:"DSA"`
	RSA      bool   `json:"RSA"`
	Bits     int    `json:"bits"`
	KeyOwner string `json:"key-owner"`
	User     string `json:"user"`
}
