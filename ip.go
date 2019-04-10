package mikrotik

import "fmt"

type ip struct {
	Address    cmd
	Route      cmd
	Firewall   firewall
	Cloud      IPCloudCMD
	DHCPClient DHCPClientCMD
	DHCPServer DHCPServerCMD
	Service    cmd
	SSH        SSHCMD
}

type firewall struct {
	NAT    cmd
	Mangle cmd
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

type IPCloudCMD struct {
	cmd
	Advanced cfg
}

func (ipc *IPCloudCMD) ForceUpdate() error {
	_, err := ipc.cmd.mikrotik.RunArgs(ipc.cmd.path + "/force-update")
	return err
}

// ====================================
//
// DHCP-Client
//
// ====================================

type DHCPClientCMD struct {
	cmd
	Option cmd
}

func (dhcpc *DHCPClientCMD) Renew(id string) error {
	_, err := dhcpc.cmd.mikrotik.RunArgs(dhcpc.cmd.path+"/renew", "=numbers="+id)
	return err
}

func (dhcpc *DHCPClientCMD) Release(id string) error {
	_, err := dhcpc.cmd.mikrotik.RunArgs(dhcpc.cmd.path+"/release", "=numbers="+id)
	return err
}

// ====================================
//
// DHCP-Server
//
// ====================================

type DHCPServerCMD struct {
	cmd
	Alert   DHCPSAlert
	Lease   DHCPSLease
	Config  cfg
	Network cmd
	Option  DHCPSOption
}

type DHCPSAlert struct {
	cmd
}

func (dhcpsa *DHCPSAlert) ResetAlert(id string) error {
	_, err := dhcpsa.cmd.mikrotik.RunArgs(dhcpsa.cmd.path+"/reset-alert", "=numbers="+id)
	return err
}

type DHCPSLease struct {
	cmd
}

func (dhcpsl *DHCPSLease) MakeStatic(id string) error {
	_, err := dhcpsl.cmd.mikrotik.RunArgs(dhcpsl.cmd.path+"/make-static", "=numbers="+id)
	return err
}

func (dhcpsl *DHCPSLease) CheckStatus(id string) error {
	_, err := dhcpsl.cmd.mikrotik.RunArgs(dhcpsl.cmd.path+"/check-status", "=numbers="+id)
	return err
}

type DHCPSOption struct {
	cmd
	Sets cmd
}

// ====================================
//
// SSH
//
// ====================================

type SSHCMD struct {
	cfg
}

func (ssh *SSHCMD) ExportHostKey(prefix string) error {
	_, err := ssh.cfg.mikrotik.RunArgs(ssh.cfg.path+"/export-host-key", "=key-file-prefix="+prefix)
	return err
}

func (ssh *SSHCMD) ImportHostKey(key string) error {
	_, err := ssh.cfg.mikrotik.RunArgs(ssh.cfg.path+"/import-host-key", "=private-key-file="+key)
	return err
}

func (ssh *SSHCMD) RegenerateHostKey(prefix string) error {
	_, err := ssh.cfg.mikrotik.RunArgs(ssh.cfg.path + "/regenerate-host-key")
	return err
}

// ====================================
//
// DNS
//
// ====================================
