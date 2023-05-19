// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package test1

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"test-1/pkg/models/operations"
	"test-1/pkg/models/shared"
	"test-1/pkg/utils"
)

// set - A Deployment Set (or just "Set") defines all of the non-Environment specific configuration for Modules and External Resources. Each of these Modules or External Resources has a unique name.
//
// Deployment Sets are immutable and their ID is a cryptographic hash of their content. This means that a Deployment Set can be globally identified based on its ID. It also means that referencing a Deployment Set by ID will always return the same Deployment Set.
//
// Deployment Sets cannot be created directly, instead they are created by applying a Deployment Delta to an existing Deployment Set.
//
// **Basic Structure**
//
// ```
//
//	{
//	  "id": <ID of the Deployment Set>,
//	  "modules" : {
//	    <ID of Module> : {
//	      "profile": <Defines how the optional "spec" property is interpreted>
//	      "spec": {
//	        <Properties that depend on the "profile" property.>
//	      }
//	      "externals": {
//	        <External Resource Name> : {
//	          "type": <Resource Type>,
//	          "params": {
//	            <Properties which parametrize the resource depending on the Resource Type.>
//	          }
//	        }
//	      }
//	    }
//	  }
//	}
//
// ```
//
// For details about how the Humanitec provided profiles work, see (Deployment Set Profiles)[].
// <SchemaDefinition schemaRef="#/components/schemas/SetRequest" />
type set struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newSet(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *set {
	return &set{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// GetSets - Get all Deployment Sets
func (s *set) GetSets(ctx context.Context, request operations.GetSetsRequest) (*operations.GetSetsResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/orgs/{orgId}/apps/{appId}/sets", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

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

	res := &operations.GetSetsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []shared.SetResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.SetResponses = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			out := string(data)
			res.GetSets404ApplicationJSONString = &out
		}
	}

	return res, nil
}

// GetOrgsOrgIDAppsAppIDSetsSetID - Get a Deployment Set
func (s *set) GetOrgsOrgIDAppsAppIDSetsSetID(ctx context.Context, request operations.GetOrgsOrgIDAppsAppIDSetsSetIDRequest) (*operations.GetOrgsOrgIDAppsAppIDSetsSetIDResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/orgs/{orgId}/apps/{appId}/sets/{setId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

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

	res := &operations.GetOrgsOrgIDAppsAppIDSetsSetIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetOrgsOrgIDAppsAppIDSetsSetID200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetOrgsOrgIDAppsAppIDSetsSetID200ApplicationJSONOneOf = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			out := string(data)
			res.GetOrgsOrgIDAppsAppIDSetsSetID404ApplicationJSONString = &out
		}
	}

	return res, nil
}

// GetOrgsOrgIDAppsAppIDSetsSetIDDiffSourceSetID - Get the difference between 2 Deployment Sets
func (s *set) GetOrgsOrgIDAppsAppIDSetsSetIDDiffSourceSetID(ctx context.Context, request operations.GetOrgsOrgIDAppsAppIDSetsSetIDDiffSourceSetIDRequest) (*operations.GetOrgsOrgIDAppsAppIDSetsSetIDDiffSourceSetIDResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/orgs/{orgId}/apps/{appId}/sets/{setId}/diff/{sourceSetId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

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

	res := &operations.GetOrgsOrgIDAppsAppIDSetsSetIDDiffSourceSetIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.PlainDeltaResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PlainDeltaResponse = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			out := string(data)
			res.GetOrgsOrgIDAppsAppIDSetsSetIDDiffSourceSetID404ApplicationJSONString = &out
		}
	}

	return res, nil
}

// PostOrgsOrgIDAppsAppIDSetsSetID - Apply a Deployment Delta to a Deployment Set
func (s *set) PostOrgsOrgIDAppsAppIDSetsSetID(ctx context.Context, request operations.PostOrgsOrgIDAppsAppIDSetsSetIDRequest) (*operations.PostOrgsOrgIDAppsAppIDSetsSetIDResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/orgs/{orgId}/apps/{appId}/sets/{setId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "DeltaRequest", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/json;q=0.7, application/json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	req.Header.Set("Content-Type", reqContentType)

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

	res := &operations.PostOrgsOrgIDAppsAppIDSetsSetIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			out := string(data)
			res.PostOrgsOrgIDAppsAppIDSetsSetID200ApplicationJSONString = &out
		}
	case httpRes.StatusCode == 400:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.HumanitecErrorResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.HumanitecErrorResponse = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			out := string(data)
			res.PostOrgsOrgIDAppsAppIDSetsSetID404ApplicationJSONString = &out
		}
	}

	return res, nil
}
