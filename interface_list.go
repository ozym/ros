package ros

func interfaceList(list string) Command {
	return Command{
		Path:    "/interface list",
		Command: "print",
		Filter: map[string]string{
			"name": list,
		},
		Detail: true,
	}
}

func (r *Ros) InterfaceList(list string) (map[string]string, error) {
	return r.First(interfaceList(list))
}

func addInterfaceList(list string, params map[string]string) Command {
	return Command{
		Path:    "/interface list",
		Command: "add",
		Params: map[string]string{
			"name": list,
		},
		Extra: params,
	}
}
func (r *Ros) AddInterfaceList(list string, params map[string]string) error {
	switch list {
	case "all":
		return nil
	default:
		return r.Exec(addInterfaceList(list, params))
	}
}

func removeInterfaceList(list string) Command {
	return Command{
		Path:    "/interface list",
		Command: "remove",
		Filter: map[string]string{
			"name": list,
		},
	}
}
func (r *Ros) RemoveInterfaceList(list string) error {
	switch list {
	case "all":
		return nil
	default:
		return r.Exec(removeInterfaceList(list))
	}
}

func setInterfaceList(list string, params map[string]string) Command {
	return Command{
		Path:    "/interface list",
		Command: "set",
		Filter: map[string]string{
			"name": list,
		},
		Params: params,
	}
}

func (r *Ros) SetInterfaceListComment(list, comment string) error {
	switch list {
	case "all":
		return nil
	default:
		return r.Exec(setInterfaceList(list, map[string]string{"comment": comment}))
	}
}
