# PluginCollectionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppbuilderMenuItems** | Pointer to **[]map[string]interface{}** | The Flask Appbuilder menu items | [optional] 
**AppbuilderViews** | Pointer to **[]map[string]interface{}** | The appuilder views | [optional] 
**Executors** | Pointer to **[]string** | The plugin executors | [optional] 
**FlaskBlueprints** | Pointer to **[]string** | The flask blueprints | [optional] 
**GlobalOperatorExtraLinks** | Pointer to **[]string** | The global operator extra links | [optional] 
**Hooks** | Pointer to **[]string** | The plugin hooks | [optional] 
**Listeners** | Pointer to **[]string** | The plugin listeners | [optional] 
**Macros** | Pointer to **[]string** | The plugin macros | [optional] 
**Name** | Pointer to **string** | The name of the plugin | [optional] 
**OperatorExtraLinks** | Pointer to **[]string** | Operator extra links | [optional] 
**Source** | Pointer to **NullableString** | The plugin source | [optional] 
**TiDeps** | Pointer to **[]string** | The plugin task instance dependencies | [optional] 
**Timetables** | Pointer to **[]string** | The plugin timetables | [optional] 

## Methods

### NewPluginCollectionItem

`func NewPluginCollectionItem() *PluginCollectionItem`

NewPluginCollectionItem instantiates a new PluginCollectionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginCollectionItemWithDefaults

`func NewPluginCollectionItemWithDefaults() *PluginCollectionItem`

NewPluginCollectionItemWithDefaults instantiates a new PluginCollectionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppbuilderMenuItems

`func (o *PluginCollectionItem) GetAppbuilderMenuItems() []*map[string]interface{}`

GetAppbuilderMenuItems returns the AppbuilderMenuItems field if non-nil, zero value otherwise.

### GetAppbuilderMenuItemsOk

`func (o *PluginCollectionItem) GetAppbuilderMenuItemsOk() (*[]*map[string]interface{}, bool)`

GetAppbuilderMenuItemsOk returns a tuple with the AppbuilderMenuItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppbuilderMenuItems

`func (o *PluginCollectionItem) SetAppbuilderMenuItems(v []*map[string]interface{})`

SetAppbuilderMenuItems sets AppbuilderMenuItems field to given value.

### HasAppbuilderMenuItems

`func (o *PluginCollectionItem) HasAppbuilderMenuItems() bool`

HasAppbuilderMenuItems returns a boolean if a field has been set.

### GetAppbuilderViews

`func (o *PluginCollectionItem) GetAppbuilderViews() []*map[string]interface{}`

GetAppbuilderViews returns the AppbuilderViews field if non-nil, zero value otherwise.

### GetAppbuilderViewsOk

`func (o *PluginCollectionItem) GetAppbuilderViewsOk() (*[]*map[string]interface{}, bool)`

GetAppbuilderViewsOk returns a tuple with the AppbuilderViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppbuilderViews

`func (o *PluginCollectionItem) SetAppbuilderViews(v []*map[string]interface{})`

SetAppbuilderViews sets AppbuilderViews field to given value.

### HasAppbuilderViews

`func (o *PluginCollectionItem) HasAppbuilderViews() bool`

HasAppbuilderViews returns a boolean if a field has been set.

### GetExecutors

`func (o *PluginCollectionItem) GetExecutors() []*string`

GetExecutors returns the Executors field if non-nil, zero value otherwise.

### GetExecutorsOk

`func (o *PluginCollectionItem) GetExecutorsOk() (*[]*string, bool)`

GetExecutorsOk returns a tuple with the Executors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutors

`func (o *PluginCollectionItem) SetExecutors(v []*string)`

SetExecutors sets Executors field to given value.

### HasExecutors

`func (o *PluginCollectionItem) HasExecutors() bool`

HasExecutors returns a boolean if a field has been set.

### GetFlaskBlueprints

`func (o *PluginCollectionItem) GetFlaskBlueprints() []*string`

GetFlaskBlueprints returns the FlaskBlueprints field if non-nil, zero value otherwise.

### GetFlaskBlueprintsOk

`func (o *PluginCollectionItem) GetFlaskBlueprintsOk() (*[]*string, bool)`

GetFlaskBlueprintsOk returns a tuple with the FlaskBlueprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlaskBlueprints

`func (o *PluginCollectionItem) SetFlaskBlueprints(v []*string)`

SetFlaskBlueprints sets FlaskBlueprints field to given value.

### HasFlaskBlueprints

`func (o *PluginCollectionItem) HasFlaskBlueprints() bool`

HasFlaskBlueprints returns a boolean if a field has been set.

### GetGlobalOperatorExtraLinks

