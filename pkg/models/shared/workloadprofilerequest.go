// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// WorkloadProfileRequest - Workload Profiles provide the baseline configuration for Workloads in Applications in Humanitec. Developers can configure various features of a workload profile to suit their needs. Examples of features might be `schedules` used in Kubernetes CronJobs or `ingress` which might be used to expose Pods controlled by a Kubernetes Deployment.
//
// Workloads in Humanitec are implemented as Helm Charts which must implement a specific schema.
type WorkloadProfileRequest struct {
	// Workload Profile ID
	ID string `json:"id"`
}
