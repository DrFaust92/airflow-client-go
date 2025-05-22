# DagWarningCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalEntries** | Pointer to **int32** | Count of total objects in the current result set before pagination parameters (limit, offset) are applied.  | [optional] 
**DagWarnings** | Pointer to [**[]DagWarning**](DagWarning.md) |  | [optional] 

## Methods

### NewDagWarningCollection

`func NewDagWarningCollection() *DagWarningCollection`

NewDagWarningCollection instantiates a new DagWarningCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDagWarningCollectionWithDefaults

`func NewDagWarningCollectionWithDefaults() *DagWarningCollection`

NewDagWarningCollectionWithDefaults instantiates a new DagWarningCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalEntries

`func (o *DagWarningCollection) GetTotalEntries() int32`

GetTotalEntries returns the TotalEntries field if non-nil, zero value otherwise.

### GetTotalEntriesOk

`func (o *DagWarningCollection) GetTotalEntriesOk() (*int32, bool)`

GetTotalEntriesOk returns a tuple with the TotalEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntries

`func (o *DagWarningCollection) SetTotalEntries(v int32)`

SetTotalEntries sets TotalEntries field to given value.

### HasTotalEntries

`func (o *DagWarningCollection) HasTotalEntries() bool`

HasTotalEntries returns a boolean if a field has been set.

### GetDagWarnings

`func (o *DagWarningCollection) GetDagWarnings() []DagWarning`

GetDagWarnings returns the DagWarnings field if non-nil, zero value otherwise.

### GetDagWarningsOk

`func (o *DagWarningCollection) GetDagWarningsOk() (*[]DagWarning, bool)`

GetDagWarningsOk returns a tuple with the DagWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagWarnings

`func (o *DagWarningCollection) SetDagWarnings(v []DagWarning)`

SetDagWarnings sets DagWarnings field to given value.

### HasDagWarnings

`func (o *DagWarningCollection) HasDagWarnings() bool`

HasDagWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


