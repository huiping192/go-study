package main

import "fmt"

type Any interface {

}

type Car struct {
	name string
}

func main() {
	var value01 Any = 1
	var value02 Any = "hello"
	var value03 Any = Car{ name:"my car"}


	if value,ok := value01.(int); ok {
		fmt.Printf("this is int, value: %d \n", value)
	}

	if value,ok := value02.(string); ok {
		fmt.Printf("this is string, value: %s \n", value)
	}

	if value,ok := value03.(Car); ok {
		fmt.Println("this is car, value:" + value.name)
	}

	switch value := value03.(type) {
	case int:
		fmt.Printf("this is int, value: %d \n", value)
	case string:
		fmt.Printf("this is string, value: %s \n", value)
	case Car:
		fmt.Println("this is car, value:" + value.name)
	}
}
