# CheckAttributeBlockBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**Operator** | Pointer to **string** | The comparison operator applied to the attribute. | 
**Attribute** | Pointer to [**map[string]interface{}**](.md) | The attribute path identifier (e.g. \&quot;$Session.Total\&quot;). | 
**Value** | Pointer to [**map[string]interface{}**](.md) | The comparison value for scalar operators. | [optional] 
**Min** | Pointer to [**map[string]interface{}**](.md) | The minimum value allowed for the &#x60;between&#x60; operator. | [optional] 
**Max** | Pointer to [**map[string]interface{}**](.md) | The maximum value allowed for the &#x60;between&#x60; operator. | [optional] 
**Start** | Pointer to [**map[string]interface{}**](.md) | The start value for the &#x60;within&#x60; operator. | [optional] 
**End** | Pointer to [**map[string]interface{}**](.md) | The end value for the &#x60;within&#x60; operator. | [optional] 
**StartInclusive** | Pointer to **bool** | When &#x60;true&#x60;, the &#x60;start&#x60; value is included in the range for the &#x60;within&#x60; operator. | [optional] 
**EndInclusive** | Pointer to **bool** | When &#x60;true&#x60;, the &#x60;end&#x60; value is included in the range for the &#x60;within&#x60; operator. | [optional] 
**TimezoneInsensitive** | Pointer to **bool** | Indicates whether the &#x60;within&#x60; operator ignores time zones and compares the wall-clock time only. When &#x60;false&#x60;, time zones are taken into account. | [optional] 
**Values** | Pointer to [**map[string]interface{}**](.md) | The set of values to match against for list operators. For location operators (&#x60;in&#x60;, &#x60;not(in)&#x60;), an array of objects with a &#x60;geometry&#x60; (see &#x60;GeoJSONGeometry&#x60;) and an optional &#x60;name&#x60;, or a string reference to a list attribute. | [optional] 
**Count** | Pointer to [**map[string]interface{}**](.md) | The count threshold for &#x60;containsAtLeast&#x60; and &#x60;containsExactly&#x60; operators. | [optional] 
**OnFailure** | Pointer to **[]map[string]interface{}** | Promotion blocks evaluated when this block fails or returns false. | [optional] 

## Methods

### NewCheckAttributeBlockBase

`func NewCheckAttributeBlockBase(type_ string, operator string, attribute map[string]interface{}, ) *CheckAttributeBlockBase`

NewCheckAttributeBlockBase instantiates a new CheckAttributeBlockBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckAttributeBlockBaseWithDefaults

`func NewCheckAttributeBlockBaseWithDefaults() *CheckAttributeBlockBase`

NewCheckAttributeBlockBaseWithDefaults instantiates a new CheckAttributeBlockBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckAttributeBlockBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckAttributeBlockBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckAttributeBlockBase) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CheckAttributeBlockBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CheckAttributeBlockBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckAttributeBlockBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckAttributeBlockBase) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *CheckAttributeBlockBase) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CheckAttributeBlockBase) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CheckAttributeBlockBase) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CheckAttributeBlockBase) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *CheckAttributeBlockBase) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CheckAttributeBlockBase) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CheckAttributeBlockBase) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetAttribute

`func (o *CheckAttributeBlockBase) GetAttribute() map[string]interface{}`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *CheckAttributeBlockBase) GetAttributeOk() (*map[string]interface{}, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *CheckAttributeBlockBase) SetAttribute(v map[string]interface{})`

SetAttribute sets Attribute field to given value.


### GetValue

`func (o *CheckAttributeBlockBase) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CheckAttributeBlockBase) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CheckAttributeBlockBase) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *CheckAttributeBlockBase) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMin

`func (o *CheckAttributeBlockBase) GetMin() map[string]interface{}`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *CheckAttributeBlockBase) GetMinOk() (*map[string]interface{}, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *CheckAttributeBlockBase) SetMin(v map[string]interface{})`

SetMin sets Min field to given value.

### HasMin

`func (o *CheckAttributeBlockBase) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *CheckAttributeBlockBase) GetMax() map[string]interface{}`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *CheckAttributeBlockBase) GetMaxOk() (*map[string]interface{}, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *CheckAttributeBlockBase) SetMax(v map[string]interface{})`

