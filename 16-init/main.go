package main

import (
	"fmt"
)

var a = func() int{
	return 100
}()

func main() {

	fmt.Println(a)
}

func init() {
	a = 1000
}

