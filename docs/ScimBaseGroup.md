# ScimBaseGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Display name of the group (Talon.One role). | [optional] 
**Members** | Pointer to [**[]ScimGroupMember**](ScimGroupMember.md) | List of members to assign to the new Talon.One role. | [optional] 

## Methods

### GetDisplayName

`func (o *ScimBaseGroup) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ScimBaseGroup) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *ScimBaseGroup) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *ScimBaseGroup) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetMembers

`func (o *ScimBaseGroup) GetMembers() []ScimGroupMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ScimBaseGroup) GetMembersOk() ([]ScimGroupMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMembers

`func (o *ScimBaseGroup) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembers

`func (o *ScimBaseGroup) SetMembers(v []ScimGroupMember)`

SetMembers gets a reference to the given []ScimGroupMember and assigns it to the Members field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


