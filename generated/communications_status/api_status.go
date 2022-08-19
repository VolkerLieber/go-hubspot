/*
Subscriptions

Subscriptions allow contacts to control what forms of communications they receive. Contacts can decide whether they want to receive communication pertaining to a specific topic, brand, or an entire HubSpot account.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package communications_status

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/clarkmcc/go-hubspot"
	"net/url"
	"strings"
)

// StatusApiService StatusApi service
type StatusApiService service

type ApiGetCommunicationPreferencesV3StatusEmailEmailAddressGetEmailStatusRequest struct {
	ctx          context.Context
	ApiService   *StatusApiService
	emailAddress string
}

func (r ApiGetCommunicationPreferencesV3StatusEmailEmailAddressGetEmailStatusRequest) Execute() (*PublicSubscriptionStatusesResponse, *http.Response, error) {
	return r.ApiService.GetCommunicationPreferencesV3StatusEmailEmailAddressGetEmailStatusExecute(r)
}

/*
GetCommunicationPreferencesV3StatusEmailEmailAddressGetEmailStatus Get subscription statuses for a contact

Returns a list of subscriptions and their status for a given contact.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param emailAddress
 @return ApiGetCommunicationPreferencesV3StatusEmailEmailAddressGetEmailStatusRequest
*/
func (a *StatusApiService) GetCommunicationPreferencesV3StatusEmailEmailAddressGetEmailStatus(ctx context.Context, emailAddress string) ApiGetCommunicationPreferencesV3StatusEmailEmailAddressGetEmailStatusRequest {
	return ApiGetCommunicationPreferencesV3StatusEmailEmailAddressGetEmailStatusRequest{
		ApiService:   a,
		ctx:          ctx,
		emailAddress: emailAddress,
	}
}

// Execute executes the request
//  @return PublicSubscriptionStatusesResponse
func (a *StatusApiService) GetCommunicationPreferencesV3StatusEmailEmailAddressGetEmailStatusExecute(r ApiGetCommunicationPreferencesV3StatusEmailEmailAddressGetEmailStatusRequest) (*PublicSubscriptionStatusesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PublicSubscriptionStatusesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatusApiService.GetCommunicationPreferencesV3StatusEmailEmailAddressGetEmailStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/communication-preferences/v3/status/email/{emailAddress}"
	localVarPath = strings.Replace(localVarPath, "{"+"emailAddress"+"}", url.PathEscape(parameterToString(r.emailAddress, "")), -1)

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

type ApiPostCommunicationPreferencesV3SubscribeSubscribeRequest struct {
	ctx                                   context.Context
	ApiService                            *StatusApiService
	publicUpdateSubscriptionStatusRequest *PublicUpdateSubscriptionStatusRequest
}

func (r ApiPostCommunicationPreferencesV3SubscribeSubscribeRequest) PublicUpdateSubscriptionStatusRequest(publicUpdateSubscriptionStatusRequest PublicUpdateSubscriptionStatusRequest) ApiPostCommunicationPreferencesV3SubscribeSubscribeRequest {
	r.publicUpdateSubscriptionStatusRequest = &publicUpdateSubscriptionStatusRequest
	return r
}

func (r ApiPostCommunicationPreferencesV3SubscribeSubscribeRequest) Execute() (*PublicSubscriptionStatus, *http.Response, error) {
	return r.ApiService.PostCommunicationPreferencesV3SubscribeSubscribeExecute(r)
}

/*
PostCommunicationPreferencesV3SubscribeSubscribe Subscribe a contact

Subscribes a contact to the given subscription type. This API is not valid to use for subscribing a contact at a brand or portal level and will return an error.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCommunicationPreferencesV3SubscribeSubscribeRequest
*/
func (a *StatusApiService) PostCommunicationPreferencesV3SubscribeSubscribe(ctx context.Context) ApiPostCommunicationPreferencesV3SubscribeSubscribeRequest {
	return ApiPostCommunicationPreferencesV3SubscribeSubscribeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return PublicSubscriptionStatus
func (a *StatusApiService) PostCommunicationPreferencesV3SubscribeSubscribeExecute(r ApiPostCommunicationPreferencesV3SubscribeSubscribeRequest) (*PublicSubscriptionStatus, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PublicSubscriptionStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatusApiService.PostCommunicationPreferencesV3SubscribeSubscribe")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/communication-preferences/v3/subscribe"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.publicUpdateSubscriptionStatusRequest == nil {
		return localVarReturnValue, nil, reportError("publicUpdateSubscriptionStatusRequest is required and must be specified")
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
	localVarPostBody = r.publicUpdateSubscriptionStatusRequest
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

type ApiPostCommunicationPreferencesV3UnsubscribeUnsubscribeRequest struct {
	ctx                                   context.Context
	ApiService                            *StatusApiService
	publicUpdateSubscriptionStatusRequest *PublicUpdateSubscriptionStatusRequest
}

func (r ApiPostCommunicationPreferencesV3UnsubscribeUnsubscribeRequest) PublicUpdateSubscriptionStatusRequest(publicUpdateSubscriptionStatusRequest PublicUpdateSubscriptionStatusRequest) ApiPostCommunicationPreferencesV3UnsubscribeUnsubscribeRequest {
	r.publicUpdateSubscriptionStatusRequest = &publicUpdateSubscriptionStatusRequest
	return r
}

func (r ApiPostCommunicationPreferencesV3UnsubscribeUnsubscribeRequest) Execute() (*PublicSubscriptionStatus, *http.Response, error) {
	return r.ApiService.PostCommunicationPreferencesV3UnsubscribeUnsubscribeExecute(r)
}

/*
PostCommunicationPreferencesV3UnsubscribeUnsubscribe Unsubscribe a contact

Unsubscribes a contact from the given subscription type. This API is not valid to use for unsubscribing a contact at a brand or portal level and will return an error.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCommunicationPreferencesV3UnsubscribeUnsubscribeRequest
*/
func (a *StatusApiService) PostCommunicationPreferencesV3UnsubscribeUnsubscribe(ctx context.Context) ApiPostCommunicationPreferencesV3UnsubscribeUnsubscribeRequest {
	return ApiPostCommunicationPreferencesV3UnsubscribeUnsubscribeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return PublicSubscriptionStatus
func (a *StatusApiService) PostCommunicationPreferencesV3UnsubscribeUnsubscribeExecute(r ApiPostCommunicationPreferencesV3UnsubscribeUnsubscribeRequest) (*PublicSubscriptionStatus, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PublicSubscriptionStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatusApiService.PostCommunicationPreferencesV3UnsubscribeUnsubscribe")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/communication-preferences/v3/unsubscribe"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.publicUpdateSubscriptionStatusRequest == nil {
		return localVarReturnValue, nil, reportError("publicUpdateSubscriptionStatusRequest is required and must be specified")
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
	localVarPostBody = r.publicUpdateSubscriptionStatusRequest
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
