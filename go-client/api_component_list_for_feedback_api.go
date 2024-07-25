/*
Hella Gutmann - macsDS (Data Services) - API collection

This document contains all relevant APIs for diagnostics (incl. DTCs), repair & maintenance information (RMI) and vehicle identification.

API version: V20240702-130718
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package go-macds

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// ComponentListForFeedbackAPIAPIService ComponentListForFeedbackAPIAPI service
type ComponentListForFeedbackAPIAPIService service

type ApiFeedbackpredictioncomponentsRequest struct {
	ctx context.Context
	ApiService *ComponentListForFeedbackAPIAPIService
	language *string
}

// Two characters defining the language code; allowed values: de, en, nl, zh, it, hu, hr, fr, fi, es, el, tr, da, sv, sk, cs, ru, ro, pt, pl
func (r ApiFeedbackpredictioncomponentsRequest) Language(language string) ApiFeedbackpredictioncomponentsRequest {
	r.language = &language
	return r
}

func (r ApiFeedbackpredictioncomponentsRequest) Execute() ([]DtcFeedbackPredictionComponent, *http.Response, error) {
	return r.ApiService.FeedbackpredictioncomponentsExecute(r)
}

/*
Feedbackpredictioncomponents FEEDBACK PREDICTION COMPONENTS API

API to get a list of available vehicle components to choose from for the feedback API

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFeedbackpredictioncomponentsRequest
*/
func (a *ComponentListForFeedbackAPIAPIService) Feedbackpredictioncomponents(ctx context.Context) ApiFeedbackpredictioncomponentsRequest {
	return ApiFeedbackpredictioncomponentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []DtcFeedbackPredictionComponent
func (a *ComponentListForFeedbackAPIAPIService) FeedbackpredictioncomponentsExecute(r ApiFeedbackpredictioncomponentsRequest) ([]DtcFeedbackPredictionComponent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []DtcFeedbackPredictionComponent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentListForFeedbackAPIAPIService.Feedbackpredictioncomponents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/feedback/prediction/components"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.language == nil {
		return localVarReturnValue, nil, reportError("language is required and must be specified")
	}

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
