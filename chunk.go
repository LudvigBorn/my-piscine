package test

import (
	"github.com/01-edu/z01"
)

func Chunk(slice []int, size int) {
	if len(slice) == 0 {
		z01.PrintRune('[')
		z01.PrintRune(']')
		z01.PrintRune('\n')
		return
	}
	if size == 0 {
		z01.PrintRune('\n')
		return
	}
	z01.PrintRune('[')
	z01.PrintRune('[')
	for i, v := range slice {
		if (i+1)%size == 0 && i != len(slice)-1 {
			z01.PrintRune(rune(v) + '0')
			// fmt.Println("work")
			z01.PrintRune(']')
			z01.PrintRune(' ')
			z01.PrintRune('[')
		} else if i != len(slice)-1 {
			z01.PrintRune(rune(v) + '0')
			z01.PrintRune(' ')
		} else if i == len(slice)-1 {
			z01.PrintRune(rune(v) + '0')
		}
	}
	z01.PrintRune(']')
	z01.PrintRune(']')
	z01.PrintRune('\n')
}
