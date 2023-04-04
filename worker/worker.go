package worker

import (
	"fmt"
	c "myhttptool/client"
	"sync"
)

// Init sets everything up. creating the workerpool calling the md5 function and then printing the response
func Init(urlList []string, numOfParallelChannels int, client c.ClientAPI) {
	numOfUrl := len(urlList)
	outputs := make(chan c.RequestedPage, numOfUrl)
	urlChan := make(chan string, numOfUrl)
	finished := make(chan struct{})

	go setUrlList(urlList, urlChan)
	go PrintOutputs(finished, outputs)

	if numOfUrl < numOfParallelChannels {
		numOfParallelChannels = numOfUrl
	}
	setUpWorkerPool(urlChan, numOfParallelChannels, client, outputs)
	<-finished
}

// seturlList sets the urlList in the channel
func setUrlList(urlList []string, urlChan chan string) {
	for _, url := range urlList {
		urlChan <- url
	}
	close(urlChan)
}

// worker gets the response from the client
func worker(wg *sync.WaitGroup, urlList chan string, client c.ClientAPI, outputs chan c.RequestedPage) {
	for url := range urlList {
		result := client.GetMD5Page(url)
		outputs <- result
	}
	wg.Done()
}

// setUpWorkerPool makes the necessary workers
func setUpWorkerPool(urlList chan string, numOfParallelChannels int, client c.ClientAPI, outputs chan c.RequestedPage) {
	var wg sync.WaitGroup
	for i := 0; i < numOfParallelChannels; i++ {
		wg.Add(1)
		go worker(&wg, urlList, client, outputs)
	}

	wg.Wait()
	close(outputs)
}

// PrintOutputs prints the output of the pages to the console
func PrintOutputs(done chan struct{}, requestedPages chan c.RequestedPage) {
	for requestedPage := range requestedPages {
		fmt.Println(requestedPage.ToString())
	}
	done <- struct{}{}
}
