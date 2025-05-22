# DagStatsCollectionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DagId** | Pointer to **string** | The DAG ID. | [optional] 
**Stats** | Pointer to [**[]DagStatsStateCollectionItem**](DagStatsStateCollectionItem.md) |  | [optional] 

## Methods

### NewDagStatsCollectionItem

`func NewDagStatsCollectionItem() *DagStatsCollectionItem`

NewDagStatsCollectionItem instantiates a new DagStatsCollectionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDagStatsCollectionItemWithDefaults

`func NewDagStatsCollectionItemWithDefaults() *DagStatsCollectionItem`

NewDagStatsCollectionItemWithDefaults instantiates a new DagStatsCollectionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDagId

`func (o *DagStatsCollectionItem) GetDagId() string`

GetDagId returns the DagId field if non-nil, zero value otherwise.

### GetDagIdOk

`func (o *DagStatsCollectionItem) GetDagIdOk() (*string, bool)`

GetDagIdOk returns a tuple with the DagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagId

`func (o *DagStatsCollectionItem) SetDagId(v string)`

SetDagId sets DagId field to given value.

### HasDagId

`func (o *DagStatsCollectionItem) HasDagId() bool`

HasDagId returns a boolean if a field has been set.

### GetStats

`func (o *DagStatsCollectionItem) GetStats() []DagStatsStateCollectionItem`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *DagStatsCollectionItem) GetStatsOk() (*[]DagStatsStateCollectionItem, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *DagStatsCollectionItem) SetStats(v []DagStatsStateCollectionItem)`

SetStats sets Stats field to given value.

### HasStats

`func (o *DagStatsCollectionItem) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStatsNil

`func (o *DagStatsCollectionItem) SetStatsNil(b bool)`

 SetStatsNil sets the value for Stats to be an explicit nil

### UnsetStats
`func (o *DagStatsCollectionItem) UnsetStats()`

UnsetStats ensures that no value is present for Stats, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


