package ros

func routingOSPFInterfaces() Command {
	return Command{
		Path:    "/routing ospf interface",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) RoutingOSPFInterfaces() ([]map[string]string, error) {
	return r.List(routingOSPFInterfaces())
}

func routingOSPFInterface(iface string) Command {
	return Command{
		Path:    "/routing ospf interface",
		Command: "print",
		Filter: map[string]string{
			"interface": iface,
		},
		Detail: true,
	}
}

func (r Ros) RoutingOSPFInterface(iface string) (map[string]string, error) {
	return r.First(routingOSPFInterface(iface))
}

func setRoutingOSPFInterface(iface, key, value string) Command {
	return Command{
		Path:    "/routing ospf interface",
		Command: "set",
		Filter: map[string]string{
			"interface": iface,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetRoutingOSPFInterfaceComment(iface, comment string) error {
	return r.Exec(setRoutingOSPFInterface(iface, "comment", comment))
}
func (r Ros) SetRoutingOSPFInterfaceCost(iface, cost string) error {
	return r.Exec(setRoutingOSPFInterface(iface, "cost", cost))
}
