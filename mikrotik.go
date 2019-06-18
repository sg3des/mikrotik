package mikrotik

import (
	"errors"
	"sync"

	routeros "gopkg.in/routeros.v2"
)

//Mikrotik is common struct contains connection to device and API-tree
type Mikrotik struct {
	Conn      *routeros.Client
	connMutex sync.Mutex

	IP        ip
	System    system
	Interface netinterface
	PPP       ppp

	debug bool
}

// ====================================
//
// Mikrotik Utilities
//
// ====================================

// setMikrotikCommands sets the relative paths of the commands.
func (mik *Mikrotik) setMikrotikCommands() {
	mik.IP = ip{
		Address: cmd{mikrotik: mik, path: "/ip/address"},
		Route:   cmd{mikrotik: mik, path: "/ip/route"},
		Firewall: firewallCMD{
			NAT:    cmd{mikrotik: mik, path: "/ip/firewall/nat"},
			Mangle: cmd{mikrotik: mik, path: "/ip/firewall/mangle"},
			Filter: firewallFilterCMD{
				cmd: cmd{mikrotik: mik, path: "/ip/firewall/filter"},
			},
		},
		Cloud: ipCloudCMD{
			cmd: cmd{
				mikrotik: mik,
				path:     "/ip/cloud",
			},
			Advanced: cfg{mikrotik: mik, path: "/ip/cloud/advanced"},
		},
		DHCPClient: dhcpClientCMD{
			cmd: cmd{
				mikrotik: mik,
				path:     "/ip/dhcp-client",
			},
			Option: cmd{mikrotik: mik, path: "/ip/dhcp-client/option"},
		},
		DHCPServer: dhcpServerCMD{
			cmd: cmd{
				mikrotik: mik,
				path:     "/ip/dhcp-server",
			},
			Alert: dhcpSAlert{
				cmd: cmd{
					mikrotik: mik,
					path:     "/ip/dhcp-server/alert",
				},
			},
			Lease: dhcpSLease{
				cmd: cmd{
					mikrotik: mik,
					path:     "/ip/dhcp-server/lease",
				},
			},
		},
		DNS: dnsCMD{
			cfg:    cfg{mikrotik: mik, path: "/ip/dns"},
			Static: cmd{mikrotik: mik, path: "/ip/dns/static"},
			Cache: dnsCacheCMD{
				cfg: cfg{mikrotik: mik, path: "/ip/dns/cache"},
				All: cfg{mikrotik: mik, path: "/ip/dns/cache/all"},
			},
		},
		Service: cmd{mikrotik: mik, path: "/ip/service"},
		SSH: sshCMD{
			cfg{mikrotik: mik, path: "/ip/ssh"},
		},
		Neighbor: neighborCMD{
			cmd:               cmd{mikrotik: mik, path: "/ip/neighbor"},
			DiscoverySettings: cmd{mikrotik: mik, path: "/ip/neighbor/discovery-settings"},
		},
	}

	mik.System = system{
		Identity: identity{mikrotik: mik, path: "/system/identity"},
		NTP: ntp{
			Client: cfg{mikrotik: mik, path: "/system/ntp/client"},
		},
		File: file{mikrotik: mik, path: "/file"},
	}

	mik.Interface = netinterface{
		mikrotik:   mik,
		path:       "/interface",
		SSTPClient: cmd{mikrotik: mik, path: "/interface/sstp-server"},
		SSTPServer: netsstpserver{
			cmd: cmd{
				mikrotik: mik,
				path:     "/interface/sstp-client",
			},
			Server: cfg{mikrotik: mik, path: "/interface/sstp-server/server"},
		},
		PPPOEClient: cmd{mikrotik: mik, path: "/interface/pppoe-client"},
		Ethernet:    cmd{mikrotik: mik, path: "/interface/ethernet"},
		List: netlist{
			cmd: cmd{
				mikrotik: mik,
				path:     "/interface/list",
			},
			Member: cmd{mikrotik: mik, path: "/interface/list/member"},
		},
		Bridge: netbridge{
			cmd: cmd{
				mikrotik: mik,
				path:     "/interface/bridge",
			},
			Settings: cfg{mikrotik: mik, path: "/interface/bridge/settings"},
			Port:     cmd{mikrotik: mik, path: "/interface/bridge/port"},
		},
		VLAN: cmd{mikrotik: mik, path: "/interface/vlan"},
		Wireless: wireless{
			cmd: cmd{
				mikrotik: mik,
				path:     "/interface/wireless",
			},
			SecurityProfiles: cmd{mikrotik: mik, path: "/interface/wireless/security-profiles"},
		},
	}

	mik.PPP = ppp{
		AAA:        cfg{mikrotik: mik, path: "/ppp/aaa"},
		Secret:     cmd{mikrotik: mik, path: "/ppp/secret"},
		L2tpSecret: cmd{mikrotik: mik, path: "/ppp/l2tp-secret"},
		Profile:    cmd{mikrotik: mik, path: "/ppp/profile"},
	}
}

func (mik *Mikrotik) Ping(addr string, count int) *Ping {
	return &Ping{mikrotik: mik, Address: addr, Count: count}
}

func (ping *Ping) Start() ([]*PingResponse, error) {
	re, err := ping.mikrotik.RunArgs("/ping", ToArgs(ping)...)
	if err != nil {
		return nil, err
	}

	var pingResp []*PingResponse
	for _, resp := range re.Re {
		ValuesFrom(resp.Map).To(&pingResp)
	}

	if len(pingResp) > 0 && pingResp[0].PacketLoss == 100 {
		return pingResp, errors.New("timeout")
	}

	return pingResp, nil
}
