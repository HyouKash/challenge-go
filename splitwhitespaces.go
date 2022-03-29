package piscine

func SplitWhiteSpaces(s string) []string {
	rep := []string{}
	start := 0
	for number, k := range s {
		if k == ' ' {
			rep = append(rep, s[start:number])
			start = number + 1
		}
		if number == len(s)-1 {
			rep = append(rep, s[start:number+1])
			start = number + 1
		}
	}
	finalrep := []string{}
	for _, k := range rep {
		if k != "" {
			finalrep = append(finalrep, k)
		}
	}
	return finalrep
}
