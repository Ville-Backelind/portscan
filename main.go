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
	o := fmt.Sprintf("%v, is Open", p)
	fmt.Println(o)
}
