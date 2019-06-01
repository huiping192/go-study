package main

import "fmt"

func main() {
	go test()

	var input string

	fmt.Scanln(&input)
}

func test()  {
	for i :=0 ; i < 10 ; i++  {
		fmt.Println(i)
	}
}
