# ScimGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Display name of the group (Talon.One role). | [optional] 
**Members** | Pointer to [**[]ScimGroupMember**](ScimGroupMember.md) | List of members to assign to the new Talon.One role. | [optional] 
**Id** | Pointer to **string** | ID of the group. | 

## Methods

### GetDisplayName

`func (o *ScimGroup) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ScimGroup) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *ScimGroup) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *ScimGroup) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetMembers

`func (o *ScimGroup) GetMembers() []ScimGroupMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ScimGroup) GetMembersOk() ([]ScimGroupMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMembers

`func (o *ScimGroup) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembers

`func (o *ScimGroup) SetMembers(v []ScimGroupMember)`

SetMembers gets a reference to the given []ScimGroupMember and assigns it to the Members field.

### GetId

`func (o *ScimGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScimGroup) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *ScimGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *ScimGroup) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


