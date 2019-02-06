package ros

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	//TODO: remove after firware update
	"github.com/ScriptRock/crypto/ssh"
	//TODO: add after firware update
	//"golang.org/x/crypto/ssh"
)

var re = regexp.MustCompile("[" + string([]byte{27, 91}) + "K]")

func clean(s string) string {
	return re.ReplaceAllString(s, "\n")
}

// Comand represents a basic CLI request.
type Command struct {
	Path    string            // base path to use
	Command string            // command to run
	UParam  *string           // possible unnamed parameter
	Params  map[string]string // possible named parameters
	Extra   map[string]string // extra unsearchable parameters
	Flags   map[string]bool   // possible boolean options to set
	Filter  map[string]string // possible recovery filter to apply
	Detail  bool              // is a detailed print required
	Once    bool              // run a loop once
}

var setTmpl = template.Must(template.New("set").Parse(
	`{{.Path}} set{{if .Filter}} [find{{range $k,$v := .Filter}}{{if ne "" $v}} {{$k}}="{{$v}}"{{end}}{{end}}{{range $k,$v := .Filter}}{{if eq "" $v}} {{$k}}{{end}}{{end}}]{{end}}{{if .UParam}} {{.UParam}}{{end}}{{range $k,$v := .Params}} {{$k}}="{{$v}}"{{end}}{{if .Flags}}{{range $k,$v := .Flags}} {{if not $v}}!{{end}}{{$k}}{{end}}{{end}}`))
var removeTmpl = template.Must(template.New("remove").Parse(
	`{{.Path}} remove{{if .Filter}} [find where{{range $k,$v := .Filter}} {{if eq $v ""}}!{{$k}}~""{{else}}{{$k}}="{{$v}}"{{end}}{{end}}{{range $k,$v := .Flags}} {{if not $v}}!{{end}}{{$k}}{{end}}]{{end}}{{if .UParam}} {{.UParam}}{{end}}`))
var addTmpl = template.Must(template.New("add").Parse(
	`{{if .Params}}:if ([:len [{{.Path}} find{{range $k,$v := .Params}} {{$k}}="{{$v}}"{{end}}]] = 0) do={{"{"}}{{end}}{{.Path}} add{{range $k,$v := .Params}} {{$k}}="{{$v}}"{{end}}{{range $k,$v := .Extra}} {{$k}}="{{$v}}"{{end}}{{if .Params}}{{"}"}}{{end}}`))
var printTmpl = template.Must(template.New("print").Parse(
	`{{.Path}} print{{if .Detail}} detail{{end}}{{if or .Filter .Flags}} where{{range $k,$v := .Filter}} {{$k}}="{{$v}}"{{end}}{{range $k,$v := .Flags}} {{if not $v}}!{{end}}{{$k}}{{end}}{{end}}`))
var exportTmpl = template.Must(template.New("export").Parse(
	`{{.Path}} export{{range $k,$v := .Filter}} {{$k}}="{{$v}}"{{end}}{{range $k,$v := .Flags}}{{if $v}} {{$k}}{{end}}{{end}}`))
var infoTmpl = template.Must(template.New("print").Parse(
	`{{.Path}} info {{if .UParam}} {{.UParam}}{{end}} {{if .Once}} once{{end}}`))

// /routing filter move numbers=2 destination=1

func (c Command) Parse() (string, error) {
	var res bytes.Buffer
	switch c.Command {
	case "add":
		if err := addTmpl.Execute(&res, c); err != nil {
			return "", err
		}
	case "remove":
		if err := removeTmpl.Execute(&res, c); err != nil {
			return "", err
		}
	case "set":
		if err := setTmpl.Execute(&res, c); err != nil {
			return "", err
		}
	case "print":
		if err := printTmpl.Execute(&res, c); err != nil {
			return "", err
		}
	case "export":
		if err := exportTmpl.Execute(&res, c); err != nil {
			return "", err
		}
	case "info":
		if err := infoTmpl.Execute(&res, c); err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("Unknown type: %s", c.Command)
	}
	return res.String(), nil
}

