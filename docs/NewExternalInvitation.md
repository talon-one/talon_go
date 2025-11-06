# NewExternalInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the user. | [optional] 
**UserGroups** | Pointer to **[]string** | List of user groups in the external identity provider.  If there are roles in Talon.One whose names match these user groups, those roles will be automatically assigned to the user upon invitation.  | [optional] 
**Email** | Pointer to **string** | Email address of the user. | 

## Methods

### NewNewExternalInvitation

`func NewNewExternalInvitation(email string, ) *NewExternalInvitation`

NewNewExternalInvitation instantiates a new NewExternalInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewExternalInvitationWithDefaults

`func NewNewExternalInvitationWithDefaults() *NewExternalInvitation`

NewNewExternalInvitationWithDefaults instantiates a new NewExternalInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewExternalInvitation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewExternalInvitation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewExternalInvitation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NewExternalInvitation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserGroups

`func (o *NewExternalInvitation) GetUserGroups() []string`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *NewExternalInvitation) GetUserGroupsOk() (*[]string, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *NewExternalInvitation) SetUserGroups(v []string)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *NewExternalInvitation) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.

### GetEmail

`func (o *NewExternalInvitation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NewExternalInvitation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NewExternalInvitation) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


