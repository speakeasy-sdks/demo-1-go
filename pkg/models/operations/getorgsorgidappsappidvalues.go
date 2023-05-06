// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"test-1/pkg/models/shared"
)

type GetOrgsOrgIDAppsAppIDValuesRequest struct {
	// The Application ID.
	//
	//
	AppID string `pathParam:"style=simple,explode=false,name=appId"`
	// The Organization ID.
	//
	//
	OrgID string `pathParam:"style=simple,explode=false,name=orgId"`
}

type GetOrgsOrgIDAppsAppIDValuesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// A possibly empty list of Values.
	//
	//
	ValueResponses []shared.ValueResponse
}
