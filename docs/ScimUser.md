# ScimUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Status of the user. | [optional] 
**DisplayName** | Pointer to **string** | Display name of the user. | [optional] 
**UserName** | Pointer to **string** | Unique identifier of the user. This is usually an email address. | [optional] 
**Name** | Pointer to [**ScimBaseUserName**](ScimBaseUser_name.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the user. | 

## Methods

### GetActive

`func (o *ScimUser) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ScimUser) GetActiveOk() (bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActive

`func (o *ScimUser) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActive

`func (o *ScimUser) SetActive(v bool)`

SetActive gets a reference to the given bool and assigns it to the Active field.

### GetDisplayName

`func (o *ScimUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ScimUser) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *ScimUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *ScimUser) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetUserName

`func (o *ScimUser) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ScimUser) GetUserNameOk() (string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserName

`func (o *ScimUser) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserName

`func (o *ScimUser) SetUserName(v string)`

SetUserName gets a reference to the given string and assigns it to the UserName field.

### GetName

`func (o *ScimUser) GetName() ScimBaseUserName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScimUser) GetNameOk() (ScimBaseUserName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *ScimUser) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *ScimUser) SetName(v ScimBaseUserName)`

SetName gets a reference to the given ScimBaseUserName and assigns it to the Name field.

### GetId

`func (o *ScimUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScimUser) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *ScimUser) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *ScimUser) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


