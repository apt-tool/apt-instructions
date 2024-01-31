package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/http2"
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

	return resp.StatusCode == http.StatusOK
}

func dirListing(url string) bool {
	paths := []string{
		"a/b",
		"../a/b",
		"../a/../b",
		"a//c",
		"/",
		"a/c/.",
		"/var/",
		"/etc/",
		"/dev/",
		"/usr/",
		"/lib/",
		"/proc/",
		"/Users/",
		"/System",
	}

	for _, subPath := range paths {
		host := fmt.Sprintf("%s/%s", url, subPath)

		if Get(host) {
			return true
		}
	}

	return false
}

func protocol(url string) bool {
	// Create an HTTP/2-enabled client
	client := &http.Client{
		Transport: &http2.Transport{},
	}

	// Create an HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating request:", err)

		return false
	}

	// Set HTTP/2 protocol for the request
	req.Header.Set("Connection", "Upgrade, HTTP2-Settings")
	req.Header.Set("Upgrade", "h2c")
	req.Header.Set("HTTP2-Settings", "")

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)

		return true
	}

	defer resp.Body.Close()

	// Print the response status code
	log.Println("Response Status:", resp.Status)

	return true
}

func fileFetch(url string) bool {
	paths := []string{
		"/bin",
		"/boot",
		"/etc/bashrc",
		"/etc/fstab",
		"/etc/group",
		"/etc/hostname",
		"/etc/hosts",
		"/etc/passwd",
		"/etc/profile",
		"/etc/resolv.conf",
		"/etc/shadow",
		"/etc/sudoers",
		"/proc",
		"/sys",
		"/usr/bin",
		"/usr/share/doc",
		"/var/lib/dpkg/status",
		"/var/lib/rpm",
		"/var/log",
		"/var/spool/mail",
		"/etc/apache2/apache2.conf",
		"/etc/apache2/sites-available/",
		"/var/log/apache2/access.log",
		"/var/log/apache2/error.log",
		"/etc/apache2/mods-available/",
		"/etc/apache2/ssl/",
		"/var/www/html/",
		"/etc/apache2/conf-available/",
		"/etc/apache2/mime.types",
		"/etc/apache2/httpd.conf.dpkg-dist",
		"/etc/nginx/nginx.conf",
		"/etc/nginx/sites-available/",
		"/var/log/nginx/access.log",
		"/var/log/nginx/error.log",
		"/etc/nginx/conf.d/",
		"/etc/nginx/mime.types",
		"/etc/nginx/fastcgi_params",
		"/etc/nginx/scgi_params",
		"/etc/nginx/uwsgi_params",
		"C:\\Windows\\System32\\config\\SYSTEM",
		"C:\\Windows\\System32\\config\\SOFTWARE",
		"C:\\Windows\\System32\\config\\SAM",
		"C:\\Windows\\System32\\config\\SECURITY",
		"C:\\Windows\\System32\\config\\DEFAULT",
		"C:\\Windows\\System32\\drivers\\etc\\hosts",
		"C:\\Windows\\System32\\drivers\\etc\\services",
		"C:\\Windows\\System32\\drivers\\etc\\networks",
		"C:\\Windows\\System32\\drivers\\etc\\lmhosts",
		"C:\\Windows\\System32\\GroupPolicy\\",
		"C:\\Windows\\System32\\GroupPolicyUsers\\",
		"C:\\Windows\\System32\\Tasks\\",
		"C:\\Windows\\System32\\drivers\\etc\\protocol",
		"C:\\Windows\\System32\\taskschd.msc",
		"C:\\Windows\\System32\\inetsrv\\config\\applicationHost.config",
	}

	for _, subPath := range paths {
		host := fmt.Sprintf("%s/%s", url, subPath)

		if Get(host) {
			return true
		}
	}

	return false
}

func main() {
	log.SetOutput(os.Stdout)

	var (
		hostFlag      = flag.String("host", "localhost", "target host address")
		endpointsFlag = flag.String("endpoints", "/database/list/2", "target specific endpoints")
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

	callbacks := []CallBack{
		fileFetch,
		dirListing,
		protocol,
		fileSend,
	}

	for _, callback := range callbacks {
		for _, endpoint := range endpoints {
			url := fmt.Sprintf("%s%s", *hostFlag, endpoint)

			if callback(url) {
				log.Println("Found a bug in error-handling attack!")

				os.Exit(1)
			}
		}
	}

	log.Println("safe against error-handling attacks!")

	os.Exit(0)
}
