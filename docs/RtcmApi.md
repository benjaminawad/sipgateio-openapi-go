# \RtcmApi

All URIs are relative to *https://api.sipgate.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCalls**](RtcmApi.md#GetCalls) | **Get** /calls | List all calls currently established calls. The response does not contain ringing calls or currently running voicemail recordings. 
[**HangupCall**](RtcmApi.md#HangupCall) | **Delete** /calls/{callId} | Hang up call
[**SendCallDtmf**](RtcmApi.md#SendCallDtmf) | **Post** /calls/{callId}/dtmf | Send DTMF tone sequences to all participants
[**SetCallHold**](RtcmApi.md#SetCallHold) | **Put** /calls/{callId}/hold | Hold/Resume all participants
[**SetCallMuted**](RtcmApi.md#SetCallMuted) | **Put** /calls/{callId}/muted | Mute/Unmute yourself
[**SetCallRecording**](RtcmApi.md#SetCallRecording) | **Put** /calls/{callId}/recording | Start or stop a call recording
[**StartCallAnnouncement**](RtcmApi.md#StartCallAnnouncement) | **Post** /calls/{callId}/announcements | Start playing a new announcement to all participants
[**TransferCall**](RtcmApi.md#TransferCall) | **Post** /calls/{callId}/transfer | Transfer call



## GetCalls

> Calls GetCalls(ctx).Execute()

List all calls currently established calls. The response does not contain ringing calls or currently running voicemail recordings. 

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
    resp, r, err := apiClient.RtcmApi.GetCalls(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RtcmApi.GetCalls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCalls`: Calls
    fmt.Fprintf(os.Stdout, "Response from `RtcmApi.GetCalls`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCallsRequest struct via the builder pattern


### Return type

[**Calls**](Calls.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HangupCall

> HangupCall(ctx, callId).Execute()

Hang up call

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
    callId := "callId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RtcmApi.HangupCall(context.Background(), callId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RtcmApi.HangupCall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHangupCallRequest struct via the builder pattern


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


## SendCallDtmf

> SendCallDtmf(ctx, callId).Body(body).Execute()

Send DTMF tone sequences to all participants

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
    callId := "callId_example" // string | 
    body := *openapiclient.NewDtmfRequest("123456") // DtmfRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RtcmApi.SendCallDtmf(context.Background(), callId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RtcmApi.SendCallDtmf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendCallDtmfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DtmfRequest**](DtmfRequest.md) |  | 

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


## SetCallHold

> SetCallHold(ctx, callId).Body(body).Execute()

Hold/Resume all participants

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
    callId := "callId_example" // string | 
    body := *openapiclient.NewBooleanRequest(true) // BooleanRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RtcmApi.SetCallHold(context.Background(), callId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RtcmApi.SetCallHold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCallHoldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**BooleanRequest**](BooleanRequest.md) |  | 

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


## SetCallMuted

> SetCallMuted(ctx, callId).Body(body).Execute()

Mute/Unmute yourself

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
    callId := "callId_example" // string | 
    body := *openapiclient.NewBooleanRequest(true) // BooleanRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RtcmApi.SetCallMuted(context.Background(), callId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RtcmApi.SetCallMuted``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCallMutedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**BooleanRequest**](BooleanRequest.md) |  | 

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


## SetCallRecording

> SetCallRecording(ctx, callId).Body(body).Execute()

Start or stop a call recording

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
    callId := "callId_example" // string | 
    body := *openapiclient.NewRecordingRequest(true, true) // RecordingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RtcmApi.SetCallRecording(context.Background(), callId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RtcmApi.SetCallRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCallRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RecordingRequest**](RecordingRequest.md) |  | 

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


## StartCallAnnouncement

> StartCallAnnouncement(ctx, callId).Body(body).Execute()

Start playing a new announcement to all participants



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
    callId := "callId_example" // string | 
    body := *openapiclient.NewAnnouncementRequest("https://static.sipgate.com/examples/wav/example.wav") // AnnouncementRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RtcmApi.StartCallAnnouncement(context.Background(), callId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RtcmApi.StartCallAnnouncement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartCallAnnouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AnnouncementRequest**](AnnouncementRequest.md) |  | 

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


## TransferCall

> TransferCall(ctx, callId).Body(body).Execute()

Transfer call

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
    callId := "callId_example" // string | 
    body := *openapiclient.NewTransferRequest(false, "+4915799912345") // TransferRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RtcmApi.TransferCall(context.Background(), callId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RtcmApi.TransferCall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TransferRequest**](TransferRequest.md) |  | 

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

