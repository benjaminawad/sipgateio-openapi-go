# \RestrictionsApi

All URIs are relative to *https://api.sipgate.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRestrictions**](RestrictionsApi.md#GetRestrictions) | **Get** /restrictions | List all restrictions based on the users product



## GetRestrictions

> RestrictionsResponse GetRestrictions(ctx).UserId(userId).Restriction(restriction).Execute()

List all restrictions based on the users product



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
    restriction := []string{"Restriction_example"} // []string | The requested restrictions (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RestrictionsApi.GetRestrictions(context.Background()).UserId(userId).Restriction(restriction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestrictionsApi.GetRestrictions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRestrictions`: RestrictionsResponse
    fmt.Fprintf(os.Stdout, "Response from `RestrictionsApi.GetRestrictions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | The unique user identifier | 
 **restriction** | **[]string** | The requested restrictions | 

### Return type

[**RestrictionsResponse**](RestrictionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

