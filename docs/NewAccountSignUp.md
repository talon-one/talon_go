# NewAccountSignUp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The email address associated with the user profile. | 
**Password** | Pointer to **string** | The password for your account. | 
**CompanyName** | Pointer to **string** |  | 

## Methods

### NewNewAccountSignUp

`func NewNewAccountSignUp(email string, password string, companyName string, ) *NewAccountSignUp`

NewNewAccountSignUp instantiates a new NewAccountSignUp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAccountSignUpWithDefaults

`func NewNewAccountSignUpWithDefaults() *NewAccountSignUp`

NewNewAccountSignUpWithDefaults instantiates a new NewAccountSignUp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *NewAccountSignUp) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NewAccountSignUp) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NewAccountSignUp) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *NewAccountSignUp) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NewAccountSignUp) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NewAccountSignUp) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetCompanyName

`func (o *NewAccountSignUp) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *NewAccountSignUp) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *NewAccountSignUp) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


