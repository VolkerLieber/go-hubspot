/*
CMS Source Code

Endpoints for interacting with files in the CMS Developer File System. These files include HTML templates, CSS, JS, modules, and other assets which are used to create CMS content.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package source_code

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"

	"github.com/clarkmcc/go-hubspot/authorization"
	_neturl "net/url"
	"os"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ContentApiService ContentApi service
type ContentApiService service

type ApiDeleteCmsV3SourceCodeEnvironmentContentPathRequest struct {
	ctx         _context.Context
	ApiService  *ContentApiService
	environment string
	path        string
}

func (r ApiDeleteCmsV3SourceCodeEnvironmentContentPathRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteCmsV3SourceCodeEnvironmentContentPathExecute(r)
}

/*
DeleteCmsV3SourceCodeEnvironmentContentPath Delete a file

Deletes the file at the specified path in the specified environment.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environment The environment of the file (\"draft\" or \"published\").
 @param path The file system location of the file.
 @return ApiDeleteCmsV3SourceCodeEnvironmentContentPathRequest
*/
func (a *ContentApiService) DeleteCmsV3SourceCodeEnvironmentContentPath(ctx _context.Context, environment string, path string) ApiDeleteCmsV3SourceCodeEnvironmentContentPathRequest {
	return ApiDeleteCmsV3SourceCodeEnvironmentContentPathRequest{
		ApiService:  a,
		ctx:         ctx,
		environment: environment,
		path:        path,
	}
}

