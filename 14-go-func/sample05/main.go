package main

import "fmt"

type user struct {
	name string
	password string
}


func login(user user) bool {
	if user.name == "a" && user.password == "b" {
		return true
	}

	return false
}

func main() {
	user := user{name:"a",password:"b"}
	ok := login(user)


	fmt.Println(ok)
}