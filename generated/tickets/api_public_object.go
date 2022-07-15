/*
Tickets

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tickets

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// PublicObjectApiService PublicObjectApi service
type PublicObjectApiService service

type ApiPostCrmV3ObjectsTicketsMergeMergeRequest struct {
	ctx              context.Context
	ApiService       *PublicObjectApiService
	publicMergeInput *PublicMergeInput
}

func (r ApiPostCrmV3ObjectsTicketsMergeMergeRequest) PublicMergeInput(publicMergeInput PublicMergeInput) ApiPostCrmV3ObjectsTicketsMergeMergeRequest {
	r.publicMergeInput = &publicMergeInput
	return r
}

func (r ApiPostCrmV3ObjectsTicketsMergeMergeRequest) Execute() (*SimplePublicObject, *http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsTicketsMergeMergeExecute(r)
}

/*
PostCrmV3ObjectsTicketsMergeMerge Merge two tickets with same type

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCrmV3ObjectsTicketsMergeMergeRequest
*/
func (a *PublicObjectApiService) PostCrmV3ObjectsTicketsMergeMerge(ctx context.Context) ApiPostCrmV3ObjectsTicketsMergeMergeRequest {
	return ApiPostCrmV3ObjectsTicketsMergeMergeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return SimplePublicObject
func (a *PublicObjectApiService) PostCrmV3ObjectsTicketsMergeMergeExecute(r ApiPostCrmV3ObjectsTicketsMergeMergeRequest) (*SimplePublicObject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicObjectApiService.PostCrmV3ObjectsTicketsMergeMerge")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/tickets/merge"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.publicMergeInput == nil {
		return localVarReturnValue, nil, reportError("publicMergeInput is required and must be specified")
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
	localVarPostBody = r.publicMergeInput
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
