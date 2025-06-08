package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	// num := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for hundreds := '0'; hundreds <= '7'; hundreds++ {
		for tens := hundreds + 1; tens <= '8'; tens++ {
			for ones := tens + 1; ones <= '9'; ones++ {

				z01.PrintRune(hundreds)
				z01.PrintRune(tens)
				z01.PrintRune(ones)
				if hundreds == '7' && tens == '8' && ones == '9' {
				} else {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
