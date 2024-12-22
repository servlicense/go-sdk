package gosdk

import (
	"testing"

	"github.com/servlicense/go-sdk/client"
	"github.com/servlicense/go-sdk/config"
)

const remote = "http://localhost:8080"
const key = "1234:test"

func TestClientInit(t *testing.T) {
	conf := config.WithAuth(remote, key)
	client.New(*conf)
}
