# TaskInstanceHistoryCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalEntries** | Pointer to **int32** | Count of total objects in the current result set before pagination parameters (limit, offset) are applied.  | [optional] 
**TaskInstancesHistory** | Pointer to [**[]TaskInstanceHistory**](TaskInstanceHistory.md) |  | [optional] 

## Methods

### NewTaskInstanceHistoryCollection

`func NewTaskInstanceHistoryCollection() *TaskInstanceHistoryCollection`

NewTaskInstanceHistoryCollection instantiates a new TaskInstanceHistoryCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskInstanceHistoryCollectionWithDefaults

`func NewTaskInstanceHistoryCollectionWithDefaults() *TaskInstanceHistoryCollection`

NewTaskInstanceHistoryCollectionWithDefaults instantiates a new TaskInstanceHistoryCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalEntries

`func (o *TaskInstanceHistoryCollection) GetTotalEntries() int32`

GetTotalEntries returns the TotalEntries field if non-nil, zero value otherwise.

### GetTotalEntriesOk

`func (o *TaskInstanceHistoryCollection) GetTotalEntriesOk() (*int32, bool)`

GetTotalEntriesOk returns a tuple with the TotalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntries

`func (o *TaskInstanceHistoryCollection) SetTotalEntries(v int32)`

SetTotalEntries sets TotalEntries field to given value.

### HasTotalEntries

`func (o *TaskInstanceHistoryCollection) HasTotalEntries() bool`

HasTotalEntries returns a boolean if a field has been set.

### GetTaskInstancesHistory

`func (o *TaskInstanceHistoryCollection) GetTaskInstancesHistory() []TaskInstanceHistory`

GetTaskInstancesHistory returns the TaskInstancesHistory field if non-nil, zero value otherwise.

### GetTaskInstancesHistoryOk

`func (o *TaskInstanceHistoryCollection) GetTaskInstancesHistoryOk() (*[]TaskInstanceHistory, bool)`

GetTaskInstancesHistoryOk returns a tuple with the TaskInstancesHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskInstancesHistory

`func (o *TaskInstanceHistoryCollection) SetTaskInstancesHistory(v []TaskInstanceHistory)`

SetTaskInstancesHistory sets TaskInstancesHistory field to given value.

### HasTaskInstancesHistory

`func (o *TaskInstanceHistoryCollection) HasTaskInstancesHistory() bool`

HasTaskInstancesHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


