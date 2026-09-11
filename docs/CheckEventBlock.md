# CheckEventBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**EventType** | Pointer to **string** | The event type to check against. | 
**Matchers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**OnFailure** | Pointer to **[]map[string]interface{}** | Promotion blocks evaluated when this block fails or returns false. | [optional] 

## Methods

### NewCheckEventBlock

`func NewCheckEventBlock(type_ string, eventType string, ) *CheckEventBlock`

NewCheckEventBlock instantiates a new CheckEventBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckEventBlockWithDefaults

`func NewCheckEventBlockWithDefaults() *CheckEventBlock`

NewCheckEventBlockWithDefaults instantiates a new CheckEventBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckEventBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckEventBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckEventBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CheckEventBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CheckEventBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckEventBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckEventBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *CheckEventBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CheckEventBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CheckEventBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CheckEventBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEventType

`func (o *CheckEventBlock) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CheckEventBlock) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CheckEventBlock) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetMatchers

`func (o *CheckEventBlock) GetMatchers() []map[string]interface{}`

GetMatchers returns the Matchers field if non-nil, zero value otherwise.

### GetMatchersOk

`func (o *CheckEventBlock) GetMatchersOk() (*[]map[string]interface{}, bool)`

GetMatchersOk returns a tuple with the Matchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchers

`func (o *CheckEventBlock) SetMatchers(v []map[string]interface{})`

SetMatchers sets Matchers field to given value.

### HasMatchers

`func (o *CheckEventBlock) HasMatchers() bool`

HasMatchers returns a boolean if a field has been set.

### GetOnFailure

`func (o *CheckEventBlock) GetOnFailure() []map[string]interface{}`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *CheckEventBlock) GetOnFailureOk() (*[]map[string]interface{}, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *CheckEventBlock) SetOnFailure(v []map[string]interface{})`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *CheckEventBlock) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


