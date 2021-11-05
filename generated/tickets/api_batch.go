/*
Tickets

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tickets

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"

	"github.com/clarkmcc/go-hubspot/authorization"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// BatchApiService BatchApi service
type BatchApiService service

type ApiPostCrmV3ObjectsTicketsBatchArchiveArchiveRequest struct {
	ctx                            _context.Context
	ApiService                     *BatchApiService
	batchInputSimplePublicObjectId *BatchInputSimplePublicObjectId
}

func (r ApiPostCrmV3ObjectsTicketsBatchArchiveArchiveRequest) BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId BatchInputSimplePublicObjectId) ApiPostCrmV3ObjectsTicketsBatchArchiveArchiveRequest {
	r.batchInputSimplePublicObjectId = &batchInputSimplePublicObjectId
	return r
}

func (r ApiPostCrmV3ObjectsTicketsBatchArchiveArchiveRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PostCrmV3ObjectsTicketsBatchArchiveArchiveExecute(r)
}

/*
PostCrmV3ObjectsTicketsBatchArchiveArchive Archive a batch of tickets by ID

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCrmV3ObjectsTicketsBatchArchiveArchiveRequest
*/
func (a *BatchApiService) PostCrmV3ObjectsTicketsBatchArchiveArchive(ctx _context.Context) ApiPostCrmV3ObjectsTicketsBatchArchiveArchiveRequest {
	return ApiPostCrmV3ObjectsTicketsBatchArchiveArchiveRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *BatchApiService) PostCrmV3ObjectsTicketsBatchArchiveArchiveExecute(r ApiPostCrmV3ObjectsTicketsBatchArchiveArchiveRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.PostCrmV3ObjectsTicketsBatchArchiveArchive")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/tickets/batch/archive"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.batchInputSimplePublicObjectId == nil {
		return nil, reportError("batchInputSimplePublicObjectId is required and must be specified")
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
	localVarPostBody = r.batchInputSimplePublicObjectId
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(authorization.ContextAPIKeys).(map[string]authorization.APIKey); ok {
			if apiKey, ok := auth["hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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

type ApiPostCrmV3ObjectsTicketsBatchCreateCreateRequest struct {
	ctx                               _context.Context
	ApiService                        *BatchApiService
	batchInputSimplePublicObjectInput *BatchInputSimplePublicObjectInput
}

func (r ApiPostCrmV3ObjectsTicketsBatchCreateCreateRequest) BatchInputSimplePublicObjectInput(batchInputSimplePublicObjectInput BatchInputSimplePublicObjectInput) ApiPostCrmV3ObjectsTicketsBatchCreateCreateRequest {
	r.batchInputSimplePublicObjectInput = &batchInputSimplePublicObjectInput
	return r
}

func (r ApiPostCrmV3ObjectsTicketsBatchCreateCreateRequest) Execute() (BatchResponseSimplePublicObject, *_nethttp.Response, error) {
	return r.ApiService.PostCrmV3ObjectsTicketsBatchCreateCreateExecute(r)
}

/*
PostCrmV3ObjectsTicketsBatchCreateCreate Create a batch of tickets

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCrmV3ObjectsTicketsBatchCreateCreateRequest
*/
func (a *BatchApiService) PostCrmV3ObjectsTicketsBatchCreateCreate(ctx _context.Context) ApiPostCrmV3ObjectsTicketsBatchCreateCreateRequest {
	return ApiPostCrmV3ObjectsTicketsBatchCreateCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BatchResponseSimplePublicObject
func (a *BatchApiService) PostCrmV3ObjectsTicketsBatchCreateCreateExecute(r ApiPostCrmV3ObjectsTicketsBatchCreateCreateRequest) (BatchResponseSimplePublicObject, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BatchResponseSimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.PostCrmV3ObjectsTicketsBatchCreateCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/tickets/batch/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.batchInputSimplePublicObjectInput == nil {
		return localVarReturnValue, nil, reportError("batchInputSimplePublicObjectInput is required and must be specified")
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
	localVarPostBody = r.batchInputSimplePublicObjectInput
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(authorization.ContextAPIKeys).(map[string]authorization.APIKey); ok {
			if apiKey, ok := auth["hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostCrmV3ObjectsTicketsBatchReadReadRequest struct {
	ctx                                _context.Context
	ApiService                         *BatchApiService
	batchReadInputSimplePublicObjectId *BatchReadInputSimplePublicObjectId
	archived                           *bool
}

func (r ApiPostCrmV3ObjectsTicketsBatchReadReadRequest) BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId BatchReadInputSimplePublicObjectId) ApiPostCrmV3ObjectsTicketsBatchReadReadRequest {
	r.batchReadInputSimplePublicObjectId = &batchReadInputSimplePublicObjectId
	return r
}

// Whether to return only results that have been archived.
func (r ApiPostCrmV3ObjectsTicketsBatchReadReadRequest) Archived(archived bool) ApiPostCrmV3ObjectsTicketsBatchReadReadRequest {
	r.archived = &archived
	return r
}

func (r ApiPostCrmV3ObjectsTicketsBatchReadReadRequest) Execute() (BatchResponseSimplePublicObject, *_nethttp.Response, error) {
	return r.ApiService.PostCrmV3ObjectsTicketsBatchReadReadExecute(r)
}

/*
PostCrmV3ObjectsTicketsBatchReadRead Read a batch of tickets by internal ID, or unique property values

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCrmV3ObjectsTicketsBatchReadReadRequest
*/
func (a *BatchApiService) PostCrmV3ObjectsTicketsBatchReadRead(ctx _context.Context) ApiPostCrmV3ObjectsTicketsBatchReadReadRequest {
	return ApiPostCrmV3ObjectsTicketsBatchReadReadRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BatchResponseSimplePublicObject
func (a *BatchApiService) PostCrmV3ObjectsTicketsBatchReadReadExecute(r ApiPostCrmV3ObjectsTicketsBatchReadReadRequest) (BatchResponseSimplePublicObject, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BatchResponseSimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.PostCrmV3ObjectsTicketsBatchReadRead")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/tickets/batch/read"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.batchReadInputSimplePublicObjectId == nil {
		return localVarReturnValue, nil, reportError("batchReadInputSimplePublicObjectId is required and must be specified")
	}

	if r.archived != nil {
		localVarQueryParams.Add("archived", parameterToString(*r.archived, ""))
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
	localVarPostBody = r.batchReadInputSimplePublicObjectId
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(authorization.ContextAPIKeys).(map[string]authorization.APIKey); ok {
			if apiKey, ok := auth["hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostCrmV3ObjectsTicketsBatchUpdateUpdateRequest struct {
	ctx                                    _context.Context
	ApiService                             *BatchApiService
	batchInputSimplePublicObjectBatchInput *BatchInputSimplePublicObjectBatchInput
}

func (r ApiPostCrmV3ObjectsTicketsBatchUpdateUpdateRequest) BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput BatchInputSimplePublicObjectBatchInput) ApiPostCrmV3ObjectsTicketsBatchUpdateUpdateRequest {
	r.batchInputSimplePublicObjectBatchInput = &batchInputSimplePublicObjectBatchInput
	return r
}

func (r ApiPostCrmV3ObjectsTicketsBatchUpdateUpdateRequest) Execute() (BatchResponseSimplePublicObject, *_nethttp.Response, error) {
	return r.ApiService.PostCrmV3ObjectsTicketsBatchUpdateUpdateExecute(r)
}

/*
PostCrmV3ObjectsTicketsBatchUpdateUpdate Update a batch of tickets

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCrmV3ObjectsTicketsBatchUpdateUpdateRequest
*/
func (a *BatchApiService) PostCrmV3ObjectsTicketsBatchUpdateUpdate(ctx _context.Context) ApiPostCrmV3ObjectsTicketsBatchUpdateUpdateRequest {
	return ApiPostCrmV3ObjectsTicketsBatchUpdateUpdateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BatchResponseSimplePublicObject
func (a *BatchApiService) PostCrmV3ObjectsTicketsBatchUpdateUpdateExecute(r ApiPostCrmV3ObjectsTicketsBatchUpdateUpdateRequest) (BatchResponseSimplePublicObject, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BatchResponseSimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.PostCrmV3ObjectsTicketsBatchUpdateUpdate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/tickets/batch/update"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.batchInputSimplePublicObjectBatchInput == nil {
		return localVarReturnValue, nil, reportError("batchInputSimplePublicObjectBatchInput is required and must be specified")
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
	localVarPostBody = r.batchInputSimplePublicObjectBatchInput
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(authorization.ContextAPIKeys).(map[string]authorization.APIKey); ok {
			if apiKey, ok := auth["hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
