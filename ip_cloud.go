package ros

func ipCloud() Command {
	return Command{
		Path:    "/ip cloud",
		Command: "print",
	}
}

func (r *Ros) HasCloud() bool {

	switch {
	case !r.Routerboard():
		return false
	case r.Major() < 6:
		return false
	case r.Major() > 6:
		return true
	case r.Minor() < 14:
		return false
	default:
		return true
	}
}

func (r *Ros) HasLegacyCloud() bool {

	switch {
	case r.Major() < 6:
		return false
	case r.Major() > 6:
		return false
	case r.Minor() >= 27:
		return false
	default:
		return true
	}
}

func (r *Ros) IpCloud() (map[string]string, error) {
	return r.Values(ipCloud())
}

func setIpCloud(key, value string) Command {
	return Command{
		Path:    "/ip cloud",
		Command: "set",
		Params: map[string]string{
			key: value,
		},
	}
}

func (r *Ros) SetIpCloudDdnsEnabled(enabled bool) error {
	if r.HasCloud() {
		if r.HasLegacyCloud() {
			return r.Exec(setIpCloud("enabled", FormatBool(enabled)))
		}
		return r.Exec(setIpCloud("ddns-enabled", FormatBool(enabled)))
	}
	return nil
}

func (r *Ros) SetIpCloudUpdateTime(update bool) error {
	if r.HasCloud() {
		return r.Exec(setIpCloud("update-time", FormatBool(update)))
	}
	return nil
}
