# QueuedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | The creation time of QueuedEvent | [optional] 
**DagId** | Pointer to **string** | The DAG ID. | [optional] 
**Uri** | Pointer to **string** | The datata uri. | [optional] 

## Methods

### NewQueuedEvent

`func NewQueuedEvent() *QueuedEvent`

NewQueuedEvent instantiates a new QueuedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueuedEventWithDefaults

`func NewQueuedEventWithDefaults() *QueuedEvent`

NewQueuedEventWithDefaults instantiates a new QueuedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *QueuedEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QueuedEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QueuedEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *QueuedEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDagId

`func (o *QueuedEvent) GetDagId() string`

GetDagId returns the DagId field if non-nil, zero value otherwise.

### GetDagIdOk

`func (o *QueuedEvent) GetDagIdOk() (*string, bool)`

GetDagIdOk returns a tuple with the DagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagId

`func (o *QueuedEvent) SetDagId(v string)`

SetDagId sets DagId field to given value.

### HasDagId

`func (o *QueuedEvent) HasDagId() bool`

HasDagId returns a boolean if a field has been set.

### GetUri

`func (o *QueuedEvent) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *QueuedEvent) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *QueuedEvent) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *QueuedEvent) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


