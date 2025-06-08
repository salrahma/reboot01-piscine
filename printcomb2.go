package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for tens1 := '0'; tens1 <= '9'; tens1++ {
		for ones1 := '0'; ones1 <= '9'; ones1++ {
			for tens2 := '0'; tens2 <= '9'; tens2++ {
				for ones2 := '0'; ones2 <= '9'; ones2++ {
					z01.PrintRune(tens1)
					z01.PrintRune(ones1)
					z01.PrintRune(' ')

					z01.PrintRune(tens2)
					z01.PrintRune(ones2)

					if tens1 == '9' && ones1 == '8' && tens2 == '9' && ones2 == '9' {
					} else {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
