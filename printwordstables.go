package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, k := range a {
		for _, x := range k {
			z01.PrintRune(x)
		}
		z01.PrintRune('\n')
	}
}
