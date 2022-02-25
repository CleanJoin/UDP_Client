package main

import (
	"fmt"
	"net"
	"runtime"
	"strconv"
)

var CH2 = make(chan string)

func Writer(ch chan<- string) {
	for i := 0; i < 5000; i++ {
		ch <- strconv.Itoa(i)
	}

}
func Read(ch <-chan string, conn *net.UDPConn) {
	for msg := range ch {
		WriteUdp(msg, conn)
	}

}

func WriteUdp(msg string, conn *net.UDPConn) {
	conn.Write([]byte(msg))
}

func main() {
	numcpu := runtime.NumCPU()
	fmt.Println("NumCPU", numcpu)
	runtime.GOMAXPROCS(2)
	ch := make(chan string, 1)
	RemoteAddr, err := net.ResolveUDPAddr("udp", ":6000")
	if err != nil {
		fmt.Println(err)
	}
	conn, err := net.DialUDP("udp", nil, RemoteAddr)
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	go Writer(ch)
	go Read(ch, conn)

	var input string
	fmt.Scan(&input)
}
