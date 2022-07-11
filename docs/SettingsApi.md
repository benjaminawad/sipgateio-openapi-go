# \SettingsApi

All URIs are relative to *https://api.sipgate.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSipgateIoUrls**](SettingsApi.md#GetSipgateIoUrls) | **Get** /settings/sipgateio | List sipgate.io global settings
[**SetSipgateIoUrls**](SettingsApi.md#SetSipgateIoUrls) | **Put** /settings/sipgateio | Update sipgate.io global settings



## GetSipgateIoUrls

> SipgateIoUrlResponse GetSipgateIoUrls(ctx).Execute()

List sipgate.io global settings

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.GetSipgateIoUrls(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetSipgateIoUrls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSipgateIoUrls`: SipgateIoUrlResponse
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetSipgateIoUrls`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSipgateIoUrlsRequest struct via the builder pattern


### Return type

[**SipgateIoUrlResponse**](SipgateIoUrlResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSipgateIoUrls

> SetSipgateIoUrls(ctx).Body(body).Execute()

Update sipgate.io global settings

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
    body := *openapiclient.NewSetSipgateIoAccountSettingsRequest("https://io.sipgate.beer/my/incoming/url", "https://io.sipgate.rehab/my/outgoing/url") // SetSipgateIoAccountSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.SetSipgateIoUrls(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.SetSipgateIoUrls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetSipgateIoUrlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SetSipgateIoAccountSettingsRequest**](SetSipgateIoAccountSettingsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

