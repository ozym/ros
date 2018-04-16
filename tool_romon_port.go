package ros

import (
	"strconv"
)

func toolRomonPortDefault(legacy bool) Command {
	return Command{
		Path: func() string {
			if legacy {
				return "/romon port"
			}
			return "/tool romon port"
		}(),
		Command: "print",
		Flags: map[string]bool{
			"default": true,
		},
		Detail: true,
	}
}

func (r *Ros) ToolRomonPortDefault() (map[string]string, error) {
	return r.First(toolRomonPortDefault(r.HasLegacyRomon()))
}

func toolRomonPort(iface string, legacy bool) Command {
	return Command{
		Path: func() string {
			if legacy {
				return "/romon port"
			}
			return "/tool romon port"
		}(),
		Command: "print",
		Filter: map[string]string{
			"interface": iface,
		},
		Flags: map[string]bool{
			"default": false,
		},
		Detail: true,
	}
}

func (r *Ros) ToolRomonPort(iface string) (map[string]string, error) {
	return r.First(toolRomonPort(iface, r.HasLegacyRomon()))
}

func addToolRomonPort(iface string, legacy bool) Command {
	return Command{
		Path: func() string {
			if legacy {
				return "/romon port"
			}
			return "/tool romon port"
		}(),
		Command: "add",
		Params: map[string]string{
			"interface": iface,
		},
	}
}
func (r *Ros) AddToolRomonPort(iface string) error {
	return r.Exec(addToolRomonPort(iface, r.HasLegacyRomon()))
}

func removeToolRomonPort(iface string, legacy bool) Command {
	return Command{
		Path: func() string {
			if legacy {
				return "/romon port"
			}
			return "/tool romon port"
		}(),
		Command: "remove",
		Flags: map[string]bool{
			"default": false,
		},
		Filter: map[string]string{
			"interface": iface,
		},
	}
}
func (r *Ros) RemoveToolRomonPort(iface string) error {
	return r.Exec(removeToolRomonPort(iface, r.HasLegacyRomon()))
}

func setToolRomonPortDefault(key, value string, legacy bool) Command {
	return Command{
		Path: func() string {
			if legacy {
				return "/romon port"
			}
			return "/tool romon port"
		}(),
		Command: "set",
		Filter: map[string]string{
			"default": "",
		},
		Params: map[string]string{
			key: value,
		},
	}
}

func (r *Ros) SetToolRomonPortDefaultSecrets(secrets string) error {
	return r.Exec(setToolRomonPortDefault("secrets", secrets, r.HasLegacyRomon()))
}
func (r *Ros) SetToolRomonPortDefaultCost(cost int) error {
	return r.Exec(setToolRomonPortDefault("cost", strconv.Itoa(cost), r.HasLegacyRomon()))
}
func (r *Ros) SetToolRomonPortDefaultDisabled(disabled bool) error {
	return r.Exec(setToolRomonPortDefault("disabled", FormatBool(disabled), r.HasLegacyRomon()))
}
func (r *Ros) SetToolRomonPortDefaultForbid(forbid bool) error {
	return r.Exec(setToolRomonPortDefault("forbid", FormatBool(forbid), r.HasLegacyRomon()))
}

func setToolRomonPort(iface, key, value string, legacy bool) Command {
	return Command{
		Path: func() string {
			if legacy {
				return "/romon port"
			}
			return "/tool romon port"
		}(),
		Command: "set",
		Filter: map[string]string{
			"interface": iface,
			"!default":  "",
		},
		Params: map[string]string{
			key: value,
		},
	}
}

func (r *Ros) SetToolRomonPortSecrets(iface, secrets string) error {
	return r.Exec(setToolRomonPort(iface, "secrets", secrets, r.HasLegacyRomon()))
}
func (r *Ros) SetToolRomonPortCost(iface string, cost int) error {
	return r.Exec(setToolRomonPort(iface, "cost", strconv.Itoa(cost), r.HasLegacyRomon()))
}
func (r *Ros) SetToolRomonPortDisabled(iface string, disabled bool) error {
	return r.Exec(setToolRomonPort(iface, "disabled", FormatBool(disabled), r.HasLegacyRomon()))
}
func (r *Ros) SetToolRomonPortForbid(iface string, forbid bool) error {
	return r.Exec(setToolRomonPort(iface, "forbid", FormatBool(forbid), r.HasLegacyRomon()))
}
