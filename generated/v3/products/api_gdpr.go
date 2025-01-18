/*
Products

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package products

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/clarkmcc/go-hubspot"
	"net/url"
)

// GDPRApiService GDPRApi service
type GDPRApiService service

type ApiPostCrmV3ObjectsProductsGdprDeletePurgeRequest struct {
	ctx                   context.Context
	ApiService            *GDPRApiService
	publicGdprDeleteInput *PublicGdprDeleteInput
}

func (r ApiPostCrmV3ObjectsProductsGdprDeletePurgeRequest) PublicGdprDeleteInput(publicGdprDeleteInput PublicGdprDeleteInput) ApiPostCrmV3ObjectsProductsGdprDeletePurgeRequest {
	r.publicGdprDeleteInput = &publicGdprDeleteInput
	return r
}

func (r ApiPostCrmV3ObjectsProductsGdprDeletePurgeRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsProductsGdprDeletePurgeExecute(r)
}

/*
PostCrmV3ObjectsProductsGdprDeletePurge GDPR DELETE

Permanently delete a contact and all associated content to follow GDPR. Use optional property 'idProperty' set to 'email' to identify contact by email address. If email address is not found, the email address will be added to a blocklist and prevent it from being used in the future.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostCrmV3ObjectsProductsGdprDeletePurgeRequest
*/
func (a *GDPRApiService) PostCrmV3ObjectsProductsGdprDeletePurge(ctx context.Context) ApiPostCrmV3ObjectsProductsGdprDeletePurgeRequest {
	return ApiPostCrmV3ObjectsProductsGdprDeletePurgeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *GDPRApiService) PostCrmV3ObjectsProductsGdprDeletePurgeExecute(r ApiPostCrmV3ObjectsProductsGdprDeletePurgeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GDPRApiService.PostCrmV3ObjectsProductsGdprDeletePurge")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/products/gdpr-delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.publicGdprDeleteInput == nil {
		return nil, reportError("publicGdprDeleteInput is required and must be specified")
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
	localVarPostBody = r.publicGdprDeleteInput
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
