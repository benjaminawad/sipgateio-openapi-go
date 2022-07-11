# \HistoryApi

All URIs are relative to *https://api.sipgate.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteHistoryEntries**](HistoryApi.md#DeleteHistoryEntries) | **Delete** /history | Delete multiple calls, faxes, SMS or voicemails
[**DeleteHistoryEntry**](HistoryApi.md#DeleteHistoryEntry) | **Delete** /history/{entryId} | Delete a call, fax, SMS or voicemail
[**GetHistory**](HistoryApi.md#GetHistory) | **Get** /history | List all calls, faxes, SMS and voicemails
[**GetHistoryById**](HistoryApi.md#GetHistoryById) | **Get** /history/{entryId} | Get a specific call, fax, SMS or voicemail
[**GetHistoryCSV**](HistoryApi.md#GetHistoryCSV) | **Get** /history/export | List all calls, faxes, SMS and voicemails as csv
[**SetArchive**](HistoryApi.md#SetArchive) | **Put** /history/{entryId}/archive | Archive a specific entry.
[**SetHistoryEntryNote**](HistoryApi.md#SetHistoryEntryNote) | **Put** /history/{entryId}/note | Get a specific call, fax, SMS or voicemail
[**SetReadProperty**](HistoryApi.md#SetReadProperty) | **Put** /history/{entryId}/read | Mark an specific entry as read or unread
[**UpdateHistoryEntries**](HistoryApi.md#UpdateHistoryEntries) | **Put** /history | Update multiple calls, faxes, SMS or voicemails
[**UpdateHistoryEntry**](HistoryApi.md#UpdateHistoryEntry) | **Put** /history/{entryId} | Update a specific call, fax, SMS or voicemail



## DeleteHistoryEntries

> DeleteHistoryEntries(ctx).Id(id).Execute()

Delete multiple calls, faxes, SMS or voicemails

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
    id := []string{"Inner_example"} // []string | List of entry ids to delete (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.DeleteHistoryEntries(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.DeleteHistoryEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHistoryEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** | List of entry ids to delete | 

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


## DeleteHistoryEntry

> DeleteHistoryEntry(ctx, entryId).Execute()

Delete a call, fax, SMS or voicemail

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
    entryId := "abc123" // string | The unique call, fax, sms or voicemail identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.DeleteHistoryEntry(context.Background(), entryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.DeleteHistoryEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryId** | **string** | The unique call, fax, sms or voicemail identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHistoryEntryRequest struct via the builder pattern


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


## GetHistory

> HistoryResponse GetHistory(ctx).ConnectionIds(connectionIds).Types(types).Directions(directions).Offset(offset).Limit(limit).Archived(archived).Starred(starred).From(from).To(to).Phonenumber(phonenumber).Execute()

List all calls, faxes, SMS and voicemails

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
    connectionIds := []string{"Inner_example"} // []string | List of extensions (optional)
    types := []string{"Types_example"} // []string | Filter by type (optional)
    directions := []string{"Directions_example"} // []string | Filter for incoming/outgoing/missed (optional)
    offset := int32(56) // int32 | The offset used for pagination (optional) (default to 0)
    limit := int32(56) // int32 | Limit result rows (optional) (default to 10)
    archived := true // bool | Only show archived events (optional) (default to false)
    starred := []string{"Starred_example"} // []string | Filter for starred/unstarred (optional)
    from := "from_example" // string | Filter 'from' date time in ISO8601 format (optional)
    to := "to_example" // string | Filter 'to' date time in ISO8601 format (optional)
    phonenumber := "phonenumber_example" // string | Filter events to or from given number (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.GetHistory(context.Background()).ConnectionIds(connectionIds).Types(types).Directions(directions).Offset(offset).Limit(limit).Archived(archived).Starred(starred).From(from).To(to).Phonenumber(phonenumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.GetHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHistory`: HistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.GetHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionIds** | **[]string** | List of extensions | 
 **types** | **[]string** | Filter by type | 
 **directions** | **[]string** | Filter for incoming/outgoing/missed | 
 **offset** | **int32** | The offset used for pagination | [default to 0]
 **limit** | **int32** | Limit result rows | [default to 10]
 **archived** | **bool** | Only show archived events | [default to false]
 **starred** | **[]string** | Filter for starred/unstarred | 
 **from** | **string** | Filter &#39;from&#39; date time in ISO8601 format | 
 **to** | **string** | Filter &#39;to&#39; date time in ISO8601 format | 
 **phonenumber** | **string** | Filter events to or from given number | 

### Return type

[**HistoryResponse**](HistoryResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoryById

> HistoryEntryResponse GetHistoryById(ctx, entryId).Execute()

Get a specific call, fax, SMS or voicemail

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
    entryId := "abc123" // string | The unique call, fax, sms or voicemail identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.GetHistoryById(context.Background(), entryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.GetHistoryById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHistoryById`: HistoryEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.GetHistoryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryId** | **string** | The unique call, fax, sms or voicemail identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HistoryEntryResponse**](HistoryEntryResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoryCSV

> HistoryResponse GetHistoryCSV(ctx).ConnectionIds(connectionIds).Types(types).Directions(directions).Offset(offset).Limit(limit).Archived(archived).Starred(starred).From(from).To(to).Execute()

List all calls, faxes, SMS and voicemails as csv

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
    connectionIds := []string{"Inner_example"} // []string | List of extensions (optional)
    types := []string{"Types_example"} // []string | Filter by type (optional)
    directions := []string{"Directions_example"} // []string | Filter for incoming/outgoing/missed (optional)
    offset := int32(56) // int32 | The offset used for pagination (optional) (default to 0)
    limit := int32(56) // int32 | Limit result rows (optional) (default to 10)
    archived := true // bool | Only show archived events (optional) (default to false)
    starred := []string{"Starred_example"} // []string | Filter for starred/unstarred (optional)
    from := "from_example" // string | Filter 'from' date time in ISO8601 format (optional)
    to := "to_example" // string | Filter 'to' date time in ISO8601 format (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.GetHistoryCSV(context.Background()).ConnectionIds(connectionIds).Types(types).Directions(directions).Offset(offset).Limit(limit).Archived(archived).Starred(starred).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.GetHistoryCSV``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHistoryCSV`: HistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.GetHistoryCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoryCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionIds** | **[]string** | List of extensions | 
 **types** | **[]string** | Filter by type | 
 **directions** | **[]string** | Filter for incoming/outgoing/missed | 
 **offset** | **int32** | The offset used for pagination | [default to 0]
 **limit** | **int32** | Limit result rows | [default to 10]
 **archived** | **bool** | Only show archived events | [default to false]
 **starred** | **[]string** | Filter for starred/unstarred | 
 **from** | **string** | Filter &#39;from&#39; date time in ISO8601 format | 
 **to** | **string** | Filter &#39;to&#39; date time in ISO8601 format | 

### Return type

[**HistoryResponse**](HistoryResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetArchive

> SetArchive(ctx, entryId).Body(body).Execute()

Archive a specific entry.

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
    entryId := "12345" // string | The unique call, fax, sms or voicemail identifier
    body := *openapiclient.NewSetHistoryEntryArchiveRequest() // SetHistoryEntryArchiveRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.SetArchive(context.Background(), entryId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.SetArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryId** | **string** | The unique call, fax, sms or voicemail identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SetHistoryEntryArchiveRequest**](SetHistoryEntryArchiveRequest.md) |  | 

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


## SetHistoryEntryNote

> HistoryEntryResponse SetHistoryEntryNote(ctx, entryId).Body(body).Execute()

Get a specific call, fax, SMS or voicemail

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
    entryId := "abc123" // string | The unique call, fax, sms or voicemail identifier
    body := *openapiclient.NewSetHistoryEntryNoteRequest("Note_example") // SetHistoryEntryNoteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.SetHistoryEntryNote(context.Background(), entryId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.SetHistoryEntryNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetHistoryEntryNote`: HistoryEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.SetHistoryEntryNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryId** | **string** | The unique call, fax, sms or voicemail identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetHistoryEntryNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SetHistoryEntryNoteRequest**](SetHistoryEntryNoteRequest.md) |  | 

### Return type

[**HistoryEntryResponse**](HistoryEntryResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetReadProperty

> SetReadProperty(ctx, entryId).Body(body).Execute()

Mark an specific entry as read or unread

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
    entryId := "12345" // string | The unique call, fax, sms or voicemail identifier
    body := *openapiclient.NewSetHistoryEntryReadRequest() // SetHistoryEntryReadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.SetReadProperty(context.Background(), entryId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.SetReadProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryId** | **string** | The unique call, fax, sms or voicemail identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetReadPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SetHistoryEntryReadRequest**](SetHistoryEntryReadRequest.md) |  | 

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


## UpdateHistoryEntries

> UpdateHistoryEntries(ctx).Body(body).Execute()

Update multiple calls, faxes, SMS or voicemails

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
    body := []openapiclient.UpdateHistoryEntriesRequest{*openapiclient.NewUpdateHistoryEntriesRequest("Id_example")} // []UpdateHistoryEntriesRequest | A list of events which should be updated. Must be less than 150 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.UpdateHistoryEntries(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.UpdateHistoryEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHistoryEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]UpdateHistoryEntriesRequest**](UpdateHistoryEntriesRequest.md) | A list of events which should be updated. Must be less than 150 | 

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


## UpdateHistoryEntry

> UpdateHistoryEntry(ctx, entryId).Body(body).Execute()

Update a specific call, fax, SMS or voicemail

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
    entryId := "abc123" // string | The unique call, fax, sms or voicemail identifier
    body := *openapiclient.NewUpdateHistoryEntryRequest() // UpdateHistoryEntryRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.UpdateHistoryEntry(context.Background(), entryId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.UpdateHistoryEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryId** | **string** | The unique call, fax, sms or voicemail identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHistoryEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateHistoryEntryRequest**](UpdateHistoryEntryRequest.md) |  | 

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

