package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)
	bypass := []string{"x = ", ", y = "}
	for k := 0; k < len(bypass); k++ {
		for s := 0; s < len(bypass[k]); s++ {
			z01.PrintRune(rune(bypass[k][s]))
		}
		if k == 0 {
			n := points.x
			finalNumber := 0
			lenNumber := 0
			for nb := 1; nb < n; nb *= 10 {
				lenNumber += 1
				finalNumber = nb
			}
			for nb := finalNumber; nb > 0; nb /= 10 {
				z01.PrintRune(rune(n/nb) + 48)
				n -= (n / nb) * nb
			}
		} else if k == 1 {
			n := points.y
			finalNumber := 0
			lenNumber := 0
			for nb := 1; nb < n; nb *= 10 {
				lenNumber += 1
				finalNumber = nb
			}
			for nb := finalNumber; nb > 0; nb /= 10 {
				z01.PrintRune(rune(n/nb) + 48)
				n -= (n / nb) * nb
			}
		}

	}
	z01.PrintRune('\n')
}
