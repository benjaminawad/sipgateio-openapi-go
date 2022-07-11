# \CallrestrictionsApi

All URIs are relative to *https://api.sipgate.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCallRestrictions**](CallrestrictionsApi.md#GetCallRestrictions) | **Get** /callrestrictions | Get call restrictions data
[**SetCallRestriction**](CallrestrictionsApi.md#SetCallRestriction) | **Post** /{userId}/callrestrictions/{restriction} | Enable/disable restriction



## GetCallRestrictions

> CallRestrictionResponse GetCallRestrictions(ctx).UserIds(userIds).Execute()

Get call restrictions data

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userIds := []string{"Inner_example"} // []string | List of extensions (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CallrestrictionsApi.GetCallRestrictions(context.Background()).UserIds(userIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallrestrictionsApi.GetCallRestrictions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCallRestrictions`: CallRestrictionResponse
    fmt.Fprintf(os.Stdout, "Response from `CallrestrictionsApi.GetCallRestrictions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCallRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userIds** | **[]string** | List of extensions | 

### Return type

[**CallRestrictionResponse**](CallRestrictionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCallRestriction

> CallRestrictionResponse SetCallRestriction(ctx, userId, restriction).Body(body).Execute()

Enable/disable restriction

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userId := "w0" // string | The unique user identifier
    restriction := "roaming" // string | The name of the call restriction
    body := *openapiclient.NewSetCallRestrictionRequest() // SetCallRestrictionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CallrestrictionsApi.SetCallRestriction(context.Background(), userId, restriction).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallrestrictionsApi.SetCallRestriction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetCallRestriction`: CallRestrictionResponse
    fmt.Fprintf(os.Stdout, "Response from `CallrestrictionsApi.SetCallRestriction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**restriction** | **string** | The name of the call restriction | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCallRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**SetCallRestrictionRequest**](SetCallRestrictionRequest.md) |  | 

### Return type

[**CallRestrictionResponse**](CallRestrictionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

