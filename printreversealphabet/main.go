package main

import "github.com/01-edu/z01"

func main() {
	alphabet := "zyxwvutsrqponmlkjihgfedcba\n"
	for _, k := range alphabet {
		z01.PrintRune(k)
	}
}
