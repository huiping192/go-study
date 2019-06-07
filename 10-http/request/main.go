package main

import (
	"fmt"
	"net/http"
)

func main() {

	doGetRequest01()

	doGetRequest02()
}



func doGetRequest01() {
	res,err := http.Get("https://golang.org")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Status)
}


func doGetRequest02() {
	req,err := http.NewRequest("GET","https://golang.org",nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	client := http.DefaultClient

	res,err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res.Status)
}