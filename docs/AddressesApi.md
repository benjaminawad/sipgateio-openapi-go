# \AddressesApi

All URIs are relative to *https://api.sipgate.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAddress**](AddressesApi.md#GetAddress) | **Get** /addresses/{addressId} | Get a specific address
[**GetAddressNumbers**](AddressesApi.md#GetAddressNumbers) | **Get** /addresses/{addressId}/numbers | List all phone numbers of a specific address
[**GetAddresses**](AddressesApi.md#GetAddresses) | **Get** /addresses | List all addresses
[**GetItemCount**](AddressesApi.md#GetItemCount) | **Get** /addresses/itemCount | Get amount of addresses
[**ModifyAddress**](AddressesApi.md#ModifyAddress) | **Put** /addresses/{addressId} | Update an existing address



## GetAddress

> AddressResponse GetAddress(ctx, addressId).Execute()

Get a specific address

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
    addressId := int32(123) // int32 | The unique address identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressesApi.GetAddress(context.Background(), addressId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressesApi.GetAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddress`: AddressResponse
    fmt.Fprintf(os.Stdout, "Response from `AddressesApi.GetAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addressId** | **int32** | The unique address identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddressResponse**](AddressResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddressNumbers

> NumbersResponse GetAddressNumbers(ctx, addressId).Execute()

List all phone numbers of a specific address

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
    addressId := int32(123) // int32 | The unique address identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressesApi.GetAddressNumbers(context.Background(), addressId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressesApi.GetAddressNumbers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddressNumbers`: NumbersResponse
    fmt.Fprintf(os.Stdout, "Response from `AddressesApi.GetAddressNumbers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addressId** | **int32** | The unique address identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NumbersResponse**](NumbersResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddresses

> AddressesResponse GetAddresses(ctx).Execute()

List all addresses

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
    resp, r, err := apiClient.AddressesApi.GetAddresses(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressesApi.GetAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddresses`: AddressesResponse
    fmt.Fprintf(os.Stdout, "Response from `AddressesApi.GetAddresses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressesRequest struct via the builder pattern


### Return type

[**AddressesResponse**](AddressesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemCount

> ItemCountResponse GetItemCount(ctx).Execute()

Get amount of addresses

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
    resp, r, err := apiClient.AddressesApi.GetItemCount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressesApi.GetItemCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemCount`: ItemCountResponse
    fmt.Fprintf(os.Stdout, "Response from `AddressesApi.GetItemCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemCountRequest struct via the builder pattern


### Return type

[**ItemCountResponse**](ItemCountResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyAddress

> ModifyAddress(ctx, addressId).Body(body).Execute()

Update an existing address



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
    addressId := int32(abc123) // int32 | The unique address identifier
    body := *openapiclient.NewModifyAddressRequest("DE", "40219", "Düsseldorf") // ModifyAddressRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressesApi.ModifyAddress(context.Background(), addressId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressesApi.ModifyAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addressId** | **int32** | The unique address identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ModifyAddressRequest**](ModifyAddressRequest.md) |  | 

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

