package api

import (
	"net"
	"strings"

	"github.com/go-routeros/routeros"
	"github.com/ozym/ros"
)

type API struct {
	client *routeros.Client
}

func New(hostname, user, password string) (*ros.MikroTik, error) {

	// required to be host:port
	hostname = func() string {
		if strings.Contains(hostname, ":") {
			return hostname
		}
		return hostname + ":8728"
	}()

	// tcp connection to the api port
	client, err := routeros.Dial(hostname, user, password)
	if err != nil {
		return nil, err
	}

	// allow multiple requests
	client.Async()

	hostname, _, err = net.SplitHostPort(hostname)
	if err != nil {
		return nil, err
	}

	return &ros.MikroTik{
		Client: &API{
			client: client,
		},
		Protocol: "api",
		Hostname: hostname,
	}, nil
}

func (a API) FormatBool(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
func (a API) ParseBool(s string) bool {
	if s == "true" {
		return true
	}
	return false
}

func (a API) Print(base string, filter map[string]string, properties []string, detail bool) (map[string]string, error) {
	command := []string{base + "/print"}
	if detail {
		command = append(command, "=detail=")
	}
	for k, v := range filter {
		command = append(command, "?="+k+"="+v)
	}
	if len(properties) > 0 {
		command = append(command, "=.proplist="+strings.Join([]string(properties), ","))
	}
	return apiRunResult(a.client, command)
}

func (a API) List(base string, filter map[string]string, properties []string, detail bool) ([]map[string]string, error) {
	command := []string{base + "/print"}
	if detail {
		command = append(command, "=detail=")
	}
	for k, v := range filter {
		command = append(command, "?="+k+"="+v)
	}
	if len(properties) > 0 {
		command = append(command, "=.proplist="+strings.Join([]string(properties), ","))
	}
	return apiRunResultList(a.client, command)
}

func (a API) Set(base string, filter map[string]string, settings map[string]string) error {
	command := []string{base + "/set"}
	for k, v := range filter {
		command = append(command, "?="+k+"="+v)
	}
	for k, v := range settings {
		command = append(command, "="+k+"="+v)
	}
	return apiRun(a.client, command)
}

func (a API) Add(base string, options map[string]string) error {
	command := []string{base + "/add"}
	for k, v := range options {
		command = append(command, "="+k+"="+v)
	}
	return apiRun(a.client, command)
}

func (a API) Remove(base string, filter map[string]string) error {
	command := []string{base + "/remove"}
	for k, v := range filter {
		command = append(command, "?="+k+"="+v)
	}
	return apiRun(a.client, command)
}

// Run a command with no results required.
func apiRun(client *routeros.Client, command []string) error {
	_, err := client.RunArgs(command)
	return err
}

// Run a command with a results map.
func apiRunResult(client *routeros.Client, command []string) (map[string]string, error) {
	reply, err := client.RunArgs(command)
	if err != nil {
		return nil, err
	}

	res := make(map[string]string)
	for _, re := range reply.Re {
		for k, v := range re.Map {
			res[k] = v
		}
	}

	return res, nil
}

// Run a command with a results slice of maps.
func apiRunResultList(client *routeros.Client, command []string) ([]map[string]string, error) {
	reply, err := client.RunArgs(command)
	if err != nil {
		return nil, err
	}

	var res []map[string]string
	for _, re := range reply.Re {
		ans := make(map[string]string)
		for k, v := range re.Map {
			ans[k] = v
		}
		res = append(res, ans)
	}

	return res, nil
}
