package ros

import (
	"strconv"
)

func ipSsh() Command {
	return Command{
		Path:    "/ip ssh",
		Command: "print",
	}
}

func (r *Ros) IpSsh() (map[string]string, error) {
	return r.Values(ipSsh())
}

func setIpSsh(key, value string) Command {
	return Command{
		Path:    "/ip ssh",
		Command: "set",
		Params: map[string]string{
			key: value,
		},
	}
}
func (r *Ros) SetIpSshStrongCrypto(strong bool) error {
	return r.Exec(setIpSsh("strong-crypto", FormatBool(strong)))
}
func (r *Ros) SetIpSshForwardingEnabled(forwarding bool) error {
	return r.Exec(setIpSsh("forwarding-enabled", FormatBool(forwarding)))
}
func (r *Ros) SetIpSshAlwaysAllowPasswordLogin(allow bool) error {
	return r.Exec(setIpSsh("always-allow-password-login", FormatBool(allow)))
}
func (r *Ros) SetIpSshHostKeySize(size int) error {
	return r.Exec(setIpSsh("host-key-size", strconv.Itoa(size)))
}
