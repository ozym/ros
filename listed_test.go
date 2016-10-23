package ros

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestListed(t *testing.T) {

	raw, err := ioutil.ReadFile("testdata/listed.prt")
	if err != nil {
		t.Fatal(err)
	}
	list, err := scanItems(string(raw))
	if err != nil {
		t.Fatal(err)
	}

	ans, err := ioutil.ReadFile("testdata/listed.json")
	if err != nil {
		t.Fatal(err)
	}
	var res map[string]string
	err = json.Unmarshal(ans, &res)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(list, res) {
		t.Errorf("%v", list)
	}

}
