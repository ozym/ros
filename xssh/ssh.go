package xssh

import (
	"bufio"
	"net"
	"strconv"
	"strings"

	"github.com/ScriptRock/crypto/ssh"
	"github.com/ozym/ros"
)

type SSH struct {
	client *ssh.Client
}

func New(hostname, user, password string) (*ros.MikroTik, error) {

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		Config: ssh.Config{
			Ciphers: ssh.AllSupportedCiphers(),
		},
	}
	hostname = func() string {
		if strings.Contains(hostname, ":") {
			return hostname
		}
		return hostname + ":22"
	}()

	client, err := ssh.Dial("tcp", hostname, config)
	if err != nil {
		return nil, err
	}

	hostname, _, err = net.SplitHostPort(hostname)
	if err != nil {
		return nil, err
	}

	return &ros.MikroTik{
		Client: &SSH{
			client: client,
		},
		Protocol: "xssh",
		Hostname: hostname,
	}, nil
}

func (s SSH) FormatBool(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}
func (s SSH) ParseBool(x string) bool {
	if x == "yes" {
		return true
	}
	return false
}

func (s SSH) Print(base string, filter map[string]string, properties []string, detail bool) (map[string]string, error) {

	b := "/" + strings.Join(strings.Split(strings.TrimPrefix(base, "/"), "/"), " ") + " print"
	if detail {
		b = b + " detail"
	}
	if filter != nil {
		b = b + " where"
		for k, v := range filter {
			b = b + " " + k + "=" + strconv.Quote(v)
		}
	}
	r, err := sshRunResults(s.client, b)
	if err != nil {
		return nil, err
	}
	if !(len(properties) > 0) {
		return r, nil
	}

	res := make(map[string]string)
	for _, k := range properties {
		if v, ok := r[k]; ok {
			res[k] = v
		}
	}

	return res, nil
}

func (s SSH) List(base string, properties []string) ([]map[string]string, error) {

	b := "/" + strings.Join(strings.Split(strings.TrimPrefix(base, "/"), "/"), " ") + " print"

	lines, err := sshRunLines(s.client, b)
	if err != nil {
		return nil, err
	}
	list, err := ros.ScanNumberedItemList(strings.Join(lines, "\n"))
	if err != nil {
		return nil, err
	}

	if len(properties) > 0 {
		var res []map[string]string
		for _, l := range list {
			r := make(map[string]string)
			for _, k := range properties {
				if v, ok := l[k]; ok {
					r[k] = v
				}
			}
			res = append(res, r)
		}
		list = res
	}

	return list, nil
}

func (s SSH) Set(base string, filter map[string]string, settings map[string]string) error {
	b := "/" + strings.Join(strings.Split(strings.TrimPrefix(base, "/"), "/"), " ") + " set"
	if filter != nil {
		b = b + " [find where"
		for k, v := range filter {
			b = b + " " + k + "=" + strconv.Quote(v)
		}
		b = b + " ]"
	}
	for k, v := range settings {
		b = b + " " + k + "=" + strconv.Quote(v)
	}
	err := sshRun(s.client, b)
	if err != nil {
		return err
	}

	return nil
}
func (s SSH) Add(base string, options map[string]string) error {
	b := "/" + strings.Join(strings.Split(strings.TrimPrefix(base, "/"), "/"), " ") + " add"
	for k, v := range options {
		b = b + " " + k + "=" + strconv.Quote(v)
	}
	err := sshRun(s.client, b)
	if err != nil {
		return err
	}
	return nil
}
func (s SSH) Remove(base string, filter map[string]string) error {
	b := "/" + strings.Join(strings.Split(strings.TrimPrefix(base, "/"), "/"), " ") + " remove"
	if filter != nil {
		b = b + " where"
		for k, v := range filter {
			b = b + " " + k + "=" + strconv.Quote(v)
		}
	}
	err := sshRun(s.client, b)
	if err != nil {
		return err
	}
	return nil
}

func sshRunLines(client *ssh.Client, command string) ([]string, error) {
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	var lines []string

	stdout, err := session.StdoutPipe()
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(stdout)

	err = session.Run(command)
	if err != nil {
		return nil, err
	}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func sshRun(client *ssh.Client, command string) error {
	_, err := sshRunLines(client, command)
	return err
}

func sshRunResults(client *ssh.Client, command string) (map[string]string, error) {
	lines, err := sshRunLines(client, command)
	if err != nil {
		return nil, err
	}

	res, err := ros.ScanItems(strings.Join(lines, "\n"))
	if err != nil {
		return nil, err
	}

	return res, nil
}
