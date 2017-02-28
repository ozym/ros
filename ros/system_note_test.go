package ros

import (
	"testing"
)

func TestSystemNote(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{systemNote(), `/system note print`},
		{setSystemNote("test"), `/system note set note="test"`},
		{setSystemNote(""), `/system note set note=""`},
		{setSystemNoteShowAtLogin(true), `/system note set show-at-login="yes"`},
		{setSystemNoteShowAtLogin(false), `/system note set show-at-login="no"`},
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
