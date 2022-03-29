package piscine

func Any(f func(string) bool, a []string) bool {
	for _, k := range a {
		if f(k) == true {
			return true
		}
	}
	return false
}
