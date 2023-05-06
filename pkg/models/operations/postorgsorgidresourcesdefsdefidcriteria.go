// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"test-1/pkg/models/shared"
)

type PostOrgsOrgIDResourcesDefsDefIDCriteriaRequest struct {
	// Matching Criteria rules.
	//
	//
	MatchingCriteriaRuleRequest shared.MatchingCriteriaRuleRequest `request:"mediaType=application/json"`
	// The Resource Definition ID.
	//
	//
	DefID string `pathParam:"style=simple,explode=false,name=defId"`
	// The Organization ID.
	//
	//
	OrgID string `pathParam:"style=simple,explode=false,name=orgId"`
}

type PostOrgsOrgIDResourcesDefsDefIDCriteriaResponse struct {
	ContentType string
	// One or more request parameters is missing or invalid.
	//
	//
	HumanitecErrorResponse *shared.HumanitecErrorResponse
	// The newly added Matching Criteria details.
	//
	//
	MatchingCriteriaResponse *shared.MatchingCriteriaResponse
	StatusCode               int
	RawResponse              *http.Response
}
