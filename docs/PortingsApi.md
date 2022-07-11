# \PortingsApi

All URIs are relative to *https://api.sipgate.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPorting**](PortingsApi.md#GetPorting) | **Get** /portings/{portingId} | Get an existing porting
[**GetPortings**](PortingsApi.md#GetPortings) | **Get** /portings | List all phone number portings
[**RevokePorting**](PortingsApi.md#RevokePorting) | **Delete** /portings/{portingId} | Delete/revoke an existing porting



## GetPorting

> PortingResponse GetPorting(ctx, portingId).Execute()

Get an existing porting

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
    portingId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortingsApi.GetPorting(context.Background(), portingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortingsApi.GetPorting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPorting`: PortingResponse
    fmt.Fprintf(os.Stdout, "Response from `PortingsApi.GetPorting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portingId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PortingResponse**](PortingResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortings

> PortingsResponse GetPortings(ctx).Execute()

List all phone number portings



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
    resp, r, err := apiClient.PortingsApi.GetPortings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortingsApi.GetPortings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPortings`: PortingsResponse
    fmt.Fprintf(os.Stdout, "Response from `PortingsApi.GetPortings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortingsRequest struct via the builder pattern


### Return type

[**PortingsResponse**](PortingsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokePorting

> PortingsResponse RevokePorting(ctx, portingId).Execute()

Delete/revoke an existing porting

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
    portingId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortingsApi.RevokePorting(context.Background(), portingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortingsApi.RevokePorting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokePorting`: PortingsResponse
    fmt.Fprintf(os.Stdout, "Response from `PortingsApi.RevokePorting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portingId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokePortingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PortingsResponse**](PortingsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

