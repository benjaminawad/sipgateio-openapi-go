# \AutorecordingApi

All URIs are relative to *https://api.sipgate.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAutorecordingGreeting**](AutorecordingApi.md#AddAutorecordingGreeting) | **Post** /autorecordings/greetings | save a greeting to be used for automated call recordings
[**DeleteGreeting**](AutorecordingApi.md#DeleteGreeting) | **Delete** /autorecordings/greetings/{greetingId} | delete the greeting used for automated call recordings
[**GetGreeting**](AutorecordingApi.md#GetGreeting) | **Get** /autorecordings/greetings | Fetch the current automated call recording greeting
[**GetSetting**](AutorecordingApi.md#GetSetting) | **Get** /autorecordings/{extension}/settings | Get the recording setting of a device
[**UpdateSetting**](AutorecordingApi.md#UpdateSetting) | **Put** /autorecordings/{extension}/settings | Set the recording setting of a device



## AddAutorecordingGreeting

> AddAutorecordingGreeting(ctx).Body(body).Execute()

save a greeting to be used for automated call recordings

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
    body := *openapiclient.NewAddGreetingRequest() // AddGreetingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutorecordingApi.AddAutorecordingGreeting(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutorecordingApi.AddAutorecordingGreeting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAutorecordingGreetingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AddGreetingRequest**](AddGreetingRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGreeting

> DeleteGreeting(ctx, greetingId).Execute()

delete the greeting used for automated call recordings

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
    greetingId := "abc123" // string | The unique greeting identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutorecordingApi.DeleteGreeting(context.Background(), greetingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutorecordingApi.DeleteGreeting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**greetingId** | **string** | The unique greeting identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGreetingRequest struct via the builder pattern


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


## GetGreeting

> GreetingResponse GetGreeting(ctx).Execute()

Fetch the current automated call recording greeting

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
    resp, r, err := apiClient.AutorecordingApi.GetGreeting(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutorecordingApi.GetGreeting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGreeting`: GreetingResponse
    fmt.Fprintf(os.Stdout, "Response from `AutorecordingApi.GetGreeting`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGreetingRequest struct via the builder pattern


### Return type

[**GreetingResponse**](GreetingResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSetting

> GetRecordingSettingResponse GetSetting(ctx, extension).Execute()

Get the recording setting of a device

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
    extension := "extension_example" // string | unique device identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutorecordingApi.GetSetting(context.Background(), extension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutorecordingApi.GetSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSetting`: GetRecordingSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `AutorecordingApi.GetSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extension** | **string** | unique device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRecordingSettingResponse**](GetRecordingSettingResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSetting

> GetRecordingSettingResponse UpdateSetting(ctx, extension).Body(body).Execute()

Set the recording setting of a device

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
    extension := "extension_example" // string | unique device identifier
    body := *openapiclient.NewRecordingSettingsIncomingRequest() // RecordingSettingsIncomingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutorecordingApi.UpdateSetting(context.Background(), extension).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutorecordingApi.UpdateSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSetting`: GetRecordingSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `AutorecordingApi.UpdateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extension** | **string** | unique device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RecordingSettingsIncomingRequest**](RecordingSettingsIncomingRequest.md) |  | 

### Return type

[**GetRecordingSettingResponse**](GetRecordingSettingResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

