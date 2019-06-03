package main

import "fmt"

func main() {
  a,b := doSomeThing()

  fmt.Println("value: ",a,b)


	c,d := doSomeThingYes()

	fmt.Println("value: ",c,d)

}

func doSomeThing() (int,string) {
	return 10,"hehe"
}



func doSomeThingYes() (a int,b string) {
	return 10,"hehe"
}
