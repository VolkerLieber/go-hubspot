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

// FunctionsApiService FunctionsApi service
type FunctionsApiService service

type ApiFunctionsArchiveRequest struct {
	ctx          context.Context
	ApiService   *FunctionsApiService
	definitionId string
	functionType string
	functionId   string
	appId        int32
}

func (r ApiFunctionsArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.FunctionsArchiveExecute(r)
}

/*
FunctionsArchive Archive a function for a definition

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param definitionId
	@param functionType
	@param functionId
	@param appId
	@return ApiFunctionsArchiveRequest
*/
func (a *FunctionsApiService) FunctionsArchive(ctx context.Context, definitionId string, functionType string, functionId string, appId int32) ApiFunctionsArchiveRequest {
	return ApiFunctionsArchiveRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		functionType: functionType,
		functionId:   functionId,
		appId:        appId,
	}
}

// Execute executes the request
func (a *FunctionsApiService) FunctionsArchiveExecute(r ApiFunctionsArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FunctionsApiService.FunctionsArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterToString(r.definitionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functionType"+"}", url.PathEscape(parameterToString(r.functionType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functionId"+"}", url.PathEscape(parameterToString(r.functionId, "")), -1)
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
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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

type ApiFunctionsArchiveByTypeRequest struct {
	ctx          context.Context
	ApiService   *FunctionsApiService
	definitionId string
	functionType string
	appId        int32
}

func (r ApiFunctionsArchiveByTypeRequest) Execute() (*http.Response, error) {
	return r.ApiService.FunctionsArchiveByTypeExecute(r)
}

/*
FunctionsArchiveByType Delete a function for a definition

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param definitionId
	@param functionType
	@param appId
	@return ApiFunctionsArchiveByTypeRequest
*/
func (a *FunctionsApiService) FunctionsArchiveByType(ctx context.Context, definitionId string, functionType string, appId int32) ApiFunctionsArchiveByTypeRequest {
	return ApiFunctionsArchiveByTypeRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		functionType: functionType,
		appId:        appId,
	}
}

// Execute executes the request
func (a *FunctionsApiService) FunctionsArchiveByTypeExecute(r ApiFunctionsArchiveByTypeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FunctionsApiService.FunctionsArchiveByType")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}/functions/{functionType}"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterToString(r.definitionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functionType"+"}", url.PathEscape(parameterToString(r.functionType, "")), -1)
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
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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

type ApiFunctionsCreateOrReplaceRequest struct {
	ctx          context.Context
	ApiService   *FunctionsApiService
	definitionId string
	functionType string
	functionId   string
	appId        int32
	body         *string
}

func (r ApiFunctionsCreateOrReplaceRequest) Body(body string) ApiFunctionsCreateOrReplaceRequest {
	r.body = &body
	return r
}

func (r ApiFunctionsCreateOrReplaceRequest) Execute() (*PublicActionFunctionIdentifier, *http.Response, error) {
	return r.ApiService.FunctionsCreateOrReplaceExecute(r)
}

/*
FunctionsCreateOrReplace Insert a function for a definition

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param definitionId
	@param functionType
	@param functionId
	@param appId
	@return ApiFunctionsCreateOrReplaceRequest
*/
func (a *FunctionsApiService) FunctionsCreateOrReplace(ctx context.Context, definitionId string, functionType string, functionId string, appId int32) ApiFunctionsCreateOrReplaceRequest {
	return ApiFunctionsCreateOrReplaceRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		functionType: functionType,
		functionId:   functionId,
		appId:        appId,
	}
}

