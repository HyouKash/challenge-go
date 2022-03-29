package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	a := 0
	b := 1
	c := 2
	for k := 0; k <= 999; k++ {
		if c > b && b > a {
			z01.PrintRune(rune(a + 48))
			z01.PrintRune(rune(b + 48))
			z01.PrintRune(rune(c + 48))
			if a == 7 && b == 8 && c == 9 {
			} else {
				z01.PrintRune(32)
				z01.PrintRune(44)
			}
		}
		c++
		if c == 10 {
			c = 0
			b++
		}
		if b == 10 {
			b = 0
			a++
		}
	}
	z01.PrintRune('\n')
}
