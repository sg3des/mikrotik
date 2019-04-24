package mikrotik

type system struct {
	Identity    identity
	NTP         ntp
	File        file
	Led         leds
	User        user
	Resource    resource
	Routerboard sysrouterboard
	Package     syspackage
	Scheduler   cmd
}

// ====================================
//
// Identity
//
// ====================================

type identity struct {
	mikrotik *Mikrotik
	path     string
}

func (c *identity) Name() (string, error) {
	var resp struct {
		Name string
	}
	err := c.mikrotik.Print(c.path+"/print", &resp)
	return resp.Name, err
}

func (c *identity) SetName(name string) error {
	return c.mikrotik.SetOne(c.path+"/set", "name", name)
}

// ====================================
//
// File
//
// ====================================

type file struct {
	mikrotik *Mikrotik
	path     string
}

func (f *file) Print(v interface{}) error {
	return f.mikrotik.Print(f.path+"/print", v)
}

func (f *file) Set(name, value string) error {
	return f.mikrotik.SetOne(f.path+"/set", name, value)
}

func (f *file) SetContents(id, contents string) error {
	return f.mikrotik.Set(f.path+"/set", id, "=contents="+contents)
}

func (f *file) New(name string) (files []SystemFile, err error) {
	re, err := f.mikrotik.RunArgs(f.path+"/print", "=file="+name)
	if err != nil {
		return nil, err
	}

	err = f.mikrotik.ParseResponce(re, &files)

	return files, err
}

// ====================================
//
// SNTP Client
//
// ====================================

type ntp struct {
	Client cfg
}

// ====================================
//
// LEDs
//
// ====================================

type leds struct {
	cmd
	Settings cfg
}

// ====================================
//
// User
//
// ====================================

type user struct {
	cmd
	SSHKeys sshkeys
	AAA     cfg
	Active  cfg
	Group   cmd
}

type sshkeys struct {
	sshcmds
	Private sshcmds
}

// ====================================
//
// Logging
//
// ====================================

type logging struct {
	cmd
	Action cmd
}

// ====================================
//
// Resource
//
// ====================================

type resource struct {
	USB cfg
	CPU cfg
	IRQ cfg
	PCI respci
}

type respci struct {
	cfg
}

// USBPowerReset to implement

// ====================================
//
// Routerboard
//
// ====================================

type sysrouterboard struct {
	mikrotik   *Mikrotik
	path       string
	Settings   cfg
	ModeButton cfg
	USB        usbcfg
}

// Print simply calls the mikrotik's Print for the commands that use cmd struct.
func (router *sysrouterboard) Print(v interface{}) error {
	return router.mikrotik.Print(router.path+"/print", v)
}

// usbcfg is a struct for System/Routerboard/USB that has cfg's and PowerReset methods
type usbcfg struct {
	cfg
}

// PowerReset wants duration and bus in order to stop the power to the USB (bus) for an amount of time (duration). Default: Duration= 3s and Bus = 1.
func (usb *usbcfg) PowerReset(duration, bus string) error {
	if duration == "" {
		duration = "3s"
	} else if bus == "" {
		bus = "1"
	}

	_, err := usb.cfg.mikrotik.RunArgs(usb.path+"/power-reset", "=duration="+duration, "=bus="+bus)

	return err
}

// ====================================
//
// Package
//
// ====================================

type syspackage struct {
	cmd
	Update sysupdate
}

func (sys *syspackage) Uninstall(name string) error {
	_, err := sys.cmd.mikrotik.RunArgs(sys.cmd.path+"/uninstall", "=numbers="+name)

	return err
}

func (sys *syspackage) Unschedule(name string) error {
	_, err := sys.cmd.mikrotik.RunArgs(sys.cmd.path+"/unschedule", "=numbers="+name)

	return err
}

func (sys *syspackage) Downgrade() error {
	_, err := sys.cmd.mikrotik.RunArgs(sys.cmd.path + "/downgrade")

	return err
}

type sysupdate struct {
	cfg
}

func (sys *sysupdate) Install() error {
	_, err := sys.cfg.mikrotik.RunArgs(sys.cfg.path + "/install")

	return err
}

func (sys *sysupdate) Download() error {
	_, err := sys.cfg.mikrotik.RunArgs(sys.cfg.path + "/download")

	return err
}

func (sys *sysupdate) Cancel() error {
	_, err := sys.cfg.mikrotik.RunArgs(sys.cfg.path + "/cancel")

	return err
}

func (sys *sysupdate) CheckForUpdates() (updates []SystemCheckForUpdates, err error) {

	re, err := sys.cfg.mikrotik.RunArgs(sys.cfg.path + "/check-for-updates")
	if err != nil {
		return nil, err
	}

	err = sys.cfg.mikrotik.ParseResponce(re, &updates)

	return updates, err
}

// ====================================
//
// Script
//
// ====================================

type script struct {
	cmd
	Job         cmd
	Environment cmd
}

func (scr *script) Run(name string) error {
	_, err := scr.cmd.mikrotik.RunArgs(scr.cmd.path+"/run", "=number="+name)

	return err
}
