package client

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"net/http"
	"strings"

	"github.com/servlicense/go-sdk/config"
	"github.com/servlicense/servlicense/api/models"
	"github.com/servlicense/servlicense/api/types"
)

// Client represents the main interaction with servlicense and holds all
// methods to call servlicense
type Client struct {
	config config.Config
}

// New creates an instance of Client configured with config
func New(config config.Config) *Client {
	return &Client{config}
}

// Me returns the scopes available for the currently authenticated api key, requires authentication
func (c *Client) Me(ctx context.Context) ([]types.ApiKeyScope, error) {
	return nil, nil
}

// CreateLicense creates a new license, requires authentication and admin/create_license scope
func (c *Client) CreateLicense(ctx context.Context) (bool, error) {
	return false, nil
}

// CheckKey returns if the given key is valid, does not require authentication
func (c *Client) CheckLicense(ctx context.Context, license string) (bool, error) {
	return false, nil
}

// ListLicenses lists the available licenses, requires authentication and admin/list_licenses scope
func (c *Client) ListLicenses(ctx context.Context) ([]models.Apikey, error) {
	return nil, nil
}

func (c *Client) CreateApiKey(ctx context.Context) {}

// fetch makes an authenticated request to the servlicense API
func (c *Client) makeRequest(ctx context.Context, endpoint string, method string, body []byte) ([]byte, error) {
	client := http.Client{
		Timeout: c.config.Timeout,
	}
	writer := bytes.NewReader(body)
	req, err := http.NewRequest(method, strings.Join([]string{c.config.Server, endpoint}, "/"), writer)
	if err != nil {
		return nil, err
	}

	byteBuff := bytes.Buffer{}
	byteBuff.WriteString(c.config.Identifier)
	byteBuff.WriteRune(':')
	byteBuff.WriteString(c.config.ApiKey)

	req.Header.Set("Authorization", base64.RawStdEncoding.EncodeToString(byteBuff.Bytes()))

	res, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return io.ReadAll(res.Body)
}
