// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteTokensTokenIDRequest struct {
	// The token ID
	//
	//
	TokenID string `pathParam:"style=simple,explode=false,name=tokenId"`
}

type DeleteTokensTokenIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
