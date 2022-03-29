package piscine

func BasicAtoi(s string) int {
	a := []byte(s)
	grandeur := 1
	rep := 0
	for k := len(a) - 1; k >= 0; k-- {
		rep += int((a[k])-48) * grandeur
		grandeur *= 10
	}
	return rep
}
