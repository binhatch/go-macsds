/*
Hella Gutmann - macsDS (Data Services) - API collection

This document contains all relevant APIs for diagnostics (incl. DTCs), repair & maintenance information (RMI) and vehicle identification.

API version: V20240702-130718
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github.com/binhatch/go-macds

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// KTypeComponentMappingAPIService KTypeComponentMappingAPI service
type KTypeComponentMappingAPIService service

type ApiComponentMappingRequest struct {
	ctx context.Context
	ApiService *KTypeComponentMappingAPIService
	kType *string
	systemIdent *string
	componentId *string
}

// KType
func (r ApiComponentMappingRequest) KType(kType string) ApiComponentMappingRequest {
	r.kType = &kType
	return r
}

// System Identification
func (r ApiComponentMappingRequest) SystemIdent(systemIdent string) ApiComponentMappingRequest {
	r.systemIdent = &systemIdent
	return r
}

// Component Id
func (r ApiComponentMappingRequest) ComponentId(componentId string) ApiComponentMappingRequest {
	r.componentId = &componentId
	return r
}

func (r ApiComponentMappingRequest) Execute() (*KtypeComponentMappingResponse, *http.Response, error) {
	return r.ApiService.ComponentMappingExecute(r)
}

/*
ComponentMapping KType Component Mapping

The KType Component Mapping

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiComponentMappingRequest
*/
func (a *KTypeComponentMappingAPIService) ComponentMapping(ctx context.Context) ApiComponentMappingRequest {
	return ApiComponentMappingRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KtypeComponentMappingResponse
func (a *KTypeComponentMappingAPIService) ComponentMappingExecute(r ApiComponentMappingRequest) (*KtypeComponentMappingResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KtypeComponentMappingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KTypeComponentMappingAPIService.ComponentMapping")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/ktype/component/mapping"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.kType == nil {
		return localVarReturnValue, nil, reportError("kType is required and must be specified")
	}
	if r.systemIdent == nil {
		return localVarReturnValue, nil, reportError("systemIdent is required and must be specified")
	}
	if r.componentId == nil {
		return localVarReturnValue, nil, reportError("componentId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "kType", r.kType, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "systemIdent", r.systemIdent, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "componentId", r.componentId, "")
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