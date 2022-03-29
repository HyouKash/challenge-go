package piscine

func IterativePower(nb int, power int) int {
	result := nb
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		for k := 0; k < power-1; k++ {
			result *= nb
		}
		return result
	}
}
