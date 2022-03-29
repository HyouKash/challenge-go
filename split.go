package piscine

func Split(s, sep string) []string {
	rep := []string{}
	start := 0
	for k := 0; k < len(s)-len(sep); k++ {
		if s[k:k+len(sep)] == sep {
			rep = append(rep, s[start:k])
			start = k + len(sep)
		}
	}
	rep = append(rep, s[start:])
	return rep
}
