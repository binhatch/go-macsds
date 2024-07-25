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


// RecallCampaignsAPIService RecallCampaignsAPI service
type RecallCampaignsAPIService service

type ApiRecallCampaigns1treeRequest struct {
	ctx context.Context
	ApiService *RecallCampaignsAPIService
	kType *string
	language *string
	country *string
	includeOutdatedInfos *string
}

// KType
func (r ApiRecallCampaigns1treeRequest) KType(kType string) ApiRecallCampaigns1treeRequest {
	r.kType = &kType
	return r
}

// Two characters defining the language code; allowed values: de, en, nl, zh, it, hu, hr, fr, fi, es, el, tr, da, sv, sk, cs, ru, ro, pt, pl
func (r ApiRecallCampaigns1treeRequest) Language(language string) ApiRecallCampaigns1treeRequest {
	r.language = &language
	return r
}

// Two characters defining the ISO country code
func (r ApiRecallCampaigns1treeRequest) Country(country string) ApiRecallCampaigns1treeRequest {
	r.country = &country
	return r
}

// Include Outdated Info
func (r ApiRecallCampaigns1treeRequest) IncludeOutdatedInfos(includeOutdatedInfos string) ApiRecallCampaigns1treeRequest {
	r.includeOutdatedInfos = &includeOutdatedInfos
	return r
}

func (r ApiRecallCampaigns1treeRequest) Execute() ([]RecallGroup, *http.Response, error) {
	return r.ApiService.RecallCampaigns1treeExecute(r)
}

/*
RecallCampaigns1tree Get recall campaigns tree

Get security-related actions coordinated via federal instituations (e.g. the Kraftfahrbundesamt (KBA) in Germany).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRecallCampaigns1treeRequest
*/
func (a *RecallCampaignsAPIService) RecallCampaigns1tree(ctx context.Context) ApiRecallCampaigns1treeRequest {
	return ApiRecallCampaigns1treeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []RecallGroup
func (a *RecallCampaignsAPIService) RecallCampaigns1treeExecute(r ApiRecallCampaigns1treeRequest) ([]RecallGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []RecallGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecallCampaignsAPIService.RecallCampaigns1tree")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/recall-campaigns/tree"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.kType == nil {
		return localVarReturnValue, nil, reportError("kType is required and must be specified")
	}
	if r.language == nil {
		return localVarReturnValue, nil, reportError("language is required and must be specified")
	}
	if r.country == nil {
		return localVarReturnValue, nil, reportError("country is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "kType", r.kType, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "language", r.language, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "")
	if r.includeOutdatedInfos != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeOutdatedInfos", r.includeOutdatedInfos, "")
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
