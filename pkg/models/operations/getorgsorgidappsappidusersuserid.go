// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"test-1/pkg/models/shared"
)

type GetOrgsOrgIDAppsAppIDUsersUserIDRequest struct {
	// The Application ID.
	//
	//
	AppID string `pathParam:"style=simple,explode=false,name=appId"`
	// The Organization ID.
	//
	//
	OrgID string `pathParam:"style=simple,explode=false,name=orgId"`
	// The User ID
	//
	//
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
}

type GetOrgsOrgIDAppsAppIDUsersUserIDResponse struct {
	ContentType string
	// The request was invalid or the payload malformed.
	//
	//
	HumanitecErrorResponse *shared.HumanitecErrorResponse
	StatusCode             int
	RawResponse            *http.Response
	// The information on the user.
	//
	//
	UserRoleResponse *shared.UserRoleResponse
}
