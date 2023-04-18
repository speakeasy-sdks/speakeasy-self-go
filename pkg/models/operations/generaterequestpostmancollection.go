// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"registry/pkg/models/shared"
)

type GenerateRequestPostmanCollectionRequest struct {
	// The ID of the request to retrieve.
	RequestID string `pathParam:"style=simple,explode=false,name=requestID"`
}

type GenerateRequestPostmanCollectionResponse struct {
	ContentType string
	// Default error response
	Error *shared.Error
	// OK
	PostmanCollection []byte
	StatusCode        int
	RawResponse       *http.Response
}
