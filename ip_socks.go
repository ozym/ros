package ros

import (
	"strconv"
)

// only manage the default mac-server
func ipSocks() Command {
	return Command{
		Path:    "/ip socks",
		Command: "print",
	}
}

func (r *Ros) IpSocks() (map[string]string, error) {
	return r.Values(ipSocks())
}

func setIpSocks(key, value string) Command {
	return Command{
		Path:    "/ip socks",
		Command: "set",
		Params: map[string]string{
			key: value,
		},
	}
}

func (r *Ros) SetIpSocksEnabled(enabled bool) error {
	return r.Exec(setIpSocks("enabled", FormatBool(enabled)))
}
func (r *Ros) SetIpSocksPort(port int) error {
	return r.Exec(setIpSocks("port", strconv.Itoa(port)))
}
func (r *Ros) SetIpSocksConnectionIdleTimeout(timeout string) error {
	return r.Exec(setIpSocks("connection-idle-timeout", timeout))
}
func (r *Ros) SetIpSocksMaxConnections(connections int) error {
	return r.Exec(setIpSocks("max-connections", strconv.Itoa(connections)))
}
