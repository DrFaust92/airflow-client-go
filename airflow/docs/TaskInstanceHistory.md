# TaskInstanceHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DagId** | Pointer to **string** |  | [optional] 
**DagRunId** | Pointer to **string** | The DagRun ID for this task instance  *New in version 2.3.0*  | [optional] 
**Duration** | Pointer to **NullableFloat32** |  | [optional] 
**EndDate** | Pointer to **NullableString** |  | [optional] 
**Executor** | Pointer to **NullableString** | Executor the task is configured to run on or None (which indicates the default executor)  *New in version 2.10.0*  | [optional] 
**ExecutorConfig** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**MapIndex** | Pointer to **int32** |  | [optional] 
**MaxTries** | Pointer to **int32** |  | [optional] 
**Operator** | Pointer to **NullableString** | *Changed in version 2.1.1*&amp;#58; Field becomes nullable.  | [optional] 
**Pid** | Pointer to **NullableInt32** |  | [optional] 
**Pool** | Pointer to **string** |  | [optional] 
**PoolSlots** | Pointer to **int32** |  | [optional] 
**PriorityWeight** | Pointer to **NullableInt32** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**QueuedWhen** | Pointer to **NullableString** | The datetime that the task enter the state QUEUE, also known as queue_at  | [optional] 
**StartDate** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to [**NullableTaskState**](TaskState.md) |  | [optional] 
**TaskDisplayName** | Pointer to **string** | Human centric display text for the task.  *New in version 2.9.0*  | [optional] 
**TaskId** | Pointer to **string** |  | [optional] 
**TryNumber** | Pointer to **int32** |  | [optional] 
**Unixname** | Pointer to **string** |  | [optional] 

## Methods

### NewTaskInstanceHistory

`func NewTaskInstanceHistory() *TaskInstanceHistory`

NewTaskInstanceHistory instantiates a new TaskInstanceHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskInstanceHistoryWithDefaults

`func NewTaskInstanceHistoryWithDefaults() *TaskInstanceHistory`

NewTaskInstanceHistoryWithDefaults instantiates a new TaskInstanceHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDagId

`func (o *TaskInstanceHistory) GetDagId() string`

GetDagId returns the DagId field if non-nil, zero value otherwise.

### GetDagIdOk

`func (o *TaskInstanceHistory) GetDagIdOk() (*string, bool)`

GetDagIdOk returns a tuple with the DagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagId

`func (o *TaskInstanceHistory) SetDagId(v string)`

SetDagId sets DagId field to given value.

### HasDagId

`func (o *TaskInstanceHistory) HasDagId() bool`

HasDagId returns a boolean if a field has been set.

### GetDagRunId

`func (o *TaskInstanceHistory) GetDagRunId() string`

GetDagRunId returns the DagRunId field if non-nil, zero value otherwise.

### GetDagRunIdOk

`func (o *TaskInstanceHistory) GetDagRunIdOk() (*string, bool)`

GetDagRunIdOk returns a tuple with the DagRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagRunId

`func (o *TaskInstanceHistory) SetDagRunId(v string)`

SetDagRunId sets DagRunId field to given value.

### HasDagRunId

`func (o *TaskInstanceHistory) HasDagRunId() bool`

HasDagRunId returns a boolean if a field has been set.

### GetDuration

`func (o *TaskInstanceHistory) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TaskInstanceHistory) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TaskInstanceHistory) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TaskInstanceHistory) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *TaskInstanceHistory) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *TaskInstanceHistory) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetEndDate

`func (o *TaskInstanceHistory) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TaskInstanceHistory) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TaskInstanceHistory) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TaskInstanceHistory) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *TaskInstanceHistory) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *TaskInstanceHistory) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetExecutor

`func (o *TaskInstanceHistory) GetExecutor() string`

GetExecutor returns the Executor field if non-nil, zero value otherwise.

### GetExecutorOk

`func (o *TaskInstanceHistory) GetExecutorOk() (*string, bool)`

GetExecutorOk returns a tuple with the Executor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutor

`func (o *TaskInstanceHistory) SetExecutor(v string)`

SetExecutor sets Executor field to given value.

### HasExecutor

`func (o *TaskInstanceHistory) HasExecutor() bool`

HasExecutor returns a boolean if a field has been set.

### SetExecutorNil

`func (o *TaskInstanceHistory) SetExecutorNil(b bool)`

 SetExecutorNil sets the value for Executor to be an explicit nil

