// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package test1

import (
	"context"
	"fmt"
	"net/http"
	"test-1/pkg/models/operations"
	"test-1/pkg/models/shared"
	"test-1/pkg/utils"
)

// artefact - Artefacts can be registered with Humanitec. Continuous Integration (CI) pipelines notify Humanitec when a new version of an Artefact becomes available. Humanitec tracks the Artefact along with metadata about how it was built.
// <SchemaDefinition schemaRef="#/components/schemas/ArtefactRequest" />
type artefact struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newArtefact(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *artefact {
	return &artefact{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// DeleteOrgsOrgIDArtefactsArtefactID - Delete Artefact and all related Artefact Versions
// The specified Artefact and its Artefact Versions will be permanently deleted. Only Administrators can delete an Artefact.
func (s *artefact) DeleteOrgsOrgIDArtefactsArtefactID(ctx context.Context, request operations.DeleteOrgsOrgIDArtefactsArtefactIDRequest) (*operations.DeleteOrgsOrgIDArtefactsArtefactIDResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/orgs/{orgId}/artefacts/{artefactId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.defaultClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DeleteOrgsOrgIDArtefactsArtefactIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 204:
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.HumanitecErrorResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.HumanitecErrorResponse = out
		}
	}

	return res, nil
}

// GetOrgsOrgIDArtefacts - List all Artefacts.
// Returns the Artefacts registered with your organization. If no elements are found, an empty list is returned.
func (s *artefact) GetOrgsOrgIDArtefacts(ctx context.Context, request operations.GetOrgsOrgIDArtefactsRequest) (*operations.GetOrgsOrgIDArtefactsResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/orgs/{orgId}/artefacts", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.defaultClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetOrgsOrgIDArtefactsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []shared.ArtefactResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ArtefactResponses = out
		}
	}

	return res, nil
}
