package test

import "fmt"

func SwapBits(octet byte) byte {
	massiv := make([]byte, 8)
	var res byte
	for i := 7; i >= 0; i-- {
		massiv[i] = octet % 2
		octet = octet / 2

	}
	fmt.Println(massiv)
	for i := 0; i < 4; i++ {
		massiv[i], massiv[i+4] = massiv[i+4], massiv[i]
	}
	fmt.Println(massiv)
	// for i := len(massiv) - 1; i >= 0; i-- {
	// 	res += Pow(2, len(massiv)-1-i) * massiv[i]
	// 	fmt.Println(res)
	// }
	for i := 7; i >= 0; i-- {
		res += Pow(2, len(massiv)-i-1) * massiv[i]
		fmt.Println(Pow(2, len(massiv)-i-1) * massiv[i])
		// fmt.Println(res)
	}

	return res
}

func Pow(a byte, b int) byte {
	aSave := a
	if b == 0 {
		return 1
	}
	for i := 1; i < b; i++ {
		a *= aSave
	}
	return a
}
