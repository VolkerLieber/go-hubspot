/*
Contacts

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package contacts

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/clarkmcc/go-hubspot"
	"net/url"
)

// SearchApiService SearchApi service
type SearchApiService service

type ApiPostCrmV3ObjectsContactsSearchDoSearchRequest struct {
	ctx                       context.Context
	ApiService                *SearchApiService
	publicObjectSearchRequest *PublicObjectSearchRequest
}

func (r ApiPostCrmV3ObjectsContactsSearchDoSearchRequest) PublicObjectSearchRequest(publicObjectSearchRequest PublicObjectSearchRequest) ApiPostCrmV3ObjectsContactsSearchDoSearchRequest {
	r.publicObjectSearchRequest = &publicObjectSearchRequest
	return r
}

func (r ApiPostCrmV3ObjectsContactsSearchDoSearchRequest) Execute() (*CollectionResponseWithTotalSimplePublicObjectForwardPaging, *http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsContactsSearchDoSearchExecute(r)
}

/*
PostCrmV3ObjectsContactsSearchDoSearch Method for PostCrmV3ObjectsContactsSearchDoSearch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCrmV3ObjectsContactsSearchDoSearchRequest
*/
func (a *SearchApiService) PostCrmV3ObjectsContactsSearchDoSearch(ctx context.Context) ApiPostCrmV3ObjectsContactsSearchDoSearchRequest {
	return ApiPostCrmV3ObjectsContactsSearchDoSearchRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return CollectionResponseWithTotalSimplePublicObjectForwardPaging
func (a *SearchApiService) PostCrmV3ObjectsContactsSearchDoSearchExecute(r ApiPostCrmV3ObjectsContactsSearchDoSearchRequest) (*CollectionResponseWithTotalSimplePublicObjectForwardPaging, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponseWithTotalSimplePublicObjectForwardPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchApiService.PostCrmV3ObjectsContactsSearchDoSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/contacts/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.publicObjectSearchRequest == nil {
		return localVarReturnValue, nil, reportError("publicObjectSearchRequest is required and must be specified")
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
	localVarPostBody = r.publicObjectSearchRequest
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
