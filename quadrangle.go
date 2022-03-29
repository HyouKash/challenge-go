package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	for s := 1; s <= y; s++ {
		for k := 1; k <= x; k++ {
			if k == 1 && s == 1 {
				if k == x {
					z01.PrintRune(111)
					z01.PrintRune('\n')
				} else {
					z01.PrintRune(111)
				}
			} else if k == x && s == 1 {
				z01.PrintRune(111)
				z01.PrintRune('\n')
			} else if k == 1 && s == y {
				z01.PrintRune(111)
			} else if k == x && s == y {
				z01.PrintRune(111)
				z01.PrintRune('\n')
			} else if s == 1 || s == y {
				z01.PrintRune('-')
			} else if k == 1 || k == x {
				if s != 1 || s != y {
					z01.PrintRune('|')
					if k == x {
						z01.PrintRune('\n')
					}
				}
			} else {
				z01.PrintRune(32)
			}
		}
	}
}

func QuadE(x, y int) {
	for s := 1; s <= y; s++ {
		for k := 1; k <= x; k++ {
			if k == 1 && s == 1 {
				if k == x {
					z01.PrintRune('A')
					z01.PrintRune('\n')
				} else {
					z01.PrintRune('A')
				}
			} else if k == x && s == 1 {
				z01.PrintRune('C')
				z01.PrintRune('\n')
			} else if k == 1 && s == y {
				z01.PrintRune('C')
			} else if k == x && s == y {
				z01.PrintRune('A')
				z01.PrintRune('\n')
			} else if s == 1 || s == y {
				z01.PrintRune('B')
			} else if k == 1 || k == x {
				if s != 1 || s != y {
					z01.PrintRune('B')
					if k == x {
						z01.PrintRune('\n')
					}
				}
			} else {
				z01.PrintRune(32)
			}
		}
	}
}

func QuadD(x, y int) {
	for s := 1; s <= y; s++ {
		for k := 1; k <= x; k++ {
			if k == 1 && s == 1 {
				if k == x {
					z01.PrintRune('A')
					z01.PrintRune('\n')
				} else {
					z01.PrintRune('A')
				}
			} else if k == x && s == 1 {
				z01.PrintRune('C')
				z01.PrintRune('\n')
			} else if k == 1 && s == y {
				z01.PrintRune('A')
			} else if k == x && s == y {
				z01.PrintRune('C')
				z01.PrintRune('\n')
			} else if s == 1 || s == y {
				z01.PrintRune('B')
			} else if k == 1 || k == x {
				if s != 1 || s != y {
					z01.PrintRune('B')
					if k == x {
						z01.PrintRune('\n')
					}
				}
			} else {
				z01.PrintRune(32)
			}
		}
	}
}

func QuadC(x, y int) {
	for s := 1; s <= y; s++ {
		for k := 1; k <= x; k++ {
			if k == 1 && s == 1 {
				if k == x {
					z01.PrintRune('A')
					z01.PrintRune('\n')
				} else {
					z01.PrintRune('A')
				}
			} else if k == x && s == 1 {
				z01.PrintRune('A')
				z01.PrintRune('\n')
			} else if k == 1 && s == y {
				z01.PrintRune('C')
			} else if k == x && s == y {
				z01.PrintRune('C')
				z01.PrintRune('\n')
			} else if s == 1 || s == y {
				z01.PrintRune('B')
			} else if k == 1 || k == x {
				if s != 1 || s != y {
					z01.PrintRune('B')
					if k == x {
						z01.PrintRune('\n')
					}
				}
			} else {
				z01.PrintRune(32)
			}
		}
	}
}

func QuadB(x, y int) {
	for s := 1; s <= y; s++ {
		for k := 1; k <= x; k++ {
			if k == 1 && s == 1 {
				if k == x {
					z01.PrintRune('/')
					z01.PrintRune('\n')
				} else {
					z01.PrintRune('/')
				}
			} else if k == x && s == 1 {
				z01.PrintRune('\\')
				z01.PrintRune('\n')
			} else if k == 1 && s == y {
				z01.PrintRune('\\')
			} else if k == x && s == y {
				z01.PrintRune('/')
				z01.PrintRune('\n')
			} else if s == 1 || s == y {
				z01.PrintRune('*')
			} else if k == 1 || k == x {
				if s != 1 || s != y {
					z01.PrintRune('*')
					if k == x {
						z01.PrintRune('\n')
					}
				}
			} else {
				z01.PrintRune(32)
			}
		}
	}
}
