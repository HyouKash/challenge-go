package piscine

func AlphaCount(s string) int {
	ss := []byte(s)
	count := 0
	for k := 0; k < len(ss); k++ {
		if ss[k] >= 65 && ss[k] <= 90 || ss[k] >= 97 && ss[k] <= 122 {
			count += 1
		}
	}
	return count
}
