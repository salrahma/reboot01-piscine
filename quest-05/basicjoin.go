package piscine

func BasicJoin(elems []string) string {
	s := ""
	for i := 0; i <= len(elems)-1; i++ {
		s = s + elems[i]
	}
	return s
}
