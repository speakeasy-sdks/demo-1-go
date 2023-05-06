// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type EnvironmentBaseRequest struct {
	// The ID the Environment is referenced as.
	ID string `json:"id"`
	// The Human-friendly name for the Environment.
	Name string `json:"name"`
	// The Environment Type. This is used for organizing and managing Environments.
	Type string `json:"type"`
}
