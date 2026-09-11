# TriggerCustomEffectBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**CustomEffect** | Pointer to [**TriggerCustomEffectBlockCustomEffect**](TriggerCustomEffectBlock_customEffect.md) |  | 
**Params** | Pointer to [**map[string]interface{}**](.md) | The custom effect&#39;s parameters, in configured order. Each property name is the parameter&#39;s title, lowercased with spaces replaced by underscores (for example, &#x60;Order ID&#x60; becomes &#x60;order_id&#x60;); falls back to &#x60;param_0&#x60;, &#x60;param_1&#x60;, and so on if a title is blank or collides with another. | [optional] 
**Target** | Pointer to [**TriggerCustomEffectBlockTarget**](TriggerCustomEffectBlock_target.md) |  | 
**OnError** | Pointer to [**map[string][]map[string]interface{}**](array.md) | Named error handlers evaluated when a specific error occurs. | [optional] 

## Methods

### NewTriggerCustomEffectBlock

`func NewTriggerCustomEffectBlock(type_ string, customEffect TriggerCustomEffectBlockCustomEffect, target TriggerCustomEffectBlockTarget, ) *TriggerCustomEffectBlock`

NewTriggerCustomEffectBlock instantiates a new TriggerCustomEffectBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerCustomEffectBlockWithDefaults

`func NewTriggerCustomEffectBlockWithDefaults() *TriggerCustomEffectBlock`

NewTriggerCustomEffectBlockWithDefaults instantiates a new TriggerCustomEffectBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TriggerCustomEffectBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TriggerCustomEffectBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TriggerCustomEffectBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TriggerCustomEffectBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TriggerCustomEffectBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TriggerCustomEffectBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TriggerCustomEffectBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *TriggerCustomEffectBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TriggerCustomEffectBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TriggerCustomEffectBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TriggerCustomEffectBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomEffect

`func (o *TriggerCustomEffectBlock) GetCustomEffect() TriggerCustomEffectBlockCustomEffect`

GetCustomEffect returns the CustomEffect field if non-nil, zero value otherwise.

### GetCustomEffectOk

`func (o *TriggerCustomEffectBlock) GetCustomEffectOk() (*TriggerCustomEffectBlockCustomEffect, bool)`

GetCustomEffectOk returns a tuple with the CustomEffect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEffect

`func (o *TriggerCustomEffectBlock) SetCustomEffect(v TriggerCustomEffectBlockCustomEffect)`

SetCustomEffect sets CustomEffect field to given value.


### GetParams

`func (o *TriggerCustomEffectBlock) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *TriggerCustomEffectBlock) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *TriggerCustomEffectBlock) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *TriggerCustomEffectBlock) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetTarget

`func (o *TriggerCustomEffectBlock) GetTarget() TriggerCustomEffectBlockTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *TriggerCustomEffectBlock) GetTargetOk() (*TriggerCustomEffectBlockTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *TriggerCustomEffectBlock) SetTarget(v TriggerCustomEffectBlockTarget)`

SetTarget sets Target field to given value.


### GetOnError

`func (o *TriggerCustomEffectBlock) GetOnError() map[string][]map[string]interface{}`

GetOnError returns the OnError field if non-nil, zero value otherwise.

### GetOnErrorOk

`func (o *TriggerCustomEffectBlock) GetOnErrorOk() (*map[string][]map[string]interface{}, bool)`

GetOnErrorOk returns a tuple with the OnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnError

`func (o *TriggerCustomEffectBlock) SetOnError(v map[string][]map[string]interface{})`

SetOnError sets OnError field to given value.

### HasOnError

`func (o *TriggerCustomEffectBlock) HasOnError() bool`

HasOnError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


