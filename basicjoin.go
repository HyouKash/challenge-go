package piscine

func BasicJoin(elems []string) string {
	rep := ""
	for k := 0; k < len(elems); k++ {
		rep += elems[k]
	}
	return rep
}
