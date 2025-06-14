package piscine

func IsPrime(nb int) bool {
	count := 0
	for i := 1; i <= nb; i++ {
		if nb%i == 0 {
			count++
		}
	}
	return count == 2
}
