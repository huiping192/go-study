package main

import "fmt"

type car struct {
	name string
	price string
	color string
}

func (c car) showName() {
	fmt.Println(c.name)
}

func (c *car) SetName01(s string) {
	fmt.Printf("SetName01: car address: %p\n", c)
	c.name = s
}


func (c car) SetName02(s string) {
	fmt.Printf("SetName01: car address: %p\n", c)
	c.name = s
}

func main() {
	yoyocar := &car{name:"yo",price:"100",color:"black"}

	// yo
	yoyocar.showName()

	yoyocar.SetName01("new name")

	// new name
	yoyocar.showName()

	yoyocar.SetName02("name2")

	// new name
	yoyocar.showName()
}