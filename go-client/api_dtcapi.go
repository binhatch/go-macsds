/*
Hella Gutmann - macsDS (Data Services) - API collection

This document contains all relevant APIs for diagnostics (incl. DTCs), repair & maintenance information (RMI) and vehicle identification.

API version: V20240702-130718
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gomacds

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// DTCAPIAPIService DTCAPIAPI service
type DTCAPIAPIService service

type ApiDtcsRequest struct {
	ctx context.Context
	ApiService *DTCAPIAPIService
	dtcRequestDto *DtcRequestDto
}

// The request requires at least the VIN, a list of DTCs and the language.
func (r ApiDtcsRequest) DtcRequestDto(dtcRequestDto DtcRequestDto) ApiDtcsRequest {
	r.dtcRequestDto = &dtcRequestDto
	return r
}

func (r ApiDtcsRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.DtcsExecute(r)
}

/*
Dtcs DTC API

API to get trouble code (DTC) interpretations as well as a prediction of a root cause.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDtcsRequest
*/
func (a *DTCAPIAPIService) Dtcs(ctx context.Context) ApiDtcsRequest {
	return ApiDtcsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Response
func (a *DTCAPIAPIService) DtcsExecute(r ApiDtcsRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DTCAPIAPIService.Dtcs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v4/dtcs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dtcRequestDto == nil {
		return localVarReturnValue, nil, reportError("dtcRequestDto is required and must be specified")
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
	// body params
	localVarPostBody = r.dtcRequestDto
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
