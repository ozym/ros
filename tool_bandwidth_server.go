package ros

import (
	"strconv"
)

// only manage the default mac-server
func toolBandwidthServer() Command {
	return Command{
		Path:    "/tool bandwidth-server",
		Command: "print",
	}
}

func (r *Ros) ToolBandwidthServer() (map[string]string, error) {
	return r.Values(toolBandwidthServer())
}

func setToolBandwidthServer(key, value string) Command {
	return Command{
		Path:    "/tool bandwidth-server",
		Command: "set",
		Params: map[string]string{
			key: value,
		},
	}
}

func (r *Ros) SetToolBandwidthServerEnabled(enabled bool) error {
	return r.Exec(setToolBandwidthServer("enabled", FormatBool(enabled)))
}
func (r *Ros) SetToolBandwidthServerAuthenticate(auth bool) error {
	return r.Exec(setToolBandwidthServer("authenticate", FormatBool(auth)))
}
func (r *Ros) SetToolBandwidthServerAllocateUdpPortsFrom(port int) error {
	return r.Exec(setToolBandwidthServer("allocate-udp-ports-from", strconv.Itoa(port)))
}
func (r *Ros) SetToolBandwidthServerMaxSessions(sessions int) error {
	return r.Exec(setToolBandwidthServer("max-sessions", strconv.Itoa(sessions)))
}
