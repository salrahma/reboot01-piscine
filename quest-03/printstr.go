package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	letter := []rune(s)
	for i := 0; i < len(s); i++ {
		z01.PrintRune(letter[i])
	}
}