### UnsetExecutor
`func (o *TaskInstanceHistory) UnsetExecutor()`

UnsetExecutor ensures that no value is present for Executor, not even an explicit nil
### GetExecutorConfig

`func (o *TaskInstanceHistory) GetExecutorConfig() string`

GetExecutorConfig returns the ExecutorConfig field if non-nil, zero value otherwise.

### GetExecutorConfigOk

`func (o *TaskInstanceHistory) GetExecutorConfigOk() (*string, bool)`

GetExecutorConfigOk returns a tuple with the ExecutorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutorConfig

`func (o *TaskInstanceHistory) SetExecutorConfig(v string)`

SetExecutorConfig sets ExecutorConfig field to given value.

### HasExecutorConfig

`func (o *TaskInstanceHistory) HasExecutorConfig() bool`

HasExecutorConfig returns a boolean if a field has been set.

### GetHostname

`func (o *TaskInstanceHistory) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *TaskInstanceHistory) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *TaskInstanceHistory) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *TaskInstanceHistory) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetMapIndex

`func (o *TaskInstanceHistory) GetMapIndex() int32`

GetMapIndex returns the MapIndex field if non-nil, zero value otherwise.

### GetMapIndexOk

`func (o *TaskInstanceHistory) GetMapIndexOk() (*int32, bool)`

GetMapIndexOk returns a tuple with the MapIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapIndex

`func (o *TaskInstanceHistory) SetMapIndex(v int32)`

SetMapIndex sets MapIndex field to given value.

### HasMapIndex

`func (o *TaskInstanceHistory) HasMapIndex() bool`

HasMapIndex returns a boolean if a field has been set.

### GetMaxTries

`func (o *TaskInstanceHistory) GetMaxTries() int32`

GetMaxTries returns the MaxTries field if non-nil, zero value otherwise.

### GetMaxTriesOk

`func (o *TaskInstanceHistory) GetMaxTriesOk() (*int32, bool)`

GetMaxTriesOk returns a tuple with the MaxTries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTries

`func (o *TaskInstanceHistory) SetMaxTries(v int32)`

SetMaxTries sets MaxTries field to given value.

### HasMaxTries

`func (o *TaskInstanceHistory) HasMaxTries() bool`

HasMaxTries returns a boolean if a field has been set.

### GetOperator

`func (o *TaskInstanceHistory) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *TaskInstanceHistory) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *TaskInstanceHistory) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *TaskInstanceHistory) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *TaskInstanceHistory) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *TaskInstanceHistory) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetPid

`func (o *TaskInstanceHistory) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *TaskInstanceHistory) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *TaskInstanceHistory) SetPid(v int32)`

SetPid sets Pid field to given value.

### HasPid

`func (o *TaskInstanceHistory) HasPid() bool`

HasPid returns a boolean if a field has been set.

### SetPidNil

`func (o *TaskInstanceHistory) SetPidNil(b bool)`

 SetPidNil sets the value for Pid to be an explicit nil

### UnsetPid
`func (o *TaskInstanceHistory) UnsetPid()`

UnsetPid ensures that no value is present for Pid, not even an explicit nil
### GetPool

`func (o *TaskInstanceHistory) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *TaskInstanceHistory) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *TaskInstanceHistory) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *TaskInstanceHistory) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolSlots

`func (o *TaskInstanceHistory) GetPoolSlots() int32`

GetPoolSlots returns the PoolSlots field if non-nil, zero value otherwise.

### GetPoolSlotsOk

`func (o *TaskInstanceHistory) GetPoolSlotsOk() (*int32, bool)`

GetPoolSlotsOk returns a tuple with the PoolSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSlots

`func (o *TaskInstanceHistory) SetPoolSlots(v int32)`

SetPoolSlots sets PoolSlots field to given value.

### HasPoolSlots

`func (o *TaskInstanceHistory) HasPoolSlots() bool`

HasPoolSlots returns a boolean if a field has been set.

### GetPriorityWeight

`func (o *TaskInstanceHistory) GetPriorityWeight() int32`

GetPriorityWeight returns the PriorityWeight field if non-nil, zero value otherwise.

### GetPriorityWeightOk

`func (o *TaskInstanceHistory) GetPriorityWeightOk() (*int32, bool)`

GetPriorityWeightOk returns a tuple with the PriorityWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityWeight

`func (o *TaskInstanceHistory) SetPriorityWeight(v int32)`

SetPriorityWeight sets PriorityWeight field to given value.

### HasPriorityWeight

`func (o *TaskInstanceHistory) HasPriorityWeight() bool`

HasPriorityWeight returns a boolean if a field has been set.

### SetPriorityWeightNil

`func (o *TaskInstanceHistory) SetPriorityWeightNil(b bool)`

 SetPriorityWeightNil sets the value for PriorityWeight to be an explicit nil

### UnsetPriorityWeight
`func (o *TaskInstanceHistory) UnsetPriorityWeight()`

UnsetPriorityWeight ensures that no value is present for PriorityWeight, not even an explicit nil
### GetQueue

`func (o *TaskInstanceHistory) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *TaskInstanceHistory) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *TaskInstanceHistory) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *TaskInstanceHistory) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *TaskInstanceHistory) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *TaskInstanceHistory) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetQueuedWhen

