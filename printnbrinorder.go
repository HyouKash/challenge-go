package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	bignumber := false
	lenNumber := 0
	finalNumber := 0
	table := []int{}
	if n == 0 {
	}
	if n < 0 {
		if n == -9223372036854775808 {
			bignumber = true
			n = -n - 1
		} else {
			n = n * -1
		}
	}
	if n > 1000000000000000000 {
		lenNumber = 19
		finalNumber = 1000000000000000000
	} else {
		for nb := 1; nb < n; nb *= 10 {
			lenNumber += 1
			finalNumber = nb
		}
	}
	for nb := finalNumber; nb > 0; nb /= 10 {
		if bignumber && n < 9 {
			n += 1
			bignumber = false
		}
		table = append(table, n/nb)
		n -= (n / nb) * nb
	}

	for k := 0; k < len(table)-1; k++ {
		iMin := k
		for i := k + 1; i < len(table); i++ {
			if table[i] < table[iMin] {
				iMin = i
			}
		}
		table[k], table[iMin] = table[iMin], table[k]
	}
	for k := 0; k < len(table); k++ {
		z01.PrintRune(rune(table[k]) + 48)
	}
}
