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

	printWelcome("Portscanner Extreme")
	

	var wg sync.WaitGroup
	const workers = 100
	wg.Add(workers)
	jobs := make(chan int, 2 * workers)
	
	for i := 1; i <= workers; i++ { 
		go tcpPortChecker(*urlFlag, jobs, &wg	)
	}

	for p := *portStartFlag; p <= *portEndFlag; p++ {
		jobs <- p
	}

	close(jobs)
	wg.Wait()
}
