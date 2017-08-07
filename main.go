package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup
var data chan string

func fetchURL(count int) {
	fmt.Printf("URL fetcher #%d has started ...\n", count)
	defer func() {
		fmt.Printf("URL fetcher #%d closing down ...\n", count)
		waitGroup.Done()
	}()
	for {
		url, ok := <-data
		var method string
		var req *http.Request
		var err error

		if !ok {
			fmt.Println("The channel is closed!")
			break
		}

		client := &http.Client{}

		// Random method: mix things up a bit!
		switch rand.Intn(7) {
		case 0:
			method = "HEAD"
			req, err = http.NewRequest(method, url, nil)
		case 1:
			method = "POST"
			randomJson := []byte(`{"why":"I do not know.."}`)
			req, err = http.NewRequest(method, url, bytes.NewBuffer(randomJson))
		case 2:
			method = "PUT"
			randomJson := []byte(`{"why":"Ooh! Ooh! I know!"}`)
			req, err = http.NewRequest(method, url, bytes.NewBuffer(randomJson))
		case 3:
			method = "PATCH"
			randomJson := []byte(`{"why":"Ah damn, I lost it again."}`)
			req, err = http.NewRequest(method, url, bytes.NewBuffer(randomJson))
		case 4:
			method = "HELP"
			randomJson := []byte(`{"why":"Really. I lost it. Send help. Kthxbye."}`)
			req, err = http.NewRequest(method, url, bytes.NewBuffer(randomJson))
		case 5:
			method = "MATTIAS"
			randomJson := []byte(`{"why":"Now I remember. The METHOD can be _anything_."}`)
			req, err = http.NewRequest(method, url, bytes.NewBuffer(randomJson))
		default:
			method = "GET"
			req, err = http.NewRequest(method, url, nil)
		}

		if err != nil {
			fmt.Println(err)
		}

		client.Do(req)

		// Don't care about exit codes, just get that HTTP call out the door
		fmt.Printf("#%d: fetched %s \n", count, url)
	}
}

func main() {
	var baseURL string
	var concurrentRequests int
	var maxRequests int
	rand.Seed(time.Now().Unix())

	if len(os.Args) != 4 {
		log.Fatalf("%s CONCURRENCY MAXREQUESTS http://yourtarget.tld\nIe: %s 10 50 https://www.google.com\nThis starts 50 HTTP calls to the URL, doing 10 at a time.\n", os.Args[0], os.Args[0])
	}

	_, err := url.ParseRequestURI(os.Args[3])
	if err != nil {
		log.Fatalf("Parameter not a URL. Error: %s", err)
	}
	concurrentRequests, _ = strconv.Atoi(os.Args[1])
	maxRequests, _ = strconv.Atoi(os.Args[2])
	baseURL = os.Args[3]

	// Loop with some random parameters
	data = make(chan string)
	fmt.Printf("Heads-up, starting %d requests to %s, doing %d at a time.\n\n", maxRequests, baseURL, concurrentRequests)

	// Start X amount of concurrent fetchers
	for i := 1; i < concurrentRequests+1; i++ {
		waitGroup.Add(1)
		go fetchURL(i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("\nStarting all requests in 5 seconds ...")
	time.Sleep(5 * time.Second)

	// No fetch X amount of URLs using those fetchers
	for i := 0; i < maxRequests; i++ {
		// Randomise the URL a bit, bypass caching
		data <- (baseURL)
	}
}
