package piscine

func RecursiveFactorial(nb int) int {
	num := nb
	if nb < 0 || nb >= 26 {
		return 0
	}
	if nb == 0 {
		return 1
	}

	return num * RecursiveFactorial(num-1)
}
