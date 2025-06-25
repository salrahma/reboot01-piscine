package piscine

func AlphaCount(s string) int {
	result := 0
	for _, v := range []rune(s) {
		if (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') {
			result = result + 1
		}
	}
	return result
}
