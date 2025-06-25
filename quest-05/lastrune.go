package piscine

func LastRune(s string) rune {
	letters := []rune(s)
	return letters[len(s)-1]
}
