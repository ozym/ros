package ros

func ipNeighbors() Command {
	return Command{
		Path:    "/ip neighbor",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) IpNeighbors() ([]map[string]string, error) {
	return r.List(ipNeighbors())
}