SetMax sets Max field to given value.

### HasMax

`func (o *CheckAttributeBlockBase) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetStart

`func (o *CheckAttributeBlockBase) GetStart() map[string]interface{}`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CheckAttributeBlockBase) GetStartOk() (*map[string]interface{}, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CheckAttributeBlockBase) SetStart(v map[string]interface{})`

SetStart sets Start field to given value.

### HasStart

`func (o *CheckAttributeBlockBase) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *CheckAttributeBlockBase) GetEnd() map[string]interface{}`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CheckAttributeBlockBase) GetEndOk() (*map[string]interface{}, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CheckAttributeBlockBase) SetEnd(v map[string]interface{})`

SetEnd sets End field to given value.

### HasEnd

`func (o *CheckAttributeBlockBase) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStartInclusive

`func (o *CheckAttributeBlockBase) GetStartInclusive() bool`

GetStartInclusive returns the StartInclusive field if non-nil, zero value otherwise.

### GetStartInclusiveOk

`func (o *CheckAttributeBlockBase) GetStartInclusiveOk() (*bool, bool)`

GetStartInclusiveOk returns a tuple with the StartInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartInclusive

`func (o *CheckAttributeBlockBase) SetStartInclusive(v bool)`

SetStartInclusive sets StartInclusive field to given value.

### HasStartInclusive

`func (o *CheckAttributeBlockBase) HasStartInclusive() bool`

HasStartInclusive returns a boolean if a field has been set.

### GetEndInclusive

`func (o *CheckAttributeBlockBase) GetEndInclusive() bool`

GetEndInclusive returns the EndInclusive field if non-nil, zero value otherwise.

### GetEndInclusiveOk

`func (o *CheckAttributeBlockBase) GetEndInclusiveOk() (*bool, bool)`

GetEndInclusiveOk returns a tuple with the EndInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndInclusive

`func (o *CheckAttributeBlockBase) SetEndInclusive(v bool)`

SetEndInclusive sets EndInclusive field to given value.

### HasEndInclusive

`func (o *CheckAttributeBlockBase) HasEndInclusive() bool`

HasEndInclusive returns a boolean if a field has been set.

### GetTimezoneInsensitive

`func (o *CheckAttributeBlockBase) GetTimezoneInsensitive() bool`

GetTimezoneInsensitive returns the TimezoneInsensitive field if non-nil, zero value otherwise.

### GetTimezoneInsensitiveOk

`func (o *CheckAttributeBlockBase) GetTimezoneInsensitiveOk() (*bool, bool)`

GetTimezoneInsensitiveOk returns a tuple with the TimezoneInsensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneInsensitive

`func (o *CheckAttributeBlockBase) SetTimezoneInsensitive(v bool)`

SetTimezoneInsensitive sets TimezoneInsensitive field to given value.

### HasTimezoneInsensitive

`func (o *CheckAttributeBlockBase) HasTimezoneInsensitive() bool`

HasTimezoneInsensitive returns a boolean if a field has been set.

### GetValues

`func (o *CheckAttributeBlockBase) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CheckAttributeBlockBase) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CheckAttributeBlockBase) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *CheckAttributeBlockBase) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetCount

`func (o *CheckAttributeBlockBase) GetCount() map[string]interface{}`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CheckAttributeBlockBase) GetCountOk() (*map[string]interface{}, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CheckAttributeBlockBase) SetCount(v map[string]interface{})`

SetCount sets Count field to given value.

### HasCount

`func (o *CheckAttributeBlockBase) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetOnFailure

`func (o *CheckAttributeBlockBase) GetOnFailure() []map[string]interface{}`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *CheckAttributeBlockBase) GetOnFailureOk() (*[]map[string]interface{}, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *CheckAttributeBlockBase) SetOnFailure(v []map[string]interface{})`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *CheckAttributeBlockBase) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


