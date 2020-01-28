package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {

	//addresses, _ := net.LookupIP("localhost") // take from 1st argument
	//for i := 0; i < len(addresses); i++ {
	//
	//	segments := strings.SplitAfter(addresses[i].String(), " ") //<--- here!
	//
	//	fmt.Printf("IP address #%d : %s \n", i, segments)
	//}
	//
	//a := addresses[1].String()

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:7777")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))

}


func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
