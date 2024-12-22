package client

import (
	"context"

	"github.com/servlicense/go-sdk/config"
	"github.com/servlicense/servlicense/api/models"
	"github.com/servlicense/servlicense/api/types"
)

// Client represents the main interaction with servlicense and holds all
// methods to call servlicense
type Client struct{}

// New creates an instance of Client configured with config
func New(config config.Config) *Client {
	return nil
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

// TODO: should this be available at this point?
func (c *Client) CreateKey(ctx context.Context) {}

// fetch makes an authenticated request to the servlicense API
func (c *Client) MakeRequest(ctx context.Context, endpoint string, method string, body interface{}) ([]byte, error) {

	return nil, nil
}
