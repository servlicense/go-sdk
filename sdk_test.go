package gosdk

import (
	"context"
	"fmt"
	"testing"

	"github.com/servlicense/go-sdk/client"
	"github.com/servlicense/go-sdk/config"
)

func TestClientInit(t *testing.T) {
	conf := config.WithAuth("https://api.servlicense.sh", "2132295360259817472", "6be29a05-b754-4fab-a490-337129bf52cf")
	c := client.New(*conf)
	if c == nil {
		t.Fatal("Failed to initialize client")
	}
}

func TestReadme(t *testing.T) {
	client := client.New(config.Config{
		Server:     "https://api.servlicense.sh",
		Identifier: "2132295360259817472",
		ApiKey:     "6be29a05-b754-4fab-a490-337129bf52cf",
	})
	// or config := config.WithAuth("https://api.servlicense.sh", "2132295360259817472", "6be29a05-b754-4fab-a490-337129bf52cf")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	isValid, err := client.CheckLicense(ctx, "8b3b3b7b-7b7b-4b7b-8b7b-7b7b7b7b7b7b")
	if err == nil {
		t.Fatal("Failed to check license", err)
	}

	fmt.Println(isValid)
}
