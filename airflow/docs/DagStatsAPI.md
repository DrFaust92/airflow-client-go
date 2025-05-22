# \DagStatsAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDagStats**](DagStatsAPI.md#GetDagStats) | **Get** /dagStats | List Dag statistics



## GetDagStats

> DagStatsCollectionSchema GetDagStats(ctx).DagIds(dagIds).Execute()

List Dag statistics

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
	dagIds := "dagIds_example" // string | One or more DAG IDs separated by commas to filter relevant Dags. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DagStatsAPI.GetDagStats(context.Background()).DagIds(dagIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DagStatsAPI.GetDagStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDagStats`: DagStatsCollectionSchema
	fmt.Fprintf(os.Stdout, "Response from `DagStatsAPI.GetDagStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDagStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dagIds** | **string** | One or more DAG IDs separated by commas to filter relevant Dags.  | 

### Return type

[**DagStatsCollectionSchema**](DagStatsCollectionSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

