package ros

// only manage the default mac-server mac-winbox
func toolMacServerMacWinbox(legacy bool) Command {
	if legacy {
		return Command{
			Path:    "/tool mac-server mac-winbox",
			Command: "print",
			Flags: map[string]bool{
				"default": true,
			},
			Detail: true,
		}
	}
	return Command{
		Path:    "/tool mac-server mac-winbox",
		Command: "print",
		Detail:  true,
	}
}

func (r *Ros) ToolMacServerMacWinbox() (map[string]string, error) {
	return r.First(toolMacServerMacWinbox(!r.AtLeast(6, 41)))
}

func setToolMacServerMacWinbox(key, value string, legacy bool) Command {
	if legacy {
		return Command{
			Path:    "/tool mac-server mac-winbox",
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
		Path:    "/tool mac-server mac-winbox",
		Command: "set",
		Params: map[string]string{
			key: value,
		},
	}
}

func (r *Ros) SetToolMacServerMacWinboxDisabled(disabled bool) error {
	if !r.AtLeast(6, 41) {
		return r.Exec(setToolMacServerMacWinbox("disabled", FormatBool(disabled), true))
	}
	return nil
}

func (r *Ros) SetToolMacServerMacWinboxAllowedInterfaceList(list string) error {
	if r.AtLeast(6, 41) {
		return r.Exec(setToolMacServerMacWinbox("allowed-interface-list", list, false))
	}
	return nil
}
