# ClearDagRun200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conf** | Pointer to **map[string]interface{}** | JSON object describing additional configuration parameters.  The value of this field can be set only when creating the object. If you try to modify the field of an existing object, the request fails with an BAD_REQUEST error.  | [optional] 
**DagId** | Pointer to **string** |  | [optional] [readonly] 
**DagRunId** | Pointer to **NullableString** | Run ID.  The value of this field can be set only when creating the object. If you try to modify the field of an existing object, the request fails with an BAD_REQUEST error.  If not provided, a value will be generated based on execution_date.  If the specified dag_run_id is in use, the creation request fails with an ALREADY_EXISTS error.  This together with DAG_ID are a unique key.  | [optional] 
**DataIntervalEnd** | Pointer to **NullableTime** | The end of the interval the DAG run covers.  | [optional] 
**DataIntervalStart** | Pointer to **NullableTime** | The beginning of the interval the DAG run covers.  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] [readonly] 
**ExecutionDate** | Pointer to **NullableTime** | The execution date. This is the same as logical_date, kept for backwards compatibility. If both this field and logical_date are provided but with different values, the request will fail with an BAD_REQUEST error.  *Changed in version 2.2.0*&amp;#58; Field becomes nullable.  *Deprecated since version 2.2.0*&amp;#58; Use &#39;logical_date&#39; instead.  | [optional] 
**ExternalTrigger** | Pointer to **bool** |  | [optional] [readonly] 
**LastSchedulingDecision** | Pointer to **NullableTime** |  | [optional] [readonly] 
**LogicalDate** | Pointer to **NullableTime** | The logical date (previously called execution date). This is the time or interval covered by this DAG run, according to the DAG definition.  The value of this field can be set only when creating the object. If you try to modify the field of an existing object, the request fails with an BAD_REQUEST error.  This together with DAG_ID are a unique key.  *New in version 2.2.0*  | [optional] 
**Note** | Pointer to **NullableString** | Contains manually entered notes by the user about the DagRun.  *New in version 2.5.0*  | [optional] 
**RunType** | Pointer to **string** |  | [optional] [readonly] 
**StartDate** | Pointer to **NullableTime** | The start time. The time when DAG run was actually created.  *Changed in version 2.1.3*&amp;#58; Field becomes nullable.  | [optional] [readonly] 
**State** | Pointer to [**DagState**](DagState.md) |  | [optional] 
**TaskInstances** | Pointer to [**[]TaskInstance**](TaskInstance.md) |  | [optional] 
**TotalEntries** | Pointer to **int32** | Count of total objects in the current result set before pagination parameters (limit, offset) are applied.  | [optional] 

## Methods

### NewClearDagRun200Response

`func NewClearDagRun200Response() *ClearDagRun200Response`

NewClearDagRun200Response instantiates a new ClearDagRun200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearDagRun200ResponseWithDefaults

`func NewClearDagRun200ResponseWithDefaults() *ClearDagRun200Response`

NewClearDagRun200ResponseWithDefaults instantiates a new ClearDagRun200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConf

`func (o *ClearDagRun200Response) GetConf() map[string]interface{}`

GetConf returns the Conf field if non-nil, zero value otherwise.

### GetConfOk

`func (o *ClearDagRun200Response) GetConfOk() (*map[string]interface{}, bool)`

GetConfOk returns a tuple with the Conf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConf

`func (o *ClearDagRun200Response) SetConf(v map[string]interface{})`

SetConf sets Conf field to given value.

### HasConf

`func (o *ClearDagRun200Response) HasConf() bool`

HasConf returns a boolean if a field has been set.

### GetDagId

`func (o *ClearDagRun200Response) GetDagId() string`

GetDagId returns the DagId field if non-nil, zero value otherwise.

### GetDagIdOk

`func (o *ClearDagRun200Response) GetDagIdOk() (*string, bool)`

GetDagIdOk returns a tuple with the DagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagId

`func (o *ClearDagRun200Response) SetDagId(v string)`

SetDagId sets DagId field to given value.

### HasDagId

`func (o *ClearDagRun200Response) HasDagId() bool`

HasDagId returns a boolean if a field has been set.

### GetDagRunId

`func (o *ClearDagRun200Response) GetDagRunId() string`

GetDagRunId returns the DagRunId field if non-nil, zero value otherwise.

### GetDagRunIdOk

`func (o *ClearDagRun200Response) GetDagRunIdOk() (*string, bool)`

GetDagRunIdOk returns a tuple with the DagRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagRunId

`func (o *ClearDagRun200Response) SetDagRunId(v string)`

