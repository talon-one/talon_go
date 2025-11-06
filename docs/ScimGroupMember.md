# ScimGroupMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | Unique identifier of the member. | [optional] 
**Display** | Pointer to **string** | Identifier of the user. This is usually an email address. | [optional] 

## Methods

### NewScimGroupMember

`func NewScimGroupMember() *ScimGroupMember`

NewScimGroupMember instantiates a new ScimGroupMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimGroupMemberWithDefaults

`func NewScimGroupMemberWithDefaults() *ScimGroupMember`

NewScimGroupMemberWithDefaults instantiates a new ScimGroupMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ScimGroupMember) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ScimGroupMember) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ScimGroupMember) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ScimGroupMember) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDisplay

`func (o *ScimGroupMember) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ScimGroupMember) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ScimGroupMember) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *ScimGroupMember) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


