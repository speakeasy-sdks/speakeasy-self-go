// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"registry/pkg/models/shared"
)

type RegisterSchemaRequestBodyFile struct {
	Content []byte `multipartForm:"content"`
	File    string `multipartForm:"name=file"`
}

// RegisterSchemaRequestBody - The schema file to upload provided as a multipart/form-data file segment.
type RegisterSchemaRequestBody struct {
	File RegisterSchemaRequestBodyFile `multipartForm:"file"`
}

type RegisterSchemaRequest struct {
	// The schema file to upload provided as a multipart/form-data file segment.
	RequestBody RegisterSchemaRequestBody `request:"mediaType=multipart/form-data"`
	// The ID of the Api to get the schema for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The version ID of the Api to delete metadata for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type RegisterSchemaResponse struct {
	ContentType string
	// Default error response
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
