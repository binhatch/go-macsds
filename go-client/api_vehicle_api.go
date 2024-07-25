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


// VehicleAPIAPIService VehicleAPIAPI service
type VehicleAPIAPIService service

type ApiGetVehiclesRequest struct {
	ctx context.Context
	ApiService *VehicleAPIAPIService
	kType *string
	modelId *string
}

// KType
func (r ApiGetVehiclesRequest) KType(kType string) ApiGetVehiclesRequest {
	r.kType = &kType
	return r
}

// Model Id
func (r ApiGetVehiclesRequest) ModelId(modelId string) ApiGetVehiclesRequest {
	r.modelId = &modelId
	return r
}

func (r ApiGetVehiclesRequest) Execute() ([]VehicleReferenceX, *http.Response, error) {
	return r.ApiService.GetVehiclesExecute(r)
}

/*
GetVehicles Vehicles

Get available vehicles based on a Ktype

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetVehiclesRequest
*/
func (a *VehicleAPIAPIService) GetVehicles(ctx context.Context) ApiGetVehiclesRequest {
	return ApiGetVehiclesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []VehicleReferenceX
func (a *VehicleAPIAPIService) GetVehiclesExecute(r ApiGetVehiclesRequest) ([]VehicleReferenceX, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []VehicleReferenceX
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VehicleAPIAPIService.GetVehicles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/vehicles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.kType == nil {
		return localVarReturnValue, nil, reportError("kType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "kType", r.kType, "")
	if r.modelId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "modelId", r.modelId, "")
	}
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