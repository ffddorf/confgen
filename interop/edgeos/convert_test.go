package edgeos_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ffddorf/confgen/interop/edgeos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type SM = map[string]interface{}

func TestMapConversion(t *testing.T) {
	in := SM{
		"smart-queue internal": SM{
			"download": SM{
				"rate": "39mbit",
			},
			"wan-interface": "eth1.2",
		},
	}

	out := &strings.Builder{}
	err := edgeos.ConfigFromMap(out, in, 0)
	require.NoError(t, err)

	fmt.Println(out.String())
	assert.Equal(t, `smart-queue internal {
  download {
    rate 39mbit
  }
  wan-interface eth1.2
}
`, out.String())
}
