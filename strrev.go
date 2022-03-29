package piscine

func StrRev(s string) string {
	a := []byte(s)
	r := ""
	for k := len(a) - 1; k >= 0; k-- {
		r += string(a[k])
	}
	return r
}
