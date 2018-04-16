package ros

import (
	"testing"
)

func TestIpFirewallAddressList(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{ipFirewallAddressList("address", "list"), `/ip firewall address-list print detail where address="address" list="list"`},
		{addIpFirewallAddressList("address", "list", map[string]string{"comment": "value"}), `:if ([:len [/ip firewall address-list find address="address" list="list"]] = 0) do={/ip firewall address-list add address="address" list="list" comment="value"}`},
		{removeIpFirewallAddressList("address", "list"), `/ip firewall address-list remove [find address="address" list="list"]`},
		{setIpFirewallAddressList("address", "list", map[string]string{"comment": "value"}), `/ip firewall address-list set [find address="address" list="list"] comment="value"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("ip firewall address-list mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
