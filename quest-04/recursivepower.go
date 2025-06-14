package piscine

func RecursivePower(nb int, power int) int {
	num := nb
	if power == 0 {
		return 1
	}

	if nb == 0 {
		return 0
	}

	if power < 0 {
		return 0
	}

	return num * RecursivePower(num, power-1)
}
