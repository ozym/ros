package ros

func interfaceListMember(iface, list string) Command {
	return Command{
		Path:    "/interface list member",
		Command: "print",
		Filter: map[string]string{
			"interface": iface,
			"list":      list,
		},
		Detail: true,
	}
}

func (r *Ros) InterfaceListMember(iface, list string) (map[string]string, error) {
	return r.First(interfaceListMember(iface, list))
}

func addInterfaceListMember(iface, list string, params map[string]string) Command {
	return Command{
		Path:    "/interface list member",
		Command: "add",
		Params: map[string]string{
			"interface": iface,
			"list":      list,
		},
		Extra: params,
	}
}
func (r *Ros) AddInterfaceListMember(iface, list string, params map[string]string) error {
	switch list {
	case "all":
		return nil
	default:
		return r.Exec(addInterfaceListMember(iface, list, params))
	}
}

func removeInterfaceListMember(iface, list string) Command {
	return Command{
		Path:    "/interface list member",
		Command: "remove",
		Filter: map[string]string{
			"interface": iface,
			"list":      list,
		},
	}
}
func (r *Ros) RemoveInterfaceListMember(iface, list string) error {
	switch list {
	case "all":
		return nil
	default:
		return r.Exec(removeInterfaceListMember(iface, list))
	}
}

func setInterfaceListMember(iface, list string, params map[string]string) Command {
	return Command{
		Path:    "/interface list member",
		Command: "set",
		Filter: map[string]string{
			"interface": iface,
			"list":      list,
		},
		Params: params,
	}
}

func (r *Ros) SetInterfaceListMemberComment(iface, list, comment string) error {
	switch list {
	case "all":
		return nil
	default:
		return r.Exec(setInterfaceListMember(iface, list, map[string]string{"comment": comment}))
	}
}

func (r *Ros) SetInterfaceListMemberDisabled(iface, list string, disabled bool) error {
	switch list {
	case "all":
		return nil
	default:
		return r.Exec(setInterfaceListMember(iface, list, map[string]string{"disabled": FormatBool(disabled)}))
	}
}
