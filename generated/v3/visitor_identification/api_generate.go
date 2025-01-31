/*
Visitor Identification

The Visitor Identification API allows you to pass identification information to the HubSpot chat widget for otherwise unknown visitors that were verified by your own authentication system.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package visitor_identification

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/clarkmcc/go-hubspot"
	"net/url"
)

// GenerateApiService GenerateApi service
type GenerateApiService service

type ApiPostConversationsV3VisitorIdentificationTokensCreateGenerateTokenRequest struct {
	ctx                                  context.Context
	ApiService                           *GenerateApiService
	identificationTokenGenerationRequest *IdentificationTokenGenerationRequest
}

func (r ApiPostConversationsV3VisitorIdentificationTokensCreateGenerateTokenRequest) IdentificationTokenGenerationRequest(identificationTokenGenerationRequest IdentificationTokenGenerationRequest) ApiPostConversationsV3VisitorIdentificationTokensCreateGenerateTokenRequest {
	r.identificationTokenGenerationRequest = &identificationTokenGenerationRequest
	return r
}

func (r ApiPostConversationsV3VisitorIdentificationTokensCreateGenerateTokenRequest) Execute() (*IdentificationTokenResponse, *http.Response, error) {
	return r.ApiService.PostConversationsV3VisitorIdentificationTokensCreateGenerateTokenExecute(r)
}

/*
PostConversationsV3VisitorIdentificationTokensCreateGenerateToken Generate a token

Generates a new visitor identification token. This token will be unique every time this endpoint is called, even if called with the same email address. This token is temporary and will expire after 12 hours

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostConversationsV3VisitorIdentificationTokensCreateGenerateTokenRequest
*/
func (a *GenerateApiService) PostConversationsV3VisitorIdentificationTokensCreateGenerateToken(ctx context.Context) ApiPostConversationsV3VisitorIdentificationTokensCreateGenerateTokenRequest {
	return ApiPostConversationsV3VisitorIdentificationTokensCreateGenerateTokenRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IdentificationTokenResponse
func (a *GenerateApiService) PostConversationsV3VisitorIdentificationTokensCreateGenerateTokenExecute(r ApiPostConversationsV3VisitorIdentificationTokensCreateGenerateTokenRequest) (*IdentificationTokenResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IdentificationTokenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GenerateApiService.PostConversationsV3VisitorIdentificationTokensCreateGenerateToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversations/v3/visitor-identification/tokens/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.identificationTokenGenerationRequest == nil {
		return localVarReturnValue, nil, reportError("identificationTokenGenerationRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.identificationTokenGenerationRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