SetDagRunId sets DagRunId field to given value.

### HasDagRunId

`func (o *ClearDagRun200Response) HasDagRunId() bool`

HasDagRunId returns a boolean if a field has been set.

### SetDagRunIdNil

`func (o *ClearDagRun200Response) SetDagRunIdNil(b bool)`

 SetDagRunIdNil sets the value for DagRunId to be an explicit nil

### UnsetDagRunId
`func (o *ClearDagRun200Response) UnsetDagRunId()`

UnsetDagRunId ensures that no value is present for DagRunId, not even an explicit nil
### GetDataIntervalEnd

`func (o *ClearDagRun200Response) GetDataIntervalEnd() time.Time`

GetDataIntervalEnd returns the DataIntervalEnd field if non-nil, zero value otherwise.

### GetDataIntervalEndOk

`func (o *ClearDagRun200Response) GetDataIntervalEndOk() (*time.Time, bool)`

GetDataIntervalEndOk returns a tuple with the DataIntervalEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIntervalEnd

`func (o *ClearDagRun200Response) SetDataIntervalEnd(v time.Time)`

SetDataIntervalEnd sets DataIntervalEnd field to given value.

### HasDataIntervalEnd

`func (o *ClearDagRun200Response) HasDataIntervalEnd() bool`

HasDataIntervalEnd returns a boolean if a field has been set.

### SetDataIntervalEndNil

`func (o *ClearDagRun200Response) SetDataIntervalEndNil(b bool)`

 SetDataIntervalEndNil sets the value for DataIntervalEnd to be an explicit nil

### UnsetDataIntervalEnd
`func (o *ClearDagRun200Response) UnsetDataIntervalEnd()`

UnsetDataIntervalEnd ensures that no value is present for DataIntervalEnd, not even an explicit nil
### GetDataIntervalStart

`func (o *ClearDagRun200Response) GetDataIntervalStart() time.Time`

GetDataIntervalStart returns the DataIntervalStart field if non-nil, zero value otherwise.

### GetDataIntervalStartOk

`func (o *ClearDagRun200Response) GetDataIntervalStartOk() (*time.Time, bool)`

GetDataIntervalStartOk returns a tuple with the DataIntervalStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIntervalStart

`func (o *ClearDagRun200Response) SetDataIntervalStart(v time.Time)`

SetDataIntervalStart sets DataIntervalStart field to given value.

### HasDataIntervalStart

`func (o *ClearDagRun200Response) HasDataIntervalStart() bool`

HasDataIntervalStart returns a boolean if a field has been set.

### SetDataIntervalStartNil

`func (o *ClearDagRun200Response) SetDataIntervalStartNil(b bool)`

 SetDataIntervalStartNil sets the value for DataIntervalStart to be an explicit nil

### UnsetDataIntervalStart
`func (o *ClearDagRun200Response) UnsetDataIntervalStart()`

UnsetDataIntervalStart ensures that no value is present for DataIntervalStart, not even an explicit nil
### GetEndDate