`func (o *TaskInstanceHistory) GetQueuedWhen() string`

GetQueuedWhen returns the QueuedWhen field if non-nil, zero value otherwise.

### GetQueuedWhenOk

`func (o *TaskInstanceHistory) GetQueuedWhenOk() (*string, bool)`

GetQueuedWhenOk returns a tuple with the QueuedWhen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedWhen

`func (o *TaskInstanceHistory) SetQueuedWhen(v string)`

SetQueuedWhen sets QueuedWhen field to given value.

### HasQueuedWhen

`func (o *TaskInstanceHistory) HasQueuedWhen() bool`

HasQueuedWhen returns a boolean if a field has been set.

### SetQueuedWhenNil

`func (o *TaskInstanceHistory) SetQueuedWhenNil(b bool)`

 SetQueuedWhenNil sets the value for QueuedWhen to be an explicit nil

### UnsetQueuedWhen
`func (o *TaskInstanceHistory) UnsetQueuedWhen()`

UnsetQueuedWhen ensures that no value is present for QueuedWhen, not even an explicit nil
### GetStartDate

`func (o *TaskInstanceHistory) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TaskInstanceHistory) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TaskInstanceHistory) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TaskInstanceHistory) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *TaskInstanceHistory) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *TaskInstanceHistory) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetState

`func (o *TaskInstanceHistory) GetState() TaskState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TaskInstanceHistory) GetStateOk() (*TaskState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TaskInstanceHistory) SetState(v TaskState)`

SetState sets State field to given value.

### HasState

`func (o *TaskInstanceHistory) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *TaskInstanceHistory) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *TaskInstanceHistory) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetTaskDisplayName

`func (o *TaskInstanceHistory) GetTaskDisplayName() string`

GetTaskDisplayName returns the TaskDisplayName field if non-nil, zero value otherwise.

### GetTaskDisplayNameOk

`func (o *TaskInstanceHistory) GetTaskDisplayNameOk() (*string, bool)`

GetTaskDisplayNameOk returns a tuple with the TaskDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDisplayName

`func (o *TaskInstanceHistory) SetTaskDisplayName(v string)`

SetTaskDisplayName sets TaskDisplayName field to given value.

### HasTaskDisplayName

`func (o *TaskInstanceHistory) HasTaskDisplayName() bool`

HasTaskDisplayName returns a boolean if a field has been set.

### GetTaskId

`func (o *TaskInstanceHistory) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TaskInstanceHistory) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TaskInstanceHistory) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *TaskInstanceHistory) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTryNumber

`func (o *TaskInstanceHistory) GetTryNumber() int32`

GetTryNumber returns the TryNumber field if non-nil, zero value otherwise.

### GetTryNumberOk

`func (o *TaskInstanceHistory) GetTryNumberOk() (*int32, bool)`

GetTryNumberOk returns a tuple with the TryNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryNumber

`func (o *TaskInstanceHistory) SetTryNumber(v int32)`

SetTryNumber sets TryNumber field to given value.

### HasTryNumber

`func (o *TaskInstanceHistory) HasTryNumber() bool`

HasTryNumber returns a boolean if a field has been set.

### GetUnixname

`func (o *TaskInstanceHistory) GetUnixname() string`

GetUnixname returns the Unixname field if non-nil, zero value otherwise.

### GetUnixnameOk

`func (o *TaskInstanceHistory) GetUnixnameOk() (*string, bool)`

GetUnixnameOk returns a tuple with the Unixname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixname

`func (o *TaskInstanceHistory) SetUnixname(v string)`

SetUnixname sets Unixname field to given value.

### HasUnixname

`func (o *TaskInstanceHistory) HasUnixname() bool`

HasUnixname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


