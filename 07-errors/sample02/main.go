package main

import "fmt"

type MyCustomError struct {
	name string
}


func (e MyCustomError) Error() string {
	return fmt.Sprintf("this is myu custom error, name: %s", e.name)
}


func main() {
	var err error = MyCustomError{name: "test error"}

	fmt.Println(err.Error())
}
