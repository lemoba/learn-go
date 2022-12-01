package strings

func HasPrefix(s string, ch string) bool {
	return len(s) >= len(ch) && s[0:len(ch)] == ch
}
