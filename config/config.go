package config

import "time"

type Config struct {
	// Server points to an instance of servlicense
	Server string
	// Identifier
	Identifier string
	// Key is the API key
	ApiKey string
	// Timeout defines the timeout for the httpclient
	Timeout time.Duration
}

func New(server string) *Config {
	return &Config{
		Server:  server,
		Timeout: time.Second * 5,
	}
}

func WithAuth(server string, identifier string, key string) *Config {
	c := New(server)
	c.Identifier = identifier
	c.ApiKey = key
	return c
}
