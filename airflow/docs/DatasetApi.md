# \DatasetAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDatasetEvent**](DatasetAPI.md#CreateDatasetEvent) | **Post** /datasets/events | Create dataset event
[**DeleteDagDatasetQueuedEvent**](DatasetAPI.md#DeleteDagDatasetQueuedEvent) | **Delete** /dags/{dag_id}/datasets/queuedEvent/{uri} | Delete a queued Dataset event for a DAG.
[**DeleteDagDatasetQueuedEvents**](DatasetAPI.md#DeleteDagDatasetQueuedEvents) | **Delete** /dags/{dag_id}/datasets/queuedEvent | Delete queued Dataset events for a DAG.
[**DeleteDatasetQueuedEvents**](DatasetAPI.md#DeleteDatasetQueuedEvents) | **Delete** /datasets/queuedEvent/{uri} | Delete queued Dataset events for a Dataset.
[**GetDagDatasetQueuedEvent**](DatasetAPI.md#GetDagDatasetQueuedEvent) | **Get** /dags/{dag_id}/datasets/queuedEvent/{uri} | Get a queued Dataset event for a DAG
[**GetDagDatasetQueuedEvents**](DatasetAPI.md#GetDagDatasetQueuedEvents) | **Get** /dags/{dag_id}/datasets/queuedEvent | Get queued Dataset events for a DAG.
[**GetDataset**](DatasetAPI.md#GetDataset) | **Get** /datasets/{uri} | Get a dataset
[**GetDatasetEvents**](DatasetAPI.md#GetDatasetEvents) | **Get** /datasets/events | Get dataset events
[**GetDatasetQueuedEvents**](DatasetAPI.md#GetDatasetQueuedEvents) | **Get** /datasets/queuedEvent/{uri} | Get queued Dataset events for a Dataset.
[**GetDatasets**](DatasetAPI.md#GetDatasets) | **Get** /datasets | List datasets
[**GetUpstreamDatasetEvents**](DatasetAPI.md#GetUpstreamDatasetEvents) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/upstreamDatasetEvents | Get dataset events for a DAG run



## CreateDatasetEvent

> DatasetEvent CreateDatasetEvent(ctx).CreateDatasetEvent(createDatasetEvent).Execute()

Create dataset event



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
	createDatasetEvent := *openapiclient.NewCreateDatasetEvent("DatasetUri_example") // CreateDatasetEvent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatasetAPI.CreateDatasetEvent(context.Background()).CreateDatasetEvent(createDatasetEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetAPI.CreateDatasetEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDatasetEvent`: DatasetEvent
	fmt.Fprintf(os.Stdout, "Response from `DatasetAPI.CreateDatasetEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatasetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDatasetEvent** | [**CreateDatasetEvent**](CreateDatasetEvent.md) |  | 

### Return type

[**DatasetEvent**](DatasetEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDagDatasetQueuedEvent

> DeleteDagDatasetQueuedEvent(ctx, dagId, uri).Before(before).Execute()

Delete a queued Dataset event for a DAG.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "./airflow"
)

func main() {
	dagId := "dagId_example" // string | The DAG ID.
	uri := "uri_example" // string | The encoded Dataset URI
	before := time.Now() // time.Time | Timestamp to select event logs occurring before. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DatasetAPI.DeleteDagDatasetQueuedEvent(context.Background(), dagId, uri).Before(before).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetAPI.DeleteDagDatasetQueuedEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 
**uri** | **string** | The encoded Dataset URI | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDagDatasetQueuedEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **before** | **time.Time** | Timestamp to select event logs occurring before. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDagDatasetQueuedEvents

> DeleteDagDatasetQueuedEvents(ctx, dagId).Before(before).Execute()

Delete queued Dataset events for a DAG.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "./airflow"
)

func main() {
	dagId := "dagId_example" // string | The DAG ID.
	before := time.Now() // time.Time | Timestamp to select event logs occurring before. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DatasetAPI.DeleteDagDatasetQueuedEvents(context.Background(), dagId).Before(before).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetAPI.DeleteDagDatasetQueuedEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDagDatasetQueuedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **before** | **time.Time** | Timestamp to select event logs occurring before. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatasetQueuedEvents

> DeleteDatasetQueuedEvents(ctx, uri).Before(before).Execute()

Delete queued Dataset events for a Dataset.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "./airflow"
)

func main() {
	uri := "uri_example" // string | The encoded Dataset URI
	before := time.Now() // time.Time | Timestamp to select event logs occurring before. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DatasetAPI.DeleteDatasetQueuedEvents(context.Background(), uri).Before(before).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetAPI.DeleteDatasetQueuedEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uri** | **string** | The encoded Dataset URI | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatasetQueuedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **before** | **time.Time** | Timestamp to select event logs occurring before. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDagDatasetQueuedEvent

> QueuedEvent GetDagDatasetQueuedEvent(ctx, dagId, uri).Before(before).Execute()

Get a queued Dataset event for a DAG



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "./airflow"
)

func main() {
	dagId := "dagId_example" // string | The DAG ID.
	uri := "uri_example" // string | The encoded Dataset URI
	before := time.Now() // time.Time | Timestamp to select event logs occurring before. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatasetAPI.GetDagDatasetQueuedEvent(context.Background(), dagId, uri).Before(before).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetAPI.GetDagDatasetQueuedEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDagDatasetQueuedEvent`: QueuedEvent
	fmt.Fprintf(os.Stdout, "Response from `DatasetAPI.GetDagDatasetQueuedEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 
**uri** | **string** | The encoded Dataset URI | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDagDatasetQueuedEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **before** | **time.Time** | Timestamp to select event logs occurring before. | 

### Return type

[**QueuedEvent**](QueuedEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDagDatasetQueuedEvents

> QueuedEventCollection GetDagDatasetQueuedEvents(ctx, dagId).Before(before).Execute()

Get queued Dataset events for a DAG.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "./airflow"
)

func main() {
	dagId := "dagId_example" // string | The DAG ID.
	before := time.Now() // time.Time | Timestamp to select event logs occurring before. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatasetAPI.GetDagDatasetQueuedEvents(context.Background(), dagId).Before(before).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetAPI.GetDagDatasetQueuedEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDagDatasetQueuedEvents`: QueuedEventCollection
	fmt.Fprintf(os.Stdout, "Response from `DatasetAPI.GetDagDatasetQueuedEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDagDatasetQueuedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **before** | **time.Time** | Timestamp to select event logs occurring before. | 

### Return type

[**QueuedEventCollection**](QueuedEventCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataset

> Dataset GetDataset(ctx, uri).Execute()

Get a dataset



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
	uri := "uri_example" // string | The encoded Dataset URI

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatasetAPI.GetDataset(context.Background(), uri).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetAPI.GetDataset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataset`: Dataset
	fmt.Fprintf(os.Stdout, "Response from `DatasetAPI.GetDataset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uri** | **string** | The encoded Dataset URI | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatasetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dataset**](Dataset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatasetEvents

> DatasetEventCollection GetDatasetEvents(ctx).Limit(limit).Offset(offset).OrderBy(orderBy).DatasetId(datasetId).SourceDagId(sourceDagId).SourceTaskId(sourceTaskId).SourceRunId(sourceRunId).SourceMapIndex(sourceMapIndex).Execute()

Get dataset events



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
	datasetId := int32(56) // int32 | The Dataset ID that updated the dataset. (optional)
	sourceDagId := "sourceDagId_example" // string | The DAG ID that updated the dataset. (optional)
	sourceTaskId := "sourceTaskId_example" // string | The task ID that updated the dataset. (optional)
	sourceRunId := "sourceRunId_example" // string | The DAG run ID that updated the dataset. (optional)
	sourceMapIndex := int32(56) // int32 | The map index that updated the dataset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatasetAPI.GetDatasetEvents(context.Background()).Limit(limit).Offset(offset).OrderBy(orderBy).DatasetId(datasetId).SourceDagId(sourceDagId).SourceTaskId(sourceTaskId).SourceRunId(sourceRunId).SourceMapIndex(sourceMapIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetAPI.GetDatasetEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatasetEvents`: DatasetEventCollection
	fmt.Fprintf(os.Stdout, "Response from `DatasetAPI.GetDatasetEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDatasetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The numbers of items to return. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **orderBy** | **string** | The name of the field to order the results by. Prefix a field name with &#x60;-&#x60; to reverse the sort order.  *New in version 2.1.0*  | 
 **datasetId** | **int32** | The Dataset ID that updated the dataset. | 
 **sourceDagId** | **string** | The DAG ID that updated the dataset. | 
 **sourceTaskId** | **string** | The task ID that updated the dataset. | 
 **sourceRunId** | **string** | The DAG run ID that updated the dataset. | 
 **sourceMapIndex** | **int32** | The map index that updated the dataset. | 

### Return type

[**DatasetEventCollection**](DatasetEventCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatasetQueuedEvents

> QueuedEventCollection GetDatasetQueuedEvents(ctx, uri).Before(before).Execute()

Get queued Dataset events for a Dataset.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "./airflow"
)

func main() {
	uri := "uri_example" // string | The encoded Dataset URI
	before := time.Now() // time.Time | Timestamp to select event logs occurring before. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatasetAPI.GetDatasetQueuedEvents(context.Background(), uri).Before(before).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetAPI.GetDatasetQueuedEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatasetQueuedEvents`: QueuedEventCollection
	fmt.Fprintf(os.Stdout, "Response from `DatasetAPI.GetDatasetQueuedEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uri** | **string** | The encoded Dataset URI | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatasetQueuedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **before** | **time.Time** | Timestamp to select event logs occurring before. | 

### Return type

[**QueuedEventCollection**](QueuedEventCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatasets

> DatasetCollection GetDatasets(ctx).Limit(limit).Offset(offset).OrderBy(orderBy).UriPattern(uriPattern).DagIds(dagIds).Execute()

List datasets

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
	uriPattern := "uriPattern_example" // string | If set, only return datasets with uris matching this pattern.  (optional)
	dagIds := "dagIds_example" // string | One or more DAG IDs separated by commas to filter datasets by associated DAGs either consuming or producing.  *New in version 2.9.0*  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatasetAPI.GetDatasets(context.Background()).Limit(limit).Offset(offset).OrderBy(orderBy).UriPattern(uriPattern).DagIds(dagIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetAPI.GetDatasets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatasets`: DatasetCollection
	fmt.Fprintf(os.Stdout, "Response from `DatasetAPI.GetDatasets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDatasetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The numbers of items to return. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **orderBy** | **string** | The name of the field to order the results by. Prefix a field name with &#x60;-&#x60; to reverse the sort order.  *New in version 2.1.0*  | 
 **uriPattern** | **string** | If set, only return datasets with uris matching this pattern.  | 
 **dagIds** | **string** | One or more DAG IDs separated by commas to filter datasets by associated DAGs either consuming or producing.  *New in version 2.9.0*  | 

### Return type

[**DatasetCollection**](DatasetCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpstreamDatasetEvents

> DatasetEventCollection GetUpstreamDatasetEvents(ctx, dagId, dagRunId).Execute()

Get dataset events for a DAG run



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
	dagId := "dagId_example" // string | The DAG ID.
	dagRunId := "dagRunId_example" // string | The DAG run ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatasetAPI.GetUpstreamDatasetEvents(context.Background(), dagId, dagRunId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasetAPI.GetUpstreamDatasetEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpstreamDatasetEvents`: DatasetEventCollection
	fmt.Fprintf(os.Stdout, "Response from `DatasetAPI.GetUpstreamDatasetEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 
**dagRunId** | **string** | The DAG run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpstreamDatasetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DatasetEventCollection**](DatasetEventCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

