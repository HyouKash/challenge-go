package piscine

func ToUpper(s string) string {
	rep := ""
	for k := 0; k < len([]rune(s)); k++ {
		if []rune(s)[k] >= 97 && []rune(s)[k] <= 122 {
			rep += string([]rune(s)[k] - 32)
		} else {
			rep += string([]rune(s)[k])
		}
	}
	return rep
}
