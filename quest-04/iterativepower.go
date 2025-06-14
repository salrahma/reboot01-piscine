package piscine

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	num := 1
	if nb == 0 {
		return 0
	}

	if power < 0 {
		return 0
	}

	for n := 1; n <= power; n++ {
		num = num * nb
	}

	return num
}
