package ros

import (
	"testing"
)

func TestToolRomonPort(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{toolRomonPort("iface", true), `/romon port print detail where interface="iface"`},
		{toolRomonPort("iface", false), `/tool romon port print detail where interface="iface"`},
		{addToolRomonPort("iface", true), `:if (:len [/romon port find interface="iface"] = 0) do={/romon port add interface="iface"}`},
		{addToolRomonPort("iface", false), `:if (:len [/tool romon port find interface="iface"] = 0) do={/tool romon port add interface="iface"}`},
		{removeToolRomonPort("iface", true), `/romon port remove [find interface="iface"]`},
		{removeToolRomonPort("iface", false), `/tool romon port remove [find interface="iface"]`},
		{setToolRomonPort("iface", "key", "value", true), `/romon port set [find interface="iface"] key="value"`},
		{setToolRomonPort("iface", "key", "value", false), `/tool romon port set [find interface="iface"] key="value"`},
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
