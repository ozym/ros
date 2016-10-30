package ros

import (
	"testing"
)

func TestSNMP(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{snmp(), `/snmp print`},
		{setSNMP("key", "value"), `/snmp set key="value"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("snmp mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
