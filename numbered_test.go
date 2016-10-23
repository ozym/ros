package ros

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestNumbered(t *testing.T) {

	raw, err := ioutil.ReadFile("testdata/numbered.prt")
	if err != nil {
		t.Fatal(err)
	}
	list, err := scanNumberedItemList(string(raw))
	if err != nil {
		t.Fatal(err)
	}

	ans, err := ioutil.ReadFile("testdata/numbered.json")
	if err != nil {
		t.Fatal(err)
	}
	res := make([]map[string]string, 1)
	err = json.Unmarshal(ans, &res)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(list) && i < len(res); i++ {
		if !reflect.DeepEqual(list[i], res[i]) {
			for k, _ := range list[i] {
				if _, ok := res[i][k]; !ok {
					t.Errorf("extra: %s", k)
				}
			}
			for k, _ := range res[i] {
				if _, ok := list[i][k]; !ok {
					t.Errorf("missing: %s", k)
				}
			}
			for k, v1 := range list[i] {
				if v2, ok := res[i][k]; ok {
					if v1 != v2 {
						t.Errorf("mismatch: %s: \"%s\" != \"%s\"", k, v1, v2)
					}
				}
			}

			t.Errorf("not equal: %q vs %q", list[i], res[i])
		}
	}

	if !reflect.DeepEqual(list, res) {
		t.Errorf("not equal")
	}
}
