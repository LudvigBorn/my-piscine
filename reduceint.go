package test

import "fmt"

func ReduceInt(a []int, f func(int, int) int) {
	fmt.Println(f(a[0], a[1]))
}
