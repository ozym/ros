package ros

import (
	"fmt"
	"sync"

	"github.com/ScriptRock/crypto/ssh"
)

type Ros struct {
	Client   *ssh.Client
	Error    error
	Hostname string

	Major int
	Minor int

	mu sync.Mutex
}

func (r *Ros) Version() error {
	if r.Major == 0 {
		res, err := r.SystemResource()
		if err != nil {
			return err
		}
		if _, ok := res["version"]; !ok {
			return fmt.Errorf("no version found")
		}

		major, minor := RouterOSVersion(res["version"])

		r.Major, r.Minor = major, minor
	}
	return nil
}

func (r Ros) Id() string {
	return r.Hostname
}

func FormatBool(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}
func ParseBool(x string) bool {
	if x == "yes" {
		return true
	}
	return false
}

func (r Ros) Parse(c Command) (string, error) {
	return c.Parse()
}
func (r Ros) Run(c Command) ([]string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return c.Run(r.Client)
}
func (r Ros) Exec(c Command) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return c.Exec(r.Client)
}
func (r Ros) Values(c Command) (map[string]string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return c.Values(r.Client)
}
func (r Ros) List(c Command) ([]map[string]string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return c.List(r.Client)
}
func (r Ros) First(c Command) (map[string]string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return c.First(r.Client)
}
