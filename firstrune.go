package piscine

func FirstRune(s string) rune {
	ss := []rune(s)
	for k := 0; k < len(ss); k++ {
		return rune(ss[k])
	}
	return '\n'
}
