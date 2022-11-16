/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


type JamfProVersionApi interface {

	/*
	V1JamfProVersionGet Return information about the Jamf Pro including the current version 

	Returns information about the Jamf Pro including the current version.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1JamfProVersionGetRequest
	*/
	V1JamfProVersionGet(ctx context.Context) ApiV1JamfProVersionGetRequest

	// V1JamfProVersionGetExecute executes the request
	//  @return JamfProVersion
	V1JamfProVersionGetExecute(r ApiV1JamfProVersionGetRequest) (*JamfProVersion, *http.Response, error)
}

// JamfProVersionApiService JamfProVersionApi service
type JamfProVersionApiService service

type ApiV1JamfProVersionGetRequest struct {
	ctx context.Context
	ApiService JamfProVersionApi
}

func (r ApiV1JamfProVersionGetRequest) Execute() (*JamfProVersion, *http.Response, error) {
	return r.ApiService.V1JamfProVersionGetExecute(r)
}

/*
V1JamfProVersionGet Return information about the Jamf Pro including the current version 

Returns information about the Jamf Pro including the current version.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1JamfProVersionGetRequest
*/
func (a *JamfProVersionApiService) V1JamfProVersionGet(ctx context.Context) ApiV1JamfProVersionGetRequest {
	return ApiV1JamfProVersionGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return JamfProVersion
func (a *JamfProVersionApiService) V1JamfProVersionGetExecute(r ApiV1JamfProVersionGetRequest) (*JamfProVersion, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JamfProVersion
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfProVersionApiService.V1JamfProVersionGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/jamf-pro-version"

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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