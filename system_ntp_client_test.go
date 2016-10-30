package ros

import (
	"testing"
)

func TestSystemNTPClient(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{systemNTPClient(), `/system ntp client print`},
		{setSystemNTPClientEnabled(true), `/system ntp client set enabled="yes"`},
		{setSystemNTPClientEnabled(false), `/system ntp client set enabled="no"`},
		{setSystemNTPClientPrimaryNTP(""), `/system ntp client set primary-ntp=""`},
		{setSystemNTPClientPrimaryNTP("host"), `/system ntp client set primary-ntp="host"`},
		{setSystemNTPClientSecondaryNTP(""), `/system ntp client set secondary-ntp=""`},
		{setSystemNTPClientSecondaryNTP("host"), `/system ntp client set secondary-ntp="host"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("system ntp client mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
