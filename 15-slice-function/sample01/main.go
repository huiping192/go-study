package sample01

import "fmt"

func functionSlice() {

	var fns []func()

	fns = append(fns, beeper)
	fns = append(fns, pinger)

	for _, fn := range fns {
		fn()
	}
}


func beeper() {
	fmt.Println("beep-beep")
}

func pinger() {
	fmt.Println("ping-ping")
}

func main() {

	functionSlice()
}



