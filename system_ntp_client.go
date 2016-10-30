package ros

func systemNTPClient() Command {
	return Command{
		Path:    "/system ntp client",
		Command: "print",
	}
}

func (r Ros) SystemNTPClient() (map[string]string, error) {
	return r.Values(systemNTPClient())
}

func setSystemNTPClientEnabled(enabled bool) Command {
	return Command{
		Path:    "/system ntp client",
		Command: "set",
		Params: map[string]string{
			"enabled": FormatBool(enabled),
		},
	}
}
func (r Ros) SetSystemNTPClientEnabled(enabled bool) error {
	return r.Exec(setSystemNTPClientEnabled(enabled))
}
func setSystemNTPClientPrimaryNTP(host string) Command {
	return Command{
		Path:    "/system ntp client",
		Command: "set",
		Params: map[string]string{
			"primary-ntp": host,
		},
	}
}
func (r Ros) SetSystemNTPClientPrimaryNTP(zone string) error {
	return r.Exec(setSystemNTPClientPrimaryNTP(zone))
}
func setSystemNTPClientSecondaryNTP(host string) Command {
	return Command{
		Path:    "/system ntp client",
		Command: "set",
		Params: map[string]string{
			"secondary-ntp": host,
		},
	}
}
func (r Ros) SetSystemNTPClientSecondaryNTP(zone string) error {
	return r.Exec(setSystemNTPClientSecondaryNTP(zone))
}
