# \LogApi

All URIs are relative to *https://api.sipgate.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWebhooksLog**](LogApi.md#GetWebhooksLog) | **Get** /log/webhooks | Get your webhooks log



## GetWebhooksLog

> PhonelineSipgateIoLogsResponse GetWebhooksLog(ctx).Execute()

Get your webhooks log

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
    resp, r, err := apiClient.LogApi.GetWebhooksLog(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogApi.GetWebhooksLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhooksLog`: PhonelineSipgateIoLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `LogApi.GetWebhooksLog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksLogRequest struct via the builder pattern


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

