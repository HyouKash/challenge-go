package piscine

func Swap(a *int, b *int) {
	temp := *a
	temp2 := *b
	*a = temp2
	*b = temp
}
