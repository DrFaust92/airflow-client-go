# \DAGAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDag**](DAGAPI.md#DeleteDag) | **Delete** /dags/{dag_id} | Delete a DAG
[**GetDag**](DAGAPI.md#GetDag) | **Get** /dags/{dag_id} | Get basic information about a DAG
[**GetDagDetails**](DAGAPI.md#GetDagDetails) | **Get** /dags/{dag_id}/details | Get a simplified representation of DAG
[**GetDagSource**](DAGAPI.md#GetDagSource) | **Get** /dagSources/{file_token} | Get a source code
[**GetDags**](DAGAPI.md#GetDags) | **Get** /dags | List DAGs
[**GetTask**](DAGAPI.md#GetTask) | **Get** /dags/{dag_id}/tasks/{task_id} | Get simplified representation of a task
[**GetTasks**](DAGAPI.md#GetTasks) | **Get** /dags/{dag_id}/tasks | Get tasks for DAG
[**PatchDag**](DAGAPI.md#PatchDag) | **Patch** /dags/{dag_id} | Update a DAG
[**PatchDags**](DAGAPI.md#PatchDags) | **Patch** /dags | Update DAGs
[**PostClearTaskInstances**](DAGAPI.md#PostClearTaskInstances) | **Post** /dags/{dag_id}/clearTaskInstances | Clear a set of task instances
[**PostSetTaskInstancesState**](DAGAPI.md#PostSetTaskInstancesState) | **Post** /dags/{dag_id}/updateTaskInstancesState | Set a state of task instances
[**ReparseDagFile**](DAGAPI.md#ReparseDagFile) | **Put** /parseDagFile/{file_token} | Request re-parsing of a DAG file



## DeleteDag

> DeleteDag(ctx, dagId).Execute()

Delete a DAG



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DAGAPI.DeleteDag(context.Background(), dagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DAGAPI.DeleteDag``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetDag

> DAG GetDag(ctx, dagId).Fields(fields).Execute()

Get basic information about a DAG



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
	fields := []string{"Inner_example"} // []string | List of field for return.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DAGAPI.GetDag(context.Background(), dagId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DAGAPI.GetDag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDag`: DAG
	fmt.Fprintf(os.Stdout, "Response from `DAGAPI.GetDag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | List of field for return.  | 

### Return type

[**DAG**](DAG.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDagDetails

> DAGDetail GetDagDetails(ctx, dagId).Fields(fields).Execute()

Get a simplified representation of DAG



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
	fields := []string{"Inner_example"} // []string | List of field for return.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DAGAPI.GetDagDetails(context.Background(), dagId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DAGAPI.GetDagDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDagDetails`: DAGDetail
	fmt.Fprintf(os.Stdout, "Response from `DAGAPI.GetDagDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDagDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | List of field for return.  | 

### Return type

[**DAGDetail**](DAGDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDagSource

> GetDagSource200Response GetDagSource(ctx, fileToken).Execute()

Get a source code



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
	fileToken := "fileToken_example" // string | The key containing the encrypted path to the file. Encryption and decryption take place only on the server. This prevents the client from reading an non-DAG file. This also ensures API extensibility, because the format of encrypted data may change. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DAGAPI.GetDagSource(context.Background(), fileToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DAGAPI.GetDagSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDagSource`: GetDagSource200Response
	fmt.Fprintf(os.Stdout, "Response from `DAGAPI.GetDagSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileToken** | **string** | The key containing the encrypted path to the file. Encryption and decryption take place only on the server. This prevents the client from reading an non-DAG file. This also ensures API extensibility, because the format of encrypted data may change.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDagSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDagSource200Response**](GetDagSource200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDags

> DAGCollection GetDags(ctx).Limit(limit).Offset(offset).OrderBy(orderBy).Tags(tags).OnlyActive(onlyActive).Paused(paused).Fields(fields).DagIdPattern(dagIdPattern).Execute()

List DAGs



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
	tags := []string{"Inner_example"} // []string | List of tags to filter results.  *New in version 2.2.0*  (optional)
	onlyActive := true // bool | Only filter active DAGs.  *New in version 2.1.1*  (optional) (default to true)
	paused := true // bool | Only filter paused/unpaused DAGs. If absent or null, it returns paused and unpaused DAGs.  *New in version 2.6.0*  (optional)
	fields := []string{"Inner_example"} // []string | List of field for return.  (optional)
	dagIdPattern := "dagIdPattern_example" // string | If set, only return DAGs with dag_ids matching this pattern.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DAGAPI.GetDags(context.Background()).Limit(limit).Offset(offset).OrderBy(orderBy).Tags(tags).OnlyActive(onlyActive).Paused(paused).Fields(fields).DagIdPattern(dagIdPattern).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DAGAPI.GetDags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDags`: DAGCollection
	fmt.Fprintf(os.Stdout, "Response from `DAGAPI.GetDags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The numbers of items to return. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **orderBy** | **string** | The name of the field to order the results by. Prefix a field name with &#x60;-&#x60; to reverse the sort order.  *New in version 2.1.0*  | 
 **tags** | **[]string** | List of tags to filter results.  *New in version 2.2.0*  | 
 **onlyActive** | **bool** | Only filter active DAGs.  *New in version 2.1.1*  | [default to true]
 **paused** | **bool** | Only filter paused/unpaused DAGs. If absent or null, it returns paused and unpaused DAGs.  *New in version 2.6.0*  | 
 **fields** | **[]string** | List of field for return.  | 
 **dagIdPattern** | **string** | If set, only return DAGs with dag_ids matching this pattern.  | 

### Return type

[**DAGCollection**](DAGCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTask

> Task GetTask(ctx, dagId, taskId).Execute()

Get simplified representation of a task

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
	taskId := "taskId_example" // string | The task ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DAGAPI.GetTask(context.Background(), dagId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DAGAPI.GetTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTask`: Task
	fmt.Fprintf(os.Stdout, "Response from `DAGAPI.GetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 
**taskId** | **string** | The task ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasks

> TaskCollection GetTasks(ctx, dagId).OrderBy(orderBy).Execute()

Get tasks for DAG

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
	orderBy := "orderBy_example" // string | The name of the field to order the results by. Prefix a field name with `-` to reverse the sort order.  *New in version 2.1.0*  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DAGAPI.GetTasks(context.Background(), dagId).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DAGAPI.GetTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTasks`: TaskCollection
	fmt.Fprintf(os.Stdout, "Response from `DAGAPI.GetTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderBy** | **string** | The name of the field to order the results by. Prefix a field name with &#x60;-&#x60; to reverse the sort order.  *New in version 2.1.0*  | 

### Return type

[**TaskCollection**](TaskCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDag

> DAG PatchDag(ctx, dagId).DAG(dAG).UpdateMask(updateMask).Execute()

Update a DAG

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
	dAG := *openapiclient.NewDAG() // DAG | 
	updateMask := []string{"Inner_example"} // []string | The fields to update on the resource. If absent or empty, all modifiable fields are updated. A comma-separated list of fully qualified names of fields.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DAGAPI.PatchDag(context.Background(), dagId).DAG(dAG).UpdateMask(updateMask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DAGAPI.PatchDag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDag`: DAG
	fmt.Fprintf(os.Stdout, "Response from `DAGAPI.PatchDag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dAG** | [**DAG**](DAG.md) |  | 
 **updateMask** | **[]string** | The fields to update on the resource. If absent or empty, all modifiable fields are updated. A comma-separated list of fully qualified names of fields.  | 

### Return type

[**DAG**](DAG.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDags

> DAGCollection PatchDags(ctx).DagIdPattern(dagIdPattern).DAG(dAG).Limit(limit).Offset(offset).Tags(tags).UpdateMask(updateMask).OnlyActive(onlyActive).Execute()

Update DAGs



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
	dagIdPattern := "dagIdPattern_example" // string | If set, only update DAGs with dag_ids matching this pattern. 
	dAG := *openapiclient.NewDAG() // DAG | 
	limit := int32(56) // int32 | The numbers of items to return. (optional) (default to 100)
	offset := int32(56) // int32 | The number of items to skip before starting to collect the result set. (optional)
	tags := []string{"Inner_example"} // []string | List of tags to filter results.  *New in version 2.2.0*  (optional)
	updateMask := []string{"Inner_example"} // []string | The fields to update on the resource. If absent or empty, all modifiable fields are updated. A comma-separated list of fully qualified names of fields.  (optional)
	onlyActive := true // bool | Only filter active DAGs.  *New in version 2.1.1*  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DAGAPI.PatchDags(context.Background()).DagIdPattern(dagIdPattern).DAG(dAG).Limit(limit).Offset(offset).Tags(tags).UpdateMask(updateMask).OnlyActive(onlyActive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DAGAPI.PatchDags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDags`: DAGCollection
	fmt.Fprintf(os.Stdout, "Response from `DAGAPI.PatchDags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchDagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dagIdPattern** | **string** | If set, only update DAGs with dag_ids matching this pattern.  | 
 **dAG** | [**DAG**](DAG.md) |  | 
 **limit** | **int32** | The numbers of items to return. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **tags** | **[]string** | List of tags to filter results.  *New in version 2.2.0*  | 
 **updateMask** | **[]string** | The fields to update on the resource. If absent or empty, all modifiable fields are updated. A comma-separated list of fully qualified names of fields.  | 
 **onlyActive** | **bool** | Only filter active DAGs.  *New in version 2.1.1*  | [default to true]

### Return type

[**DAGCollection**](DAGCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClearTaskInstances

> TaskInstanceReferenceCollection PostClearTaskInstances(ctx, dagId).ClearTaskInstances(clearTaskInstances).Execute()

Clear a set of task instances



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
	clearTaskInstances := *openapiclient.NewClearTaskInstances() // ClearTaskInstances | Parameters of action

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DAGAPI.PostClearTaskInstances(context.Background(), dagId).ClearTaskInstances(clearTaskInstances).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DAGAPI.PostClearTaskInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClearTaskInstances`: TaskInstanceReferenceCollection
	fmt.Fprintf(os.Stdout, "Response from `DAGAPI.PostClearTaskInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClearTaskInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clearTaskInstances** | [**ClearTaskInstances**](ClearTaskInstances.md) | Parameters of action | 

### Return type

[**TaskInstanceReferenceCollection**](TaskInstanceReferenceCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSetTaskInstancesState

> TaskInstanceReferenceCollection PostSetTaskInstancesState(ctx, dagId).UpdateTaskInstancesState(updateTaskInstancesState).Execute()

Set a state of task instances



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
	updateTaskInstancesState := *openapiclient.NewUpdateTaskInstancesState() // UpdateTaskInstancesState | Parameters of action

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DAGAPI.PostSetTaskInstancesState(context.Background(), dagId).UpdateTaskInstancesState(updateTaskInstancesState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DAGAPI.PostSetTaskInstancesState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSetTaskInstancesState`: TaskInstanceReferenceCollection
	fmt.Fprintf(os.Stdout, "Response from `DAGAPI.PostSetTaskInstancesState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSetTaskInstancesStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTaskInstancesState** | [**UpdateTaskInstancesState**](UpdateTaskInstancesState.md) | Parameters of action | 

### Return type

[**TaskInstanceReferenceCollection**](TaskInstanceReferenceCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReparseDagFile

> ReparseDagFile(ctx, fileToken).Execute()

Request re-parsing of a DAG file



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
	fileToken := "fileToken_example" // string | The key containing the encrypted path to the file. Encryption and decryption take place only on the server. This prevents the client from reading an non-DAG file. This also ensures API extensibility, because the format of encrypted data may change. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DAGAPI.ReparseDagFile(context.Background(), fileToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DAGAPI.ReparseDagFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileToken** | **string** | The key containing the encrypted path to the file. Encryption and decryption take place only on the server. This prevents the client from reading an non-DAG file. This also ensures API extensibility, because the format of encrypted data may change.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReparseDagFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

