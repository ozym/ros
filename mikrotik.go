package ros

type BoolParser interface {
	ParseBool(string) bool
}
type BoolFormatter interface {
	FormatBool(bool) string
}

type Printer interface {
	Print(base string, filter map[string]string, properties []string, detail bool) (map[string]string, error)
}
type Lister interface {
	List(base string, filter map[string]string, properties []string, detail bool) ([]map[string]string, error)
}
type Setter interface {
	Set(base string, filter map[string]string, settings map[string]string) error
}
type BoolSetter interface {
	BoolFormatter
	Setter
}
type Adder interface {
	Add(base string, options map[string]string) error
}
type Remover interface {
	Remove(base string, filter map[string]string) error
}

type Client interface {
	BoolFormatter
	BoolParser
	Printer
	Lister
	Setter
	Adder
	Remover
}

type MikroTik struct {
	Client

	Protocol string
	Hostname string
}

// Recover Router Name
func (m MikroTik) Id() string {
	return m.Hostname
}

// Recover Router Access Method
func (m MikroTik) Method() string {
	return m.Protocol
}

// Recover Router Resources
func (m MikroTik) SystemResource() (map[string]string, error) {
	return SystemResource(m)
}

// Manage Router Identity
func (m MikroTik) SystemIdentity() (map[string]string, error) {
	return SystemIdentity(m)
}
func (m MikroTik) SetSystemIdentityName(name string) error {
	return SetSystemIdentityName(m, name)
}

// Manage Router Notes
func (m MikroTik) SystemNote() (map[string]string, error) {
	return SystemNote(m)
}
func (m MikroTik) SetSystemNote(note string) error {
	return SetSystemNote(m, note)
}
func (m MikroTik) SetSystemNoteShowAtLogin(show bool) error {
	return SetSystemNoteShowAtLogin(m, show)
}

// Manage Router Clock Settings
func (m MikroTik) SystemClock() (map[string]string, error) {
	return SystemClock(m)
}
func (m MikroTik) SetSystemClockTimeZoneName(name string) error {
	return SetSystemClockTimeZoneName(m, name)
}
func (m MikroTik) SetSystemClockTimeZoneAutodetect(auto bool) error {
	return SetSystemClockTimeZoneAutodetect(m, auto)
}

// Manage Router NTP Client Settings
func (m MikroTik) SystemNTPClient() (map[string]string, error) {
	return SystemNTPClient(m)
}
func (m MikroTik) SetSystemNTPClientEnabled(enable bool) error {
	return SetSystemNTPClientEnabled(m, enable)
}
func (m MikroTik) SetSystemNTPClientPrimaryNTP(name string) error {
	return SetSystemNTPClientPrimaryNTP(m, name)
}
func (m MikroTik) SetSystemNTPClientSecondaryNTP(name string) error {
	return SetSystemNTPClientSecondaryNTP(m, name)
}

// Manage Router Logging Settings
func (m MikroTik) SystemLogging(action, topics string) (map[string]string, error) {
	return SystemLogging(m, action, topics)
}
func (m MikroTik) AddSystemLogging(action, topics string) error {
	return AddSystemLogging(m, action, topics)
}
func (m MikroTik) RemoveSystemLogging(action, topics string) error {
	return RemoveSystemLogging(m, action, topics)
}
func (m MikroTik) SetSystemLoggingPrefix(action, topics, prefix string) error {
	return SetSystemLoggingPrefix(m, action, topics, prefix)
}

