# NewPassword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | The new password for your account. | 
**ResetToken** | Pointer to **string** |  | 

## Methods

### NewNewPassword

`func NewNewPassword(password string, resetToken string, ) *NewPassword`

NewNewPassword instantiates a new NewPassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewPasswordWithDefaults

`func NewNewPasswordWithDefaults() *NewPassword`

NewNewPasswordWithDefaults instantiates a new NewPassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *NewPassword) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NewPassword) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NewPassword) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetResetToken

`func (o *NewPassword) GetResetToken() string`

GetResetToken returns the ResetToken field if non-nil, zero value otherwise.

### GetResetTokenOk

`func (o *NewPassword) GetResetTokenOk() (*string, bool)`

GetResetTokenOk returns a tuple with the ResetToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetToken

`func (o *NewPassword) SetResetToken(v string)`

SetResetToken sets ResetToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


