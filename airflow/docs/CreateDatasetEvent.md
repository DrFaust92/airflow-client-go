# CreateDatasetEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetUri** | **string** | The URI of the dataset | 
**Extra** | Pointer to **map[string]interface{}** | The dataset event extra | [optional] 

## Methods

### NewCreateDatasetEvent

`func NewCreateDatasetEvent(datasetUri string, ) *CreateDatasetEvent`

NewCreateDatasetEvent instantiates a new CreateDatasetEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDatasetEventWithDefaults

`func NewCreateDatasetEventWithDefaults() *CreateDatasetEvent`

NewCreateDatasetEventWithDefaults instantiates a new CreateDatasetEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetUri

`func (o *CreateDatasetEvent) GetDatasetUri() string`

GetDatasetUri returns the DatasetUri field if non-nil, zero value otherwise.

### GetDatasetUriOk

`func (o *CreateDatasetEvent) GetDatasetUriOk() (*string, bool)`

GetDatasetUriOk returns a tuple with the DatasetUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetUri

`func (o *CreateDatasetEvent) SetDatasetUri(v string)`

SetDatasetUri sets DatasetUri field to given value.


### GetExtra

`func (o *CreateDatasetEvent) GetExtra() map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *CreateDatasetEvent) GetExtraOk() (*map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *CreateDatasetEvent) SetExtra(v map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *CreateDatasetEvent) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### SetExtraNil

`func (o *CreateDatasetEvent) SetExtraNil(b bool)`

 SetExtraNil sets the value for Extra to be an explicit nil

### UnsetExtra
`func (o *CreateDatasetEvent) UnsetExtra()`

UnsetExtra ensures that no value is present for Extra, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


