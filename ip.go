package mikrotik

import "fmt"

type ip struct {
	Address    cmd
	Route      cmd
	Firewall   firewallCMD
	Cloud      ipCloudCMD
	DNS        dnsCMD
	DHCPClient dhcpClientCMD
	DHCPServer dhcpServerCMD
	Service    cmd
	SSH        sshCMD
	Neighbor   neighborCMD
}

// ====================================
//
// Address
//
// ====================================

func (ipa IPAddress) String() string {
	return fmt.Sprintf("%s dev %s", ipa.Address, ipa.Interface)
}

// ====================================
//
// Cloud
//
// ====================================

type ipCloudCMD struct {
	cmd
	Advanced cfg
}

func (ipc *ipCloudCMD) ForceUpdate() error {
	_, err := ipc.cmd.mikrotik.RunArgs(ipc.cmd.path + "/force-update")
	return err
}

// ====================================
//
// DHCP-Client
//
// ====================================

type dhcpClientCMD struct {
	cmd
	Option cmd
}

func (dhcpc *dhcpClientCMD) Renew(id string) error {
	_, err := dhcpc.cmd.mikrotik.RunArgs(dhcpc.cmd.path+"/renew", "=numbers="+id)
	return err
}

func (dhcpc *dhcpClientCMD) Release(id string) error {
	_, err := dhcpc.cmd.mikrotik.RunArgs(dhcpc.cmd.path+"/release", "=numbers="+id)
	return err
}

// ====================================
//
// DHCP-Server
//
// ====================================

type dhcpServerCMD struct {
	cmd
	Alert   dhcpSAlert
	Lease   dhcpSLease
	Config  cfg
	Network cmd
	Option  dhcpSOption
}

type dhcpSAlert struct {
	cmd
}

func (dhcpsa *dhcpSAlert) ResetAlert(id string) error {
	_, err := dhcpsa.cmd.mikrotik.RunArgs(dhcpsa.cmd.path+"/reset-alert", "=numbers="+id)
	return err
}

type dhcpSLease struct {
	cmd
}

func (dhcpsl *dhcpSLease) MakeStatic(id string) error {
	_, err := dhcpsl.cmd.mikrotik.RunArgs(dhcpsl.cmd.path+"/make-static", "=numbers="+id)
	return err
}

func (dhcpsl *dhcpSLease) CheckStatus(id string) error {
	_, err := dhcpsl.cmd.mikrotik.RunArgs(dhcpsl.cmd.path+"/check-status", "=numbers="+id)
	return err
}

type dhcpSOption struct {
	cmd
	Sets cmd
}

// ====================================
//
// SSH
//
// ====================================

type sshCMD struct {
	cfg
}

func (ssh *sshCMD) ExportHostKey(prefix string) error {
	_, err := ssh.cfg.mikrotik.RunArgs(ssh.cfg.path+"/export-host-key", "=key-file-prefix="+prefix)
	return err
}

func (ssh *sshCMD) ImportHostKey(key string) error {
	_, err := ssh.cfg.mikrotik.RunArgs(ssh.cfg.path+"/import-host-key", "=private-key-file="+key)
	return err
}

func (ssh *sshCMD) RegenerateHostKey(prefix string) error {
	_, err := ssh.cfg.mikrotik.RunArgs(ssh.cfg.path + "/regenerate-host-key")
	return err
}

// ====================================
//
// DNS
//
// ====================================

type dnsCMD struct {
	cfg
	Static cmd
	Cache  dnsCacheCMD
}

type dnsCacheCMD struct {
	cfg
	All cfg
}

func (dnsc *dnsCacheCMD) Flush() error {
	_, err := dnsc.cfg.mikrotik.RunArgs(dnsc.cfg.path + "/flush")
	return err
}

// ====================================
//
// Firewall
//
// ====================================

type firewallCMD struct {
	NAT    cmd
	Mangle cmd
	Filter firewallFilterCMD
}

type firewallFilterCMD struct {
	cmd
}

func (ff *firewallFilterCMD) ResetCounters(id string) error {
	_, err := ff.cmd.mikrotik.RunArgs(ff.cmd.path+"/reset-counters", "=numbers="+id)
	return err
}

func (ff *firewallFilterCMD) ResetCountersAll() error {
	_, err := ff.cmd.mikrotik.RunArgs(ff.cmd.path + "/reset-counters-all")
	return err
}

// ====================================
//
// Neighbor
//
// ====================================

type neighborCMD struct {
	cmd
	DiscoverySettings cmd
}

// ====================================
//
// Pools
//
// ====================================

type poolCMD struct {
	cmd
	Used cfg
}

// ====================================
//
// Route
//
// ====================================
