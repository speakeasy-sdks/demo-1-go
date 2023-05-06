// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"test-1/pkg/models/shared"
)

type DeleteOrgsOrgIDResourcesDefsDefIDCriteriaCriteriaIDRequest struct {
	// The Matching Criteria ID.
	//
	//
	CriteriaID string `pathParam:"style=simple,explode=false,name=criteriaId"`
	// The Resource Definition ID.
	//
	//
	DefID string `pathParam:"style=simple,explode=false,name=defId"`
	// If set to `true`, the Matching Criteria is deleted immediately, even if this action affects existing Active Resources.
	//
	//
	Force *bool `queryParam:"style=form,explode=true,name=force"`
	// The Organization ID.
	//
	//
	OrgID string `pathParam:"style=simple,explode=false,name=orgId"`
}

type DeleteOrgsOrgIDResourcesDefsDefIDCriteriaCriteriaIDResponse struct {
	ContentType string
	// Internal application error.
	//
	//
	HumanitecErrorResponse *shared.HumanitecErrorResponse
	StatusCode             int
	RawResponse            *http.Response
	// One or more Active Resources reference the Resource Definition through this Matching Criteria.
	//
	//
	ResourceDefinitionChangeResponses []shared.ResourceDefinitionChangeResponse
	// The Resource Definition is not found.
	//
	//
	DeleteOrgsOrgIDResourcesDefsDefIDCriteriaCriteriaID404ApplicationJSONString *string
}
