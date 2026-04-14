package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func tcpPortChecker(serv string, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		//build address
		servNport := fmt.Sprintf("%s:%d",serv,j)
		
		connex, err := net.DialTimeout("tcp", servNport, 500 * time.Millisecond)
		if err != nil {
			continue
		}
		connex.Close()
		//output if port is open.
		fmt.Printf("\033[36m%v\033[32m, is Open\n\033[97m", j)
	}	
}	
