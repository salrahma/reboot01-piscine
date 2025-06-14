package piscine

func IterativeFactorial(nb int) int {
	num := 1
	if nb < 0 || nb > 26 {
		return 0
	}
	if nb == 0 {
		return 1
	}

	for n := 1; n <= nb; n++ {
		num = num * n
	}
	if num < 0 {
		return 0
	}
	return num
}
