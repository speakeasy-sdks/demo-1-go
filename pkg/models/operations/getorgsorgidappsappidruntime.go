// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"test-1/pkg/models/shared"
)

type GetOrgsOrgIDAppsAppIDRuntimeRequest struct {
	// The Application ID.
	//
	//
	AppID string `pathParam:"style=simple,explode=false,name=appId"`
	// Filter environments by ID (required). Up to 5 ids can be supplied per request.
	//
	//
	ID *string `queryParam:"style=form,explode=true,name=id"`
	// The Organization ID.
	//
	//
	OrgID string `pathParam:"style=simple,explode=false,name=orgId"`
}

type GetOrgsOrgIDAppsAppIDRuntimeResponse struct {
	ContentType string
	// A list of the RuntimeInfo of the environments specified by the id parameter.
	//
	//
	EnvironmentRuntimeInfoResponses []shared.EnvironmentRuntimeInfoResponse
	StatusCode                      int
	RawResponse                     *http.Response
}
