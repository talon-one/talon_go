# UpdateRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the role | [optional] 
**Description** | Pointer to **string** | Description of the role | [optional] 
**Acl** | Pointer to **string** | Role Policy this should be a stringified blob of json | [optional] 
**Members** | Pointer to **[]int32** | An array of user identifiers | [optional] 

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

### GetMembers

`func (o *UpdateRole) GetMembers() []int32`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *UpdateRole) GetMembersOk() ([]int32, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMembers

`func (o *UpdateRole) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembers

`func (o *UpdateRole) SetMembers(v []int32)`

SetMembers gets a reference to the given []int32 and assigns it to the Members field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


