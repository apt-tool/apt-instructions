package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

type CallBack func(url string) bool

func Post(url string) bool {
	log.Println("POST URL:>", url)

	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)

		return false
	}

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	body, _ := io.ReadAll(resp.Body)

	log.Println("response Body:", string(body))

	return resp.StatusCode == http.StatusServiceUnavailable
}

func Get(url string) bool {
	log.Println("GET URL:>", url)

	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)

		return false
	}

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	body, _ := io.ReadAll(resp.Body)

	log.Println("response Body:", string(body))

	return resp.StatusCode == http.StatusServiceUnavailable
}

func worker(input int, url string, cb CallBack) {
	wg := sync.WaitGroup{}
	signal := make(chan int)
	terminate := make(chan int)

	for i := 0; i < input; i++ {
		wg.Add(1)
		go func() {
			if cb(url) {
				signal <- 1
			}

			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		terminate <- 1
	}()

	select {
	case <-terminate:
		return
	case <-signal:
		os.Exit(1)
	}
}

func main() {
	log.SetOutput(os.Stdout)

	var (
		hostFlag      = flag.String("host", "localhost", "target host address")
		endpointsFlag = flag.String("endpoints", "/", "target specific endpoints")
		paramsFlag    = flag.String("params", "", "system parameters for testing")
	)

	flag.Parse()

	endpoints := strings.Split(*endpointsFlag, ",")
	paramSet := strings.Split(*paramsFlag, "&")

	params := make(map[string]string)

	for _, item := range paramSet {
		parts := strings.Split(item, "=")
		params[parts[0]] = parts[1]
	}

	log.Println(*hostFlag)
	log.Println(endpoints)
	log.Println(params)

	wg := sync.WaitGroup{}

	for _, endpoint := range endpoints {
		url := fmt.Sprintf("%s%s", *hostFlag, endpoint)

		wg.Add(2)

		go func() {
			worker(1000, url, Get)
			wg.Done()
		}()

		go func() {
			worker(1000, url, Post)
			wg.Done()
		}()

		wg.Wait()
	}

	log.Println("safe against dos attack!")

	os.Exit(0)
}
