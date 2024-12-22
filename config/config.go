package config

type Config struct {
	// Remote points to an instance of servlicense
	Remote string
	// Key signifies
	Key string
}

func New(remote string) *Config {
	return &Config{
		Remote: remote,
	}
}

func WithAuth(remote string, key string) *Config {
	c := New(remote)
	c.Key = key
	return c
}
