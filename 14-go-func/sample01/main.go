package main

import "fmt"

func main() {
	doSomething()

	doSomethingWith("haha")

	a := doSomethingAndReturn()

	fmt.Println(a)

}

func doSomething() {
	fmt.Println("do some thing")
}

func doSomethingWith(value string) {
	fmt.Println("do some thing" + value)
}

func doSomethingAndReturn() string {
	return fmt.Sprint("do some thing and return")
}