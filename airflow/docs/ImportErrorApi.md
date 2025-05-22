# \ImportErrorAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImportError**](ImportErrorAPI.md#GetImportError) | **Get** /importErrors/{import_error_id} | Get an import error
[**GetImportErrors**](ImportErrorAPI.md#GetImportErrors) | **Get** /importErrors | List import errors



## GetImportError

> ImportError GetImportError(ctx, importErrorId).Execute()

Get an import error

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
	importErrorId := int32(56) // int32 | The import error ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportErrorAPI.GetImportError(context.Background(), importErrorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportErrorAPI.GetImportError``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImportError`: ImportError
	fmt.Fprintf(os.Stdout, "Response from `ImportErrorAPI.GetImportError`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**importErrorId** | **int32** | The import error ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImportErrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImportError**](ImportError.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImportErrors

> ImportErrorCollection GetImportErrors(ctx).Limit(limit).Offset(offset).OrderBy(orderBy).Execute()

List import errors

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
	limit := int32(56) // int32 | The numbers of items to return. (optional) (default to 100)
	offset := int32(56) // int32 | The number of items to skip before starting to collect the result set. (optional)
	orderBy := "orderBy_example" // string | The name of the field to order the results by. Prefix a field name with `-` to reverse the sort order.  *New in version 2.1.0*  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportErrorAPI.GetImportErrors(context.Background()).Limit(limit).Offset(offset).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportErrorAPI.GetImportErrors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImportErrors`: ImportErrorCollection
	fmt.Fprintf(os.Stdout, "Response from `ImportErrorAPI.GetImportErrors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImportErrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The numbers of items to return. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **orderBy** | **string** | The name of the field to order the results by. Prefix a field name with &#x60;-&#x60; to reverse the sort order.  *New in version 2.1.0*  | 

### Return type

[**ImportErrorCollection**](ImportErrorCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