// Execute executes the request
func (a *ContentApiService) DeleteCmsV3SourceCodeEnvironmentContentPathExecute(r ApiDeleteCmsV3SourceCodeEnvironmentContentPathRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentApiService.DeleteCmsV3SourceCodeEnvironmentContentPath")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/source-code/{environment}/content/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"environment"+"}", _neturl.PathEscape(parameterToString(r.environment, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", _neturl.PathEscape(parameterToString(r.path, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type ApiGetCmsV3SourceCodeEnvironmentContentPathRequest struct {
	ctx         _context.Context
	ApiService  *ContentApiService
	environment string
	path        string
}

func (r ApiGetCmsV3SourceCodeEnvironmentContentPathRequest) Execute() (Error, *_nethttp.Response, error) {
	return r.ApiService.GetCmsV3SourceCodeEnvironmentContentPathExecute(r)
}

/*
GetCmsV3SourceCodeEnvironmentContentPath Download a file

Downloads the byte contents of the file at the specified path in the specified environment.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environment The environment of the file (\"draft\" or \"published\").
 @param path The file system location of the file.
 @return ApiGetCmsV3SourceCodeEnvironmentContentPathRequest
*/
func (a *ContentApiService) GetCmsV3SourceCodeEnvironmentContentPath(ctx _context.Context, environment string, path string) ApiGetCmsV3SourceCodeEnvironmentContentPathRequest {
	return ApiGetCmsV3SourceCodeEnvironmentContentPathRequest{
		ApiService:  a,
		ctx:         ctx,
		environment: environment,
		path:        path,
	}
}

// Execute executes the request
//  @return Error
func (a *ContentApiService) GetCmsV3SourceCodeEnvironmentContentPathExecute(r ApiGetCmsV3SourceCodeEnvironmentContentPathRequest) (Error, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Error
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentApiService.GetCmsV3SourceCodeEnvironmentContentPath")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/source-code/{environment}/content/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"environment"+"}", _neturl.PathEscape(parameterToString(r.environment, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", _neturl.PathEscape(parameterToString(r.path, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type ApiPostCmsV3SourceCodeEnvironmentContentPathRequest struct {
	ctx         _context.Context
	ApiService  *ContentApiService
	environment string
	path        string
	file        **os.File
}

// The file to upload.
func (r ApiPostCmsV3SourceCodeEnvironmentContentPathRequest) File(file *os.File) ApiPostCmsV3SourceCodeEnvironmentContentPathRequest {
	r.file = &file
	return r
}

func (r ApiPostCmsV3SourceCodeEnvironmentContentPathRequest) Execute() (AssetFileMetadata, *_nethttp.Response, error) {
	return r.ApiService.PostCmsV3SourceCodeEnvironmentContentPathExecute(r)
}

/*
PostCmsV3SourceCodeEnvironmentContentPath Create a file

Creates a file at the specified path in the specified environment. Accepts multipart/form-data content type. Throws an error if a file already exists at the specified path.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environment The environment of the file (\"draft\" or \"published\").
 @param path The file system location of the file.
 @return ApiPostCmsV3SourceCodeEnvironmentContentPathRequest

Deprecated
*/
func (a *ContentApiService) PostCmsV3SourceCodeEnvironmentContentPath(ctx _context.Context, environment string, path string) ApiPostCmsV3SourceCodeEnvironmentContentPathRequest {
	return ApiPostCmsV3SourceCodeEnvironmentContentPathRequest{
		ApiService:  a,
		ctx:         ctx,
		environment: environment,
		path:        path,
	}
}

// Execute executes the request
//  @return AssetFileMetadata
// Deprecated
func (a *ContentApiService) PostCmsV3SourceCodeEnvironmentContentPathExecute(r ApiPostCmsV3SourceCodeEnvironmentContentPathRequest) (AssetFileMetadata, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AssetFileMetadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentApiService.PostCmsV3SourceCodeEnvironmentContentPath")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/source-code/{environment}/content/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"environment"+"}", _neturl.PathEscape(parameterToString(r.environment, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", _neturl.PathEscape(parameterToString(r.path, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	localVarFormFileName = "file"
	var localVarFile *os.File
	if r.file != nil {
		localVarFile = *r.file
	}
	if localVarFile != nil {
		fbs, _ := _ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
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

type ApiPutCmsV3SourceCodeEnvironmentContentPathRequest struct {
	ctx         _context.Context
	ApiService  *ContentApiService
	environment string
	path        string
	file        **os.File
}

// The file to upload.
func (r ApiPutCmsV3SourceCodeEnvironmentContentPathRequest) File(file *os.File) ApiPutCmsV3SourceCodeEnvironmentContentPathRequest {
	r.file = &file
	return r
}

func (r ApiPutCmsV3SourceCodeEnvironmentContentPathRequest) Execute() (AssetFileMetadata, *_nethttp.Response, error) {
	return r.ApiService.PutCmsV3SourceCodeEnvironmentContentPathExecute(r)
}

/*
PutCmsV3SourceCodeEnvironmentContentPath Create or update a file

Upserts a file at the specified path in the specified environment. Accepts multipart/form-data content type.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environment The environment of the file (\"draft\" or \"published\").
 @param path The file system location of the file.
 @return ApiPutCmsV3SourceCodeEnvironmentContentPathRequest
*/
func (a *ContentApiService) PutCmsV3SourceCodeEnvironmentContentPath(ctx _context.Context, environment string, path string) ApiPutCmsV3SourceCodeEnvironmentContentPathRequest {
	return ApiPutCmsV3SourceCodeEnvironmentContentPathRequest{
		ApiService:  a,
		ctx:         ctx,
		environment: environment,
		path:        path,
	}
}

// Execute executes the request
//  @return AssetFileMetadata
func (a *ContentApiService) PutCmsV3SourceCodeEnvironmentContentPathExecute(r ApiPutCmsV3SourceCodeEnvironmentContentPathRequest) (AssetFileMetadata, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AssetFileMetadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentApiService.PutCmsV3SourceCodeEnvironmentContentPath")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/source-code/{environment}/content/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"environment"+"}", _neturl.PathEscape(parameterToString(r.environment, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", _neturl.PathEscape(parameterToString(r.path, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	localVarFormFileName = "file"
	var localVarFile *os.File
	if r.file != nil {
		localVarFile = *r.file
	}
	if localVarFile != nil {
		fbs, _ := _ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
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
