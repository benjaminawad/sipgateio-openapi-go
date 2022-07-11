# \PhonelinesApi

All URIs are relative to *https://api.sipgate.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDeviceToPhoneline**](PhonelinesApi.md#AddDeviceToPhoneline) | **Post** /{userId}/phonelines/{phonelineId}/devices | Add device to phone line
[**AddGreeting**](PhonelinesApi.md#AddGreeting) | **Post** /{userId}/phonelines/{phonelineId}/voicemails/{voicemailId}/greetings | Create a new greeting for a specific phone line voicemail
[**AddParallelForwarding**](PhonelinesApi.md#AddParallelForwarding) | **Post** /{userId}/phonelines/{phonelineId}/parallelforwardings | Create parallel forwarding for an existing phone line
[**ChangeParallelForwardingSettings**](PhonelinesApi.md#ChangeParallelForwardingSettings) | **Put** /{userId}/phonelines/{phonelineId}/parallelforwardings/{parallelForwardingId} | Update an existing parallel forwarding
[**ChangePhonelineVoicemailGreeting**](PhonelinesApi.md#ChangePhonelineVoicemailGreeting) | **Put** /{userId}/phonelines/{phonelineId}/voicemails/{voicemailId}/greetings/{greetingId} | Update an existing phone line voicemail greeting
[**ChangePhonelineVoicemailTranscriptionSetting**](PhonelinesApi.md#ChangePhonelineVoicemailTranscriptionSetting) | **Put** /{userId}/phonelines/{phonelineId}/voicemails/{voicemailId}/transcriptions | Update an existing phone line voicemail transcription setting
[**ChangeVoicemailSettings**](PhonelinesApi.md#ChangeVoicemailSettings) | **Put** /{userId}/phonelines/{phonelineId}/voicemails/{voicemailId} | Update voicemail settings for a specific phone line
[**CreatePhoneline**](PhonelinesApi.md#CreatePhoneline) | **Post** /{userId}/phonelines | Create a new phone line
[**DeleteParallelForwardingSettings**](PhonelinesApi.md#DeleteParallelForwardingSettings) | **Delete** /{userId}/phonelines/{phonelineId}/parallelforwardings/{parallelForwardingId} | Delete parallel forwarding
[**DeletePhoneline**](PhonelinesApi.md#DeletePhoneline) | **Delete** /{userId}/phonelines/{phonelineId} | Delete an existing phone line
[**DeletePhonlineVoicemailGreeting**](PhonelinesApi.md#DeletePhonlineVoicemailGreeting) | **Delete** /{userId}/phonelines/{phonelineId}/voicemails/{voicemailId}/greetings/{greetingId} | Delete an existing phone line voicemail greeting
[**GetBlockAnonymousSetting**](PhonelinesApi.md#GetBlockAnonymousSetting) | **Get** /{userId}/phonelines/{phonelineId}/blockanonymous | Get the block anonymous setting of a specific phone line
[**GetDevicesForPhoneline**](PhonelinesApi.md#GetDevicesForPhoneline) | **Get** /{userId}/phonelines/{phonelineId}/devices | List all devices of a specific phone line
[**GetPhoneline**](PhonelinesApi.md#GetPhoneline) | **Get** /{userId}/phonelines/{phonelineId} | Get a single phoneline
[**GetPhonelineForwardings**](PhonelinesApi.md#GetPhonelineForwardings) | **Get** /{userId}/phonelines/{phonelineId}/forwardings | List all forwarding settings of a specific phone line
[**GetPhonelineNumbers**](PhonelinesApi.md#GetPhonelineNumbers) | **Get** /{userId}/phonelines/{phonelineId}/numbers | List all phone numbers routed to phone line
[**GetPhonelineParallelForwardings**](PhonelinesApi.md#GetPhonelineParallelForwardings) | **Get** /{userId}/phonelines/{phonelineId}/parallelforwardings | List all parallel forwardings of a specific phone line
[**GetPhonelineVoicemailGreetings**](PhonelinesApi.md#GetPhonelineVoicemailGreetings) | **Get** /{userId}/phonelines/{phonelineId}/voicemails/{voicemailId}/greetings | List all greetings for a specific phone line voicemail
[**GetPhonelineVoicemails**](PhonelinesApi.md#GetPhonelineVoicemails) | **Get** /{userId}/phonelines/{phonelineId}/voicemails | List all voicemails of a specific phone line
[**GetPhonelines**](PhonelinesApi.md#GetPhonelines) | **Get** /{userId}/phonelines | List all phone lines
[**GetSipgateIoLogs**](PhonelinesApi.md#GetSipgateIoLogs) | **Get** /{userId}/phonelines/{phonelineId}/sipgateio/log | List sipgate.io debug log for a specific phone line
[**GetSipgateIoSettings**](PhonelinesApi.md#GetSipgateIoSettings) | **Get** /{userId}/phonelines/{phonelineId}/sipgateio | List sipgate.io settings of a specific phone line
[**RemoveDeviceFromPhoneline**](PhonelinesApi.md#RemoveDeviceFromPhoneline) | **Delete** /{userId}/phonelines/{phonelineId}/devices/{deviceId} | Remove device from phone line
[**SetBlockAnonymousSetting**](PhonelinesApi.md#SetBlockAnonymousSetting) | **Put** /{userId}/phonelines/{phonelineId}/blockanonymous | Update the block anonymous setting for an existing phone line
[**SetPhonelineAlias**](PhonelinesApi.md#SetPhonelineAlias) | **Put** /{userId}/phonelines/{phonelineId} | Update alias for an existing phone line
[**SetPhonelineForwardingSettings**](PhonelinesApi.md#SetPhonelineForwardingSettings) | **Put** /{userId}/phonelines/{phonelineId}/forwardings | Update forwarding settings for an existing phone line
[**SetSipgateIoSettings**](PhonelinesApi.md#SetSipgateIoSettings) | **Put** /{userId}/phonelines/{phonelineId}/sipgateio | Update sipgate.io settings for an existing phone line



## AddDeviceToPhoneline

> AddDeviceToPhoneline(ctx, userId, phonelineId).Body(body).Execute()

Add device to phone line

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
    phonelineId := "p0" // string | The unique phone line identifier
    body := *openapiclient.NewRouteDeviceRequest("e0") // RouteDeviceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.AddDeviceToPhoneline(context.Background(), userId, phonelineId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.AddDeviceToPhoneline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDeviceToPhonelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**RouteDeviceRequest**](RouteDeviceRequest.md) |  | 

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


## AddGreeting

> GreetingsResponse AddGreeting(ctx, userId, phonelineId, voicemailId).Body(body).Execute()

Create a new greeting for a specific phone line voicemail

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
    phonelineId := "p0" // string | The unique phone line identifier
    voicemailId := "v0" // string | The unique voicemail identifier
    body := *openapiclient.NewAddGreetingRequest() // AddGreetingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.AddGreeting(context.Background(), userId, phonelineId, voicemailId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.AddGreeting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddGreeting`: GreetingsResponse
    fmt.Fprintf(os.Stdout, "Response from `PhonelinesApi.AddGreeting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 
**voicemailId** | **string** | The unique voicemail identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGreetingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**AddGreetingRequest**](AddGreetingRequest.md) |  | 

### Return type

[**GreetingsResponse**](GreetingsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddParallelForwarding

> AddParallelForwarding(ctx, userId, phonelineId).Body(body).Execute()

Create parallel forwarding for an existing phone line

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
    phonelineId := "p0" // string | The unique phone line identifier
    body := *openapiclient.NewSetParallelForwardingRequest() // SetParallelForwardingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.AddParallelForwarding(context.Background(), userId, phonelineId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.AddParallelForwarding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddParallelForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**SetParallelForwardingRequest**](SetParallelForwardingRequest.md) |  | 

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


## ChangeParallelForwardingSettings

> ChangeParallelForwardingSettings(ctx, userId, phonelineId, parallelForwardingId).Body(body).Execute()

Update an existing parallel forwarding

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
    phonelineId := "p0" // string | The unique phone line identifier
    parallelForwardingId := "x0" // string | The unique external device identifier
    body := *openapiclient.NewSetParallelForwardingRequest() // SetParallelForwardingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.ChangeParallelForwardingSettings(context.Background(), userId, phonelineId, parallelForwardingId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.ChangeParallelForwardingSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 
**parallelForwardingId** | **string** | The unique external device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeParallelForwardingSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**SetParallelForwardingRequest**](SetParallelForwardingRequest.md) |  | 

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


## ChangePhonelineVoicemailGreeting

> ChangePhonelineVoicemailGreeting(ctx, userId, phonelineId, voicemailId, greetingId).Body(body).Execute()

Update an existing phone line voicemail greeting

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
    phonelineId := "p0" // string | The unique phone line identifier
    voicemailId := "v0" // string | The unique voicemail identifier
    greetingId := "abc123" // string | The unique greeting identifier
    body := *openapiclient.NewChangeVoicemailGreetingRequest() // ChangeVoicemailGreetingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.ChangePhonelineVoicemailGreeting(context.Background(), userId, phonelineId, voicemailId, greetingId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.ChangePhonelineVoicemailGreeting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 
**voicemailId** | **string** | The unique voicemail identifier | 
**greetingId** | **string** | The unique greeting identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangePhonelineVoicemailGreetingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**ChangeVoicemailGreetingRequest**](ChangeVoicemailGreetingRequest.md) |  | 

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


## ChangePhonelineVoicemailTranscriptionSetting

> ChangePhonelineVoicemailTranscriptionSetting(ctx, userId, phonelineId, voicemailId).Body(body).Execute()

Update an existing phone line voicemail transcription setting

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
    phonelineId := "p0" // string | The unique phone line identifier
    voicemailId := "v0" // string | The unique voicemail identifier
    body := *openapiclient.NewChangeVoicemailTranscriptionRequest() // ChangeVoicemailTranscriptionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.ChangePhonelineVoicemailTranscriptionSetting(context.Background(), userId, phonelineId, voicemailId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.ChangePhonelineVoicemailTranscriptionSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 
**voicemailId** | **string** | The unique voicemail identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangePhonelineVoicemailTranscriptionSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**ChangeVoicemailTranscriptionRequest**](ChangeVoicemailTranscriptionRequest.md) |  | 

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


## ChangeVoicemailSettings

> ChangeVoicemailSettings(ctx, userId, phonelineId, voicemailId).Body(body).Execute()

Update voicemail settings for a specific phone line

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
    phonelineId := "p0" // string | The unique phone line identifier
    voicemailId := "v0" // string | The unique voicemail identifier
    body := *openapiclient.NewChangeVoicemailRequest(false, false) // ChangeVoicemailRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.ChangeVoicemailSettings(context.Background(), userId, phonelineId, voicemailId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.ChangeVoicemailSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 
**voicemailId** | **string** | The unique voicemail identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeVoicemailSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**ChangeVoicemailRequest**](ChangeVoicemailRequest.md) |  | 

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


## CreatePhoneline

> SimplePhonelineResponse CreatePhoneline(ctx, userId).Execute()

Create a new phone line

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
    resp, r, err := apiClient.PhonelinesApi.CreatePhoneline(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.CreatePhoneline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePhoneline`: SimplePhonelineResponse
    fmt.Fprintf(os.Stdout, "Response from `PhonelinesApi.CreatePhoneline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePhonelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimplePhonelineResponse**](SimplePhonelineResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteParallelForwardingSettings

> DeleteParallelForwardingSettings(ctx, userId, phonelineId, parallelForwardingId).Execute()

Delete parallel forwarding

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
    phonelineId := "p0" // string | The unique phone line identifier
    parallelForwardingId := "x0" // string | The unique external device identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.DeleteParallelForwardingSettings(context.Background(), userId, phonelineId, parallelForwardingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.DeleteParallelForwardingSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 
**parallelForwardingId** | **string** | The unique external device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteParallelForwardingSettingsRequest struct via the builder pattern


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


## DeletePhoneline

> DeletePhoneline(ctx, userId, phonelineId).Execute()

Delete an existing phone line

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
    phonelineId := "p0" // string | The unique phone line identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.DeletePhoneline(context.Background(), userId, phonelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.DeletePhoneline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePhonelineRequest struct via the builder pattern


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


## DeletePhonlineVoicemailGreeting

> DeletePhonlineVoicemailGreeting(ctx, userId, phonelineId, voicemailId, greetingId).Execute()

Delete an existing phone line voicemail greeting

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
    phonelineId := "p0" // string | The unique phone line identifier
    voicemailId := "v0" // string | The unique voicemail identifier
    greetingId := "abc123" // string | The unique greeting identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.DeletePhonlineVoicemailGreeting(context.Background(), userId, phonelineId, voicemailId, greetingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.DeletePhonlineVoicemailGreeting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 
**voicemailId** | **string** | The unique voicemail identifier | 
**greetingId** | **string** | The unique greeting identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePhonlineVoicemailGreetingRequest struct via the builder pattern


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


## GetBlockAnonymousSetting

> PhonelineBlockAnonymousResponse GetBlockAnonymousSetting(ctx, userId, phonelineId).Execute()

Get the block anonymous setting of a specific phone line

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
    phonelineId := "p0" // string | The unique phone line identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.GetBlockAnonymousSetting(context.Background(), userId, phonelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.GetBlockAnonymousSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockAnonymousSetting`: PhonelineBlockAnonymousResponse
    fmt.Fprintf(os.Stdout, "Response from `PhonelinesApi.GetBlockAnonymousSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockAnonymousSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PhonelineBlockAnonymousResponse**](PhonelineBlockAnonymousResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevicesForPhoneline

> DevicesResponse GetDevicesForPhoneline(ctx, userId, phonelineId).Execute()

List all devices of a specific phone line

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
    phonelineId := "p0" // string | The unique phone line identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.GetDevicesForPhoneline(context.Background(), userId, phonelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.GetDevicesForPhoneline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevicesForPhoneline`: DevicesResponse
    fmt.Fprintf(os.Stdout, "Response from `PhonelinesApi.GetDevicesForPhoneline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesForPhonelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DevicesResponse**](DevicesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPhoneline

> PhonelineResponse GetPhoneline(ctx, userId, phonelineId).Execute()

Get a single phoneline

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
    phonelineId := "p0" // string | The unique phone line identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.GetPhoneline(context.Background(), userId, phonelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.GetPhoneline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPhoneline`: PhonelineResponse
    fmt.Fprintf(os.Stdout, "Response from `PhonelinesApi.GetPhoneline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPhonelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PhonelineResponse**](PhonelineResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPhonelineForwardings

> ForwardingsResponse GetPhonelineForwardings(ctx, userId, phonelineId).Execute()

List all forwarding settings of a specific phone line

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
    phonelineId := "p0" // string | The unique phone line identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.GetPhonelineForwardings(context.Background(), userId, phonelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.GetPhonelineForwardings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPhonelineForwardings`: ForwardingsResponse
    fmt.Fprintf(os.Stdout, "Response from `PhonelinesApi.GetPhonelineForwardings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPhonelineForwardingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ForwardingsResponse**](ForwardingsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPhonelineNumbers

> EndpointNumbersResponse GetPhonelineNumbers(ctx, userId, phonelineId).Execute()

List all phone numbers routed to phone line

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
    phonelineId := "p0" // string | The unique phone line identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.GetPhonelineNumbers(context.Background(), userId, phonelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.GetPhonelineNumbers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPhonelineNumbers`: EndpointNumbersResponse
    fmt.Fprintf(os.Stdout, "Response from `PhonelinesApi.GetPhonelineNumbers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPhonelineNumbersRequest struct via the builder pattern


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


## GetPhonelineParallelForwardings

> ParallelForwardingsResponse GetPhonelineParallelForwardings(ctx, userId, phonelineId).Execute()

List all parallel forwardings of a specific phone line

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
    phonelineId := "p0" // string | The unique phone line identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.GetPhonelineParallelForwardings(context.Background(), userId, phonelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.GetPhonelineParallelForwardings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPhonelineParallelForwardings`: ParallelForwardingsResponse
    fmt.Fprintf(os.Stdout, "Response from `PhonelinesApi.GetPhonelineParallelForwardings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPhonelineParallelForwardingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ParallelForwardingsResponse**](ParallelForwardingsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPhonelineVoicemailGreetings

> GreetingsResponse GetPhonelineVoicemailGreetings(ctx, userId, phonelineId, voicemailId).Execute()

List all greetings for a specific phone line voicemail

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
    phonelineId := "p0" // string | The unique phone line identifier
    voicemailId := "v0" // string | The unique voicemail identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.GetPhonelineVoicemailGreetings(context.Background(), userId, phonelineId, voicemailId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.GetPhonelineVoicemailGreetings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPhonelineVoicemailGreetings`: GreetingsResponse
    fmt.Fprintf(os.Stdout, "Response from `PhonelinesApi.GetPhonelineVoicemailGreetings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 
**voicemailId** | **string** | The unique voicemail identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPhonelineVoicemailGreetingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GreetingsResponse**](GreetingsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPhonelineVoicemails

> PhonelineVoicemailsResponse GetPhonelineVoicemails(ctx, userId, phonelineId).Execute()

List all voicemails of a specific phone line

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
    phonelineId := "p0" // string | The unique phone line identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.GetPhonelineVoicemails(context.Background(), userId, phonelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.GetPhonelineVoicemails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPhonelineVoicemails`: PhonelineVoicemailsResponse
    fmt.Fprintf(os.Stdout, "Response from `PhonelinesApi.GetPhonelineVoicemails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPhonelineVoicemailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PhonelineVoicemailsResponse**](PhonelineVoicemailsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPhonelines

> SimplePhonelinesResponse GetPhonelines(ctx, userId).Execute()

List all phone lines

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
    resp, r, err := apiClient.PhonelinesApi.GetPhonelines(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.GetPhonelines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPhonelines`: SimplePhonelinesResponse
    fmt.Fprintf(os.Stdout, "Response from `PhonelinesApi.GetPhonelines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPhonelinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimplePhonelinesResponse**](SimplePhonelinesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSipgateIoLogs

> PhonelineSipgateIoLogsResponse GetSipgateIoLogs(ctx, userId, phonelineId).Execute()

List sipgate.io debug log for a specific phone line

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
    phonelineId := "p0" // string | The unique phone line identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.GetSipgateIoLogs(context.Background(), userId, phonelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.GetSipgateIoLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSipgateIoLogs`: PhonelineSipgateIoLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `PhonelinesApi.GetSipgateIoLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSipgateIoLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PhonelineSipgateIoLogsResponse**](PhonelineSipgateIoLogsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSipgateIoSettings

> PhonelineSipgateIoResponse GetSipgateIoSettings(ctx, userId, phonelineId).Execute()

List sipgate.io settings of a specific phone line

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
    phonelineId := "p0" // string | The unique phone line identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.GetSipgateIoSettings(context.Background(), userId, phonelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.GetSipgateIoSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSipgateIoSettings`: PhonelineSipgateIoResponse
    fmt.Fprintf(os.Stdout, "Response from `PhonelinesApi.GetSipgateIoSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSipgateIoSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PhonelineSipgateIoResponse**](PhonelineSipgateIoResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDeviceFromPhoneline

> RemoveDeviceFromPhoneline(ctx, userId, phonelineId, deviceId).Execute()

Remove device from phone line

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
    phonelineId := "p0" // string | The unique phone line identifier
    deviceId := "e0" // string | The unique device identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.RemoveDeviceFromPhoneline(context.Background(), userId, phonelineId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.RemoveDeviceFromPhoneline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 
**deviceId** | **string** | The unique device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDeviceFromPhonelineRequest struct via the builder pattern


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


## SetBlockAnonymousSetting

> PhonelineBlockAnonymousResponse SetBlockAnonymousSetting(ctx, userId, phonelineId).Body(body).Execute()

Update the block anonymous setting for an existing phone line

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
    phonelineId := "p0" // string | The unique phone line identifier
    body := *openapiclient.NewSetBlockAnonymousSettingsRequest() // SetBlockAnonymousSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.SetBlockAnonymousSetting(context.Background(), userId, phonelineId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.SetBlockAnonymousSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetBlockAnonymousSetting`: PhonelineBlockAnonymousResponse
    fmt.Fprintf(os.Stdout, "Response from `PhonelinesApi.SetBlockAnonymousSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetBlockAnonymousSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**SetBlockAnonymousSettingsRequest**](SetBlockAnonymousSettingsRequest.md) |  | 

### Return type

[**PhonelineBlockAnonymousResponse**](PhonelineBlockAnonymousResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPhonelineAlias

> SetPhonelineAlias(ctx, userId, phonelineId).Body(body).Execute()

Update alias for an existing phone line

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
    phonelineId := "p0" // string | The unique phone line identifier
    body := *openapiclient.NewSetAliasRequest() // SetAliasRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.SetPhonelineAlias(context.Background(), userId, phonelineId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.SetPhonelineAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPhonelineAliasRequest struct via the builder pattern


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


## SetPhonelineForwardingSettings

> SetPhonelineForwardingSettings(ctx, userId, phonelineId).Body(body).Execute()

Update forwarding settings for an existing phone line

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
    phonelineId := "p0" // string | The unique phone line identifier
    body := *openapiclient.NewSetForwardingsRequest() // SetForwardingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.SetPhonelineForwardingSettings(context.Background(), userId, phonelineId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.SetPhonelineForwardingSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPhonelineForwardingSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**SetForwardingsRequest**](SetForwardingsRequest.md) |  | 

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


## SetSipgateIoSettings

> SetSipgateIoSettings(ctx, userId, phonelineId).Body(body).Execute()

Update sipgate.io settings for an existing phone line

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
    phonelineId := "p0" // string | The unique phone line identifier
    body := *openapiclient.NewSetSipgateIoPhonelineSettingsRequest("https://io.sipgate.beer/my/incoming/url", "https://io.sipgate.rehab/my/outgoing/url") // SetSipgateIoPhonelineSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhonelinesApi.SetSipgateIoSettings(context.Background(), userId, phonelineId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhonelinesApi.SetSipgateIoSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**phonelineId** | **string** | The unique phone line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSipgateIoSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**SetSipgateIoPhonelineSettingsRequest**](SetSipgateIoPhonelineSettingsRequest.md) |  | 

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

