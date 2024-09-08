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
	"io"
	"net/http"
	"net/url"
)


type SlasaAPI interface {

	/*
	V1SlasaGet Get the status of SLASA

	Get if SLASA has been accepted or not

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SlasaAPIV1SlasaGetRequest
	*/
	V1SlasaGet(ctx context.Context) SlasaAPIV1SlasaGetRequest

	// V1SlasaGetExecute executes the request
	//  @return SlasaAcceptance
	V1SlasaGetExecute(r SlasaAPIV1SlasaGetRequest) (*SlasaAcceptance, *http.Response, error)

	/*
	V1SlasaPost Accept the SLASA

	Accept the SLASA for Jamf Pro.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SlasaAPIV1SlasaPostRequest
	*/
	V1SlasaPost(ctx context.Context) SlasaAPIV1SlasaPostRequest

	// V1SlasaPostExecute executes the request
	V1SlasaPostExecute(r SlasaAPIV1SlasaPostRequest) (*http.Response, error)
}

// SlasaAPIService SlasaAPI service
type SlasaAPIService service

type SlasaAPIV1SlasaGetRequest struct {
	ctx context.Context
	ApiService SlasaAPI
}

func (r SlasaAPIV1SlasaGetRequest) Execute() (*SlasaAcceptance, *http.Response, error) {
	return r.ApiService.V1SlasaGetExecute(r)
}

/*
V1SlasaGet Get the status of SLASA

Get if SLASA has been accepted or not

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SlasaAPIV1SlasaGetRequest
*/
func (a *SlasaAPIService) V1SlasaGet(ctx context.Context) SlasaAPIV1SlasaGetRequest {
	return SlasaAPIV1SlasaGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SlasaAcceptance
func (a *SlasaAPIService) V1SlasaGetExecute(r SlasaAPIV1SlasaGetRequest) (*SlasaAcceptance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SlasaAcceptance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SlasaAPIService.V1SlasaGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/slasa"

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type SlasaAPIV1SlasaPostRequest struct {
	ctx context.Context
	ApiService SlasaAPI
}

func (r SlasaAPIV1SlasaPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1SlasaPostExecute(r)
}

/*
V1SlasaPost Accept the SLASA

Accept the SLASA for Jamf Pro.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SlasaAPIV1SlasaPostRequest
*/
func (a *SlasaAPIService) V1SlasaPost(ctx context.Context) SlasaAPIV1SlasaPostRequest {
	return SlasaAPIV1SlasaPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *SlasaAPIService) V1SlasaPostExecute(r SlasaAPIV1SlasaPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SlasaAPIService.V1SlasaPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/slasa"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}