package main

import (
	"flag"
	"fmt"
	"myhttptool/client"
	"myhttptool/worker"
	"net/http"
)

func main() {
	numOfParallelChannels := flag.Int("parallel", 10, "Parallel number of requests permitted (default 10)")
	flag.Parse()
	urlList := flag.Args()

	if len(urlList) == 0 {
		fmt.Println("No arguments passed to program")
	} else {
		// Load client with httpclient from http library
		client := client.ClientAPI{Client: &http.Client{}}

		worker.Init(flag.Args(), *numOfParallelChannels, client)
	}
}
