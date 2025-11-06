# StrikethroughCustomEffectPerItemProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectId** | Pointer to **int64** | ID of the effect. | 
**Name** | Pointer to **string** | The type of the custom effect. | 
**Payload** | Pointer to [**map[string]interface{}**](.md) | The JSON payload of the custom effect. | 

## Methods

### NewStrikethroughCustomEffectPerItemProps

`func NewStrikethroughCustomEffectPerItemProps(effectId int64, name string, payload map[string]interface{}, ) *StrikethroughCustomEffectPerItemProps`

NewStrikethroughCustomEffectPerItemProps instantiates a new StrikethroughCustomEffectPerItemProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrikethroughCustomEffectPerItemPropsWithDefaults

`func NewStrikethroughCustomEffectPerItemPropsWithDefaults() *StrikethroughCustomEffectPerItemProps`

NewStrikethroughCustomEffectPerItemPropsWithDefaults instantiates a new StrikethroughCustomEffectPerItemProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffectId

`func (o *StrikethroughCustomEffectPerItemProps) GetEffectId() int64`

GetEffectId returns the EffectId field if non-nil, zero value otherwise.

### GetEffectIdOk

`func (o *StrikethroughCustomEffectPerItemProps) GetEffectIdOk() (*int64, bool)`

GetEffectIdOk returns a tuple with the EffectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectId

`func (o *StrikethroughCustomEffectPerItemProps) SetEffectId(v int64)`

SetEffectId sets EffectId field to given value.


### GetName

`func (o *StrikethroughCustomEffectPerItemProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StrikethroughCustomEffectPerItemProps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StrikethroughCustomEffectPerItemProps) SetName(v string)`

SetName sets Name field to given value.


### GetPayload

`func (o *StrikethroughCustomEffectPerItemProps) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *StrikethroughCustomEffectPerItemProps) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *StrikethroughCustomEffectPerItemProps) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


