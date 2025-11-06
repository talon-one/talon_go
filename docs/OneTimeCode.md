# OneTimeCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int64** | The ID of the user. | 
**AccountId** | Pointer to **int64** | The ID of the account. | 
**Token** | Pointer to **string** | The two-factor authentication token created during sign-in. This token is used to ensure that the correct user is trying to sign in with a given one-time security code. | 
**Code** | Pointer to **string** | The one-time security code used for signing in. | [optional] 

## Methods

### NewOneTimeCode

`func NewOneTimeCode(userId int64, accountId int64, token string, ) *OneTimeCode`

NewOneTimeCode instantiates a new OneTimeCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneTimeCodeWithDefaults

`func NewOneTimeCodeWithDefaults() *OneTimeCode`

NewOneTimeCodeWithDefaults instantiates a new OneTimeCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *OneTimeCode) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OneTimeCode) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OneTimeCode) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetAccountId

`func (o *OneTimeCode) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *OneTimeCode) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *OneTimeCode) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetToken

`func (o *OneTimeCode) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OneTimeCode) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OneTimeCode) SetToken(v string)`

SetToken sets Token field to given value.


### GetCode

`func (o *OneTimeCode) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OneTimeCode) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OneTimeCode) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *OneTimeCode) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


