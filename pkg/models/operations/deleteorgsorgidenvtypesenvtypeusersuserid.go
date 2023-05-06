// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"test-1/pkg/models/shared"
)

type DeleteOrgsOrgIDEnvTypesEnvTypeUsersUserIDRequest struct {
	// The Environment Type.
	//
	//
	EnvType string `pathParam:"style=simple,explode=false,name=envType"`
	// The Organization ID.
	//
	//
	OrgID string `pathParam:"style=simple,explode=false,name=orgId"`
	// The User ID
	//
	//
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
}

type DeleteOrgsOrgIDEnvTypesEnvTypeUsersUserIDResponse struct {
	ContentType string
	// The request was invalid or the payload malformed.
	//
	//
	HumanitecErrorResponse *shared.HumanitecErrorResponse
	StatusCode             int
	RawResponse            *http.Response
}
