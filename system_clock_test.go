package ros

import (
	"testing"
)

func TestSystemClock(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{systemClock(), `/system clock print`},
		{setSystemClockTimeZoneName("zone"), `/system clock set time-zone-name="zone"`},
		{setSystemClockTimeZoneAutodetect(true), `/system clock set time-zone-autodetect="yes"`},
		{setSystemClockTimeZoneAutodetect(false), `/system clock set time-zone-autodetect="no"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("system note mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
