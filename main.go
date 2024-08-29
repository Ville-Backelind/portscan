/*
		Features to implement:
[]

*/

package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

func main() {

	urlFlag := flag.String("url", "", "Give an URL or IP to be scanned")
	portStartFlag := flag.Int("start", 1, "Starting port number")
	portEndFlag := flag.Int("end", 65535, "Ending port number")
	flag.Parse()
	printWelcome()
	var wg sync.WaitGroup

	for p := *portStartFlag; p <= *portEndFlag; p++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tcpPortChecker(*urlFlag, p)
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
	fmt.Printf("\033[36m%v"+"\033[32m, is Open\n\033[97m", p)
}
