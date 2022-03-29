package piscine

func ToLower(s string) string {
	rep := ""
	for k := 0; k < len([]rune(s)); k++ {
		if []rune(s)[k] >= 65 && []rune(s)[k] <= 90 {
			rep += string([]rune(s)[k] + 32)
		} else {
			rep += string([]rune(s)[k])
		}
	}
	return rep
}
