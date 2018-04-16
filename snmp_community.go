package ros

func snmpCommunityDefault() Command {
	return Command{
		Path: func() string {
			return "/snmp community"
		}(),
		Command: "print",
		Flags: map[string]bool{
			"default": true,
		},
	}
}
func (r *Ros) SnmpCommunityDefault() (map[string]string, error) {
	return r.First(snmpCommunityDefault())
}

func snmpCommunity(name string) Command {
	return Command{
		Path: func() string {
			return "/snmp community"
		}(),
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Flags: map[string]bool{
			"default": false,
		},
	}
}

func (r *Ros) SnmpCommunity(name string) (map[string]string, error) {
	return r.First(snmpCommunity(name))
}

func addSnmpCommunity(name string) Command {
	return Command{
		Path: func() string {
			return "/snmp community"
		}(),
		Command: "add",
		Params: map[string]string{
			"name": name,
		},
	}
}
func (r *Ros) AddSnmpCommunity(name string) error {
	return r.Exec(addSnmpCommunity(name))
}

func removeSnmpCommunity(name string) Command {
	return Command{
		Path: func() string {
			return "/snmp community"
		}(),
		Command: "remove",
		Flags: map[string]bool{
			"default": false,
		},
		Filter: map[string]string{
			"name": name,
		},
	}
}
func (r *Ros) RemoveSnmpCommunity(name string) error {
	return r.Exec(removeSnmpCommunity(name))
}

func setSnmpCommunityDefault(key, value string) Command {
	return Command{
		Path: func() string {
			return "/snmp community"
		}(),
		Command: "set",
		Filter: map[string]string{
			"default": "",
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setSnmpCommunity(name, key, value string) Command {
	return Command{
		Path: func() string {
			return "/snmp community"
		}(),
		Command: "set",
		Filter: map[string]string{
			"name":     name,
			"!default": "",
		},
		Params: map[string]string{
			key: value,
		},
	}
}

func (r *Ros) SetSnmpCommunityDefaultName(name string) error {
	return r.Exec(setSnmpCommunityDefault("name", name))
}
func (r *Ros) SetSnmpCommunityDefaultAddresses(addresses string) error {
	return r.Exec(setSnmpCommunityDefault("addresses", addresses))
}
func (r *Ros) SetSnmpCommunityDefaultAuthenticationPassword(password string) error {
	return r.Exec(setSnmpCommunityDefault("authentication-password", password))
}
func (r *Ros) SetSnmpCommunityDefaultAuthenticationProtocol(protocol string) error {
	return r.Exec(setSnmpCommunityDefault("authentication-protocol", protocol))
}
func (r *Ros) SetSnmpCommunityDefaultEncryptionPassword(password string) error {
	return r.Exec(setSnmpCommunityDefault("encryption-password", password))
}
func (r *Ros) SetSnmpCommunityDefaultEncryptionProtocol(protocol string) error {
	return r.Exec(setSnmpCommunityDefault("encryption-protocol", protocol))
}
func (r *Ros) SetSnmpCommunityDefaultSecurity(security string) error {
	return r.Exec(setSnmpCommunityDefault("security", security))
}
func (r *Ros) SetSnmpCommunityDefaultReadAccess(access bool) error {
	return r.Exec(setSnmpCommunityDefault("read-access", FormatBool(access)))
}
func (r *Ros) SetSnmpCommunityDefaultWriteAccess(access bool) error {
	return r.Exec(setSnmpCommunityDefault("write-access", FormatBool(access)))
}

func (r *Ros) SetSnmpCommunityAddresses(name, addresses string) error {
	return r.Exec(setSnmpCommunity(name, "addresses", addresses))
}
func (r *Ros) SetSnmpCommunityAuthenticationPassword(name, password string) error {
	return r.Exec(setSnmpCommunity(name, "authentication-password", password))
}
func (r *Ros) SetSnmpCommunityAuthenticationProtocol(name, protocol string) error {
	return r.Exec(setSnmpCommunity(name, "authentication-protocol", protocol))
}
func (r *Ros) SetSnmpCommunityEncryptionPassword(name, password string) error {
	return r.Exec(setSnmpCommunity(name, "encryption-password", password))
}
func (r *Ros) SetSnmpCommunityEncryptionProtocol(name, protocol string) error {
	return r.Exec(setSnmpCommunity(name, "encryption-protocol", protocol))
}
func (r *Ros) SetSnmpCommunitySecurity(name, security string) error {
	return r.Exec(setSnmpCommunity(name, "security", security))
}
func (r *Ros) SetSnmpCommunityReadAccess(name string, access bool) error {
	return r.Exec(setSnmpCommunity(name, "read-access", FormatBool(access)))
}
func (r *Ros) SetSnmpCommunityWriteAccess(name string, access bool) error {
	return r.Exec(setSnmpCommunity(name, "write-access", FormatBool(access)))
}
