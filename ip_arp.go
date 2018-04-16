package ros

func ipArps() Command {
	return Command{
		Path:    "/ip arp",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) IpArps() ([]map[string]string, error) {
	return r.List(ipArps())
}
