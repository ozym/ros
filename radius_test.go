package ros

import (
	"testing"
)

func TestRadius(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{radius("login"), `/radius print detail where service="login"`},
		{addRadius("login"), `:if ([:len [/radius find service="login"]] = 0) do={/radius add service="login"}`},
		{setRadiusAddress("login", "192.168.0.1"), `/radius set [find service="login"] address="192.168.0.1"`},
		{setRadiusSecret("login", "secret"), `/radius set [find service="login"] secret="secret"`},
		{setRadiusSrcAddress("login", "192.168.0.1"), `/radius set [find service="login"] src-address="192.168.0.1"`},
		{setRadiusTimeout("login", "100ms"), `/radius set [find service="login"] timeout="100ms"`},
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
