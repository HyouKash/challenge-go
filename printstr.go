package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	a := []byte(s)
	for k := 0; k < len(s); k++ {
		z01.PrintRune(rune(a[k]))
	}
}
