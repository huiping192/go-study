package constant

import "fmt"

const (
	// iotaつけることで順番になる
	monday = iota + 1
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

const (
	a = iota
	b
)


func PrintIotaNumber() {
	fmt.Println(monday)
	fmt.Println(tuesday)
	fmt.Println(sunday)

	fmt.Println(a)
	fmt.Println(b)
}