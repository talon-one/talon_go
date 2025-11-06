# HiddenConditionsEffects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuiltInEffects** | Pointer to **[]string** | List of hidden built-in effects. | [optional] 
**Conditions** | Pointer to **[]string** | List of hidden conditions. | [optional] 
**CustomEffects** | Pointer to **[]int64** | List of the IDs of hidden custom effects. | [optional] 
**Webhooks** | Pointer to **[]int64** | List of the IDs of hidden webhooks. | [optional] 

## Methods

### NewHiddenConditionsEffects

`func NewHiddenConditionsEffects() *HiddenConditionsEffects`

NewHiddenConditionsEffects instantiates a new HiddenConditionsEffects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiddenConditionsEffectsWithDefaults

`func NewHiddenConditionsEffectsWithDefaults() *HiddenConditionsEffects`

NewHiddenConditionsEffectsWithDefaults instantiates a new HiddenConditionsEffects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuiltInEffects

`func (o *HiddenConditionsEffects) GetBuiltInEffects() []string`

GetBuiltInEffects returns the BuiltInEffects field if non-nil, zero value otherwise.

### GetBuiltInEffectsOk

`func (o *HiddenConditionsEffects) GetBuiltInEffectsOk() (*[]string, bool)`

GetBuiltInEffectsOk returns a tuple with the BuiltInEffects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltInEffects

`func (o *HiddenConditionsEffects) SetBuiltInEffects(v []string)`

SetBuiltInEffects sets BuiltInEffects field to given value.

### HasBuiltInEffects

`func (o *HiddenConditionsEffects) HasBuiltInEffects() bool`

HasBuiltInEffects returns a boolean if a field has been set.

### GetConditions

`func (o *HiddenConditionsEffects) GetConditions() []string`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *HiddenConditionsEffects) GetConditionsOk() (*[]string, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *HiddenConditionsEffects) SetConditions(v []string)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *HiddenConditionsEffects) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCustomEffects

`func (o *HiddenConditionsEffects) GetCustomEffects() []int64`

GetCustomEffects returns the CustomEffects field if non-nil, zero value otherwise.

### GetCustomEffectsOk

`func (o *HiddenConditionsEffects) GetCustomEffectsOk() (*[]int64, bool)`

GetCustomEffectsOk returns a tuple with the CustomEffects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEffects

`func (o *HiddenConditionsEffects) SetCustomEffects(v []int64)`

SetCustomEffects sets CustomEffects field to given value.

### HasCustomEffects

`func (o *HiddenConditionsEffects) HasCustomEffects() bool`

HasCustomEffects returns a boolean if a field has been set.

### GetWebhooks

`func (o *HiddenConditionsEffects) GetWebhooks() []int64`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *HiddenConditionsEffects) GetWebhooksOk() (*[]int64, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *HiddenConditionsEffects) SetWebhooks(v []int64)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *HiddenConditionsEffects) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


