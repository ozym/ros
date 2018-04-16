package ros

import (
	"testing"
)

func TestToolRomonPortDefault(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{toolRomonPortDefault(true), `/romon port print detail where default`},
		{toolRomonPortDefault(false), `/tool romon port print detail where default`},
		{setToolRomonPortDefault("key", "value", true), `/romon port set [find default] key="value"`},
		{setToolRomonPortDefault("key", "value", false), `/tool romon port set [find default] key="value"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("tool romon port default mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
func TestToolRomonPort(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{toolRomonPort("iface", true), `/romon port print detail where interface="iface" !default`},
		{toolRomonPort("iface", false), `/tool romon port print detail where interface="iface" !default`},
		{addToolRomonPort("iface", false), `:if ([:len [/tool romon port find interface="iface"]] = 0) do={/tool romon port add interface="iface"}`},
		{removeToolRomonPort("iface", false), `/tool romon port remove [find interface="iface" !default]`},
		{setToolRomonPort("iface", "key", "value", true), `/romon port set [find interface="iface" !default] key="value"`},
		{setToolRomonPort("iface", "key", "value", false), `/tool romon port set [find interface="iface" !default] key="value"`},
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
