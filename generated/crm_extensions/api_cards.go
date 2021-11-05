/*
CRM cards

Allows an app to extend the CRM UI by surfacing custom cards in the sidebar of record pages. These cards are defined up-front as part of app configuration, then populated by external data fetch requests when the record page is accessed by a user.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crm_extensions

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"

	"github.com/clarkmcc/go-hubspot/authorization"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// CardsApiService CardsApi service
type CardsApiService service

type ApiDeleteCrmV3ExtensionsCardsAppIdCardIdArchiveRequest struct {
	ctx        _context.Context
	ApiService *CardsApiService
	appId      int32
	cardId     string
}

func (r ApiDeleteCrmV3ExtensionsCardsAppIdCardIdArchiveRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteCrmV3ExtensionsCardsAppIdCardIdArchiveExecute(r)
}

/*
DeleteCrmV3ExtensionsCardsAppIdCardIdArchive Delete a card

Permanently deletes a card definition with the given ID. Once deleted, data fetch requests for this card will no longer be sent to your service. This can't be undone.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId The ID of the target app.
 @param cardId The ID of the card to delete.
 @return ApiDeleteCrmV3ExtensionsCardsAppIdCardIdArchiveRequest
*/
func (a *CardsApiService) DeleteCrmV3ExtensionsCardsAppIdCardIdArchive(ctx _context.Context, appId int32, cardId string) ApiDeleteCrmV3ExtensionsCardsAppIdCardIdArchiveRequest {
	return ApiDeleteCrmV3ExtensionsCardsAppIdCardIdArchiveRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
		cardId:     cardId,
	}
}

