# BetweenCheckAttributeBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** | The range comparison operator. Must be &#x60;between&#x60;. | [optional] 
**Min** | Pointer to [**map[string]interface{}**](.md) | The minimum value allowed for the &#x60;between&#x60; operator. | 
**Max** | Pointer to [**map[string]interface{}**](.md) | The maximum value allowed for the &#x60;between&#x60; operator. | 

## Methods

### NewBetweenCheckAttributeBlock

`func NewBetweenCheckAttributeBlock(min map[string]interface{}, max map[string]interface{}, ) *BetweenCheckAttributeBlock`

NewBetweenCheckAttributeBlock instantiates a new BetweenCheckAttributeBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetweenCheckAttributeBlockWithDefaults

`func NewBetweenCheckAttributeBlockWithDefaults() *BetweenCheckAttributeBlock`

NewBetweenCheckAttributeBlockWithDefaults instantiates a new BetweenCheckAttributeBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *BetweenCheckAttributeBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *BetweenCheckAttributeBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *BetweenCheckAttributeBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *BetweenCheckAttributeBlock) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetMin

`func (o *BetweenCheckAttributeBlock) GetMin() map[string]interface{}`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *BetweenCheckAttributeBlock) GetMinOk() (*map[string]interface{}, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *BetweenCheckAttributeBlock) SetMin(v map[string]interface{})`

SetMin sets Min field to given value.


### GetMax

`func (o *BetweenCheckAttributeBlock) GetMax() map[string]interface{}`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *BetweenCheckAttributeBlock) GetMaxOk() (*map[string]interface{}, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *BetweenCheckAttributeBlock) SetMax(v map[string]interface{})`

SetMax sets Max field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


