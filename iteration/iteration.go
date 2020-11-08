package iteration


func Repeat5Times(inp string) string {
	out := ""
	for i := 0; i < 5; i++ {
		out += inp
	}
	return out
}

func RepeatNTimes(inp string, x int) string {
	out := ""
	for i := 0; i < x; i++ {
		out += inp
	}
	return out
}
