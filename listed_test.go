package ros

import (
	"reflect"
	"testing"
)

func TestListed(t *testing.T) {

	tests := []struct {
		s string
		m map[string]string
	}{
		{
			s: `                   uptime: 15w6d1h30m32s
                  version: 5.20
              free-memory: 101096KiB
             total-memory: 127160KiB
                      cpu: MIPS 74Kc V4.12
                cpu-count: 1
            cpu-frequency: 600MHz
                 cpu-load: 6%
           free-hdd-space: 29820KiB
          total-hdd-space: 61440KiB
  write-sect-since-reboot: 1026778
         write-sect-total: 29326944
               bad-blocks: 0%
        architecture-name: mipsbe
               board-name: RB2011UAS
                 platform: MikroTik
`,
			m: map[string]string{
				"uptime":                  "15w6d1h30m32s",
				"version":                 "5.20",
				"free-memory":             "101096KiB",
				"total-memory":            "127160KiB",
				"cpu":                     "MIPS 74Kc V4.12",
				"cpu-count":               "1",
				"cpu-frequency":           "600MHz",
				"cpu-load":                "6%",
				"free-hdd-space":          "29820KiB",
				"total-hdd-space":         "61440KiB",
				"write-sect-since-reboot": "1026778",
				"write-sect-total":        "29326944",
				"bad-blocks":              "0%",
				"architecture-name":       "mipsbe",
				"board-name":              "RB2011UAS",
				"platform":                "MikroTik",
			},
		},
		{
			s: `       routerboard: yes
             model: 2011LS
     serial-number: 39A50206C6B8
     firmware-type: ar9344
  current-firmware: 3.24
  upgrade-firmware: 3.24
  `,
			m: map[string]string{
				"routerboard":      "yes",
				"model":            "2011LS",
				"serial-number":    "39A50206C6B8",
				"firmware-type":    "ar9344",
				"current-firmware": "3.24",
				"upgrade-firmware": "3.24",
			},
		},
	}

	for _, x := range tests {
		list, err := ScanItems(x.s)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(list, x.m) {
			t.Errorf("not equal: %v != %v", list, x.m)
		}
	}
}
