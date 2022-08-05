package edgeos_test

import (
	"strings"
	"testing"

	"github.com/ffddorf/confgen/interop/edgeos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	edgeos.ForceConsistentMapOrdering = true
}

type SM = map[string]interface{}

func TestMapConversion(t *testing.T) {
	testCases := map[string]struct {
		in  map[string]interface{}
		out string
	}{
		"smart queue": {
			in: SM{
				"smart-queue internal": SM{
					"download": SM{
						"rate": "39mbit",
					},
					"wan-interface": "eth1.2",
				},
			},
			out: `smart-queue internal {
  download {
    rate 39mbit
  }
  wan-interface eth1.2
}
`,
		},
		"quoted values": {
			in: SM{
				"ethernet eth0": SM{
					"description": "Some interface doing something",
				},
			},
			out: `ethernet eth0 {
  description "Some interface doing something"
}
`,
		},
		"multivalue": {
			in: SM{
				"interfaces": SM{
					"ethernet eth1": SM{
						"address": []interface{}{"10.1.0.1/16", "fde4:4d90:9ebf::1/64"},
					},
				},
			},
			out: `interfaces {
  ethernet eth1 {
    address 10.1.0.1/16
    address fde4:4d90:9ebf::1/64
  }
}
`,
		},
		"numbers": {
			in: SM{
				"protocols": SM{
					"bgp 207871": SM{
						"maximum-paths": SM{
							"ibgp": 2,
						},
					},
				},
			},
			out: `protocols {
  bgp 207871 {
    maximum-paths {
      ibgp 2
    }
  }
}
`,
		},
		"multivalue numbers": {
			in: SM{
				"snmp": SM{
					"community public": SM{
						"authorization": "ro",
					},
					"contact": "support@freifunk-duesseldorf.de",
					"listen-address 2001:678:b7c::3": SM{
						"port": []interface{}{161, "345"},
					},
				},
			},
			out: `snmp {
  community public {
    authorization ro
  }
  contact support@freifunk-duesseldorf.de
  listen-address 2001:678:b7c::3 {
    port 161
    port 345
  }
}
`,
		},
		"booleans": {
			in: SM{
				"ethernet eth3": SM{
					"disable": true,
					"duplex":  "auto",
					"speed":   "auto",
				},
			},
			out: `ethernet eth3 {
  disable
  duplex auto
  speed auto
}
`,
		},
		"boolean off": {
			in: SM{
				"ethernet eth3": SM{
					"disable": false,
					"duplex":  "auto",
					"speed":   "auto",
				},
			},
			out: `ethernet eth3 {
  duplex auto
  speed auto
}
`,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			out := &strings.Builder{}
			err := edgeos.ConfigFromMap(out, tc.in, 0)
			require.NoError(t, err)
			assert.Equal(t, tc.out, out.String())
		})
	}
}
