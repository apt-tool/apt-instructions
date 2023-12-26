package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type CallBack func(url string) bool

func protocol(url string) bool {
	return false
}

func fileSend(url string) bool {
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
		protocol,
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
