package mikrotik

//netinterface not have add method
type netinterface struct {
	mikrotik *Mikrotik
	path     string

	SSTPClient  cmd
	SSTPServer  netsstpserver
	Ethernet    cmd
	List        netlist
	Bridge      netbridge
	PPPOEClient cmd
	VLAN        cmd

	Wireless wireless
}

type netsstpserver struct {
	cmd
	Server cfg
}

type netlist struct {
	cmd
	Member cmd
}

type netbridge struct {
	cmd
	Settings cfg
	Port     cmd
}

type wireless struct {
	// mikrotik *Mikrotik
	// path     string
	cmd

	SecurityProfiles cmd
}

func (c *netinterface) Print(v interface{}) error {
	return c.mikrotik.Print(c.path+"/print", v)
}

func (c *netinterface) Find(where string, v interface{}) error {
	re, err := c.mikrotik.RunArgs(c.path+"/print", "?"+where)
	if err != nil {
		return err
	}

	return c.mikrotik.ParseResponce(re, v)
}

func (c *netinterface) GetByName(name string, v interface{}) error {
	re, err := c.mikrotik.RunArgs(c.path+"/print", "?name="+name)
	if err != nil {
		return err
	}

	return c.mikrotik.ParseResponce(re, v)
}

func (c *netinterface) Set(id string, v interface{}) error {
	return c.mikrotik.Set(c.path+"/set", id, v)
}

func (c *netinterface) Remove(id string) error {
	return c.mikrotik.Remove(c.path+"/remove", id)
}

func (c *netinterface) Enable(id string) error {
	return c.mikrotik.Enable(c.path+"/enable", id)
}

func (c *netinterface) Disable(id string) error {
	return c.mikrotik.Disable(c.path+"/disable", id)
}

func (c *netinterface) Comment(id, comment string) error {
	return c.mikrotik.Comment(c.path+"/comment", id, comment)
}

func (c *wireless) Scan(name, duration string) (APlist []*WirelessAP, err error) {
	re, err := c.mikrotik.RunArgs(c.path+"/scan", "=.id="+name, "=duration="+duration)
	if err != nil {
		return nil, err
	}

	var list []*WirelessAP
	err = c.mikrotik.ParseResponce(re, &list)
	if err != nil {
		return list, err
	}

	//remove duplicates
	for i := range list {
		if c.contains(APlist, list[i]) {
			continue
		}
		APlist = append(APlist, list[i])
	}

	return
}

func (c *wireless) contains(list []*WirelessAP, ap *WirelessAP) bool {
	for i := range list {
		if list[i].Address == ap.Address && list[i].SSID == ap.SSID {
			return true
		}
	}
	return false
}
