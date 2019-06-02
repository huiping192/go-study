package main

import (
	"errors"
	"fmt"
)

func doSomething() (bool,error) {

	return false, errors.New("custom error")
}

func doSomething2() (bool,error) {
	return false, fmt.Errorf("my error is you %s", "error")
}

func main() {

	result, err := doSomething()

	if !result && err != nil {
		fmt.Println(err.Error())
	}

	result2, err2 := doSomething2()

	if !result2 && err2 != nil {
		fmt.Println(err2.Error())
	}

}


