// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"test-1/pkg/models/shared"
)

type PostOrgsOrgIDAppsAppIDEnvsEnvIDValuesRequest struct {
	// Definition of the new Shared Value.
	//
	//
	ValueCreatePayloadRequest shared.ValueCreatePayloadRequest `request:"mediaType=application/json"`
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

type PostOrgsOrgIDAppsAppIDEnvsEnvIDValuesResponse struct {
	ContentType string
	// Input not valid.
	//
	//
	HumanitecErrorResponse *shared.HumanitecErrorResponse
	StatusCode             int
	RawResponse            *http.Response
	// Shared Value successfully created.
	//
	//
	ValueResponse *shared.ValueResponse
}
