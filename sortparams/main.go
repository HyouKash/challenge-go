package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	for k := 0; k < len(args)-1; k++ {
		i_min := k
		for i := k + 1; i < len(args); i++ {
			if args[i] < args[i_min] {
				i_min = i
			}
		}
		args[k], args[i_min] = args[i_min], args[k]
	}
	for _, ascii := range args {
		for _, ez := range ascii {
			z01.PrintRune(ez)
		}
		z01.PrintRune('\n')
	}
}
