# \ConfigAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfig**](ConfigAPI.md#GetConfig) | **Get** /config | Get current configuration
[**GetValue**](ConfigAPI.md#GetValue) | **Get** /config/section/{section}/option/{option} | Get a option from configuration



## GetConfig

> Config GetConfig(ctx).Section(section).Execute()

Get current configuration

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "./airflow"
)

func main() {
	section := "section_example" // string | If given, only return config of this section. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigAPI.GetConfig(context.Background()).Section(section).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigAPI.GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfig`: Config
	fmt.Fprintf(os.Stdout, "Response from `ConfigAPI.GetConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **section** | **string** | If given, only return config of this section. | 

### Return type

[**Config**](Config.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetValue

> Config GetValue(ctx, section, option).Execute()

Get a option from configuration

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "./airflow"
)

func main() {
	section := "section_example" // string | 
	option := "option_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigAPI.GetValue(context.Background(), section, option).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigAPI.GetValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetValue`: Config
	fmt.Fprintf(os.Stdout, "Response from `ConfigAPI.GetValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**section** | **string** |  | 
**option** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Config**](Config.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

