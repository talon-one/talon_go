# NewRoleV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the role. | 
**Description** | Pointer to **string** | Description of the role. | 
**Permissions** | Pointer to [**RoleV2Permissions**](RoleV2Permissions.md) |  | [optional] 
**Members** | Pointer to **[]int32** | A list of user identifiers the role is assigned to. | [optional] 

## Methods

### GetName

`func (o *NewRoleV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewRoleV2) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *NewRoleV2) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *NewRoleV2) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *NewRoleV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewRoleV2) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *NewRoleV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *NewRoleV2) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetPermissions

`func (o *NewRoleV2) GetPermissions() RoleV2Permissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *NewRoleV2) GetPermissionsOk() (RoleV2Permissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPermissions

`func (o *NewRoleV2) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissions

`func (o *NewRoleV2) SetPermissions(v RoleV2Permissions)`

SetPermissions gets a reference to the given RoleV2Permissions and assigns it to the Permissions field.

### GetMembers

`func (o *NewRoleV2) GetMembers() []int32`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *NewRoleV2) GetMembersOk() ([]int32, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMembers

`func (o *NewRoleV2) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembers

`func (o *NewRoleV2) SetMembers(v []int32)`

SetMembers gets a reference to the given []int32 and assigns it to the Members field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


