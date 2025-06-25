package piscine

func Index(s string, toFind string) int {
	for i := 0; i <= len(s)-1; i++ {
		if (len(s) - len(toFind)) <= 1 {
			return -1
		}
		if string(s[i:i+len(toFind)]) == toFind {
			return i
		}
	}
	return -1
}
