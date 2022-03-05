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
		time.Sleep(time.Millisecond + 1)
		go conn.Write(msg)
	}
}

func start() int {
	planets := make([]byte, 1024)

	numcpu := runtime.NumCPU()
	fmt.Println("NumCPU", numcpu)
	// runtime.GOMAXPROCS(2)
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
	j := 0
	for j = 0; j < 6; j++ {

		go func() {
			time.Sleep(time.Millisecond + 24)
			for i := 0; i < 5000; i++ {
				Writer(ch, planets)
			}
		}()

		go Read(ch, conn)
	}
	return j
}
func main() {
	start()
}
