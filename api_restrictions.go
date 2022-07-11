/*
sipgate API

This is the sipgate REST API documentation. We build our applications on this API and we invite you to use it too.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)


// RestrictionsApiService RestrictionsApi service
type RestrictionsApiService service

type ApiGetRestrictionsRequest struct {
	ctx context.Context
	ApiService *RestrictionsApiService
	userId *string
	restriction *[]string
}

// The unique user identifier
func (r ApiGetRestrictionsRequest) UserId(userId string) ApiGetRestrictionsRequest {
	r.userId = &userId
	return r
}

// The requested restrictions
func (r ApiGetRestrictionsRequest) Restriction(restriction []string) ApiGetRestrictionsRequest {
	r.restriction = &restriction
	return r
}

func (r ApiGetRestrictionsRequest) Execute() (*RestrictionsResponse, *http.Response, error) {
	return r.ApiService.GetRestrictionsExecute(r)
}

/*
GetRestrictions List all restrictions based on the users product

Get restrictions for actions the authenticated user can perform for himself or another user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetRestrictionsRequest
*/
func (a *RestrictionsApiService) GetRestrictions(ctx context.Context) ApiGetRestrictionsRequest {
	return ApiGetRestrictionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RestrictionsResponse
func (a *RestrictionsApiService) GetRestrictionsExecute(r ApiGetRestrictionsRequest) (*RestrictionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RestrictionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RestrictionsApiService.GetRestrictions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/restrictions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userId == nil {
		return localVarReturnValue, nil, reportError("userId is required and must be specified")
	}

	localVarQueryParams.Add("userId", parameterToString(*r.userId, ""))
	if r.restriction != nil {
		t := *r.restriction
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("restriction", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("restriction", parameterToString(t, "multi"))
		}
	}
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
