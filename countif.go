package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, k := range tab {
		if f(k) == true {
			count += 1
		}
	}
	return count
}
