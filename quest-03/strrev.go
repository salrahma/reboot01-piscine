package piscine

// "github.com/01-edu/z01"

func StrRev(s string) string {
	arr1 := []rune(s)
	arr2 := []rune{}
	l := len(arr1)

	for n := 0; n <= l-1; n++ {
		arr2 = append(arr2, arr1[l-n-1])
	}
	return string(arr2)
}
