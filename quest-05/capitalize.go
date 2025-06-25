package piscine

func Capitalize(s string) string {
	s2 := []rune(s)
	for i, letter := range s2 {
		if isAlphaNumeric(letter) {
			if i == 0 || !isAlphaNumeric(s2[i-1]) {
				if s2[i] >= 'a' && letter <= 'z' {
					s2[i] = letter - 32
				}
			} else {
				if s2[i] >= 'A' && letter <= 'Z' {
					s2[i] = letter + 32
				}
			}
		}
	}
	return string(s2)
}

func isAlphaNumeric(ch rune) bool {
	if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
		return true
	}
	return false
}
