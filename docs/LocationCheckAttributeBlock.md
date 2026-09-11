# LocationCheckAttributeBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** | The location membership operator applied to the attribute. | [optional] 
**Values** | Pointer to [**map[string]interface{}**](.md) | The geometric areas to check the location against. | 

## Methods

### NewLocationCheckAttributeBlock

`func NewLocationCheckAttributeBlock(values map[string]interface{}, ) *LocationCheckAttributeBlock`

NewLocationCheckAttributeBlock instantiates a new LocationCheckAttributeBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationCheckAttributeBlockWithDefaults

`func NewLocationCheckAttributeBlockWithDefaults() *LocationCheckAttributeBlock`

NewLocationCheckAttributeBlockWithDefaults instantiates a new LocationCheckAttributeBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *LocationCheckAttributeBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *LocationCheckAttributeBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *LocationCheckAttributeBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *LocationCheckAttributeBlock) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *LocationCheckAttributeBlock) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *LocationCheckAttributeBlock) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *LocationCheckAttributeBlock) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


