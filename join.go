package piscine

func Join(strs []string, sep string) string {
	rep := ""
	for k := 0; k < len(strs); k++ {
		if k == len(strs)-1 {
			rep += strs[k]
		} else {
			rep += strs[k] + sep
		}
	}
	return rep
}
