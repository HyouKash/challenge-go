package piscine

func SortIntegerTable(table []int) {
	for k := 0; k < len(table)-1; k++ {
		iMin := k
		for i := k + 1; i < len(table); i++ {
			if table[i] < table[iMin] {
				iMin = i
			}
		}
		table[k], table[iMin] = table[iMin], table[k]
	}
}
