package ros

func routingBGPNetworks() Command {
	return Command{
		Path:    "/routing bgp network",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) RoutingBGPNetworks() ([]map[string]string, error) {
	return r.List(routingBGPNetworks())
}

func routingBGPNetwork(network string) Command {
	return Command{
		Path:    "/routing bgp network",
		Command: "print",
		Filter: map[string]string{
			"network": network,
		},
		Detail: true,
	}
}

func (r Ros) RoutingBGPNetwork(network string) (map[string]string, error) {
	return r.First(routingBGPNetwork(network))
}

func setRoutingBGPNetwork(network, key, value string) Command {
	return Command{
		Path:    "/routing bgp network",
		Command: "set",
		Filter: map[string]string{
			"network": network,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetRoutingBGPNetworkComment(network, comment string) error {
	return r.Exec(setRoutingBGPNetwork(network, "comment", comment))
}
