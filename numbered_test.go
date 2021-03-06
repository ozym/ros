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
				{
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
			`Flags: I - invalid
			 0   name="lk2hillroad-climie_up" owner="admin+cte" policy=ftp,reboot,read,write,policy,test,password,sniff,sensitive run-count=0 source={/ip address set [find address="10.54.242.2/28" ] disabled=no}

			  1   name="lk2hillroad-climie_down" owner="admin+cte" policy=ftp,reboot,read,write,policy,test,password,sniff,sensitive run-count=0 source={/ip address set [find address="10.54.242.2/28" ] disabled=yes}
			  `,
			[]map[string]string{
				{
					"number":    "0",
					"name":      "lk2hillroad-climie_up",
					"owner":     "admin+cte",
					"policy":    "ftp,reboot,read,write,policy,test,password,sniff,sensitive",
					"run-count": "0",
					"source":    "{/ip address set [find address=\"10.54.242.2/28\" ] disabled=no}",
					"comment":   "",
					"invalid":   "no",
				},
				{
					"number":    "1",
					"name":      "lk2hillroad-climie_down",
					"owner":     "admin+cte",
					"policy":    "ftp,reboot,read,write,policy,test,password,sniff,sensitive",
					"run-count": "0",
					"source":    "{/ip address set [find address=\"10.54.242.2/28\" ] disabled=yes}",
					"comment":   "",
					"invalid":   "no",
				},
			},
		},
		{
			`Flags: X - disabled
			 0   ;;; Monitor Link Between Cotton and Climie
			      host=10.242.0.17 timeout=998ms interval=10s since=nov/13/2016 11:37:06 status=up up-script=lk2cotton-climie_up down-script=lk2cotton-climie_down
			`,
			[]map[string]string{
				{
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
				{
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
				{
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
				{
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
				{
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
				{
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
				{
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
		{`Flags: X - disabled, * - default
 0  * name="default" router-id=10.242.0.2 distribute-default=never redistribute-connected=as-type-1 redistribute-static=no redistribute-rip=no
      redistribute-bgp=no redistribute-other-ospf=no metric-default=1 metric-connected=20 metric-static=20 metric-rip=20 metric-bgp=auto
      metric-other-ospf=auto in-filter=ospf-in out-filter=ospf-out
`,
			[]map[string]string{
				{
					"number":                  "0",
					"comment":                 "",
					"name":                    "default",
					"disabled":                "no",
					"default":                 "yes",
					"router-id":               "10.242.0.2",
					"distribute-default":      "never",
					"redistribute-connected":  "as-type-1",
					"redistribute-static":     "no",
					"redistribute-rip":        "no",
					"redistribute-bgp":        "no",
					"redistribute-other-ospf": "no",
					"metric-default":          "1",
					"metric-connected":        "20",
					"metric-static":           "20",
					"metric-rip":              "20",
					"metric-bgp":              "auto",
					"metric-other-ospf":       "auto",
					"in-filter":               "ospf-in",
					"out-filter":              "ospf-out",
				},
			},
		},
		{`Flags: X - disabled, A - active, D - dynamic, C - connect, S - static, r - rip, b - bgp, o - ospf, m - mme, B - blackhole, U - unreachable, P - prohibit
		 0 A S  dst-address=0.0.0.0/0 gateway=192.168.1.254 gateway-status=192.168.1.254 reachable via  ether5 distance=1 scope=30 target-scope=10

		  1 A S  dst-address=161.65.59.1/32 gateway=192.168.1.254 gateway-status=192.168.1.254 reachable via  ether5 distance=1 scope=30 target-scope=10

		   2 A S  dst-address=161.65.63.1/32 gateway=192.168.1.254 gateway-status=192.168.1.254 reachable via  ether5 distance=1 scope=30 target-scope=10
`,
			[]map[string]string{
				{
					"number":         "0",
					"comment":        "",
					"active":         "yes",
					"dynamic":        "no",
					"connect":        "no",
					"static":         "yes",
					"rip":            "no",
					"bgp":            "no",
					"ospf":           "no",
					"mme":            "no",
					"blackhole":      "no",
					"unreachable":    "no",
					"prohibit":       "no",
					"disabled":       "no",
					"dst-address":    "0.0.0.0/0",
					"gateway":        "192.168.1.254",
					"gateway-status": "192.168.1.254 reachable via ether5",
					"distance":       "1",
					"scope":          "30",
					"target-scope":   "10",
				},
				{
					"number":         "1",
					"comment":        "",
					"active":         "yes",
					"dynamic":        "no",
					"connect":        "no",
					"static":         "yes",
					"rip":            "no",
					"bgp":            "no",
					"ospf":           "no",
					"mme":            "no",
					"blackhole":      "no",
					"unreachable":    "no",
					"prohibit":       "no",
					"disabled":       "no",
					"dst-address":    "161.65.59.1/32",
					"gateway":        "192.168.1.254",
					"gateway-status": "192.168.1.254 reachable via ether5",
					"distance":       "1",
					"scope":          "30",
					"target-scope":   "10",
				},
				{
					"number":         "2",
					"comment":        "",
					"active":         "yes",
					"dynamic":        "no",
					"connect":        "no",
					"static":         "yes",
					"rip":            "no",
					"bgp":            "no",
					"ospf":           "no",
					"mme":            "no",
					"blackhole":      "no",
					"unreachable":    "no",
					"prohibit":       "no",
					"disabled":       "no",
					"dst-address":    "161.65.63.1/32",
					"gateway":        "192.168.1.254",
					"gateway-status": "192.168.1.254 reachable via ether5",
					"distance":       "1",
					"scope":          "30",
					"target-scope":   "10",
				},
			},
		},
		{`Flags: I - invalid
		 0   name="ospf-restart" owner="admin" policy=ftp,reboot,read,write,policy,test,password,sniff,sensitive run-count=0 source=
		       :local State;
		       :local Interface;
		       :local Network;
		       :local Address;
		       :local helpAdd;
		       :local help0;
		       :local help1;
		       :local ipNetAdd;
		       :local networkName;
		       :local fileName value="ospf_restart.txt";

		       # If previous data file not found, then create new one
		       :if ([:len [/file find where name=$fileName]] < 1 ) do={
		       /file print file=$fileName where name=$fileName;
		       # Add some delay, for slow or high load routers
		       /delay delay-time=2;
		       # Set begginning count to 0
		       /file set $fileName contents="0";
		       }


		       :foreach i in=[/routing ospf neighbor find] do={
		           :set State [/routing ospf neighbor get $i value-name=state];
		           :put "\nNetwork state: $State";


		           :if ($State = "Init" || $State = "Down") do={
		               :put "Condition met";
		       # read some working variables
		               :set Interface [/routing ospf neighbor get $i value-name=interface];
		               :set Address [/routing ospf neighbor get $i value-name=address];
		               :put "OSPF neighbor interface:                              $Interface";
		               :put "OSPF neighbor address:                                $Address";
		               :set helpAdd ($Address&255.255.255.0);
		               :put "OSPF neighbor addres without last octet:              $helpAdd";
		       # find all networks in ip addresses matching OSPF neighbor address without last octet - Intermediate Step
		               :set ipNetAdd [/ip address find where (network&255.255.255.0)=$helpAdd];
		               :put "IP address item where is OSPF neightbor matching address without last octet: "
		               :put $ipNetAdd;
		       # find the only network in ip addresses matching OSPF neighbor address and interface
		               :set help1 [/ip address find where interface=$Interface && (network&255.255.255.0)=$helpAdd ];
		               :put "IP address item matching network AND interface:       $help1";
		       # find ospf network to be resetet
		               :set Network [/ip address get $help1 value-name=network];
		               :put "Which OSPF network should be reseted:                 $Network";
		       # find ospf network item to disable + enable
		               :set help0 [/routing ospf network find where network~"$Network/*"];
		               :set networkName [/routing ospf network get $help0 value-name=comment];
		               :put "OSPF network item number to be reseted:               $help0";
		       # Restart OSPF network
		               /routing ospf network set $help0 disabled=yes;
		               :put "Network has been DISABLED";
		               /routing ospf network print;
		               [/routing ospf network set $help0 disabled=no];
		               :put "Network has been ENABLED";
		       # Add record to log
		               :log info "OSPF network $Network - $networkName has been RESTARTED";
		       # Update restart count in file
		               :local before value=[/file get $fileName contents];
		               /file set $fileName contents= ($before + 1);

		               /routing ospf network print;
		           } else={
		               :put "Condition NOT met";
		             }
		       }
		`,
			[]map[string]string{
				{
					"number":    "0",
					"name":      "ospf-restart",
					"owner":     "admin",
					"policy":    "ftp,reboot,read,write,policy,test,password,sniff,sensitive",
					"run-count": "0",
					"source": `
		       :local State;
		       :local Interface;
		       :local Network;
		       :local Address;
		       :local helpAdd;
		       :local help0;
		       :local help1;
		       :local ipNetAdd;
		       :local networkName;
		       :local fileName value="ospf_restart.txt";

		       # If previous data file not found, then create new one
		       :if ([:len [/file find where name=$fileName]] < 1 ) do={
		       /file print file=$fileName where name=$fileName;
		       # Add some delay, for slow or high load routers
		       /delay delay-time=2;
		       # Set begginning count to 0
		       /file set $fileName contents="0";
		       }


		       :foreach i in=[/routing ospf neighbor find] do={
		           :set State [/routing ospf neighbor get $i value-name=state];
		           :put "\nNetwork state: $State";


		           :if ($State = "Init" || $State = "Down") do={
		               :put "Condition met";
		       # read some working variables
		               :set Interface [/routing ospf neighbor get $i value-name=interface];
		               :set Address [/routing ospf neighbor get $i value-name=address];
		               :put "OSPF neighbor interface:                              $Interface";
		               :put "OSPF neighbor address:                                $Address";
		               :set helpAdd ($Address&255.255.255.0);
		               :put "OSPF neighbor addres without last octet:              $helpAdd";
		       # find all networks in ip addresses matching OSPF neighbor address without last octet - Intermediate Step
		               :set ipNetAdd [/ip address find where (network&255.255.255.0)=$helpAdd];
		               :put "IP address item where is OSPF neightbor matching address without last octet: "
		               :put $ipNetAdd;
		       # find the only network in ip addresses matching OSPF neighbor address and interface
		               :set help1 [/ip address find where interface=$Interface && (network&255.255.255.0)=$helpAdd ];
		               :put "IP address item matching network AND interface:       $help1";
		       # find ospf network to be resetet
		               :set Network [/ip address get $help1 value-name=network];
		               :put "Which OSPF network should be reseted:                 $Network";
		       # find ospf network item to disable + enable
		               :set help0 [/routing ospf network find where network~"$Network/*"];
		               :set networkName [/routing ospf network get $help0 value-name=comment];
		               :put "OSPF network item number to be reseted:               $help0";
		       # Restart OSPF network
		               /routing ospf network set $help0 disabled=yes;
		               :put "Network has been DISABLED";
		               /routing ospf network print;
		               [/routing ospf network set $help0 disabled=no];
		               :put "Network has been ENABLED";
		       # Add record to log
		               :log info "OSPF network $Network - $networkName has been RESTARTED";
		       # Update restart count in file
		               :local before value=[/file get $fileName contents];
		               /file set $fileName contents= ($before + 1);

		               /routing ospf network print;
		           } else={
		               :put "Condition NOT met";
		             }
		       }`,
					"comment": "",
					"invalid": "no",
				},
			},
		},
		{
			`Flags: X - disabled, * - default
			 0  * interface=all
			`,
			[]map[string]string{
				{
					"number":    "0",
					"default":   "yes",
					"disabled":  "no",
					"interface": "all",
					"comment":   "",
				},
			},
		},
		{
			`Flags: X - disabled, * - default
			 0 X* interface=all
			`,
			[]map[string]string{
				{
					"number":    "0",
					"default":   "yes",
					"disabled":  "yes",
					"interface": "all",
					"comment":   "",
				},
			},
		},
		{
			` 0 interface=ether1 address=10.236.0.81 address4=10.236.0.81 mac-address=00:0C:42:EB:5E:54 identity="wanrt-it05avl" platform="MikroTik" version="6.40.5 (stable)" unpack=none age=28s uptime=10w2d20h42m32s
   software-id="PD8V-GGMC" board="RB1100AH" interface-name="ether8" system-description="MikroTik RouterOS 6.40.5 (stable) RB1100AH" system-caps=bridge,router system-caps-enabled=bridge,router

 1 interface=ether2 address=10.240.0.30 address4=10.240.0.30 mac-address=4C:5E:0C:60:F9:23 identity="wanrt1-cottonbuildingvuw" platform="MikroTik" version="6.35.1 (stable)" unpack=none age=40s
   uptime=62w2d20h27m27s software-id="LRZA-SSJI" board="RB2011UiAS" ipv6=no interface-name="ether2" system-caps="" system-caps-enabled=""

 2 interface=ether4 address=10.239.0.20 address4=10.239.0.20 mac-address=4C:5E:0C:41:3C:80 identity="rf2hillroad-soundstage" platform="MikroTik" version="6.36 (stable)" unpack=none age=5s
   uptime=41w2d22h53m54s software-id="3RXV-YCE2" board="RB Metal 5SHPn" ipv6=no interface-name="bridge1" system-caps="" system-caps-enabled=""

 3 interface=ether4 address=10.239.0.25 address4=10.239.0.25 mac-address=D4:CA:6D:A2:36:61 identity="rf2soundstage-hillroad" platform="MikroTik" version="6.36 (stable)" unpack=none age=23s
   uptime=1w3d17h33m1s software-id="L2WP-FCDH" board="RB912UAG-5HPnD" ipv6=no interface-name="bridge1" system-caps="" system-caps-enabled=""
`,
			[]map[string]string{
				{
					"address":             "10.236.0.81",
					"version":             "6.40.5 (stable)",
					"address4":            "10.236.0.81",
					"unpack":              "none",
					"uptime":              "10w2d20h42m32s",
					"comment":             "",
					"software-id":         "PD8V-GGMC",
					"system-description":  "MikroTik RouterOS 6.40.5 (stable) RB1100AH",
					"system-caps-enabled": "bridge,router",
					"interface-name":      "ether8",
					"number":              "0",
					"board":               "RB1100AH",
					"interface":           "ether1",
					"mac-address":         "00:0C:42:EB:5E:54",
					"platform":            "MikroTik",
					"system-caps":         "bridge,router",
					"age":                 "28s",
					"identity":            "wanrt-it05avl",
				},
				{
					"interface-name":      "ether2",
					"platform":            "MikroTik",
					"unpack":              "none",
					"uptime":              "62w2d20h27m27s",
					"ipv6":                "no",
					"identity":            "wanrt1-cottonbuildingvuw",
					"software-id":         "LRZA-SSJI",
					"comment":             "",
					"board":               "RB2011UiAS",
					"version":             "6.35.1 (stable)",
					"system-caps":         "",
					"age":                 "40s",
					"mac-address":         "4C:5E:0C:60:F9:23",
					"number":              "1",
					"address":             "10.240.0.30",
					"system-caps-enabled": "",
					"interface":           "ether2",
					"address4":            "10.240.0.30",
				},
				{
					"mac-address":         "4C:5E:0C:41:3C:80",
					"identity":            "rf2hillroad-soundstage",
					"comment":             "",
					"software-id":         "3RXV-YCE2",
					"interface-name":      "bridge1",
					"age":                 "5s",
					"address4":            "10.239.0.20",
					"uptime":              "41w2d22h53m54s",
					"system-caps-enabled": "",
					"interface":           "ether4",
					"number":              "2",
					"version":             "6.36 (stable)",
					"board":               "RB Metal 5SHPn",
					"platform":            "MikroTik",
					"unpack":              "none",
					"ipv6":                "no",
					"system-caps":         "",
					"address":             "10.239.0.20",
				},
				{
					"number":              "3",
					"mac-address":         "D4:CA:6D:A2:36:61",
					"version":             "6.36 (stable)",
					"identity":            "rf2soundstage-hillroad",
					"uptime":              "1w3d17h33m1s",
					"age":                 "23s",
					"ipv6":                "no",
					"unpack":              "none",
					"interface-name":      "bridge1",
					"address":             "10.239.0.25",
					"comment":             "",
					"system-caps-enabled": "",
					"platform":            "MikroTik",
					"board":               "RB912UAG-5HPnD",
					"software-id":         "L2WP-FCDH",
					"system-caps":         "",
					"interface":           "ether4",
					"address4":            "10.239.0.25",
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
				for k := range test.m[i] {
					if _, ok := m[i][k]; !ok {
						t.Errorf("extra: %s", k)
					}
				}
				for k := range m[i] {
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
