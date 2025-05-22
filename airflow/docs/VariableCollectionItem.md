# VariableCollectionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | The description of the variable.  *New in version 2.4.0*  | [optional] 
**Key** | Pointer to **string** |  | [optional] 

## Methods

### NewVariableCollectionItem

`func NewVariableCollectionItem() *VariableCollectionItem`

NewVariableCollectionItem instantiates a new VariableCollectionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableCollectionItemWithDefaults

`func NewVariableCollectionItemWithDefaults() *VariableCollectionItem`

NewVariableCollectionItemWithDefaults instantiates a new VariableCollectionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *VariableCollectionItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VariableCollectionItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VariableCollectionItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VariableCollectionItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VariableCollectionItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VariableCollectionItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetKey

`func (o *VariableCollectionItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VariableCollectionItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VariableCollectionItem) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *VariableCollectionItem) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