`func (o *ClearDagRun200Response) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ClearDagRun200Response) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ClearDagRun200Response) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ClearDagRun200Response) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *ClearDagRun200Response) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *ClearDagRun200Response) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetExecutionDate

`func (o *ClearDagRun200Response) GetExecutionDate() time.Time`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *ClearDagRun200Response) GetExecutionDateOk() (*time.Time, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *ClearDagRun200Response) SetExecutionDate(v time.Time)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *ClearDagRun200Response) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### SetExecutionDateNil

`func (o *ClearDagRun200Response) SetExecutionDateNil(b bool)`

 SetExecutionDateNil sets the value for ExecutionDate to be an explicit nil

### UnsetExecutionDate
`func (o *ClearDagRun200Response) UnsetExecutionDate()`

UnsetExecutionDate ensures that no value is present for ExecutionDate, not even an explicit nil
### GetExternalTrigger

`func (o *ClearDagRun200Response) GetExternalTrigger() bool`

GetExternalTrigger returns the ExternalTrigger field if non-nil, zero value otherwise.

### GetExternalTriggerOk

`func (o *ClearDagRun200Response) GetExternalTriggerOk() (*bool, bool)`

GetExternalTriggerOk returns a tuple with the ExternalTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTrigger

`func (o *ClearDagRun200Response) SetExternalTrigger(v bool)`

SetExternalTrigger sets ExternalTrigger field to given value.

### HasExternalTrigger

`func (o *ClearDagRun200Response) HasExternalTrigger() bool`

HasExternalTrigger returns a boolean if a field has been set.

### GetLastSchedulingDecision

`func (o *ClearDagRun200Response) GetLastSchedulingDecision() time.Time`

GetLastSchedulingDecision returns the LastSchedulingDecision field if non-nil, zero value otherwise.

### GetLastSchedulingDecisionOk

`func (o *ClearDagRun200Response) GetLastSchedulingDecisionOk() (*time.Time, bool)`

GetLastSchedulingDecisionOk returns a tuple with the LastSchedulingDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSchedulingDecision

`func (o *ClearDagRun200Response) SetLastSchedulingDecision(v time.Time)`

SetLastSchedulingDecision sets LastSchedulingDecision field to given value.

### HasLastSchedulingDecision

`func (o *ClearDagRun200Response) HasLastSchedulingDecision() bool`

HasLastSchedulingDecision returns a boolean if a field has been set.

### SetLastSchedulingDecisionNil

`func (o *ClearDagRun200Response) SetLastSchedulingDecisionNil(b bool)`

 SetLastSchedulingDecisionNil sets the value for LastSchedulingDecision to be an explicit nil

### UnsetLastSchedulingDecision
`func (o *ClearDagRun200Response) UnsetLastSchedulingDecision()`

UnsetLastSchedulingDecision ensures that no value is present for LastSchedulingDecision, not even an explicit nil
### GetLogicalDate

`func (o *ClearDagRun200Response) GetLogicalDate() time.Time`

GetLogicalDate returns the LogicalDate field if non-nil, zero value otherwise.

### GetLogicalDateOk

`func (o *ClearDagRun200Response) GetLogicalDateOk() (*time.Time, bool)`

GetLogicalDateOk returns a tuple with the LogicalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalDate

`func (o *ClearDagRun200Response) SetLogicalDate(v time.Time)`

SetLogicalDate sets LogicalDate field to given value.

### HasLogicalDate

`func (o *ClearDagRun200Response) HasLogicalDate() bool`

HasLogicalDate returns a boolean if a field has been set.

### SetLogicalDateNil

`func (o *ClearDagRun200Response) SetLogicalDateNil(b bool)`

 SetLogicalDateNil sets the value for LogicalDate to be an explicit nil

### UnsetLogicalDate
`func (o *ClearDagRun200Response) UnsetLogicalDate()`

UnsetLogicalDate ensures that no value is present for LogicalDate, not even an explicit nil
### GetNote

`func (o *ClearDagRun200Response) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ClearDagRun200Response) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ClearDagRun200Response) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ClearDagRun200Response) HasNote() bool`

HasNote returns a boolean if a field has been set.

### SetNoteNil

`func (o *ClearDagRun200Response) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *ClearDagRun200Response) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetRunType

`func (o *ClearDagRun200Response) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *ClearDagRun200Response) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *ClearDagRun200Response) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *ClearDagRun200Response) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### GetStartDate

`func (o *ClearDagRun200Response) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ClearDagRun200Response) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ClearDagRun200Response) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ClearDagRun200Response) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *ClearDagRun200Response) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *ClearDagRun200Response) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetState

`func (o *ClearDagRun200Response) GetState() DagState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ClearDagRun200Response) GetStateOk() (*DagState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ClearDagRun200Response) SetState(v DagState)`

SetState sets State field to given value.

### HasState

`func (o *ClearDagRun200Response) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTaskInstances

`func (o *ClearDagRun200Response) GetTaskInstances() []TaskInstance`

GetTaskInstances returns the TaskInstances field if non-nil, zero value otherwise.

### GetTaskInstancesOk

`func (o *ClearDagRun200Response) GetTaskInstancesOk() (*[]TaskInstance, bool)`

GetTaskInstancesOk returns a tuple with the TaskInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskInstances

`func (o *ClearDagRun200Response) SetTaskInstances(v []TaskInstance)`

SetTaskInstances sets TaskInstances field to given value.

### HasTaskInstances

`func (o *ClearDagRun200Response) HasTaskInstances() bool`

HasTaskInstances returns a boolean if a field has been set.

### GetTotalEntries

`func (o *ClearDagRun200Response) GetTotalEntries() int32`

GetTotalEntries returns the TotalEntries field if non-nil, zero value otherwise.

### GetTotalEntriesOk

`func (o *ClearDagRun200Response) GetTotalEntriesOk() (*int32, bool)`

GetTotalEntriesOk returns a tuple with the TotalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntries

`func (o *ClearDagRun200Response) SetTotalEntries(v int32)`

SetTotalEntries sets TotalEntries field to given value.

### HasTotalEntries

`func (o *ClearDagRun200Response) HasTotalEntries() bool`

HasTotalEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


