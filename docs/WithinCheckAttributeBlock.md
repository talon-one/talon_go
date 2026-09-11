# WithinCheckAttributeBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** | The range comparison operator. Must be &#x60;within&#x60; or &#x60;not(within)&#x60;. | [optional] 
**Start** | Pointer to [**map[string]interface{}**](.md) | The start value for the &#x60;within&#x60; operator. | 
**End** | Pointer to [**map[string]interface{}**](.md) | The end value for the &#x60;within&#x60; operator. | 
**StartInclusive** | Pointer to **bool** | When &#x60;true&#x60;, the &#x60;start&#x60; value is included in the range for the &#x60;within&#x60; operator. | [optional] 
**EndInclusive** | Pointer to **bool** | When &#x60;true&#x60;, the &#x60;end&#x60; value is included in the range for the &#x60;within&#x60; operator. | [optional] 
**TimezoneInsensitive** | Pointer to **bool** | Indicates whether the &#x60;within&#x60; operator ignores time zones and compares the wall-clock time only. When &#x60;false&#x60;, time zones are taken into account. | [optional] 

## Methods

### NewWithinCheckAttributeBlock

`func NewWithinCheckAttributeBlock(start map[string]interface{}, end map[string]interface{}, ) *WithinCheckAttributeBlock`

NewWithinCheckAttributeBlock instantiates a new WithinCheckAttributeBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithinCheckAttributeBlockWithDefaults

`func NewWithinCheckAttributeBlockWithDefaults() *WithinCheckAttributeBlock`

NewWithinCheckAttributeBlockWithDefaults instantiates a new WithinCheckAttributeBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *WithinCheckAttributeBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *WithinCheckAttributeBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *WithinCheckAttributeBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *WithinCheckAttributeBlock) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetStart

`func (o *WithinCheckAttributeBlock) GetStart() map[string]interface{}`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *WithinCheckAttributeBlock) GetStartOk() (*map[string]interface{}, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *WithinCheckAttributeBlock) SetStart(v map[string]interface{})`

SetStart sets Start field to given value.


### GetEnd

`func (o *WithinCheckAttributeBlock) GetEnd() map[string]interface{}`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *WithinCheckAttributeBlock) GetEndOk() (*map[string]interface{}, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *WithinCheckAttributeBlock) SetEnd(v map[string]interface{})`

SetEnd sets End field to given value.


### GetStartInclusive

`func (o *WithinCheckAttributeBlock) GetStartInclusive() bool`

GetStartInclusive returns the StartInclusive field if non-nil, zero value otherwise.

### GetStartInclusiveOk

`func (o *WithinCheckAttributeBlock) GetStartInclusiveOk() (*bool, bool)`

GetStartInclusiveOk returns a tuple with the StartInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartInclusive

`func (o *WithinCheckAttributeBlock) SetStartInclusive(v bool)`

SetStartInclusive sets StartInclusive field to given value.

### HasStartInclusive

`func (o *WithinCheckAttributeBlock) HasStartInclusive() bool`

HasStartInclusive returns a boolean if a field has been set.

### GetEndInclusive

`func (o *WithinCheckAttributeBlock) GetEndInclusive() bool`

GetEndInclusive returns the EndInclusive field if non-nil, zero value otherwise.

### GetEndInclusiveOk

`func (o *WithinCheckAttributeBlock) GetEndInclusiveOk() (*bool, bool)`

GetEndInclusiveOk returns a tuple with the EndInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndInclusive

`func (o *WithinCheckAttributeBlock) SetEndInclusive(v bool)`

SetEndInclusive sets EndInclusive field to given value.

### HasEndInclusive

`func (o *WithinCheckAttributeBlock) HasEndInclusive() bool`

HasEndInclusive returns a boolean if a field has been set.

### GetTimezoneInsensitive

`func (o *WithinCheckAttributeBlock) GetTimezoneInsensitive() bool`

GetTimezoneInsensitive returns the TimezoneInsensitive field if non-nil, zero value otherwise.

### GetTimezoneInsensitiveOk

`func (o *WithinCheckAttributeBlock) GetTimezoneInsensitiveOk() (*bool, bool)`

GetTimezoneInsensitiveOk returns a tuple with the TimezoneInsensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneInsensitive

`func (o *WithinCheckAttributeBlock) SetTimezoneInsensitive(v bool)`

SetTimezoneInsensitive sets TimezoneInsensitive field to given value.

### HasTimezoneInsensitive

`func (o *WithinCheckAttributeBlock) HasTimezoneInsensitive() bool`

HasTimezoneInsensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


