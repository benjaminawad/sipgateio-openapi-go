# \DevicesApi

All URIs are relative to *https://api.sipgate.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateSim**](DevicesApi.md#ActivateSim) | **Post** /{userId}/devices/{deviceId}/sim | Activate a SIM card for a specific device
[**ChangeCallerDisplayOfDevice**](DevicesApi.md#ChangeCallerDisplayOfDevice) | **Put** /devices/{deviceId}/external/incomingcalldisplay | Update the incoming call display for an existing external device
[**ChangeCallerIdSettings**](DevicesApi.md#ChangeCallerIdSettings) | **Put** /devices/{deviceId}/callerid | Update the caller ID for an existing device
[**ChangeDeviceAlias**](DevicesApi.md#ChangeDeviceAlias) | **Put** /devices/{deviceId}/alias | Update alias for an existing device
[**ChangeDevicePassword**](DevicesApi.md#ChangeDevicePassword) | **Post** /devices/{deviceId}/credentials/password | Change password for an existing device
[**ChangeDeviceSettings**](DevicesApi.md#ChangeDeviceSettings) | **Put** /devices/{deviceId} | Update settings for an existing device
[**ChangeLocalprefixSettings**](DevicesApi.md#ChangeLocalprefixSettings) | **Put** /devices/{deviceId}/localprefix | Update the local prefix setting for an existing device
[**ChangeSingleRowDisplay**](DevicesApi.md#ChangeSingleRowDisplay) | **Put** /devices/{deviceId}/singlerowdisplay | Update single row display setting for an existing device
[**ChangeTargetNumberOfDevice**](DevicesApi.md#ChangeTargetNumberOfDevice) | **Put** /devices/{deviceId}/external/targetnumber | Update the target phone number for an existing external device
[**ChangeTariffAnnouncement**](DevicesApi.md#ChangeTariffAnnouncement) | **Put** /devices/{deviceId}/tariffannouncement | Update tariff announcement setting for an existing device
[**CreateDevice**](DevicesApi.md#CreateDevice) | **Post** /{userId}/devices | Create a new device
[**CreateExternalDevice**](DevicesApi.md#CreateExternalDevice) | **Post** /{userId}/devices/external | Create a new external device
[**CreateMobileDevice**](DevicesApi.md#CreateMobileDevice) | **Post** /{userId}/devices/mobile | Create a new mobile device
[**CreateRegisterDevice**](DevicesApi.md#CreateRegisterDevice) | **Post** /{userId}/devices/register | Create a new register device
[**DeactivateSim**](DevicesApi.md#DeactivateSim) | **Delete** /{userId}/devices/{deviceId}/sim | Deactivate the SIM card for a specific device
[**DeleteDevice**](DevicesApi.md#DeleteDevice) | **Delete** /devices/{deviceId} | Delete an existing device
[**GetCallerId**](DevicesApi.md#GetCallerId) | **Get** /devices/{deviceId}/callerid | Get the caller ID of a specific device
[**GetDevice**](DevicesApi.md#GetDevice) | **Get** /devices/{deviceId} | Get a specific device
[**GetDeviceContingents**](DevicesApi.md#GetDeviceContingents) | **Get** /{userId}/devices/{deviceId}/contingents | List all device contingents
[**GetDevices**](DevicesApi.md#GetDevices) | **Get** /{userId}/devices | List all devices
[**GetLocalprefix**](DevicesApi.md#GetLocalprefix) | **Get** /devices/{deviceId}/localprefix | Get the local prefix setting of a specific device
[**GetSingleRowDisplay**](DevicesApi.md#GetSingleRowDisplay) | **Get** /devices/{deviceId}/singlerowdisplay | Get single row display setting of a specific device
[**GetTariffAnnouncement**](DevicesApi.md#GetTariffAnnouncement) | **Get** /devices/{deviceId}/tariffannouncement | Get tariff announcement setting of a specific device
[**OrderSim**](DevicesApi.md#OrderSim) | **Post** /{userId}/devices/{deviceId}/sim/orders | Order a new SIM card for a specific device



## ActivateSim

> ActivateSim(ctx, userId, deviceId).Body(body).Execute()

Activate a SIM card for a specific device

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
    deviceId := "y0" // string | The unique mobile device identifier
    body := *openapiclient.NewActivateSimRequest() // ActivateSimRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.ActivateSim(context.Background(), userId, deviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.ActivateSim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**deviceId** | **string** | The unique mobile device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateSimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ActivateSimRequest**](ActivateSimRequest.md) |  | 

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


## ChangeCallerDisplayOfDevice

> ChangeCallerDisplayOfDevice(ctx, deviceId).Body(body).Execute()

Update the incoming call display for an existing external device

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
    deviceId := "x0" // string | The unique external device identifier
    body := *openapiclient.NewChangeExternalDeviceIncomingCallDisplayRequest("IncomingCallDisplay_example") // ChangeExternalDeviceIncomingCallDisplayRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.ChangeCallerDisplayOfDevice(context.Background(), deviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.ChangeCallerDisplayOfDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The unique external device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeCallerDisplayOfDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ChangeExternalDeviceIncomingCallDisplayRequest**](ChangeExternalDeviceIncomingCallDisplayRequest.md) |  | 

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


## ChangeCallerIdSettings

> ChangeCallerIdSettings(ctx, deviceId).Body(body).Execute()

Update the caller ID for an existing device

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
    deviceId := "e0" // string | The unique device identifier
    body := *openapiclient.NewCallerIdRequest() // CallerIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.ChangeCallerIdSettings(context.Background(), deviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.ChangeCallerIdSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The unique device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeCallerIdSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CallerIdRequest**](CallerIdRequest.md) |  | 

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


## ChangeDeviceAlias

> ChangeDeviceAlias(ctx, deviceId).Body(body).Execute()

Update alias for an existing device

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
    deviceId := "e0" // string | The unique device identifier
    body := *openapiclient.NewSetDeviceAliasRequest() // SetDeviceAliasRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.ChangeDeviceAlias(context.Background(), deviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.ChangeDeviceAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The unique device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeDeviceAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SetDeviceAliasRequest**](SetDeviceAliasRequest.md) |  | 

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


## ChangeDevicePassword

> ChangeDevicePassword(ctx, deviceId).Execute()

Change password for an existing device

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
    deviceId := "e0" // string | The unique device identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.ChangeDevicePassword(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.ChangeDevicePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The unique device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeDevicePasswordRequest struct via the builder pattern


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


## ChangeDeviceSettings

> ChangeDeviceSettings(ctx, deviceId).Body(body).Execute()

Update settings for an existing device

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
    deviceId := "e0" // string | The unique device identifier
    body := *openapiclient.NewChangeDeviceRequest() // ChangeDeviceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.ChangeDeviceSettings(context.Background(), deviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.ChangeDeviceSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The unique device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeDeviceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ChangeDeviceRequest**](ChangeDeviceRequest.md) |  | 

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


## ChangeLocalprefixSettings

> ChangeLocalprefixSettings(ctx, deviceId).Body(body).Execute()

Update the local prefix setting for an existing device

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
    deviceId := "e0" // string | The unique device identifier
    body := *openapiclient.NewLocalprefixRequest() // LocalprefixRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.ChangeLocalprefixSettings(context.Background(), deviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.ChangeLocalprefixSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The unique device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeLocalprefixSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**LocalprefixRequest**](LocalprefixRequest.md) |  | 

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


## ChangeSingleRowDisplay

> ChangeSingleRowDisplay(ctx, deviceId).Body(body).Execute()

Update single row display setting for an existing device

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
    deviceId := "e0" // string | The unique device identifier
    body := *openapiclient.NewSingleRowDisplayRequest() // SingleRowDisplayRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.ChangeSingleRowDisplay(context.Background(), deviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.ChangeSingleRowDisplay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The unique device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeSingleRowDisplayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SingleRowDisplayRequest**](SingleRowDisplayRequest.md) |  | 

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


## ChangeTargetNumberOfDevice

> ChangeTargetNumberOfDevice(ctx, deviceId).Body(body).Execute()

Update the target phone number for an existing external device

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
    deviceId := "x0" // string | The unique external device identifier
    body := *openapiclient.NewChangeExternalDeviceTargetNumberRequest() // ChangeExternalDeviceTargetNumberRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.ChangeTargetNumberOfDevice(context.Background(), deviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.ChangeTargetNumberOfDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The unique external device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeTargetNumberOfDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ChangeExternalDeviceTargetNumberRequest**](ChangeExternalDeviceTargetNumberRequest.md) |  | 

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


## ChangeTariffAnnouncement

> ChangeTariffAnnouncement(ctx, deviceId).Body(body).Execute()

Update tariff announcement setting for an existing device

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
    deviceId := "e0" // string | The unique device identifier
    body := *openapiclient.NewTariffAnnouncementRequest() // TariffAnnouncementRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.ChangeTariffAnnouncement(context.Background(), deviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.ChangeTariffAnnouncement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The unique device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeTariffAnnouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TariffAnnouncementRequest**](TariffAnnouncementRequest.md) |  | 

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


## CreateDevice

> DeviceResponse CreateDevice(ctx, userId).Body(body).Execute()

Create a new device



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
    body := *openapiclient.NewCreateDeviceRequest("Type_example") // CreateDeviceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CreateDevice(context.Background(), userId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDevice`: DeviceResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateDeviceRequest**](CreateDeviceRequest.md) |  | 

### Return type

[**DeviceResponse**](DeviceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExternalDevice

> ExternalDeviceResponse CreateExternalDevice(ctx, userId).Body(body).Execute()

Create a new external device



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
    body := *openapiclient.NewCreateExternalDeviceRequest() // CreateExternalDeviceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CreateExternalDevice(context.Background(), userId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateExternalDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateExternalDevice`: ExternalDeviceResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateExternalDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateExternalDeviceRequest**](CreateExternalDeviceRequest.md) |  | 

### Return type

[**ExternalDeviceResponse**](ExternalDeviceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMobileDevice

> MobileDeviceResponse CreateMobileDevice(ctx, userId).Body(body).Execute()

Create a new mobile device



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
    body := *openapiclient.NewCreateMobileDeviceRequest() // CreateMobileDeviceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CreateMobileDevice(context.Background(), userId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateMobileDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMobileDevice`: MobileDeviceResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateMobileDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMobileDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateMobileDeviceRequest**](CreateMobileDeviceRequest.md) |  | 

### Return type

[**MobileDeviceResponse**](MobileDeviceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRegisterDevice

> RegisterDeviceResponse CreateRegisterDevice(ctx, userId).Body(body).Execute()

Create a new register device



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
    body := *openapiclient.NewCreateRegisterDeviceRequest() // CreateRegisterDeviceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.CreateRegisterDevice(context.Background(), userId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateRegisterDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRegisterDevice`: RegisterDeviceResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateRegisterDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRegisterDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateRegisterDeviceRequest**](CreateRegisterDeviceRequest.md) |  | 

### Return type

[**RegisterDeviceResponse**](RegisterDeviceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateSim

> DeactivateSim(ctx, userId, deviceId).Execute()

Deactivate the SIM card for a specific device

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
    deviceId := "y0" // string | The unique mobile device identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DeactivateSim(context.Background(), userId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DeactivateSim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**deviceId** | **string** | The unique mobile device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateSimRequest struct via the builder pattern


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


## DeleteDevice

> DeleteDevice(ctx, deviceId).Execute()

Delete an existing device

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
    deviceId := "e0" // string | The unique device identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DeleteDevice(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DeleteDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The unique device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceRequest struct via the builder pattern


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


## GetCallerId

> CallerIdResponse GetCallerId(ctx, deviceId).Execute()

Get the caller ID of a specific device

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
    deviceId := "e0" // string | The unique device identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetCallerId(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetCallerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCallerId`: CallerIdResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetCallerId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The unique device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCallerIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CallerIdResponse**](CallerIdResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevice

> DeviceResponse GetDevice(ctx, deviceId).Execute()

Get a specific device

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
    deviceId := "e0" // string | The unique device identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetDevice(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevice`: DeviceResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The unique device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceResponse**](DeviceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceContingents

> DeviceContingentsResponse GetDeviceContingents(ctx, userId, deviceId).Execute()

List all device contingents

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
    deviceId := "y0" // string | The unique device identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetDeviceContingents(context.Background(), userId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDeviceContingents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceContingents`: DeviceContingentsResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDeviceContingents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**deviceId** | **string** | The unique device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceContingentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeviceContingentsResponse**](DeviceContingentsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevices

> DevicesResponse GetDevices(ctx, userId).Execute()

List all devices

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
    resp, r, err := apiClient.DevicesApi.GetDevices(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevices`: DevicesResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesRequest struct via the builder pattern


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


## GetLocalprefix

> LocalprefixResponse GetLocalprefix(ctx, deviceId).Execute()

Get the local prefix setting of a specific device

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
    deviceId := "e0" // string | The unique device identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetLocalprefix(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetLocalprefix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocalprefix`: LocalprefixResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetLocalprefix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The unique device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalprefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LocalprefixResponse**](LocalprefixResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleRowDisplay

> SingleRowDisplayResponse GetSingleRowDisplay(ctx, deviceId).Execute()

Get single row display setting of a specific device

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
    deviceId := "e0" // string | The unique device identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetSingleRowDisplay(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetSingleRowDisplay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingleRowDisplay`: SingleRowDisplayResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetSingleRowDisplay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The unique device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleRowDisplayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SingleRowDisplayResponse**](SingleRowDisplayResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTariffAnnouncement

> TariffAnnouncementResponse GetTariffAnnouncement(ctx, deviceId).Execute()

Get tariff announcement setting of a specific device

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
    deviceId := "e0" // string | The unique device identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.GetTariffAnnouncement(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetTariffAnnouncement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTariffAnnouncement`: TariffAnnouncementResponse
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetTariffAnnouncement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The unique device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTariffAnnouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TariffAnnouncementResponse**](TariffAnnouncementResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderSim

> OrderSim(ctx, userId, deviceId).Body(body).Execute()

Order a new SIM card for a specific device

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
    deviceId := "y0" // string | The unique mobile device identifier
    body := *openapiclient.NewOrderSimRequest("123456") // OrderSimRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.OrderSim(context.Background(), userId, deviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.OrderSim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique user identifier | 
**deviceId** | **string** | The unique mobile device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderSimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**OrderSimRequest**](OrderSimRequest.md) |  | 

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

