package main

import (
	"fmt"
	"net/http"
	"sync"
)

func webGetWorker(in <-chan []byte, grt int, wg *sync.WaitGroup) {
	for {
		url := string(<-in)
		if url == "" {
			continue
		}
		fmt.Printf("GoRoutine #%d getting data from url: %s \n", grt, url)

		res, err := http.Get(url)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("GET %s: %d\n\n", url, res.StatusCode)
			wg.Done()
		}
	}
}

func main() {
	var wg sync.WaitGroup
	work := make(chan []byte, 6144)

	numWorker := 2
	for i := 0; i < numWorker; i++ {
		go webGetWorker(work, i+1, &wg)
	}

	urls := [4]string{"http://reddit.com", "http://google.com", "http://infobae.com", "http://youtube.com"}

	for _, url := range urls {
		wg.Add(1)
		fmt.Printf("Please get %s \n", url)
		work <- []byte(url)
	}

	wg.Wait()
	fmt.Println("The end!!")
}
