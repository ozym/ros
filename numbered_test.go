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
		{
			`Flags: X - disabled
			 0   ;;; Monitor Link Between Cotton and Climie
			      host=10.242.0.17 timeout=998ms interval=10s since=nov/13/2016 11:37:06 status=up up-script=lk2cotton-climie_up down-script=lk2cotton-climie_down
			`,
			[]map[string]string{
				map[string]string{
					"number":      "0",
					"comment":     "Monitor Link Between Cotton and Climie",
					"host":        "10.242.0.17",
					"timeout":     "998ms",
					"interval":    "10s",
					"status":      "up",
					"up-script":   "lk2cotton-climie_up",
					"down-script": "lk2cotton-climie_down",
					"since":       "nov/13/2016 11:37:06",
					"disabled":    "no",
				},
			},
		},
		{
			`Flags: I - invalid
			 0   name="lk2cotton-climie_up" owner="admin+cte" policy=ftp,reboot,read,write,policy,test,winbox,password,sniff,sensitive,api last-started=nov/13/2016 12:14:19
			      run-count=6 source={/ip address set [find address="10.54.242.1/28" ] disabled=no}
			`,
			[]map[string]string{
				map[string]string{
					"number":       "0",
					"name":         "lk2cotton-climie_up",
					"owner":        "admin+cte",
					"policy":       "ftp,reboot,read,write,policy,test,winbox,password,sniff,sensitive,api",
					"last-started": "nov/13/2016 12:14:19",
					"run-count":    "6",
					"source":       "{/ip address set [find address=\"10.54.242.1/28\" ] disabled=no}",
					"comment":      "",
					"invalid":      "no",
				},
			},
		},
		{
			`Flags: I - invalid
 0   name="lk2cotton-climie_up" owner="admin+cte"
     policy=ftp,reboot,read,write,policy,test,winbox,password,sniff,sensitive,
       api
     last-started=nov/13/2016 12:14:19 run-count=6
     source={/ip address set [find address="10.54.242.1/28" ] disabled=no}

`,

			[]map[string]string{
				map[string]string{
					"number":       "0",
					"name":         "lk2cotton-climie_up",
					"owner":        "admin+cte",
					"policy":       "ftp,reboot,read,write,policy,test,winbox,password,sniff,sensitive,api",
					"last-started": "nov/13/2016 12:14:19",
					"run-count":    "6",
					"source":       "{/ip address set [find address=\"10.54.242.1/28\" ] disabled=no}",
					"comment":      "",
					"invalid":      "no",
				},
			},
		},
		{
			`Flags: D - dynamic, X - disabled, R - running, S - slave
0   S name="ether1" default-name="ether1" type="ether" mtu=1500 actual-mtu=1500 l2mtu=1598 max-l2mtu=2028 mac-address=4C:5E:0C:18:C1:4D fast-path=yes
	last-link-down-time=nov/18/2016 02:30:07 last-link-up-time=nov/18/2016 02:23:19 link-downs=10

1  RS name="ether2" default-name="ether2" type="ether" mtu=1500 actual-mtu=1500 l2mtu=1598 max-l2mtu=2028 mac-address=4C:5E:0C:18:C1:4E fast-path=yes
	last-link-down-time=nov/18/2016 04:03:47 last-link-up-time=nov/18/2016 04:03:49 link-downs=3
			
55  R  name="switch" type="bridge" mtu=auto actual-mtu=1500 l2mtu=1598 mac-address=4C:5E:0C:18:C1:4D fast-path=yes last-link-up-time=nov/18/2016 01:53:02 link-downs=0

`,
			[]map[string]string{
				map[string]string{
					"comment":             "",
					"number":              "0",
					"dynamic":             "no",
					"disabled":            "no",
					"slave":               "yes",
					"running":             "no",
					"name":                "ether1",
					"default-name":        "ether1",
					"type":                "ether",
					"mtu":                 "1500",
					"actual-mtu":          "1500",
					"l2mtu":               "1598",
					"max-l2mtu":           "2028",
					"mac-address":         "4C:5E:0C:18:C1:4D",
					"fast-path":           "yes",
					"last-link-down-time": "nov/18/2016 02:30:07",
					"last-link-up-time":   "nov/18/2016 02:23:19",
					"link-downs":          "10",
				},
				map[string]string{
					"comment":             "",
					"number":              "1",
					"dynamic":             "no",
					"disabled":            "no",
					"slave":               "yes",
					"running":             "yes",
					"name":                "ether2",
					"default-name":        "ether2",
					"type":                "ether",
					"mtu":                 "1500",
					"actual-mtu":          "1500",
					"l2mtu":               "1598",
					"max-l2mtu":           "2028",
					"mac-address":         "4C:5E:0C:18:C1:4E",
					"fast-path":           "yes",
					"last-link-down-time": "nov/18/2016 04:03:47",
					"last-link-up-time":   "nov/18/2016 04:03:49",
					"link-downs":          "3",
				},
				map[string]string{
					"comment":           "",
					"number":            "55",
					"dynamic":           "no",
					"disabled":          "no",
					"slave":             "no",
					"running":           "yes",
					"name":              "switch",
					"type":              "bridge",
					"mtu":               "auto",
					"actual-mtu":        "1500",
					"l2mtu":             "1598",
					"mac-address":       "4C:5E:0C:18:C1:4D",
					"fast-path":         "yes",
					"last-link-up-time": "nov/18/2016 01:53:02",
					"link-downs":        "0",
				},
			},
		},
		{`Flags: X - disabled, I - invalid, D - dynamic
		 0   address=192.168.80.114/28 network=192.168.80.112 interface=switch actual-interface=switch
		 `,
			[]map[string]string{
				map[string]string{
					"comment":          "",
					"invalid":          "no",
					"dynamic":          "no",
					"disabled":         "no",
					"number":           "0",
					"address":          "192.168.80.114/28",
					"network":          "192.168.80.112",
					"interface":        "switch",
					"actual-interface": "switch",
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
