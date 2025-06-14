package piscine

func Sqrt(nb int) int {
	start := 0

	if nb < 1 {
		return 0
	}
	if nb == 1 {
		return 1
	}
	middle := int(float64(start+nb)/2) - ((start + nb) % 2)
	mid := middle
	for i := 0; i <= middle; i++ {
		if (mid)*(mid) > nb {
			mid = mid - 1
		}
	}
	if mid*mid == nb {
		return mid
	} else {
		return 0
	}
}
