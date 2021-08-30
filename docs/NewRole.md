# NewRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the role | 
**Description** | Pointer to **string** | Description of the role | [optional] 
**Acl** | Pointer to **string** | Role Policy this should be a stringified blob of json | 
**Members** | Pointer to **[]int32** | An array of user identifiers | 

## Methods

### GetName

`func (o *NewRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewRole) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *NewRole) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *NewRole) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *NewRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewRole) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *NewRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *NewRole) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetAcl

`func (o *NewRole) GetAcl() string`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *NewRole) GetAclOk() (string, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAcl

`func (o *NewRole) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### SetAcl

`func (o *NewRole) SetAcl(v string)`

SetAcl gets a reference to the given string and assigns it to the Acl field.

### GetMembers

`func (o *NewRole) GetMembers() []int32`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *NewRole) GetMembersOk() ([]int32, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMembers

`func (o *NewRole) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembers

`func (o *NewRole) SetMembers(v []int32)`

SetMembers gets a reference to the given []int32 and assigns it to the Members field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


