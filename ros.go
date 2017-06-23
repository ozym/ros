package ros

import (
	"fmt"
	"sync"

	"github.com/ScriptRock/crypto/ssh"
)

type Ros struct {
	sync.Mutex

	Client   *ssh.Client
	Hostname string
	Major    int
	Minor    int
}

func New(client *ssh.Client, hostname string) (*Ros, error) {
	r := Ros{
		Client:   client,
		Hostname: hostname,
	}

	res, err := r.SystemResource()
	if err != nil {
		return nil, err
	}

	if _, ok := res["version"]; !ok {
		return nil, fmt.Errorf("no version found")
	}

	minor, major := RouterOSVersion(res["version"])

	r.Major = major
	r.Minor = minor

	return &r, nil
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
	return c.Run(r.Client)
}
func (r Ros) Exec(c Command) error {
	return c.Exec(r.Client)
}
func (r Ros) Values(c Command) (map[string]string, error) {
	return c.Values(r.Client)
}
func (r Ros) List(c Command) ([]map[string]string, error) {
	return c.List(r.Client)
}
func (r Ros) First(c Command) (map[string]string, error) {
	return c.First(r.Client)
}
