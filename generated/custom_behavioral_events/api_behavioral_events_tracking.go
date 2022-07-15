/*
Custom Behavioral Events API

HTTP API for triggering instances of custom behavioral events

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package custom_behavioral_events

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// BehavioralEventsTrackingApiService BehavioralEventsTrackingApi service
type BehavioralEventsTrackingApiService service

type ApiPostEventsV3SendRequest struct {
	ctx                                  context.Context
	ApiService                           *BehavioralEventsTrackingApiService
	behavioralEventHttpCompletionRequest *BehavioralEventHttpCompletionRequest
}

func (r ApiPostEventsV3SendRequest) BehavioralEventHttpCompletionRequest(behavioralEventHttpCompletionRequest BehavioralEventHttpCompletionRequest) ApiPostEventsV3SendRequest {
	r.behavioralEventHttpCompletionRequest = &behavioralEventHttpCompletionRequest
	return r
}

func (r ApiPostEventsV3SendRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostEventsV3SendExecute(r)
}

/*
PostEventsV3Send Sends Custom Behavioral Event

Endpoint to send an instance of a behavioral event

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostEventsV3SendRequest
*/
func (a *BehavioralEventsTrackingApiService) PostEventsV3Send(ctx context.Context) ApiPostEventsV3SendRequest {
	return ApiPostEventsV3SendRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *BehavioralEventsTrackingApiService) PostEventsV3SendExecute(r ApiPostEventsV3SendRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BehavioralEventsTrackingApiService.PostEventsV3Send")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/v3/send"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.behavioralEventHttpCompletionRequest == nil {
		return nil, reportError("behavioralEventHttpCompletionRequest is required and must be specified")
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
	localVarPostBody = r.behavioralEventHttpCompletionRequest
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
