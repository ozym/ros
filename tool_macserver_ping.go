package ros

func toolMacServerPing() Command {
	return Command{
		Path:    "/tool mac-server ping",
		Command: "print",
	}
}

func (r *Ros) ToolMacServerPing() (map[string]string, error) {
	return r.Values(toolMacServerPing())
}

func setToolMacServerPing(key, value string) Command {
	return Command{
		Path:    "/tool mac-server ping",
		Command: "set",
		Params: map[string]string{
			key: value,
		},
	}
}

func (r *Ros) SetToolMacServerPingEnabled(enabled bool) error {
	return r.Exec(setToolMacServerPing("enabled", FormatBool(enabled)))
}
