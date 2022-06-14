package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	var hostPort string
	flag.StringVar(&hostPort, "c", "localhost:55355", "UDP endpoint to dial")
	flag.Parse()

	raddr, err := net.ResolveUDPAddr("udp", hostPort)
	if err != nil {
		panic(err)
	}

	var udp *net.UDPConn
	udp, err = net.DialUDP("udp", nil, raddr)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, 65530), 65530)
	for scanner.Scan() {
		line := scanner.Bytes()
		line = append(line, '\n')

		fmt.Print(string(line))
		_, err = udp.Write(line)
		if err != nil {
			panic(err)
		}
	}
}
