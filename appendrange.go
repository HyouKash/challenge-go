package piscine

func AppendRange(min, max int) []int {
	size := []int{}
	if min == 0 && max == 0 || min > max {
		return []int(nil)
	}
	for i := min; i < max; i++ {
		size = append(size, i)
	}
	return size
}
