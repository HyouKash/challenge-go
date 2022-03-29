package piscine

func Map(f func(int) bool, a []int) []bool {
	rep := []bool{}
	for _, k := range a {
		rep = append(rep, f(k))
	}
	return rep
}
