// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ControllerResponse - Controller represents deployment, stateful set etc
type ControllerResponse struct {
	Kind     string             `json:"kind"`
	Message  string             `json:"message"`
	Pods     []PodStateResponse `json:"pods"`
	Replicas int64              `json:"replicas"`
	Revision int64              `json:"revision"`
	Status   string             `json:"status"`
}
