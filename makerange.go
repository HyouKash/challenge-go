package piscine

func MakeRange(min, max int) []int {
	if min == 0 && max == 0 || min > max {
		return []int(nil)
	}
	size := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		size[i] = i + min
	}
	return size
}
