package main

const prefix = `
package ros

import (
	"strings"
)

type Param struct {
        Name       string // setting name
        Alias      string // name override (usually for "id")
        Type       string // setting storage type
        Default    string // a default value if appropriate
	Version    []int  // the minimum major and minor version numbers
	Deprecated []int  // the maximum major and minor version numbers
        Optional   bool   // is this an optional parameter
	NotEmpty   bool   // optional parameter requires a value
        Filter     bool   // required parameter needed for read or update
        Extra      bool   // required parameter needed for add
        ReadOnly   bool   // value will only be listed, or generated

	StateFunc string // func(string) string      // helper function to provide a consistent listing
	DiffFunc  string // func(string, string) bool // helper function to identify a change
}

func (p Param) Title() string {
	return strings.Replace(strings.Title(strings.Replace(p.Name, "-", " ", -1)), " ", "", -1)
}

func (p Param) Id() string {
	if p.Alias != "" {
		return strings.Replace(p.Alias, "-", "_", -1)
	}
	return strings.Replace(p.Name, "-", "_", -1)
}

func (p Param) Label() string {
	return strings.Replace(p.Name, "-", "_", -1)
}

type Menu struct {
        Path        string  // the base ROS command prefix
        Params      []Param // command parameters
	Version     []int   // the minimum major and minor version numbers
	Deprecated  []int   // the maximum major and minor version numbers
	List        bool    // has a list capability
	Ordered     string  // expects the name of an ordered parameter list
	ListOnly    bool    // only provide a list capability
	ReadOnly    bool    // only read capability
	SetOnly     bool    // no add or remove capability
	Default     bool    // command has a "default" flag
	ShowOnly    bool    // command only shows ephemeral information
	Routerboard bool    // command only applies to routerboard hardware
}

func (m Menu) Id() string {
	return strings.Replace(strings.Replace(strings.TrimLeft(m.Path, "/"), "-", " ", -1), " ", "_", -1)
}

func (m Menu) Title() string {
	return strings.Replace(strings.Title(strings.TrimLeft(strings.Replace(m.Path, "-", " ", -1), "/")), " ", "", -1)
}

func (m Menu) HasCreate() bool {
        if m.SetOnly || m.ReadOnly {
                return false
        }

        for _, f := range m.Params {
                if f.Filter {
                        return true
                }
        }

        return false
}

func (m Menu) HasDelete() bool {
	return m.HasCreate()
}
`

type Param struct {
	Name       string `yaml:"name"`
	Alias      string `yaml:"alias"`
	Type       string `yaml:"type"`
	Default    string `yaml:"default"`
	Optional   bool   `yaml:"optional"`
	NotEmpty   bool   `yaml:"notempty"`
	Filter     bool   `yaml:"filter"`
	Extra      bool   `yaml:"extra"`
	ReadOnly   bool   `yaml:"readonly"`
	Script     bool   `yaml:"script"`
	Version    []int  `yaml:"version"`
	Deprecated []int  `yaml:"deprecated"`
}

// Menu represents a settings menu
type Menu struct {
	Path        string            `yaml:"path"`
	Params      []Param           `yaml:"params"`
	Detail      bool              `yaml:"detail"`
	Version     []int             `yaml:"version"`
	Deprecated  []int             `yaml:"deprecated"`
	List        bool              `yaml:"list"`
	Ordered     string            `yaml:"ordered"`
	ListOnly    bool              `yaml:"listonly"`
	ReadOnly    bool              `yaml:"readonly"`
	SetOnly     bool              `yaml:"setonly"`
	SetFilter   map[string]string `yaml:"setfilter"`
	ReadFlags   map[string]bool   `yaml:"readflags"`
	Mapped      string            `yaml:"mapped"`
	Raw         bool              `yaml:"raw"`
	Default     bool              `yaml:"default"`
	Routerboard bool              `yaml:"routerboard"`
	ShowOnly    bool              `yaml:"showonly"`
}

func (m Menu) Filter() []string {
	var filter []string
	for _, p := range m.Params {
		if !p.Filter || p.ReadOnly {
			continue
		}
		filter = append(filter, p.Name+" "+p.Type)
	}

	return filter
}

func (m Menu) Extra() []string {
	var extra []string
	for _, p := range m.Params {
		if !p.Extra || p.ReadOnly {
			continue
		}
		extra = append(extra, p.Name+" "+p.Type)
	}

	return extra
}

func (m Menu) Required() []string {
	var required []string
	for _, p := range m.Params {
		if p.Optional || p.ReadOnly {
			continue
		}
		if p.Extra || p.Filter {
			continue
		}
		required = append(required, p.Name+" "+p.Type)
	}

	return required
}

func (m Menu) Optional() []string {
	var optional []string
	for _, p := range m.Params {
		if !p.Optional || p.ReadOnly {
			continue
		}
		if p.Extra || p.Filter {
			continue
		}
		optional = append(optional, p.Name+" "+p.Type)
	}

	return optional
}

func (m Menu) HasNotEmpty() bool {
	for _, p := range m.Params {
		if p.NotEmpty {
			return true
		}
	}
	return false
}

// Console represents a set of Menus
type Console struct {
	Menus []Menu `yaml:"menus"`
}
