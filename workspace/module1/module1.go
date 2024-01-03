package module1

import "strconv"

// Int returns the decimal reversal of the integer i.
func Int(i int) int {
	i, _ = strconv.Atoi(String(strconv.Itoa(i)))
	return i
}

func String(itoa string) string {
	return String(itoa)
}
