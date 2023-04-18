// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"fmt"
	"net/http"
	"registry/pkg/models/shared"
	"registry/pkg/utils"
	"time"
)

const (
	ServerProd string = "prod"
)

// ServerList contains the list of servers available to the SDK
var ServerList = map[string]string{
	ServerProd: "https://api.prod.speakeasyapi.dev",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// SDK - The Speakeasy API allows teams to manage common operations with their APIs
// https://docs.speakeasyapi.dev - The Speakeasy Platform Documentation
type SDK struct {
	// APIEndpoints - REST APIs for managing ApiEndpoint entities
	APIEndpoints *apiEndpoints
	// Apis - REST APIs for managing Api entities
	Apis *apis
	// Embeds - REST APIs for managing embeds
	Embeds *embeds
	// Metadata - REST APIs for managing Version Metadata entities
	Metadata *metadata
	// Requests - REST APIs for retrieving request information
	Requests *requests
	// Schemas - REST APIs for managing Schema entities
	Schemas *schemas

	// Non-idiomatic field names below are to namespace fields from the fields names above to avoid name conflicts
	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_security       *shared.Security
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

type SDKOption func(*SDK)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *SDK) {
		sdk._serverURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *SDK) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk._serverURL = serverURL
	}
}

// WithServer allows the overriding of the default server by name
func WithServer(server string) SDKOption {
	return func(sdk *SDK) {
		serverURL, ok := ServerList[server]
		if !ok {
			panic(fmt.Errorf("server %s not found", server))
		}

		WithServerURL(serverURL)(sdk)
	}
}

// WithTemplatedServer allows the overriding of the default server by name templated with the provided parameters
func WithTemplatedServer(server string, params map[string]string) SDKOption {
	return func(sdk *SDK) {
		serverURL, ok := ServerList[server]
		if !ok {
			panic(fmt.Errorf("server %s not found", server))
		}

		WithTemplatedServerURL(serverURL, params)(sdk)
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *SDK) {
		sdk._defaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *SDK) {
		sdk._security = &security
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *SDK {
	sdk := &SDK{
		_language:   "go",
		_sdkVersion: "1.0.0",
		_genVersion: "2.18.1",
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk._defaultClient == nil {
		sdk._defaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk._securityClient == nil {
		if sdk._security != nil {
			sdk._securityClient = utils.ConfigureSecurityClient(sdk._defaultClient, sdk._security)
		} else {
			sdk._securityClient = sdk._defaultClient
		}
	}

	if sdk._serverURL == "" {
		sdk._serverURL = ServerList[ServerProd]
	}

	sdk.APIEndpoints = newAPIEndpoints(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Apis = newApis(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Embeds = newEmbeds(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Metadata = newMetadata(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Requests = newRequests(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Schemas = newSchemas(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	return sdk
}
