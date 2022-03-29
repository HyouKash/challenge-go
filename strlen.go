package piscine

func StrLen(s string) int {
	a := []byte(s)
	count := 0
	for k := 0; k < len(s); k++ {
		if a[k] > 128 {
			k++
		}
		count++
	}
	return count
}
