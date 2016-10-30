package ros

import (
	"strconv"
)

func firstList(l Lister, command string, filter map[string]string, properties []string, detail bool) (map[string]string, error) {
	res, err := l.List(command, filter, properties, detail)
	if err != nil {
		return nil, err
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func SystemResource(p Printer) (map[string]string, error) {
	return p.Print("/system/resource", nil,
		[]string{
			"version",
			"cpu",
			"architecture-name",
			"board-name",
			"platform",
		},
		false,
	)
}

func SystemIdentity(p Printer) (map[string]string, error) {
	return p.Print("/system/identity", nil,
		[]string{
			"name",
		},
		false,
	)
}

func SetSystemIdentityName(s Setter, name string) error {
	return s.Set("/system/identity", nil,
		map[string]string{
			"name": name,
		})
}

func SystemNote(p Printer) (map[string]string, error) {
	return p.Print("/system/note", nil,
		[]string{
			"note",
		},
		false,
	)
}

func SetSystemNote(s Setter, note string) error {
	return s.Set("/system/note", nil,
		map[string]string{
			"note": note,
		})
}

func SetSystemNoteShowAtLogin(s BoolSetter, show bool) error {
	return s.Set("/system/note", nil,
		map[string]string{
			"show-at-login": s.FormatBool(show),
		})
}

func SystemClock(p Printer) (map[string]string, error) {
	return p.Print("/system/clock", nil,
		[]string{
			"time-zone-name",
			"time-zone-autotodetect",
		},
		false,
	)
}

func SetSystemClock(s Setter, key, value string) error {
	return s.Set("/system/clock", nil,
		map[string]string{
			key: value,
		})
}

func SetSystemClockTimeZoneName(s Setter, name string) error {
	return SetSystemClock(s, "time-zone-name", name)
}
func SetSystemClockTimeZoneAutodetect(s BoolSetter, auto bool) error {
	return SetSystemClock(s, "time-zone-autodetect", s.FormatBool(auto))
}

func SystemNTPClient(p Printer) (map[string]string, error) {
	return p.Print("/system/ntp/client", nil,
		[]string{
			"enabled",
			"primary-ntp",
			"secondary-ntp",
		},
		false,
	)
}
func SetSystemNTPClient(s Setter, key, value string) error {
	return s.Set("/system/ntp/client", nil,
		map[string]string{
			key: value,
		})
}

func SetSystemNTPClientEnabled(s BoolSetter, enabled bool) error {
	return SetSystemNTPClient(s, "enabled", s.FormatBool(enabled))
}

func SetSystemNTPClientPrimaryNTP(s Setter, ntp string) error {
	return SetSystemNTPClient(s, "primary-ntp", ntp)
}

func SetSystemNTPClientSecondaryNTP(s Setter, ntp string) error {
	return SetSystemNTPClient(s, "secondary-ntp", ntp)
}

func SystemLoggingAction(l Lister, name string) (map[string]string, error) {
	return firstList(l, "/system/logging/action",
		map[string]string{
			"name": name,
		},
		[]string{
			"remote",
			"target",
			"remote-port",
			"bds-syslog",
			"src-address",
			"syslog-facility",
			"syslog-severity",
			"syslog-time-format",
		},
		true,
	)
}

func SetSystemLoggingAction(s Setter, name, key, value string) error {
	return s.Set("/system/logging/action",
		map[string]string{
			"name": name,
		},
		map[string]string{
			key: value,
		})
}

func SetSystemLoggingActionRemote(s Setter, name, remote string) error {
	return SetSystemLoggingAction(s, name, "remote", remote)
}
func SetSystemLoggingActionTarget(s Setter, name, target string) error {
	return SetSystemLoggingAction(s, name, "target", target)
}
func SetSystemLoggingActionRemotePort(s Setter, name string, port int) error {
	return SetSystemLoggingAction(s, name, "remote-port", strconv.Itoa(port))
}
func SetSystemLoggingActionBSDSyslog(s BoolSetter, name string, bsd bool) error {
	return SetSystemLoggingAction(s, name, "bsd-syslog", s.FormatBool(bsd))
}
func SetSystemLoggingActionSrcAddress(s Setter, name, address string) error {
	return SetSystemLoggingAction(s, name, "src-address", address)
}
func SetSystemLoggingActionSyslogFacility(s Setter, name, address string) error {
	return SetSystemLoggingAction(s, name, "syslog-facility", address)
}
func SetSystemLoggingActionSyslogSeverity(s Setter, name, severity string) error {
	return SetSystemLoggingAction(s, name, "syslog-severity", severity)
}
func SetSystemLoggingActionSyslogTimeFormat(s Setter, name, format string) error {
	return SetSystemLoggingAction(s, name, "syslog-time-format", format)
}

func SystemLogging(l Lister, action, topics string) (map[string]string, error) {
	return firstList(l, "/system/logging",
		map[string]string{
			"action": action,
			"topics": topics,
		},
		[]string{
			"prefix",
		},
		true,
	)
}

func AddSystemLogging(a Adder, action, topics string) error {
	return a.Add("/system/logging",
		map[string]string{
			"action": action,
			"topics": topics,
		},
	)
}

func RemoveSystemLogging(r Remover, action, topics string) error {
	return r.Remove("/system/logging",
		map[string]string{
			"action": action,
			"topics": topics,
		},
	)
}

func SetSystemLogging(s Setter, action, topics, key, value string) error {
	return s.Set("/system/logging",
		map[string]string{
			"action": action,
			"topics": topics,
		},
		map[string]string{
			key: value,
		})
}

func SetSystemLoggingPrefix(s Setter, action, topics, prefix string) error {
	return SetSystemLogging(s, action, topics, "prefix", prefix)
}

func SNMP(p Printer) (map[string]string, error) {
	return p.Print("/snmp", nil,
		[]string{
			"enabled",
			"location",
			"contact",
		},
		false,
	)
}
func SetSNMP(s Setter, key, value string) error {
	return s.Set("/snmp", nil,
		map[string]string{
			key: value,
		})
}

func SetSNMPEnabled(s BoolSetter, enabled bool) error {
	return SetSNMP(s, "enabled", s.FormatBool(enabled))
}
func SetSNMPLocation(s Setter, location string) error {
	return SetSNMP(s, "location", location)
}
func SetSNMPContact(s Setter, contact string) error {
	return SetSNMP(s, "contact", contact)
}

func InterfaceGRE(l Lister, name string) (map[string]string, error) {
	return firstList(l, "/interface/gre",
		map[string]string{
			"name": name,
		},
		[]string{
			"comment",
			"mtu",
			"clamp-tcp-mss",
			"dont-fragment",
			"allow-fast-path",
			"keepalive",
		},
		true,
	)
}
func SetInterfaceGRE(s Setter, name, key, value string) error {
	return s.Set("/snmp",
		map[string]string{
			name: "name",
		},
		map[string]string{
			key: value,
		})
}
func SetInterfaceGREComment(s Setter, name, comment string) error {
	return SetInterfaceGRE(s, name, "comment", comment)
}
func SetInterfaceGREMTU(s Setter, name, mtu string) error {
	return SetInterfaceGRE(s, name, "mtu", mtu)
}
func SetInterfaceGREClampTCPMSS(s BoolSetter, name string, clamp bool) error {
	return SetInterfaceGRE(s, name, "clamp-tcp-mss", s.FormatBool(clamp))
}
func SetInterfaceGREDontFragment(s BoolSetter, name string, dont bool) error {
	return SetInterfaceGRE(s, name, "dont-fragment", s.FormatBool(dont))
}
func SetInterfaceGREAllowFastPath(s BoolSetter, name string, allow bool) error {
	return SetInterfaceGRE(s, name, "allow-fast-path", s.FormatBool(allow))
}
func SetInterfaceGREKeepalive(s Setter, name, alive string) error {
	return SetInterfaceGRE(s, name, "keepalive", alive)
}

func RoutingBGPPeer(l Lister, instance, address string) (map[string]string, error) {
	return firstList(l, "/routing/bgp/peer",
		map[string]string{
			"interface":      instance,
			"remote-address": address,
		},
		[]string{
			"name",
			"comment",
		},
		true,
	)
}

func SetRoutingBGPPeer(s Setter, instance, address, key, value string) error {
	return s.Set("/routing/bgp/peer",
		map[string]string{
			"interface":      instance,
			"remote-address": address,
		},
		map[string]string{
			key: value,
		})
}
func SetRoutingBGPPeerName(s Setter, instance, address, name string) error {
	return SetRoutingBGPPeer(s, instance, address, "name", name)
}
func SetRoutingBGPPeerComment(s Setter, instance, address, comment string) error {
	return SetRoutingBGPPeer(s, instance, address, "comment", comment)
}

func IPDNS(p Printer) (map[string]string, error) {
	return p.Print("/ip/dns", nil,
		[]string{
			"servers",
			"allow-remote-requests",
		},
		false,
	)
}
func SetIPDNS(s Setter, key, value string) error {
	return s.Set("/ip/dns", nil,
		map[string]string{
			key: value,
		})
}
func SetIPDNSServers(s Setter, servers string) error {
	return SetIPDNS(s, "servers", servers)
}
func SetIPDNSAllowRemoteRequests(s BoolSetter, allow bool) error {
	return SetIPDNS(s, "allow-remote-requests", s.FormatBool(allow))
}

func IPAddress(l Lister, address string) (map[string]string, error) {
	return firstList(l, "/ip/address",
		map[string]string{
			"address": address,
		},
		[]string{
			"comment",
		},
		true,
	)
}
func SetIPAddress(s Setter, address, key, value string) error {
	return s.Set("/ip/address",
		map[string]string{
			"address": address,
		},
		map[string]string{
			key: value,
		})
}
func SetIPAddressComment(s Setter, address, comment string) error {
	return SetIPAddress(s, address, "comment", comment)
}

func User(l Lister, name string) (map[string]string, error) {
	return firstList(l, "/user",
		map[string]string{
			"name": name,
		},
		[]string{
			"comment",
			"group",
		},
		true,
	)
}
func SetUser(s Setter, name, key, value string) error {
	return s.Set("/user",
		map[string]string{
			name: "name",
		},
		map[string]string{
			key: value,
		})
}

func SetUserComment(s Setter, name, comment string) error {
	return SetUser(s, name, "comment", comment)
}
func SetUserGroup(s Setter, name, group string) error {
	return SetUser(s, name, "group", group)
}
func AddUser(a Adder, name, group, password string) error {
	if name != "admin" {
		return a.Add("/user",
			map[string]string{
				"name":     name,
				"group":    group,
				"password": password,
			})
	}
	return nil
}

func RemoveUser(r Remover, name string) error {
	if name != "admin" {
		return r.Remove("/user",
			map[string]string{
				"name": name,
			},
		)
	}
	return nil
}

func toolRomonLabel(legacy bool) string {
	if legacy {
		return "/romon"
	}
	return "/tool/romon"
}

func ToolRomon(p Printer, legacy bool) (map[string]string, error) {
	return p.Print(toolRomonLabel(legacy), nil,
		[]string{
			"id",
			"enabled",
			"secrets",
		},
		true,
	)
}

func SetToolRomon(s Setter, key, value string, legacy bool) error {
	return s.Set(toolRomonLabel(legacy), nil,
		map[string]string{
			key: value,
		})
}

func SetToolRomonId(s Setter, id string, legacy bool) error {
	return SetToolRomon(s, "id", id, legacy)
}
func SetToolRomonEnabled(s BoolSetter, enabled bool, legacy bool) error {
	return SetToolRomon(s, "enabled", s.FormatBool(enabled), legacy)
}
func SetToolRomonSecrets(s Setter, secrets string, legacy bool) error {
	return SetToolRomon(s, "secrets", secrets, legacy)
}

func ToolRomonPort(l Lister, iface string, legacy bool) (map[string]string, error) {
	return firstList(l, toolRomonLabel(legacy)+"/port",
		map[string]string{
			"interface": iface,
		},
		[]string{
			"secrets",
			"costs",
			"disabled",
			"forbid",
		},
		true,
	)
}

func SetToolRomonPort(s Setter, iface, key, value string, legacy bool) error {
	return s.Set(toolRomonLabel(legacy)+"/port",
		map[string]string{
			"interface": iface,
		},
		map[string]string{
			key: value,
		})
}

func SetToolRomonPortSecrets(s Setter, iface string, secrets string, legacy bool) error {
	return SetToolRomonPort(s, iface, "secrets", secrets, legacy)
}
func SetToolRomonPortCost(s Setter, iface string, cost int, legacy bool) error {
	return SetToolRomonPort(s, iface, "cost", strconv.Itoa(cost), legacy)
}
func SetToolRomonPortDisabled(s BoolSetter, iface string, disabled bool, legacy bool) error {
	return SetToolRomonPort(s, iface, "disabled", s.FormatBool(disabled), legacy)
}
func SetToolRomonPortForbid(s BoolSetter, iface string, forbid bool, legacy bool) error {
	return SetToolRomonPort(s, iface, "forbid", s.FormatBool(forbid), legacy)
}

func AddToolRomonPort(a Adder, iface string, legacy bool) error {
	if iface != "default" {
		return a.Add("/tool/romon/port",
			map[string]string{
				"interface": iface,
			})
	}
	return nil
}

func RemoveToolRomonPort(r Remover, iface string, legacy bool) error {
	if iface != "default" {
		return r.Remove("/tool/romon/port",
			map[string]string{
				"interface": iface,
			},
		)
	}
	return nil
}

func IPService(l Lister, name string) (map[string]string, error) {
	return firstList(l, "/ip/service",
		map[string]string{
			"name": name,
		},
		[]string{
			"disabled",
			"port",
			"address",
		},
		true,
	)
}

func SetIPService(s Setter, name, key, value string) error {
	return s.Set("/ip/service",
		map[string]string{
			"name": name,
		},
		map[string]string{
			key: value,
		})
}

func SetIPServiceDisabled(s BoolSetter, name string, disabled bool) error {
	return SetIPService(s, name, "disabled", s.FormatBool(disabled))
}
func SetIPServicePort(s Setter, name string, port int) error {
	return SetIPService(s, name, "port", strconv.Itoa(port))
}
func SetIPServiceAddress(s Setter, name string, address string) error {
	return SetIPService(s, name, "address", address)
}

func InterfaceList(l Lister) ([]map[string]string, error) {
	return l.List("/interface", nil, nil, true)
}

func AddressList(l Lister) ([]map[string]string, error) {
	return l.List("/ip/address", nil, nil, true)
}

func InterfaceBridgePortList(l Lister) ([]map[string]string, error) {
	return l.List("/interface/bridge/port", nil, nil, true)
}
