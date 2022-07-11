# \GroupsApi

All URIs are relative to *https://api.sipgate.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDeviceToGroup**](GroupsApi.md#AddDeviceToGroup) | **Post** /groups/{groupId}/devices | Add device to group
[**GetGroup**](GroupsApi.md#GetGroup) | **Get** /groups/{groupId} | Get a single groups details
[**GetGroupFaxlineCallerId**](GroupsApi.md#GetGroupFaxlineCallerId) | **Get** /groupfaxlines/{faxlineId}/callerid | Get the caller ID of a specific group fax
[**GetGroupFaxlineNumbers**](GroupsApi.md#GetGroupFaxlineNumbers) | **Get** /groupfaxlines/{faxlineId}/numbers | List all phone numbers routed to group fax line
[**GetGroupFaxlines**](GroupsApi.md#GetGroupFaxlines) | **Get** /groups/{groupId}/faxlines | List all fax lines for a specific group
[**GetGroupFaxlinesForUser**](GroupsApi.md#GetGroupFaxlinesForUser) | **Get** /groupfaxlines | List all group fax lines
[**GetGroupNumbers**](GroupsApi.md#GetGroupNumbers) | **Get** /groups/{groupId}/numbers | List all phone numbers routed to group
[**GetGroupUsers**](GroupsApi.md#GetGroupUsers) | **Get** /groups/{groupId}/users | List all group members
[**GetGroupVoicemail**](GroupsApi.md#GetGroupVoicemail) | **Get** /groups/{groupId}/voicemail | List all voicemails for a specific group
[**GetGroups**](GroupsApi.md#GetGroups) | **Get** /groups | List all groups
[**RemoveDeviceFromGroup**](GroupsApi.md#RemoveDeviceFromGroup) | **Delete** /groups/{groupId}/devices/{deviceId} | Remove device from group



## AddDeviceToGroup

> AddDeviceToGroup(ctx, groupId).Body(body).Execute()

Add device to group

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
    groupId := "g0" // string | The unique group identifier
    body := *openapiclient.NewRouteDeviceRequest("e0") // RouteDeviceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.AddDeviceToGroup(context.Background(), groupId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.AddDeviceToGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The unique group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDeviceToGroupRequest struct via the builder pattern


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


## GetGroup

> GroupResponse GetGroup(ctx, groupId).Execute()

Get a single groups details

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
    groupId := "g0" // string | The unique group identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroup`: GroupResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The unique group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupResponse**](GroupResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupFaxlineCallerId

> FaxlineCallerIdResponse GetGroupFaxlineCallerId(ctx, faxlineId).Execute()

Get the caller ID of a specific group fax

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
    faxlineId := "f0" // string | The unique fax line identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupFaxlineCallerId(context.Background(), faxlineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupFaxlineCallerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupFaxlineCallerId`: FaxlineCallerIdResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupFaxlineCallerId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**faxlineId** | **string** | The unique fax line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupFaxlineCallerIdRequest struct via the builder pattern


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


## GetGroupFaxlineNumbers

> EndpointNumbersResponse GetGroupFaxlineNumbers(ctx, faxlineId).Execute()

List all phone numbers routed to group fax line

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
    faxlineId := "f0" // string | The unique fax line identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupFaxlineNumbers(context.Background(), faxlineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupFaxlineNumbers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupFaxlineNumbers`: EndpointNumbersResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupFaxlineNumbers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**faxlineId** | **string** | The unique fax line identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupFaxlineNumbersRequest struct via the builder pattern


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


## GetGroupFaxlines

> GroupFaxlinesResponse GetGroupFaxlines(ctx, groupId).Execute()

List all fax lines for a specific group

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
    groupId := "g0" // string | The unique group identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupFaxlines(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupFaxlines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupFaxlines`: GroupFaxlinesResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupFaxlines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The unique group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupFaxlinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupFaxlinesResponse**](GroupFaxlinesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupFaxlinesForUser

> GroupFaxlinesResponse GetGroupFaxlinesForUser(ctx).UserId(userId).Execute()

List all group fax lines

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
    userId := "w0" // string | The unique user identifier (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupFaxlinesForUser(context.Background()).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupFaxlinesForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupFaxlinesForUser`: GroupFaxlinesResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupFaxlinesForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupFaxlinesForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | The unique user identifier | 

### Return type

[**GroupFaxlinesResponse**](GroupFaxlinesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupNumbers

> GroupNumbersResponse GetGroupNumbers(ctx, groupId).Execute()

List all phone numbers routed to group

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
    groupId := "g0" // string | The unique group identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupNumbers(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupNumbers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupNumbers`: GroupNumbersResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupNumbers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The unique group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupNumbersResponse**](GroupNumbersResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupUsers

> GroupUsersResponse GetGroupUsers(ctx, groupId).Execute()

List all group members

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
    groupId := "g0" // string | The unique group identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupUsers(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupUsers`: GroupUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The unique group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupUsersResponse**](GroupUsersResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupVoicemail

> PhonelineVoicemailsResponse GetGroupVoicemail(ctx, groupId).Execute()

List all voicemails for a specific group

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
    groupId := "g0" // string | The unique group identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupVoicemail(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupVoicemail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupVoicemail`: PhonelineVoicemailsResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupVoicemail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The unique group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupVoicemailRequest struct via the builder pattern


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


## GetGroups

> SimpleGroupsResponse GetGroups(ctx).UserId(userId).Execute()

List all groups

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
    userId := "w0" // string | The unique user identifier (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroups(context.Background()).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroups`: SimpleGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | The unique user identifier | 

### Return type

[**SimpleGroupsResponse**](SimpleGroupsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDeviceFromGroup

> RemoveDeviceFromGroup(ctx, groupId, deviceId).Execute()

Remove device from group

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
    groupId := "g0" // string | The unique group identifier
    deviceId := "e0" // string | The unique device identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.RemoveDeviceFromGroup(context.Background(), groupId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.RemoveDeviceFromGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The unique group identifier | 
**deviceId** | **string** | The unique device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDeviceFromGroupRequest struct via the builder pattern


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

