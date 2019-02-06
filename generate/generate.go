package main

import (
	"io"
	"strconv"
	"strings"
	"text/template"
)

const settingsTemplate = `
func Resources() []Menu {
	return []Menu{
{{range $b := .Menus}}
		Menu{
			Path: "{{$b.Path}}",
                	Params: []Param{
{{range $v := $b.Params -}}
				{
					Name: "{{$v.Name}}",
					{{if $v.Alias -}}
					Alias: "{{$v.Alias}}",
					{{end -}}
					Type: "{{$v.Type}}",
					Filter: {{$v.Filter}},
					{{if $v.Version }}Version: []int{ {{array ", " $v.Version}} },
					{{end -}}
					{{if $v.Deprecated }}Deprecated: []int{ {{array ", " $v.Deprecated}} },
					{{end -}}
					Optional: {{$v.Optional}},
					NotEmpty: {{$v.NotEmpty}},
					Default: "{{$v.Default}}",
					Extra: {{$v.Extra}},
					ReadOnly: {{$v.ReadOnly}},
					{{if $v.Script -}}
					StateFunc: "func(s string) string { return ros.ParseSystemScriptSource(s)}",
					DiffFunc: "func(before, after string) bool { return ros.ParseSystemScriptSource(before) == ros.PostSystemScriptSource(after)}",
					{{end -}}
				},
{{end -}}
                	},
{{if $b.Version }}Version: []int{ {{array ", " $b.Version}} },
{{end -}}
{{if $b.Deprecated }}Deprecated: []int{ {{array ", " $b.Deprecated}} },
{{end -}}
                	List:     {{$b.List}},
			{{if $b.Ordered }}Ordered:  "{{$b.Ordered}}",{{end}}
			{{if $b.Routerboard }}Routerboard:  {{$b.Routerboard}},{{end}}
                	ListOnly: {{$b.ListOnly}},
			ReadOnly: {{$b.ReadOnly}},
                	SetOnly:  {{$b.SetOnly}},
                	ShowOnly: {{$b.ShowOnly}},
			Default:  {{$b.Default}},
		},
{{end }}
	}
}
{{range $b := .Menus }}
{{if $b.List -}}
func {{lower $b.Path}}{{plural $b.Path}}() Command {
	return Command{
		Path:    "{{$b.Path}}",
		Command: "print",
                Flags: map[string]bool{
{{range $k, $v := (flags $b.ReadFlags) -}}
			"{{$k}}": {{$v}},
{{end -}}
		},
                Detail: {{$b.Detail}},
	}
}

func (r *Ros) {{upper $b.Path}}{{plural $b.Path}}() ([]map[string]string, bool, error) {
{{if $b.Routerboard}}if r.Routerboard() { {{end}}
{{if $b.Version}}if r.AtLeast({{array ", " $b.Version}}) { {{end}}
{{if $b.Deprecated}}if r.AtMost({{array ", " $b.Deprecated}}) { {{end -}}
	res, err := r.List({{lower $b.Path}}{{plural $b.Path}}())

	return res, true, err
{{if $b.Deprecated}} } {{end}}
{{if $b.Version}} } {{end}}
{{if $b.Routerboard}} } {{end}}
{{if or (or $b.Version $b.Deprecated) $b.Routerboard }}
	return nil, false, nil
{{end }}
}
{{end -}}
{{if not $b.ListOnly -}}
func {{lower $b.Path}}({{join "," (args $b.Filter)}}) Command {
	return Command{
		Path:    "{{$b.Path}}",
		Command: "print",
{{if $b.Mapped -}}
                Filter: {{$b.Mapped}},
{{else -}}
                Filter: map[string]string{
{{range $k, $v := (raw $b.Filter) -}}
                	"{{$k}}": {{$v}},
{{end -}}
                },
{{end -}}
                Flags: map[string]bool{
{{if $b.Default -}}
			"default": false,
{{end -}}
{{range $k, $v := (flags $b.ReadFlags) -}}
			"{{$k}}": {{$v}},
{{end -}}
		},
                Detail: {{$b.Detail}},
	}
}
func (r *Ros) {{upper $b.Path}}({{join "," (args $b.Filter)}}) (map[string]string, bool, error) {
{{if $b.Routerboard}}if r.Routerboard() { {{end}}
{{if $b.Version}}if r.AtLeast({{array ", " $b.Version}}) { {{end}}
{{if $b.Deprecated}}if r.AtMost({{array ", " $b.Deprecated}}) { {{end}}
{{if $b.Raw -}}
        raw, err := r.Raw({{lower $b.Path}}({{join "," (keys $b.Filter)}}))
        if err != nil {
                return nil, true, err
        }
        if len(raw) > 0 {
                return raw[0], true, err
        }
        return nil, true, nil
{{else -}}
{{if gt (len $b.Filter) 0 -}}
	res, err :=  r.First({{lower $b.Path}}({{join "," (keys $b.Filter)}}))
{{else -}}
	res, err := r.Values({{lower $b.Path}}())
{{end -}}
        return res, true, err
{{end -}}
{{if $b.Deprecated}} } {{end}}
{{if $b.Version}} } {{end}}
{{if $b.Routerboard}} } {{end}}
{{if or (or $b.Version $b.Deprecated) $b.Routerboard }}
	return nil, false, nil {{end -}}
}
{{if not $b.ReadOnly -}}
{{if not $b.SetOnly -}}
{{if gt (len $b.Filter) 0 -}}
func add{{upper $b.Path}}({{join "," (args $b.Filter $b.Required $b.Extra)}}) Command {
        return Command{
		Path:    "{{$b.Path}}",
                Command: "add",
{{if $b.Mapped -}}
                Params: {{$b.Mapped}},
{{else -}}
                Params: map[string]string{
{{range $k, $v := (raw $b.Filter $b.Required) -}}
                	"{{$k}}": {{$v}},
{{end -}}
                },
{{end -}}
                Extra: map[string]string{
{{range $k, $v := (raw $b.Extra) -}}
                        "{{$k}}": {{$v}},
{{end -}}
                },
        }
}
func (r *Ros) Add{{upper $b.Path}}({{join "," (args $b.Filter $b.Required $b.Extra)}}) error {
{{if $b.Routerboard}}if r.Routerboard() { {{end}}
{{if $b.Version}}if r.AtLeast({{array ", " $b.Version}}) { {{end}}
{{if $b.Deprecated}}if r.AtMost({{array ", " $b.Deprecated}}) { {{end}}
	return r.Exec(add{{upper $b.Path}}({{join "," (keys $b.Filter $b.Required $b.Extra)}}))
{{if $b.Deprecated}} } {{end}}
{{if $b.Version}} } {{end}}
{{if $b.Routerboard}} } {{end}}
{{if or (or $b.Version $b.Deprecated) $b.Routerboard }}
	return nil {{end}}
}
func remove{{upper $b.Path}}({{join "," (args $b.Filter)}}) Command {
        return Command{
                Path:    "{{$b.Path}}",
                Command: "remove",
{{if $b.Mapped -}}
                Filter: {{$b.Mapped}},
{{else -}}
                Filter: map[string]string{
{{range $k, $v := (raw $b.Filter) -}}
                "{{$k}}": {{$v}},
{{end -}}
                },
{{end -}}
                Flags: map[string]bool{
{{if $b.Default -}}
                        "default": false,
{{end -}}
                },
        }
}
func (r *Ros) Remove{{upper $b.Path}}({{join "," (args $b.Filter)}}) error {
{{if $b.Routerboard}}if r.Routerboard() { {{end}}
{{if $b.Version}}if r.AtLeast({{array ", " $b.Version}}) { {{end}}
{{if $b.Deprecated}}if r.AtMost({{array ", " $b.Deprecated}}) { {{end}}
	return r.Exec(remove{{upper $b.Path}}({{join "," (keys $b.Filter)}}))
{{if $b.Deprecated}} } {{end}}
{{if $b.Version}} } {{end}}
{{if $b.Routerboard}} } {{end}}
{{if or (or $b.Version $b.Deprecated) $b.Routerboard }}
	return nil {{end}}
}
{{end -}}
{{end -}}
func set{{upper $b.Path}}({{join "," (args $b.Filter)}}{{if gt (len $b.Filter) 0}},{{end}}key, value string) Command {
        return Command{
                Path:    "{{$b.Path}}",
                Command: "set",
{{if $b.Mapped -}}
                Filter: {{$b.Mapped}},
{{else -}}
                Filter: map[string]string{
{{range $k, $v := (raw $b.Filter) -}}
                        "{{$k}}": {{$v}},
{{end -}}
{{range $k, $v := $b.SetFilter -}}
		        "{{$k}}": "{{$v}}",
{{end -}}
{{if $b.Default -}}
		        "!default": "",
{{end -}}
                },
{{end -}}
                Params: map[string]string{
                        key: value,
                },
        }
}
{{if $b.HasNotEmpty -}}
func set{{upper $b.Path}}Off({{join "," (args $b.Filter)}}{{if gt (len $b.Filter) 0}},{{end}}key string) Command {
        return Command{
                Path:    "{{$b.Path}}",
                Command: "set",
{{if $b.Mapped -}}
                Filter: {{$b.Mapped}},
{{else -}}
                Filter: map[string]string{
{{range $k, $v := (raw $b.Filter) -}}
                        "{{$k}}": {{$v}},
{{end -}}
{{range $k, $v := $b.SetFilter -}}
		        "{{$k}}": "{{$v}}",
{{end -}}
{{if $b.Default -}}
		        "!default": "",
{{end -}}
                },
{{end -}}
                Flags: map[string]bool {
                       key: false,
              },
        }
}
{{end -}}
{{range $v := $b.Params -}}{{if (or $v.Optional $v.Extra) -}}
func set{{upper $b.Path}}{{upper $v.Name}}({{join "," (args $b.Filter)}}{{if gt (len $b.Filter) 0}},{{end}}value {{$v.Type}}) Command {
	return set{{upper $b.Path}}({{join "," (keys $b.Filter)}}{{if gt (len $b.Filter) 0}},{{end}}"{{$v.Name}}", {{convert $v "value"}})
}
func (r *Ros) Set{{upper $b.Path}}{{upper $v.Name}}({{join "," (args $b.Filter)}}{{if gt (len $b.Filter) 0}},{{end}}value {{$v.Type}}) error {
{{if $b.Routerboard}}if r.Routerboard() { {{end}}
{{if $b.Version}}if r.AtLeast({{array ", " $b.Version}}) { {{end}}
{{if $b.Deprecated}}if r.AtMost({{array ", " $b.Deprecated}}) { {{end}}
{{if $v.Version}}if r.AtLeast({{array ", " $v.Version}}) { {{end}}
{{if $v.Deprecated}}if r.AtMost({{array ", " $v.Deprecated}}) { {{end}}
{{if $v.Script -}}
        return r.Exec(set{{upper $b.Path}}{{upper $v.Name}}({{join "," (keys $b.Filter)}}{{if gt (len $b.Filter) 0}},{{end}}ParseSystemScriptSource(value)))
{{else -}}
	return r.Exec(set{{upper $b.Path}}{{upper $v.Name}}({{join "," (keys $b.Filter)}}{{if gt (len $b.Filter) 0}},{{end}}value))
{{end -}}
{{if $v.Deprecated}} } {{end}}
{{if $v.Version}} } {{end}}
{{if $b.Deprecated}} } {{end}}
{{if $b.Version}} } {{end}}
{{if $b.Routerboard}} } {{end}}
{{if or (append $v.Deprecated $v.Version $b.Version $b.Deprecated) $b.Routerboard }}
	return nil {{end}}
}
{{if $v.NotEmpty -}}
func set{{upper $b.Path}}{{upper $v.Name}}Off({{join "," (args $b.Filter)}}) Command {
	return set{{upper $b.Path}}Off({{join "," (keys $b.Filter)}}{{if gt (len $b.Filter) 0}},{{end}}"{{$v.Name}}")
}
func (r *Ros) Set{{upper $b.Path}}{{upper $v.Name}}Off({{join "," (args $b.Filter)}}) error {
{{if $b.Routerboard}}if r.Routerboard() { {{end}}
{{if $b.Version}}if r.AtLeast({{array ", " $b.Version}}) { {{end}}
{{if $b.Deprecated}}if r.AtMost({{array ", " $b.Deprecated}}) { {{end}}
{{if $v.Version}}if r.AtLeast({{array ", " $v.Version}}) { {{end}}
{{if $v.Deprecated}}if r.AtMost({{array ", " $v.Deprecated}}) { {{end}}
	return r.Exec(set{{upper $b.Path}}{{upper $v.Name}}Off({{join "," (keys $b.Filter)}}))
{{if $v.Deprecated}} } {{end}}
{{if $v.Version}} } {{end}}
{{if $b.Deprecated}} } {{end}}
{{if $b.Version}} } {{end}}
{{if $b.Routerboard}} } {{end}}
{{if or (append $v.Deprecated $v.Version $b.Version $b.Deprecated) $b.Routerboard }}
	return nil {{end -}}
}
{{end -}}
{{end}}{{end -}}
{{end -}}
{{if $b.Default -}}
func {{lower $b.Path}}Default({{join "," (args $b.Extra)}}) Command {
	return Command{
		Path:    "{{$b.Path}}",
		Command: "print",
{{if $b.Mapped -}}
                Filter: {{$b.Mapped}},
{{end -}}
                Flags: map[string]bool{
			"default": true,
{{range $k, $v := (flags $b.ReadFlags) -}}
			"{{$k}}": {{$v}},
{{end -}}
		},
                Detail: {{$b.Detail}},
	}
}
func (r *Ros) {{upper $b.Path}}Default({{join "," (args $b.Extra)}}) (map[string]string, bool, error) {
{{if $b.Routerboard}}if r.Routerboard() { {{end}}
{{if $b.Version}}if r.AtLeast({{array ", " $b.Version}}) { {{end}}
{{if $b.Deprecated}}if r.AtMost({{array ", " $b.Deprecated}}) { {{end}}
	res, err := r.Values({{lower $b.Path}}Default({{join "," (keys $b.Extra)}}))

	return res, true, err
{{if $b.Deprecated}} } {{end}}
{{if $b.Version}} } {{end}}
{{if $b.Routerboard}} } {{end}}
{{if or (or $b.Version $b.Deprecated) $b.Routerboard }}
	return nil, false, nil {{end}}
}
func set{{upper $b.Path}}Default({{join "," (args $b.Extra)}}{{if gt (len $b.Extra) 0}},{{end}}key, value string) Command {
        return Command{
                Path:    "{{$b.Path}}",
                Command: "set",
                Filter: map[string]string{
{{range $k, $v := $b.SetFilter -}}
		        "{{$k}}": "{{$v}}",
{{end -}}
		        "default": "",
                },
                Params: map[string]string{
                        key: value,
                },
        }
}
{{range $v := $b.Params -}}
func set{{upper $b.Path}}Default{{upper $v.Name}}({{join "," (args $b.Extra)}}{{if gt (len $b.Extra) 0}},{{end}}value {{$v.Type}}) Command {
	return set{{upper $b.Path}}Default({{join "," (keys $b.Extra)}}{{if gt (len $b.Extra) 0}},{{end}}"{{$v.Name}}", {{convert $v "value"}})
}
func (r *Ros) Set{{upper $b.Path}}Default{{upper $v.Name}}({{join "," (args $b.Extra)}}{{if gt (len $b.Extra) 0}},{{end}}value {{$v.Type}}) error {
{{if $b.Routerboard}}if r.Routerboard() { {{end}}
{{if $b.Version}}if r.AtLeast({{array ", " $b.Version}}) { {{end}}
{{if $b.Deprecated}}if r.AtMost({{array ", " $b.Deprecated}}) { {{end}}
{{if $v.Version}}if r.AtLeast({{array ", " $v.Version}}) { {{end}}
{{if $v.Deprecated}}if r.AtMost({{array ", " $v.Deprecated}}) { {{end}}
	return r.Exec(set{{upper $b.Path}}Default{{upper $v.Name}}({{join "," (keys $b.Extra)}}{{if gt (len $b.Extra) 0}},{{end}}value))
{{if $v.Deprecated}} } {{end}}
{{if $v.Version}} } {{end}}
{{if $b.Deprecated}} } {{end}}
{{if $b.Version}} } {{end}}
{{if $b.Routerboard}} } {{end}}
{{if or (append $v.Deprecated $v.Version $b.Version $b.Deprecated) $b.Routerboard }}
	return nil {{end}}
}
{{end -}}
{{end -}}
{{end -}}
{{end -}}
`

