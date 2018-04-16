package ros

import (
	"testing"
)

func TestSystemRouterboard(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{systemRouterboard(), `/system routerboard print`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("system routerboard mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
