package piscine

func BasicAtoi2(s string) int {
	a := []byte(s)
	r := 0
	grandeur := 1
	for k := len(a) - 1; k >= 0; k-- {
		if a[k] > 57 || a[k] < 48 {
			return 0
		} else {
			r += int(a[k]-48) * grandeur
			grandeur *= 10
		}
	}
	return r
}
