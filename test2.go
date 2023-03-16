package test

import (
	"fmt"

	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	end := lastNum(n)
	start := firstNum(n)
	fmt.Println(start)
	fmt.Println(end)
	for i := start; i <= end; i++ {
		if isCorrect(i) {
			printNmbr(i, n)
			if i != end {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

func lastNum(n int) int {
	num := 0
	dev := 1
	for i := 9; n >= 1; n-- {
		num += i * dev
		dev *= 10
		i--
	}
	return num
}

func firstNum(n int) int {
	num := 0
	dev := 1
	for i := n - 1; n >= 1; n-- {
		num += i * dev
		dev *= 10
		i--
	}
	return num
}

func printNmbr(num, n int) {
	arr := make([]rune, 0)
	dev := num
	for dev >= 1 {
		arr = append(arr, rune(dev%10+'0'))
		dev /= 10
	}
	if num < powinDec(n-1) {
		arr = append(arr, '0')
	}
	for i := len(arr) - 1; i >= 0; i-- {
		z01.PrintRune(arr[i])
	}
}

func isCorrect(num int) bool {
	arr := make([]int, 0)
	for num >= 1 {
		arr = append(arr, num%10)
		num /= 10
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] <= arr[j] {
				return false
			}
		}
	}
	return true
}

func powinDec(n int) int {
	num := 1
	for n != 0 {
		num *= 10
		n--
	}
	return num
}
