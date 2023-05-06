// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PatchOrgsOrgIDAppsAppIDEnvsEnvIDRuntimeReplicasRequest struct {
	// map of replicas by modules.
	//
	//
	RequestBody map[string]int64 `request:"mediaType=application/json"`
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

type PatchOrgsOrgIDAppsAppIDEnvsEnvIDRuntimeReplicasResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