`func (o *PluginCollectionItem) GetGlobalOperatorExtraLinks() []*string`

GetGlobalOperatorExtraLinks returns the GlobalOperatorExtraLinks field if non-nil, zero value otherwise.

### GetGlobalOperatorExtraLinksOk

`func (o *PluginCollectionItem) GetGlobalOperatorExtraLinksOk() (*[]*string, bool)`

GetGlobalOperatorExtraLinksOk returns a tuple with the GlobalOperatorExtraLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalOperatorExtraLinks

`func (o *PluginCollectionItem) SetGlobalOperatorExtraLinks(v []*string)`

SetGlobalOperatorExtraLinks sets GlobalOperatorExtraLinks field to given value.

### HasGlobalOperatorExtraLinks

`func (o *PluginCollectionItem) HasGlobalOperatorExtraLinks() bool`

HasGlobalOperatorExtraLinks returns a boolean if a field has been set.

### GetHooks

`func (o *PluginCollectionItem) GetHooks() []*string`

GetHooks returns the Hooks field if non-nil, zero value otherwise.

### GetHooksOk

`func (o *PluginCollectionItem) GetHooksOk() (*[]*string, bool)`

GetHooksOk returns a tuple with the Hooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooks

`func (o *PluginCollectionItem) SetHooks(v []*string)`

SetHooks sets Hooks field to given value.

### HasHooks

`func (o *PluginCollectionItem) HasHooks() bool`

HasHooks returns a boolean if a field has been set.

### GetListeners

`func (o *PluginCollectionItem) GetListeners() []string`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *PluginCollectionItem) GetListenersOk() (*[]string, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *PluginCollectionItem) SetListeners(v []string)`

SetListeners sets Listeners field to given value.

### HasListeners

`func (o *PluginCollectionItem) HasListeners() bool`

HasListeners returns a boolean if a field has been set.

### GetMacros

`func (o *PluginCollectionItem) GetMacros() []*string`

GetMacros returns the Macros field if non-nil, zero value otherwise.

### GetMacrosOk

`func (o *PluginCollectionItem) GetMacrosOk() (*[]*string, bool)`

GetMacrosOk returns a tuple with the Macros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacros

`func (o *PluginCollectionItem) SetMacros(v []*string)`

SetMacros sets Macros field to given value.

### HasMacros

`func (o *PluginCollectionItem) HasMacros() bool`

HasMacros returns a boolean if a field has been set.

### GetName

`func (o *PluginCollectionItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PluginCollectionItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PluginCollectionItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PluginCollectionItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperatorExtraLinks

`func (o *PluginCollectionItem) GetOperatorExtraLinks() []*string`

GetOperatorExtraLinks returns the OperatorExtraLinks field if non-nil, zero value otherwise.

### GetOperatorExtraLinksOk

`func (o *PluginCollectionItem) GetOperatorExtraLinksOk() (*[]*string, bool)`

GetOperatorExtraLinksOk returns a tuple with the OperatorExtraLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorExtraLinks

`func (o *PluginCollectionItem) SetOperatorExtraLinks(v []*string)`

SetOperatorExtraLinks sets OperatorExtraLinks field to given value.

### HasOperatorExtraLinks

`func (o *PluginCollectionItem) HasOperatorExtraLinks() bool`

HasOperatorExtraLinks returns a boolean if a field has been set.

### GetSource

`func (o *PluginCollectionItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PluginCollectionItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PluginCollectionItem) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PluginCollectionItem) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *PluginCollectionItem) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *PluginCollectionItem) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetTiDeps

`func (o *PluginCollectionItem) GetTiDeps() []string`

GetTiDeps returns the TiDeps field if non-nil, zero value otherwise.

### GetTiDepsOk

`func (o *PluginCollectionItem) GetTiDepsOk() (*[]string, bool)`

GetTiDepsOk returns a tuple with the TiDeps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiDeps

`func (o *PluginCollectionItem) SetTiDeps(v []string)`

SetTiDeps sets TiDeps field to given value.

### HasTiDeps

`func (o *PluginCollectionItem) HasTiDeps() bool`

HasTiDeps returns a boolean if a field has been set.

### GetTimetables

`func (o *PluginCollectionItem) GetTimetables() []string`

GetTimetables returns the Timetables field if non-nil, zero value otherwise.

### GetTimetablesOk

`func (o *PluginCollectionItem) GetTimetablesOk() (*[]string, bool)`

GetTimetablesOk returns a tuple with the Timetables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimetables

`func (o *PluginCollectionItem) SetTimetables(v []string)`

SetTimetables sets Timetables field to given value.

### HasTimetables

`func (o *PluginCollectionItem) HasTimetables() bool`

HasTimetables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


