package ros

import (
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	//TODO: remove after firware update
	"github.com/ScriptRock/crypto/ssh"
	//TODO: add after firware update
	//"golang.org/x/crypto/ssh"
)

const (
	DefaultPort     = 22
	DefaultUsername = "admin"
	DefaultPassword = ""
	DefaultTimeout  = 30 * time.Second
)

type Ros struct {
	client *ssh.Client
	config *ssh.ClientConfig

	hostname string
	port     int

	major       int
	minor       int
	routerboard bool
	once        sync.Once

	mu  sync.Mutex
	err error
}

func Port(port int) func(*Ros) error {
	return func(r *Ros) error {
		r.port = port
		return nil
	}
}

func Username(username string) func(*Ros) error {
	return func(r *Ros) error {
		r.config.User = username
		return nil
	}
}

func Password(password string) func(*Ros) error {
	return func(r *Ros) error {
		r.config.Auth = []ssh.AuthMethod{
			ssh.Password(password),
		}
		return nil
	}
}

func Timeout(timeout time.Duration) func(*Ros) error {
	return func(r *Ros) error {
		r.config.Timeout = timeout
		return nil
	}
}

func NewRos(hostname string, options ...func(*Ros) error) (*Ros, error) {

	host, port := hostname, DefaultPort
	if strings.Contains(hostname, ":") {
		var p string
		var err error

		host, p, err = net.SplitHostPort(hostname)
		if err != nil {
			return nil, err
		}

		port, err = strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
	}

	r := &Ros{
		config: &ssh.ClientConfig{
			User: DefaultUsername,
			Auth: []ssh.AuthMethod{
				ssh.Password(DefaultPassword),
			},
			//TODO: remove after firware update
			Config: ssh.Config{
				Ciphers: ssh.AllSupportedCiphers(),
			},
			Timeout:         DefaultTimeout,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		},

		hostname: host,
		port:     port,
	}

	for _, option := range options {
		if err := option(r); err != nil {
			return nil, err
		}
	}

	return r, nil
}

func (r *Ros) Connect() error {

	if r.err != nil {
		return r.err
	}

	if r.client != nil {
		return nil
	}

	hostname := net.JoinHostPort(r.hostname, strconv.Itoa(r.port))

	client, err := ssh.Dial("tcp", hostname, r.config)
	if err != nil {
		r.err = err

		return err
	}

	r.client = client

	return nil
}

func (r *Ros) Close() error {
	if r.err != nil {
		return r.err
	}
	if r.client == nil {
		return nil
	}
	if err := r.client.Close(); err != nil {
		return err
	}

	return nil
}

func (r *Ros) Version() error {

	if r.err != nil {
		return r.err
	}

	r.once.Do(func() {
		res, err := r.SystemResource()
		if err != nil {
			r.err = err
			return
		}
		if v, ok := res["version"]; ok {
			major, minor := RouterOsVersion(v)

			r.major, r.minor = major, minor
		}
		if b, ok := res["routerboard"]; ok {
			r.routerboard = ParseBool(b)
		}
	})

	return r.err
}

func (r *Ros) Id() string {
	return r.hostname
}

func (r *Ros) Error() error {
	return r.err
}

func (r *Ros) Major() int {
	return r.major
}

func (r *Ros) Minor() int {
	return r.minor
}

func (r *Ros) AtLeast(major, minor int) bool {
	switch {
	case r.Major() < major:
		return false
	case r.Major() > major:
		return true
	case r.Minor() < minor:
		return false
	default:
		return true
	}
}
func (r *Ros) Routerboard() bool {
	return r.routerboard
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

func (r *Ros) Parse(c Command) (string, error) {
	return c.Parse()
}

func (r *Ros) Run(c Command) ([]string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.err != nil {
		return nil, r.err
	}

	if err := r.Connect(); err != nil {
		return nil, err
	}

	res, err := c.Run(r.client)
	if err != nil {
		r.err = err

		return res, err
	}

	return res, nil
}

func (r *Ros) Exec(c Command) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.err != nil {
		return r.err
	}

	if err := r.Connect(); err != nil {
		return err
	}

	err := c.Exec(r.client)
	if err != nil {
		r.err = err

		return err
	}

	return nil
}

func (r *Ros) Values(c Command) (map[string]string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.err != nil {
		return nil, r.err
	}

	if err := r.Connect(); err != nil {
		return nil, err
	}

	res, err := c.Values(r.client)
	if err != nil {
		r.err = err

		return res, err
	}
	return res, nil
}

func (r *Ros) List(c Command) ([]map[string]string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.err != nil {
		return nil, r.err
	}

	if err := r.Connect(); err != nil {
		return nil, err
	}

	res, err := c.List(r.client)
	if err != nil {
		r.err = err

		return res, err
	}

	return res, nil
}

func (r *Ros) UnnumberedList(c Command, header int) ([]map[string]string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.err != nil {
		return nil, r.err
	}

	if err := r.Connect(); err != nil {
		return nil, err
	}

	res, err := c.UnnumberedList(r.client, header)
	if err != nil {
		r.err = err

		return res, err
	}

	return res, nil
}

func (r *Ros) First(c Command) (map[string]string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.err != nil {
		return nil, r.err
	}

	if err := r.Connect(); err != nil {
		return nil, err
	}

	res, err := c.First(r.client)
	if err != nil {
		r.err = err

		return res, err
	}

	return res, nil
}
