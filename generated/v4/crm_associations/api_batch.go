/*
CrmPublicAssociationsServiceV4

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crm_associations

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/clarkmcc/go-hubspot"
	"net/url"
	"strings"
)

// BatchApiService BatchApi service
type BatchApiService service

type ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchiveRequest struct {
	ctx                                     context.Context
	ApiService                              *BatchApiService
	fromObjectType                          string
	toObjectType                            string
	batchInputPublicAssociationMultiArchive *BatchInputPublicAssociationMultiArchive
}

func (r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchiveRequest) BatchInputPublicAssociationMultiArchive(batchInputPublicAssociationMultiArchive BatchInputPublicAssociationMultiArchive) ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchiveRequest {
	r.batchInputPublicAssociationMultiArchive = &batchInputPublicAssociationMultiArchive
	return r
}

func (r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchiveExecute(r)
}

/*
PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchive Delete

Batch delete associations for objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fromObjectType
	@param toObjectType
	@return ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchiveRequest
*/
func (a *BatchApiService) PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchive(ctx context.Context, fromObjectType string, toObjectType string) ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchiveRequest {
	return ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchiveRequest{
		ApiService:     a,
		ctx:            ctx,
		fromObjectType: fromObjectType,
		toObjectType:   toObjectType,
	}
}

