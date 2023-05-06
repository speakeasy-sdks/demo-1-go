// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ResourceAccountResponse - ResourceAccount represents the account being used to access a resource.
//
// Resource Accounts hold credentials that are required to provision and manage resources.
type ResourceAccountResponse struct {
	// The timestamp of when the account was created.
	CreatedAt string `json:"created_at"`
	// The ID of the user who created the account.
	CreatedBy string `json:"created_by"`
	// Unique identifier for the account (in scope of the organization it belongs to).
	ID string `json:"id"`
	// Indicates if this account is being used (referenced) by any resource definition.
	IsUsed bool `json:"is_used"`
	// Display name.
	Name string `json:"name"`
	// The type of the account
	Type string `json:"type"`
}
