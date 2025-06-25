package piscine

func Join(strs []string, sep string) string {
	s := ""
	for i := 0; i <= len(strs)-1; i++ {
		s = s + strs[i]
		if i <= len(strs)-2 {
			s = s + sep
		}
	}
	return s
}
