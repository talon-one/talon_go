# SelectSelectorStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A step discriminator of type &#x60;select&#x60;. | 
**Operator** | Pointer to **string** | The selection operator applied to the items. | 
**From** | Pointer to [**map[string]interface{}**](.md) | The starting value of the selection. For the &#x60;many&#x60; operator this is the string &#x60;start&#x60; or &#x60;end&#x60;; for the &#x60;between&#x60; operator this is an integer start index. No discriminator is needed since the string and integer branches are distinguishable by JSON type alone. | [optional] 
**To** | Pointer to **int32** | The end index for the &#x60;between&#x60; operator. The item at this index is not included. | [optional] 
**Count** | Pointer to **int32** | The maximum number of items to select for the &#x60;many&#x60; operator. | [optional] 
**Index** | Pointer to **int32** | The exact position of the item to select for the &#x60;one&#x60; operator. | [optional] 
**Partial** | Pointer to **bool** | Indicates if the step returns fewer items than requested when the source list is shorter than the range needs. Always &#x60;true&#x60; for the &#x60;many&#x60; and &#x60;between&#x60; operators; not present for &#x60;one&#x60;, which fails instead of returning a partial result. | [optional] 

## Methods

### NewSelectSelectorStep

`func NewSelectSelectorStep(type_ string, operator string, ) *SelectSelectorStep`

NewSelectSelectorStep instantiates a new SelectSelectorStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectSelectorStepWithDefaults

`func NewSelectSelectorStepWithDefaults() *SelectSelectorStep`

NewSelectSelectorStepWithDefaults instantiates a new SelectSelectorStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SelectSelectorStep) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SelectSelectorStep) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SelectSelectorStep) SetType(v string)`

SetType sets Type field to given value.


### GetOperator

`func (o *SelectSelectorStep) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SelectSelectorStep) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SelectSelectorStep) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetFrom

`func (o *SelectSelectorStep) GetFrom() map[string]interface{}`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SelectSelectorStep) GetFromOk() (*map[string]interface{}, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SelectSelectorStep) SetFrom(v map[string]interface{})`

SetFrom sets From field to given value.

### HasFrom

`func (o *SelectSelectorStep) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *SelectSelectorStep) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SelectSelectorStep) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SelectSelectorStep) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *SelectSelectorStep) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetCount

`func (o *SelectSelectorStep) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SelectSelectorStep) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SelectSelectorStep) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SelectSelectorStep) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetIndex

`func (o *SelectSelectorStep) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *SelectSelectorStep) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *SelectSelectorStep) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *SelectSelectorStep) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetPartial

`func (o *SelectSelectorStep) GetPartial() bool`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *SelectSelectorStep) GetPartialOk() (*bool, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *SelectSelectorStep) SetPartial(v bool)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *SelectSelectorStep) HasPartial() bool`

HasPartial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