func (c Command) Run(client *ssh.Client) ([]string, error) {
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	command, err := c.Parse()
	if err != nil || command == "" {
		return nil, err
	}

	var lines []string

	stdout, err := session.StdoutPipe()
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(stdout)

	err = session.Run(string(command))
	if err != nil {
		return nil, err
	}
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "#") {
			continue
		}
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func (c Command) Exec(client *ssh.Client) error {
	res, err := c.Run(client)
	if err != nil {
		return err
	}
	if len(res) > 0 {
		p, err := c.Parse()
		if err != nil {
			return err
		}

		return fmt.Errorf("error: '%s': %s\n", p, strings.Join(res, ";"))
	}
	return nil
}

func (c Command) Values(client *ssh.Client) (map[string]string, error) {
	lines, err := c.Run(client)
	if err != nil || !(len(lines) > 0) {
		return nil, err
	}

	res, err := ScanItems(strings.Join(lines, "\n"))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c Command) List(client *ssh.Client) ([]map[string]string, error) {
	lines, err := c.Run(client)
	if err != nil || !(len(lines) > 0) {
		return nil, err
	}
	var trimmed []string
	for _, l := range lines {
		parts := strings.Fields(strings.TrimSpace(l))

		var fields []string
		for i := 0; i < len(parts); i++ {
			switch {
			// causes octal decoding errors?
			case strings.HasPrefix(parts[i], "last-link-up-time="):
				i++
			// causes octal decoding errors?
			case strings.HasPrefix(parts[i], "last-link-down-time="):
				i++
			default:
				fields = append(fields, parts[i])
			}
		}

		trimmed = append(trimmed, strings.Join(fields, " "))
	}

	list, err := ScanNumberedItemList(strings.Join(trimmed, "\n"))
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (c Command) Raw(client *ssh.Client) ([]map[string]string, error) {
	lines, err := c.Run(client)
	if err != nil || !(len(lines) > 0) {
		return nil, err
	}

	list, err := ScanNumberedItemList(strings.Join(lines, "\n"))
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (c Command) First(client *ssh.Client) (map[string]string, error) {
	list, err := c.List(client)
	if err != nil || !(len(list) > 0) {
		return nil, err
	}
	return list[0], nil
}

func (c Command) UnnumberedList(client *ssh.Client, offset int) ([]map[string]string, error) {
	lines, err := c.Run(client)
	if err != nil || !(len(lines) > 0) {
		return nil, err
	}
	var trimmed []string
	for n, l := range lines {
		parts := strings.Fields(strings.TrimSpace(l))

		var fields []string
		if !(n < offset) {
			fields = append(fields, strconv.Itoa(n-offset))
		}
		for i := 0; i < len(parts); i++ {
			switch {
			// causes octal decoding errors?
			case strings.HasPrefix(parts[i], "last-link-up-time="):
				i++
			// causes octal decoding errors?
			case strings.HasPrefix(parts[i], "last-link-down-time="):
				i++
			default:
				fields = append(fields, parts[i])
			}
		}

		trimmed = append(trimmed, strings.Join(fields, " "))
	}

	list, err := ScanNumberedItemList(strings.Join(trimmed, "\n"))
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (c Command) Info(client *ssh.Client, halt string) (map[string]string, error) {
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	command, err := c.Parse()
	if err != nil || command == "" {
		return nil, err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 40, 80, modes); err != nil {
		return nil, err
	}

	stdin, err := session.StdinPipe()
	if err != nil {
		return nil, err
	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(stdout)

	res := make(map[string]string)

	go func() {

		var key string
		for scanner.Scan() {
			if s := strings.Replace(clean(scanner.Text()), "-- [Q quit|D dump|C-z pause]", "", -1); len(s) > 0 {
				switch {
				case strings.Contains(s, ":"):
					if parts := strings.Split(s, ":"); len(parts) > 1 {
						if keys := strings.Fields(parts[0]); len(keys) > 0 {
							key = keys[len(keys)-1]
							res[key] = strings.TrimSpace(strings.Join(parts[1:], ":"))
						}
					}
				default:
					if key != "" {
						res[key] = res[key] + " " + strings.TrimSpace(s)
					}
				}
			}
			if strings.Contains(scanner.Text(), halt) {
				break
			}
		}
		_, _ = stdin.Write([]byte("\n"))
		session.Close()
	}()

	err = session.Run(command)
	if err != nil {
		return nil, err
	}

	return res, nil
}
