package ros

func routingOSPFNBMANeighbors() Command {
	return Command{
		Path:    "/routing ospf nbma-neighbor",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) RoutingOSPFNBMANeighbors() ([]map[string]string, error) {
	return r.List(routingOSPFNBMANeighbors())
}

func routingOSPFNBMANeighbor(address string) Command {
	return Command{
		Path:    "/routing ospf nbma-neighbor",
		Command: "print",
		Filter: map[string]string{
			"address": address,
		},
		Detail: true,
	}
}

func (r Ros) RoutingOSPFNBMANeighbor(address string) (map[string]string, error) {
	return r.First(routingOSPFNBMANeighbor(address))
}

func setRoutingOSPFNBMANeighbor(address, key, value string) Command {
	return Command{
		Path:    "/routing ospf nbma-neighbor",
		Command: "set",
		Filter: map[string]string{
			"address": address,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetRoutingOSPFNBMANeighborComment(address, comment string) error {
	return r.Exec(setRoutingOSPFNBMANeighbor(address, "comment", comment))
}
