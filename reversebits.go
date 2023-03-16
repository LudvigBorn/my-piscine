package test

import "fmt"

func ReverseBits(oct byte) byte {
	//	fmt.Println(oct)
	arr := make([]byte, 8)
	//	var arr byte = 66

	for i := 7; i >= 0; i-- {
		arr[i] = (oct % 2)
		oct /= 2
	}
	// fmt.Println(arr)
	for i := 0; i < len(arr)/2; i++ {
		arr[len(arr)-i-1], arr[i] = arr[i], arr[len(arr)-i-1]
	}
	fmt.Println(arr)
	var res byte
	for i := len(arr) - 1; i >= 0; i-- {

		res += pow(2, len(arr)-1-i) * arr[i]
		fmt.Println(res)
	}
	// fmt.Println(pow(2, 4))
	return res
}

func pow(a byte, b int) byte {
	aSave := a
	for i := 1; i < b; i++ {
		a *= aSave
	}
	return a
}
