package _if

import "fmt"

func calc() (int, int) {
	return 1, 2
}

func PrintIf() {

	test := true

	if test {
		fmt.Println("test is true")
	} else {
		fmt.Println("test is false")
	}

	// 違うtestになる
	if test, test2 := calc(); test+test2 < 10 {
		fmt.Println("test:", test)
		fmt.Println("test2:", test2)
		fmt.Println("test + test2 < 10")
	}

}
