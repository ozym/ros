package ros

import (
	"testing"
)

func TestCommand_Add(t *testing.T) {

	tests := []struct {
		c Command
		s string
	}{
		{
			Command{
				Command: "add",
				Path:    "/user",
				Params: map[string]string{
					"name":  "name",
					"group": "read",
				},
			},
			`:if (:len [/user find group="read" name="name"] = 0) do={/user add group="read" name="name"}`,
		},
	}

	for _, x := range tests {
		v, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if string(v) != x.s {
			t.Errorf("command mismatch: got '%s', expected '%s'", string(v), x.s)
		}
	}
}

func TestCommand_Remove(t *testing.T) {

	tests := []struct {
		c Command
		s string
	}{
		{
			Command{
				Command: "remove",
				Path:    "/user",
				Filter: map[string]string{
					"name":  "name",
					"group": "read",
				},
			},
			`/user remove [find group="read" name="name"]`,
		},
		{
			Command{
				Command: "remove",
				Path:    "/user",
				UParam:  &[]string{"1"}[0],
			},
			`/user remove 1`,
		},
	}

	for _, x := range tests {
		v, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if string(v) != x.s {
			t.Errorf("command mismatch: got '%s', expected '%s'", string(v), x.s)
		}
	}
}

func TestCommand_Set(t *testing.T) {

	tests := []struct {
		c Command
		s string
	}{
		{
			Command{
				Command: "set",
				Path:    "/user",
				Filter: map[string]string{
					"name": "name",
				},
				Params: map[string]string{
					"group": "read",
				},
			},
			`/user set [find name="name"] group="read"`,
		},
		{
			Command{
				Command: "set",
				Path:    "/user",
				Filter: map[string]string{
					"name": "name",
				},
				Params: map[string]string{
					"group": "read",
				},
				Flags: map[string]bool{
					"on":  true,
					"off": false,
				},
			},
			`/user set [find name="name"] group="read" !off on`,
		},
		{
			Command{
				Command: "set",
				Path:    "/user",
				UParam:  &[]string{"10"}[0],
				Params: map[string]string{
					"group": "read",
				},
				Flags: map[string]bool{
					"on":  true,
					"off": false,
				},
			},
			`/user set 10 group="read" !off on`,
		},
	}

	for _, x := range tests {
		v, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if string(v) != x.s {
			t.Errorf("command mismatch: got '%s', expected '%s'", string(v), x.s)
		}
	}
}

func TestCommand_Print(t *testing.T) {

	tests := []struct {
		c Command
		s string
	}{
		{
			Command{
				Command: "print",
				Path:    "/user",
			},
			`/user print`,
		},
		{
			Command{
				Command: "print",
				Path:    "/user",
				Filter: map[string]string{
					"name": "name",
				},
			},
			`/user print where name="name"`,
		},
		{
			Command{
				Command: "print",
				Path:    "/user",
				Filter: map[string]string{
					"name": "name",
				},
				Flags: map[string]bool{
					"on":  true,
					"off": false,
				},
			},
			`/user print where name="name" !off on`,
		},
	}

	for _, x := range tests {
		v, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if string(v) != x.s {
			t.Errorf("command mismatch: got '%s', expected '%s'", string(v), x.s)
		}
	}
}

func TestCommand_Detail(t *testing.T) {

	tests := []struct {
		c Command
		s string
	}{
		{
			Command{
				Command: "print",
				Detail:  true,
				Path:    "/user",
			},
			`/user print detail`,
		},
		{
			Command{
				Command: "print",
				Detail:  true,
				Path:    "/user",
				Filter: map[string]string{
					"name": "name",
				},
			},
			`/user print detail where name="name"`,
		},
		{
			Command{
				Command: "print",
				Detail:  true,
				Path:    "/user",
				Filter: map[string]string{
					"name": "name",
				},
				Flags: map[string]bool{
					"on":  true,
					"off": false,
				},
			},
			`/user print detail where name="name" !off on`,
		},
	}

	for _, x := range tests {
		v, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if string(v) != x.s {
			t.Errorf("command mismatch: got '%s', expected '%s'", string(v), x.s)
		}
	}
}
