package ros

import (
	"testing"
)

func TestRadius(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{radius(), `/radius print`},
		{setRadiusAddress("192.168.0.1"), `/radius set address="192.168.0.1"`},
		{setRadiusSecret("secret"), `/radius set secret="secret"`},
		{setRadiusService("login"), `/radius set service="login"`},
		{setRadiusSrcAddress("192.168.0.1"), `/radius set src-address="192.168.0.1"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("radius mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
