/*
Calling Extensions

Provides a way for apps to add custom calling options to a contact record. This works in conjunction with the [Calling SDK](#), which is used to build your phone/calling UI. The endpoints here allow your service to appear as an option to HubSpot users when they access the *Call* action on a contact record. Once accessed, your custom phone/calling UI will be displayed in an iframe at the specified URL with the specified dimensions on that record.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package calling

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/clarkmcc/go-hubspot"
	"net/url"
	"strings"
)

// RecordingSettingsApiService RecordingSettingsApi service
type RecordingSettingsApiService service

type ApiGetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormatRequest struct {
	ctx        context.Context
	ApiService *RecordingSettingsApiService
	appId      int32
}

func (r ApiGetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormatRequest) Execute() (*RecordingSettingsResponse, *http.Response, error) {
	return r.ApiService.GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormatExecute(r)
}

/*
GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormat Method for GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormat

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param appId
	@return ApiGetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormatRequest
*/
func (a *RecordingSettingsApiService) GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormat(ctx context.Context, appId int32) ApiGetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormatRequest {
	return ApiGetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormatRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
	}
}

// Execute executes the request
//
//	@return RecordingSettingsResponse
func (a *RecordingSettingsApiService) GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormatExecute(r ApiGetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormatRequest) (*RecordingSettingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RecordingSettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordingSettingsApiService.GetCrmV3ExtensionsCallingAppIdSettingsRecordingGetUrlFormat")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/extensions/calling/{appId}/settings/recording"
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

type ApiPatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormatRequest struct {
	ctx                           context.Context
	ApiService                    *RecordingSettingsApiService
	appId                         int32
	recordingSettingsPatchRequest *RecordingSettingsPatchRequest
}

func (r ApiPatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormatRequest) RecordingSettingsPatchRequest(recordingSettingsPatchRequest RecordingSettingsPatchRequest) ApiPatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormatRequest {
	r.recordingSettingsPatchRequest = &recordingSettingsPatchRequest
	return r
}

func (r ApiPatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormatRequest) Execute() (*RecordingSettingsResponse, *http.Response, error) {
	return r.ApiService.PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormatExecute(r)
}

/*
PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormat Method for PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormat

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param appId
	@return ApiPatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormatRequest
*/
func (a *RecordingSettingsApiService) PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormat(ctx context.Context, appId int32) ApiPatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormatRequest {
	return ApiPatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormatRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
	}
}

// Execute executes the request
//
//	@return RecordingSettingsResponse
func (a *RecordingSettingsApiService) PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormatExecute(r ApiPatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormatRequest) (*RecordingSettingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RecordingSettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordingSettingsApiService.PatchCrmV3ExtensionsCallingAppIdSettingsRecordingUpdateUrlFormat")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/extensions/calling/{appId}/settings/recording"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.recordingSettingsPatchRequest == nil {
		return localVarReturnValue, nil, reportError("recordingSettingsPatchRequest is required and must be specified")
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
	localVarPostBody = r.recordingSettingsPatchRequest
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

type ApiPostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormatRequest struct {
	ctx                      context.Context
	ApiService               *RecordingSettingsApiService
	appId                    int32
	recordingSettingsRequest *RecordingSettingsRequest
}

func (r ApiPostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormatRequest) RecordingSettingsRequest(recordingSettingsRequest RecordingSettingsRequest) ApiPostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormatRequest {
	r.recordingSettingsRequest = &recordingSettingsRequest
	return r
}

func (r ApiPostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormatRequest) Execute() (*RecordingSettingsResponse, *http.Response, error) {
	return r.ApiService.PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormatExecute(r)
}

/*
PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormat Method for PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormat

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param appId
	@return ApiPostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormatRequest
*/
func (a *RecordingSettingsApiService) PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormat(ctx context.Context, appId int32) ApiPostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormatRequest {
	return ApiPostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormatRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
	}
}

// Execute executes the request
//
//	@return RecordingSettingsResponse
func (a *RecordingSettingsApiService) PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormatExecute(r ApiPostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormatRequest) (*RecordingSettingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RecordingSettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordingSettingsApiService.PostCrmV3ExtensionsCallingAppIdSettingsRecordingRegisterUrlFormat")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/extensions/calling/{appId}/settings/recording"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.recordingSettingsRequest == nil {
		return localVarReturnValue, nil, reportError("recordingSettingsRequest is required and must be specified")
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
	localVarPostBody = r.recordingSettingsRequest
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
