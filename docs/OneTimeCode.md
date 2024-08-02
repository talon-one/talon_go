# OneTimeCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int32** | The ID of the account. | 
**Code** | Pointer to **string** | The one-time security code used for signing in. | [optional] 
**Token** | Pointer to **string** | The two-factor authentication token created during sign-in. This token is used to ensure that the correct user is trying to sign in with a given one-time security code. | 
**UserId** | Pointer to **int32** | The ID of the user. | 

## Methods

### GetAccountId

`func (o *OneTimeCode) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *OneTimeCode) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *OneTimeCode) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *OneTimeCode) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetCode

`func (o *OneTimeCode) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OneTimeCode) GetCodeOk() (string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCode

`func (o *OneTimeCode) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCode

`func (o *OneTimeCode) SetCode(v string)`

SetCode gets a reference to the given string and assigns it to the Code field.

### GetToken

`func (o *OneTimeCode) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OneTimeCode) GetTokenOk() (string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToken

`func (o *OneTimeCode) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetToken

`func (o *OneTimeCode) SetToken(v string)`

SetToken gets a reference to the given string and assigns it to the Token field.

### GetUserId

`func (o *OneTimeCode) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OneTimeCode) GetUserIdOk() (int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *OneTimeCode) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *OneTimeCode) SetUserId(v int32)`

SetUserId gets a reference to the given int32 and assigns it to the UserId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


