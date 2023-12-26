package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type CallBack func(url string) bool

func fileSend(url string) bool {
	largeUrls := []string{
		"https://speed.hetzner.de/1GB.bin",
		"https://speedtest-sgp1.digitalocean.com/5gb.test",
		"https://speed.hetzner.de/10GB.bin",
	}

	for _, file := range largeUrls {
		data, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}

		defer data.Close()

		req, err := http.NewRequest("PUT", url, data)

		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Content-Type", "text/plain")

		client := &http.Client{}

		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		if res.StatusCode == 200 {
			return true
		}
	}

	return false
}

func dirListing(url string) bool {
	return false
}

func fileFetch(url string) bool {
	return false
}

func path(url string) bool {
	return false
}

func headers(url string) bool {
	return false
}

func main() {
	var (
		hostFlag      = flag.String("host", "localhost", "target host address")
		endpointsFlag = flag.String("endpoints", "/", "target pathes")
	)

	flag.Parse()

	endpoints := strings.Split(*endpointsFlag, ",")
	callbacks := []CallBack{
		headers,
		path,
		fileFetch,
		dirListing,
		fileSend,
	}

	for _, callback := range callbacks {
		for _, endpoint := range endpoints {
			url := fmt.Sprintf("%s%s", *hostFlag, endpoint)

			if callback(url) {
				os.Exit(1)
			}
		}
	}

	os.Exit(0)
}
