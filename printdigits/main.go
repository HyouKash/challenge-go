package main

import "github.com/01-edu/z01"

func main() {
	alphabet := "0123456789\n"
	for _, k := range alphabet {
		z01.PrintRune(k)
	}
}
