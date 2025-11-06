# NewInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the user. | [optional] 
**Email** | Pointer to **string** | Email address of the user. | 
**IsAdmin** | Pointer to **bool** | Indicates whether the user is an &#x60;admin&#x60;. | [optional] 
**Roles** | Pointer to **[]int64** | A list of the IDs of the roles assigned to the user. | [optional] 
**Acl** | Pointer to **string** | Indicates the access level of the user. | [optional] 

## Methods

### NewNewInvitation

`func NewNewInvitation(email string, ) *NewInvitation`

NewNewInvitation instantiates a new NewInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewInvitationWithDefaults

`func NewNewInvitationWithDefaults() *NewInvitation`

NewNewInvitationWithDefaults instantiates a new NewInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewInvitation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewInvitation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewInvitation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NewInvitation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *NewInvitation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NewInvitation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NewInvitation) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetIsAdmin

`func (o *NewInvitation) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *NewInvitation) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *NewInvitation) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *NewInvitation) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetRoles

`func (o *NewInvitation) GetRoles() []int64`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *NewInvitation) GetRolesOk() (*[]int64, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *NewInvitation) SetRoles(v []int64)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *NewInvitation) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetAcl

`func (o *NewInvitation) GetAcl() string`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *NewInvitation) GetAclOk() (*string, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *NewInvitation) SetAcl(v string)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *NewInvitation) HasAcl() bool`

HasAcl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


