package ros

func interfaceBridgeHosts() Command {
	return Command{
		Path:    "/interface bridge host",
		Command: "print",
		Detail:  true,
	}
}

func (r *Ros) InterfaceBridgeHosts() ([]map[string]string, error) {
	return r.UnnumberedList(interfaceBridgeHosts(), 1)
}
