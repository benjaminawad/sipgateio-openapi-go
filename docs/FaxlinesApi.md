# \FaxlinesApi

All URIs are relative to *https://api.sipgate.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeFaxlineAlias**](FaxlinesApi.md#ChangeFaxlineAlias) | **Put** /{userId}/faxlines/{faxlineId} | Update alias for an existing fax line
[**ChangeFaxlineCallerIdSettings**](FaxlinesApi.md#ChangeFaxlineCallerIdSettings) | **Put** /{userId}/faxlines/{faxlineId}/callerid | Update the caller ID for an existing fax line
[**ChangeTagline**](FaxlinesApi.md#ChangeTagline) | **Put** /{userId}/faxlines/{faxlineId}/tagline | Update tagline of an existing fax line
[**CreateFaxline**](FaxlinesApi.md#CreateFaxline) | **Post** /{userId}/faxlines | Create a new fax line
[**DeleteFaxline**](FaxlinesApi.md#DeleteFaxline) | **Delete** /{userId}/faxlines/{faxlineId} | Delete an existing fax line
[**GetFaxlineCallerId**](FaxlinesApi.md#GetFaxlineCallerId) | **Get** /{userId}/faxlines/{faxlineId}/callerid | Get the caller ID of a specific fax line
[**GetFaxlineNumbers**](FaxlinesApi.md#GetFaxlineNumbers) | **Get** /{userId}/faxlines/{faxlineId}/numbers | List all phone numbers routed to fax line
[**GetUserFaxlines**](FaxlinesApi.md#GetUserFaxlines) | **Get** /{userId}/faxlines | List all fax lines



## ChangeFaxlineAlias

> ChangeFaxlineAlias(ctx, userId, faxlineId).Body(body).Execute()

Update alias for an existing fax line

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
    faxlineId := "f0" // string | The unique fax line identifier
    body := *openapiclient.NewSetAliasRequest() // SetAliasRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaxlinesApi.ChangeFaxlineAlias(context.Background(), userId, faxlineId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaxlinesApi.ChangeFaxlineAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**faxlineId** | **string** | The unique fax line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeFaxlineAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**SetAliasRequest**](SetAliasRequest.md) |  | 

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


## ChangeFaxlineCallerIdSettings

> ChangeFaxlineCallerIdSettings(ctx, userId, faxlineId).Body(body).Execute()

Update the caller ID for an existing fax line

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
    faxlineId := "f0" // string | The unique fax line identifier
    body := *openapiclient.NewFaxlineCallerIdRequest() // FaxlineCallerIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaxlinesApi.ChangeFaxlineCallerIdSettings(context.Background(), userId, faxlineId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaxlinesApi.ChangeFaxlineCallerIdSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**faxlineId** | **string** | The unique fax line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeFaxlineCallerIdSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**FaxlineCallerIdRequest**](FaxlineCallerIdRequest.md) |  | 

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


## ChangeTagline

> ChangeTagline(ctx, userId, faxlineId).Body(body).Execute()

Update tagline of an existing fax line

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
    faxlineId := "f0" // string | The unique fax line identifier
    body := *openapiclient.NewFaxlineTaglineRequest() // FaxlineTaglineRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaxlinesApi.ChangeTagline(context.Background(), userId, faxlineId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaxlinesApi.ChangeTagline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**faxlineId** | **string** | The unique fax line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeTaglineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**FaxlineTaglineRequest**](FaxlineTaglineRequest.md) |  | 

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


## CreateFaxline

> FaxlineResponse CreateFaxline(ctx, userId).Execute()

Create a new fax line

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaxlinesApi.CreateFaxline(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaxlinesApi.CreateFaxline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFaxline`: FaxlineResponse
    fmt.Fprintf(os.Stdout, "Response from `FaxlinesApi.CreateFaxline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFaxlineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FaxlineResponse**](FaxlineResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFaxline

> DeleteFaxline(ctx, userId, faxlineId).Execute()

Delete an existing fax line

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
    faxlineId := "f0" // string | The unique fax line identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaxlinesApi.DeleteFaxline(context.Background(), userId, faxlineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaxlinesApi.DeleteFaxline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**faxlineId** | **string** | The unique fax line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFaxlineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetFaxlineCallerId

> FaxlineCallerIdResponse GetFaxlineCallerId(ctx, userId, faxlineId).Execute()

Get the caller ID of a specific fax line

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
    faxlineId := "f0" // string | The unique fax line identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaxlinesApi.GetFaxlineCallerId(context.Background(), userId, faxlineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaxlinesApi.GetFaxlineCallerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFaxlineCallerId`: FaxlineCallerIdResponse
    fmt.Fprintf(os.Stdout, "Response from `FaxlinesApi.GetFaxlineCallerId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**faxlineId** | **string** | The unique fax line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFaxlineCallerIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FaxlineCallerIdResponse**](FaxlineCallerIdResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFaxlineNumbers

> EndpointNumbersResponse GetFaxlineNumbers(ctx, userId, faxlineId).Execute()

List all phone numbers routed to fax line

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
    faxlineId := "f0" // string | The unique fax line identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaxlinesApi.GetFaxlineNumbers(context.Background(), userId, faxlineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaxlinesApi.GetFaxlineNumbers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFaxlineNumbers`: EndpointNumbersResponse
    fmt.Fprintf(os.Stdout, "Response from `FaxlinesApi.GetFaxlineNumbers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**faxlineId** | **string** | The unique fax line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFaxlineNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EndpointNumbersResponse**](EndpointNumbersResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserFaxlines

> FaxlinesResponse GetUserFaxlines(ctx, userId).Execute()

List all fax lines

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FaxlinesApi.GetUserFaxlines(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FaxlinesApi.GetUserFaxlines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserFaxlines`: FaxlinesResponse
    fmt.Fprintf(os.Stdout, "Response from `FaxlinesApi.GetUserFaxlines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserFaxlinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FaxlinesResponse**](FaxlinesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

