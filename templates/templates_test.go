package templates_test

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/ffddorf/confgen/netbox"
	"github.com/ffddorf/confgen/templates"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEdgeOS(t *testing.T) {
	device := &netbox.Device{}
	require.NoError(t, json.Unmarshal([]byte(mockAPIData), &device))

	data := templates.TemplateData{Device: device}
	out := new(bytes.Buffer)
	require.NoError(t, templates.Render(out, "edgeos", data))

	assert.Equal(t,
		strings.TrimSpace(out.String()),
		strings.TrimSpace(expectedConfig),
	)
}

const expectedConfig = `
interfaces {
  ethernet eth0 {
    disable
  }
  ethernet eth1 {
    disable
  }
  ethernet eth2 {
    description "LAN management"
    address 10.1.0.3/16
    address 2001:678:b7c:201:7a8a:20ff:fe46:cb0/64
    address fe80::7a8a:20ff:fe46:cb0/64
    vif 10 {
      address 100.64.16.1/29
      address fdcb:aa6b:5532:25::1/64
    }
    vif 11 {
      address 10.129.0.76/31
      address fdcb:aa6b:5532:26::1/64
    }
  }
  ethernet eth3 {
    disable
  }
  ethernet eth4 {
    disable
  }
  ethernet eth5 {
    disable
  }
  ethernet eth6 {
    disable
  }
  ethernet eth7 {
    description "CR4 interconnect"
    address 45.151.166.10/31
  }
  ethernet eth8 {
    description "FFRL peering"
    address 185.66.192.193/31
    address 2a03:2260:0:25::2/64
  }
  loopback lo {
    description Loopback
    address 10.0.0.3/32
    address 45.151.166.3/32
    address 2001:678:b7c::3/128
  }
}
`

