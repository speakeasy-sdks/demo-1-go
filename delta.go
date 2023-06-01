// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package test1

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"test-1/pkg/models/operations"
	"test-1/pkg/models/shared"
	"test-1/pkg/utils"
)

// delta - A Deployment Delta (or just "Delta") describes the changes that must be applied to one Deployment Set to generate another Deployment Set. Deployment Deltas are the only way to create new Deployment Sets.
//
// Deployment Deltas can be created fully formed or combined together via PATCHing. They can also be generated from the difference between two Deployment Sets.
//
// **Basic Structure**
//
// ```
//
//	{
//	  "id": <ID of the Deployment Delta.>,
//	  "metadata": {
//	    <Properties such as who created the Delta, which Environment it is associated to and a Human-friendly name>
//	  }
//	  "modules" : {
//	    "add" : {
//	      <ID of Module to add to the Deployment Set> : {
//	        <An entire Modules object>
//	      }
//	    },
//	    "remove": [
//	      <An array of Module IDs that should be removed from the Deployment Set>
//	    ],
//	   "update": {
//	      <ID of Module already in the Set to be updated> : [
//	        <An array of JSON Patch (Search Results (RFC 6902) objects scoped to the module>
//	      ]
//	    }
//	  }
//	}
//
// ```
// <SchemaDefinition schemaRef="#/components/schemas/DeltaRequest" />
type delta struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newDelta(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *delta {
	return &delta{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// GetDelta - Fetch an existing Delta
func (s *delta) GetDelta(ctx context.Context, request operations.GetDeltaRequest) (*operations.GetDeltaResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/orgs/{orgId}/apps/{appId}/deltas/{deltaId}", request, nil)
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

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetDeltaResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.DeltaResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.DeltaResponse = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			out := string(rawBody)
			res.GetDelta404ApplicationJSONString = &out
		}
	}

	return res, nil
}

// GetOrgsOrgIDAppsAppIDDeltas - List Deltas in an Application
func (s *delta) GetOrgsOrgIDAppsAppIDDeltas(ctx context.Context, request operations.GetOrgsOrgIDAppsAppIDDeltasRequest) (*operations.GetOrgsOrgIDAppsAppIDDeltasResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/orgs/{orgId}/apps/{appId}/deltas", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
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

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetOrgsOrgIDAppsAppIDDeltasResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []shared.DeltaResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.DeltaResponses = out
		}
	}

	return res, nil
}

// PatchOrgsOrgIDAppsAppIDDeltasDeltaID - Update an existing Delta
func (s *delta) PatchOrgsOrgIDAppsAppIDDeltasDeltaID(ctx context.Context, request operations.PatchOrgsOrgIDAppsAppIDDeltasDeltaIDRequest) (*operations.PatchOrgsOrgIDAppsAppIDDeltasDeltaIDResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/orgs/{orgId}/apps/{appId}/deltas/{deltaId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bodyReader)
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

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PatchOrgsOrgIDAppsAppIDDeltasDeltaIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.DeltaResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.DeltaResponse = out
		}
	case httpRes.StatusCode == 400:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.HumanitecErrorResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.HumanitecErrorResponse = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			out := string(rawBody)
			res.PatchOrgsOrgIDAppsAppIDDeltasDeltaID404ApplicationJSONString = &out
		}
	}

	return res, nil
}

// PostOrgsOrgIDAppsAppIDDeltas - Create a new Delta
func (s *delta) PostOrgsOrgIDAppsAppIDDeltas(ctx context.Context, request operations.PostOrgsOrgIDAppsAppIDDeltasRequest) (*operations.PostOrgsOrgIDAppsAppIDDeltasResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/orgs/{orgId}/apps/{appId}/deltas", request, nil)
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
	req.Header.Set("Accept", "application/json;q=1, application/json;q=0")
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

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PostOrgsOrgIDAppsAppIDDeltasResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PostOrgsOrgIDAppsAppIDDeltas200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.PostOrgsOrgIDAppsAppIDDeltas200ApplicationJSONOneOf = out
		}
	case httpRes.StatusCode == 400:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.HumanitecErrorResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.HumanitecErrorResponse = out
		}
	}

	return res, nil
}

// PutDelta - Update an existing Delta
func (s *delta) PutDelta(ctx context.Context, request operations.PutDeltaRequest) (*operations.PutDeltaResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/orgs/{orgId}/apps/{appId}/deltas/{deltaId}", request, nil)
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

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/json;q=0")
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

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PutDeltaResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 204:
	case httpRes.StatusCode == 400:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.HumanitecErrorResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.HumanitecErrorResponse = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			out := string(rawBody)
			res.PutDelta404ApplicationJSONString = &out
		}
	}

	return res, nil
}

// PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataArchived - Mark a Delta as "archived"
// Archived Deltas are still accessible but can no longer be updated.
func (s *delta) PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataArchived(ctx context.Context, request operations.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataArchivedRequest) (*operations.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataArchivedResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/orgs/{orgId}/apps/{appId}/deltas/{deltaId}/metadata/archived", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/json;q=0")
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

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataArchivedResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 204:
	case httpRes.StatusCode == 400:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.HumanitecErrorResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.HumanitecErrorResponse = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			out := string(rawBody)
			res.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataArchived404ApplicationJSONString = &out
		}
	}

	return res, nil
}

// PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataEnvID - Change the Environment of a Delta
func (s *delta) PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataEnvID(ctx context.Context, request operations.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataEnvIDRequest) (*operations.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataEnvIDResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/orgs/{orgId}/apps/{appId}/deltas/{deltaId}/metadata/env_id", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "string")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/json;q=0")
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

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataEnvIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 204:
	case httpRes.StatusCode == 400:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.HumanitecErrorResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.HumanitecErrorResponse = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			out := string(rawBody)
			res.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataEnvID404ApplicationJSONString = &out
		}
	}

	return res, nil
}

// PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataName - Change the name of a Delta
func (s *delta) PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataName(ctx context.Context, request operations.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataNameRequest) (*operations.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataNameResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/orgs/{orgId}/apps/{appId}/deltas/{deltaId}/metadata/name", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "string")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/json;q=0")
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

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataNameResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 204:
	case httpRes.StatusCode == 400:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.HumanitecErrorResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.HumanitecErrorResponse = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			out := string(rawBody)
			res.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataName404ApplicationJSONString = &out
		}
	}

	return res, nil
}
