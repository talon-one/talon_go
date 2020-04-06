# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the role corresponding to the DB row | 
**AccountID** | Pointer to **int32** | The ID of the Talon.One account that owns this role. | 
**Name** | Pointer to **string** | Name of the role | [optional] 
**Description** | Pointer to **string** | Description of the role | [optional] 
**Members** | Pointer to **[]int32** | A list of user identifiers assigned to this role | [optional] 
**Acl** | Pointer to **string** | Role Policy this should be a stringified blob of json | [optional] 

## Methods

### GetId

`func (o *Role) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Role) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Role) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Role) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetAccountID

`func (o *Role) GetAccountID() int32`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *Role) GetAccountIDOk() (int32, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountID

`func (o *Role) HasAccountID() bool`

HasAccountID returns a boolean if a field has been set.

### SetAccountID

`func (o *Role) SetAccountID(v int32)`

SetAccountID gets a reference to the given int32 and assigns it to the AccountID field.

### GetName

`func (o *Role) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Role) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Role) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Role) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Role) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Role) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Role) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Role) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetMembers

`func (o *Role) GetMembers() []int32`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Role) GetMembersOk() ([]int32, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMembers

`func (o *Role) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembers

`func (o *Role) SetMembers(v []int32)`

SetMembers gets a reference to the given []int32 and assigns it to the Members field.

### GetAcl

`func (o *Role) GetAcl() string`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *Role) GetAclOk() (string, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAcl

`func (o *Role) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### SetAcl

`func (o *Role) SetAcl(v string)`

SetAcl gets a reference to the given string and assigns it to the Acl field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