const mockAPIData = `
{
	"name": "CR3",
	"rack": {
		"name": "RK1"
	},
	"location": {
		"name": "Turm CoLo"
	},
	"site": {
		"name": "DUS2"
	},
	"primary_ip4": {
		"address": "45.151.166.3/32"
	},
	"primary_ip6": {
		"address": "2001:678:b7c::3/128"
	},
	"interfaces": [
		{
			"name": "eth0",
			"description": "",
			"type": "A_1000BASE_T",
			"enabled": false,
			"speed": null,
			"duplex": null,
			"ip_addresses": [],
			"parent": null,
			"child_interfaces": []
		},
		{
			"name": "eth1",
			"description": "",
			"type": "A_1000BASE_X_SFP",
			"enabled": false,
			"speed": null,
			"duplex": null,
			"ip_addresses": [],
			"parent": null,
			"child_interfaces": []
		},
		{
			"name": "eth2",
			"description": "LAN management",
			"type": "A_10GBASE_X_SFPP",
			"enabled": true,
			"speed": null,
			"duplex": null,
			"ip_addresses": [
				{
					"address": "10.1.0.3/16"
				},
				{
					"address": "2001:678:b7c:201:7a8a:20ff:fe46:cb0/64"
				},
				{
					"address": "fe80::7a8a:20ff:fe46:cb0/64"
				}
			],
			"parent": null,
			"child_interfaces": [
				{
					"name": "eth2.10",
					"description": "",
					"type": "VIRTUAL",
					"enabled": true,
					"untagged_vlan": {
						"vid": 10
					},
					"ip_addresses": [
						{
							"address": "100.64.16.1/29"
						},
						{
							"address": "fdcb:aa6b:5532:25::1/64"
						}
					]
				},
				{
					"name": "eth2.11",
					"description": "",
					"type": "VIRTUAL",
					"enabled": true,
					"untagged_vlan": {
						"vid": 11
					},
					"ip_addresses": [
						{
							"address": "10.129.0.76/31"
						},
						{
							"address": "fdcb:aa6b:5532:26::1/64"
						}
					]
				}
			]
		},
		{
			"name": "eth2.10",
			"description": "",
			"type": "VIRTUAL",
			"enabled": true,
			"speed": null,
			"duplex": null,
			"ip_addresses": [
				{
					"address": "100.64.16.1/29"
				},
				{
					"address": "fdcb:aa6b:5532:25::1/64"
				}
			],
			"parent": {
				"id": "91"
			},
			"child_interfaces": []
		},
		{
			"name": "eth2.11",
			"description": "",
			"type": "VIRTUAL",
			"enabled": true,
			"speed": null,
			"duplex": null,
			"ip_addresses": [
				{
					"address": "10.129.0.76/31"
				},
				{
					"address": "fdcb:aa6b:5532:26::1/64"
				}
			],
			"parent": {
				"id": "91"
			},
			"child_interfaces": []
		},
		{
			"name": "eth3",
			"description": "",
			"type": "A_10GBASE_X_SFPP",
			"enabled": false,
			"speed": null,
			"duplex": null,
			"ip_addresses": [],
			"parent": null,
			"child_interfaces": []
		},
		{
			"name": "eth4",
			"description": "",
			"type": "A_10GBASE_X_SFPP",
			"enabled": false,
			"speed": null,
			"duplex": null,
			"ip_addresses": [],
			"parent": null,
			"child_interfaces": []
		},
		{
			"name": "eth5",
			"description": "",
			"type": "A_10GBASE_X_SFPP",
			"enabled": false,
			"speed": null,
			"duplex": null,
			"ip_addresses": [],
			"parent": null,
			"child_interfaces": []
		},
		{
			"name": "eth6",
			"description": "",
			"type": "A_10GBASE_X_SFPP",
			"enabled": false,
			"speed": null,
			"duplex": null,
			"ip_addresses": [],
			"parent": null,
			"child_interfaces": []
		},
		{
			"name": "eth7",
			"description": "CR4 interconnect",
			"type": "A_10GBASE_X_SFPP",
			"enabled": true,
			"speed": null,
			"duplex": null,
			"ip_addresses": [
				{
					"address": "45.151.166.10/31"
				}
			],
			"parent": null,
			"child_interfaces": []
		},
		{
			"name": "eth8",
			"description": "FFRL peering",
			"type": "A_10GBASE_X_SFPP",
			"enabled": true,
			"speed": null,
			"duplex": null,
			"ip_addresses": [
				{
					"address": "185.66.192.193/31"
				},
				{
					"address": "2a03:2260:0:25::2/64"
				}
			],
			"parent": null,
			"child_interfaces": []
		},
		{
			"name": "lo",
			"description": "Loopback",
			"type": "VIRTUAL",
			"enabled": true,
			"speed": null,
			"duplex": null,
			"ip_addresses": [
				{
					"address": "10.0.0.3/32"
				},
				{
					"address": "45.151.166.3/32"
				},
				{
					"address": "2001:678:b7c::3/128"
				}
			],
			"parent": null,
			"child_interfaces": []
		},
		{
			"name": "tun0",
			"description": "R12 dago",
			"type": "OTHER",
			"enabled": true,
			"speed": null,
			"duplex": null,
			"ip_addresses": [
				{
					"address": "10.129.0.0/31"
				},
				{
					"address": "fdcb:aa6b:5532:1::1/64"
				}
			],
			"parent": null,
			"child_interfaces": []
		},
		{
			"name": "tun1",
			"description": "R9 voelklinger",
			"type": "OTHER",
			"enabled": true,
			"speed": null,
			"duplex": null,
			"ip_addresses": [
				{
					"address": "10.129.0.2/31"
				},
				{
					"address": "fdcb:aa6b:5532:2::1/64"
				}
			],
			"parent": null,
			"child_interfaces": []
		},
		{
			"name": "tun2",
			"description": "R10 niessdonk",
			"type": "OTHER",
			"enabled": true,
			"speed": null,
			"duplex": null,
			"ip_addresses": [
				{
					"address": "10.129.0.4/31"
				},
				{
					"address": "fdcb:aa6b:5532:3::1/64"
				}
			],
			"parent": null,
			"child_interfaces": []
		},
		{
			"name": "tun3",
			"description": "R15 Gatherweg",
			"type": "OTHER",
			"enabled": true,
			"speed": null,
			"duplex": null,
			"ip_addresses": [
				{
					"address": "10.129.0.10/31"
				},
				{
					"address": "fdcb:aa6b:5532:6::1/64"
				}
			],
			"parent": null,
			"child_interfaces": []
		},
		{
			"name": "tun4",
			"description": "R14 hoeherweg",
			"type": "OTHER",
			"enabled": true,
			"speed": null,
			"duplex": null,
			"ip_addresses": [
				{
					"address": "10.129.0.14/31"
				},
				{
					"address": "fdcb:aa6b:5532:8::1/64"
				}
			],
			"parent": null,
			"child_interfaces": []
		}
	]
}
`
