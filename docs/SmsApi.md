# \SmsApi

All URIs are relative to *https://api.sipgate.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCallerId**](SmsApi.md#AddCallerId) | **Post** /{userId}/sms/{smsId}/callerids | Create a new caller ID for an existing short message service
[**DeleteSmsCallerId**](SmsApi.md#DeleteSmsCallerId) | **Delete** /{userId}/sms/{smsId}/callerids/{callerId} | Delete an existing caller ID for a specific short message service
[**EditCallerId**](SmsApi.md#EditCallerId) | **Put** /{userId}/sms/{smsId}/callerids/{callerId} | Update the caller ID for an existing short message service
[**GetSmsCallerIds**](SmsApi.md#GetSmsCallerIds) | **Get** /{userId}/sms/{smsId}/callerids | List all caller IDs for a specific short message service
[**GetSmsExtensions**](SmsApi.md#GetSmsExtensions) | **Get** /{userId}/sms | List all short message services
[**SetSmsAlias**](SmsApi.md#SetSmsAlias) | **Put** /{userId}/sms/{smsId} | Update alias for an existing short message service
[**VerifyCallerId**](SmsApi.md#VerifyCallerId) | **Post** /{userId}/sms/{smsId}/callerids/{callerId}/verification | Verify an existing caller ID for a specific short message service



## AddCallerId

> AddCallerId(ctx, userId, smsId).Body(body).Execute()

Create a new caller ID for an existing short message service

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
    smsId := "s0" // string | The unique short message service identifier
    body := *openapiclient.NewSmsCallerIdRequest() // SmsCallerIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmsApi.AddCallerId(context.Background(), userId, smsId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmsApi.AddCallerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**smsId** | **string** | The unique short message service identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCallerIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**SmsCallerIdRequest**](SmsCallerIdRequest.md) |  | 

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


## DeleteSmsCallerId

> DeleteSmsCallerId(ctx, userId, smsId, callerId).Execute()

Delete an existing caller ID for a specific short message service

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
    smsId := "s0" // string | The unique short message service identifier
    callerId := int32(123) // int32 | The caller ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmsApi.DeleteSmsCallerId(context.Background(), userId, smsId, callerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmsApi.DeleteSmsCallerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**smsId** | **string** | The unique short message service identifier | 
**callerId** | **int32** | The caller ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSmsCallerIdRequest struct via the builder pattern


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


## EditCallerId

> EditCallerId(ctx, userId, smsId, callerId).Body(body).Execute()

Update the caller ID for an existing short message service

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
    smsId := "s0" // string | The unique short message service identifier
    callerId := int32(123) // int32 | The caller ID
    body := *openapiclient.NewSetSmsCallerIdRequest() // SetSmsCallerIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmsApi.EditCallerId(context.Background(), userId, smsId, callerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmsApi.EditCallerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**smsId** | **string** | The unique short message service identifier | 
**callerId** | **int32** | The caller ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditCallerIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**SetSmsCallerIdRequest**](SetSmsCallerIdRequest.md) |  | 

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


## GetSmsCallerIds

> SmsCallerIdsResponse GetSmsCallerIds(ctx, userId, smsId).Execute()

List all caller IDs for a specific short message service

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
    smsId := "s0" // string | The unique short message service identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmsApi.GetSmsCallerIds(context.Background(), userId, smsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmsApi.GetSmsCallerIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmsCallerIds`: SmsCallerIdsResponse
    fmt.Fprintf(os.Stdout, "Response from `SmsApi.GetSmsCallerIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**smsId** | **string** | The unique short message service identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmsCallerIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SmsCallerIdsResponse**](SmsCallerIdsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmsExtensions

> SmsExtensionsResponse GetSmsExtensions(ctx, userId).Execute()

List all short message services

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
    resp, r, err := apiClient.SmsApi.GetSmsExtensions(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmsApi.GetSmsExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmsExtensions`: SmsExtensionsResponse
    fmt.Fprintf(os.Stdout, "Response from `SmsApi.GetSmsExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmsExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SmsExtensionsResponse**](SmsExtensionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSmsAlias

> SetSmsAlias(ctx, userId, smsId).Body(body).Execute()

Update alias for an existing short message service

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
    smsId := "s0" // string | The unique short message service identifier
    body := *openapiclient.NewSetAliasRequest() // SetAliasRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmsApi.SetSmsAlias(context.Background(), userId, smsId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmsApi.SetSmsAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**smsId** | **string** | The unique short message service identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSmsAliasRequest struct via the builder pattern


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


## VerifyCallerId

> VerifyCallerId(ctx, userId, smsId, callerId).Body(body).Execute()

Verify an existing caller ID for a specific short message service

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
    smsId := "s0" // string | The unique short message service identifier
    callerId := int32(123) // int32 | The caller ID
    body := *openapiclient.NewVerifySmsCallerIdRequest() // VerifySmsCallerIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmsApi.VerifyCallerId(context.Background(), userId, smsId, callerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmsApi.VerifyCallerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**smsId** | **string** | The unique short message service identifier | 
**callerId** | **int32** | The caller ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyCallerIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**VerifySmsCallerIdRequest**](VerifySmsCallerIdRequest.md) |  | 

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

