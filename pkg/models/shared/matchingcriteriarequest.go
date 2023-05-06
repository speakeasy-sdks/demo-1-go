// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// MatchingCriteriaRequest - Matching Criteria are a set of rules used to choose which Resource Definition to use to provision a particular Resource Type.
//
// Matching criteria are made up in order of specificity with least specific first:
//
// - Environment Type (`env_type`)
//
// - Application (`app_id`)
//
// - Environment (`env_id`)
//
// - Resource (`res_id`)
//
// When selecting matching criteria, the most specific one is selected. In general, this means of all the Matching Criteria fully matching the context, the Matching Criteria Rule with the most specific element filled is chosen. If there is a tie, the next most specific elements are compared and so on until one is chosen.
//
// **NOTE:**
//
// Humanitec will reject the registration of matching criteria rules that duplicate rules already present for a Resource Type.
type MatchingCriteriaRequest struct {
	// (Optional) The ID of the Application that the Resources should belong to.
	AppID *string `json:"app_id,omitempty"`
	// (Optional) The ID of the Environment that the Resources should belong to. If `env_type` is also set, it must match the Type of the Environment for the Criteria to match.
	EnvID *string `json:"env_id,omitempty"`
	// (Optional) The Type of the Environment that the Resources should belong to. If `env_id` is also set, it must have an Environment Type that matches this parameter for the Criteria to match.
	EnvType *string `json:"env_type,omitempty"`
	// Matching Criteria ID
	ID *string `json:"id,omitempty"`
	// (Optional) The ID of the Resource in the Deployment Set. The ID is normally a `.` separated path to the definition in the set, e.g. `modules.my-module.externals.my-database`.
	ResID *string `json:"res_id,omitempty"`
}