// Execute executes the request
func (a *BatchApiService) PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchiveExecute(r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v4/associations/{fromObjectType}/{toObjectType}/batch/archive"
	localVarPath = strings.Replace(localVarPath, "{"+"fromObjectType"+"}", url.PathEscape(parameterToString(r.fromObjectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterToString(r.toObjectType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputPublicAssociationMultiArchive == nil {
		return nil, reportError("batchInputPublicAssociationMultiArchive is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.batchInputPublicAssociationMultiArchive
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
			return localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefaultRequest struct {
	ctx                                         context.Context
	ApiService                                  *BatchApiService
	fromObjectType                              string
	toObjectType                                string
	batchInputPublicDefaultAssociationMultiPost *BatchInputPublicDefaultAssociationMultiPost
}

func (r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefaultRequest) BatchInputPublicDefaultAssociationMultiPost(batchInputPublicDefaultAssociationMultiPost BatchInputPublicDefaultAssociationMultiPost) ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefaultRequest {
	r.batchInputPublicDefaultAssociationMultiPost = &batchInputPublicDefaultAssociationMultiPost
	return r
}

func (r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefaultRequest) Execute() (*BatchResponsePublicDefaultAssociation, *http.Response, error) {
	return r.ApiService.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefaultExecute(r)
}

/*
PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefault  Create Default Associations

Create the default (most generic) association type between two object types

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fromObjectType
	@param toObjectType
	@return ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefaultRequest
*/
func (a *BatchApiService) PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefault(ctx context.Context, fromObjectType string, toObjectType string) ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefaultRequest {
	return ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefaultRequest{
		ApiService:     a,
		ctx:            ctx,
		fromObjectType: fromObjectType,
		toObjectType:   toObjectType,
	}
}

// Execute executes the request
//
//	@return BatchResponsePublicDefaultAssociation
func (a *BatchApiService) PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefaultExecute(r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefaultRequest) (*BatchResponsePublicDefaultAssociation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponsePublicDefaultAssociation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchAssociateDefaultCreateDefault")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v4/associations/{fromObjectType}/{toObjectType}/batch/associate/default"
	localVarPath = strings.Replace(localVarPath, "{"+"fromObjectType"+"}", url.PathEscape(parameterToString(r.fromObjectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterToString(r.toObjectType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputPublicDefaultAssociationMultiPost == nil {
		return localVarReturnValue, nil, reportError("batchInputPublicDefaultAssociationMultiPost is required and must be specified")
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
	localVarPostBody = r.batchInputPublicDefaultAssociationMultiPost
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

type ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreateRequest struct {
	ctx                                  context.Context
	ApiService                           *BatchApiService
	fromObjectType                       string
	toObjectType                         string
	batchInputPublicAssociationMultiPost *BatchInputPublicAssociationMultiPost
}

func (r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreateRequest) BatchInputPublicAssociationMultiPost(batchInputPublicAssociationMultiPost BatchInputPublicAssociationMultiPost) ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreateRequest {
	r.batchInputPublicAssociationMultiPost = &batchInputPublicAssociationMultiPost
	return r
}

func (r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreateRequest) Execute() (*BatchResponseLabelsBetweenObjectPair, *http.Response, error) {
	return r.ApiService.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreateExecute(r)
}

/*
PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreate Create

Batch create associations for objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fromObjectType
	@param toObjectType
	@return ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreateRequest
*/
func (a *BatchApiService) PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreate(ctx context.Context, fromObjectType string, toObjectType string) ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreateRequest {
	return ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreateRequest{
		ApiService:     a,
		ctx:            ctx,
		fromObjectType: fromObjectType,
		toObjectType:   toObjectType,
	}
}

// Execute executes the request
//
//	@return BatchResponseLabelsBetweenObjectPair
func (a *BatchApiService) PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreateExecute(r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreateRequest) (*BatchResponseLabelsBetweenObjectPair, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponseLabelsBetweenObjectPair
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v4/associations/{fromObjectType}/{toObjectType}/batch/create"
	localVarPath = strings.Replace(localVarPath, "{"+"fromObjectType"+"}", url.PathEscape(parameterToString(r.fromObjectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterToString(r.toObjectType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputPublicAssociationMultiPost == nil {
		return localVarReturnValue, nil, reportError("batchInputPublicAssociationMultiPost is required and must be specified")
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
	localVarPostBody = r.batchInputPublicAssociationMultiPost
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

type ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabelsRequest struct {
	ctx                                  context.Context
	ApiService                           *BatchApiService
	fromObjectType                       string
	toObjectType                         string
	batchInputPublicAssociationMultiPost *BatchInputPublicAssociationMultiPost
}

func (r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabelsRequest) BatchInputPublicAssociationMultiPost(batchInputPublicAssociationMultiPost BatchInputPublicAssociationMultiPost) ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabelsRequest {
	r.batchInputPublicAssociationMultiPost = &batchInputPublicAssociationMultiPost
	return r
}

func (r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabelsRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabelsExecute(r)
}

/*
PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabels Delete Specific Labels

Batch delete specific association labels for objects. Deleting an unlabeled association will also delete all labeled associations between those two objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fromObjectType
	@param toObjectType
	@return ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabelsRequest
*/
func (a *BatchApiService) PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabels(ctx context.Context, fromObjectType string, toObjectType string) ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabelsRequest {
	return ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabelsRequest{
		ApiService:     a,
		ctx:            ctx,
		fromObjectType: fromObjectType,
		toObjectType:   toObjectType,
	}
}

// Execute executes the request
func (a *BatchApiService) PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabelsExecute(r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabelsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveArchiveLabels")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v4/associations/{fromObjectType}/{toObjectType}/batch/labels/archive"
	localVarPath = strings.Replace(localVarPath, "{"+"fromObjectType"+"}", url.PathEscape(parameterToString(r.fromObjectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterToString(r.toObjectType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputPublicAssociationMultiPost == nil {
		return nil, reportError("batchInputPublicAssociationMultiPost is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.batchInputPublicAssociationMultiPost
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
			return localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPageRequest struct {
	ctx                                           context.Context
	ApiService                                    *BatchApiService
	fromObjectType                                string
	toObjectType                                  string
	batchInputPublicFetchAssociationsBatchRequest *BatchInputPublicFetchAssociationsBatchRequest
}

func (r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPageRequest) BatchInputPublicFetchAssociationsBatchRequest(batchInputPublicFetchAssociationsBatchRequest BatchInputPublicFetchAssociationsBatchRequest) ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPageRequest {
	r.batchInputPublicFetchAssociationsBatchRequest = &batchInputPublicFetchAssociationsBatchRequest
	return r
}

func (r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPageRequest) Execute() (*BatchResponsePublicAssociationMultiWithLabel, *http.Response, error) {
	return r.ApiService.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPageExecute(r)
}

/*
PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPage Read

Batch read associations for objects to specific object type. The 'after' field in a returned paging object  can be added alongside the 'id' to retrieve the next page of associations from that objectId. The 'link' field is deprecated and should be ignored.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fromObjectType
	@param toObjectType
	@return ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPageRequest
*/
func (a *BatchApiService) PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPage(ctx context.Context, fromObjectType string, toObjectType string) ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPageRequest {
	return ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPageRequest{
		ApiService:     a,
		ctx:            ctx,
		fromObjectType: fromObjectType,
		toObjectType:   toObjectType,
	}
}

// Execute executes the request
//
//	@return BatchResponsePublicAssociationMultiWithLabel
func (a *BatchApiService) PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPageExecute(r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPageRequest) (*BatchResponsePublicAssociationMultiWithLabel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponsePublicAssociationMultiWithLabel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadGetPage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v4/associations/{fromObjectType}/{toObjectType}/batch/read"
	localVarPath = strings.Replace(localVarPath, "{"+"fromObjectType"+"}", url.PathEscape(parameterToString(r.fromObjectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterToString(r.toObjectType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputPublicFetchAssociationsBatchRequest == nil {
		return localVarReturnValue, nil, reportError("batchInputPublicFetchAssociationsBatchRequest is required and must be specified")
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
	localVarPostBody = r.batchInputPublicFetchAssociationsBatchRequest
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
