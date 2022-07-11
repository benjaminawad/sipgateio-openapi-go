# \SessionsApi

All URIs are relative to *https://api.sipgate.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NewCall**](SessionsApi.md#NewCall) | **Post** /sessions/calls | Initiate a new call
[**NewIvrRecording**](SessionsApi.md#NewIvrRecording) | **Post** /sessions/ivr/recording | Initiate a new call to record an ivr message
[**NewRecording**](SessionsApi.md#NewRecording) | **Post** /sessions/voicemail/recording | Initiate a new call to record a greeting
[**PlayRecording**](SessionsApi.md#PlayRecording) | **Post** /sessions/voicemail/play | Initiate a new call to playback a voicemail
[**ResendFax**](SessionsApi.md#ResendFax) | **Post** /sessions/fax/resend | Resend a faX
[**SendFax**](SessionsApi.md#SendFax) | **Post** /sessions/fax | Send a fax
[**SendWebSms**](SessionsApi.md#SendWebSms) | **Post** /sessions/sms | Send a text message



## NewCall

> InitiateNewCallSessionResponse NewCall(ctx).Body(body).Execute()

Initiate a new call



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
    body := *openapiclient.NewNewCallRequest("e0", "+4915799912345") // NewCallRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsApi.NewCall(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.NewCall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewCall`: InitiateNewCallSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.NewCall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewCallRequest**](NewCallRequest.md) |  | 

### Return type

[**InitiateNewCallSessionResponse**](InitiateNewCallSessionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewIvrRecording

> InitiateNewCallSessionResponse NewIvrRecording(ctx).Body(body).Execute()

Initiate a new call to record an ivr message

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
    body := *openapiclient.NewNewIvrRecordingRequest() // NewIvrRecordingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsApi.NewIvrRecording(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.NewIvrRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewIvrRecording`: InitiateNewCallSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.NewIvrRecording`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewIvrRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewIvrRecordingRequest**](NewIvrRecordingRequest.md) |  | 

### Return type

[**InitiateNewCallSessionResponse**](InitiateNewCallSessionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewRecording

> InitiateNewCallSessionResponse NewRecording(ctx).Body(body).Execute()

Initiate a new call to record a greeting

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
    body := *openapiclient.NewNewVoicemailRecordingRecordingRequest() // NewVoicemailRecordingRecordingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsApi.NewRecording(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.NewRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewRecording`: InitiateNewCallSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.NewRecording`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewVoicemailRecordingRecordingRequest**](NewVoicemailRecordingRecordingRequest.md) |  | 

### Return type

[**InitiateNewCallSessionResponse**](InitiateNewCallSessionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlayRecording

> InitiateNewCallSessionResponse PlayRecording(ctx).Body(body).Execute()

Initiate a new call to playback a voicemail

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
    body := *openapiclient.NewPlayRecordingRequest() // PlayRecordingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsApi.PlayRecording(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.PlayRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlayRecording`: InitiateNewCallSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.PlayRecording`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlayRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PlayRecordingRequest**](PlayRecordingRequest.md) |  | 

### Return type

[**InitiateNewCallSessionResponse**](InitiateNewCallSessionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendFax

> ResendFax(ctx).Body(body).Execute()

Resend a faX

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
    body := *openapiclient.NewResendFaxRequest("100018428") // ResendFaxRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsApi.ResendFax(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.ResendFax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResendFaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ResendFaxRequest**](ResendFaxRequest.md) |  | 

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


## SendFax

> SendFaxSessionResponse SendFax(ctx).Body(body).Execute()

Send a fax

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
    body := *openapiclient.NewSendFaxRequest("f0", "+4921112345678", "fax.pdf", "TWF5IHRoZSBmb3VydGggYmUgd2l0aCB5b3U=") // SendFaxRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsApi.SendFax(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.SendFax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendFax`: SendFaxSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.SendFax`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendFaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SendFaxRequest**](SendFaxRequest.md) |  | 

### Return type

[**SendFaxSessionResponse**](SendFaxSessionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendWebSms

> SendWebSms(ctx).Body(body).Execute()

Send a text message

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
    body := *openapiclient.NewSendWebSmsRequest("s0", "+4915799912345", "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua.") // SendWebSmsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsApi.SendWebSms(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.SendWebSms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendWebSmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SendWebSmsRequest**](SendWebSmsRequest.md) |  | 

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

