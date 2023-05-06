// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"test-1/pkg/models/shared"
)

type PostOrgsOrgIDAppsAppIDEnvsEnvIDRulesRequest struct {
	// The definition of the Automation Rule.
	//
	//
	AutomationRuleRequest shared.AutomationRuleRequest `request:"mediaType=application/json"`
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

type PostOrgsOrgIDAppsAppIDEnvsEnvIDRulesResponse struct {
	// The AutomationRule
	//
	//
	AutomationRuleResponse *shared.AutomationRuleResponse
	ContentType            string
	// The input was not a valid Automation Rule.
	//
	//
	ErrorInfoResponse *shared.ErrorInfoResponse
	StatusCode        int
	RawResponse       *http.Response
}
