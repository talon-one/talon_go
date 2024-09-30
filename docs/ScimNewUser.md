# ScimNewUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Status of the user. | [optional] 
**DisplayName** | Pointer to **string** | Display name of the user. | [optional] 
**UserName** | Pointer to **string** | Unique identifier of the user. This is usually an email address. | [optional] 
**Name** | Pointer to [**ScimBaseUserName**](ScimBaseUser_name.md) |  | [optional] 

## Methods

### GetActive

`func (o *ScimNewUser) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ScimNewUser) GetActiveOk() (bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActive

`func (o *ScimNewUser) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActive

`func (o *ScimNewUser) SetActive(v bool)`

SetActive gets a reference to the given bool and assigns it to the Active field.

### GetDisplayName

`func (o *ScimNewUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ScimNewUser) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *ScimNewUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *ScimNewUser) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetUserName

`func (o *ScimNewUser) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ScimNewUser) GetUserNameOk() (string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserName

`func (o *ScimNewUser) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserName

`func (o *ScimNewUser) SetUserName(v string)`

SetUserName gets a reference to the given string and assigns it to the UserName field.

### GetName

`func (o *ScimNewUser) GetName() ScimBaseUserName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScimNewUser) GetNameOk() (ScimBaseUserName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *ScimNewUser) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *ScimNewUser) SetName(v ScimBaseUserName)`

SetName gets a reference to the given ScimBaseUserName and assigns it to the Name field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


