package ros

import (
	"testing"
)

func TestIpFirewallNat(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{ipFirewallNat(map[string]string{"key": "value"}), `/ip firewall nat print detail where key="value"`},
		{addIpFirewallNat(map[string]string{"key": "value"}), `:if ([:len [/ip firewall nat find key="value"]] = 0) do={/ip firewall nat add key="value"}`},
		{removeIpFirewallNat(map[string]string{"key": "value"}), `/ip firewall nat remove [find key="value"]`},
		{setIpFirewallNat(map[string]string{"key": "value"}, map[string]string{"set": "to"}), `/ip firewall nat set [find key="value"] set="to"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("ip firewall nat mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