// Execute executes the request
func (a *CardsApiService) DeleteCrmV3ExtensionsCardsAppIdCardIdArchiveExecute(r ApiDeleteCrmV3ExtensionsCardsAppIdCardIdArchiveRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardsApiService.DeleteCrmV3ExtensionsCardsAppIdCardIdArchive")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/extensions/cards/{appId}/{cardId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", _neturl.PathEscape(parameterToString(r.appId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cardId"+"}", _neturl.PathEscape(parameterToString(r.cardId, "")), -1)

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
			if apiKey, ok := auth["developer_hapikey"]; ok {
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

type ApiGetCrmV3ExtensionsCardsAppIdCardIdGetByIdRequest struct {
	ctx        _context.Context
	ApiService *CardsApiService
	appId      int32
	cardId     string
}

func (r ApiGetCrmV3ExtensionsCardsAppIdCardIdGetByIdRequest) Execute() (CardResponse, *_nethttp.Response, error) {
	return r.ApiService.GetCrmV3ExtensionsCardsAppIdCardIdGetByIdExecute(r)
}

/*
GetCrmV3ExtensionsCardsAppIdCardIdGetById Get a card.

Returns the definition for a card with the given ID.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId The ID of the target app.
 @param cardId The ID of the target card.
 @return ApiGetCrmV3ExtensionsCardsAppIdCardIdGetByIdRequest
*/
func (a *CardsApiService) GetCrmV3ExtensionsCardsAppIdCardIdGetById(ctx _context.Context, appId int32, cardId string) ApiGetCrmV3ExtensionsCardsAppIdCardIdGetByIdRequest {
	return ApiGetCrmV3ExtensionsCardsAppIdCardIdGetByIdRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
		cardId:     cardId,
	}
}

// Execute executes the request
//  @return CardResponse
func (a *CardsApiService) GetCrmV3ExtensionsCardsAppIdCardIdGetByIdExecute(r ApiGetCrmV3ExtensionsCardsAppIdCardIdGetByIdRequest) (CardResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardsApiService.GetCrmV3ExtensionsCardsAppIdCardIdGetById")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/extensions/cards/{appId}/{cardId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", _neturl.PathEscape(parameterToString(r.appId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cardId"+"}", _neturl.PathEscape(parameterToString(r.cardId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(authorization.ContextAPIKeys).(map[string]authorization.APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
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

type ApiGetCrmV3ExtensionsCardsAppIdGetAllRequest struct {
	ctx        _context.Context
	ApiService *CardsApiService
	appId      int32
}

func (r ApiGetCrmV3ExtensionsCardsAppIdGetAllRequest) Execute() (CardListResponse, *_nethttp.Response, error) {
	return r.ApiService.GetCrmV3ExtensionsCardsAppIdGetAllExecute(r)
}

/*
GetCrmV3ExtensionsCardsAppIdGetAll Get all cards

Returns a list of cards for a given app.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId The ID of the target app.
 @return ApiGetCrmV3ExtensionsCardsAppIdGetAllRequest
*/
func (a *CardsApiService) GetCrmV3ExtensionsCardsAppIdGetAll(ctx _context.Context, appId int32) ApiGetCrmV3ExtensionsCardsAppIdGetAllRequest {
	return ApiGetCrmV3ExtensionsCardsAppIdGetAllRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
	}
}

// Execute executes the request
//  @return CardListResponse
func (a *CardsApiService) GetCrmV3ExtensionsCardsAppIdGetAllExecute(r ApiGetCrmV3ExtensionsCardsAppIdGetAllRequest) (CardListResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CardListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardsApiService.GetCrmV3ExtensionsCardsAppIdGetAll")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/extensions/cards/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", _neturl.PathEscape(parameterToString(r.appId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(authorization.ContextAPIKeys).(map[string]authorization.APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
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

type ApiPatchCrmV3ExtensionsCardsAppIdCardIdUpdateRequest struct {
	ctx              _context.Context
	ApiService       *CardsApiService
	appId            int32
	cardId           string
	cardPatchRequest *CardPatchRequest
}

// Card definition fields to be updated.
func (r ApiPatchCrmV3ExtensionsCardsAppIdCardIdUpdateRequest) CardPatchRequest(cardPatchRequest CardPatchRequest) ApiPatchCrmV3ExtensionsCardsAppIdCardIdUpdateRequest {
	r.cardPatchRequest = &cardPatchRequest
	return r
}

func (r ApiPatchCrmV3ExtensionsCardsAppIdCardIdUpdateRequest) Execute() (CardResponse, *_nethttp.Response, error) {
	return r.ApiService.PatchCrmV3ExtensionsCardsAppIdCardIdUpdateExecute(r)
}

/*
PatchCrmV3ExtensionsCardsAppIdCardIdUpdate Update a card

Update a card definition with new details.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId The ID of the target app.
 @param cardId The ID of the card to update.
 @return ApiPatchCrmV3ExtensionsCardsAppIdCardIdUpdateRequest
*/
func (a *CardsApiService) PatchCrmV3ExtensionsCardsAppIdCardIdUpdate(ctx _context.Context, appId int32, cardId string) ApiPatchCrmV3ExtensionsCardsAppIdCardIdUpdateRequest {
	return ApiPatchCrmV3ExtensionsCardsAppIdCardIdUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
		cardId:     cardId,
	}
}

// Execute executes the request
//  @return CardResponse
func (a *CardsApiService) PatchCrmV3ExtensionsCardsAppIdCardIdUpdateExecute(r ApiPatchCrmV3ExtensionsCardsAppIdCardIdUpdateRequest) (CardResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardsApiService.PatchCrmV3ExtensionsCardsAppIdCardIdUpdate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/extensions/cards/{appId}/{cardId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", _neturl.PathEscape(parameterToString(r.appId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cardId"+"}", _neturl.PathEscape(parameterToString(r.cardId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.cardPatchRequest == nil {
		return localVarReturnValue, nil, reportError("cardPatchRequest is required and must be specified")
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
	localVarPostBody = r.cardPatchRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(authorization.ContextAPIKeys).(map[string]authorization.APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
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

type ApiPostCrmV3ExtensionsCardsAppIdCreateRequest struct {
	ctx               _context.Context
	ApiService        *CardsApiService
	appId             int32
	cardCreateRequest *CardCreateRequest
}

// The new card definition.
func (r ApiPostCrmV3ExtensionsCardsAppIdCreateRequest) CardCreateRequest(cardCreateRequest CardCreateRequest) ApiPostCrmV3ExtensionsCardsAppIdCreateRequest {
	r.cardCreateRequest = &cardCreateRequest
	return r
}

func (r ApiPostCrmV3ExtensionsCardsAppIdCreateRequest) Execute() (CardResponse, *_nethttp.Response, error) {
	return r.ApiService.PostCrmV3ExtensionsCardsAppIdCreateExecute(r)
}

/*
PostCrmV3ExtensionsCardsAppIdCreate Create a new card

Defines a new card that will become active on an account when this app is installed.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId The ID of the target app.
 @return ApiPostCrmV3ExtensionsCardsAppIdCreateRequest
*/
func (a *CardsApiService) PostCrmV3ExtensionsCardsAppIdCreate(ctx _context.Context, appId int32) ApiPostCrmV3ExtensionsCardsAppIdCreateRequest {
	return ApiPostCrmV3ExtensionsCardsAppIdCreateRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
	}
}

// Execute executes the request
//  @return CardResponse
func (a *CardsApiService) PostCrmV3ExtensionsCardsAppIdCreateExecute(r ApiPostCrmV3ExtensionsCardsAppIdCreateRequest) (CardResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardsApiService.PostCrmV3ExtensionsCardsAppIdCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/extensions/cards/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", _neturl.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.cardCreateRequest == nil {
		return localVarReturnValue, nil, reportError("cardCreateRequest is required and must be specified")
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
	localVarPostBody = r.cardCreateRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(authorization.ContextAPIKeys).(map[string]authorization.APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
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
