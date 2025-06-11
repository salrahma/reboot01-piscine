package piscine

import "fmt"

func BasicAtoi(s string) int {
	temp := ""
	// arr1 := []rune(s)
	// arr2 := []rune{}
	for i, v := range []rune(s) {
		if v >= '0' && v <= '9' {
			if v == '0' && temp == "" {
			} else {
				b := int(v)
				fmt.Println((i + 1) * 10)
				temp = string(b)
				fmt.Println(temp)
				// arr2 := append(arr2, v)
			}

			// temp = string(arr2)
			// 01.PrintRune(i)
		}
	}

	return 0
}
