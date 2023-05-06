// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"test-1/pkg/models/shared"
)

type GetOrgsOrgIDImagesRequest struct {
	// The organization ID.
	//
	//
	OrgID string `pathParam:"style=simple,explode=false,name=orgId"`
}

type GetOrgsOrgIDImagesResponse struct {
	ContentType string
	// Possibly empty list of Container Images.
	//
	//
	ImageResponses []shared.ImageResponse
	StatusCode     int
	RawResponse    *http.Response
}
