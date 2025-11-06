# RoleV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 
**AccountId** | Pointer to **int64** | The ID of the account that owns this entity. | 
**Name** | Pointer to **string** | Name of the role. | [optional] 
**Description** | Pointer to **string** | Description of the role. | [optional] 
**Permissions** | Pointer to [**RoleV2Permissions**](RoleV2Permissions.md) |  | [optional] 
**Members** | Pointer to **[]int64** | A list of user IDs the role is assigned to. | [optional] 

## Methods

### NewRoleV2

`func NewRoleV2(id int64, created time.Time, modified time.Time, accountId int64, ) *RoleV2`

NewRoleV2 instantiates a new RoleV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleV2WithDefaults

`func NewRoleV2WithDefaults() *RoleV2`

NewRoleV2WithDefaults instantiates a new RoleV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoleV2) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleV2) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleV2) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *RoleV2) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RoleV2) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RoleV2) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *RoleV2) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *RoleV2) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *RoleV2) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetAccountId

`func (o *RoleV2) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RoleV2) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RoleV2) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *RoleV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoleV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RoleV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPermissions

`func (o *RoleV2) GetPermissions() RoleV2Permissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RoleV2) GetPermissionsOk() (*RoleV2Permissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RoleV2) SetPermissions(v RoleV2Permissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *RoleV2) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetMembers

`func (o *RoleV2) GetMembers() []int64`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *RoleV2) GetMembersOk() (*[]int64, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *RoleV2) SetMembers(v []int64)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *RoleV2) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


