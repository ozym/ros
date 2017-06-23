package ros

import (
	"testing"
)

func TestSystemIdentity(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{systemIdentity(), `/system identity print`},
		{setSystemIdentityName("test"), `/system identity set name="test"`},
		{setSystemIdentityName(""), `/system identity set name=""`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("system identity mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
