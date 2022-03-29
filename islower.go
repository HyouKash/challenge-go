package piscine

func IsLower(s string) bool {
	for k := 0; k < len([]rune(s)); k++ {
		if []rune(s)[k] < 97 || []rune(s)[k] > 122 {
			return false
		}
	}
	return true
}
