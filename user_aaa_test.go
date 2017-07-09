package ros

import (
	"testing"
)

func TestUserAaa(t *testing.T) {

	tests := []struct {
		c Command
		v string
	}{
		{userAaa(), `/user aaa print`},
		{setUserAaaUseRadius(true), `/user aaa set use-radius="yes"`},
		{setUserAaaUseRadius(false), `/user aaa set use-radius="no"`},
		{setUserAaaAccounting(true), `/user aaa set accounting="yes"`},
		{setUserAaaAccounting(false), `/user aaa set accounting="no"`},
		{setUserAaaInterimUpdate("0s"), `/user aaa set interim-update="0s"`},
		{setUserAaaDefaultGroup("read"), `/user aaa set default-group="read"`},
		{setUserAaaExcludeGroups("groups"), `/user aaa set exclude-groups="groups"`},
	}

	for _, x := range tests {
		r, err := x.c.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if r != x.v {
			t.Errorf("user aaa mismatch: got '%s', expected '%s'", r, x.v)
		}
	}
}
