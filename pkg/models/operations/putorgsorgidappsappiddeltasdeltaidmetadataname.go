// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"test-1/pkg/models/shared"
)

type PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataNameRequest struct {
	// The new name.(NOTE: The string must still be JSON encoded.)
	//
	//
	RequestBody string `request:"mediaType=application/json"`
	// The Application ID.
	//
	//
	AppID string `pathParam:"style=simple,explode=false,name=appId"`
	// ID of the Deployment Delta.
	//
	//
	DeltaID string `pathParam:"style=simple,explode=false,name=deltaId"`
	// The Organization ID.
	//
	//
	OrgID string `pathParam:"style=simple,explode=false,name=orgId"`
}

type PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataNameResponse struct {
	ContentType string
	// The request was invalid.
	//
	//
	HumanitecErrorResponse *shared.HumanitecErrorResponse
	StatusCode             int
	RawResponse            *http.Response
	// No Deployment Delta with ID `deltaId` found in Application.
	//
	//
	PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataName404ApplicationJSONString *string
}
