package sample01

import "fmt"

type car interface {
	desc() string
	test()
}


type mycar struct {
	name string
}

func (car mycar) test() {
	fmt.Println("test")
}

func (car mycar) desc() string {
	return "this is " + car.name
}


func main() {
	var car01 car = mycar{name:"super car"}

	fmt.Println(car01.desc())
}
