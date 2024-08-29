package main

import (
	"flag"
	"sync"
)

func main() {

	urlFlag := flag.String("url", "127.0.0.1", "Give an URL or IP to be scanned")
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
