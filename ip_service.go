package ros

import (
	"strconv"
)

func ipService(name string) Command {
	return Command{
		Path:    "/ip service",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Detail: true,
	}
}

func (r Ros) IPService(name string) (map[string]string, error) {
	return r.First(ipService(name))
}

func setIPService(name, key, value string) Command {
	return Command{
		Path:    "/ip service",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetIPServiceDisabled(name string, disabled bool) error {
	return r.Exec(setIPService(name, "disabled", FormatBool(disabled)))
}

func (r Ros) SetIPServicePort(name string, port int) error {
	return r.Exec(setIPService(name, "port", strconv.Itoa(port)))
}

func (r Ros) SetIPServiceAddress(name, address string) error {
	return r.Exec(setIPService(name, "address", address))
}
