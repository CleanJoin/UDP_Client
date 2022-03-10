package main

import (
	"fmt"
	"net"
	"runtime"
	"time"
)

func main() {
	numcpu := runtime.NumCPU()
	fmt.Println("NumCPU", numcpu)
	// runtime.GOMAXPROCS(1)
	planets := make([]byte, 1)
	ch := make(chan []byte, 1024)

	m := &runtime.MemStats{}
	runtime.ReadMemStats(m)

	RemoteAddr, err := net.ResolveUDPAddr("udp", ":6000")
	if err != nil {
		fmt.Println(err)
	}
	conn, err := net.DialUDP("udp", nil, RemoteAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	var j int
	for j = 0; j < 50; j++ {

		for i := 0; i < 5000; i++ {
			go func() {
				planets[0]++
				ch <- planets
				conn.Write(<-ch)
			}()
			time.Sleep(time.Microsecond)
		}
	}
}
