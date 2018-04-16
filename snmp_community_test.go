package ros

import (
	"testing"
)

func TestSnmpCommunity(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{snmpCommunity("name"), `/snmp community print where name="name" !default`},
		{setSnmpCommunity("name", "key", "value"), `/snmp community set [find name="name" !default] key="value"`},
		{addSnmpCommunity("name"), `:if ([:len [/snmp community find name="name"]] = 0) do={/snmp community add name="name"}`},
		{removeSnmpCommunity("name"), `/snmp community remove [find name="name" !default]`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("snmp community mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}

func TestSnmpCommunityDefault(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{snmpCommunityDefault(), `/snmp community print where default`},
		{setSnmpCommunityDefault("key", "value"), `/snmp community set [find default] key="value"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("snmp community mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
