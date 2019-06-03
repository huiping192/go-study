package main

import (
	"fmt"
	"strconv"
)

func main() {

	func01 := func(a int) string {
		return "hello world" + strconv.Itoa(a)
	}

	fmt.Println(func01(100))



	value := func(a,b int) string {
		return "hello world"+ strconv.Itoa(a)+ strconv.Itoa(b)
	}(1,2)

	fmt.Println(value)

}
