package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"

		msg := <- c
		fmt.Println(msg)
	}
}

func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)

		c <- "response"
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go printer(c)
	go pinger(c)

	var input string
	fmt.Scanln(&input)
}