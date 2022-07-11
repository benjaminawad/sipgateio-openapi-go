# \VoicemailsApi

All URIs are relative to *https://api.sipgate.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVoicemail**](VoicemailsApi.md#GetVoicemail) | **Get** /voicemails/{voicemailId} | Get a single voicemail
[**GetVoicemails**](VoicemailsApi.md#GetVoicemails) | **Get** /voicemails | List all voicemails



## GetVoicemail

> VoicemailResponse GetVoicemail(ctx, voicemailId).Execute()

Get a single voicemail

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
    voicemailId := "v0" // string | The unique voicemail identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VoicemailsApi.GetVoicemail(context.Background(), voicemailId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsApi.GetVoicemail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVoicemail`: VoicemailResponse
    fmt.Fprintf(os.Stdout, "Response from `VoicemailsApi.GetVoicemail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voicemailId** | **string** | The unique voicemail identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoicemailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VoicemailResponse**](VoicemailResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoicemails

> VoicemailsResponse GetVoicemails(ctx).Execute()

List all voicemails

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
    resp, r, err := apiClient.VoicemailsApi.GetVoicemails(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoicemailsApi.GetVoicemails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVoicemails`: VoicemailsResponse
    fmt.Fprintf(os.Stdout, "Response from `VoicemailsApi.GetVoicemails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoicemailsRequest struct via the builder pattern


### Return type

[**VoicemailsResponse**](VoicemailsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

