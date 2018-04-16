package ros

func ipFirewallNat(filter map[string]string) Command {
	return Command{
		Path:    "/ip firewall nat",
		Command: "print",
		Filter:  filter,
		Detail:  true,
	}
}

func (r *Ros) IpFirewallNat(filter map[string]string) (map[string]string, error) {
	return r.First(ipFirewallNat(filter))
}

func addIpFirewallNat(params map[string]string) Command {
	return Command{
		Path:    "/ip firewall nat",
		Command: "add",
		Params:  params,
	}
}
func (r *Ros) AddIpFirewallNat(params map[string]string) error {
	return r.Exec(addIpFirewallNat(params))
}

func removeIpFirewallNat(filter map[string]string) Command {
	return Command{
		Path:    "/ip firewall nat",
		Command: "remove",
		Filter:  filter,
	}
}
func (r *Ros) RemoveIpFirewallNat(filter map[string]string) error {
	return r.Exec(removeIpFirewallNat(filter))
}

func setIpFirewallNat(filter, params map[string]string) Command {
	return Command{
		Path:    "/ip firewall nat",
		Command: "set",
		Filter:  filter,
		Params:  params,
	}
}

func (r *Ros) SetIpFirewallNatDisabled(filter map[string]string, disabled bool) error {
	return r.Exec(setIpFirewallNat(filter, map[string]string{
		"disabled": FormatBool(disabled),
	}))
}

func (r *Ros) SetIpFirewallNatComment(filter map[string]string, comment string) error {
	return r.Exec(setIpFirewallNat(filter, map[string]string{
		"comment": comment,
	}))
}
