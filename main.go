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
	"sync"
)

func main() {
	portList := []int{20,21,22,23,24,25,26,27,28,29,30}
	serv := "scanme.nmap.org"
	var wg sync.WaitGroup
	
	for _, P := range portList {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tcpPortChecker(serv, P)
		}()
		
	}
	wg.Wait()
}

func tcpPortChecker(serv string, p int) {
	servNport := fmt.Sprintf(serv+":%v", p)
	connex, err := net.Dial("tcp", servNport)
	if err != nil {
		return
	}
	defer connex.Close()
	fmt.Printf("%v, is Open\n", p)
}






