package piscine

func Atoi(s string) int {
	a := []byte(s)
	r := 0
	isNega := false
	grandeur := 1
	if len(s) == 0 {
		return 0
	}
	if a[0] == 45 {
		isNega = true
		tempory := a
		tempory[0] = '0'
		s = string(tempory)
	}
	if a[0] == 43 {
		tempory := a
		tempory[0] = '0'
		s = string(tempory)
	}
	for k := len(a) - 1; k >= 0; k-- {
		if a[k] > 57 || a[k] < 48 {
			return 0
		} else {
			r += int(a[k]-48) * grandeur
			grandeur *= 10
		}
	}
	if isNega {
		return r * -1
	}
	return r
}