func (c Console) Generate(w io.Writer) error {

	t, err := template.New("generate").Funcs(
		template.FuncMap{
			"plural": func(s string) string {
				if strings.HasSuffix(strings.ToLower(s), "s") {
					return "es"
				}
				return "s"
			},
			"lower": func(s string) string {
				if u := strings.Join(strings.Fields(strings.Title(strings.Replace(strings.TrimLeft(s, "/"), "-", " ", -1))), ""); len(u) > 0 {
					switch s := strings.Replace(u, u[:1], strings.ToLower(u[:1]), 1); s {
					case "interface":
						return "iface"
					default:
						return s
					}
				}
				return ""
			},
			"upper": func(s string) string {
				return strings.Join(strings.Fields(strings.Title(strings.Replace(strings.TrimLeft(s, "/"), "-", " ", -1))), "")
			},
			"join": func(s string, r ...[]string) string {
				var args []string
				for _, a := range r {
					args = append(args, a...)
				}
				return strings.Join(args, s)
			},
			"append": func(r ...[]int) []int {
				var res []int
				for _, a := range r {
					res = append(res, a...)
				}
				return res
			},
			"array": func(s string, r ...[]int) string {
				var args []string
				for _, a := range r {
					for _, i := range a {
						args = append(args, strconv.Itoa(i))
					}
				}
				return strings.Join(args, s)
			},
			"convert": func(p Param, v string) string {
				switch p.Type {
				case "bool":
					return "FormatBool(" + v + ")"
				case "int":
					return "FormatInt(" + v + ")"
				default:
					return v
				}
			},
			"args": func(m ...[]string) []string {
				var args []string
				for _, r := range m {
					for _, s := range r {
						if parts := strings.Fields(s); len(parts) > 1 {
							if u := strings.Join(strings.Fields(strings.Title(strings.Replace(parts[0], "-", " ", -1))), ""); len(u) > 0 {
								if s := strings.Replace(u, u[:1], strings.ToLower(u[:1]), 1); len(s) > 0 {
									switch s {
									case "interface":
										args = append(args, "iface "+parts[1])
									default:
										args = append(args, s+" "+parts[1])
									}
								}
							}
						}
					}
				}
				return args
			},
			"flags": func(m ...map[string]bool) map[string]bool {
				flags := make(map[string]bool)
				for _, r := range m {
					for k, v := range r {
						flags[k] = v
					}
				}
				return flags
			},
			"raw": func(m ...[]string) map[string]string {
				raw := make(map[string]string)
				for _, r := range m {
					for _, s := range r {
						if parts := strings.Fields(s); len(parts) > 1 {
							if u := strings.Join(strings.Fields(strings.Title(strings.Replace(strings.TrimLeft(parts[0], "/"), "-", " ", -1))), ""); len(u) > 0 {
								v := func() string {
									switch s := strings.Replace(u, u[:1], strings.ToLower(u[:1]), 1); s {
									case "interface":
										return "iface"
									default:
										return s
									}
								}()
								switch parts[1] {
								case "bool":
									raw[parts[0]] = "FormatBool(" + v + ")"
								case "int":
									raw[parts[0]] = "FormatInt(" + v + ")"
								default:
									raw[parts[0]] = v
								}
							}
						}
					}
				}
				return raw
			},
			"keys": func(m ...[]string) []string {
				var keys []string
				for _, r := range m {
					for _, s := range r {
						if parts := strings.Fields(s); len(parts) > 1 {
							if u := strings.Join(strings.Fields(strings.Title(strings.Replace(parts[0], "-", " ", -1))), ""); len(u) > 0 {
								if s := strings.Replace(u, u[:1], strings.ToLower(u[:1]), 1); len(s) > 0 {
									switch s {
									case "interface":
										keys = append(keys, "iface")
									default:
										keys = append(keys, s)
									}
								}
							}
						}
					}
				}
				return keys
			},
		},
	).Parse(settingsTemplate)
	if err != nil {
		return err
	}
	if err := t.Execute(w, c); err != nil {
		return err
	}

	return nil
}
