package ros

func routingOSPFNetworks() Command {
	return Command{
		Path:    "/routing ospf network",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) RoutingOSPFNetworks() ([]map[string]string, error) {
	return r.List(routingOSPFNetworks())
}

func routingOSPFNetwork(network string) Command {
	return Command{
		Path:    "/routing ospf network",
		Command: "print",
		Filter: map[string]string{
			"network": network,
		},
		Detail: true,
	}
}

func (r Ros) RoutingOSPFNetwork(network string) (map[string]string, error) {
	return r.First(routingOSPFNetwork(network))
}

func setRoutingOSPFNetwork(network, key, value string) Command {
	return Command{
		Path:    "/routing ospf network",
		Command: "set",
		Filter: map[string]string{
			"network": network,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetRoutingOSPFNetworkComment(network, comment string) error {
	return r.Exec(setRoutingOSPFNetwork(network, "comment", comment))
}
