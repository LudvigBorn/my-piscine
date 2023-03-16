package main

import (
	"fmt"
	piscine "piscine"
)

func main() {
	fmt.Println(piscine.SwapBits(125))
}

// test

// package main

// import (
// 	"fmt"
// 	piscine "piscine"
// )

// func main() {
// 	fmt.Println(piscine.ReverseBits(122))
// 	// fmt.Println(test.FindPrevPrime(430))
// }

//************* doop *************

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fir := atoi(os.Args[1])
// 	ope := os.Args[2]
// 	sec := atoi(os.Args[3])
// 	if !isNum(os.Args[1]) {
// 		return
// 	}
// 	if !isNum(os.Args[3]) {
// 		return
// 	}

// 	if isOver(fir, sec) {
// 		return
// 	}

// 	switch ope {
// 	case "+":
// 		fmt.Println(fir + sec)
// 		return
// 	case "-":
// 		fmt.Println(fir - sec)
// 		return
// 	case "/":
// 		if sec == 0 {
// 			fmt.Println("No division by 0")
// 			return
// 		}
// 		fmt.Println(fir / sec)
// 		return
// 	case "*":
// 		fmt.Println(fir * sec)
// 		return
// 	case "%":
// 		if sec == 0 {
// 			fmt.Println("No modulo by 0")
// 			return
// 		}
// 		fmt.Println(fir % sec)
// 		return
// 	default:
// 		return
// 	}
// }

// func isOver(a, b int) bool {
// 	if (a+b < 0) && (a > 0 && b > 0) {
// 		return true
// 	}

// 	if (a-b > 0) && (a < 0 && b < 0) {
// 		return true
// 	}

// 	if b != (a*b)/a {
// 		return true
// 	}

// 	return false
// }

// func atoi(s string) int {
// 	res := 0
// 	for _, v := range s {
// 		res = res*10 + int(v-'0')
// 	}
// 	return res
// }

// func isNum(s string) bool {
// 	for _, v := range s {
// 		if v > '9' || v < '0' {
// 			return false
// 		}
// 	}
// 	return true
// }

//************* searchreplace *************

// package main

// import (
// 	"fmt"
// 	"os"
// )package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	s := os.Args[1]
// 	f := os.Args[2]
// 	rep := os.Args[3]
// 	str := []rune(s)
// 	new := []rune{}
// 	for _, v := range str {
// 		if string(v) == f {
// 			// fmt.Println(string(v))
// 			v = []rune(rep)[0]
// 		}
// 		new = append(new, v)

// 	}
// 	fmt.Println(string(new))
// }
// 	s := os.Args[1]
// 	f := os.Args[2]
// 	rep := os.Args[3]
// 	str := []rune(s)
// 	new := []rune{}
// 	for _, v := range str {
// 		if string(v) == f {
// 			// fmt.Println(string(v))
// 			v = []rune(rep)[0]
// 		}
// 		new = append(new, v)

// 	}
// 	fmt.Println(string(new))
// }

//************* tabmult *************

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	n := atoi(os.Args[1])
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("%v x %v = %v", i, n, i*n)
// 		fmt.Println()
// 	}
// }

// func atoi(s string) int {
// 	res := 0
// 	for _, v := range s {
// 		res = res*10 + int(v-'0')
// 	}
// 	return res
// }

// ************* expandstr *************
// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	str := os.Args[1]
// 	res := ""
// 	flag := false
// 	block := len(str)
// 	if str[len(str)-1] == ' ' {
// 		for i := len(str) - 1; i >= 0; i-- {
// 			if str[i] != ' ' {
// 				block = i + 1
// 				break
// 			}
// 		}
// 	}

// 	for i := 0; i < block; i++ {
// 		if str[0] == ' ' && flag == false {

// 			for str[i] == ' ' {
// 				i++
// 			}
// 			flag = true
// 		}

// 		res += string(str[i])
// 	}

// 	fmt.Println(res)
// }

// ************* switchcase *************

// package main

// import (
// 	"fmt"expandstr
// 	"os"
// )

// func main() {
// 	str := os.Args[1]
// 	res := []rune{}
// 	for _, v := range str {
// 		if v >= 'a' && v <= 'z' {
// 			v -= 'a' - 'A'
// 		} else if v >= 'A' && v <= 'Z' {
// 			v += 'a' - 'A'
// 		}
// 		res = append(res, v)
// 	}
// 	fmt.Println(string(res))
// }

// ************* lastword *************

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {expandstr
// 	s := os.Args[1]
// 	str := []string{}
// 	word := []rune{}

// 	for i, v := range s {
// 		if v != ' ' {
// 			word = append(word, v)
// 		} else {
// 			if len(word) != 0 {
// 				str = append(str, string(word))
// 			}
// 			word = make([]rune, 0)
// 		}
// 		if i == len(s)-1 && len(word) != 0 {
// 			str = append(str, string(word))
// 		}
// 	}
// 	fmt.Println(str[len(str)-1])
// }

// ************* rot13 *************

// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	if len(os.Args[1:]) != 1 {
// 		return
// 	}expandstr
// 		if v >= 'a' && v <= 'z' {
// 			if v < 'n' {
// 				v += 13
// 			} else {
// 				v = 'a' + ((v - 'a') - 13)
// 			}
// 		}

// 		z01.PrintRune(v)
// 	}
// }

//************* ispowerof2 *************

// package main

// import (
// 	"fmpackage main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	if len(os.Args[1:]) != 1 {
// 		return
// 	}
// 	n := atoi(os.Args[1])
// 	for n != 1 {
// 		if n%2 != 0 {
// 			fmt.Println(false)
// 			return
// 		}
// 		n /= 2
// 	}
// 	fmt.Println(true)
// }

// func atoi(s string) int {
// 	res := 0
// 	for _, v := range s {
// 		res = res*10 + int(v-'0')
// 	}
// 	return res
// }t"
// 	"os"
// )

// func main() {
// 	if len(os.Args[1:]) != 1 {
// 		return
// 	}
// 	n := atoi(os.Args[1])
// 	for n != 1 {
// 		if n%2 != 0 {
// 			fmt.Println(false)
// 			return
// 		}
// 		n /= 2
// 	}
// 	fmt.Println(true)
// }

// func atoi(s string) int {
// 	res := 0
// 	for _, v := range s {
// 		res = res*10 + int(v-'0')
// 	}
// 	return res
// }

//************* wdmatch *************

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	args := os.Args[1:]

// 	if len(args) != 2 {
// 		return
// 	}
// 	f := args[0]
// 	str := args[1]

// 	for _, v := range f {
// 		found := false
// 		for j, v2 := range str {
// 			if v == v2 {
// 				str = str[j:]
// 				found = true
// 				// fmt.Println(str)
// 				break
// 			}
// 		}
// 		if !found {
// 			return
// 		}
// 	}
// 	fmt.Println(f)
// }
