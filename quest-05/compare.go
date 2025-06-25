package piscine

func Compare(a, b string) int {
	// fmt.Println(strings.Compare("a", "b"))
	str1 := []rune(a)
	str2 := []rune(b)
	l := len(a)
	Result := 0
	if l > len(b) {
		l = len(b)
	}
	for n := 0; n < l; n++ {
		if str1[n] == str2[n] {
			Result = 0
		} else if str1[n] < str2[n] {
			return -1
		} else {
			return 1
		}
	}
	if len(a) > len(b) {
		return 1
	}
	return Result
}
