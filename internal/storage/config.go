package storage

import "strings"

type Config struct {
	Endpoint string
	Bucket   string
	Access   string
	Secret   string
	SSL      bool
}

func LoadConfig(input string) Config {
	cfg := Config{}

	parts := strings.Split(input, "@")

	credentials := strings.Split(parts[0], ":")
	cfg.Access = credentials[0]
	cfg.Secret = credentials[1]

	host := strings.Split(parts[1], "&")
	cfg.Endpoint = host[0]
	cfg.Bucket = host[1]

	if host[2] == "true" {
		cfg.SSL = true
	} else {
		cfg.SSL = false
	}

	return cfg
}