func (m MikroTik) SystemLoggingAction(name string) (map[string]string, error) {
	return SystemLoggingAction(m, name)
}
func (m MikroTik) SetSystemLoggingActionRemote(name, remote string) error {
	return SetSystemLoggingActionRemote(m, name, remote)
}
func (m MikroTik) SetSystemLoggingActionTarget(name, target string) error {
	return SetSystemLoggingActionTarget(m, name, target)
}
func (m MikroTik) SetSystemLoggingActionRemotePort(name string, port int) error {
	return SetSystemLoggingActionRemotePort(m, name, port)
}
func (m MikroTik) SetSystemLoggingActionBSDSyslog(name string, bsd bool) error {
	return SetSystemLoggingActionBSDSyslog(m, name, bsd)
}
func (m MikroTik) SetSystemLoggingActionSrcAddress(name, address string) error {
	return SetSystemLoggingActionSrcAddress(m, name, address)
}
func (m MikroTik) SetSystemLoggingActionSyslogFacility(name, facility string) error {
	return SetSystemLoggingActionSyslogFacility(m, name, facility)
}
func (m MikroTik) SetSystemLoggingActionSyslogSeverity(name, severity string) error {
	return SetSystemLoggingActionSyslogSeverity(m, name, severity)
}
func (m MikroTik) SetSystemLoggingActionSyslogTimeFormat(name, format string) error {
	return SetSystemLoggingActionSyslogTimeFormat(m, name, format)
}

// Manage Router SNMP Settings
func (m MikroTik) SNMP() (map[string]string, error) {
	return SNMP(m)
}
func (m MikroTik) SetSNMPEnabled(enabled bool) error {
	return SetSNMPEnabled(m, enabled)
}
func (m MikroTik) SetSNMPContact(contact string) error {
	return SetSNMPContact(m, contact)
}
func (m MikroTik) SetSNMPLocation(location string) error {
	return SetSNMPLocation(m, location)
}

// Manage Router GRE Tunnel Settings
func (m MikroTik) InterfaceGRE(name string) (map[string]string, error) {
	return InterfaceGRE(m, name)
}
func (m MikroTik) SetInterfaceGREComment(name, comment string) error {
	return SetInterfaceGREComment(m, name, comment)
}
func (m MikroTik) SetInterfaceGREMTU(name, mtu string) error {
	return SetInterfaceGREMTU(m, name, mtu)
}
func (m MikroTik) SetInterfaceGREClampTCPMSS(name string, clamp bool) error {
	return SetInterfaceGREClampTCPMSS(m, name, clamp)
}
func (m MikroTik) SetInterfaceGREDontFragment(name string, dont bool) error {
	return SetInterfaceGREDontFragment(m, name, dont)
}
func (m MikroTik) SetInterfaceGREAllowFastPath(name string, allow bool) error {
	return SetInterfaceGREAllowFastPath(m, name, allow)
}
func (m MikroTik) SetInterfaceGREKeepalive(name, alive string) error {
	return SetInterfaceGREKeepalive(m, name, alive)
}

// Manage BGP Peer Settings
func (m MikroTik) RoutingBGPPeer(instance, address string) (map[string]string, error) {
	return RoutingBGPPeer(m, instance, address)
}
func (m MikroTik) SetRoutingBGPPeerName(instance, address, name string) error {
	return SetRoutingBGPPeerName(m, instance, address, name)
}
func (m MikroTik) SetRoutingBGPPeerComment(instance, address, comment string) error {
	return SetRoutingBGPPeerComment(m, instance, address, comment)
}

// Manage Router DNS Settings
func (m MikroTik) IPDNS() (map[string]string, error) {
	return IPDNS(m)
}
func (m MikroTik) SetIPDNSServers(servers string) error {
	return SetIPDNSServers(m, servers)
}
func (m MikroTik) SetIPDNSAllowRemoteRequests(allow bool) error {
	return SetIPDNSAllowRemoteRequests(m, allow)
}

// Manage Router IP Addresses
func (m MikroTik) IPAddress(address string) (map[string]string, error) {
	return IPAddress(m, address)
}
func (m MikroTik) SetIPAddressComment(address, comment string) error {
	return SetIPAddressComment(m, address, comment)
}

// Manage Router Users
func (m MikroTik) User(name string) (map[string]string, error) {
	return User(m, name)
}
func (m MikroTik) SetUserComment(name, comment string) error {
	return SetUserComment(m, name, comment)
}
func (m MikroTik) SetUserGroup(name, group string) error {
	return SetUserGroup(m, name, group)
}
func (m MikroTik) AddUser(name, group, password string) error {
	return AddUser(m, name, group, password)
}
func (m MikroTik) RemoveUser(name string) error {
	return RemoveUser(m, name)
}

