# RoleV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the role. | [optional] 
**Description** | Pointer to **string** | Description of the role. | [optional] 
**IsAdmin** | Pointer to **bool** | Indicates whether the role grants admin permissions. | [optional] 
**Permissions** | Pointer to [**RoleV2Permissions**](RoleV2Permissions.md) |  | [optional] 
**Members** | Pointer to **[]int32** | An array of user identifiers. | [optional] 

## Methods

### GetName

`func (o *RoleV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleV2) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *RoleV2) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *RoleV2) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *RoleV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleV2) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *RoleV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *RoleV2) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetIsAdmin

`func (o *RoleV2) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *RoleV2) GetIsAdminOk() (bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsAdmin

`func (o *RoleV2) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### SetIsAdmin

`func (o *RoleV2) SetIsAdmin(v bool)`

SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.

### GetPermissions

`func (o *RoleV2) GetPermissions() RoleV2Permissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RoleV2) GetPermissionsOk() (RoleV2Permissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPermissions

`func (o *RoleV2) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissions

`func (o *RoleV2) SetPermissions(v RoleV2Permissions)`

SetPermissions gets a reference to the given RoleV2Permissions and assigns it to the Permissions field.

### GetMembers

`func (o *RoleV2) GetMembers() []int32`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *RoleV2) GetMembersOk() ([]int32, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMembers

`func (o *RoleV2) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembers

`func (o *RoleV2) SetMembers(v []int32)`

SetMembers gets a reference to the given []int32 and assigns it to the Members field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


