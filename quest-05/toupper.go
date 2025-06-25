package piscine

func ToUpper(s string) string {
	s2 := []rune{}
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			s2 = append(s2, letter-32)
		} else {
			s2 = append(s2, letter)
		}
	}
	return string(s2)
}
