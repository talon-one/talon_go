# RoleV2Permissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionSets** | Pointer to [**[]RoleV2PermissionSet**](RoleV2PermissionSet.md) | List of grouped logical operations to use as a reference in the roles section. Each group of logical operations has a name. | [optional] 
**Roles** | Pointer to [**RoleV2RolesGroup**](RoleV2RolesGroup.md) |  | [optional] 

## Methods

### GetPermissionSets

`func (o *RoleV2Permissions) GetPermissionSets() []RoleV2PermissionSet`

GetPermissionSets returns the PermissionSets field if non-nil, zero value otherwise.

### GetPermissionSetsOk

`func (o *RoleV2Permissions) GetPermissionSetsOk() ([]RoleV2PermissionSet, bool)`

GetPermissionSetsOk returns a tuple with the PermissionSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPermissionSets

`func (o *RoleV2Permissions) HasPermissionSets() bool`

HasPermissionSets returns a boolean if a field has been set.

### SetPermissionSets

`func (o *RoleV2Permissions) SetPermissionSets(v []RoleV2PermissionSet)`

SetPermissionSets gets a reference to the given []RoleV2PermissionSet and assigns it to the PermissionSets field.

### GetRoles

`func (o *RoleV2Permissions) GetRoles() RoleV2RolesGroup`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *RoleV2Permissions) GetRolesOk() (RoleV2RolesGroup, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoles

`func (o *RoleV2Permissions) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRoles

`func (o *RoleV2Permissions) SetRoles(v RoleV2RolesGroup)`

SetRoles gets a reference to the given RoleV2RolesGroup and assigns it to the Roles field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


