package main

import (
	"fmt"
	"net"
	"runtime"
	"time"
)

func Writer(ch chan<- []byte, planets []byte) {

	planets[0]++
	ch <- planets

}

func Read(ch <-chan []byte, conn *net.UDPConn) {
	for msg := range ch {
		WriteUdp(msg, conn)
	}

}

func WriteUdp(msg []byte, conn *net.UDPConn) {
	// time.Sleep(time.Millisecond + 24)
	conn.Write(msg)
}

func main() {
	planets := make([]byte, 1024)

	numcpu := runtime.NumCPU()
	fmt.Println("NumCPU", numcpu)
	// runtime.GOMAXPROCS(3)
	ch := make(chan []byte, 1024)

	RemoteAddr, err := net.ResolveUDPAddr("udp", ":6000")
	if err != nil {
		fmt.Println(err)
	}
	conn, err := net.DialUDP("udp", nil, RemoteAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	for {
		time.Sleep(time.Millisecond + 24)
		go func() {

			for i := 0; i < 5000; i++ {
				Writer(ch, planets)
			}
		}()

		go Read(ch, conn)

		// var input string
		// fmt.Scan(&input)

	}
}
