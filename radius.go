package ros

func radius(service string) Command {
	return Command{
		Path:    "/radius",
		Command: "print",
		Filter: map[string]string{
			"service": service,
		},
		Detail: true,
	}
}

func (r *Ros) Radius(service string) (map[string]string, error) {
	return r.First(radius(service))
}

func addRadius(service string) Command {
	return Command{
		Path:    "/radius",
		Command: "add",
		Params: map[string]string{
			"service": service,
		},
	}
}

func (r *Ros) AddRadius(service string) error {
	return r.Exec(addRadius(service))
}

func removeRadius(service string) Command {
	return Command{
		Path:    "/radius",
		Command: "remove",
		Filter: map[string]string{
			"service": service,
		},
	}
}

func (r *Ros) RemoveRadius(service string) error {
	return r.Exec(removeRadius(service))
}

func setRadiusAddress(service, address string) Command {
	return Command{
		Path:    "/radius",
		Command: "set",
		Filter: map[string]string{
			"service": service,
		},
		Params: map[string]string{
			"address": address,
		},
	}
}
func (r *Ros) SetRadiusAddress(service, address string) error {
	return r.Exec(setRadiusAddress(service, address))
}

func setRadiusSecret(service, secret string) Command {
	return Command{
		Path:    "/radius",
		Command: "set",
		Filter: map[string]string{
			"service": service,
		},
		Params: map[string]string{
			"secret": secret,
		},
	}
}
func (r *Ros) SetRadiusSecret(service, secret string) error {
	return r.Exec(setRadiusSecret(service, secret))
}

func setRadiusTimeout(service, timeout string) Command {
	return Command{
		Path:    "/radius",
		Command: "set",
		Filter: map[string]string{
			"service": service,
		},
		Params: map[string]string{
			"timeout": timeout,
		},
	}
}
func (r *Ros) SetRadiusTimeout(service, timeout string) error {
	return r.Exec(setRadiusTimeout(service, timeout))
}

func setRadiusSrcAddress(service, address string) Command {
	return Command{
		Path:    "/radius",
		Command: "set",
		Filter: map[string]string{
			"service": service,
		},
		Params: map[string]string{
			"src-address": address,
		},
	}
}

func (r *Ros) SetRadiusSrcAddress(service, address string) error {
	return r.Exec(setRadiusSrcAddress(service, address))
}

func setRadiusDisabled(service, disabled string) Command {
	return Command{
		Path:    "/radius",
		Command: "set",
		Filter: map[string]string{
			"service": service,
		},
		Params: map[string]string{
			"disabled": disabled,
		},
	}
}

func (r *Ros) SetRadiusDisabled(service string, disabled bool) error {
	return r.Exec(setRadiusDisabled(service, FormatBool(disabled)))
}