// Execute executes the request
//
//	@return PublicActionFunctionIdentifier
func (a *FunctionsApiService) FunctionsCreateOrReplaceExecute(r ApiFunctionsCreateOrReplaceRequest) (*PublicActionFunctionIdentifier, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PublicActionFunctionIdentifier
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FunctionsApiService.FunctionsCreateOrReplace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterToString(r.definitionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functionType"+"}", url.PathEscape(parameterToString(r.functionType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functionId"+"}", url.PathEscape(parameterToString(r.functionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

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
	localVarPostBody = r.body
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

type ApiFunctionsCreateOrReplaceByTypeRequest struct {
	ctx          context.Context
	ApiService   *FunctionsApiService
	definitionId string
	functionType string
	appId        int32
	body         *string
}

func (r ApiFunctionsCreateOrReplaceByTypeRequest) Body(body string) ApiFunctionsCreateOrReplaceByTypeRequest {
	r.body = &body
	return r
}

func (r ApiFunctionsCreateOrReplaceByTypeRequest) Execute() (*PublicActionFunctionIdentifier, *http.Response, error) {
	return r.ApiService.FunctionsCreateOrReplaceByTypeExecute(r)
}

/*
FunctionsCreateOrReplaceByType Insert a function for a definition

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param definitionId
	@param functionType
	@param appId
	@return ApiFunctionsCreateOrReplaceByTypeRequest
*/
func (a *FunctionsApiService) FunctionsCreateOrReplaceByType(ctx context.Context, definitionId string, functionType string, appId int32) ApiFunctionsCreateOrReplaceByTypeRequest {
	return ApiFunctionsCreateOrReplaceByTypeRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		functionType: functionType,
		appId:        appId,
	}
}

// Execute executes the request
//
//	@return PublicActionFunctionIdentifier
func (a *FunctionsApiService) FunctionsCreateOrReplaceByTypeExecute(r ApiFunctionsCreateOrReplaceByTypeRequest) (*PublicActionFunctionIdentifier, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PublicActionFunctionIdentifier
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FunctionsApiService.FunctionsCreateOrReplaceByType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}/functions/{functionType}"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterToString(r.definitionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functionType"+"}", url.PathEscape(parameterToString(r.functionType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

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
	localVarPostBody = r.body
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

type ApiFunctionsGetByIDRequest struct {
	ctx          context.Context
	ApiService   *FunctionsApiService
	definitionId string
	functionType string
	functionId   string
	appId        int32
}

func (r ApiFunctionsGetByIDRequest) Execute() (*PublicActionFunction, *http.Response, error) {
	return r.ApiService.FunctionsGetByIDExecute(r)
}

/*
FunctionsGetByID Get a function for a given definition

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param definitionId
	@param functionType
	@param functionId
	@param appId
	@return ApiFunctionsGetByIDRequest
*/
func (a *FunctionsApiService) FunctionsGetByID(ctx context.Context, definitionId string, functionType string, functionId string, appId int32) ApiFunctionsGetByIDRequest {
	return ApiFunctionsGetByIDRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		functionType: functionType,
		functionId:   functionId,
		appId:        appId,
	}
}

// Execute executes the request
//
//	@return PublicActionFunction
func (a *FunctionsApiService) FunctionsGetByIDExecute(r ApiFunctionsGetByIDRequest) (*PublicActionFunction, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PublicActionFunction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FunctionsApiService.FunctionsGetByID")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterToString(r.definitionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functionType"+"}", url.PathEscape(parameterToString(r.functionType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functionId"+"}", url.PathEscape(parameterToString(r.functionId, "")), -1)
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

type ApiFunctionsGetByTypeRequest struct {
	ctx          context.Context
	ApiService   *FunctionsApiService
	definitionId string
	functionType string
	appId        int32
}

func (r ApiFunctionsGetByTypeRequest) Execute() (*PublicActionFunction, *http.Response, error) {
	return r.ApiService.FunctionsGetByTypeExecute(r)
}

/*
FunctionsGetByType Get all functions by a type for a given definition

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param definitionId
	@param functionType
	@param appId
	@return ApiFunctionsGetByTypeRequest
*/
func (a *FunctionsApiService) FunctionsGetByType(ctx context.Context, definitionId string, functionType string, appId int32) ApiFunctionsGetByTypeRequest {
	return ApiFunctionsGetByTypeRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		functionType: functionType,
		appId:        appId,
	}
}

// Execute executes the request
//
//	@return PublicActionFunction
func (a *FunctionsApiService) FunctionsGetByTypeExecute(r ApiFunctionsGetByTypeRequest) (*PublicActionFunction, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PublicActionFunction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FunctionsApiService.FunctionsGetByType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}/functions/{functionType}"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterToString(r.definitionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functionType"+"}", url.PathEscape(parameterToString(r.functionType, "")), -1)
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

type ApiFunctionsGetPageRequest struct {
	ctx          context.Context
	ApiService   *FunctionsApiService
	definitionId string
	appId        int32
}

func (r ApiFunctionsGetPageRequest) Execute() (*CollectionResponsePublicActionFunctionIdentifierNoPaging, *http.Response, error) {
	return r.ApiService.FunctionsGetPageExecute(r)
}

/*
FunctionsGetPage Get all functions for a given definition

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param definitionId
	@param appId
	@return ApiFunctionsGetPageRequest
*/
func (a *FunctionsApiService) FunctionsGetPage(ctx context.Context, definitionId string, appId int32) ApiFunctionsGetPageRequest {
	return ApiFunctionsGetPageRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		appId:        appId,
	}
}

// Execute executes the request
//
//	@return CollectionResponsePublicActionFunctionIdentifierNoPaging
func (a *FunctionsApiService) FunctionsGetPageExecute(r ApiFunctionsGetPageRequest) (*CollectionResponsePublicActionFunctionIdentifierNoPaging, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponsePublicActionFunctionIdentifierNoPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FunctionsApiService.FunctionsGetPage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}/functions"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterToString(r.definitionId, "")), -1)
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
