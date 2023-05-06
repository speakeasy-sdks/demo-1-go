// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"test-1/pkg/models/shared"
)

type PatchOrgsOrgIDArtefactsArtefactIDVersionsVersionIDRequest struct {
	// The Artefact Version Update Request. Only the field `archive` can be updated.
	//
	//
	UpdateArtefactVersionPayloadRequest shared.UpdateArtefactVersionPayloadRequest `request:"mediaType=application/json"`
	// The Artefact ID.
	//
	//
	ArtefactID string `pathParam:"style=simple,explode=false,name=artefactId"`
	// The organization ID.
	//
	//
	OrgID string `pathParam:"style=simple,explode=false,name=orgId"`
	// The Version ID.
	//
	//
	VersionID string `pathParam:"style=simple,explode=false,name=versionId"`
}

type PatchOrgsOrgIDArtefactsArtefactIDVersionsVersionIDResponse struct {
	// The updated Artefact Version.
	//
	//
	ArtefactVersionResponse *shared.ArtefactVersionResponse
	ContentType             string
	// One or more request parameters are missing or invalid, or the requested payload is not provided or malformed.
	//
	//
	HumanitecErrorResponse *shared.HumanitecErrorResponse
	StatusCode             int
	RawResponse            *http.Response
}
