# Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int32** | The ID of the user of this session. | 
**Token** | Pointer to **string** | The token to use as a bearer token to query Management API endpoints. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | Unix timestamp indicating when the session was first created. | 

## Methods

### GetUserId

`func (o *Session) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Session) GetUserIdOk() (int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *Session) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *Session) SetUserId(v int32)`

SetUserId gets a reference to the given int32 and assigns it to the UserId field.

### GetToken

`func (o *Session) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Session) GetTokenOk() (string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToken

`func (o *Session) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetToken

`func (o *Session) SetToken(v string)`

SetToken gets a reference to the given string and assigns it to the Token field.

### GetCreated

`func (o *Session) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Session) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Session) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Session) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


