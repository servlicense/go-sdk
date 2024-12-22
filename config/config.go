package config

type Config struct {
	// Server points to an instance of servlicense
	Server string
	// Identifier
	Identifier string
	// Key is the API key
	ApiKey string
}

func New(server string) *Config {
	return &Config{
		Server: server,
	}
}

func WithAuth(server string, identifier string, key string) *Config {
	c := New(server)
	c.Identifier = identifier
	c.ApiKey = key
	return c
}
