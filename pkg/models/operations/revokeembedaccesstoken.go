// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"registry/pkg/models/shared"
)

type RevokeEmbedAccessTokenRequest struct {
	// The ID of the EmbedToken to revoke.
	TokenID string `pathParam:"style=simple,explode=false,name=tokenID"`
}

type RevokeEmbedAccessTokenResponse struct {
	ContentType string
	// Default error response
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
