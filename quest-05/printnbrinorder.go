package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	arr := []rune{}
	var num int
	for n > 0 {
		// segregate by place order
		num = n % 10
		arr = append(arr, rune(num+'0'))
		n = n / 10
	}
	for a1 := 0; a1 < len(arr); a1++ {
		for a2 := a1 + 1; a2 < len(arr); a2++ {
			if arr[a1] > arr[a2] {
				arr[a1], arr[a2] = arr[a2], arr[a1]
			}
		}
	}
	for _, order := range arr {
		z01.PrintRune(order)
	}
}
