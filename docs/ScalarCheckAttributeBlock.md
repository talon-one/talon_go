# ScalarCheckAttributeBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** | The comparison operator applied to the attribute. | [optional] 
**Value** | Pointer to [**map[string]interface{}**](.md) | The comparison value for this operator. | 

## Methods

### NewScalarCheckAttributeBlock

`func NewScalarCheckAttributeBlock(value map[string]interface{}, ) *ScalarCheckAttributeBlock`

NewScalarCheckAttributeBlock instantiates a new ScalarCheckAttributeBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScalarCheckAttributeBlockWithDefaults

`func NewScalarCheckAttributeBlockWithDefaults() *ScalarCheckAttributeBlock`

NewScalarCheckAttributeBlockWithDefaults instantiates a new ScalarCheckAttributeBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *ScalarCheckAttributeBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ScalarCheckAttributeBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ScalarCheckAttributeBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *ScalarCheckAttributeBlock) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValue

`func (o *ScalarCheckAttributeBlock) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ScalarCheckAttributeBlock) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ScalarCheckAttributeBlock) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


