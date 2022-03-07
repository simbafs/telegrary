package util

import "strconv"

// AddZero add zero to the front of the number
// e.g 2 -> 02
//     12 -> 12
func AddZero(i int) string {
	if i < 10 {
		return "0" + strconv.Itoa(i)
	}
	return strconv.Itoa(i)
}

