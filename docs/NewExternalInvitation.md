# NewExternalInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email address of the user. | 
**Name** | Pointer to **string** | Name of the user. | [optional] 
**UserGroups** | Pointer to **[]string** | List of user groups in the external identity provider.  If there are roles in Talon.One whose names match these user groups, those roles will be automatically assigned to the user upon invitation.  | [optional] 

## Methods

### GetEmail

`func (o *NewExternalInvitation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NewExternalInvitation) GetEmailOk() (string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmail

`func (o *NewExternalInvitation) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmail

`func (o *NewExternalInvitation) SetEmail(v string)`

SetEmail gets a reference to the given string and assigns it to the Email field.

### GetName

`func (o *NewExternalInvitation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewExternalInvitation) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *NewExternalInvitation) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *NewExternalInvitation) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetUserGroups

`func (o *NewExternalInvitation) GetUserGroups() []string`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *NewExternalInvitation) GetUserGroupsOk() ([]string, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserGroups

`func (o *NewExternalInvitation) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.

### SetUserGroups

`func (o *NewExternalInvitation) SetUserGroups(v []string)`

SetUserGroups gets a reference to the given []string and assigns it to the UserGroups field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


