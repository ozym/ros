package ros

func ipAddresses() Command {
	return Command{
		Path:    "/ip address",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) IPAddresses() ([]map[string]string, error) {
	return r.List(ipAddresses())
}

func ipAddress(address string) Command {
	return Command{
		Path:    "/ip address",
		Command: "print",
		Filter: map[string]string{
			"address": address,
		},
		Detail: true,
	}
}

func (r Ros) IPAddress(address string) (map[string]string, error) {
	return r.First(ipAddress(address))
}

func setIPAddress(address, key, value string) Command {
	return Command{
		Path:    "/ip address",
		Command: "set",
		Filter: map[string]string{
			"address": address,
		},
		Params: map[string]string{
			key: value,
		},
	}
}

func (r Ros) SetIPAddressComment(address, comment string) error {
	return r.Exec(setIPAddress(address, "comment", comment))
}
