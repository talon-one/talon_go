# ListCheckAttributeBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** | The list membership operator applied to the attribute. | [optional] 
**Values** | Pointer to [**map[string]interface{}**](.md) | The set of values to match against. | 

## Methods

### NewListCheckAttributeBlock

`func NewListCheckAttributeBlock(values map[string]interface{}, ) *ListCheckAttributeBlock`

NewListCheckAttributeBlock instantiates a new ListCheckAttributeBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCheckAttributeBlockWithDefaults

`func NewListCheckAttributeBlockWithDefaults() *ListCheckAttributeBlock`

NewListCheckAttributeBlockWithDefaults instantiates a new ListCheckAttributeBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *ListCheckAttributeBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ListCheckAttributeBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ListCheckAttributeBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *ListCheckAttributeBlock) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *ListCheckAttributeBlock) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ListCheckAttributeBlock) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ListCheckAttributeBlock) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


