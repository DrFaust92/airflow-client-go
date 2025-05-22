# TaskInstanceDependencyCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dependencies** | Pointer to [**[]TaskFailedDependency**](TaskFailedDependency.md) |  | [optional] 

## Methods

### NewTaskInstanceDependencyCollection

`func NewTaskInstanceDependencyCollection() *TaskInstanceDependencyCollection`

NewTaskInstanceDependencyCollection instantiates a new TaskInstanceDependencyCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskInstanceDependencyCollectionWithDefaults

`func NewTaskInstanceDependencyCollectionWithDefaults() *TaskInstanceDependencyCollection`

NewTaskInstanceDependencyCollectionWithDefaults instantiates a new TaskInstanceDependencyCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependencies

`func (o *TaskInstanceDependencyCollection) GetDependencies() []TaskFailedDependency`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *TaskInstanceDependencyCollection) GetDependenciesOk() (*[]TaskFailedDependency, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *TaskInstanceDependencyCollection) SetDependencies(v []TaskFailedDependency)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *TaskInstanceDependencyCollection) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


