package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	a := 0
	b := 0
	c := 0
	d := 0
	for k := 0; k < 9999; k++ {
		if ((a * 10) + b) < ((c * 10) + d) {
			z01.PrintRune(rune(a + 48))
			z01.PrintRune(rune(b + 48))
			z01.PrintRune(rune(32))
			z01.PrintRune(rune(c + 48))
			z01.PrintRune(rune(d + 48))
			if a == 9 && b == 8 && c == 9 && d == 9 {
			} else {
				z01.PrintRune(rune(44))
				z01.PrintRune(rune(32))
			}
		}
		d += 1

		if d == 10 {
			d = 0
			c += 1
		}
		if c == 10 {
			c = 0
			b += 1
		}
		if b == 10 {
			b = 0
			a += 1
		}
	}
	z01.PrintRune(rune('\n'))
}
