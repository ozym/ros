package ros

func routingBGPAggregates() Command {
	return Command{
		Path:    "/routing bgp aggregate",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) RoutingBGPAggregates() ([]map[string]string, error) {
	return r.List(routingBGPAggregates())
}

func routingBGPAggregate(instance, prefix string) Command {
	return Command{
		Path:    "/routing bgp aggregate",
		Command: "print",
		Filter: map[string]string{
			"instance": instance,
			"prefix":   prefix,
		},
		Detail: true,
	}
}

func (r Ros) RoutingBGPAggregate(instance, prefix string) (map[string]string, error) {
	return r.First(routingBGPAggregate(instance, prefix))
}

func setRoutingBGPAggregate(instance, prefix, key, value string) Command {
	return Command{
		Path:    "/routing bgp aggregate",
		Command: "set",
		Filter: map[string]string{
			"instance": instance,
			"prefix":   prefix,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetRoutingBGPAggregateComment(instance, prefix, comment string) error {
	return r.Exec(setRoutingBGPAggregate(instance, prefix, "comment", comment))
}
