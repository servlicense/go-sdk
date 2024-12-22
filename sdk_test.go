package gosdk

import (
	"testing"

	"github.com/servlicense/go-sdk/client"
	"github.com/servlicense/go-sdk/config"
)

const remote = "http://localhost:8080"
const identifier = "187"
const key = "123456"

func TestClientInit(t *testing.T) {
	conf := config.WithAuth(remote, identifier, key)
	c := client.New(*conf)
	if c == nil {
		t.Fatal("Failed to initialize client")
	}
}
