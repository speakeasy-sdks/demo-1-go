// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"test-1/pkg/models/shared"
)

type GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployIDErrorsRequest struct {
	// The Application ID.
	//
	//
	AppID string `pathParam:"style=simple,explode=false,name=appId"`
	// The Deployment ID.
	//
	//
	DeployID string `pathParam:"style=simple,explode=false,name=deployId"`
	// The Environment ID.
	//
	//
	EnvID string `pathParam:"style=simple,explode=false,name=envId"`
	// The Organization ID.
	//
	//
	OrgID string `pathParam:"style=simple,explode=false,name=orgId"`
}

type GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployIDErrorsResponse struct {
	ContentType string
	// A List of deployment errors, could be empty.
	//
	//
	DeploymentErrorResponses []shared.DeploymentErrorResponse
	StatusCode               int
	RawResponse              *http.Response
}
