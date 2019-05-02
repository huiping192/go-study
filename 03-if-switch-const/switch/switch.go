package _switch

import "fmt"

func PrintSwitch() {
	a := 1
	switch {
	case a >= 0:
		fmt.Println("a is true")
		// default is no fallthrough
		//fallthrough
	case a > 200:
		fmt.Println("a is false")
	}

	b := 1

	switch b {
	case 1:
		fmt.Println("b is 1")
	default:
		fmt.Println("b is not 1")
	}
}
