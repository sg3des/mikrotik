package mikrotik

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

var mikrotik *Mikrotik

func init() {
	log.SetFlags(log.Lshortfile)

	addr := "192.168.107.78:8728"
	user := "admin"
	pass := ""

	var err error
	mikrotik, err = DialTimeout(addr, user, pass, 3e9)
	if err != nil {
		log.Fatal(err)
	}
}

func TestRun(t *testing.T) {
	_, err := mikrotik.Run("/system/identity/print")
	if err != nil {
		t.Error(err)
	}

	// t.Log(re)
}

func TestIPAddress(t *testing.T) {
	ip := IPAddress{
		Address:   "10.0.0.3/24",
		Interface: "bridge1",
	}

	var list []*IPAddress
	if err := mikrotik.IP.Address.List(&list); err != nil {
		t.Error(err)
	}

	// for _, ip := range list {
	// 	t.Log(ip)
	// }

	if err := mikrotik.IP.Address.Add(&ip); err != nil {
		t.Error(err)
	}

	// t.Log(ip)

	if err := mikrotik.IP.Address.Remove(ip.ID); err != nil {
		t.Error(err)
	}
}

func TestSystem(t *testing.T) {
	name, err := mikrotik.System.Identity.Name()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Logf("/system/idenitity name: %s", name)

	if err := mikrotik.System.Identity.SetName(name); err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestInterface(t *testing.T) {
	var list []*Interface
	if err := mikrotik.Interface.List(&list); err != nil {
		t.Error(err)
	}

	// for _, intf := range list {
	// 	t.Log(intf)
	// }
}

func TestWirelessInterface(t *testing.T) {
	var list []*WirelessInterface
	err := mikrotik.Interface.Wireless.List(&list)
	if err != nil {
		t.Error(err)
	}

	for _, intf := range list {
		t.Log(intf)
	}
}

func TestWirelessScan(t *testing.T) {
	list, err := mikrotik.Interface.Wireless.Scan("wlan1", "5s")
	if err != nil {
		t.Error(err)
	}

	for _, ap := range list {
		t.Log(ap)
	}
}

func TestWirelessSecurityProfiles(t *testing.T) {
	var list []*WirelessSecurityProfile
	err := mikrotik.Interface.Wireless.SecurityProfiles.List(&list)
	if err != nil {
		t.Error(err)
	}

	for _, item := range list {
		t.Logf("%+v", item)
	}

	var profile *WirelessSecurityProfile
	err = mikrotik.Interface.Wireless.SecurityProfiles.Find("?default=true", &profile)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	err = mikrotik.Interface.Wireless.SecurityProfiles.Set(profile.ID, WirelessSecurityProfile{Wpa2PreSharedKey: "12345678"})
	if err != nil {
		t.Error(err)
	}
}

func TestFind(t *testing.T) {
	var list []*Interface
	if err := mikrotik.Interface.Find("type=ether", &list); err != nil {
		t.Error(err)
	}

	if len(list) == 0 {
		t.Error("failed find items")
	}

	// for _, intf := range list {
	// 	t.Log(intf)
	// }

	var intf *Interface
	if err := mikrotik.Interface.Find("name=ether1", &intf); err != nil {
		t.Error(err)
	}
	t.Log(intf)

	var ipaddrs []*IPAddress
	err := mikrotik.IP.Address.Find(`interface=ppp-out51`, &ipaddrs)
	if err != nil {
		t.Error(err)
	}
	if len(ipaddrs) == 0 {
		t.Error("ipaddresses not found")
	}
	for _, ipa := range ipaddrs {
		t.Log(ipa)
	}
}

// func TestGet(t *testing.T) {
// 	var intf *Interface
// 	if err := mikrotik.Interface.GetByName("ether1", &intf); err != nil {
// 		t.Error(err)
// 	}

// 	t.Log(intf)
// }

func TestRoute(t *testing.T) {
	var list []*Route
	if err := mikrotik.IP.Route.List(&list); err != nil {
		t.Error(err)
	}

	for _, item := range list {
		t.Log(item)
	}

	if len(list) == 0 {
		t.Error("routes not found")
	}

	route := list[0]
	err := mikrotik.IP.Route.Set(route.ID, &Route{Comment: "newcomment"})
	if err != nil {
		t.Error(err)
	}

	// err = mikrotik.IP.Route.Set(route.ID, Route{Distance: route.Distance})
	// if err != nil {
	// 	t.Error(err)
	// }
}

func TestFirewallNAT(t *testing.T) {
	// mikrotik.Debug(true)

	var rules []NATRule
	if err := mikrotik.IP.Firewall.NAT.List(&rules); err != nil {
		t.Error(err)
	}

	// for _, rule := range rules {
	// 	t.Logf("%+v", rule)
	// }

	rule := NATRule{
		Chain:       ChainDstNAT,
		Protocol:    "tcp",
		DstPort:     2222,
		Action:      FirewallActionNetmap,
		ToAddresses: "192.168.241.1",
		ToPorts:     "22",
	}
	if err := mikrotik.IP.Firewall.NAT.Add(&rule); err != nil {
		t.Error(err)
	}
}

func TestSSTPServer(t *testing.T) {
	s := SSTPserver{
		Name: "test-sstp-server",
		User: "test-sstp-user",
	}
	err := mikrotik.Interface.SSTPServer.Add(&s)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", s)

	var s2 SSTPserver
	if err = mikrotik.Interface.SSTPServer.Find(".id="+s.ID, &s2); err != nil {
		t.Error(err)
	}

	t.Logf("%+v", s2)

	if err = mikrotik.Interface.SSTPServer.Remove(s.ID); err != nil {
		t.Error(err)
	}
}

func TestPPPprofile(t *testing.T) {
	var profiles []*PPPprofile
	if err := mikrotik.PPP.Profile.List(&profiles); err != nil {
		t.Error(err)
	}

	for _, p := range profiles {
		t.Logf("%+v", p)
	}
}

func TestPPPsecret(t *testing.T) {
	secret := Secret{
		Name:     "test-secret",
		Password: "test-password",
		Service:  PPPServiceAny,
		Profile:  "default",
	}
	if err := mikrotik.PPP.Secret.Add(&secret); err != nil {
		t.Error(err)
	}

	secret.Profile = "default-encryption"
	if err := mikrotik.PPP.Secret.Set(secret.ID, &secret); err != nil {
		t.Error(err)
	}

	var secrets []*Secret
	if err := mikrotik.PPP.Secret.List(&secrets); err != nil {
		t.Error(err)
	}

	for _, p := range secrets {
		t.Logf("%+v", p)
	}

	if err := mikrotik.PPP.Secret.Remove(secret.ID); err != nil {
		t.Error(err)
	}
}

func TestSSTPClient(t *testing.T) {
	c := SSTPclient{
		Name:      "test",
		ConnectTo: "10.0.0.1:443",
		User:      "username",
		Password:  "password",
	}

	if err := mikrotik.Interface.SSTPClient.Add(&c); err != nil {
		t.Error(err)
		// t.FailNow()
	}

	if err := mikrotik.Interface.SSTPClient.Find("name=test", &c); err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(c.ID)

	if err := mikrotik.Interface.SSTPClient.Remove(c.ID); err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestPing(t *testing.T) {
	resp, err := mikrotik.Ping("8.8.8.8", 1).Start()
	if err != nil {
		t.Error(err)
	}

	for _, r := range resp {
		t.Logf("%+v", r)
	}

	ping := mikrotik.Ping("8.8.8.8", 1)
	ping.SrcAddress = "192.168.240.1"
	resp, err = ping.Start()
	if err != nil {
		t.Error(err)
	}

	for _, r := range resp {
		t.Logf("%+v", r)
	}
}
