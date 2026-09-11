# SortSelectorStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A step discriminator of type &#x60;sort&#x60;. | 
**Fields** | Pointer to [**[]SortSelectorStepField**](SortSelectorStepField.md) | One or more fields to sort by, applied in order. Each field has its own direction. | 

## Methods

### NewSortSelectorStep

`func NewSortSelectorStep(type_ string, fields []SortSelectorStepField, ) *SortSelectorStep`

NewSortSelectorStep instantiates a new SortSelectorStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSortSelectorStepWithDefaults

`func NewSortSelectorStepWithDefaults() *SortSelectorStep`

NewSortSelectorStepWithDefaults instantiates a new SortSelectorStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SortSelectorStep) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SortSelectorStep) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SortSelectorStep) SetType(v string)`

SetType sets Type field to given value.


### GetFields

`func (o *SortSelectorStep) GetFields() []SortSelectorStepField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SortSelectorStep) GetFieldsOk() (*[]SortSelectorStepField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SortSelectorStep) SetFields(v []SortSelectorStepField)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


