# MapSelectorStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A step discriminator of type &#x60;map&#x60;. | 
**Expression** | Pointer to **string** | The attribute path each item is mapped to. | 

## Methods

### NewMapSelectorStep

`func NewMapSelectorStep(type_ string, expression string, ) *MapSelectorStep`

NewMapSelectorStep instantiates a new MapSelectorStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapSelectorStepWithDefaults

`func NewMapSelectorStepWithDefaults() *MapSelectorStep`

NewMapSelectorStepWithDefaults instantiates a new MapSelectorStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MapSelectorStep) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MapSelectorStep) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MapSelectorStep) SetType(v string)`

SetType sets Type field to given value.


### GetExpression

`func (o *MapSelectorStep) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *MapSelectorStep) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *MapSelectorStep) SetExpression(v string)`

SetExpression sets Expression field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


