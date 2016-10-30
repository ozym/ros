package ros

import (
	"testing"
)

func TestSystemUser(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{user("name"), `/user print detail where name="name"`},
		{addUser("name", "group", "password"), `:if ([:len [/user find group="group" name="name"]] = 0) do={/user add group="group" name="name" password="password"}`},
		{removeUser("name"), `/user remove [find name="name"]`},
		{setUser("name", "key", "value"), `/user set [find name="name"] key="value"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("tool romon port mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
