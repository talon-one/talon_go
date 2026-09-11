# RoleV2Permissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionSets** | Pointer to [**[]RoleV2PermissionSet**](RoleV2PermissionSet.md) | List of grouped logical operations referenced by roles. | [optional] 
**Roles** | Pointer to [**RoleV2RolesGroup**](RoleV2RolesGroup.md) |  | [optional] 
**Thresholds** | Pointer to [**[]RolesV2Thresholds**](RolesV2Thresholds.md) | Support user limits for actions that require admin approval within the given application. | [optional] 

## Methods

### NewRoleV2Permissions

`func NewRoleV2Permissions() *RoleV2Permissions`

NewRoleV2Permissions instantiates a new RoleV2Permissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleV2PermissionsWithDefaults

`func NewRoleV2PermissionsWithDefaults() *RoleV2Permissions`

NewRoleV2PermissionsWithDefaults instantiates a new RoleV2Permissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionSets

`func (o *RoleV2Permissions) GetPermissionSets() []RoleV2PermissionSet`

GetPermissionSets returns the PermissionSets field if non-nil, zero value otherwise.

### GetPermissionSetsOk

`func (o *RoleV2Permissions) GetPermissionSetsOk() (*[]RoleV2PermissionSet, bool)`

GetPermissionSetsOk returns a tuple with the PermissionSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSets

`func (o *RoleV2Permissions) SetPermissionSets(v []RoleV2PermissionSet)`

SetPermissionSets sets PermissionSets field to given value.

### HasPermissionSets

`func (o *RoleV2Permissions) HasPermissionSets() bool`

HasPermissionSets returns a boolean if a field has been set.

### GetRoles

`func (o *RoleV2Permissions) GetRoles() RoleV2RolesGroup`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *RoleV2Permissions) GetRolesOk() (*RoleV2RolesGroup, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *RoleV2Permissions) SetRoles(v RoleV2RolesGroup)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *RoleV2Permissions) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetThresholds

`func (o *RoleV2Permissions) GetThresholds() []RolesV2Thresholds`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *RoleV2Permissions) GetThresholdsOk() (*[]RolesV2Thresholds, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *RoleV2Permissions) SetThresholds(v []RolesV2Thresholds)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *RoleV2Permissions) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


