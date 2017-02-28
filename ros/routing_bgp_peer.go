package ros

func routingBGPPeers() Command {
	return Command{
		Path:    "/routing bgp peer",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) RoutingBGPPeers() ([]map[string]string, error) {
	return r.List(routingBGPPeers())
}

func routingBGPPeer(addr string) Command {
	return Command{
		Path:    "/routing bgp peer",
		Command: "print",
		Filter: map[string]string{
			"remote-address": addr,
		},
		Detail: true,
	}
}

func (r Ros) RoutingBGPPeer(addr string) (map[string]string, error) {
	return r.First(routingBGPPeer(addr))
}

func setRoutingBGPPeer(addr, key, value string) Command {
	return Command{
		Path:    "/routing bgp peer",
		Command: "set",
		Filter: map[string]string{
			"remote-address": addr,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetRoutingBGPPeerComment(addr, comment string) error {
	return r.Exec(setRoutingBGPPeer(addr, "comment", comment))
}
func (r Ros) SetRoutingBGPPeerName(addr, name string) error {
	return r.Exec(setRoutingBGPPeer(addr, "name", name))
}
