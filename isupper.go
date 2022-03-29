package piscine

func IsUpper(s string) bool {
	for k := 0; k < len([]rune(s)); k++ {
		if []rune(s)[k] <= 65 || []rune(s)[k] >= 90 {
			return false
		}
	}
	return true
}
