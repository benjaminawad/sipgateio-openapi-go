# \IdentityVerificationApi

All URIs are relative to *https://api.sipgate.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVerification**](IdentityVerificationApi.md#GetVerification) | **Get** /identityVerification | 



## GetVerification

> IdentityVerificationResponse GetVerification(ctx).AllowStart(allowStart).Execute()



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
    allowStart := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityVerificationApi.GetVerification(context.Background()).AllowStart(allowStart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityVerificationApi.GetVerification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVerification`: IdentityVerificationResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentityVerificationApi.GetVerification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allowStart** | **bool** |  | [default to false]

### Return type

[**IdentityVerificationResponse**](IdentityVerificationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

