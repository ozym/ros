package ros

import (
//	"strings"
)

func systemScheduler(name string) Command {
	return Command{
		Path:    "/system scheduler",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Detail: true,
	}
}

func (r *Ros) SystemScheduler(name string) (map[string]string, error) {
	return r.First(systemScheduler(name))
}

func addSystemScheduler(name, interval, policy, onevent string) Command {
	return Command{
		Path:    "/system scheduler",
		Command: "add",
		Params: map[string]string{
			"name":     name,
			"interval": interval,
			"policy":   policy,
			"on-event": onevent,
		},
	}
}

func (r *Ros) AddSystemScheduler(name, interval, policy, onevent string) error {
	return r.Exec(addSystemScheduler(name, interval, policy, onevent))
}

func removeSystemScheduler(name string) Command {
	return Command{
		Path:    "/system scheduler",
		Command: "remove",
		Filter: map[string]string{
			"name": name,
		},
	}
}

func (r *Ros) RemoveSystemScheduler(name string) error {
	return r.Exec(removeSystemScheduler(name))
}

func setSystemScheduler(name, key, value string) Command {
	return Command{
		Path:    "/system scheduler",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}

func (r *Ros) SetSystemSchedulerInterval(name, interval string) error {
	return r.Exec(setSystemScheduler(name, "interval", interval))
}
func (r *Ros) SetSystemSchedulerPolicy(name, policy string) error {
	return r.Exec(setSystemScheduler(name, "policy", policy))
}
func (r *Ros) SetSystemSchedulerComment(name, comment string) error {
	return r.Exec(setSystemScheduler(name, "comment", comment))
}
func (r *Ros) SetSystemSchedulerOnEvent(name, onevent string) error {
	return r.Exec(setSystemScheduler(name, "on-event", onevent))
}
func (r *Ros) SetSystemSchedulerStartDate(name, startdate string) error {
	return r.Exec(setSystemScheduler(name, "start_date", startdate))
}
func (r *Ros) SetSystemSchedulerStartTime(name, starttime string) error {
	return r.Exec(setSystemScheduler(name, "start_time", starttime))
}
func (r *Ros) SetSystemSchedulerStartup(name string, startup bool) error {
	return r.Exec(setSystemScheduler(name, "startup", FormatBool(startup)))
}
