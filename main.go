/*
		Features to implement:
[] read url and port from system arguments
[] scan with go routines
[]

*/

package main

import (
	"fmt"
	"net"
)

func main() {
	p := 80
	serv := "scanme.nmap.org"
	tcpPortChecker(serv, p)

}

func tcpPortChecker(serv string, p int) {
	servNport := fmt.Sprintf(serv+":%v", p)
	connex, err := net.Dial("tcp", servNport)
	if err != nil {
		return
	}
	defer connex.Close()
	fmt.Printf(`%v, is Open`, p)
}
