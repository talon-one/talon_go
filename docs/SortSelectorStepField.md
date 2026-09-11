# SortSelectorStepField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | Pointer to **string** | The attribute path the items are sorted by. | 
**Direction** | Pointer to **string** | The sort direction for this field. | 

## Methods

### NewSortSelectorStepField

`func NewSortSelectorStepField(expression string, direction string, ) *SortSelectorStepField`

NewSortSelectorStepField instantiates a new SortSelectorStepField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSortSelectorStepFieldWithDefaults

`func NewSortSelectorStepFieldWithDefaults() *SortSelectorStepField`

NewSortSelectorStepFieldWithDefaults instantiates a new SortSelectorStepField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *SortSelectorStepField) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *SortSelectorStepField) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *SortSelectorStepField) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetDirection

`func (o *SortSelectorStepField) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SortSelectorStepField) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SortSelectorStepField) SetDirection(v string)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


