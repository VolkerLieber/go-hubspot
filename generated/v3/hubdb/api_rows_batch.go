/*
Hubdb

HubDB is a relational data store that presents data as rows, columns, and cells in a table, much like a spreadsheet. HubDB tables can be added or modified [in the HubSpot CMS](https://knowledge.hubspot.com/cos-general/how-to-edit-hubdb-tables), but you can also use the API endpoints documented here. For more information on HubDB tables and using their data on a HubSpot site, see the [CMS developers site](https://designers.hubspot.com/docs/tools/hubdb). You can also see the [documentation for dynamic pages](https://designers.hubspot.com/docs/tutorials/how-to-build-dynamic-pages-with-hubdb) for more details about the `useForPages` field.  HubDB tables support `draft` and `published` versions. This allows you to update data in the table, either for testing or to allow for a manual approval process, without affecting any live pages using the existing data. Draft data can be reviewed, and published by a user working in HubSpot or published via the API. Draft data can also be discarded, allowing users to go back to the published version of the data without disrupting it. If a table is set to be `allowed for public access`, you can access the published version of the table and rows without any authentication by specifying the portal id via the query parameter `portalId`.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hubdb

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/clarkmcc/go-hubspot"
	"net/url"
	"strings"
)

// RowsBatchApiService RowsBatchApi service
type RowsBatchApiService service

type ApiPostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRowsRequest struct {
	ctx              context.Context
	ApiService       *RowsBatchApiService
	tableIdOrName    string
	batchInputString *BatchInputString
}

// The JSON array of row ids
func (r ApiPostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRowsRequest) BatchInputString(batchInputString BatchInputString) ApiPostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRowsRequest {
	r.batchInputString = &batchInputString
	return r
}

func (r ApiPostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRowsRequest) Execute() (*BatchResponseHubDbTableRowV3, *http.Response, error) {
	return r.ApiService.PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRowsExecute(r)
}

/*
PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRows Get a set of rows

Returns rows in the `published` version of the specified table, given a set of row ids.
**Note:** This endpoint can be accessed without any authentication if the table is set to be allowed for public access.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tableIdOrName The ID or name of the table to query.
	@return ApiPostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRowsRequest
*/
func (a *RowsBatchApiService) PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRows(ctx context.Context, tableIdOrName string) ApiPostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRowsRequest {
	return ApiPostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRowsRequest{
		ApiService:    a,
		ctx:           ctx,
		tableIdOrName: tableIdOrName,
	}
}

