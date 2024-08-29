package main

import (
	"fmt"
	"net"
)

func tcpPortChecker(serv string, p int) {
	servNport := fmt.Sprintf(serv+":%v", p)
	connex, err := net.Dial("tcp", servNport)
	if err != nil {
		return
	}
	defer connex.Close()
	fmt.Printf("\033[36m%v"+"\033[32m, is Open\n\033[97m", p)
}
