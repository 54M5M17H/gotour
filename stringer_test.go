package gotour

import (
	"testing"
)

type ipTestCase struct {
	expected string
	ip       IPAddr
}

func TestStringer(t *testing.T) {
	hosts := map[string]ipTestCase{
		"loopback": {
			"127.0.0.1",
			IPAddr{127, 0, 0, 1},
		},
		"googleDNS": {
			"8.8.8.8",
			IPAddr{8, 8, 8, 8},
		},
	}

	for _, testCase := range hosts {
		ipAsString := testCase.ip.String()
		if ipAsString != testCase.expected {
			t.Errorf("Expected: %s. Instead got: %s", testCase.expected, ipAsString)
		}
	}
}
