package main

import (
	"./test"
	"./util"
	"fmt"
)

func main() {
	// test package
	test.SayHello()

	helloString := test.GetHello()

	fmt.Println(helloString)

	// util package
	util.Print("hello world")
}
