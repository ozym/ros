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

func Resources() []Menu {
	return []Menu{

		{
			Path:   "/interface",
			Params: []Param{},
			List:   true,

			ListOnly: true,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/interface bridge",
			Params: []Param{
				{
					Name:     "name",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "protocol-mode",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "priority",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: true,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  true,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/interface ethernet",
			Params: []Param{
				{
					Name:     "name",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "mtu",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "auto",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: true,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  true,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path:   "/interface bridge host",
			Params: []Param{},
			List:   true,

			ListOnly: true,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: true,
			Default:  false,
		},

		{
			Path: "/interface bridge port",
			Params: []Param{
				{
					Name:     "bridge",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "interface",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: true,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/interface gre",
			Params: []Param{
				{
					Name:     "remote-address",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "name",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "mtu",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "auto",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "keepalive",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "clamp-tcp-mss",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "true",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "allow-fast-path",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "true",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: true,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/interface list",
			Params: []Param{
				{
					Name:     "name",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/interface list member",
			Params: []Param{
				{
					Name:     "interface",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "list",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "disabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/interface wireless",
			Params: []Param{
				{
					Name:     "name",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: true,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  true,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path:   "/interface wireless monitor",
			Params: []Param{},
			List:   true,

			ListOnly: true,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: true,
			Default:  false,
		},

		{
			Path:   "/ip arp",
			Params: []Param{},
			List:   true,

			ListOnly: true,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: true,
			Default:  false,
		},

		{
			Path: "/ip address",
			Params: []Param{
				{
					Name:     "address",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: true,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  true,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/ip cloud",
			Params: []Param{
				{
					Name:     "ddns-enabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "update-time",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
			},
			Version: []int{6, 27},
			List:    false,

			Routerboard: true,
			ListOnly:    false,
			ReadOnly:    false,
			SetOnly:     false,
			ShowOnly:    false,
			Default:     false,
		},

		{
			Path: "/ip dns",
			Params: []Param{
				{
					Name:     "servers",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "allow-remote-requests",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/ip dns static",
			Params: []Param{
				{
					Name:     "name",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "address",
					Type:     "string",
					Filter:   false,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    true,
					ReadOnly: false,
				},
				{
					Name:     "ttl",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "1d",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "disabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/ip firewall address-list",
			Params: []Param{
				{
					Name:     "address",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "list",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "disabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/ip firewall nat",
			Params: []Param{
				{
					Name:     "filter",
					Type:     "map[string]string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "disabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path:   "/ip neighbor",
			Params: []Param{},
			List:   true,

			ListOnly: true,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: true,
			Default:  false,
		},

		{
			Path: "/ip service",
			Params: []Param{
				{
					Name:     "name",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "disabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "port",
					Type:     "int",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "0",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "address",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/ip route",
			Params: []Param{
				{
					Name:     "dst-address",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: true,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  true,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/ip socks",
			Params: []Param{
				{
					Name:     "enabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "port",
					Type:     "int",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "1080",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "connection-idle-timeout",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "2m",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "max-connections",
					Type:     "int",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "200",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/ip smb",
			Params: []Param{
				{
					Name:     "enabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "domain",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "MSHOME",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "MikrotikSMB",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "allow-guests",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "true",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "interface",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "all",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/ip smb shares",
			Params: []Param{
				{
					Name:     "name",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "directory",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "/pub",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "default share",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "disabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "true",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "max-sessions",
					Type:     "int",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "10",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: true,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  true,
		},

		{
			Path: "/ip smb users",
			Params: []Param{
				{
					Name:     "name",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "password",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "read-only",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "true",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "disabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: true,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  true,
		},

		{
			Path: "/ip ssh",
			Params: []Param{
				{
					Name:     "strong-crypto",
					Type:     "bool",
					Filter:   false,
					Version:  []int{6, 36},
					Optional: true,
					NotEmpty: false,
					Default:  "true",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "forwarding-enabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "always-allow-password-login",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "host-key-size",
					Type:     "int",
					Filter:   false,
					Version:  []int{6, 36},
					Optional: true,
					NotEmpty: false,
					Default:  "2048",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/ip traffic-flow",
			Params: []Param{
				{
					Name:     "enabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "interfaces",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "all",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "cache-entries",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "active-flow-timeout",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "60s",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "inactive-flow-timeout",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "60s",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/ip traffic-flow target",
			Params: []Param{
				{
					Name:     "dst-address",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "port",
					Type:     "int",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "src-address",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "version",
					Type:     "int",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "9",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "v9-template-refresh",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "v9-template-timeout",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			Version: []int{6, 31},
			List:    false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/queue simple",
			Params: []Param{
				{
					Name:     "name",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "target",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "disabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "time",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "dst-address",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "packet-marks",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "parent",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "none",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "priority",
					Type:     "int",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "1",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "queue",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "limit-at",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "max-limit",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "burst-limit",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "burst-time",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "burst-threshold",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "total-queue",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "total-limit-at",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "total-max-limit",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "total-burst-limit",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "total-burst-time",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "total-burst-threshold",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "bucket-size",
					Type:     "string",
					Filter:   false,
					Version:  []int{6, 35},
					Optional: true,
					NotEmpty: false,
					Default:  "0.1",
					Extra:    false,
					ReadOnly: false,
				},
			},
			Version: []int{6, 0},
			List:    true,
			Ordered: "queue",

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/queue tree",
			Params: []Param{
				{
					Name:     "name",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "packet-marks",
					Type:     "string",
					Filter:   false,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    true,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "disabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "parent",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "none",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "priority",
					Type:     "int",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "1",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "queue",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "limit-at",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "max-limit",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "burst-limit",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "burst-time",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "burst-threshold",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "bucket-size",
					Type:     "string",
					Filter:   false,
					Version:  []int{6, 35},
					Optional: true,
					NotEmpty: false,
					Default:  "0.1",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: true,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/radius",
			Params: []Param{
				{
					Name:     "service",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "secret",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "address",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "src-address",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "timeout",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "disabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/radius incoming",
			Params: []Param{
				{
					Name:     "accept",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "port",
					Type:     "int",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "0",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/routing bgp network",
			Params: []Param{
				{
					Name:     "address",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: true,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  true,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/routing bgp aggregate",
			Params: []Param{
				{
					Name:     "instance",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "prefix",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: true,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  true,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/routing bgp instance",
			Params: []Param{
				{
					Name:     "name",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "router-id",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: true,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  true,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/routing bgp peer",
			Params: []Param{
				{
					Name:     "remote-address",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "name",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: true,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/routing filter",
			Params: []Param{
				{
					Name:     "comment",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "action",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "chain",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "prefix",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: true,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "prefix-length",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: true,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "protocol",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: true,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "invert-match",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "set-bgp-prepend-path",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List:    true,
			Ordered: "rule",

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/routing ospf instance",
			Params: []Param{
				{
					Name:     "name",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "router-id",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "redistribute-connected",
					Type:     "string",
					Filter:   false,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: true,
				},
				{
					Name:     "redistribute-static",
					Type:     "string",
					Filter:   false,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: true,
				},
				{
					Name:     "disabled",
					Type:     "bool",
					Filter:   false,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: true,
				},
			},
			List: true,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  true,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/routing ospf interface",
			Params: []Param{
				{
					Name:     "interface",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "cost",
					Type:     "int",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "100",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "network-type",
					Type:     "string",
					Filter:   false,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: true,
				},
				{
					Name:     "hello-interval",
					Type:     "string",
					Filter:   false,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: true,
				},
				{
					Name:     "dead-interval",
					Type:     "string",
					Filter:   false,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: true,
				},
				{
					Name:     "authentication",
					Type:     "string",
					Filter:   false,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: true,
				},
				{
					Name:     "authentication-key",
					Type:     "string",
					Filter:   false,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: true,
				},
				{
					Name:     "disabled",
					Type:     "bool",
					Filter:   false,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: true,
				},
			},
			List: true,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  true,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/routing ospf network",
			Params: []Param{
				{
					Name:     "network",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: true,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  true,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/routing ospf nbma-neighbor",
			Params: []Param{
				{
					Name:     "address",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "priority",
					Type:     "int",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "1",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "poll-interval",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "5s",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: true,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/snmp",
			Params: []Param{
				{
					Name:     "enabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "location",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "contact",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "engine-id",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "trap-community",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "public",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "trap-generators",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "trap-target",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "trap-version",
					Type:     "int",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "1",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/snmp community",
			Params: []Param{
				{
					Name:     "name",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "public",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "authentication-password",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "authentication-protocol",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "MD5",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "encryption-password",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "encryption-protocol",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "DES",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "addresses",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "security",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "none",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "read-access",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "true",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "write-access",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  true,
		},

		{
			Path: "/system clock",
			Params: []Param{
				{
					Name:     "time-zone-name",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "Pacific/Auckland",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "time-zone-autodetect",
					Type:     "bool",
					Filter:   false,
					Version:  []int{6, 27},
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/system identity",
			Params: []Param{
				{
					Name:     "name",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/system logging",
			Params: []Param{
				{
					Name:     "action",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "topics",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "prefix",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/system logging action",
			Params: []Param{
				{
					Name:     "name",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "target",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "remote",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "remote",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "remote-port",
					Type:     "int",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "0",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "src-address",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "bsd-syslog",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "true",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "syslog-severity",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "auto",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "syslog-facility",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "daemon",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/system note",
			Params: []Param{
				{
					Name:     "note",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "show-at-login",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/system ntp client",
			Params: []Param{
				{
					Name:     "enabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "primary-ntp",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "secondary-ntp",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/system resource",
			Params: []Param{
				{
					Name:     "version",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "cpu",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "architecture-name",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "board-name",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "platform",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: true,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/user",
			Params: []Param{
				{
					Name:     "name",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "group",
					Type:     "string",
					Filter:   false,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    true,
					ReadOnly: false,
				},
				{
					Name:     "password",
					Type:     "string",
					Filter:   false,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    true,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "address",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path:   "/system routerboard",
			Params: []Param{},
			List:   false,

			ListOnly: false,
			ReadOnly: true,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/system scheduler",
			Params: []Param{
				{
					Name:     "name",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "interval",
					Type:     "string",
					Filter:   false,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    true,
					ReadOnly: false,
				},
				{
					Name:     "policy",
					Type:     "string",
					Filter:   false,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    true,
					ReadOnly: false,
				},
				{
					Name:     "on-event",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    true,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "start-date",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "start-time",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "startup",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/system script",
			Params: []Param{
				{
					Name:     "name",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "owner",
					Type:     "string",
					Filter:   false,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    true,
					ReadOnly: false,
				},
				{
					Name:     "policy",
					Type:     "string",
					Filter:   false,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    true,
					ReadOnly: false,
				},
				{
					Name:      "source",
					Type:      "string",
					Filter:    false,
					Optional:  false,
					NotEmpty:  false,
					Default:   "",
					Extra:     true,
					ReadOnly:  false,
					StateFunc: "func(s string) string { return ros.ParseSystemScriptSource(s)}",
					DiffFunc:  "func(before, after string) bool { return ros.ParseSystemScriptSource(before) == ros.PostSystemScriptSource(after)}",
				},
			},
			Version: []int{6, 41},
			List:    false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/tool bandwidth-server",
			Params: []Param{
				{
					Name:     "enabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "true",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "authenticate",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "true",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "allocate-udp-ports-from",
					Type:     "int",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "2000",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "max-sessions",
					Type:     "int",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "100",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/tool mac-server",
			Params: []Param{
				{
					Name:     "allowed-interface-list",
					Type:     "string",
					Filter:   false,
					Version:  []int{6, 41},
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/tool mac-server mac-winbox",
			Params: []Param{
				{
					Name:     "allowed-interface-list",
					Type:     "string",
					Filter:   false,
					Version:  []int{6, 41},
					Optional: true,
					NotEmpty: false,
					Default:  "all",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/tool mac-server ping",
			Params: []Param{
				{
					Name:     "enabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "true",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/tool netwatch",
			Params: []Param{
				{
					Name:     "host",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "up-script",
					Type:     "string",
					Filter:   false,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    true,
					ReadOnly: false,
				},
				{
					Name:     "down-script",
					Type:     "string",
					Filter:   false,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    true,
					ReadOnly: false,
				},
				{
					Name:     "interval",
					Type:     "string",
					Filter:   false,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    true,
					ReadOnly: false,
				},
				{
					Name:     "timeout",
					Type:     "string",
					Filter:   false,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    true,
					ReadOnly: false,
				},
				{
					Name:     "comment",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "disabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/tool romon",
			Params: []Param{
				{
					Name:     "enabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "id",
					Alias:    "mac-id",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "00:00:00:00:00:00",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "secrets",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			Version: []int{6, 29},
			List:    false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},

		{
			Path: "/tool romon port",
			Params: []Param{
				{
					Name:     "interface",
					Type:     "string",
					Filter:   true,
					Optional: false,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "secrets",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "cost",
					Type:     "int",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "0",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "disabled",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "forbid",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
			},
			Version: []int{6, 29},
			List:    false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  true,
		},

		{
			Path: "/user aaa",
			Params: []Param{
				{
					Name:     "use-radius",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "false",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "accounting",
					Type:     "bool",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "true",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "interim-update",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "0s",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "default-group",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "read",
					Extra:    false,
					ReadOnly: false,
				},
				{
					Name:     "exclude-groups",
					Type:     "string",
					Filter:   false,
					Optional: true,
					NotEmpty: false,
					Default:  "",
					Extra:    false,
					ReadOnly: false,
				},
			},
			List: false,

			ListOnly: false,
			ReadOnly: false,
			SetOnly:  false,
			ShowOnly: false,
			Default:  false,
		},
	}
}

func ifaces() Command {
	return Command{
		Path:    "/interface",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) Interfaces() ([]map[string]string, bool, error) {

	res, err := r.List(ifaces())

	return res, true, err

}

func interfaceBridges() Command {
	return Command{
		Path:    "/interface bridge",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) InterfaceBridges() ([]map[string]string, bool, error) {

	res, err := r.List(interfaceBridges())

	return res, true, err

}
func interfaceBridge(name string) Command {
	return Command{
		Path:    "/interface bridge",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) InterfaceBridge(name string) (map[string]string, bool, error) {

	res, err := r.First(interfaceBridge(name))
	return res, true, err

}
func setInterfaceBridge(name string, key, value string) Command {
	return Command{
		Path:    "/interface bridge",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setInterfaceBridgeComment(name string, value string) Command {
	return setInterfaceBridge(name, "comment", value)
}
func (r *Ros) SetInterfaceBridgeComment(name string, value string) error {

	return r.Exec(setInterfaceBridgeComment(name, value))

}
func setInterfaceBridgeProtocolMode(name string, value string) Command {
	return setInterfaceBridge(name, "protocol-mode", value)
}
func (r *Ros) SetInterfaceBridgeProtocolMode(name string, value string) error {

	return r.Exec(setInterfaceBridgeProtocolMode(name, value))

}
func setInterfaceBridgePriority(name string, value string) Command {
	return setInterfaceBridge(name, "priority", value)
}
func (r *Ros) SetInterfaceBridgePriority(name string, value string) error {

	return r.Exec(setInterfaceBridgePriority(name, value))

}

func interfaceEthernets() Command {
	return Command{
		Path:    "/interface ethernet",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) InterfaceEthernets() ([]map[string]string, bool, error) {

	res, err := r.List(interfaceEthernets())

	return res, true, err

}
func interfaceEthernet(name string) Command {
	return Command{
		Path:    "/interface ethernet",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) InterfaceEthernet(name string) (map[string]string, bool, error) {

	res, err := r.First(interfaceEthernet(name))
	return res, true, err

}
func setInterfaceEthernet(name string, key, value string) Command {
	return Command{
		Path:    "/interface ethernet",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setInterfaceEthernetComment(name string, value string) Command {
	return setInterfaceEthernet(name, "comment", value)
}
func (r *Ros) SetInterfaceEthernetComment(name string, value string) error {

	return r.Exec(setInterfaceEthernetComment(name, value))

}
func setInterfaceEthernetMtu(name string, value string) Command {
	return setInterfaceEthernet(name, "mtu", value)
}
func (r *Ros) SetInterfaceEthernetMtu(name string, value string) error {

	return r.Exec(setInterfaceEthernetMtu(name, value))

}

func interfaceBridgeHosts() Command {
	return Command{
		Path:    "/interface bridge host",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) InterfaceBridgeHosts() ([]map[string]string, bool, error) {

	res, err := r.List(interfaceBridgeHosts())

	return res, true, err

}

func interfaceBridgePorts() Command {
	return Command{
		Path:    "/interface bridge port",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) InterfaceBridgePorts() ([]map[string]string, bool, error) {

	res, err := r.List(interfaceBridgePorts())

	return res, true, err

}
func interfaceBridgePort(bridge string, iface string) Command {
	return Command{
		Path:    "/interface bridge port",
		Command: "print",
		Filter: map[string]string{
			"bridge":    bridge,
			"interface": iface,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) InterfaceBridgePort(bridge string, iface string) (map[string]string, bool, error) {

	res, err := r.First(interfaceBridgePort(bridge, iface))
	return res, true, err

}
func addInterfaceBridgePort(bridge string, iface string) Command {
	return Command{
		Path:    "/interface bridge port",
		Command: "add",
		Params: map[string]string{
			"bridge":    bridge,
			"interface": iface,
		},
		Extra: map[string]string{},
	}
}
func (r *Ros) AddInterfaceBridgePort(bridge string, iface string) error {

	return r.Exec(addInterfaceBridgePort(bridge, iface))

}
func removeInterfaceBridgePort(bridge string, iface string) Command {
	return Command{
		Path:    "/interface bridge port",
		Command: "remove",
		Filter: map[string]string{
			"bridge":    bridge,
			"interface": iface,
		},
		Flags: map[string]bool{},
	}
}
func (r *Ros) RemoveInterfaceBridgePort(bridge string, iface string) error {

	return r.Exec(removeInterfaceBridgePort(bridge, iface))

}
func setInterfaceBridgePort(bridge string, iface string, key, value string) Command {
	return Command{
		Path:    "/interface bridge port",
		Command: "set",
		Filter: map[string]string{
			"bridge":    bridge,
			"interface": iface,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setInterfaceBridgePortComment(bridge string, iface string, value string) Command {
	return setInterfaceBridgePort(bridge, iface, "comment", value)
}
func (r *Ros) SetInterfaceBridgePortComment(bridge string, iface string, value string) error {

	return r.Exec(setInterfaceBridgePortComment(bridge, iface, value))

}

func interfaceGres() Command {
	return Command{
		Path:    "/interface gre",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) InterfaceGres() ([]map[string]string, bool, error) {

	res, err := r.List(interfaceGres())

	return res, true, err

}
func interfaceGre(remoteAddress string) Command {
	return Command{
		Path:    "/interface gre",
		Command: "print",
		Filter: map[string]string{
			"remote-address": remoteAddress,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) InterfaceGre(remoteAddress string) (map[string]string, bool, error) {

	res, err := r.First(interfaceGre(remoteAddress))
	return res, true, err

}
func addInterfaceGre(remoteAddress string) Command {
	return Command{
		Path:    "/interface gre",
		Command: "add",
		Params: map[string]string{
			"remote-address": remoteAddress,
		},
		Extra: map[string]string{},
	}
}
func (r *Ros) AddInterfaceGre(remoteAddress string) error {

	return r.Exec(addInterfaceGre(remoteAddress))

}
func removeInterfaceGre(remoteAddress string) Command {
	return Command{
		Path:    "/interface gre",
		Command: "remove",
		Filter: map[string]string{
			"remote-address": remoteAddress,
		},
		Flags: map[string]bool{},
	}
}
func (r *Ros) RemoveInterfaceGre(remoteAddress string) error {

	return r.Exec(removeInterfaceGre(remoteAddress))

}
func setInterfaceGre(remoteAddress string, key, value string) Command {
	return Command{
		Path:    "/interface gre",
		Command: "set",
		Filter: map[string]string{
			"remote-address": remoteAddress,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setInterfaceGreName(remoteAddress string, value string) Command {
	return setInterfaceGre(remoteAddress, "name", value)
}
func (r *Ros) SetInterfaceGreName(remoteAddress string, value string) error {

	return r.Exec(setInterfaceGreName(remoteAddress, value))

}
func setInterfaceGreComment(remoteAddress string, value string) Command {
	return setInterfaceGre(remoteAddress, "comment", value)
}
func (r *Ros) SetInterfaceGreComment(remoteAddress string, value string) error {

	return r.Exec(setInterfaceGreComment(remoteAddress, value))

}
func setInterfaceGreMtu(remoteAddress string, value string) Command {
	return setInterfaceGre(remoteAddress, "mtu", value)
}
func (r *Ros) SetInterfaceGreMtu(remoteAddress string, value string) error {

	return r.Exec(setInterfaceGreMtu(remoteAddress, value))

}
func setInterfaceGreKeepalive(remoteAddress string, value string) Command {
	return setInterfaceGre(remoteAddress, "keepalive", value)
}
func (r *Ros) SetInterfaceGreKeepalive(remoteAddress string, value string) error {

	return r.Exec(setInterfaceGreKeepalive(remoteAddress, value))

}
func setInterfaceGreClampTcpMss(remoteAddress string, value bool) Command {
	return setInterfaceGre(remoteAddress, "clamp-tcp-mss", FormatBool(value))
}
func (r *Ros) SetInterfaceGreClampTcpMss(remoteAddress string, value bool) error {

	return r.Exec(setInterfaceGreClampTcpMss(remoteAddress, value))

}
func setInterfaceGreAllowFastPath(remoteAddress string, value bool) Command {
	return setInterfaceGre(remoteAddress, "allow-fast-path", FormatBool(value))
}
func (r *Ros) SetInterfaceGreAllowFastPath(remoteAddress string, value bool) error {

	return r.Exec(setInterfaceGreAllowFastPath(remoteAddress, value))

}

func interfaceList(name string) Command {
	return Command{
		Path:    "/interface list",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) InterfaceList(name string) (map[string]string, bool, error) {

	res, err := r.First(interfaceList(name))
	return res, true, err

}
func addInterfaceList(name string) Command {
	return Command{
		Path:    "/interface list",
		Command: "add",
		Params: map[string]string{
			"name": name,
		},
		Extra: map[string]string{},
	}
}
func (r *Ros) AddInterfaceList(name string) error {

	return r.Exec(addInterfaceList(name))

}
func removeInterfaceList(name string) Command {
	return Command{
		Path:    "/interface list",
		Command: "remove",
		Filter: map[string]string{
			"name": name,
		},
		Flags: map[string]bool{},
	}
}
func (r *Ros) RemoveInterfaceList(name string) error {

	return r.Exec(removeInterfaceList(name))

}
func setInterfaceList(name string, key, value string) Command {
	return Command{
		Path:    "/interface list",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setInterfaceListComment(name string, value string) Command {
	return setInterfaceList(name, "comment", value)
}
func (r *Ros) SetInterfaceListComment(name string, value string) error {

	return r.Exec(setInterfaceListComment(name, value))

}

func interfaceListMember(iface string, list string) Command {
	return Command{
		Path:    "/interface list member",
		Command: "print",
		Filter: map[string]string{
			"interface": iface,
			"list":      list,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) InterfaceListMember(iface string, list string) (map[string]string, bool, error) {

	res, err := r.First(interfaceListMember(iface, list))
	return res, true, err

}
func addInterfaceListMember(iface string, list string) Command {
	return Command{
		Path:    "/interface list member",
		Command: "add",
		Params: map[string]string{
			"interface": iface,
			"list":      list,
		},
		Extra: map[string]string{},
	}
}
func (r *Ros) AddInterfaceListMember(iface string, list string) error {

	return r.Exec(addInterfaceListMember(iface, list))

}
func removeInterfaceListMember(iface string, list string) Command {
	return Command{
		Path:    "/interface list member",
		Command: "remove",
		Filter: map[string]string{
			"interface": iface,
			"list":      list,
		},
		Flags: map[string]bool{},
	}
}
func (r *Ros) RemoveInterfaceListMember(iface string, list string) error {

	return r.Exec(removeInterfaceListMember(iface, list))

}
func setInterfaceListMember(iface string, list string, key, value string) Command {
	return Command{
		Path:    "/interface list member",
		Command: "set",
		Filter: map[string]string{
			"interface": iface,
			"list":      list,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setInterfaceListMemberComment(iface string, list string, value string) Command {
	return setInterfaceListMember(iface, list, "comment", value)
}
func (r *Ros) SetInterfaceListMemberComment(iface string, list string, value string) error {

	return r.Exec(setInterfaceListMemberComment(iface, list, value))

}
func setInterfaceListMemberDisabled(iface string, list string, value bool) Command {
	return setInterfaceListMember(iface, list, "disabled", FormatBool(value))
}
func (r *Ros) SetInterfaceListMemberDisabled(iface string, list string, value bool) error {

	return r.Exec(setInterfaceListMemberDisabled(iface, list, value))

}

func interfaceWirelesses() Command {
	return Command{
		Path:    "/interface wireless",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) InterfaceWirelesses() ([]map[string]string, bool, error) {

	res, err := r.List(interfaceWirelesses())

	return res, true, err

}
func interfaceWireless(name string) Command {
	return Command{
		Path:    "/interface wireless",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) InterfaceWireless(name string) (map[string]string, bool, error) {

	res, err := r.First(interfaceWireless(name))
	return res, true, err

}
func setInterfaceWireless(name string, key, value string) Command {
	return Command{
		Path:    "/interface wireless",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setInterfaceWirelessComment(name string, value string) Command {
	return setInterfaceWireless(name, "comment", value)
}
func (r *Ros) SetInterfaceWirelessComment(name string, value string) error {

	return r.Exec(setInterfaceWirelessComment(name, value))

}

func interfaceWirelessMonitors() Command {
	return Command{
		Path:    "/interface wireless monitor",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) InterfaceWirelessMonitors() ([]map[string]string, bool, error) {

	res, err := r.List(interfaceWirelessMonitors())

	return res, true, err

}

func ipArps() Command {
	return Command{
		Path:    "/ip arp",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) IpArps() ([]map[string]string, bool, error) {

	res, err := r.List(ipArps())

	return res, true, err

}

func ipAddresses() Command {
	return Command{
		Path:    "/ip address",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) IpAddresses() ([]map[string]string, bool, error) {

	res, err := r.List(ipAddresses())

	return res, true, err

}
func ipAddress(address string) Command {
	return Command{
		Path:    "/ip address",
		Command: "print",
		Filter: map[string]string{
			"address": address,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) IpAddress(address string) (map[string]string, bool, error) {

	res, err := r.First(ipAddress(address))
	return res, true, err

}
func setIpAddress(address string, key, value string) Command {
	return Command{
		Path:    "/ip address",
		Command: "set",
		Filter: map[string]string{
			"address": address,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setIpAddressComment(address string, value string) Command {
	return setIpAddress(address, "comment", value)
}
func (r *Ros) SetIpAddressComment(address string, value string) error {

	return r.Exec(setIpAddressComment(address, value))

}

func ipCloud() Command {
	return Command{
		Path:    "/ip cloud",
		Command: "print",
		Filter:  map[string]string{},
		Flags:   map[string]bool{},
		Detail:  false,
	}
}
func (r *Ros) IpCloud() (map[string]string, bool, error) {
	if r.Routerboard() {
		if r.AtLeast(6, 27) {

			res, err := r.Values(ipCloud())
			return res, true, err

		}
	}

	return nil, false, nil
}
func setIpCloud(key, value string) Command {
	return Command{
		Path:    "/ip cloud",
		Command: "set",
		Filter:  map[string]string{},
		Params: map[string]string{
			key: value,
		},
	}
}
func setIpCloudDdnsEnabled(value bool) Command {
	return setIpCloud("ddns-enabled", FormatBool(value))
}
func (r *Ros) SetIpCloudDdnsEnabled(value bool) error {
	if r.Routerboard() {
		if r.AtLeast(6, 27) {

			return r.Exec(setIpCloudDdnsEnabled(value))

		}
	}

	return nil
}
func setIpCloudUpdateTime(value bool) Command {
	return setIpCloud("update-time", FormatBool(value))
}
func (r *Ros) SetIpCloudUpdateTime(value bool) error {
	if r.Routerboard() {
		if r.AtLeast(6, 27) {

			return r.Exec(setIpCloudUpdateTime(value))

		}
	}

	return nil
}

func ipDns() Command {
	return Command{
		Path:    "/ip dns",
		Command: "print",
		Filter:  map[string]string{},
		Flags:   map[string]bool{},
		Detail:  false,
	}
}
func (r *Ros) IpDns() (map[string]string, bool, error) {

	res, err := r.Values(ipDns())
	return res, true, err

}
func setIpDns(key, value string) Command {
	return Command{
		Path:    "/ip dns",
		Command: "set",
		Filter:  map[string]string{},
		Params: map[string]string{
			key: value,
		},
	}
}
func setIpDnsServers(value string) Command {
	return setIpDns("servers", value)
}
func (r *Ros) SetIpDnsServers(value string) error {

	return r.Exec(setIpDnsServers(value))

}
func setIpDnsAllowRemoteRequests(value bool) Command {
	return setIpDns("allow-remote-requests", FormatBool(value))
}
func (r *Ros) SetIpDnsAllowRemoteRequests(value bool) error {

	return r.Exec(setIpDnsAllowRemoteRequests(value))

}

func ipDnsStatic(name string) Command {
	return Command{
		Path:    "/ip dns static",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) IpDnsStatic(name string) (map[string]string, bool, error) {

	res, err := r.First(ipDnsStatic(name))
	return res, true, err

}
func addIpDnsStatic(name string, address string) Command {
	return Command{
		Path:    "/ip dns static",
		Command: "add",
		Params: map[string]string{
			"name": name,
		},
		Extra: map[string]string{
			"address": address,
		},
	}
}
func (r *Ros) AddIpDnsStatic(name string, address string) error {

	return r.Exec(addIpDnsStatic(name, address))

}
func removeIpDnsStatic(name string) Command {
	return Command{
		Path:    "/ip dns static",
		Command: "remove",
		Filter: map[string]string{
			"name": name,
		},
		Flags: map[string]bool{},
	}
}
func (r *Ros) RemoveIpDnsStatic(name string) error {

	return r.Exec(removeIpDnsStatic(name))

}
func setIpDnsStatic(name string, key, value string) Command {
	return Command{
		Path:    "/ip dns static",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setIpDnsStaticAddress(name string, value string) Command {
	return setIpDnsStatic(name, "address", value)
}
func (r *Ros) SetIpDnsStaticAddress(name string, value string) error {

	return r.Exec(setIpDnsStaticAddress(name, value))

}
func setIpDnsStaticTtl(name string, value string) Command {
	return setIpDnsStatic(name, "ttl", value)
}
func (r *Ros) SetIpDnsStaticTtl(name string, value string) error {

	return r.Exec(setIpDnsStaticTtl(name, value))

}
func setIpDnsStaticComment(name string, value string) Command {
	return setIpDnsStatic(name, "comment", value)
}
func (r *Ros) SetIpDnsStaticComment(name string, value string) error {

	return r.Exec(setIpDnsStaticComment(name, value))

}
func setIpDnsStaticDisabled(name string, value bool) Command {
	return setIpDnsStatic(name, "disabled", FormatBool(value))
}
func (r *Ros) SetIpDnsStaticDisabled(name string, value bool) error {

	return r.Exec(setIpDnsStaticDisabled(name, value))

}

func ipFirewallAddressList(address string, list string) Command {
	return Command{
		Path:    "/ip firewall address-list",
		Command: "print",
		Filter: map[string]string{
			"address": address,
			"list":    list,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) IpFirewallAddressList(address string, list string) (map[string]string, bool, error) {

	res, err := r.First(ipFirewallAddressList(address, list))
	return res, true, err

}
func addIpFirewallAddressList(address string, list string) Command {
	return Command{
		Path:    "/ip firewall address-list",
		Command: "add",
		Params: map[string]string{
			"address": address,
			"list":    list,
		},
		Extra: map[string]string{},
	}
}
func (r *Ros) AddIpFirewallAddressList(address string, list string) error {

	return r.Exec(addIpFirewallAddressList(address, list))

}
func removeIpFirewallAddressList(address string, list string) Command {
	return Command{
		Path:    "/ip firewall address-list",
		Command: "remove",
		Filter: map[string]string{
			"address": address,
			"list":    list,
		},
		Flags: map[string]bool{},
	}
}
func (r *Ros) RemoveIpFirewallAddressList(address string, list string) error {

	return r.Exec(removeIpFirewallAddressList(address, list))

}
func setIpFirewallAddressList(address string, list string, key, value string) Command {
	return Command{
		Path:    "/ip firewall address-list",
		Command: "set",
		Filter: map[string]string{
			"address": address,
			"list":    list,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setIpFirewallAddressListComment(address string, list string, value string) Command {
	return setIpFirewallAddressList(address, list, "comment", value)
}
func (r *Ros) SetIpFirewallAddressListComment(address string, list string, value string) error {

	return r.Exec(setIpFirewallAddressListComment(address, list, value))

}
func setIpFirewallAddressListDisabled(address string, list string, value bool) Command {
	return setIpFirewallAddressList(address, list, "disabled", FormatBool(value))
}
func (r *Ros) SetIpFirewallAddressListDisabled(address string, list string, value bool) error {

	return r.Exec(setIpFirewallAddressListDisabled(address, list, value))

}

func ipFirewallNat(filter map[string]string) Command {
	return Command{
		Path:    "/ip firewall nat",
		Command: "print",
		Filter:  filter,
		Flags:   map[string]bool{},
		Detail:  true,
	}
}
func (r *Ros) IpFirewallNat(filter map[string]string) (map[string]string, bool, error) {

	res, err := r.First(ipFirewallNat(filter))
	return res, true, err

}
func addIpFirewallNat(filter map[string]string) Command {
	return Command{
		Path:    "/ip firewall nat",
		Command: "add",
		Params:  filter,
		Extra:   map[string]string{},
	}
}
func (r *Ros) AddIpFirewallNat(filter map[string]string) error {

	return r.Exec(addIpFirewallNat(filter))

}
func removeIpFirewallNat(filter map[string]string) Command {
	return Command{
		Path:    "/ip firewall nat",
		Command: "remove",
		Filter:  filter,
		Flags:   map[string]bool{},
	}
}
func (r *Ros) RemoveIpFirewallNat(filter map[string]string) error {

	return r.Exec(removeIpFirewallNat(filter))

}
func setIpFirewallNat(filter map[string]string, key, value string) Command {
	return Command{
		Path:    "/ip firewall nat",
		Command: "set",
		Filter:  filter,
		Params: map[string]string{
			key: value,
		},
	}
}
func setIpFirewallNatComment(filter map[string]string, value string) Command {
	return setIpFirewallNat(filter, "comment", value)
}
func (r *Ros) SetIpFirewallNatComment(filter map[string]string, value string) error {

	return r.Exec(setIpFirewallNatComment(filter, value))

}
func setIpFirewallNatDisabled(filter map[string]string, value bool) Command {
	return setIpFirewallNat(filter, "disabled", FormatBool(value))
}
func (r *Ros) SetIpFirewallNatDisabled(filter map[string]string, value bool) error {

	return r.Exec(setIpFirewallNatDisabled(filter, value))

}

func ipNeighbors() Command {
	return Command{
		Path:    "/ip neighbor",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) IpNeighbors() ([]map[string]string, bool, error) {

	res, err := r.List(ipNeighbors())

	return res, true, err

}

func ipService(name string) Command {
	return Command{
		Path:    "/ip service",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) IpService(name string) (map[string]string, bool, error) {

	res, err := r.First(ipService(name))
	return res, true, err

}
func addIpService(name string) Command {
	return Command{
		Path:    "/ip service",
		Command: "add",
		Params: map[string]string{
			"name": name,
		},
		Extra: map[string]string{},
	}
}
func (r *Ros) AddIpService(name string) error {

	return r.Exec(addIpService(name))

}
func removeIpService(name string) Command {
	return Command{
		Path:    "/ip service",
		Command: "remove",
		Filter: map[string]string{
			"name": name,
		},
		Flags: map[string]bool{},
	}
}
func (r *Ros) RemoveIpService(name string) error {

	return r.Exec(removeIpService(name))

}
func setIpService(name string, key, value string) Command {
	return Command{
		Path:    "/ip service",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setIpServiceDisabled(name string, value bool) Command {
	return setIpService(name, "disabled", FormatBool(value))
}
func (r *Ros) SetIpServiceDisabled(name string, value bool) error {

	return r.Exec(setIpServiceDisabled(name, value))

}
func setIpServicePort(name string, value int) Command {
	return setIpService(name, "port", FormatInt(value))
}
func (r *Ros) SetIpServicePort(name string, value int) error {

	return r.Exec(setIpServicePort(name, value))

}
func setIpServiceAddress(name string, value string) Command {
	return setIpService(name, "address", value)
}
func (r *Ros) SetIpServiceAddress(name string, value string) error {

	return r.Exec(setIpServiceAddress(name, value))

}

func ipRoutes() Command {
	return Command{
		Path:    "/ip route",
		Command: "print",
		Flags: map[string]bool{
			"static": true,
		},
		Detail: true,
	}
}

func (r *Ros) IpRoutes() ([]map[string]string, bool, error) {

	res, err := r.List(ipRoutes())

	return res, true, err

}
func ipRoute(dstAddress string) Command {
	return Command{
		Path:    "/ip route",
		Command: "print",
		Filter: map[string]string{
			"dst-address": dstAddress,
		},
		Flags: map[string]bool{
			"static": true,
		},
		Detail: true,
	}
}
func (r *Ros) IpRoute(dstAddress string) (map[string]string, bool, error) {

	res, err := r.First(ipRoute(dstAddress))
	return res, true, err

}
func setIpRoute(dstAddress string, key, value string) Command {
	return Command{
		Path:    "/ip route",
		Command: "set",
		Filter: map[string]string{
			"dst-address": dstAddress,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setIpRouteComment(dstAddress string, value string) Command {
	return setIpRoute(dstAddress, "comment", value)
}
func (r *Ros) SetIpRouteComment(dstAddress string, value string) error {

	return r.Exec(setIpRouteComment(dstAddress, value))

}

func ipSocks() Command {
	return Command{
		Path:    "/ip socks",
		Command: "print",
		Filter:  map[string]string{},
		Flags:   map[string]bool{},
		Detail:  false,
	}
}
func (r *Ros) IpSocks() (map[string]string, bool, error) {

	res, err := r.Values(ipSocks())
	return res, true, err

}
func setIpSocks(key, value string) Command {
	return Command{
		Path:    "/ip socks",
		Command: "set",
		Filter:  map[string]string{},
		Params: map[string]string{
			key: value,
		},
	}
}
func setIpSocksEnabled(value bool) Command {
	return setIpSocks("enabled", FormatBool(value))
}
func (r *Ros) SetIpSocksEnabled(value bool) error {

	return r.Exec(setIpSocksEnabled(value))

}
func setIpSocksPort(value int) Command {
	return setIpSocks("port", FormatInt(value))
}
func (r *Ros) SetIpSocksPort(value int) error {

	return r.Exec(setIpSocksPort(value))

}
func setIpSocksConnectionIdleTimeout(value string) Command {
	return setIpSocks("connection-idle-timeout", value)
}
func (r *Ros) SetIpSocksConnectionIdleTimeout(value string) error {

	return r.Exec(setIpSocksConnectionIdleTimeout(value))

}
func setIpSocksMaxConnections(value int) Command {
	return setIpSocks("max-connections", FormatInt(value))
}
func (r *Ros) SetIpSocksMaxConnections(value int) error {

	return r.Exec(setIpSocksMaxConnections(value))

}

func ipSmb() Command {
	return Command{
		Path:    "/ip smb",
		Command: "print",
		Filter:  map[string]string{},
		Flags:   map[string]bool{},
		Detail:  false,
	}
}
func (r *Ros) IpSmb() (map[string]string, bool, error) {

	res, err := r.Values(ipSmb())
	return res, true, err

}
func setIpSmb(key, value string) Command {
	return Command{
		Path:    "/ip smb",
		Command: "set",
		Filter:  map[string]string{},
		Params: map[string]string{
			key: value,
		},
	}
}
func setIpSmbEnabled(value bool) Command {
	return setIpSmb("enabled", FormatBool(value))
}
func (r *Ros) SetIpSmbEnabled(value bool) error {

	return r.Exec(setIpSmbEnabled(value))

}
func setIpSmbDomain(value string) Command {
	return setIpSmb("domain", value)
}
func (r *Ros) SetIpSmbDomain(value string) error {

	return r.Exec(setIpSmbDomain(value))

}
func setIpSmbComment(value string) Command {
	return setIpSmb("comment", value)
}
func (r *Ros) SetIpSmbComment(value string) error {

	return r.Exec(setIpSmbComment(value))

}
func setIpSmbAllowGuests(value bool) Command {
	return setIpSmb("allow-guests", FormatBool(value))
}
func (r *Ros) SetIpSmbAllowGuests(value bool) error {

	return r.Exec(setIpSmbAllowGuests(value))

}
func setIpSmbInterface(value string) Command {
	return setIpSmb("interface", value)
}
func (r *Ros) SetIpSmbInterface(value string) error {

	return r.Exec(setIpSmbInterface(value))

}

func ipSmbShareses() Command {
	return Command{
		Path:    "/ip smb shares",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) IpSmbShareses() ([]map[string]string, bool, error) {

	res, err := r.List(ipSmbShareses())

	return res, true, err

}
func ipSmbShares(name string) Command {
	return Command{
		Path:    "/ip smb shares",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Flags: map[string]bool{
			"default": false,
		},
		Detail: true,
	}
}
func (r *Ros) IpSmbShares(name string) (map[string]string, bool, error) {

	res, err := r.First(ipSmbShares(name))
	return res, true, err

}
func addIpSmbShares(name string) Command {
	return Command{
		Path:    "/ip smb shares",
		Command: "add",
		Params: map[string]string{
			"name": name,
		},
		Extra: map[string]string{},
	}
}
func (r *Ros) AddIpSmbShares(name string) error {

	return r.Exec(addIpSmbShares(name))

}
func removeIpSmbShares(name string) Command {
	return Command{
		Path:    "/ip smb shares",
		Command: "remove",
		Filter: map[string]string{
			"name": name,
		},
		Flags: map[string]bool{
			"default": false,
		},
	}
}
func (r *Ros) RemoveIpSmbShares(name string) error {

	return r.Exec(removeIpSmbShares(name))

}
func setIpSmbShares(name string, key, value string) Command {
	return Command{
		Path:    "/ip smb shares",
		Command: "set",
		Filter: map[string]string{
			"name":     name,
			"!default": "",
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setIpSmbSharesDirectory(name string, value string) Command {
	return setIpSmbShares(name, "directory", value)
}
func (r *Ros) SetIpSmbSharesDirectory(name string, value string) error {

	return r.Exec(setIpSmbSharesDirectory(name, value))

}
func setIpSmbSharesComment(name string, value string) Command {
	return setIpSmbShares(name, "comment", value)
}
func (r *Ros) SetIpSmbSharesComment(name string, value string) error {

	return r.Exec(setIpSmbSharesComment(name, value))

}
func setIpSmbSharesDisabled(name string, value bool) Command {
	return setIpSmbShares(name, "disabled", FormatBool(value))
}
func (r *Ros) SetIpSmbSharesDisabled(name string, value bool) error {

	return r.Exec(setIpSmbSharesDisabled(name, value))

}
func setIpSmbSharesMaxSessions(name string, value int) Command {
	return setIpSmbShares(name, "max-sessions", FormatInt(value))
}
func (r *Ros) SetIpSmbSharesMaxSessions(name string, value int) error {

	return r.Exec(setIpSmbSharesMaxSessions(name, value))

}
func ipSmbSharesDefault() Command {
	return Command{
		Path:    "/ip smb shares",
		Command: "print",
		Flags: map[string]bool{
			"default": true,
		},
		Detail: true,
	}
}
func (r *Ros) IpSmbSharesDefault() (map[string]string, bool, error) {

	res, err := r.Values(ipSmbSharesDefault())

	return res, true, err

}
func setIpSmbSharesDefault(key, value string) Command {
	return Command{
		Path:    "/ip smb shares",
		Command: "set",
		Filter: map[string]string{
			"default": "",
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setIpSmbSharesDefaultName(value string) Command {
	return setIpSmbSharesDefault("name", value)
}
func (r *Ros) SetIpSmbSharesDefaultName(value string) error {

	return r.Exec(setIpSmbSharesDefaultName(value))

}
func setIpSmbSharesDefaultDirectory(value string) Command {
	return setIpSmbSharesDefault("directory", value)
}
func (r *Ros) SetIpSmbSharesDefaultDirectory(value string) error {

	return r.Exec(setIpSmbSharesDefaultDirectory(value))

}
func setIpSmbSharesDefaultComment(value string) Command {
	return setIpSmbSharesDefault("comment", value)
}
func (r *Ros) SetIpSmbSharesDefaultComment(value string) error {

	return r.Exec(setIpSmbSharesDefaultComment(value))

}
func setIpSmbSharesDefaultDisabled(value bool) Command {
	return setIpSmbSharesDefault("disabled", FormatBool(value))
}
func (r *Ros) SetIpSmbSharesDefaultDisabled(value bool) error {

	return r.Exec(setIpSmbSharesDefaultDisabled(value))

}
func setIpSmbSharesDefaultMaxSessions(value int) Command {
	return setIpSmbSharesDefault("max-sessions", FormatInt(value))
}
func (r *Ros) SetIpSmbSharesDefaultMaxSessions(value int) error {

	return r.Exec(setIpSmbSharesDefaultMaxSessions(value))

}

func ipSmbUserses() Command {
	return Command{
		Path:    "/ip smb users",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) IpSmbUserses() ([]map[string]string, bool, error) {

	res, err := r.List(ipSmbUserses())

	return res, true, err

}
func ipSmbUsers(name string) Command {
	return Command{
		Path:    "/ip smb users",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Flags: map[string]bool{
			"default": false,
		},
		Detail: true,
	}
}
func (r *Ros) IpSmbUsers(name string) (map[string]string, bool, error) {

	res, err := r.First(ipSmbUsers(name))
	return res, true, err

}
func addIpSmbUsers(name string) Command {
	return Command{
		Path:    "/ip smb users",
		Command: "add",
		Params: map[string]string{
			"name": name,
		},
		Extra: map[string]string{},
	}
}
func (r *Ros) AddIpSmbUsers(name string) error {

	return r.Exec(addIpSmbUsers(name))

}
func removeIpSmbUsers(name string) Command {
	return Command{
		Path:    "/ip smb users",
		Command: "remove",
		Filter: map[string]string{
			"name": name,
		},
		Flags: map[string]bool{
			"default": false,
		},
	}
}
func (r *Ros) RemoveIpSmbUsers(name string) error {

	return r.Exec(removeIpSmbUsers(name))

}
func setIpSmbUsers(name string, key, value string) Command {
	return Command{
		Path:    "/ip smb users",
		Command: "set",
		Filter: map[string]string{
			"name":     name,
			"!default": "",
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setIpSmbUsersPassword(name string, value string) Command {
	return setIpSmbUsers(name, "password", value)
}
func (r *Ros) SetIpSmbUsersPassword(name string, value string) error {

	return r.Exec(setIpSmbUsersPassword(name, value))

}
func setIpSmbUsersReadOnly(name string, value bool) Command {
	return setIpSmbUsers(name, "read-only", FormatBool(value))
}
func (r *Ros) SetIpSmbUsersReadOnly(name string, value bool) error {

	return r.Exec(setIpSmbUsersReadOnly(name, value))

}
func setIpSmbUsersDisabled(name string, value bool) Command {
	return setIpSmbUsers(name, "disabled", FormatBool(value))
}
func (r *Ros) SetIpSmbUsersDisabled(name string, value bool) error {

	return r.Exec(setIpSmbUsersDisabled(name, value))

}
func ipSmbUsersDefault() Command {
	return Command{
		Path:    "/ip smb users",
		Command: "print",
		Flags: map[string]bool{
			"default": true,
		},
		Detail: true,
	}
}
func (r *Ros) IpSmbUsersDefault() (map[string]string, bool, error) {

	res, err := r.Values(ipSmbUsersDefault())

	return res, true, err

}
func setIpSmbUsersDefault(key, value string) Command {
	return Command{
		Path:    "/ip smb users",
		Command: "set",
		Filter: map[string]string{
			"default": "",
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setIpSmbUsersDefaultName(value string) Command {
	return setIpSmbUsersDefault("name", value)
}
func (r *Ros) SetIpSmbUsersDefaultName(value string) error {

	return r.Exec(setIpSmbUsersDefaultName(value))

}
func setIpSmbUsersDefaultPassword(value string) Command {
	return setIpSmbUsersDefault("password", value)
}
func (r *Ros) SetIpSmbUsersDefaultPassword(value string) error {

	return r.Exec(setIpSmbUsersDefaultPassword(value))

}
func setIpSmbUsersDefaultReadOnly(value bool) Command {
	return setIpSmbUsersDefault("read-only", FormatBool(value))
}
func (r *Ros) SetIpSmbUsersDefaultReadOnly(value bool) error {

	return r.Exec(setIpSmbUsersDefaultReadOnly(value))

}
func setIpSmbUsersDefaultDisabled(value bool) Command {
	return setIpSmbUsersDefault("disabled", FormatBool(value))
}
func (r *Ros) SetIpSmbUsersDefaultDisabled(value bool) error {

	return r.Exec(setIpSmbUsersDefaultDisabled(value))

}

func ipSsh() Command {
	return Command{
		Path:    "/ip ssh",
		Command: "print",
		Filter:  map[string]string{},
		Flags:   map[string]bool{},
		Detail:  false,
	}
}
func (r *Ros) IpSsh() (map[string]string, bool, error) {

	res, err := r.Values(ipSsh())
	return res, true, err

}
func setIpSsh(key, value string) Command {
	return Command{
		Path:    "/ip ssh",
		Command: "set",
		Filter:  map[string]string{},
		Params: map[string]string{
			key: value,
		},
	}
}
func setIpSshStrongCrypto(value bool) Command {
	return setIpSsh("strong-crypto", FormatBool(value))
}
func (r *Ros) SetIpSshStrongCrypto(value bool) error {

	if r.AtLeast(6, 36) {

		return r.Exec(setIpSshStrongCrypto(value))

	}

	return nil
}
func setIpSshForwardingEnabled(value bool) Command {
	return setIpSsh("forwarding-enabled", FormatBool(value))
}
func (r *Ros) SetIpSshForwardingEnabled(value bool) error {

	return r.Exec(setIpSshForwardingEnabled(value))

}
func setIpSshAlwaysAllowPasswordLogin(value bool) Command {
	return setIpSsh("always-allow-password-login", FormatBool(value))
}
func (r *Ros) SetIpSshAlwaysAllowPasswordLogin(value bool) error {

	return r.Exec(setIpSshAlwaysAllowPasswordLogin(value))

}
func setIpSshHostKeySize(value int) Command {
	return setIpSsh("host-key-size", FormatInt(value))
}
func (r *Ros) SetIpSshHostKeySize(value int) error {

	if r.AtLeast(6, 36) {

		return r.Exec(setIpSshHostKeySize(value))

	}

	return nil
}

func ipTrafficFlow() Command {
	return Command{
		Path:    "/ip traffic-flow",
		Command: "print",
		Filter:  map[string]string{},
		Flags:   map[string]bool{},
		Detail:  true,
	}
}
func (r *Ros) IpTrafficFlow() (map[string]string, bool, error) {

	res, err := r.Values(ipTrafficFlow())
	return res, true, err

}
func setIpTrafficFlow(key, value string) Command {
	return Command{
		Path:    "/ip traffic-flow",
		Command: "set",
		Filter:  map[string]string{},
		Params: map[string]string{
			key: value,
		},
	}
}
func setIpTrafficFlowEnabled(value bool) Command {
	return setIpTrafficFlow("enabled", FormatBool(value))
}
func (r *Ros) SetIpTrafficFlowEnabled(value bool) error {

	return r.Exec(setIpTrafficFlowEnabled(value))

}
func setIpTrafficFlowInterfaces(value string) Command {
	return setIpTrafficFlow("interfaces", value)
}
func (r *Ros) SetIpTrafficFlowInterfaces(value string) error {

	return r.Exec(setIpTrafficFlowInterfaces(value))

}
func setIpTrafficFlowCacheEntries(value string) Command {
	return setIpTrafficFlow("cache-entries", value)
}
func (r *Ros) SetIpTrafficFlowCacheEntries(value string) error {

	return r.Exec(setIpTrafficFlowCacheEntries(value))

}
func setIpTrafficFlowActiveFlowTimeout(value string) Command {
	return setIpTrafficFlow("active-flow-timeout", value)
}
func (r *Ros) SetIpTrafficFlowActiveFlowTimeout(value string) error {

	return r.Exec(setIpTrafficFlowActiveFlowTimeout(value))

}
func setIpTrafficFlowInactiveFlowTimeout(value string) Command {
	return setIpTrafficFlow("inactive-flow-timeout", value)
}
func (r *Ros) SetIpTrafficFlowInactiveFlowTimeout(value string) error {

	return r.Exec(setIpTrafficFlowInactiveFlowTimeout(value))

}

func ipTrafficFlowTarget(dstAddress string, port int) Command {
	return Command{
		Path:    "/ip traffic-flow target",
		Command: "print",
		Filter: map[string]string{
			"dst-address": dstAddress,
			"port":        FormatInt(port),
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) IpTrafficFlowTarget(dstAddress string, port int) (map[string]string, bool, error) {

	if r.AtLeast(6, 31) {

		res, err := r.First(ipTrafficFlowTarget(dstAddress, port))
		return res, true, err

	}

	return nil, false, nil
}
func addIpTrafficFlowTarget(dstAddress string, port int) Command {
	return Command{
		Path:    "/ip traffic-flow target",
		Command: "add",
		Params: map[string]string{
			"dst-address": dstAddress,
			"port":        FormatInt(port),
		},
		Extra: map[string]string{},
	}
}
func (r *Ros) AddIpTrafficFlowTarget(dstAddress string, port int) error {

	if r.AtLeast(6, 31) {

		return r.Exec(addIpTrafficFlowTarget(dstAddress, port))

	}

	return nil
}
func removeIpTrafficFlowTarget(dstAddress string, port int) Command {
	return Command{
		Path:    "/ip traffic-flow target",
		Command: "remove",
		Filter: map[string]string{
			"dst-address": dstAddress,
			"port":        FormatInt(port),
		},
		Flags: map[string]bool{},
	}
}
func (r *Ros) RemoveIpTrafficFlowTarget(dstAddress string, port int) error {

	if r.AtLeast(6, 31) {

		return r.Exec(removeIpTrafficFlowTarget(dstAddress, port))

	}

	return nil
}
func setIpTrafficFlowTarget(dstAddress string, port int, key, value string) Command {
	return Command{
		Path:    "/ip traffic-flow target",
		Command: "set",
		Filter: map[string]string{
			"dst-address": dstAddress,
			"port":        FormatInt(port),
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setIpTrafficFlowTargetSrcAddress(dstAddress string, port int, value string) Command {
	return setIpTrafficFlowTarget(dstAddress, port, "src-address", value)
}
func (r *Ros) SetIpTrafficFlowTargetSrcAddress(dstAddress string, port int, value string) error {

	if r.AtLeast(6, 31) {

		return r.Exec(setIpTrafficFlowTargetSrcAddress(dstAddress, port, value))

	}

	return nil
}
func setIpTrafficFlowTargetVersion(dstAddress string, port int, value int) Command {
	return setIpTrafficFlowTarget(dstAddress, port, "version", FormatInt(value))
}
func (r *Ros) SetIpTrafficFlowTargetVersion(dstAddress string, port int, value int) error {

	if r.AtLeast(6, 31) {

		return r.Exec(setIpTrafficFlowTargetVersion(dstAddress, port, value))

	}

	return nil
}
func setIpTrafficFlowTargetV9TemplateRefresh(dstAddress string, port int, value string) Command {
	return setIpTrafficFlowTarget(dstAddress, port, "v9-template-refresh", value)
}
func (r *Ros) SetIpTrafficFlowTargetV9TemplateRefresh(dstAddress string, port int, value string) error {

	if r.AtLeast(6, 31) {

		return r.Exec(setIpTrafficFlowTargetV9TemplateRefresh(dstAddress, port, value))

	}

	return nil
}
func setIpTrafficFlowTargetV9TemplateTimeout(dstAddress string, port int, value string) Command {
	return setIpTrafficFlowTarget(dstAddress, port, "v9-template-timeout", value)
}
func (r *Ros) SetIpTrafficFlowTargetV9TemplateTimeout(dstAddress string, port int, value string) error {

	if r.AtLeast(6, 31) {

		return r.Exec(setIpTrafficFlowTargetV9TemplateTimeout(dstAddress, port, value))

	}

	return nil
}

func queueSimples() Command {
	return Command{
		Path:    "/queue simple",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) QueueSimples() ([]map[string]string, bool, error) {

	if r.AtLeast(6, 0) {
		res, err := r.List(queueSimples())

		return res, true, err

	}

	return nil, false, nil

}
func queueSimple(name string, target string) Command {
	return Command{
		Path:    "/queue simple",
		Command: "print",
		Filter: map[string]string{
			"name":   name,
			"target": target,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) QueueSimple(name string, target string) (map[string]string, bool, error) {

	if r.AtLeast(6, 0) {

		res, err := r.First(queueSimple(name, target))
		return res, true, err

	}

	return nil, false, nil
}
func addQueueSimple(name string, target string) Command {
	return Command{
		Path:    "/queue simple",
		Command: "add",
		Params: map[string]string{
			"name":   name,
			"target": target,
		},
		Extra: map[string]string{},
	}
}
func (r *Ros) AddQueueSimple(name string, target string) error {

	if r.AtLeast(6, 0) {

		return r.Exec(addQueueSimple(name, target))

	}

	return nil
}
func removeQueueSimple(name string, target string) Command {
	return Command{
		Path:    "/queue simple",
		Command: "remove",
		Filter: map[string]string{
			"name":   name,
			"target": target,
		},
		Flags: map[string]bool{},
	}
}
func (r *Ros) RemoveQueueSimple(name string, target string) error {

	if r.AtLeast(6, 0) {

		return r.Exec(removeQueueSimple(name, target))

	}

	return nil
}
func setQueueSimple(name string, target string, key, value string) Command {
	return Command{
		Path:    "/queue simple",
		Command: "set",
		Filter: map[string]string{
			"name":   name,
			"target": target,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setQueueSimpleComment(name string, target string, value string) Command {
	return setQueueSimple(name, target, "comment", value)
}
func (r *Ros) SetQueueSimpleComment(name string, target string, value string) error {

	if r.AtLeast(6, 0) {

		return r.Exec(setQueueSimpleComment(name, target, value))

	}

	return nil
}
func setQueueSimpleDisabled(name string, target string, value bool) Command {
	return setQueueSimple(name, target, "disabled", FormatBool(value))
}
func (r *Ros) SetQueueSimpleDisabled(name string, target string, value bool) error {

	if r.AtLeast(6, 0) {

		return r.Exec(setQueueSimpleDisabled(name, target, value))

	}

	return nil
}
func setQueueSimpleTime(name string, target string, value string) Command {
	return setQueueSimple(name, target, "time", value)
}
func (r *Ros) SetQueueSimpleTime(name string, target string, value string) error {

	if r.AtLeast(6, 0) {

		return r.Exec(setQueueSimpleTime(name, target, value))

	}

	return nil
}
func setQueueSimpleDstAddress(name string, target string, value string) Command {
	return setQueueSimple(name, target, "dst-address", value)
}
func (r *Ros) SetQueueSimpleDstAddress(name string, target string, value string) error {

	if r.AtLeast(6, 0) {

		return r.Exec(setQueueSimpleDstAddress(name, target, value))

	}

	return nil
}
func setQueueSimplePacketMarks(name string, target string, value string) Command {
	return setQueueSimple(name, target, "packet-marks", value)
}
func (r *Ros) SetQueueSimplePacketMarks(name string, target string, value string) error {

	if r.AtLeast(6, 0) {

		return r.Exec(setQueueSimplePacketMarks(name, target, value))

	}

	return nil
}
func setQueueSimpleParent(name string, target string, value string) Command {
	return setQueueSimple(name, target, "parent", value)
}
func (r *Ros) SetQueueSimpleParent(name string, target string, value string) error {

	if r.AtLeast(6, 0) {

		return r.Exec(setQueueSimpleParent(name, target, value))

	}

	return nil
}
func setQueueSimplePriority(name string, target string, value int) Command {
	return setQueueSimple(name, target, "priority", FormatInt(value))
}
func (r *Ros) SetQueueSimplePriority(name string, target string, value int) error {

	if r.AtLeast(6, 0) {

		return r.Exec(setQueueSimplePriority(name, target, value))

	}

	return nil
}
func setQueueSimpleQueue(name string, target string, value string) Command {
	return setQueueSimple(name, target, "queue", value)
}
func (r *Ros) SetQueueSimpleQueue(name string, target string, value string) error {

	if r.AtLeast(6, 0) {

		return r.Exec(setQueueSimpleQueue(name, target, value))

	}

	return nil
}
func setQueueSimpleLimitAt(name string, target string, value string) Command {
	return setQueueSimple(name, target, "limit-at", value)
}
func (r *Ros) SetQueueSimpleLimitAt(name string, target string, value string) error {

	if r.AtLeast(6, 0) {

		return r.Exec(setQueueSimpleLimitAt(name, target, value))

	}

	return nil
}
func setQueueSimpleMaxLimit(name string, target string, value string) Command {
	return setQueueSimple(name, target, "max-limit", value)
}
func (r *Ros) SetQueueSimpleMaxLimit(name string, target string, value string) error {

	if r.AtLeast(6, 0) {

		return r.Exec(setQueueSimpleMaxLimit(name, target, value))

	}

	return nil
}
func setQueueSimpleBurstLimit(name string, target string, value string) Command {
	return setQueueSimple(name, target, "burst-limit", value)
}
func (r *Ros) SetQueueSimpleBurstLimit(name string, target string, value string) error {

	if r.AtLeast(6, 0) {

		return r.Exec(setQueueSimpleBurstLimit(name, target, value))

	}

	return nil
}
func setQueueSimpleBurstTime(name string, target string, value string) Command {
	return setQueueSimple(name, target, "burst-time", value)
}
func (r *Ros) SetQueueSimpleBurstTime(name string, target string, value string) error {

	if r.AtLeast(6, 0) {

		return r.Exec(setQueueSimpleBurstTime(name, target, value))

	}

	return nil
}
func setQueueSimpleBurstThreshold(name string, target string, value string) Command {
	return setQueueSimple(name, target, "burst-threshold", value)
}
func (r *Ros) SetQueueSimpleBurstThreshold(name string, target string, value string) error {

	if r.AtLeast(6, 0) {

		return r.Exec(setQueueSimpleBurstThreshold(name, target, value))

	}

	return nil
}
func setQueueSimpleTotalQueue(name string, target string, value string) Command {
	return setQueueSimple(name, target, "total-queue", value)
}
func (r *Ros) SetQueueSimpleTotalQueue(name string, target string, value string) error {

	if r.AtLeast(6, 0) {

		return r.Exec(setQueueSimpleTotalQueue(name, target, value))

	}

	return nil
}
func setQueueSimpleTotalLimitAt(name string, target string, value string) Command {
	return setQueueSimple(name, target, "total-limit-at", value)
}
func (r *Ros) SetQueueSimpleTotalLimitAt(name string, target string, value string) error {

	if r.AtLeast(6, 0) {

		return r.Exec(setQueueSimpleTotalLimitAt(name, target, value))

	}

	return nil
}
func setQueueSimpleTotalMaxLimit(name string, target string, value string) Command {
	return setQueueSimple(name, target, "total-max-limit", value)
}
func (r *Ros) SetQueueSimpleTotalMaxLimit(name string, target string, value string) error {

	if r.AtLeast(6, 0) {

		return r.Exec(setQueueSimpleTotalMaxLimit(name, target, value))

	}

	return nil
}
func setQueueSimpleTotalBurstLimit(name string, target string, value string) Command {
	return setQueueSimple(name, target, "total-burst-limit", value)
}
func (r *Ros) SetQueueSimpleTotalBurstLimit(name string, target string, value string) error {

	if r.AtLeast(6, 0) {

		return r.Exec(setQueueSimpleTotalBurstLimit(name, target, value))

	}

	return nil
}
func setQueueSimpleTotalBurstTime(name string, target string, value string) Command {
	return setQueueSimple(name, target, "total-burst-time", value)
}
func (r *Ros) SetQueueSimpleTotalBurstTime(name string, target string, value string) error {

	if r.AtLeast(6, 0) {

		return r.Exec(setQueueSimpleTotalBurstTime(name, target, value))

	}

	return nil
}
func setQueueSimpleTotalBurstThreshold(name string, target string, value string) Command {
	return setQueueSimple(name, target, "total-burst-threshold", value)
}
func (r *Ros) SetQueueSimpleTotalBurstThreshold(name string, target string, value string) error {

	if r.AtLeast(6, 0) {

		return r.Exec(setQueueSimpleTotalBurstThreshold(name, target, value))

	}

	return nil
}
func setQueueSimpleBucketSize(name string, target string, value string) Command {
	return setQueueSimple(name, target, "bucket-size", value)
}
func (r *Ros) SetQueueSimpleBucketSize(name string, target string, value string) error {

	if r.AtLeast(6, 0) {

		if r.AtLeast(6, 35) {

			return r.Exec(setQueueSimpleBucketSize(name, target, value))

		}

	}

	return nil
}

func queueTrees() Command {
	return Command{
		Path:    "/queue tree",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) QueueTrees() ([]map[string]string, bool, error) {

	res, err := r.List(queueTrees())

	return res, true, err

}
func queueTree(name string) Command {
	return Command{
		Path:    "/queue tree",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) QueueTree(name string) (map[string]string, bool, error) {

	res, err := r.First(queueTree(name))
	return res, true, err

}
func addQueueTree(name string, packetMarks string) Command {
	return Command{
		Path:    "/queue tree",
		Command: "add",
		Params: map[string]string{
			"name": name,
		},
		Extra: map[string]string{
			"packet-marks": packetMarks,
		},
	}
}
func (r *Ros) AddQueueTree(name string, packetMarks string) error {

	return r.Exec(addQueueTree(name, packetMarks))

}
func removeQueueTree(name string) Command {
	return Command{
		Path:    "/queue tree",
		Command: "remove",
		Filter: map[string]string{
			"name": name,
		},
		Flags: map[string]bool{},
	}
}
func (r *Ros) RemoveQueueTree(name string) error {

	return r.Exec(removeQueueTree(name))

}
func setQueueTree(name string, key, value string) Command {
	return Command{
		Path:    "/queue tree",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setQueueTreePacketMarks(name string, value string) Command {
	return setQueueTree(name, "packet-marks", value)
}
func (r *Ros) SetQueueTreePacketMarks(name string, value string) error {

	return r.Exec(setQueueTreePacketMarks(name, value))

}
func setQueueTreeComment(name string, value string) Command {
	return setQueueTree(name, "comment", value)
}
func (r *Ros) SetQueueTreeComment(name string, value string) error {

	return r.Exec(setQueueTreeComment(name, value))

}
func setQueueTreeDisabled(name string, value bool) Command {
	return setQueueTree(name, "disabled", FormatBool(value))
}
func (r *Ros) SetQueueTreeDisabled(name string, value bool) error {

	return r.Exec(setQueueTreeDisabled(name, value))

}
func setQueueTreeParent(name string, value string) Command {
	return setQueueTree(name, "parent", value)
}
func (r *Ros) SetQueueTreeParent(name string, value string) error {

	return r.Exec(setQueueTreeParent(name, value))

}
func setQueueTreePriority(name string, value int) Command {
	return setQueueTree(name, "priority", FormatInt(value))
}
func (r *Ros) SetQueueTreePriority(name string, value int) error {

	return r.Exec(setQueueTreePriority(name, value))

}
func setQueueTreeQueue(name string, value string) Command {
	return setQueueTree(name, "queue", value)
}
func (r *Ros) SetQueueTreeQueue(name string, value string) error {

	return r.Exec(setQueueTreeQueue(name, value))

}
func setQueueTreeLimitAt(name string, value string) Command {
	return setQueueTree(name, "limit-at", value)
}
func (r *Ros) SetQueueTreeLimitAt(name string, value string) error {

	return r.Exec(setQueueTreeLimitAt(name, value))

}
func setQueueTreeMaxLimit(name string, value string) Command {
	return setQueueTree(name, "max-limit", value)
}
func (r *Ros) SetQueueTreeMaxLimit(name string, value string) error {

	return r.Exec(setQueueTreeMaxLimit(name, value))

}
func setQueueTreeBurstLimit(name string, value string) Command {
	return setQueueTree(name, "burst-limit", value)
}
func (r *Ros) SetQueueTreeBurstLimit(name string, value string) error {

	return r.Exec(setQueueTreeBurstLimit(name, value))

}
func setQueueTreeBurstTime(name string, value string) Command {
	return setQueueTree(name, "burst-time", value)
}
func (r *Ros) SetQueueTreeBurstTime(name string, value string) error {

	return r.Exec(setQueueTreeBurstTime(name, value))

}
func setQueueTreeBurstThreshold(name string, value string) Command {
	return setQueueTree(name, "burst-threshold", value)
}
func (r *Ros) SetQueueTreeBurstThreshold(name string, value string) error {

	return r.Exec(setQueueTreeBurstThreshold(name, value))

}
func setQueueTreeBucketSize(name string, value string) Command {
	return setQueueTree(name, "bucket-size", value)
}
func (r *Ros) SetQueueTreeBucketSize(name string, value string) error {

	if r.AtLeast(6, 35) {

		return r.Exec(setQueueTreeBucketSize(name, value))

	}

	return nil
}

func radius(service string) Command {
	return Command{
		Path:    "/radius",
		Command: "print",
		Filter: map[string]string{
			"service": service,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) Radius(service string) (map[string]string, bool, error) {

	res, err := r.First(radius(service))
	return res, true, err

}
func addRadius(service string) Command {
	return Command{
		Path:    "/radius",
		Command: "add",
		Params: map[string]string{
			"service": service,
		},
		Extra: map[string]string{},
	}
}
func (r *Ros) AddRadius(service string) error {

	return r.Exec(addRadius(service))

}
func removeRadius(service string) Command {
	return Command{
		Path:    "/radius",
		Command: "remove",
		Filter: map[string]string{
			"service": service,
		},
		Flags: map[string]bool{},
	}
}
func (r *Ros) RemoveRadius(service string) error {

	return r.Exec(removeRadius(service))

}
func setRadius(service string, key, value string) Command {
	return Command{
		Path:    "/radius",
		Command: "set",
		Filter: map[string]string{
			"service": service,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setRadiusSecret(service string, value string) Command {
	return setRadius(service, "secret", value)
}
func (r *Ros) SetRadiusSecret(service string, value string) error {

	return r.Exec(setRadiusSecret(service, value))

}
func setRadiusAddress(service string, value string) Command {
	return setRadius(service, "address", value)
}
func (r *Ros) SetRadiusAddress(service string, value string) error {

	return r.Exec(setRadiusAddress(service, value))

}
func setRadiusSrcAddress(service string, value string) Command {
	return setRadius(service, "src-address", value)
}
func (r *Ros) SetRadiusSrcAddress(service string, value string) error {

	return r.Exec(setRadiusSrcAddress(service, value))

}
func setRadiusTimeout(service string, value string) Command {
	return setRadius(service, "timeout", value)
}
func (r *Ros) SetRadiusTimeout(service string, value string) error {

	return r.Exec(setRadiusTimeout(service, value))

}
func setRadiusDisabled(service string, value bool) Command {
	return setRadius(service, "disabled", FormatBool(value))
}
func (r *Ros) SetRadiusDisabled(service string, value bool) error {

	return r.Exec(setRadiusDisabled(service, value))

}

func radiusIncoming() Command {
	return Command{
		Path:    "/radius incoming",
		Command: "print",
		Filter:  map[string]string{},
		Flags:   map[string]bool{},
		Detail:  false,
	}
}
func (r *Ros) RadiusIncoming() (map[string]string, bool, error) {

	res, err := r.Values(radiusIncoming())
	return res, true, err

}
func setRadiusIncoming(key, value string) Command {
	return Command{
		Path:    "/radius incoming",
		Command: "set",
		Filter:  map[string]string{},
		Params: map[string]string{
			key: value,
		},
	}
}
func setRadiusIncomingAccept(value bool) Command {
	return setRadiusIncoming("accept", FormatBool(value))
}
func (r *Ros) SetRadiusIncomingAccept(value bool) error {

	return r.Exec(setRadiusIncomingAccept(value))

}
func setRadiusIncomingPort(value int) Command {
	return setRadiusIncoming("port", FormatInt(value))
}
func (r *Ros) SetRadiusIncomingPort(value int) error {

	return r.Exec(setRadiusIncomingPort(value))

}

func routingBgpNetworks() Command {
	return Command{
		Path:    "/routing bgp network",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) RoutingBgpNetworks() ([]map[string]string, bool, error) {

	res, err := r.List(routingBgpNetworks())

	return res, true, err

}
func routingBgpNetwork(address string) Command {
	return Command{
		Path:    "/routing bgp network",
		Command: "print",
		Filter: map[string]string{
			"address": address,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) RoutingBgpNetwork(address string) (map[string]string, bool, error) {

	res, err := r.First(routingBgpNetwork(address))
	return res, true, err

}
func setRoutingBgpNetwork(address string, key, value string) Command {
	return Command{
		Path:    "/routing bgp network",
		Command: "set",
		Filter: map[string]string{
			"address": address,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setRoutingBgpNetworkComment(address string, value string) Command {
	return setRoutingBgpNetwork(address, "comment", value)
}
func (r *Ros) SetRoutingBgpNetworkComment(address string, value string) error {

	return r.Exec(setRoutingBgpNetworkComment(address, value))

}

func routingBgpAggregates() Command {
	return Command{
		Path:    "/routing bgp aggregate",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) RoutingBgpAggregates() ([]map[string]string, bool, error) {

	res, err := r.List(routingBgpAggregates())

	return res, true, err

}
func routingBgpAggregate(instance string, prefix string) Command {
	return Command{
		Path:    "/routing bgp aggregate",
		Command: "print",
		Filter: map[string]string{
			"instance": instance,
			"prefix":   prefix,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) RoutingBgpAggregate(instance string, prefix string) (map[string]string, bool, error) {

	res, err := r.First(routingBgpAggregate(instance, prefix))
	return res, true, err

}
func setRoutingBgpAggregate(instance string, prefix string, key, value string) Command {
	return Command{
		Path:    "/routing bgp aggregate",
		Command: "set",
		Filter: map[string]string{
			"instance": instance,
			"prefix":   prefix,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setRoutingBgpAggregateComment(instance string, prefix string, value string) Command {
	return setRoutingBgpAggregate(instance, prefix, "comment", value)
}
func (r *Ros) SetRoutingBgpAggregateComment(instance string, prefix string, value string) error {

	return r.Exec(setRoutingBgpAggregateComment(instance, prefix, value))

}

func routingBgpInstances() Command {
	return Command{
		Path:    "/routing bgp instance",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) RoutingBgpInstances() ([]map[string]string, bool, error) {

	res, err := r.List(routingBgpInstances())

	return res, true, err

}
func routingBgpInstance(name string) Command {
	return Command{
		Path:    "/routing bgp instance",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) RoutingBgpInstance(name string) (map[string]string, bool, error) {

	res, err := r.First(routingBgpInstance(name))
	return res, true, err

}
func setRoutingBgpInstance(name string, key, value string) Command {
	return Command{
		Path:    "/routing bgp instance",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setRoutingBgpInstanceRouterId(name string, value string) Command {
	return setRoutingBgpInstance(name, "router-id", value)
}
func (r *Ros) SetRoutingBgpInstanceRouterId(name string, value string) error {

	return r.Exec(setRoutingBgpInstanceRouterId(name, value))

}
func setRoutingBgpInstanceComment(name string, value string) Command {
	return setRoutingBgpInstance(name, "comment", value)
}
func (r *Ros) SetRoutingBgpInstanceComment(name string, value string) error {

	return r.Exec(setRoutingBgpInstanceComment(name, value))

}

func routingBgpPeers() Command {
	return Command{
		Path:    "/routing bgp peer",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) RoutingBgpPeers() ([]map[string]string, bool, error) {

	res, err := r.List(routingBgpPeers())

	return res, true, err

}
func routingBgpPeer(remoteAddress string) Command {
	return Command{
		Path:    "/routing bgp peer",
		Command: "print",
		Filter: map[string]string{
			"remote-address": remoteAddress,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) RoutingBgpPeer(remoteAddress string) (map[string]string, bool, error) {

	res, err := r.First(routingBgpPeer(remoteAddress))
	return res, true, err

}
func addRoutingBgpPeer(remoteAddress string) Command {
	return Command{
		Path:    "/routing bgp peer",
		Command: "add",
		Params: map[string]string{
			"remote-address": remoteAddress,
		},
		Extra: map[string]string{},
	}
}
func (r *Ros) AddRoutingBgpPeer(remoteAddress string) error {

	return r.Exec(addRoutingBgpPeer(remoteAddress))

}
func removeRoutingBgpPeer(remoteAddress string) Command {
	return Command{
		Path:    "/routing bgp peer",
		Command: "remove",
		Filter: map[string]string{
			"remote-address": remoteAddress,
		},
		Flags: map[string]bool{},
	}
}
func (r *Ros) RemoveRoutingBgpPeer(remoteAddress string) error {

	return r.Exec(removeRoutingBgpPeer(remoteAddress))

}
func setRoutingBgpPeer(remoteAddress string, key, value string) Command {
	return Command{
		Path:    "/routing bgp peer",
		Command: "set",
		Filter: map[string]string{
			"remote-address": remoteAddress,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setRoutingBgpPeerComment(remoteAddress string, value string) Command {
	return setRoutingBgpPeer(remoteAddress, "comment", value)
}
func (r *Ros) SetRoutingBgpPeerComment(remoteAddress string, value string) error {

	return r.Exec(setRoutingBgpPeerComment(remoteAddress, value))

}
func setRoutingBgpPeerName(remoteAddress string, value string) Command {
	return setRoutingBgpPeer(remoteAddress, "name", value)
}
func (r *Ros) SetRoutingBgpPeerName(remoteAddress string, value string) error {

	return r.Exec(setRoutingBgpPeerName(remoteAddress, value))

}

func routingFilters() Command {
	return Command{
		Path:    "/routing filter",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) RoutingFilters() ([]map[string]string, bool, error) {

	res, err := r.List(routingFilters())

	return res, true, err

}
func routingFilter(comment string, action string, chain string) Command {
	return Command{
		Path:    "/routing filter",
		Command: "print",
		Filter: map[string]string{
			"action":  action,
			"chain":   chain,
			"comment": comment,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) RoutingFilter(comment string, action string, chain string) (map[string]string, bool, error) {

	res, err := r.First(routingFilter(comment, action, chain))
	return res, true, err

}
func addRoutingFilter(comment string, action string, chain string) Command {
	return Command{
		Path:    "/routing filter",
		Command: "add",
		Params: map[string]string{
			"action":  action,
			"chain":   chain,
			"comment": comment,
		},
		Extra: map[string]string{},
	}
}
func (r *Ros) AddRoutingFilter(comment string, action string, chain string) error {

	return r.Exec(addRoutingFilter(comment, action, chain))

}
func removeRoutingFilter(comment string, action string, chain string) Command {
	return Command{
		Path:    "/routing filter",
		Command: "remove",
		Filter: map[string]string{
			"action":  action,
			"chain":   chain,
			"comment": comment,
		},
		Flags: map[string]bool{},
	}
}
func (r *Ros) RemoveRoutingFilter(comment string, action string, chain string) error {

	return r.Exec(removeRoutingFilter(comment, action, chain))

}
func setRoutingFilter(comment string, action string, chain string, key, value string) Command {
	return Command{
		Path:    "/routing filter",
		Command: "set",
		Filter: map[string]string{
			"action":  action,
			"chain":   chain,
			"comment": comment,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setRoutingFilterOff(comment string, action string, chain string, key string) Command {
	return Command{
		Path:    "/routing filter",
		Command: "set",
		Filter: map[string]string{
			"action":  action,
			"chain":   chain,
			"comment": comment,
		},
		Flags: map[string]bool{
			key: false,
		},
	}
}
func setRoutingFilterPrefix(comment string, action string, chain string, value string) Command {
	return setRoutingFilter(comment, action, chain, "prefix", value)
}
func (r *Ros) SetRoutingFilterPrefix(comment string, action string, chain string, value string) error {

	return r.Exec(setRoutingFilterPrefix(comment, action, chain, value))

}
func setRoutingFilterPrefixOff(comment string, action string, chain string) Command {
	return setRoutingFilterOff(comment, action, chain, "prefix")
}
func (r *Ros) SetRoutingFilterPrefixOff(comment string, action string, chain string) error {

	return r.Exec(setRoutingFilterPrefixOff(comment, action, chain))

}
func setRoutingFilterPrefixLength(comment string, action string, chain string, value string) Command {
	return setRoutingFilter(comment, action, chain, "prefix-length", value)
}
func (r *Ros) SetRoutingFilterPrefixLength(comment string, action string, chain string, value string) error {

	return r.Exec(setRoutingFilterPrefixLength(comment, action, chain, value))

}
func setRoutingFilterPrefixLengthOff(comment string, action string, chain string) Command {
	return setRoutingFilterOff(comment, action, chain, "prefix-length")
}
func (r *Ros) SetRoutingFilterPrefixLengthOff(comment string, action string, chain string) error {

	return r.Exec(setRoutingFilterPrefixLengthOff(comment, action, chain))

}
func setRoutingFilterProtocol(comment string, action string, chain string, value string) Command {
	return setRoutingFilter(comment, action, chain, "protocol", value)
}
func (r *Ros) SetRoutingFilterProtocol(comment string, action string, chain string, value string) error {

	return r.Exec(setRoutingFilterProtocol(comment, action, chain, value))

}
func setRoutingFilterProtocolOff(comment string, action string, chain string) Command {
	return setRoutingFilterOff(comment, action, chain, "protocol")
}
func (r *Ros) SetRoutingFilterProtocolOff(comment string, action string, chain string) error {

	return r.Exec(setRoutingFilterProtocolOff(comment, action, chain))

}
func setRoutingFilterInvertMatch(comment string, action string, chain string, value bool) Command {
	return setRoutingFilter(comment, action, chain, "invert-match", FormatBool(value))
}
func (r *Ros) SetRoutingFilterInvertMatch(comment string, action string, chain string, value bool) error {

	return r.Exec(setRoutingFilterInvertMatch(comment, action, chain, value))

}
func setRoutingFilterSetBgpPrependPath(comment string, action string, chain string, value string) Command {
	return setRoutingFilter(comment, action, chain, "set-bgp-prepend-path", value)
}
func (r *Ros) SetRoutingFilterSetBgpPrependPath(comment string, action string, chain string, value string) error {

	return r.Exec(setRoutingFilterSetBgpPrependPath(comment, action, chain, value))

}

func routingOspfInstances() Command {
	return Command{
		Path:    "/routing ospf instance",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) RoutingOspfInstances() ([]map[string]string, bool, error) {

	res, err := r.List(routingOspfInstances())

	return res, true, err

}
func routingOspfInstance(name string) Command {
	return Command{
		Path:    "/routing ospf instance",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) RoutingOspfInstance(name string) (map[string]string, bool, error) {

	res, err := r.First(routingOspfInstance(name))
	return res, true, err

}
func setRoutingOspfInstance(name string, key, value string) Command {
	return Command{
		Path:    "/routing ospf instance",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setRoutingOspfInstanceRouterId(name string, value string) Command {
	return setRoutingOspfInstance(name, "router-id", value)
}
func (r *Ros) SetRoutingOspfInstanceRouterId(name string, value string) error {

	return r.Exec(setRoutingOspfInstanceRouterId(name, value))

}
func setRoutingOspfInstanceComment(name string, value string) Command {
	return setRoutingOspfInstance(name, "comment", value)
}
func (r *Ros) SetRoutingOspfInstanceComment(name string, value string) error {

	return r.Exec(setRoutingOspfInstanceComment(name, value))

}

func routingOspfInterfaces() Command {
	return Command{
		Path:    "/routing ospf interface",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) RoutingOspfInterfaces() ([]map[string]string, bool, error) {

	res, err := r.List(routingOspfInterfaces())

	return res, true, err

}
func routingOspfInterface(iface string) Command {
	return Command{
		Path:    "/routing ospf interface",
		Command: "print",
		Filter: map[string]string{
			"interface": iface,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) RoutingOspfInterface(iface string) (map[string]string, bool, error) {

	res, err := r.First(routingOspfInterface(iface))
	return res, true, err

}
func setRoutingOspfInterface(iface string, key, value string) Command {
	return Command{
		Path:    "/routing ospf interface",
		Command: "set",
		Filter: map[string]string{
			"interface": iface,
			"dynamic":   "no",
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setRoutingOspfInterfaceCost(iface string, value int) Command {
	return setRoutingOspfInterface(iface, "cost", FormatInt(value))
}
func (r *Ros) SetRoutingOspfInterfaceCost(iface string, value int) error {

	return r.Exec(setRoutingOspfInterfaceCost(iface, value))

}
func setRoutingOspfInterfaceComment(iface string, value string) Command {
	return setRoutingOspfInterface(iface, "comment", value)
}
func (r *Ros) SetRoutingOspfInterfaceComment(iface string, value string) error {

	return r.Exec(setRoutingOspfInterfaceComment(iface, value))

}

func routingOspfNetworks() Command {
	return Command{
		Path:    "/routing ospf network",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) RoutingOspfNetworks() ([]map[string]string, bool, error) {

	res, err := r.List(routingOspfNetworks())

	return res, true, err

}
func routingOspfNetwork(network string) Command {
	return Command{
		Path:    "/routing ospf network",
		Command: "print",
		Filter: map[string]string{
			"network": network,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) RoutingOspfNetwork(network string) (map[string]string, bool, error) {

	res, err := r.First(routingOspfNetwork(network))
	return res, true, err

}
func setRoutingOspfNetwork(network string, key, value string) Command {
	return Command{
		Path:    "/routing ospf network",
		Command: "set",
		Filter: map[string]string{
			"network": network,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setRoutingOspfNetworkComment(network string, value string) Command {
	return setRoutingOspfNetwork(network, "comment", value)
}
func (r *Ros) SetRoutingOspfNetworkComment(network string, value string) error {

	return r.Exec(setRoutingOspfNetworkComment(network, value))

}

func routingOspfNbmaNeighbors() Command {
	return Command{
		Path:    "/routing ospf nbma-neighbor",
		Command: "print",
		Flags:   map[string]bool{},
		Detail:  true,
	}
}

func (r *Ros) RoutingOspfNbmaNeighbors() ([]map[string]string, bool, error) {

	res, err := r.List(routingOspfNbmaNeighbors())

	return res, true, err

}
func routingOspfNbmaNeighbor(address string) Command {
	return Command{
		Path:    "/routing ospf nbma-neighbor",
		Command: "print",
		Filter: map[string]string{
			"address": address,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) RoutingOspfNbmaNeighbor(address string) (map[string]string, bool, error) {

	res, err := r.First(routingOspfNbmaNeighbor(address))
	return res, true, err

}
func addRoutingOspfNbmaNeighbor(address string) Command {
	return Command{
		Path:    "/routing ospf nbma-neighbor",
		Command: "add",
		Params: map[string]string{
			"address": address,
		},
		Extra: map[string]string{},
	}
}
func (r *Ros) AddRoutingOspfNbmaNeighbor(address string) error {

	return r.Exec(addRoutingOspfNbmaNeighbor(address))

}
func removeRoutingOspfNbmaNeighbor(address string) Command {
	return Command{
		Path:    "/routing ospf nbma-neighbor",
		Command: "remove",
		Filter: map[string]string{
			"address": address,
		},
		Flags: map[string]bool{},
	}
}
func (r *Ros) RemoveRoutingOspfNbmaNeighbor(address string) error {

	return r.Exec(removeRoutingOspfNbmaNeighbor(address))

}
func setRoutingOspfNbmaNeighbor(address string, key, value string) Command {
	return Command{
		Path:    "/routing ospf nbma-neighbor",
		Command: "set",
		Filter: map[string]string{
			"address": address,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setRoutingOspfNbmaNeighborComment(address string, value string) Command {
	return setRoutingOspfNbmaNeighbor(address, "comment", value)
}
func (r *Ros) SetRoutingOspfNbmaNeighborComment(address string, value string) error {

	return r.Exec(setRoutingOspfNbmaNeighborComment(address, value))

}
func setRoutingOspfNbmaNeighborPriority(address string, value int) Command {
	return setRoutingOspfNbmaNeighbor(address, "priority", FormatInt(value))
}
func (r *Ros) SetRoutingOspfNbmaNeighborPriority(address string, value int) error {

	return r.Exec(setRoutingOspfNbmaNeighborPriority(address, value))

}
func setRoutingOspfNbmaNeighborPollInterval(address string, value string) Command {
	return setRoutingOspfNbmaNeighbor(address, "poll-interval", value)
}
func (r *Ros) SetRoutingOspfNbmaNeighborPollInterval(address string, value string) error {

	return r.Exec(setRoutingOspfNbmaNeighborPollInterval(address, value))

}

func snmp() Command {
	return Command{
		Path:    "/snmp",
		Command: "print",
		Filter:  map[string]string{},
		Flags:   map[string]bool{},
		Detail:  false,
	}
}
func (r *Ros) Snmp() (map[string]string, bool, error) {

	res, err := r.Values(snmp())
	return res, true, err

}
func setSnmp(key, value string) Command {
	return Command{
		Path:    "/snmp",
		Command: "set",
		Filter:  map[string]string{},
		Params: map[string]string{
			key: value,
		},
	}
}
func setSnmpEnabled(value bool) Command {
	return setSnmp("enabled", FormatBool(value))
}
func (r *Ros) SetSnmpEnabled(value bool) error {

	return r.Exec(setSnmpEnabled(value))

}
func setSnmpLocation(value string) Command {
	return setSnmp("location", value)
}
func (r *Ros) SetSnmpLocation(value string) error {

	return r.Exec(setSnmpLocation(value))

}
func setSnmpContact(value string) Command {
	return setSnmp("contact", value)
}
func (r *Ros) SetSnmpContact(value string) error {

	return r.Exec(setSnmpContact(value))

}
func setSnmpEngineId(value string) Command {
	return setSnmp("engine-id", value)
}
func (r *Ros) SetSnmpEngineId(value string) error {

	return r.Exec(setSnmpEngineId(value))

}
func setSnmpTrapCommunity(value string) Command {
	return setSnmp("trap-community", value)
}
func (r *Ros) SetSnmpTrapCommunity(value string) error {

	return r.Exec(setSnmpTrapCommunity(value))

}
func setSnmpTrapGenerators(value string) Command {
	return setSnmp("trap-generators", value)
}
func (r *Ros) SetSnmpTrapGenerators(value string) error {

	return r.Exec(setSnmpTrapGenerators(value))

}
func setSnmpTrapTarget(value string) Command {
	return setSnmp("trap-target", value)
}
func (r *Ros) SetSnmpTrapTarget(value string) error {

	return r.Exec(setSnmpTrapTarget(value))

}
func setSnmpTrapVersion(value int) Command {
	return setSnmp("trap-version", FormatInt(value))
}
func (r *Ros) SetSnmpTrapVersion(value int) error {

	return r.Exec(setSnmpTrapVersion(value))

}

func snmpCommunity(name string) Command {
	return Command{
		Path:    "/snmp community",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Flags: map[string]bool{
			"default": false,
		},
		Detail: false,
	}
}
func (r *Ros) SnmpCommunity(name string) (map[string]string, bool, error) {

	res, err := r.First(snmpCommunity(name))
	return res, true, err

}
func addSnmpCommunity(name string) Command {
	return Command{
		Path:    "/snmp community",
		Command: "add",
		Params: map[string]string{
			"name": name,
		},
		Extra: map[string]string{},
	}
}
func (r *Ros) AddSnmpCommunity(name string) error {

	return r.Exec(addSnmpCommunity(name))

}
func removeSnmpCommunity(name string) Command {
	return Command{
		Path:    "/snmp community",
		Command: "remove",
		Filter: map[string]string{
			"name": name,
		},
		Flags: map[string]bool{
			"default": false,
		},
	}
}
func (r *Ros) RemoveSnmpCommunity(name string) error {

	return r.Exec(removeSnmpCommunity(name))

}
func setSnmpCommunity(name string, key, value string) Command {
	return Command{
		Path:    "/snmp community",
		Command: "set",
		Filter: map[string]string{
			"name":     name,
			"!default": "",
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setSnmpCommunityAuthenticationPassword(name string, value string) Command {
	return setSnmpCommunity(name, "authentication-password", value)
}
func (r *Ros) SetSnmpCommunityAuthenticationPassword(name string, value string) error {

	return r.Exec(setSnmpCommunityAuthenticationPassword(name, value))

}
func setSnmpCommunityAuthenticationProtocol(name string, value string) Command {
	return setSnmpCommunity(name, "authentication-protocol", value)
}
func (r *Ros) SetSnmpCommunityAuthenticationProtocol(name string, value string) error {

	return r.Exec(setSnmpCommunityAuthenticationProtocol(name, value))

}
func setSnmpCommunityEncryptionPassword(name string, value string) Command {
	return setSnmpCommunity(name, "encryption-password", value)
}
func (r *Ros) SetSnmpCommunityEncryptionPassword(name string, value string) error {

	return r.Exec(setSnmpCommunityEncryptionPassword(name, value))

}
func setSnmpCommunityEncryptionProtocol(name string, value string) Command {
	return setSnmpCommunity(name, "encryption-protocol", value)
}
func (r *Ros) SetSnmpCommunityEncryptionProtocol(name string, value string) error {

	return r.Exec(setSnmpCommunityEncryptionProtocol(name, value))

}
func setSnmpCommunityAddresses(name string, value string) Command {
	return setSnmpCommunity(name, "addresses", value)
}
func (r *Ros) SetSnmpCommunityAddresses(name string, value string) error {

	return r.Exec(setSnmpCommunityAddresses(name, value))

}
func setSnmpCommunitySecurity(name string, value string) Command {
	return setSnmpCommunity(name, "security", value)
}
func (r *Ros) SetSnmpCommunitySecurity(name string, value string) error {

	return r.Exec(setSnmpCommunitySecurity(name, value))

}
func setSnmpCommunityReadAccess(name string, value bool) Command {
	return setSnmpCommunity(name, "read-access", FormatBool(value))
}
func (r *Ros) SetSnmpCommunityReadAccess(name string, value bool) error {

	return r.Exec(setSnmpCommunityReadAccess(name, value))

}
func setSnmpCommunityWriteAccess(name string, value bool) Command {
	return setSnmpCommunity(name, "write-access", FormatBool(value))
}
func (r *Ros) SetSnmpCommunityWriteAccess(name string, value bool) error {

	return r.Exec(setSnmpCommunityWriteAccess(name, value))

}
func snmpCommunityDefault() Command {
	return Command{
		Path:    "/snmp community",
		Command: "print",
		Flags: map[string]bool{
			"default": true,
		},
		Detail: false,
	}
}
func (r *Ros) SnmpCommunityDefault() (map[string]string, bool, error) {

	res, err := r.Values(snmpCommunityDefault())

	return res, true, err

}
func setSnmpCommunityDefault(key, value string) Command {
	return Command{
		Path:    "/snmp community",
		Command: "set",
		Filter: map[string]string{
			"default": "",
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setSnmpCommunityDefaultName(value string) Command {
	return setSnmpCommunityDefault("name", value)
}
func (r *Ros) SetSnmpCommunityDefaultName(value string) error {

	return r.Exec(setSnmpCommunityDefaultName(value))

}
func setSnmpCommunityDefaultAuthenticationPassword(value string) Command {
	return setSnmpCommunityDefault("authentication-password", value)
}
func (r *Ros) SetSnmpCommunityDefaultAuthenticationPassword(value string) error {

	return r.Exec(setSnmpCommunityDefaultAuthenticationPassword(value))

}
func setSnmpCommunityDefaultAuthenticationProtocol(value string) Command {
	return setSnmpCommunityDefault("authentication-protocol", value)
}
func (r *Ros) SetSnmpCommunityDefaultAuthenticationProtocol(value string) error {

	return r.Exec(setSnmpCommunityDefaultAuthenticationProtocol(value))

}
func setSnmpCommunityDefaultEncryptionPassword(value string) Command {
	return setSnmpCommunityDefault("encryption-password", value)
}
func (r *Ros) SetSnmpCommunityDefaultEncryptionPassword(value string) error {

	return r.Exec(setSnmpCommunityDefaultEncryptionPassword(value))

}
func setSnmpCommunityDefaultEncryptionProtocol(value string) Command {
	return setSnmpCommunityDefault("encryption-protocol", value)
}
func (r *Ros) SetSnmpCommunityDefaultEncryptionProtocol(value string) error {

	return r.Exec(setSnmpCommunityDefaultEncryptionProtocol(value))

}
func setSnmpCommunityDefaultAddresses(value string) Command {
	return setSnmpCommunityDefault("addresses", value)
}
func (r *Ros) SetSnmpCommunityDefaultAddresses(value string) error {

	return r.Exec(setSnmpCommunityDefaultAddresses(value))

}
func setSnmpCommunityDefaultSecurity(value string) Command {
	return setSnmpCommunityDefault("security", value)
}
func (r *Ros) SetSnmpCommunityDefaultSecurity(value string) error {

	return r.Exec(setSnmpCommunityDefaultSecurity(value))

}
func setSnmpCommunityDefaultReadAccess(value bool) Command {
	return setSnmpCommunityDefault("read-access", FormatBool(value))
}
func (r *Ros) SetSnmpCommunityDefaultReadAccess(value bool) error {

	return r.Exec(setSnmpCommunityDefaultReadAccess(value))

}
func setSnmpCommunityDefaultWriteAccess(value bool) Command {
	return setSnmpCommunityDefault("write-access", FormatBool(value))
}
func (r *Ros) SetSnmpCommunityDefaultWriteAccess(value bool) error {

	return r.Exec(setSnmpCommunityDefaultWriteAccess(value))

}

func systemClock() Command {
	return Command{
		Path:    "/system clock",
		Command: "print",
		Filter:  map[string]string{},
		Flags:   map[string]bool{},
		Detail:  false,
	}
}
func (r *Ros) SystemClock() (map[string]string, bool, error) {

	res, err := r.Values(systemClock())
	return res, true, err

}
func setSystemClock(key, value string) Command {
	return Command{
		Path:    "/system clock",
		Command: "set",
		Filter:  map[string]string{},
		Params: map[string]string{
			key: value,
		},
	}
}
func setSystemClockTimeZoneName(value string) Command {
	return setSystemClock("time-zone-name", value)
}
func (r *Ros) SetSystemClockTimeZoneName(value string) error {

	return r.Exec(setSystemClockTimeZoneName(value))

}
func setSystemClockTimeZoneAutodetect(value bool) Command {
	return setSystemClock("time-zone-autodetect", FormatBool(value))
}
func (r *Ros) SetSystemClockTimeZoneAutodetect(value bool) error {

	if r.AtLeast(6, 27) {

		return r.Exec(setSystemClockTimeZoneAutodetect(value))

	}

	return nil
}

func systemIdentity() Command {
	return Command{
		Path:    "/system identity",
		Command: "print",
		Filter:  map[string]string{},
		Flags:   map[string]bool{},
		Detail:  false,
	}
}
func (r *Ros) SystemIdentity() (map[string]string, bool, error) {

	res, err := r.Values(systemIdentity())
	return res, true, err

}
func setSystemIdentity(key, value string) Command {
	return Command{
		Path:    "/system identity",
		Command: "set",
		Filter:  map[string]string{},
		Params: map[string]string{
			key: value,
		},
	}
}
func setSystemIdentityName(value string) Command {
	return setSystemIdentity("name", value)
}
func (r *Ros) SetSystemIdentityName(value string) error {

	return r.Exec(setSystemIdentityName(value))

}

func systemLogging(action string, topics string) Command {
	return Command{
		Path:    "/system logging",
		Command: "print",
		Filter: map[string]string{
			"action": action,
			"topics": topics,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) SystemLogging(action string, topics string) (map[string]string, bool, error) {

	res, err := r.First(systemLogging(action, topics))
	return res, true, err

}
func addSystemLogging(action string, topics string) Command {
	return Command{
		Path:    "/system logging",
		Command: "add",
		Params: map[string]string{
			"action": action,
			"topics": topics,
		},
		Extra: map[string]string{},
	}
}
func (r *Ros) AddSystemLogging(action string, topics string) error {

	return r.Exec(addSystemLogging(action, topics))

}
func removeSystemLogging(action string, topics string) Command {
	return Command{
		Path:    "/system logging",
		Command: "remove",
		Filter: map[string]string{
			"action": action,
			"topics": topics,
		},
		Flags: map[string]bool{},
	}
}
func (r *Ros) RemoveSystemLogging(action string, topics string) error {

	return r.Exec(removeSystemLogging(action, topics))

}
func setSystemLogging(action string, topics string, key, value string) Command {
	return Command{
		Path:    "/system logging",
		Command: "set",
		Filter: map[string]string{
			"action": action,
			"topics": topics,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setSystemLoggingPrefix(action string, topics string, value string) Command {
	return setSystemLogging(action, topics, "prefix", value)
}
func (r *Ros) SetSystemLoggingPrefix(action string, topics string, value string) error {

	return r.Exec(setSystemLoggingPrefix(action, topics, value))

}

func systemLoggingAction(name string) Command {
	return Command{
		Path:    "/system logging action",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) SystemLoggingAction(name string) (map[string]string, bool, error) {

	res, err := r.First(systemLoggingAction(name))
	return res, true, err

}
func addSystemLoggingAction(name string) Command {
	return Command{
		Path:    "/system logging action",
		Command: "add",
		Params: map[string]string{
			"name": name,
		},
		Extra: map[string]string{},
	}
}
func (r *Ros) AddSystemLoggingAction(name string) error {

	return r.Exec(addSystemLoggingAction(name))

}
func removeSystemLoggingAction(name string) Command {
	return Command{
		Path:    "/system logging action",
		Command: "remove",
		Filter: map[string]string{
			"name": name,
		},
		Flags: map[string]bool{},
	}
}
func (r *Ros) RemoveSystemLoggingAction(name string) error {

	return r.Exec(removeSystemLoggingAction(name))

}
func setSystemLoggingAction(name string, key, value string) Command {
	return Command{
		Path:    "/system logging action",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setSystemLoggingActionTarget(name string, value string) Command {
	return setSystemLoggingAction(name, "target", value)
}
func (r *Ros) SetSystemLoggingActionTarget(name string, value string) error {

	return r.Exec(setSystemLoggingActionTarget(name, value))

}
func setSystemLoggingActionRemote(name string, value string) Command {
	return setSystemLoggingAction(name, "remote", value)
}
func (r *Ros) SetSystemLoggingActionRemote(name string, value string) error {

	return r.Exec(setSystemLoggingActionRemote(name, value))

}
func setSystemLoggingActionRemotePort(name string, value int) Command {
	return setSystemLoggingAction(name, "remote-port", FormatInt(value))
}
func (r *Ros) SetSystemLoggingActionRemotePort(name string, value int) error {

	return r.Exec(setSystemLoggingActionRemotePort(name, value))

}
func setSystemLoggingActionSrcAddress(name string, value string) Command {
	return setSystemLoggingAction(name, "src-address", value)
}
func (r *Ros) SetSystemLoggingActionSrcAddress(name string, value string) error {

	return r.Exec(setSystemLoggingActionSrcAddress(name, value))

}
func setSystemLoggingActionBsdSyslog(name string, value bool) Command {
	return setSystemLoggingAction(name, "bsd-syslog", FormatBool(value))
}
func (r *Ros) SetSystemLoggingActionBsdSyslog(name string, value bool) error {

	return r.Exec(setSystemLoggingActionBsdSyslog(name, value))

}
func setSystemLoggingActionSyslogSeverity(name string, value string) Command {
	return setSystemLoggingAction(name, "syslog-severity", value)
}
func (r *Ros) SetSystemLoggingActionSyslogSeverity(name string, value string) error {

	return r.Exec(setSystemLoggingActionSyslogSeverity(name, value))

}
func setSystemLoggingActionSyslogFacility(name string, value string) Command {
	return setSystemLoggingAction(name, "syslog-facility", value)
}
func (r *Ros) SetSystemLoggingActionSyslogFacility(name string, value string) error {

	return r.Exec(setSystemLoggingActionSyslogFacility(name, value))

}

func systemNote() Command {
	return Command{
		Path:    "/system note",
		Command: "print",
		Filter:  map[string]string{},
		Flags:   map[string]bool{},
		Detail:  false,
	}
}
func (r *Ros) SystemNote() (map[string]string, bool, error) {

	res, err := r.Values(systemNote())
	return res, true, err

}
func setSystemNote(key, value string) Command {
	return Command{
		Path:    "/system note",
		Command: "set",
		Filter:  map[string]string{},
		Params: map[string]string{
			key: value,
		},
	}
}
func setSystemNoteNote(value string) Command {
	return setSystemNote("note", value)
}
func (r *Ros) SetSystemNoteNote(value string) error {

	return r.Exec(setSystemNoteNote(value))

}
func setSystemNoteShowAtLogin(value bool) Command {
	return setSystemNote("show-at-login", FormatBool(value))
}
func (r *Ros) SetSystemNoteShowAtLogin(value bool) error {

	return r.Exec(setSystemNoteShowAtLogin(value))

}

func systemNtpClient() Command {
	return Command{
		Path:    "/system ntp client",
		Command: "print",
		Filter:  map[string]string{},
		Flags:   map[string]bool{},
		Detail:  false,
	}
}
func (r *Ros) SystemNtpClient() (map[string]string, bool, error) {

	res, err := r.Values(systemNtpClient())
	return res, true, err

}
func setSystemNtpClient(key, value string) Command {
	return Command{
		Path:    "/system ntp client",
		Command: "set",
		Filter:  map[string]string{},
		Params: map[string]string{
			key: value,
		},
	}
}
func setSystemNtpClientEnabled(value bool) Command {
	return setSystemNtpClient("enabled", FormatBool(value))
}
func (r *Ros) SetSystemNtpClientEnabled(value bool) error {

	return r.Exec(setSystemNtpClientEnabled(value))

}
func setSystemNtpClientPrimaryNtp(value string) Command {
	return setSystemNtpClient("primary-ntp", value)
}
func (r *Ros) SetSystemNtpClientPrimaryNtp(value string) error {

	return r.Exec(setSystemNtpClientPrimaryNtp(value))

}
func setSystemNtpClientSecondaryNtp(value string) Command {
	return setSystemNtpClient("secondary-ntp", value)
}
func (r *Ros) SetSystemNtpClientSecondaryNtp(value string) error {

	return r.Exec(setSystemNtpClientSecondaryNtp(value))

}

func systemResource() Command {
	return Command{
		Path:    "/system resource",
		Command: "print",
		Filter:  map[string]string{},
		Flags:   map[string]bool{},
		Detail:  false,
	}
}
func (r *Ros) SystemResource() (map[string]string, bool, error) {

	res, err := r.Values(systemResource())
	return res, true, err

}

func user(name string) Command {
	return Command{
		Path:    "/user",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) User(name string) (map[string]string, bool, error) {

	res, err := r.First(user(name))
	return res, true, err

}
func addUser(name string, group string, password string) Command {
	return Command{
		Path:    "/user",
		Command: "add",
		Params: map[string]string{
			"name": name,
		},
		Extra: map[string]string{
			"group":    group,
			"password": password,
		},
	}
}
func (r *Ros) AddUser(name string, group string, password string) error {

	return r.Exec(addUser(name, group, password))

}
func removeUser(name string) Command {
	return Command{
		Path:    "/user",
		Command: "remove",
		Filter: map[string]string{
			"name": name,
		},
		Flags: map[string]bool{},
	}
}
func (r *Ros) RemoveUser(name string) error {

	return r.Exec(removeUser(name))

}
func setUser(name string, key, value string) Command {
	return Command{
		Path:    "/user",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setUserGroup(name string, value string) Command {
	return setUser(name, "group", value)
}
func (r *Ros) SetUserGroup(name string, value string) error {

	return r.Exec(setUserGroup(name, value))

}
func setUserPassword(name string, value string) Command {
	return setUser(name, "password", value)
}
func (r *Ros) SetUserPassword(name string, value string) error {

	return r.Exec(setUserPassword(name, value))

}
func setUserComment(name string, value string) Command {
	return setUser(name, "comment", value)
}
func (r *Ros) SetUserComment(name string, value string) error {

	return r.Exec(setUserComment(name, value))

}
func setUserAddress(name string, value string) Command {
	return setUser(name, "address", value)
}
func (r *Ros) SetUserAddress(name string, value string) error {

	return r.Exec(setUserAddress(name, value))

}

func systemRouterboard() Command {
	return Command{
		Path:    "/system routerboard",
		Command: "print",
		Filter:  map[string]string{},
		Flags:   map[string]bool{},
		Detail:  false,
	}
}
func (r *Ros) SystemRouterboard() (map[string]string, bool, error) {

	res, err := r.Values(systemRouterboard())
	return res, true, err

}

func systemScheduler(name string) Command {
	return Command{
		Path:    "/system scheduler",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) SystemScheduler(name string) (map[string]string, bool, error) {

	res, err := r.First(systemScheduler(name))
	return res, true, err

}
func addSystemScheduler(name string, interval string, policy string, onEvent string) Command {
	return Command{
		Path:    "/system scheduler",
		Command: "add",
		Params: map[string]string{
			"name": name,
		},
		Extra: map[string]string{
			"interval": interval,
			"on-event": onEvent,
			"policy":   policy,
		},
	}
}
func (r *Ros) AddSystemScheduler(name string, interval string, policy string, onEvent string) error {

	return r.Exec(addSystemScheduler(name, interval, policy, onEvent))

}
func removeSystemScheduler(name string) Command {
	return Command{
		Path:    "/system scheduler",
		Command: "remove",
		Filter: map[string]string{
			"name": name,
		},
		Flags: map[string]bool{},
	}
}
func (r *Ros) RemoveSystemScheduler(name string) error {

	return r.Exec(removeSystemScheduler(name))

}
func setSystemScheduler(name string, key, value string) Command {
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
func setSystemSchedulerInterval(name string, value string) Command {
	return setSystemScheduler(name, "interval", value)
}
func (r *Ros) SetSystemSchedulerInterval(name string, value string) error {

	return r.Exec(setSystemSchedulerInterval(name, value))

}
func setSystemSchedulerPolicy(name string, value string) Command {
	return setSystemScheduler(name, "policy", value)
}
func (r *Ros) SetSystemSchedulerPolicy(name string, value string) error {

	return r.Exec(setSystemSchedulerPolicy(name, value))

}
func setSystemSchedulerOnEvent(name string, value string) Command {
	return setSystemScheduler(name, "on-event", value)
}
func (r *Ros) SetSystemSchedulerOnEvent(name string, value string) error {

	return r.Exec(setSystemSchedulerOnEvent(name, value))

}
func setSystemSchedulerComment(name string, value string) Command {
	return setSystemScheduler(name, "comment", value)
}
func (r *Ros) SetSystemSchedulerComment(name string, value string) error {

	return r.Exec(setSystemSchedulerComment(name, value))

}
func setSystemSchedulerStartDate(name string, value string) Command {
	return setSystemScheduler(name, "start-date", value)
}
func (r *Ros) SetSystemSchedulerStartDate(name string, value string) error {

	return r.Exec(setSystemSchedulerStartDate(name, value))

}
func setSystemSchedulerStartTime(name string, value string) Command {
	return setSystemScheduler(name, "start-time", value)
}
func (r *Ros) SetSystemSchedulerStartTime(name string, value string) error {

	return r.Exec(setSystemSchedulerStartTime(name, value))

}

func systemScript(name string) Command {
	return Command{
		Path:    "/system script",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) SystemScript(name string) (map[string]string, bool, error) {

	if r.AtLeast(6, 41) {

		raw, err := r.Raw(systemScript(name))
		if err != nil {
			return nil, true, err
		}
		if len(raw) > 0 {
			return raw[0], true, err
		}
		return nil, true, nil

	}

	return nil, false, nil
}
func addSystemScript(name string, owner string, policy string, source string) Command {
	return Command{
		Path:    "/system script",
		Command: "add",
		Params: map[string]string{
			"name": name,
		},
		Extra: map[string]string{
			"owner":  owner,
			"policy": policy,
			"source": source,
		},
	}
}
func (r *Ros) AddSystemScript(name string, owner string, policy string, source string) error {

	if r.AtLeast(6, 41) {

		return r.Exec(addSystemScript(name, owner, policy, source))

	}

	return nil
}
func removeSystemScript(name string) Command {
	return Command{
		Path:    "/system script",
		Command: "remove",
		Filter: map[string]string{
			"name": name,
		},
		Flags: map[string]bool{},
	}
}
func (r *Ros) RemoveSystemScript(name string) error {

	if r.AtLeast(6, 41) {

		return r.Exec(removeSystemScript(name))

	}

	return nil
}
func setSystemScript(name string, key, value string) Command {
	return Command{
		Path:    "/system script",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setSystemScriptOwner(name string, value string) Command {
	return setSystemScript(name, "owner", value)
}
func (r *Ros) SetSystemScriptOwner(name string, value string) error {

	if r.AtLeast(6, 41) {

		return r.Exec(setSystemScriptOwner(name, value))

	}

	return nil
}
func setSystemScriptPolicy(name string, value string) Command {
	return setSystemScript(name, "policy", value)
}
func (r *Ros) SetSystemScriptPolicy(name string, value string) error {

	if r.AtLeast(6, 41) {

		return r.Exec(setSystemScriptPolicy(name, value))

	}

	return nil
}
func setSystemScriptSource(name string, value string) Command {
	return setSystemScript(name, "source", value)
}
func (r *Ros) SetSystemScriptSource(name string, value string) error {

	if r.AtLeast(6, 41) {

		return r.Exec(setSystemScriptSource(name, ParseSystemScriptSource(value)))

	}

	return nil
}

func toolBandwidthServer() Command {
	return Command{
		Path:    "/tool bandwidth-server",
		Command: "print",
		Filter:  map[string]string{},
		Flags:   map[string]bool{},
		Detail:  false,
	}
}
func (r *Ros) ToolBandwidthServer() (map[string]string, bool, error) {

	res, err := r.Values(toolBandwidthServer())
	return res, true, err

}
func setToolBandwidthServer(key, value string) Command {
	return Command{
		Path:    "/tool bandwidth-server",
		Command: "set",
		Filter:  map[string]string{},
		Params: map[string]string{
			key: value,
		},
	}
}
func setToolBandwidthServerEnabled(value bool) Command {
	return setToolBandwidthServer("enabled", FormatBool(value))
}
func (r *Ros) SetToolBandwidthServerEnabled(value bool) error {

	return r.Exec(setToolBandwidthServerEnabled(value))

}
func setToolBandwidthServerAuthenticate(value bool) Command {
	return setToolBandwidthServer("authenticate", FormatBool(value))
}
func (r *Ros) SetToolBandwidthServerAuthenticate(value bool) error {

	return r.Exec(setToolBandwidthServerAuthenticate(value))

}
func setToolBandwidthServerAllocateUdpPortsFrom(value int) Command {
	return setToolBandwidthServer("allocate-udp-ports-from", FormatInt(value))
}
func (r *Ros) SetToolBandwidthServerAllocateUdpPortsFrom(value int) error {

	return r.Exec(setToolBandwidthServerAllocateUdpPortsFrom(value))

}
func setToolBandwidthServerMaxSessions(value int) Command {
	return setToolBandwidthServer("max-sessions", FormatInt(value))
}
func (r *Ros) SetToolBandwidthServerMaxSessions(value int) error {

	return r.Exec(setToolBandwidthServerMaxSessions(value))

}

func toolMacServer() Command {
	return Command{
		Path:    "/tool mac-server",
		Command: "print",
		Filter:  map[string]string{},
		Flags:   map[string]bool{},
		Detail:  false,
	}
}
func (r *Ros) ToolMacServer() (map[string]string, bool, error) {

	res, err := r.Values(toolMacServer())
	return res, true, err

}
func setToolMacServer(key, value string) Command {
	return Command{
		Path:    "/tool mac-server",
		Command: "set",
		Filter:  map[string]string{},
		Params: map[string]string{
			key: value,
		},
	}
}
func setToolMacServerAllowedInterfaceList(value string) Command {
	return setToolMacServer("allowed-interface-list", value)
}
func (r *Ros) SetToolMacServerAllowedInterfaceList(value string) error {

	if r.AtLeast(6, 41) {

		return r.Exec(setToolMacServerAllowedInterfaceList(value))

	}

	return nil
}

func toolMacServerMacWinbox() Command {
	return Command{
		Path:    "/tool mac-server mac-winbox",
		Command: "print",
		Filter:  map[string]string{},
		Flags:   map[string]bool{},
		Detail:  false,
	}
}
func (r *Ros) ToolMacServerMacWinbox() (map[string]string, bool, error) {

	res, err := r.Values(toolMacServerMacWinbox())
	return res, true, err

}
func setToolMacServerMacWinbox(key, value string) Command {
	return Command{
		Path:    "/tool mac-server mac-winbox",
		Command: "set",
		Filter:  map[string]string{},
		Params: map[string]string{
			key: value,
		},
	}
}
func setToolMacServerMacWinboxAllowedInterfaceList(value string) Command {
	return setToolMacServerMacWinbox("allowed-interface-list", value)
}
func (r *Ros) SetToolMacServerMacWinboxAllowedInterfaceList(value string) error {

	if r.AtLeast(6, 41) {

		return r.Exec(setToolMacServerMacWinboxAllowedInterfaceList(value))

	}

	return nil
}

func toolMacServerPing() Command {
	return Command{
		Path:    "/tool mac-server ping",
		Command: "print",
		Filter:  map[string]string{},
		Flags:   map[string]bool{},
		Detail:  false,
	}
}
func (r *Ros) ToolMacServerPing() (map[string]string, bool, error) {

	res, err := r.Values(toolMacServerPing())
	return res, true, err

}
func setToolMacServerPing(key, value string) Command {
	return Command{
		Path:    "/tool mac-server ping",
		Command: "set",
		Filter:  map[string]string{},
		Params: map[string]string{
			key: value,
		},
	}
}
func setToolMacServerPingEnabled(value bool) Command {
	return setToolMacServerPing("enabled", FormatBool(value))
}
func (r *Ros) SetToolMacServerPingEnabled(value bool) error {

	return r.Exec(setToolMacServerPingEnabled(value))

}

func toolNetwatch(host string) Command {
	return Command{
		Path:    "/tool netwatch",
		Command: "print",
		Filter: map[string]string{
			"host": host,
		},
		Flags:  map[string]bool{},
		Detail: true,
	}
}
func (r *Ros) ToolNetwatch(host string) (map[string]string, bool, error) {

	res, err := r.First(toolNetwatch(host))
	return res, true, err

}
func addToolNetwatch(host string, upScript string, downScript string, interval string, timeout string) Command {
	return Command{
		Path:    "/tool netwatch",
		Command: "add",
		Params: map[string]string{
			"host": host,
		},
		Extra: map[string]string{
			"down-script": downScript,
			"interval":    interval,
			"timeout":     timeout,
			"up-script":   upScript,
		},
	}
}
func (r *Ros) AddToolNetwatch(host string, upScript string, downScript string, interval string, timeout string) error {

	return r.Exec(addToolNetwatch(host, upScript, downScript, interval, timeout))

}
func removeToolNetwatch(host string) Command {
	return Command{
		Path:    "/tool netwatch",
		Command: "remove",
		Filter: map[string]string{
			"host": host,
		},
		Flags: map[string]bool{},
	}
}
func (r *Ros) RemoveToolNetwatch(host string) error {

	return r.Exec(removeToolNetwatch(host))

}
func setToolNetwatch(host string, key, value string) Command {
	return Command{
		Path:    "/tool netwatch",
		Command: "set",
		Filter: map[string]string{
			"host": host,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setToolNetwatchUpScript(host string, value string) Command {
	return setToolNetwatch(host, "up-script", value)
}
func (r *Ros) SetToolNetwatchUpScript(host string, value string) error {

	return r.Exec(setToolNetwatchUpScript(host, value))

}
func setToolNetwatchDownScript(host string, value string) Command {
	return setToolNetwatch(host, "down-script", value)
}
func (r *Ros) SetToolNetwatchDownScript(host string, value string) error {

	return r.Exec(setToolNetwatchDownScript(host, value))

}
func setToolNetwatchInterval(host string, value string) Command {
	return setToolNetwatch(host, "interval", value)
}
func (r *Ros) SetToolNetwatchInterval(host string, value string) error {

	return r.Exec(setToolNetwatchInterval(host, value))

}
func setToolNetwatchTimeout(host string, value string) Command {
	return setToolNetwatch(host, "timeout", value)
}
func (r *Ros) SetToolNetwatchTimeout(host string, value string) error {

	return r.Exec(setToolNetwatchTimeout(host, value))

}
func setToolNetwatchComment(host string, value string) Command {
	return setToolNetwatch(host, "comment", value)
}
func (r *Ros) SetToolNetwatchComment(host string, value string) error {

	return r.Exec(setToolNetwatchComment(host, value))

}
func setToolNetwatchDisabled(host string, value bool) Command {
	return setToolNetwatch(host, "disabled", FormatBool(value))
}
func (r *Ros) SetToolNetwatchDisabled(host string, value bool) error {

	return r.Exec(setToolNetwatchDisabled(host, value))

}

func toolRomon() Command {
	return Command{
		Path:    "/tool romon",
		Command: "print",
		Filter:  map[string]string{},
		Flags:   map[string]bool{},
		Detail:  false,
	}
}
func (r *Ros) ToolRomon() (map[string]string, bool, error) {

	if r.AtLeast(6, 29) {

		res, err := r.Values(toolRomon())
		return res, true, err

	}

	return nil, false, nil
}
func setToolRomon(key, value string) Command {
	return Command{
		Path:    "/tool romon",
		Command: "set",
		Filter:  map[string]string{},
		Params: map[string]string{
			key: value,
		},
	}
}
func setToolRomonEnabled(value bool) Command {
	return setToolRomon("enabled", FormatBool(value))
}
func (r *Ros) SetToolRomonEnabled(value bool) error {

	if r.AtLeast(6, 29) {

		return r.Exec(setToolRomonEnabled(value))

	}

	return nil
}
func setToolRomonId(value string) Command {
	return setToolRomon("id", value)
}
func (r *Ros) SetToolRomonId(value string) error {

	if r.AtLeast(6, 29) {

		return r.Exec(setToolRomonId(value))

	}

	return nil
}
func setToolRomonSecrets(value string) Command {
	return setToolRomon("secrets", value)
}
func (r *Ros) SetToolRomonSecrets(value string) error {

	if r.AtLeast(6, 29) {

		return r.Exec(setToolRomonSecrets(value))

	}

	return nil
}

func toolRomonPort(iface string) Command {
	return Command{
		Path:    "/tool romon port",
		Command: "print",
		Filter: map[string]string{
			"interface": iface,
		},
		Flags: map[string]bool{
			"default": false,
		},
		Detail: true,
	}
}
func (r *Ros) ToolRomonPort(iface string) (map[string]string, bool, error) {

	if r.AtLeast(6, 29) {

		res, err := r.First(toolRomonPort(iface))
		return res, true, err

	}

	return nil, false, nil
}
func addToolRomonPort(iface string) Command {
	return Command{
		Path:    "/tool romon port",
		Command: "add",
		Params: map[string]string{
			"interface": iface,
		},
		Extra: map[string]string{},
	}
}
func (r *Ros) AddToolRomonPort(iface string) error {

	if r.AtLeast(6, 29) {

		return r.Exec(addToolRomonPort(iface))

	}

	return nil
}
func removeToolRomonPort(iface string) Command {
	return Command{
		Path:    "/tool romon port",
		Command: "remove",
		Filter: map[string]string{
			"interface": iface,
		},
		Flags: map[string]bool{
			"default": false,
		},
	}
}
func (r *Ros) RemoveToolRomonPort(iface string) error {

	if r.AtLeast(6, 29) {

		return r.Exec(removeToolRomonPort(iface))

	}

	return nil
}
func setToolRomonPort(iface string, key, value string) Command {
	return Command{
		Path:    "/tool romon port",
		Command: "set",
		Filter: map[string]string{
			"interface": iface,
			"!default":  "",
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setToolRomonPortSecrets(iface string, value string) Command {
	return setToolRomonPort(iface, "secrets", value)
}
func (r *Ros) SetToolRomonPortSecrets(iface string, value string) error {

	if r.AtLeast(6, 29) {

		return r.Exec(setToolRomonPortSecrets(iface, value))

	}

	return nil
}
func setToolRomonPortCost(iface string, value int) Command {
	return setToolRomonPort(iface, "cost", FormatInt(value))
}
func (r *Ros) SetToolRomonPortCost(iface string, value int) error {

	if r.AtLeast(6, 29) {

		return r.Exec(setToolRomonPortCost(iface, value))

	}

	return nil
}
func setToolRomonPortDisabled(iface string, value bool) Command {
	return setToolRomonPort(iface, "disabled", FormatBool(value))
}
func (r *Ros) SetToolRomonPortDisabled(iface string, value bool) error {

	if r.AtLeast(6, 29) {

		return r.Exec(setToolRomonPortDisabled(iface, value))

	}

	return nil
}
func setToolRomonPortForbid(iface string, value bool) Command {
	return setToolRomonPort(iface, "forbid", FormatBool(value))
}
func (r *Ros) SetToolRomonPortForbid(iface string, value bool) error {

	if r.AtLeast(6, 29) {

		return r.Exec(setToolRomonPortForbid(iface, value))

	}

	return nil
}
func toolRomonPortDefault() Command {
	return Command{
		Path:    "/tool romon port",
		Command: "print",
		Flags: map[string]bool{
			"default": true,
		},
		Detail: true,
	}
}
func (r *Ros) ToolRomonPortDefault() (map[string]string, bool, error) {

	if r.AtLeast(6, 29) {

		res, err := r.Values(toolRomonPortDefault())

		return res, true, err

	}

	return nil, false, nil
}
func setToolRomonPortDefault(key, value string) Command {
	return Command{
		Path:    "/tool romon port",
		Command: "set",
		Filter: map[string]string{
			"default": "",
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func setToolRomonPortDefaultInterface(value string) Command {
	return setToolRomonPortDefault("interface", value)
}
func (r *Ros) SetToolRomonPortDefaultInterface(value string) error {

	if r.AtLeast(6, 29) {

		return r.Exec(setToolRomonPortDefaultInterface(value))

	}

	return nil
}
func setToolRomonPortDefaultSecrets(value string) Command {
	return setToolRomonPortDefault("secrets", value)
}
func (r *Ros) SetToolRomonPortDefaultSecrets(value string) error {

	if r.AtLeast(6, 29) {

		return r.Exec(setToolRomonPortDefaultSecrets(value))

	}

	return nil
}
func setToolRomonPortDefaultCost(value int) Command {
	return setToolRomonPortDefault("cost", FormatInt(value))
}
func (r *Ros) SetToolRomonPortDefaultCost(value int) error {

	if r.AtLeast(6, 29) {

		return r.Exec(setToolRomonPortDefaultCost(value))

	}

	return nil
}
func setToolRomonPortDefaultDisabled(value bool) Command {
	return setToolRomonPortDefault("disabled", FormatBool(value))
}
func (r *Ros) SetToolRomonPortDefaultDisabled(value bool) error {

	if r.AtLeast(6, 29) {

		return r.Exec(setToolRomonPortDefaultDisabled(value))

	}

	return nil
}
func setToolRomonPortDefaultForbid(value bool) Command {
	return setToolRomonPortDefault("forbid", FormatBool(value))
}
func (r *Ros) SetToolRomonPortDefaultForbid(value bool) error {

	if r.AtLeast(6, 29) {

		return r.Exec(setToolRomonPortDefaultForbid(value))

	}

	return nil
}

func userAaa() Command {
	return Command{
		Path:    "/user aaa",
		Command: "print",
		Filter:  map[string]string{},
		Flags:   map[string]bool{},
		Detail:  false,
	}
}
func (r *Ros) UserAaa() (map[string]string, bool, error) {

	res, err := r.Values(userAaa())
	return res, true, err

}
func setUserAaa(key, value string) Command {
	return Command{
		Path:    "/user aaa",
		Command: "set",
		Filter:  map[string]string{},
		Params: map[string]string{
			key: value,
		},
	}
}
func setUserAaaUseRadius(value bool) Command {
	return setUserAaa("use-radius", FormatBool(value))
}
func (r *Ros) SetUserAaaUseRadius(value bool) error {

	return r.Exec(setUserAaaUseRadius(value))

}
func setUserAaaAccounting(value bool) Command {
	return setUserAaa("accounting", FormatBool(value))
}
func (r *Ros) SetUserAaaAccounting(value bool) error {

	return r.Exec(setUserAaaAccounting(value))

}
func setUserAaaInterimUpdate(value string) Command {
	return setUserAaa("interim-update", value)
}
func (r *Ros) SetUserAaaInterimUpdate(value string) error {

	return r.Exec(setUserAaaInterimUpdate(value))

}
func setUserAaaDefaultGroup(value string) Command {
	return setUserAaa("default-group", value)
}
func (r *Ros) SetUserAaaDefaultGroup(value string) error {

	return r.Exec(setUserAaaDefaultGroup(value))

}
func setUserAaaExcludeGroups(value string) Command {
	return setUserAaa("exclude-groups", value)
}
func (r *Ros) SetUserAaaExcludeGroups(value string) error {

	return r.Exec(setUserAaaExcludeGroups(value))

}
