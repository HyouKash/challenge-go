package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	newLine := ""
	voy := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	voyels := []rune{}
	if len(args) == 0 {
		z01.PrintRune('\n')
	} else {
		for _, k := range args {
			newLine += k
			newLine += " "
		}
		for _, i := range newLine {
			for _, x := range voy {
				if i == x {
					voyels = append(voyels, i)
				}
			}
		}
		for i := 0; i < (len(voyels) / 2); i++ {
			voyels[i], voyels[len(voyels)-1-i] = voyels[len(voyels)-1-i], voyels[i]
		}
		voyE := 0
		test := []byte(newLine)
		for number, x := range test {
			for _, s := range voy {
				if rune(x) == s {
					test[number] = byte(voyels[voyE])
					voyE++
				}
			}
		}
		for _, s := range test {
			z01.PrintRune(rune(s))
		}
		z01.PrintRune('\n')
	}
}
