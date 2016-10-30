package ros

func routingBGPPeer(iface, address string) Command {
	return Command{
		Path:    "/routing bgp peer",
		Command: "print",
		Filter: map[string]string{
			"interface":      iface,
			"remote-address": address,
		},
		Detail: true,
	}
}

func (r Ros) RoutingBGPPeer(iface, address string) (map[string]string, error) {
	return r.First(routingBGPPeer(iface, address))
}

func setRoutingBGPPeer(iface, address, key, value string) Command {
	return Command{
		Path:    "/routing bgp peer",
		Command: "set",
		Filter: map[string]string{
			"interface":      iface,
			"remote-address": address,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetRoutingBGPPeerComment(iface, address, comment string) error {
	return r.Exec(setRoutingBGPPeer(iface, address, "comment", comment))
}
func (r Ros) SetRoutingBGPPeerName(iface, address, name string) error {
	return r.Exec(setRoutingBGPPeer(iface, address, "name", name))
}
