package ros

func ipFirewallFilter(chain string) Command {
	return Command{
		Path:    "/routing filter",
		Command: "print",
		Filter: map[string]string{
			"chain": chain,
		},
		Detail: true,
	}
}

func (r *Ros) IpFirewallFilter(chain string) ([]map[string]string, error) {
	return r.List(ipFirewallFilter(chain))
}

func addIpFirewallFilter(params map[string]string) Command {
	return Command{
		Path:    "/routing filter",
		Command: "add",
		Params:  params,
	}
}
func (r *Ros) AddIpFirewallFilter(chain string, params []map[string]string) error {
	//return r.Exec(addIpFirewallFilter(params))
	return nil
}

func removeIpFirewallFilter(chain string) Command {
	/*
		return Command{
			Path:    "/routing filter",
			Command: "remove",
			Filter:  filter,
		}
	*/
	return Command{}
}

func (r *Ros) RemoveIpFirewallFilter(chain string) error {
	return nil
	//return r.Exec(removeIpFirewallFilter(filter))
}

func setIpFirewallFilter(chain, params map[string]string) Command {
	/*
		return Command{
			Path:    "/routing filter",
			Command: "set",
			Filter:  filter,
			Params:  params,
		}
	*/
	return Command{}
}

/*
func (r *Ros) SetIpFirewallFilterDisabled(chain filter map[string]string, disabled bool) error {
	return r.Exec(setIpFirewallFilter(filter, map[string]string{
		"disabled": FormatBool(disabled),
	}))
}

func (r *Ros) SetIpFirewallFilterComment(filter map[string]string, comment string) error {
	return r.Exec(setIpFirewallFilter(filter, map[string]string{
		"comment": comment,
	}))
}
*/
