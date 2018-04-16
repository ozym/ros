package ros

// only manage the default mac-server
func toolMacServer(legacy bool) Command {
	if legacy {
		return Command{
			Path:    "/tool mac-server",
			Command: "print",
			Flags: map[string]bool{
				"default": true,
			},
			Detail: true,
		}
	}
	return Command{
		Path:    "/tool mac-server",
		Command: "print",
		Detail:  true,
	}
}

func (r *Ros) ToolMacServer() (map[string]string, error) {
	return r.First(toolMacServer(!r.AtLeast(6, 41)))
}

func setToolMacServer(key, value string, legacy bool) Command {
	if legacy {
		return Command{
			Path:    "/tool mac-server",
			Command: "set",
			Filter: map[string]string{
				"default": "",
			},
			Params: map[string]string{
				key: value,
			},
		}
	}
	return Command{
		Path:    "/tool mac-server",
		Command: "set",
		Params: map[string]string{
			key: value,
		},
	}
}

func (r *Ros) SetToolMacServerDisabled(disabled bool) error {
	if !r.AtLeast(6, 41) {
		return r.Exec(setToolMacServer("disabled", FormatBool(disabled), true))
	}
	return nil
}

func (r *Ros) SetToolMacServerAllowedInterfaceList(list string) error {
	if r.AtLeast(6, 41) {
		return r.Exec(setToolMacServer("allowed-interface-list", list, false))
	}
	return nil
}
