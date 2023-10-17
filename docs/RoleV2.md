# RoleV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 
**AccountId** | Pointer to **int32** | The ID of the account that owns this entity. | 
**Name** | Pointer to **string** | Name of the role. | [optional] 
**Description** | Pointer to **string** | Description of the role. | [optional] 
**Permissions** | Pointer to [**RoleV2Permissions**](RoleV2Permissions.md) |  | [optional] 
**Members** | Pointer to **[]int32** | A list of user identifiers the role is assigned to. | [optional] 

## Methods

### GetId

`func (o *RoleV2) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleV2) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *RoleV2) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *RoleV2) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *RoleV2) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RoleV2) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *RoleV2) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *RoleV2) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetModified

`func (o *RoleV2) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *RoleV2) GetModifiedOk() (time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *RoleV2) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *RoleV2) SetModified(v time.Time)`

SetModified gets a reference to the given time.Time and assigns it to the Modified field.

### GetAccountId

`func (o *RoleV2) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RoleV2) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *RoleV2) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *RoleV2) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

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


