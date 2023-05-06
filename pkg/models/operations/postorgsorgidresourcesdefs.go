// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"test-1/pkg/models/shared"
)

type PostOrgsOrgIDResourcesDefsRequest struct {
	// The Resource Definition details.
	//
	//
	CreateResourceDefinitionRequestRequest shared.CreateResourceDefinitionRequestRequest `request:"mediaType=application/json"`
	// The Organization ID.
	//
	//
	OrgID string `pathParam:"style=simple,explode=false,name=orgId"`
}

type PostOrgsOrgIDResourcesDefsResponse struct {
	ContentType string
	// One or more request parameters is missing or invalid.
	//
	//
	HumanitecErrorResponse *shared.HumanitecErrorResponse
	// The newly created Resources Definition details.
	//
	//
	ResourceDefinitionResponse *shared.ResourceDefinitionResponse
	StatusCode                 int
	RawResponse                *http.Response
}
