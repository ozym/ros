package ros

import (
	"reflect"
	"testing"
)

func TestNumbered(t *testing.T) {

	tests := []struct {
		l string
		m []map[string]string
	}{
		{
			`Flags: X - disabled, I - invalid, D - dynamic
		 0   address=10.242.0.41/28 network=10.242.0.32 interface=ether1 actual-interface=bridge1
		  `,

			[]map[string]string{
				map[string]string{
					"number":           "0",
					"address":          "10.242.0.41/28",
					"network":          "10.242.0.32",
					"interface":        "ether1",
					"actual-interface": "bridge1",
					"disabled":         "no",
					"invalid":          "no",
					"dynamic":          "no",
					"comment":          "",
				},
			},
		},
	}

	for _, test := range tests {
		m, err := ScanNumberedItemList(test.l)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < len(test.m) && i < len(m); i++ {
			if !reflect.DeepEqual(test.m[i], m[i]) {
				for k, _ := range test.m[i] {
					if _, ok := m[i][k]; !ok {
						t.Errorf("extra: %s", k)
					}
				}
				for k, _ := range m[i] {
					if _, ok := test.m[i][k]; !ok {
						t.Errorf("missing: %s", k)
					}
				}
				for k, v1 := range test.m[i] {
					if v2, ok := m[i][k]; ok {
						if v1 != v2 {
							t.Errorf("mismatch: %s: \"%s\" != \"%s\"", k, v1, v2)
						}
					}
				}

				t.Errorf("not equal: %q vs %q", test.m[i], m[i])
			}
		}

		if !reflect.DeepEqual(test.m, m) {
			t.Errorf("mismatch: %v != %s", test.m, m)
		}
	}
}