// Manage Router Romon Service
func (m MikroTik) ToolRomon() (map[string]string, error) {
	return ToolRomon(m, false)
}
func (m MikroTik) SetToolRomonId(id string) error {
	return SetToolRomonId(m, id, false)
}
func (m MikroTik) SetToolRomonEnabled(enabled bool) error {
	return SetToolRomonEnabled(m, enabled, false)
}
func (m MikroTik) SetToolRomonSecrets(secrets string) error {
	return SetToolRomonSecrets(m, secrets, false)
}
func (m MikroTik) AddToolRomonPort(iface string) error {
	return AddToolRomonPort(m, iface, false)
}
func (m MikroTik) RemoveToolRomonPort(iface string) error {
	return RemoveToolRomonPort(m, iface, false)
}
func (m MikroTik) ToolRomonPort(iface string) (map[string]string, error) {
	return ToolRomonPort(m, iface, false)
}
func (m MikroTik) SetToolRomonPortDisabled(iface string, disabled bool) error {
	return SetToolRomonPortDisabled(m, iface, disabled, false)
}
func (m MikroTik) SetToolRomonPortCost(iface string, cost int) error {
	return SetToolRomonPortCost(m, iface, cost, false)
}
func (m MikroTik) SetToolRomonPortForbid(iface string, forbid bool) error {
	return SetToolRomonPortForbid(m, iface, forbid, false)
}
func (m MikroTik) SetToolRomonPortSecrets(iface string, secrets string) error {
	return SetToolRomonPortSecrets(m, iface, secrets, false)
}

// Manage Route Legacy Romon Service
func (m MikroTik) LegacyToolRomon() (map[string]string, error) {
	return ToolRomon(m, true)
}
func (m MikroTik) SetLegacyToolRomonId(id string) error {
	return SetToolRomonId(m, id, true)
}
func (m MikroTik) SetLegacyToolRomonEnabled(enabled bool) error {
	return SetToolRomonEnabled(m, enabled, true)
}
func (m MikroTik) SetLegacyToolRomonSecrets(secrets string) error {
	return SetToolRomonSecrets(m, secrets, true)
}
func (m MikroTik) AddLegacyToolRomonPort(iface string) error {
	return AddToolRomonPort(m, iface, true)
}
func (m MikroTik) RemoveLegacyToolRomonPort(iface string) error {
	return RemoveToolRomonPort(m, iface, true)
}
func (m MikroTik) ToolLegacyRomonPort(iface string) (map[string]string, error) {
	return ToolRomonPort(m, iface, true)
}
func (m MikroTik) SetLegacyToolRomonPortDisabled(iface string, disabled bool) error {
	return SetToolRomonPortDisabled(m, iface, disabled, true)
}
func (m MikroTik) SetLegacyToolRomonPortCost(iface string, cost int) error {
	return SetToolRomonPortCost(m, iface, cost, true)
}
func (m MikroTik) SetLegacyToolRomonPortForbid(iface string, forbid bool) error {
	return SetToolRomonPortForbid(m, iface, forbid, true)
}
func (m MikroTik) SetLegacyToolRomonPortSecrets(iface string, secrets string) error {
	return SetToolRomonPortSecrets(m, iface, secrets, true)
}

// Manage Router IP Services
func (m MikroTik) IPService(name string) (map[string]string, error) {
	return IPService(m, name)
}

func (m MikroTik) SetIPServiceDisabled(name string, disabled bool) error {
	if name != "ssh" {
		return SetIPServiceDisabled(m, name, disabled)
	}
	return nil
}
func (m MikroTik) SetIPServicePort(name string, port int) error {
	if name != "ssh" {
		return SetIPServicePort(m, name, port)
	}
	return nil
}
func (m MikroTik) SetIPServiceAddress(name string, address string) error {
	return SetIPServiceAddress(m, name, address)
}

// List Router Interfaces
func (m MikroTik) InterfaceList() ([]map[string]string, error) {
	return InterfaceList(m)
}

// List Router Addresses
func (m MikroTik) AddressList() ([]map[string]string, error) {
	return AddressList(m)
}

// List Router Interface Bridge Ports
func (m MikroTik) InterfaceBridgePortList() ([]map[string]string, error) {
	return InterfaceBridgePortList(m)
}
