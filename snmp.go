package ros

import (
	"strconv"
)

func snmp() Command {
	return Command{
		Path:    "/snmp",
		Command: "print",
	}
}

func (r *Ros) Snmp() (map[string]string, error) {
	return r.Values(snmp())
}

func setSnmp(key, value string) Command {
	return Command{
		Path:    "/snmp",
		Command: "set",
		Params: map[string]string{
			key: value,
		},
	}
}

func (r *Ros) SetSnmpEnabled(enabled bool) error {
	return r.Exec(setSnmp("enabled", FormatBool(enabled)))
}
func (r *Ros) SetSnmpLocation(location string) error {
	return r.Exec(setSnmp("location", location))
}
func (r *Ros) SetSnmpContact(contact string) error {
	return r.Exec(setSnmp("contact", contact))
}
func (r *Ros) SetSnmpEngineId(id string) error {
	return r.Exec(setSnmp("engine-id", id))
}
func (r *Ros) SetSnmpTrapCommunity(community string) error {
	return r.Exec(setSnmp("trap-community", community))
}
func (r *Ros) SetSnmpTrapGenerators(generators string) error {
	return r.Exec(setSnmp("trap-generators", generators))
}
func (r *Ros) SetSnmpTrapTarget(target string) error {
	return r.Exec(setSnmp("trap-target", target))
}
func (r *Ros) SetSnmpTrapVersion(version int) error {
	return r.Exec(setSnmp("trap-version", strconv.Itoa(version)))
}
