# ListWithCountCheckAttributeBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** | The list membership operator with a count threshold applied to the attribute. | [optional] 
**Values** | Pointer to [**map[string]interface{}**](.md) | The set of values to match against. | 
**Count** | Pointer to [**map[string]interface{}**](.md) | The count threshold for this operator. | 

## Methods

### NewListWithCountCheckAttributeBlock

`func NewListWithCountCheckAttributeBlock(values map[string]interface{}, count map[string]interface{}, ) *ListWithCountCheckAttributeBlock`

NewListWithCountCheckAttributeBlock instantiates a new ListWithCountCheckAttributeBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWithCountCheckAttributeBlockWithDefaults

`func NewListWithCountCheckAttributeBlockWithDefaults() *ListWithCountCheckAttributeBlock`

NewListWithCountCheckAttributeBlockWithDefaults instantiates a new ListWithCountCheckAttributeBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *ListWithCountCheckAttributeBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ListWithCountCheckAttributeBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ListWithCountCheckAttributeBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *ListWithCountCheckAttributeBlock) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *ListWithCountCheckAttributeBlock) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ListWithCountCheckAttributeBlock) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ListWithCountCheckAttributeBlock) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.


### GetCount

`func (o *ListWithCountCheckAttributeBlock) GetCount() map[string]interface{}`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListWithCountCheckAttributeBlock) GetCountOk() (*map[string]interface{}, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListWithCountCheckAttributeBlock) SetCount(v map[string]interface{})`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


