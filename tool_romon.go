package ros

func toolRomon(legacy bool) Command {
	return Command{
		Path: func() string {
			if legacy {
				return "/romon"
			}
			return "/tool romon"
		}(),
		Command: "print",
	}
}

func (r *Ros) HasRomon() bool {

	switch {
	case r.Major() < 6:
		return false
	case r.Major() > 6:
		return true
	case r.Minor() < 28:
		return false
	default:
		return true
	}
}

func (r *Ros) HasLegacyRomon() bool {

	switch {
	case r.Major() < 6:
		return false
	case r.Major() > 6:
		return false
	case r.Minor() != 28:
		return false
	default:
		return true
	}
}

func (r *Ros) ToolRomon() (map[string]string, error) {
	return r.Values(toolRomon(r.HasLegacyRomon()))
}

func setToolRomon(key, value string, legacy bool) Command {
	return Command{
		Path: func() string {
			if legacy {
				return "/romon"
			}
			return "/tool romon"
		}(),
		Command: "set",
		Params: map[string]string{
			key: value,
		},
	}
}

func (r *Ros) SetToolRomonId(id string) error {
	return r.Exec(setToolRomon("id", id, r.HasLegacyRomon()))
}
func (r *Ros) SetToolRomonEnabled(enabled bool) error {
	return r.Exec(setToolRomon("enabled", FormatBool(enabled), r.HasLegacyRomon()))
}
func (r *Ros) SetToolRomonSecrets(secrets string) error {
	return r.Exec(setToolRomon("secrets", secrets, r.HasLegacyRomon()))
}
