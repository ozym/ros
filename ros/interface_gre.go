package ros

func interfaceGREs() Command {
	return Command{
		Path:    "/interface gre",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) InterfaceGREs() ([]map[string]string, error) {
	return r.List(interfaceGREs())
}

func interfaceGRE(address string) Command {
	return Command{
		Path:    "/interface gre",
		Command: "print",
		Filter: map[string]string{
			"remote-address": address,
		},
		Detail: true,
	}
}

func (r Ros) InterfaceGRE(address string) (map[string]string, error) {
	return r.First(interfaceGRE(address))
}

func setInterfaceGRE(address, key, value string) Command {
	return Command{
		Path:    "/interface gre",
		Command: "set",
		Filter: map[string]string{
			"remote-address": address,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetInterfaceGREName(address, name string) error {
	return r.Exec(setInterfaceGRE(address, "name", name))
}
func (r Ros) SetInterfaceGREComment(address, comment string) error {
	return r.Exec(setInterfaceGRE(address, "comment", comment))
}
func (r Ros) SetInterfaceGREMTU(address, mtu string) error {
	return r.Exec(setInterfaceGRE(address, "mtu", mtu))
}
func (r Ros) SetInterfaceGREKeepalive(address, alive string) error {
	return r.Exec(setInterfaceGRE(address, "keepalive", alive))
}
func (r Ros) SetInterfaceGREClampTCPMSS(address string, clamp bool) error {
	return r.Exec(setInterfaceGRE(address, "clamp-tcp-mss", FormatBool(clamp)))
}
func (r Ros) SetInterfaceGREAllowFastPath(address string, allow bool) error {
	return r.Exec(setInterfaceGRE(address, "allow-fast-path", FormatBool(allow)))
}