// Execute executes the request
//
//	@return BatchResponseHubDbTableRowV3
func (a *RowsBatchApiService) PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRowsExecute(r ApiPostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRowsRequest) (*BatchResponseHubDbTableRowV3, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponseHubDbTableRowV3
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RowsBatchApiService.PostCmsV3HubdbTablesTableIdOrNameRowsBatchReadReadTableRows")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/hubdb/tables/{tableIdOrName}/rows/batch/read"
	localVarPath = strings.Replace(localVarPath, "{"+"tableIdOrName"+"}", url.PathEscape(parameterToString(r.tableIdOrName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputString == nil {
		return localVarReturnValue, nil, reportError("batchInputString is required and must be specified")
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
	localVarPostBody = r.batchInputString
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

type ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRowsRequest struct {
	ctx              context.Context
	ApiService       *RowsBatchApiService
	tableIdOrName    string
	batchInputString *BatchInputString
}

// The JSON array of row ids
func (r ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRowsRequest) BatchInputString(batchInputString BatchInputString) ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRowsRequest {
	r.batchInputString = &batchInputString
	return r
}

func (r ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRowsRequest) Execute() (*BatchResponseHubDbTableRowV3, *http.Response, error) {
	return r.ApiService.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRowsExecute(r)
}

/*
PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRows Clone rows in batch

Clones rows in the `draft` version of the specified table, given a set of row ids. Maximum of 100 row ids per call.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tableIdOrName The ID or name of the table
	@return ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRowsRequest
*/
func (a *RowsBatchApiService) PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRows(ctx context.Context, tableIdOrName string) ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRowsRequest {
	return ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRowsRequest{
		ApiService:    a,
		ctx:           ctx,
		tableIdOrName: tableIdOrName,
	}
}

// Execute executes the request
//
//	@return BatchResponseHubDbTableRowV3
func (a *RowsBatchApiService) PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRowsExecute(r ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRowsRequest) (*BatchResponseHubDbTableRowV3, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponseHubDbTableRowV3
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RowsBatchApiService.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneCloneDraftTableRows")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/clone"
	localVarPath = strings.Replace(localVarPath, "{"+"tableIdOrName"+"}", url.PathEscape(parameterToString(r.tableIdOrName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputString == nil {
		return localVarReturnValue, nil, reportError("batchInputString is required and must be specified")
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
	localVarPostBody = r.batchInputString
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

type ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRowsRequest struct {
	ctx                              context.Context
	ApiService                       *RowsBatchApiService
	tableIdOrName                    string
	batchInputHubDbTableRowV3Request *BatchInputHubDbTableRowV3Request
}

// JSON array of row objects
func (r ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRowsRequest) BatchInputHubDbTableRowV3Request(batchInputHubDbTableRowV3Request BatchInputHubDbTableRowV3Request) ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRowsRequest {
	r.batchInputHubDbTableRowV3Request = &batchInputHubDbTableRowV3Request
	return r
}

func (r ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRowsRequest) Execute() (*BatchResponseHubDbTableRowV3, *http.Response, error) {
	return r.ApiService.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRowsExecute(r)
}

/*
PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRows Create rows in batch

Creates rows in the `draft` version of the specified table, given an array of row objects. Maximum of 100 row object per call. See the overview section for more details with an example.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tableIdOrName The ID or name of the table
	@return ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRowsRequest
*/
func (a *RowsBatchApiService) PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRows(ctx context.Context, tableIdOrName string) ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRowsRequest {
	return ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRowsRequest{
		ApiService:    a,
		ctx:           ctx,
		tableIdOrName: tableIdOrName,
	}
}

// Execute executes the request
//
//	@return BatchResponseHubDbTableRowV3
func (a *RowsBatchApiService) PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRowsExecute(r ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRowsRequest) (*BatchResponseHubDbTableRowV3, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponseHubDbTableRowV3
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RowsBatchApiService.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateCreateDraftTableRows")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/create"
	localVarPath = strings.Replace(localVarPath, "{"+"tableIdOrName"+"}", url.PathEscape(parameterToString(r.tableIdOrName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputHubDbTableRowV3Request == nil {
		return localVarReturnValue, nil, reportError("batchInputHubDbTableRowV3Request is required and must be specified")
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
	localVarPostBody = r.batchInputHubDbTableRowV3Request
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

type ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgePurgeDraftTableRowsRequest struct {
	ctx              context.Context
	ApiService       *RowsBatchApiService
	tableIdOrName    string
	batchInputString *BatchInputString
}

// JSON array of row ids.
func (r ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgePurgeDraftTableRowsRequest) BatchInputString(batchInputString BatchInputString) ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgePurgeDraftTableRowsRequest {
	r.batchInputString = &batchInputString
	return r
}

func (r ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgePurgeDraftTableRowsRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgePurgeDraftTableRowsExecute(r)
}

/*
PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgePurgeDraftTableRows Permanently deletes rows

Permanently deletes rows from the `draft` version of the table, given a set of row ids. Maximum of 100 row ids per call.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tableIdOrName The ID or name of the table
	@return ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgePurgeDraftTableRowsRequest
*/
func (a *RowsBatchApiService) PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgePurgeDraftTableRows(ctx context.Context, tableIdOrName string) ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgePurgeDraftTableRowsRequest {
	return ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgePurgeDraftTableRowsRequest{
		ApiService:    a,
		ctx:           ctx,
		tableIdOrName: tableIdOrName,
	}
}

// Execute executes the request
func (a *RowsBatchApiService) PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgePurgeDraftTableRowsExecute(r ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgePurgeDraftTableRowsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RowsBatchApiService.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgePurgeDraftTableRows")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/purge"
	localVarPath = strings.Replace(localVarPath, "{"+"tableIdOrName"+"}", url.PathEscape(parameterToString(r.tableIdOrName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputString == nil {
		return nil, reportError("batchInputString is required and must be specified")
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
	localVarPostBody = r.batchInputString
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

type ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRowsRequest struct {
	ctx              context.Context
	ApiService       *RowsBatchApiService
	tableIdOrName    string
	batchInputString *BatchInputString
}

// JSON array of row ids.
func (r ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRowsRequest) BatchInputString(batchInputString BatchInputString) ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRowsRequest {
	r.batchInputString = &batchInputString
	return r
}

func (r ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRowsRequest) Execute() (*BatchResponseHubDbTableRowV3, *http.Response, error) {
	return r.ApiService.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRowsExecute(r)
}

/*
PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRows Get a set of rows from draft table

Returns rows in the `draft` version of the specified table, given a set of row ids.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tableIdOrName The ID or name of the table
	@return ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRowsRequest
*/
func (a *RowsBatchApiService) PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRows(ctx context.Context, tableIdOrName string) ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRowsRequest {
	return ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRowsRequest{
		ApiService:    a,
		ctx:           ctx,
		tableIdOrName: tableIdOrName,
	}
}

// Execute executes the request
//
//	@return BatchResponseHubDbTableRowV3
func (a *RowsBatchApiService) PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRowsExecute(r ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRowsRequest) (*BatchResponseHubDbTableRowV3, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponseHubDbTableRowV3
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RowsBatchApiService.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadReadDraftTableRows")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/read"
	localVarPath = strings.Replace(localVarPath, "{"+"tableIdOrName"+"}", url.PathEscape(parameterToString(r.tableIdOrName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputString == nil {
		return localVarReturnValue, nil, reportError("batchInputString is required and must be specified")
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
	localVarPostBody = r.batchInputString
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

type ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRowsRequest struct {
	ctx                                         context.Context
	ApiService                                  *RowsBatchApiService
	tableIdOrName                               string
	batchInputHubDbTableRowV3BatchUpdateRequest *BatchInputHubDbTableRowV3BatchUpdateRequest
}

// JSON array of row objects.
func (r ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRowsRequest) BatchInputHubDbTableRowV3BatchUpdateRequest(batchInputHubDbTableRowV3BatchUpdateRequest BatchInputHubDbTableRowV3BatchUpdateRequest) ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRowsRequest {
	r.batchInputHubDbTableRowV3BatchUpdateRequest = &batchInputHubDbTableRowV3BatchUpdateRequest
	return r
}

func (r ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRowsRequest) Execute() (*BatchResponseHubDbTableRowV3, *http.Response, error) {
	return r.ApiService.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRowsExecute(r)
}

/*
PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRows Replace rows in batch in draft table

Replaces multiple rows as a batch in the `draft` version of the table, with a maximum of 100 rows per call. See the endpoint `PUT /tables/{tableIdOrName}/rows/{rowId}/draft` for details on updating a single row.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tableIdOrName The ID or name of the table
	@return ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRowsRequest
*/
func (a *RowsBatchApiService) PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRows(ctx context.Context, tableIdOrName string) ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRowsRequest {
	return ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRowsRequest{
		ApiService:    a,
		ctx:           ctx,
		tableIdOrName: tableIdOrName,
	}
}

// Execute executes the request
//
//	@return BatchResponseHubDbTableRowV3
func (a *RowsBatchApiService) PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRowsExecute(r ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRowsRequest) (*BatchResponseHubDbTableRowV3, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponseHubDbTableRowV3
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RowsBatchApiService.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceReplaceDraftTableRows")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/replace"
	localVarPath = strings.Replace(localVarPath, "{"+"tableIdOrName"+"}", url.PathEscape(parameterToString(r.tableIdOrName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputHubDbTableRowV3BatchUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("batchInputHubDbTableRowV3BatchUpdateRequest is required and must be specified")
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
	localVarPostBody = r.batchInputHubDbTableRowV3BatchUpdateRequest
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

type ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRowsRequest struct {
	ctx                                         context.Context
	ApiService                                  *RowsBatchApiService
	tableIdOrName                               string
	batchInputHubDbTableRowV3BatchUpdateRequest *BatchInputHubDbTableRowV3BatchUpdateRequest
}

// JSON array of row objects.
func (r ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRowsRequest) BatchInputHubDbTableRowV3BatchUpdateRequest(batchInputHubDbTableRowV3BatchUpdateRequest BatchInputHubDbTableRowV3BatchUpdateRequest) ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRowsRequest {
	r.batchInputHubDbTableRowV3BatchUpdateRequest = &batchInputHubDbTableRowV3BatchUpdateRequest
	return r
}

func (r ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRowsRequest) Execute() (*BatchResponseHubDbTableRowV3, *http.Response, error) {
	return r.ApiService.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRowsExecute(r)
}

/*
PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRows Update rows in batch in draft table

Updates multiple rows as a batch in the `draft` version of the table, with a maximum of 100 rows per call. See the endpoint `PATCH /tables/{tableIdOrName}/rows/{rowId}/draft` for details on updating a single row.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tableIdOrName The ID or name of the table
	@return ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRowsRequest
*/
func (a *RowsBatchApiService) PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRows(ctx context.Context, tableIdOrName string) ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRowsRequest {
	return ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRowsRequest{
		ApiService:    a,
		ctx:           ctx,
		tableIdOrName: tableIdOrName,
	}
}

// Execute executes the request
//
//	@return BatchResponseHubDbTableRowV3
func (a *RowsBatchApiService) PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRowsExecute(r ApiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRowsRequest) (*BatchResponseHubDbTableRowV3, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponseHubDbTableRowV3
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RowsBatchApiService.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateUpdateDraftTableRows")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/update"
	localVarPath = strings.Replace(localVarPath, "{"+"tableIdOrName"+"}", url.PathEscape(parameterToString(r.tableIdOrName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputHubDbTableRowV3BatchUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("batchInputHubDbTableRowV3BatchUpdateRequest is required and must be specified")
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
	localVarPostBody = r.batchInputHubDbTableRowV3BatchUpdateRequest
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
