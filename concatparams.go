package piscine

func ConcatParams(args []string) string {
	rep := ""
	for _, k := range args[:len(args)-1] {
		rep += k + "\n"
	}
	return rep + args[len(args)-1]
}
