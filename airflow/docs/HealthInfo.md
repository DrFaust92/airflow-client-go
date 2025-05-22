# HealthInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DagProcessor** | Pointer to [**DagProcessorStatus**](DagProcessorStatus.md) |  | [optional] 
**Metadatabase** | Pointer to [**MetadatabaseStatus**](MetadatabaseStatus.md) |  | [optional] 
**Scheduler** | Pointer to [**SchedulerStatus**](SchedulerStatus.md) |  | [optional] 
**Triggerer** | Pointer to [**TriggererStatus**](TriggererStatus.md) |  | [optional] 

## Methods

### NewHealthInfo

`func NewHealthInfo() *HealthInfo`

NewHealthInfo instantiates a new HealthInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthInfoWithDefaults

`func NewHealthInfoWithDefaults() *HealthInfo`

NewHealthInfoWithDefaults instantiates a new HealthInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDagProcessor

`func (o *HealthInfo) GetDagProcessor() DagProcessorStatus`

GetDagProcessor returns the DagProcessor field if non-nil, zero value otherwise.

### GetDagProcessorOk

`func (o *HealthInfo) GetDagProcessorOk() (*DagProcessorStatus, bool)`

GetDagProcessorOk returns a tuple with the DagProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagProcessor

`func (o *HealthInfo) SetDagProcessor(v DagProcessorStatus)`

SetDagProcessor sets DagProcessor field to given value.

### HasDagProcessor

`func (o *HealthInfo) HasDagProcessor() bool`

HasDagProcessor returns a boolean if a field has been set.

### GetMetadatabase

`func (o *HealthInfo) GetMetadatabase() MetadatabaseStatus`

GetMetadatabase returns the Metadatabase field if non-nil, zero value otherwise.

### GetMetadatabaseOk

`func (o *HealthInfo) GetMetadatabaseOk() (*MetadatabaseStatus, bool)`

GetMetadatabaseOk returns a tuple with the Metadatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadatabase

`func (o *HealthInfo) SetMetadatabase(v MetadatabaseStatus)`

SetMetadatabase sets Metadatabase field to given value.

### HasMetadatabase

`func (o *HealthInfo) HasMetadatabase() bool`

HasMetadatabase returns a boolean if a field has been set.

### GetScheduler

`func (o *HealthInfo) GetScheduler() SchedulerStatus`

GetScheduler returns the Scheduler field if non-nil, zero value otherwise.

### GetSchedulerOk

`func (o *HealthInfo) GetSchedulerOk() (*SchedulerStatus, bool)`

GetSchedulerOk returns a tuple with the Scheduler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduler

`func (o *HealthInfo) SetScheduler(v SchedulerStatus)`

SetScheduler sets Scheduler field to given value.

### HasScheduler

`func (o *HealthInfo) HasScheduler() bool`

HasScheduler returns a boolean if a field has been set.

### GetTriggerer

`func (o *HealthInfo) GetTriggerer() TriggererStatus`

GetTriggerer returns the Triggerer field if non-nil, zero value otherwise.

### GetTriggererOk

`func (o *HealthInfo) GetTriggererOk() (*TriggererStatus, bool)`

GetTriggererOk returns a tuple with the Triggerer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerer

`func (o *HealthInfo) SetTriggerer(v TriggererStatus)`

SetTriggerer sets Triggerer field to given value.

### HasTriggerer

`func (o *HealthInfo) HasTriggerer() bool`

HasTriggerer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


