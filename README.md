# servlicense/go-sdk

## Installation

## Usage

```go
package main

import (
	"context"
	"fmt"

	"github.com/servlicense/go-sdk/client"
	"github.com/servlicense/go-sdk/config"
)

func main() {
	cl := client.New(config.Config{
		Server:     "https://api.servlicense.sh",
		Identifier: "2132295360259817472",
		ApiKey:     "6be29a05-b754-4fab-a490-337129bf52cf",
	})
	// or config := config.WithAuth("https://api.servlicense.sh", "2132295360259817472", "6be29a05-b754-4fab-a490-337129bf52cf")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	isValid, err := cl.CheckLicense(ctx, "8b3b3b7b-7b7b-4b7b-8b7b-7b7b7b7b7b7b")
	if err == nil {
		t.Fatal("Failed to check license", err)
	}
	fmt.Println(isValid)
}
```
