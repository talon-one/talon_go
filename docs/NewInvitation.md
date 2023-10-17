# NewInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the user being invited. | [optional] 
**Email** | Pointer to **string** |  | 
**Acl** | Pointer to **string** | The &#x60;Access Control List&#x60; json defining the role of the user.  This represents the access control on the user level. Use one of the following: - normal user: &#x60;{\&quot;Role\&quot;: 0}&#x60; - admin: &#x60;{\&quot;Role\&quot;: 127}&#x60;  | [optional] 
**IsAdmin** | Pointer to **bool** | An indication of whether the user has admin permissions. We recommend using this flag over using the &#x60;acl&#x60; with value &#x60;{\&quot;Role\&quot;: 127}&#x60;.  | [optional] 
**Roles** | Pointer to **[]int32** | An array of role IDs to assign to the new user. | [optional] 

## Methods

### GetName

`func (o *NewInvitation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewInvitation) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *NewInvitation) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *NewInvitation) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetEmail

`func (o *NewInvitation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NewInvitation) GetEmailOk() (string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmail

`func (o *NewInvitation) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmail

`func (o *NewInvitation) SetEmail(v string)`

SetEmail gets a reference to the given string and assigns it to the Email field.

### GetAcl

`func (o *NewInvitation) GetAcl() string`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *NewInvitation) GetAclOk() (string, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAcl

`func (o *NewInvitation) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### SetAcl

`func (o *NewInvitation) SetAcl(v string)`

SetAcl gets a reference to the given string and assigns it to the Acl field.

### GetIsAdmin

`func (o *NewInvitation) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *NewInvitation) GetIsAdminOk() (bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsAdmin

`func (o *NewInvitation) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### SetIsAdmin

`func (o *NewInvitation) SetIsAdmin(v bool)`

SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.

### GetRoles

`func (o *NewInvitation) GetRoles() []int32`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *NewInvitation) GetRolesOk() ([]int32, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoles

`func (o *NewInvitation) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRoles

`func (o *NewInvitation) SetRoles(v []int32)`

SetRoles gets a reference to the given []int32 and assigns it to the Roles field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


