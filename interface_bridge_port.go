package ros

func interfaceBridgePorts() Command {
	return Command{
		Path:    "/interface bridge port",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) InterfaceBridgePorts() ([]map[string]string, error) {
	return r.List(interfaceBridgePorts())
}
