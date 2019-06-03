package main

import "fmt"

func main() {

	a := test()


	value := a()

	fmt.Println(value)



	test2 := func() string {
		return "hello world 2"
	}

	value2 := test2()

	fmt.Println(value2)

}


func test() func() string {
	return func () string {
		return "hello world"
	}
}
