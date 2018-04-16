package ros

func systemRouterboard() Command {
	return Command{
		Path:    "/system routerboard",
		Command: "print",
	}
}

func (r *Ros) SystemRouterboard() (map[string]string, error) {
	return r.Values(systemRouterboard())
}
