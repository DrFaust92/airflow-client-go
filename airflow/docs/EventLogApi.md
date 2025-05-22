# \EventLogAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventLog**](EventLogAPI.md#GetEventLog) | **Get** /eventLogs/{event_log_id} | Get a log entry
[**GetEventLogs**](EventLogAPI.md#GetEventLogs) | **Get** /eventLogs | List log entries



## GetEventLog

> EventLog GetEventLog(ctx, eventLogId).Execute()

Get a log entry

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
	eventLogId := int32(56) // int32 | The event log ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventLogAPI.GetEventLog(context.Background(), eventLogId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventLogAPI.GetEventLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventLog`: EventLog
	fmt.Fprintf(os.Stdout, "Response from `EventLogAPI.GetEventLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventLogId** | **int32** | The event log ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventLog**](EventLog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventLogs

> EventLogCollection GetEventLogs(ctx).Limit(limit).Offset(offset).OrderBy(orderBy).DagId(dagId).TaskId(taskId).RunId(runId).MapIndex(mapIndex).TryNumber(tryNumber).Event(event).Owner(owner).Before(before).After(after).IncludedEvents(includedEvents).ExcludedEvents(excludedEvents).Execute()

List log entries



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
	limit := int32(56) // int32 | The numbers of items to return. (optional) (default to 100)
	offset := int32(56) // int32 | The number of items to skip before starting to collect the result set. (optional)
	orderBy := "orderBy_example" // string | The name of the field to order the results by. Prefix a field name with `-` to reverse the sort order.  *New in version 2.1.0*  (optional)
	dagId := "dagId_example" // string | Returns objects matched by the DAG ID. (optional)
	taskId := "taskId_example" // string | Returns objects matched by the Task ID. (optional)
	runId := "runId_example" // string | Returns objects matched by the Run ID. (optional)
	mapIndex := int32(56) // int32 | Filter on map index for mapped task. (optional)
	tryNumber := int32(56) // int32 | Filter on try_number for task instance. (optional)
	event := "event_example" // string | The name of event log. (optional)
	owner := "owner_example" // string | The owner's name of event log. (optional)
	before := time.Now() // time.Time | Timestamp to select event logs occurring before. (optional)
	after := time.Now() // time.Time | Timestamp to select event logs occurring after. (optional)
	includedEvents := "includedEvents_example" // string | One or more event names separated by commas. If set, only return event logs with events matching this pattern. *New in version 2.9.0*  (optional)
	excludedEvents := "excludedEvents_example" // string | One or more event names separated by commas. If set, only return event logs with events that do not match this pattern. *New in version 2.9.0*  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventLogAPI.GetEventLogs(context.Background()).Limit(limit).Offset(offset).OrderBy(orderBy).DagId(dagId).TaskId(taskId).RunId(runId).MapIndex(mapIndex).TryNumber(tryNumber).Event(event).Owner(owner).Before(before).After(after).IncludedEvents(includedEvents).ExcludedEvents(excludedEvents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventLogAPI.GetEventLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventLogs`: EventLogCollection
	fmt.Fprintf(os.Stdout, "Response from `EventLogAPI.GetEventLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The numbers of items to return. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **orderBy** | **string** | The name of the field to order the results by. Prefix a field name with &#x60;-&#x60; to reverse the sort order.  *New in version 2.1.0*  | 
 **dagId** | **string** | Returns objects matched by the DAG ID. | 
 **taskId** | **string** | Returns objects matched by the Task ID. | 
 **runId** | **string** | Returns objects matched by the Run ID. | 
 **mapIndex** | **int32** | Filter on map index for mapped task. | 
 **tryNumber** | **int32** | Filter on try_number for task instance. | 
 **event** | **string** | The name of event log. | 
 **owner** | **string** | The owner&#39;s name of event log. | 
 **before** | **time.Time** | Timestamp to select event logs occurring before. | 
 **after** | **time.Time** | Timestamp to select event logs occurring after. | 
 **includedEvents** | **string** | One or more event names separated by commas. If set, only return event logs with events matching this pattern. *New in version 2.9.0*  | 
 **excludedEvents** | **string** | One or more event names separated by commas. If set, only return event logs with events that do not match this pattern. *New in version 2.9.0*  | 

### Return type

[**EventLogCollection**](EventLogCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

