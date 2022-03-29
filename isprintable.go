package piscine

func IsPrintable(s string) bool {
	for k := 0; k < len([]rune(s)); k++ {
		if []rune(s)[k] >= 0 && []rune(s)[k] <= 31 {
			return false
		}
	}
	return true
}
