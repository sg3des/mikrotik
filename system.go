package mikrotik

type system struct {
	Identity identity
	NTP      ntp
	File     file
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
// Resource
//
// ====================================

type sysresource struct {
	mikrotik   *Mikrotik
	path       string
	Settings   cfg
	ModeButton cfg
	USB        usbcfg
}

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
