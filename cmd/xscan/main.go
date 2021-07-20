// Author: Stanislav Yakush <st.yakush@yandex.ru>

package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

var address string
var startPort uint64
var endPort uint64
var timeout time.Duration
var workers uint64
var verbose bool

func parseFlags() {
	flag.StringVar(&address, "address", "127.0.0.1", "Address")
	flag.Uint64Var(&startPort, "start", 1, "Start port range")
	flag.Uint64Var(&endPort, "end", 65535, "End port range")
	flag.DurationVar(&timeout, "timeout", time.Second, "Timeout")
	flag.Uint64Var(&workers, "workers", 1, "Workers")
	flag.BoolVar(&verbose, "verbose", false, "Verbose")

	flag.Parse()
}

func showFlags() {
	fmt.Println("=========================================================")
	fmt.Println("================== TCP PORT SCANNER =====================")
	fmt.Println("=========================================================")

	fmt.Println("> Address:", address)
	fmt.Printf("> Port range: %d-%d\n", startPort, endPort)
	fmt.Println("> Timeout:", timeout)
	fmt.Println("> Workers:", workers)
	fmt.Println("> Verbose:", verbose)

	fmt.Println("=========================================================")
}

func main() {
	parseFlags()
	showFlags()

	var wg sync.WaitGroup

	addresses := make(chan string)

	for i := uint64(0); i <= workers; i++ {
		wg.Add(1)

		go func(addresses <-chan string) {
			defer wg.Done()

			for address := range addresses {
				_, err := net.DialTimeout("tcp", address, timeout)

				if err == nil {
					fmt.Println("Open:", address)
					continue
				}

				if verbose {
					fmt.Println("Closed:", address)
				}
			}
		}(addresses)
	}

	for i := startPort; i <= endPort; i++ {
		addresses <- fmt.Sprintf("%s:%d", address, i)
	}

	close(addresses)

	wg.Wait()
}
