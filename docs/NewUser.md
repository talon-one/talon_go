# NewUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The email address associated with the user profile. | 
**Password** | Pointer to **string** | The password for your account. | 
**Name** | Pointer to **string** | Your name. | [optional] 
**InviteToken** | Pointer to **string** |  | 

## Methods

### NewNewUser

`func NewNewUser(email string, password string, inviteToken string, ) *NewUser`

NewNewUser instantiates a new NewUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewUserWithDefaults

`func NewNewUserWithDefaults() *NewUser`

NewNewUserWithDefaults instantiates a new NewUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *NewUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NewUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NewUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *NewUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NewUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NewUser) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetName

`func (o *NewUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NewUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInviteToken

`func (o *NewUser) GetInviteToken() string`

GetInviteToken returns the InviteToken field if non-nil, zero value otherwise.

### GetInviteTokenOk

`func (o *NewUser) GetInviteTokenOk() (*string, bool)`

GetInviteTokenOk returns a tuple with the InviteToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteToken

`func (o *NewUser) SetInviteToken(v string)`

SetInviteToken sets InviteToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


