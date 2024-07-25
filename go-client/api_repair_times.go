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


// RepairTimesAPIService RepairTimesAPI service
type RepairTimesAPIService service

type RepairTimesAPIRepairTimes1treeRequest struct {
	ctx context.Context
	ApiService *RepairTimesAPIService
	kType *string
	language *string
}

// KType
func (r RepairTimesAPIRepairTimes1treeRequest) KType(kType string) RepairTimesAPIRepairTimes1treeRequest {
	r.kType = &kType
	return r
}

// Two characters defining the language code; allowed values: de, en, nl, zh, it, hu, hr, fr, fi, es, el, tr, da, sv, sk, cs, ru, ro, pt, pl
func (r RepairTimesAPIRepairTimes1treeRequest) Language(language string) RepairTimesAPIRepairTimes1treeRequest {
	r.language = &language
	return r
}

func (r RepairTimesAPIRepairTimes1treeRequest) Execute() (*RepairTimesTreeResponse, *http.Response, error) {
	return r.ApiService.RepairTimes1treeExecute(r)
}

/*
RepairTimes1tree Get repair times tree

Returns category tree for repair times for a specified vehicle.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RepairTimesAPIRepairTimes1treeRequest
*/
func (a *RepairTimesAPIService) RepairTimes1tree(ctx context.Context) RepairTimesAPIRepairTimes1treeRequest {
	return RepairTimesAPIRepairTimes1treeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RepairTimesTreeResponse
func (a *RepairTimesAPIService) RepairTimes1treeExecute(r RepairTimesAPIRepairTimes1treeRequest) (*RepairTimesTreeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RepairTimesTreeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepairTimesAPIService.RepairTimes1tree")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/repair-times-tree"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.kType == nil {
		return localVarReturnValue, nil, reportError("kType is required and must be specified")
	}
	if r.language == nil {
		return localVarReturnValue, nil, reportError("language is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "kType", r.kType, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "language", r.language, "")
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

type RepairTimesAPIRepairTimes2detailsRequest struct {
	ctx context.Context
	ApiService *RepairTimesAPIService
	kType *string
	language *string
	categoryId *string
	subCategoryId *string
	genericArticleNo *string
}

// KType
func (r RepairTimesAPIRepairTimes2detailsRequest) KType(kType string) RepairTimesAPIRepairTimes2detailsRequest {
	r.kType = &kType
	return r
}

// Two characters defining the language code; allowed values: de, en, nl, zh, it, hu, hr, fr, fi, es, el, tr, da, sv, sk, cs, ru, ro, pt, pl
func (r RepairTimesAPIRepairTimes2detailsRequest) Language(language string) RepairTimesAPIRepairTimes2detailsRequest {
	r.language = &language
	return r
}

// The id of the desired category, as describe in the /tree endpoint.  A valid combination of query parameters can be: - categoryId &amp; subCategoryId - categoryId &amp; subCategoryId &amp; genericArticleNo - categoryId &amp; genericArticleNo - genericArticleNo
func (r RepairTimesAPIRepairTimes2detailsRequest) CategoryId(categoryId string) RepairTimesAPIRepairTimes2detailsRequest {
	r.categoryId = &categoryId
	return r
}

// The id of the desired sub category, as describe in the /tree endpoint.  A valid combination of query parameters can be: - categoryId &amp; subCategoryId - categoryId &amp; subCategoryId &amp; genericArticleNo - categoryId &amp; genericArticleNo - genericArticleNo
func (r RepairTimesAPIRepairTimes2detailsRequest) SubCategoryId(subCategoryId string) RepairTimesAPIRepairTimes2detailsRequest {
	r.subCategoryId = &subCategoryId
	return r
}

// The number of the generic article number (spare part/component) to get specific repair time jobs related to the component.  A valid combination of query parameters can be: - categoryId &amp; subCategoryId - categoryId &amp; subCategoryId &amp; genericArticleNo - categoryId &amp; genericArticleNo - genericArticleNo
func (r RepairTimesAPIRepairTimes2detailsRequest) GenericArticleNo(genericArticleNo string) RepairTimesAPIRepairTimes2detailsRequest {
	r.genericArticleNo = &genericArticleNo
	return r
}

func (r RepairTimesAPIRepairTimes2detailsRequest) Execute() (*RepairTimesJobResponse, *http.Response, error) {
	return r.ApiService.RepairTimes2detailsExecute(r)
}

/*
RepairTimes2details Get repair times details

Returns all jobs and their repair times for a specified vehicle.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RepairTimesAPIRepairTimes2detailsRequest
*/
func (a *RepairTimesAPIService) RepairTimes2details(ctx context.Context) RepairTimesAPIRepairTimes2detailsRequest {
	return RepairTimesAPIRepairTimes2detailsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RepairTimesJobResponse
func (a *RepairTimesAPIService) RepairTimes2detailsExecute(r RepairTimesAPIRepairTimes2detailsRequest) (*RepairTimesJobResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RepairTimesJobResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepairTimesAPIService.RepairTimes2details")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/repair-times"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.kType == nil {
		return localVarReturnValue, nil, reportError("kType is required and must be specified")
	}
	if r.language == nil {
		return localVarReturnValue, nil, reportError("language is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "kType", r.kType, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "language", r.language, "")
	if r.categoryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "categoryId", r.categoryId, "")
	}
	if r.subCategoryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subCategoryId", r.subCategoryId, "")
	}
	if r.genericArticleNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "genericArticleNo", r.genericArticleNo, "")
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
