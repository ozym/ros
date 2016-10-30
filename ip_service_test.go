package ros

import (
	"testing"
)

func TestIPService(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{ipService("name"), `/ip service print detail where name="name"`},
		{setIPService("name", "key", "value"), `/ip service set [find name="name"] key="value"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("ip service mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
