package ros

import (
	"strings"
)

func ipFirewallAddressList(address, list string) Command {
	return Command{
		Path:    "/ip firewall address-list",
		Command: "print",
		Filter: map[string]string{
			"address": strings.TrimSuffix(address, "/32"),
			"list":    list,
		},
		Detail: true,
	}
}

func (r *Ros) IpFirewallAddressList(address, list string) (map[string]string, error) {
	return r.First(ipFirewallAddressList(address, list))
}

func addIpFirewallAddressList(address, list string, params map[string]string) Command {
	return Command{
		Path:    "/ip firewall address-list",
		Command: "add",
		Params: map[string]string{
			"address": strings.TrimSuffix(address, "/32"),
			"list":    list,
		},
		Extra: params,
	}
}
func (r *Ros) AddIpFirewallAddressList(address, list string, params map[string]string) error {
	return r.Exec(addIpFirewallAddressList(address, list, params))
}

func removeIpFirewallAddressList(address, list string) Command {
	return Command{
		Path:    "/ip firewall address-list",
		Command: "remove",
		Filter: map[string]string{
			"address": strings.TrimSuffix(address, "/32"),
			"list":    list,
		},
	}
}
func (r *Ros) RemoveIpFirewallAddressList(address, list string) error {
	return r.Exec(removeIpFirewallAddressList(address, list))
}

func setIpFirewallAddressList(address, list string, params map[string]string) Command {
	return Command{
		Path:    "/ip firewall address-list",
		Command: "set",
		Filter: map[string]string{
			"address": address,
			"list":    list,
		},
		Params: params,
	}
}

func (r *Ros) SetIpFirewallAddressListComment(address, list, comment string) error {
	return r.Exec(setIpFirewallAddressList(address, list, map[string]string{"comment": comment}))
}

func (r *Ros) SetIpFirewallAddressListDisabled(address, list string, disabled bool) error {
	return r.Exec(setIpFirewallAddressList(address, list, map[string]string{"disabled": FormatBool(disabled)}))
}
