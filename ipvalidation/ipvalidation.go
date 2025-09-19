package ipvalidation

import (
	"strconv"
	"strings"
)

func Is_valid_ip(ip string) bool {
	octets := strings.Split(ip, ".")

	if len(octets) != 4 {
		return false
	}

	for _, octet := range octets {
		number, err := strconv.Atoi(octet)
		if err != nil {
			return false
		}

		if number < 0 || number > 255 {
			return false
		}

		if len(octet) > 1 && octet[0] == '0' {
			return false
		}
	}

	return true
}
