/*
Automation Actions V4

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package actions

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// RevisionsApiService RevisionsApi service
type RevisionsApiService service

type ApiRevisionsGetByIDRequest struct {
	ctx          context.Context
	ApiService   *RevisionsApiService
	definitionId string
	revisionId   string
	appId        int32
}

func (r ApiRevisionsGetByIDRequest) Execute() (*PublicActionRevision, *http.Response, error) {
	return r.ApiService.RevisionsGetByIDExecute(r)
}

/*
RevisionsGetByID Gets a revision for a given definition by revision id

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param definitionId
	@param revisionId
	@param appId
	@return ApiRevisionsGetByIDRequest
*/
func (a *RevisionsApiService) RevisionsGetByID(ctx context.Context, definitionId string, revisionId string, appId int32) ApiRevisionsGetByIDRequest {
	return ApiRevisionsGetByIDRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		revisionId:   revisionId,
		appId:        appId,
	}
}

// Execute executes the request
//
//	@return PublicActionRevision
func (a *RevisionsApiService) RevisionsGetByIDExecute(r ApiRevisionsGetByIDRequest) (*PublicActionRevision, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PublicActionRevision
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RevisionsApiService.RevisionsGetByID")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}/revisions/{revisionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterToString(r.definitionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"revisionId"+"}", url.PathEscape(parameterToString(r.revisionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiRevisionsGetPageRequest struct {
	ctx          context.Context
	ApiService   *RevisionsApiService
	definitionId string
	appId        int32
	limit        *int32
	after        *string
}

// The maximum number of results to display per page.
func (r ApiRevisionsGetPageRequest) Limit(limit int32) ApiRevisionsGetPageRequest {
	r.limit = &limit
	return r
}

// The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results.
func (r ApiRevisionsGetPageRequest) After(after string) ApiRevisionsGetPageRequest {
	r.after = &after
	return r
}

func (r ApiRevisionsGetPageRequest) Execute() (*CollectionResponsePublicActionRevisionForwardPaging, *http.Response, error) {
	return r.ApiService.RevisionsGetPageExecute(r)
}

/*
RevisionsGetPage Get all revisions for a given definition

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param definitionId
	@param appId
	@return ApiRevisionsGetPageRequest
*/
func (a *RevisionsApiService) RevisionsGetPage(ctx context.Context, definitionId string, appId int32) ApiRevisionsGetPageRequest {
	return ApiRevisionsGetPageRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		appId:        appId,
	}
}

// Execute executes the request
//
//	@return CollectionResponsePublicActionRevisionForwardPaging
func (a *RevisionsApiService) RevisionsGetPageExecute(r ApiRevisionsGetPageRequest) (*CollectionResponsePublicActionRevisionForwardPaging, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponsePublicActionRevisionForwardPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RevisionsApiService.RevisionsGetPage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}/revisions"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterToString(r.definitionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
