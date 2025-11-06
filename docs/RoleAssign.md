# RoleAssign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to **[]int64** | An array of user IDs. | 
**Roles** | Pointer to **[]int64** | An array of role IDs. | 

## Methods

### NewRoleAssign

`func NewRoleAssign(users []int64, roles []int64, ) *RoleAssign`

NewRoleAssign instantiates a new RoleAssign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleAssignWithDefaults

`func NewRoleAssignWithDefaults() *RoleAssign`

NewRoleAssignWithDefaults instantiates a new RoleAssign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *RoleAssign) GetUsers() []int64`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *RoleAssign) GetUsersOk() (*[]int64, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *RoleAssign) SetUsers(v []int64)`

SetUsers sets Users field to given value.


### GetRoles

`func (o *RoleAssign) GetRoles() []int64`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *RoleAssign) GetRolesOk() (*[]int64, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *RoleAssign) SetRoles(v []int64)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


