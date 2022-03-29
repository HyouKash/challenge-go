package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	a := 0
	b := 0
	c := 0
	d := 0
	e := 0
	f := 0
	g := 0
	h := 0
	z := 0
	if n == 1 {
		for k := 0; k <= 9; k++ {
			a = k
			z01.PrintRune(rune(a) + 48)
			if a == 9 {
			} else {
				z01.PrintRune(rune(44))
				z01.PrintRune(rune(32))
			}
		}
		z01.PrintRune('\n')
	}
	if n == 2 {
		for k := 1; k <= 999; k++ {
			if b > a {
				z01.PrintRune(rune(a + 48))
				z01.PrintRune(rune(b + 48))

				if a == 8 && b == 9 {
					z01.PrintRune(rune('\n'))
				} else {
					z01.PrintRune(rune(44))
					z01.PrintRune(rune(32))
				}
			}
			b += 1
			if b == 10 {
				b = 0
				a += 1
			}
		}
	}
	if n == 3 {
		for i := 0; i < 999; i++ {
			if c > b && b > a {
				z01.PrintRune(rune(a + 48))
				z01.PrintRune(rune(b + 48))
				z01.PrintRune(rune(c + 48))

				if a == 7 && b == 8 && c == 9 {
					z01.PrintRune(rune('\n'))
				} else {
					z01.PrintRune(rune(44))
					z01.PrintRune(rune(32))
				}
			}
			c += 1
			if c == 10 {
				c = 0
				b += 1
			}
			if b == 10 {
				b = 0
				a += 1
			}
		}
	}
	if n == 4 {
		for i := 0; i < 9999; i++ {
			if d > c && c > b && b > a {
				z01.PrintRune(rune(a + 48))
				z01.PrintRune(rune(b + 48))
				z01.PrintRune(rune(c + 48))
				z01.PrintRune(rune(d + 48))
				if a == 6 && b == 7 && c == 8 && d == 9 {
					z01.PrintRune(rune('\n'))
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
	}
	if n == 5 {
		for i := 0; i < 99999; i++ {
			if e > d && d > c && c > b && b > a {
				z01.PrintRune(rune(a + 48))
				z01.PrintRune(rune(b + 48))
				z01.PrintRune(rune(c + 48))
				z01.PrintRune(rune(d + 48))
				z01.PrintRune(rune(e + 48))
				if a == 5 && b == 6 && c == 7 && d == 8 && e == 9 {
					z01.PrintRune(rune('\n'))
				} else {
					z01.PrintRune(rune(44))
					z01.PrintRune(rune(32))
				}
			}
			e += 1
			if e == 10 {
				e = 0
				d += 1
			}
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
	}
	if n == 6 {
		for i := 0; i < 999999; i++ {
			if f > e && e > d && d > c && c > b && b > a {
				z01.PrintRune(rune(a + 48))
				z01.PrintRune(rune(b + 48))
				z01.PrintRune(rune(c + 48))
				z01.PrintRune(rune(d + 48))
				z01.PrintRune(rune(e + 48))
				z01.PrintRune(rune(f + 48))
				if a == 4 && b == 5 && c == 6 && d == 7 && e == 8 && f == 9 {
					z01.PrintRune(rune('\n'))
				} else {
					z01.PrintRune(rune(44))
					z01.PrintRune(rune(32))
				}
			}
			f += 1
			if f == 10 {
				f = 0
				e += 1
			}
			if e == 10 {
				e = 0
				d += 1
			}
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
	}
	if n == 7 {
		for i := 0; i < 9999999; i++ {
			if g > f && f > e && e > d && d > c && c > b && b > a {
				z01.PrintRune(rune(a + 48))
				z01.PrintRune(rune(b + 48))
				z01.PrintRune(rune(c + 48))
				z01.PrintRune(rune(d + 48))
				z01.PrintRune(rune(e + 48))
				z01.PrintRune(rune(f + 48))
				z01.PrintRune(rune(g + 48))
				if a == 3 && b == 4 && c == 5 && d == 6 && e == 7 && f == 8 && g == 9 {
					z01.PrintRune(rune('\n'))
				} else {
					z01.PrintRune(rune(44))
					z01.PrintRune(rune(32))
				}
			}
			g += 1
			if g == 10 {
				g = 0
				f += 1
			}
			if f == 10 {
				f = 0
				e += 1
			}
			if e == 10 {
				e = 0
				d += 1
			}
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
	}
	if n == 8 {
		for i := 0; i < 99999999; i++ {
			if h > g && g > f && f > e && e > d && d > c && c > b && b > a {
				z01.PrintRune(rune(a + 48))
				z01.PrintRune(rune(b + 48))
				z01.PrintRune(rune(c + 48))
				z01.PrintRune(rune(d + 48))
				z01.PrintRune(rune(e + 48))
				z01.PrintRune(rune(f + 48))
				z01.PrintRune(rune(g + 48))
				z01.PrintRune(rune(h + 48))
				if a == 2 && b == 3 && c == 4 && d == 5 && e == 6 && f == 7 && g == 8 && h == 9 {
					z01.PrintRune(rune('\n'))
				} else {
					z01.PrintRune(rune(44))
					z01.PrintRune(rune(32))
				}
			}
			h += 1
			if h == 10 {
				h = 0
				g += 1
			}
			if g == 10 {
				g = 0
				f += 1
			}
			if f == 10 {
				f = 0
				e += 1
			}
			if e == 10 {
				e = 0
				d += 1
			}
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
	}
	if n == 9 {
		for i := 0; i < 999999999; i++ {
			if z > h && h > g && g > f && f > e && e > d && d > c && c > b && b > a {
				z01.PrintRune(rune(a + 48))
				z01.PrintRune(rune(b + 48))
				z01.PrintRune(rune(c + 48))
				z01.PrintRune(rune(d + 48))
				z01.PrintRune(rune(e + 48))
				z01.PrintRune(rune(f + 48))
				z01.PrintRune(rune(g + 48))
				z01.PrintRune(rune(h + 48))
				z01.PrintRune(rune(z + 48))
				if a == 1 && b == 2 && c == 3 && d == 4 && e == 5 && f == 6 && g == 7 && h == 8 && z == 9 {
					z01.PrintRune(rune('\n'))
				} else {
					z01.PrintRune(rune(44))
					z01.PrintRune(rune(32))
				}
			}
			z += 1
			if z == 10 {
				z = 0
				h += 1
			}
			if h == 10 {
				h = 0
				g += 1
			}
			if g == 10 {
				g = 0
				f += 1
			}
			if f == 10 {
				f = 0
				e += 1
			}
			if e == 10 {
				e = 0
				d += 1
			}
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
	}
}
