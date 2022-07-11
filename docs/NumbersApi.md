# \NumbersApi

All URIs are relative to *https://api.sipgate.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddQuickDialNumber**](NumbersApi.md#AddQuickDialNumber) | **Post** /numbers/quickdial | Add a new quick dial number
[**ChangeNumberSettings**](NumbersApi.md#ChangeNumberSettings) | **Put** /numbers/{numberId} | Change phone number settings, e.g. route to endpoint
[**ChangeQuickDialNumber**](NumbersApi.md#ChangeQuickDialNumber) | **Put** /numbers/quickdial/{numberId} | Update an existing quick dial number
[**DeleteQuickDialNumber**](NumbersApi.md#DeleteQuickDialNumber) | **Delete** /numbers/quickdial/{numberId} | Delete an existing quick dial number
[**GetNumbers**](NumbersApi.md#GetNumbers) | **Get** /numbers | List all phone numbers
[**GetUserNumbers**](NumbersApi.md#GetUserNumbers) | **Get** /{userId}/numbers | List all user phone numbers of a specific user
[**ValidateQuickDialNumber**](NumbersApi.md#ValidateQuickDialNumber) | **Get** /numbers/quickdial/validation/{quickDialNumber} | Check if quick dial number is already taken



## AddQuickDialNumber

> AddQuickDialNumber(ctx).Body(body).Execute()

Add a new quick dial number

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
    body := *openapiclient.NewAddDirectDialRequest("w0") // AddDirectDialRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NumbersApi.AddQuickDialNumber(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NumbersApi.AddQuickDialNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddQuickDialNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AddDirectDialRequest**](AddDirectDialRequest.md) |  | 

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


## ChangeNumberSettings

> ChangeNumberSettings(ctx, numberId).Body(body).Execute()

Change phone number settings, e.g. route to endpoint

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
    numberId := "123" // string | The unique phone number identifier
    body := *openapiclient.NewChangeNumberSettingsRequest() // ChangeNumberSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NumbersApi.ChangeNumberSettings(context.Background(), numberId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NumbersApi.ChangeNumberSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**numberId** | **string** | The unique phone number identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeNumberSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ChangeNumberSettingsRequest**](ChangeNumberSettingsRequest.md) |  | 

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


## ChangeQuickDialNumber

> ChangeQuickDialNumber(ctx, numberId).Body(body).Execute()

Update an existing quick dial number

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
    numberId := "numberId_example" // string | 
    body := *openapiclient.NewChangeDirectDialRequest("w0") // ChangeDirectDialRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NumbersApi.ChangeQuickDialNumber(context.Background(), numberId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NumbersApi.ChangeQuickDialNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**numberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeQuickDialNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ChangeDirectDialRequest**](ChangeDirectDialRequest.md) |  | 

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


## DeleteQuickDialNumber

> DeleteQuickDialNumber(ctx, numberId).Execute()

Delete an existing quick dial number

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
    numberId := "abc123" // string | The unique quick dial number identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NumbersApi.DeleteQuickDialNumber(context.Background(), numberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NumbersApi.DeleteQuickDialNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**numberId** | **string** | The unique quick dial number identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQuickDialNumberRequest struct via the builder pattern


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


## GetNumbers

> NumbersResponse GetNumbers(ctx).Offset(offset).Limit(limit).Execute()

List all phone numbers



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
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NumbersApi.GetNumbers(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NumbersApi.GetNumbers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNumbers`: NumbersResponse
    fmt.Fprintf(os.Stdout, "Response from `NumbersApi.GetNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

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


## GetUserNumbers

> NumbersResponse GetUserNumbers(ctx, userId).Execute()

List all user phone numbers of a specific user



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
    resp, r, err := apiClient.NumbersApi.GetUserNumbers(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NumbersApi.GetUserNumbers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserNumbers`: NumbersResponse
    fmt.Fprintf(os.Stdout, "Response from `NumbersApi.GetUserNumbers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserNumbersRequest struct via the builder pattern


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


## ValidateQuickDialNumber

> ValidateQuickDialNumber(ctx, quickDialNumber).Execute()

Check if quick dial number is already taken

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
    quickDialNumber := "42" // string | The quick dial number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NumbersApi.ValidateQuickDialNumber(context.Background(), quickDialNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NumbersApi.ValidateQuickDialNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quickDialNumber** | **string** | The quick dial number | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateQuickDialNumberRequest struct via the builder pattern


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

