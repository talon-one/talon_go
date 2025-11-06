# RoleV2Base

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the role. | [optional] 
**Description** | Pointer to **string** | Description of the role. | [optional] 
**Permissions** | Pointer to [**RoleV2Permissions**](RoleV2Permissions.md) |  | [optional] 
**Members** | Pointer to **[]int64** | A list of user IDs the role is assigned to. | [optional] 

## Methods

### NewRoleV2Base

`func NewRoleV2Base() *RoleV2Base`

NewRoleV2Base instantiates a new RoleV2Base object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleV2BaseWithDefaults

`func NewRoleV2BaseWithDefaults() *RoleV2Base`

NewRoleV2BaseWithDefaults instantiates a new RoleV2Base object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RoleV2Base) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleV2Base) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleV2Base) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoleV2Base) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RoleV2Base) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleV2Base) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleV2Base) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleV2Base) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPermissions

`func (o *RoleV2Base) GetPermissions() RoleV2Permissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RoleV2Base) GetPermissionsOk() (*RoleV2Permissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RoleV2Base) SetPermissions(v RoleV2Permissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *RoleV2Base) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetMembers

`func (o *RoleV2Base) GetMembers() []int64`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *RoleV2Base) GetMembersOk() (*[]int64, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *RoleV2Base) SetMembers(v []int64)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *RoleV2Base) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


