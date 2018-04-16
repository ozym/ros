package ros

func routingFilter(filter map[string]string) Command {
	return Command{
		Path:    "/routing filter",
		Command: "print",
		Filter:  filter,
		Detail:  true,
	}
}

func (r *Ros) RoutingFilter(filter map[string]string) (map[string]string, error) {
	return r.First(routingFilter(filter))
}

func routingFilterChain(chain string) Command {
	return Command{
		Path:    "/routing filter",
		Command: "print",
		Filter: map[string]string{
			"chain": chain,
		},

		Detail: true,
	}
}

func (r *Ros) RoutingFilterChain(chain string) ([]map[string]string, error) {
	return r.List(routingFilterChain(chain))
}

func addRoutingFilter(params map[string]string) Command {
	return Command{
		Path:    "/routing filter",
		Command: "add",
		Params:  params,
	}
}
func (r *Ros) AddRoutingFilter(params map[string]string) error {
	return r.Exec(addRoutingFilter(params))
}

func removeRoutingFilter(filter map[string]string) Command {
	return Command{
		Path:    "/routing filter",
		Command: "remove",
		Filter:  filter,
	}
}
func (r *Ros) RemoveRoutingFilter(filter map[string]string) error {
	return r.Exec(removeRoutingFilter(filter))
}

func setRoutingFilter(filter, params map[string]string) Command {
	return Command{
		Path:    "/routing filter",
		Command: "set",
		Filter:  filter,
		Params:  params,
	}
}

func (r *Ros) SetRoutingFilterDisabled(filter map[string]string, disabled bool) error {
	return r.Exec(setRoutingFilter(filter, map[string]string{
		"disabled": FormatBool(disabled),
	}))
}

func (r *Ros) SetRoutingFilterComment(filter map[string]string, comment string) error {
	return r.Exec(setRoutingFilter(filter, map[string]string{
		"comment": comment,
	}))
}
