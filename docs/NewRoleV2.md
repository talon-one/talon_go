# NewRoleV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the role. | 
**Description** | Pointer to **string** | Description of the role. | 
**Permissions** | Pointer to [**RoleV2Permissions**](RoleV2Permissions.md) |  | [optional] 
**Members** | Pointer to **[]int64** | A list of user IDs the role is assigned to. | [optional] 

## Methods

### NewNewRoleV2

`func NewNewRoleV2(name string, description string, ) *NewRoleV2`

NewNewRoleV2 instantiates a new NewRoleV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewRoleV2WithDefaults

`func NewNewRoleV2WithDefaults() *NewRoleV2`

NewNewRoleV2WithDefaults instantiates a new NewRoleV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewRoleV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewRoleV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewRoleV2) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NewRoleV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewRoleV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewRoleV2) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPermissions

`func (o *NewRoleV2) GetPermissions() RoleV2Permissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *NewRoleV2) GetPermissionsOk() (*RoleV2Permissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *NewRoleV2) SetPermissions(v RoleV2Permissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *NewRoleV2) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetMembers

`func (o *NewRoleV2) GetMembers() []int64`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *NewRoleV2) GetMembersOk() (*[]int64, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *NewRoleV2) SetMembers(v []int64)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *NewRoleV2) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


