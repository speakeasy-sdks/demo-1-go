// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"test-1/pkg/models/shared"
)

type PostOrgsOrgIDEnvTypesEnvTypeUsersRequest struct {
	// The user ID and the role
	//
	//
	UserRoleRequest shared.UserRoleRequest `request:"mediaType=application/json"`
	// The Environment Type.
	//
	//
	EnvType string `pathParam:"style=simple,explode=false,name=envType"`
	// The Organization ID.
	//
	//
	OrgID string `pathParam:"style=simple,explode=false,name=orgId"`
}

type PostOrgsOrgIDEnvTypesEnvTypeUsersResponse struct {
	ContentType string
	// The request was invalid or the payload malformed.
	//
	//
	HumanitecErrorResponse *shared.HumanitecErrorResponse
	StatusCode             int
	RawResponse            *http.Response
	// The user ID and associated role.
	//
	//
	UserRoleResponse *shared.UserRoleResponse
}
