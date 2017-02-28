package ros

func snmp() Command {
	return Command{
		Path:    "/snmp",
		Command: "print",
	}
}

func (r Ros) SNMP() (map[string]string, error) {
	return r.Values(snmp())
}

func setSNMP(key, value string) Command {
	return Command{
		Path:    "/snmp",
		Command: "set",
		Params: map[string]string{
			"key": value,
		},
	}
}

func (r Ros) SetSNMPEnabled(enabled bool) error {
	return r.Exec(setSNMP("enabled", FormatBool(enabled)))
}
func (r Ros) SetSNMPLocation(location string) error {
	return r.Exec(setSNMP("location", location))
}
func (r Ros) SetSNMPContact(contact string) error {
	return r.Exec(setSNMP("contact", contact))
}
