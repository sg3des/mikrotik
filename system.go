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

type Leds struct {
	cmd
	Settings cfg
}

// ====================================
//
// User
//
// ====================================

type User struct {
	cmd
	SSHKeys sshkeys
}

type sshkeys struct {
	SSHCmds
	Private SSHCmds
}
