package mikrotik

type system struct {
	Identity identity
	NTP      ntp
}
type ntp struct {
	Client cfg
}

type cfg struct {
	mikrotik *Mikrotik
	path     string
}

type identity struct {
	mikrotik *Mikrotik
	path     string
}

func (c *cfg) Get(v interface{}) error {
	return c.mikrotik.Print(c.path+"/print", v)
}

func (c *cfg) Set(name, value string) error {
	return c.mikrotik.SetOne(c.path+"/set", name, value)
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
