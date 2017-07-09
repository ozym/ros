package ros

import (
	"testing"
)

func TestSystemNtpClient(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{systemNtpClient(), `/system ntp client print`},
		{setSystemNtpClientEnabled(true), `/system ntp client set enabled="yes"`},
		{setSystemNtpClientEnabled(false), `/system ntp client set enabled="no"`},
		{setSystemNtpClientPrimaryNtp(""), `/system ntp client set primary-ntp=""`},
		{setSystemNtpClientPrimaryNtp("host"), `/system ntp client set primary-ntp="host"`},
		{setSystemNtpClientSecondaryNtp(""), `/system ntp client set secondary-ntp=""`},
		{setSystemNtpClientSecondaryNtp("host"), `/system ntp client set secondary-ntp="host"`},
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
