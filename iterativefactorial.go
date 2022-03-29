package piscine

func IterativeFactorial(nb int) int {
	result := 1

	if nb == 0 {
		return 1
	} else if nb < 0 || nb > 1000 {
		return 0
	} else {
		for k := 1; k < nb+1; k++ {
			result *= k
		}
		return result
	}
}
