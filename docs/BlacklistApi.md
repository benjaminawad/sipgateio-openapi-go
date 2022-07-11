# \BlacklistApi

All URIs are relative to *https://api.sipgate.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIncomingEntry**](BlacklistApi.md#AddIncomingEntry) | **Post** /blacklist/incoming | Create a new incoming blacklist entry
[**GetIncomingEntries**](BlacklistApi.md#GetIncomingEntries) | **Get** /blacklist/incoming | Get your incoming blacklist entries
[**RemoveIncomingEntry**](BlacklistApi.md#RemoveIncomingEntry) | **Delete** /blacklist/incoming/{phoneNumber} | Delete an incoming blacklist entry



## AddIncomingEntry

> IncomingEntry AddIncomingEntry(ctx).Body(body).Execute()

Create a new incoming blacklist entry

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
    body := *openapiclient.NewIncomingEntry("+4902117654321") // IncomingEntry | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlacklistApi.AddIncomingEntry(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlacklistApi.AddIncomingEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddIncomingEntry`: IncomingEntry
    fmt.Fprintf(os.Stdout, "Response from `BlacklistApi.AddIncomingEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddIncomingEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IncomingEntry**](IncomingEntry.md) |  | 

### Return type

[**IncomingEntry**](IncomingEntry.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncomingEntries

> IncomingEntries GetIncomingEntries(ctx).Execute()

Get your incoming blacklist entries

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
    resp, r, err := apiClient.BlacklistApi.GetIncomingEntries(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlacklistApi.GetIncomingEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIncomingEntries`: IncomingEntries
    fmt.Fprintf(os.Stdout, "Response from `BlacklistApi.GetIncomingEntries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncomingEntriesRequest struct via the builder pattern


### Return type

[**IncomingEntries**](IncomingEntries.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveIncomingEntry

> RemoveIncomingEntry(ctx, phoneNumber).Execute()

Delete an incoming blacklist entry

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
    phoneNumber := "phoneNumber_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlacklistApi.RemoveIncomingEntry(context.Background(), phoneNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlacklistApi.RemoveIncomingEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumber** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveIncomingEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

