# TriggerCustomEffectBlockTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The scope the custom effect applies to: - &#x60;cart&#x60; applies once to the whole cart. - &#x60;allItems&#x60; applies once per cart item. - &#x60;selector&#x60; applies once per item matched by the named selector. - &#x60;globalFilter&#x60; applies once per item matched by the named global item filter. - &#x60;bundle&#x60; applies once per item in the named bundle. | 
**Name** | Pointer to **string** | The name of the targeted selector or bundle. Only set when &#x60;type&#x60; is &#x60;selector&#x60;, &#x60;globalFilter&#x60;, or &#x60;bundle&#x60;. | [optional] 

## Methods

### NewTriggerCustomEffectBlockTarget

`func NewTriggerCustomEffectBlockTarget(type_ string, ) *TriggerCustomEffectBlockTarget`

NewTriggerCustomEffectBlockTarget instantiates a new TriggerCustomEffectBlockTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerCustomEffectBlockTargetWithDefaults

`func NewTriggerCustomEffectBlockTargetWithDefaults() *TriggerCustomEffectBlockTarget`

NewTriggerCustomEffectBlockTargetWithDefaults instantiates a new TriggerCustomEffectBlockTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TriggerCustomEffectBlockTarget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TriggerCustomEffectBlockTarget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TriggerCustomEffectBlockTarget) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *TriggerCustomEffectBlockTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TriggerCustomEffectBlockTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TriggerCustomEffectBlockTarget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TriggerCustomEffectBlockTarget) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


