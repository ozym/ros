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
		{
			s: `                   uptime: 52w2d4m10s
                  version: "4.5"
              free-memory: 46032kB
             total-memory: 62432kB
                      cpu: "MIPS 24K V7.4"
                cpu-count: 1
            cpu-frequency: 680MHz
                 cpu-load: 1
           free-hdd-space: 90776kB
          total-hdd-space: 126976kB
  write-sect-since-reboot: 90739
         write-sect-total: 90739
               bad-blocks: 2
        architecture-name: "mipsbe"
               board-name: "RB411UAHR"
`,
			m: map[string]string{
				"uptime":                  "52w2d4m10s",
				"version":                 "4.5",
				"free-memory":             "46032kB",
				"total-memory":            "62432kB",
				"cpu":                     "MIPS 24K V7.4",
				"cpu-count":               "1",
				"cpu-frequency":           "680MHz",
				"cpu-load":                "1",
				"free-hdd-space":          "90776kB",
				"total-hdd-space":         "126976kB",
				"write-sect-since-reboot": "90739",
				"write-sect-total":        "90739",
				"bad-blocks":              "2",
				"architecture-name":       "mipsbe",
				"board-name":              "RB411UAHR",
			},
		},
		{
			s: `            time: 20:45:55
            date: dec/04/2016
  time-zone-name: Pacific/Auckland
      gmt-offset: +13:00
      dst-active: yes
`,
			m: map[string]string{
				"time":           "20:45:55",
				"date":           "dec/04/2016",
				"time-zone-name": "Pacific/Auckland",
				"gmt-offset":     "+13:00",
				"dst-active":     "yes",
			},
		},
		{
			s: `                   uptime: 2w6d6h54m56s
                  version: 6.34.1 (stable)
               build-time: Feb/02/2016 14:08:42
              free-memory: 41.5MiB
             total-memory: 64.0MiB
                      cpu: MIPS 24Kc V7.4
                cpu-count: 1
            cpu-frequency: 650MHz
                 cpu-load: 0%
           free-hdd-space: 4.9MiB
          total-hdd-space: 16.0MiB
  write-sect-since-reboot: 147004
         write-sect-total: 147578
               bad-blocks: 0%
        architecture-name: mipsbe
               board-name: hEX PoE lite
                 platform: MikroTik
`,
			m: map[string]string{
				"uptime":                  "2w6d6h54m56s",
				"version":                 "6.34.1 (stable)",
				"build-time":              "Feb/02/2016 14:08:42",
				"free-memory":             "41.5MiB",
				"total-memory":            "64.0MiB",
				"cpu":                     "MIPS 24Kc V7.4",
				"cpu-count":               "1",
				"cpu-frequency":           "650MHz",
				"cpu-load":                "0%",
				"free-hdd-space":          "4.9MiB",
				"total-hdd-space":         "16.0MiB",
				"write-sect-since-reboot": "147004",
				"write-sect-total":        "147578",
				"bad-blocks":              "0%",
				"architecture-name":       "mipsbe",
				"board-name":              "hEX PoE lite",
				"platform":                "MikroTik",
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
