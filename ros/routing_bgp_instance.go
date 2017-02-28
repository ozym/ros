package ros

func routingBGPInstances() Command {
	return Command{
		Path:    "/routing bgp instance",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) RoutingBGPInstances() ([]map[string]string, error) {
	return r.List(routingBGPInstances())
}

func routingBGPInstance(name string) Command {
	return Command{
		Path:    "/routing bgp instance",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Detail: true,
	}
}

func (r Ros) RoutingBGPInstance(name string) (map[string]string, error) {
	return r.First(routingBGPInstance(name))
}

func setRoutingBGPInstance(name, key, value string) Command {
	return Command{
		Path:    "/routing bgp instance",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}

func (r Ros) SetRoutingBGPInstanceRouterId(name, router_id string) error {
	return r.Exec(setRoutingBGPInstance(name, "router-id", router_id))
}

func (r Ros) SetRoutingBGPInstanceComment(name, comment string) error {
	return r.Exec(setRoutingBGPInstance(name, "comment", comment))
}
