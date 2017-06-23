package ros

func ipDNS() Command {
	return Command{
		Path:    "/ip dns",
		Command: "print",
	}
}

func (r Ros) IPDNS() (map[string]string, error) {
	return r.Values(ipDNS())
}

func setIPDNS(key, value string) Command {
	return Command{
		Path:    "/ip dns",
		Command: "set",
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetIPDNSServers(servers string) error {
	return r.Exec(setIPDNS("servers", servers))
}
func (r Ros) SetIPDNSAllowRemoteRequests(allow bool) error {
	return r.Exec(setIPDNS("allow-remote-requests", FormatBool(allow)))
}
