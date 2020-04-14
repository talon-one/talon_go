# UpdateRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the role | [optional] 
**Description** | Pointer to **string** | Description of the role | [optional] 
**Acl** | Pointer to **string** | Role Policy this should be a stringified blob of json | [optional] 
**Users** | Pointer to **[]int32** | an array of user identifiers | [optional] 

## Methods

### GetName

`func (o *UpdateRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRole) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *UpdateRole) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *UpdateRole) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *UpdateRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateRole) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *UpdateRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *UpdateRole) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetAcl

`func (o *UpdateRole) GetAcl() string`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *UpdateRole) GetAclOk() (string, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAcl

`func (o *UpdateRole) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### SetAcl

`func (o *UpdateRole) SetAcl(v string)`

SetAcl gets a reference to the given string and assigns it to the Acl field.

### GetUsers

`func (o *UpdateRole) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UpdateRole) GetUsersOk() ([]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsers

`func (o *UpdateRole) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsers

`func (o *UpdateRole) SetUsers(v []int32)`

SetUsers gets a reference to the given []int32 and assigns it to the Users field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


