package ros

func routingOSPFInstances() Command {
	return Command{
		Path:    "/routing ospf instance",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) RoutingOSPFInstances() ([]map[string]string, error) {
	return r.List(routingOSPFInstances())
}

func routingOSPFInstance(name string) Command {
	return Command{
		Path:    "/routing ospf instance",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Detail: true,
	}
}

func (r Ros) RoutingOSPFInstance(name string) (map[string]string, error) {
	return r.First(routingOSPFInstance(name))
}

func setRoutingOSPFInstance(name, key, value string) Command {
	return Command{
		Path:    "/routing ospf instance",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetRoutingOSPFInstanceRouterId(name, router_id string) error {
	return r.Exec(setRoutingOSPFInstance(name, "router-id", router_id))
}
func (r Ros) SetRoutingOSPFInstanceComment(name, comment string) error {
	return r.Exec(setRoutingOSPFInstance(name, "comment", comment))
}
