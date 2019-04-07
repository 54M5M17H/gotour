package gotour

import (
	"strconv"
	"strings"
)

// IPAddr -- represenation of an IP address
type IPAddr [4]byte

func (addr IPAddr) String() string {
	var stringified [4]string

	for i, seg := range addr {
		num := int(seg)
		newStr := strconv.Itoa(num)
		stringified[i] = newStr
	}
	return strings.Join(stringified[:], ".")
}

