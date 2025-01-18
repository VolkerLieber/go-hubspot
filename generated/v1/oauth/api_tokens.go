/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oauth

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// TokensApiService TokensApi service
type TokensApiService service

type ApiPostOauthV1TokenCreateRequest struct {
	ctx          context.Context
	ApiService   *TokensApiService
	grantType    *string
	code         *string
	redirectUri  *string
	clientId     *string
	clientSecret *string
	refreshToken *string
}

func (r ApiPostOauthV1TokenCreateRequest) GrantType(grantType string) ApiPostOauthV1TokenCreateRequest {
	r.grantType = &grantType
	return r
}

func (r ApiPostOauthV1TokenCreateRequest) Code(code string) ApiPostOauthV1TokenCreateRequest {
	r.code = &code
	return r
}

func (r ApiPostOauthV1TokenCreateRequest) RedirectUri(redirectUri string) ApiPostOauthV1TokenCreateRequest {
	r.redirectUri = &redirectUri
	return r
}

func (r ApiPostOauthV1TokenCreateRequest) ClientId(clientId string) ApiPostOauthV1TokenCreateRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostOauthV1TokenCreateRequest) ClientSecret(clientSecret string) ApiPostOauthV1TokenCreateRequest {
	r.clientSecret = &clientSecret
	return r
}

func (r ApiPostOauthV1TokenCreateRequest) RefreshToken(refreshToken string) ApiPostOauthV1TokenCreateRequest {
	r.refreshToken = &refreshToken
	return r
}

func (r ApiPostOauthV1TokenCreateRequest) Execute() (*TokenResponseIF, *http.Response, error) {
	return r.ApiService.PostOauthV1TokenCreateExecute(r)
}

/*
PostOauthV1TokenCreate Method for PostOauthV1TokenCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostOauthV1TokenCreateRequest
*/
func (a *TokensApiService) PostOauthV1TokenCreate(ctx context.Context) ApiPostOauthV1TokenCreateRequest {
	return ApiPostOauthV1TokenCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TokenResponseIF
func (a *TokensApiService) PostOauthV1TokenCreateExecute(r ApiPostOauthV1TokenCreateRequest) (*TokenResponseIF, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TokenResponseIF
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokensApiService.PostOauthV1TokenCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/v1/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

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
	if r.grantType != nil {
		localVarFormParams.Add("grant_type", parameterToString(*r.grantType, ""))
	}
	if r.code != nil {
		localVarFormParams.Add("code", parameterToString(*r.code, ""))
	}
	if r.redirectUri != nil {
		localVarFormParams.Add("redirect_uri", parameterToString(*r.redirectUri, ""))
	}
	if r.clientId != nil {
		localVarFormParams.Add("client_id", parameterToString(*r.clientId, ""))
	}
	if r.clientSecret != nil {
		localVarFormParams.Add("client_secret", parameterToString(*r.clientSecret, ""))
	}
	if r.refreshToken != nil {
		localVarFormParams.Add("refresh_token", parameterToString(*r.refreshToken, ""))
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
