/*
CRM Pipelines

Pipelines represent distinct stages in a workflow, like closing a deal or servicing a support ticket. These endpoints provide access to read and modify pipelines in HubSpot. Pipelines support `deals` and `tickets` object types.  ## Pipeline ID validation  When calling endpoints that take pipelineId as a parameter, that ID must correspond to an existing, un-archived pipeline. Otherwise the request will fail with a `404 Not Found` response.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipelines

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// PipelineStageAuditsApiService PipelineStageAuditsApi service
type PipelineStageAuditsApiService service

type ApiGetCrmV3PipelinesObjectTypePipelineIdStagesStageIdAuditGetAuditRequest struct {
	ctx        context.Context
	ApiService *PipelineStageAuditsApiService
	objectType string
	stageId    string
}

func (r ApiGetCrmV3PipelinesObjectTypePipelineIdStagesStageIdAuditGetAuditRequest) Execute() (*CollectionResponsePublicAuditInfoNoPaging, *http.Response, error) {
	return r.ApiService.GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdAuditGetAuditExecute(r)
}

/*
GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdAuditGetAudit Return an audit of all changes to the pipeline stage

Return a reverse chronological list of all mutations that have occurred on the pipeline stage identified by `{stageId}`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @param stageId
 @return ApiGetCrmV3PipelinesObjectTypePipelineIdStagesStageIdAuditGetAuditRequest
*/
func (a *PipelineStageAuditsApiService) GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdAuditGetAudit(ctx context.Context, objectType string, stageId string) ApiGetCrmV3PipelinesObjectTypePipelineIdStagesStageIdAuditGetAuditRequest {
	return ApiGetCrmV3PipelinesObjectTypePipelineIdStagesStageIdAuditGetAuditRequest{
		ApiService: a,
		ctx:        ctx,
		objectType: objectType,
		stageId:    stageId,
	}
}

// Execute executes the request
//  @return CollectionResponsePublicAuditInfoNoPaging
func (a *PipelineStageAuditsApiService) GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdAuditGetAuditExecute(r ApiGetCrmV3PipelinesObjectTypePipelineIdStagesStageIdAuditGetAuditRequest) (*CollectionResponsePublicAuditInfoNoPaging, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponsePublicAuditInfoNoPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelineStageAuditsApiService.GetCrmV3PipelinesObjectTypePipelineIdStagesStageIdAuditGetAudit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/pipelines/{objectType}/{pipelineId}/stages/{stageId}/audit"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterToString(r.objectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"stageId"+"}", url.PathEscape(parameterToString(r.stageId, "")), -1)

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
