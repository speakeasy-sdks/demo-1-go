// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"test-1/pkg/models/shared"
)

type GetOrgsOrgIDAppsAppIDEnvsEnvIDResourcesRequest struct {
	// The Application ID.
	//
	//
	AppID string `pathParam:"style=simple,explode=false,name=appId"`
	// The Environment ID.
	//
	//
	EnvID string `pathParam:"style=simple,explode=false,name=envId"`
	// The Organization ID.
	//
	//
	OrgID string `pathParam:"style=simple,explode=false,name=orgId"`
}

type GetOrgsOrgIDAppsAppIDEnvsEnvIDResourcesResponse struct {
	// A possibly empty list of Active Resources.
	//
	//
	ActiveResourceResponses []shared.ActiveResourceResponse
	ContentType             string
	// Internal application error.
	//
	//
	HumanitecErrorResponse *shared.HumanitecErrorResponse
	StatusCode             int
	RawResponse            *http.Response
}
