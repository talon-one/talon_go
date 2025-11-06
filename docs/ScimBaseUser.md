# ScimBaseUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Status of the user. | [optional] 
**DisplayName** | Pointer to **string** | Display name of the user. | [optional] 
**UserName** | Pointer to **string** | Unique identifier of the user. This is usually an email address. | [optional] 
**Name** | Pointer to [**ScimBaseUserName**](ScimBaseUser_name.md) |  | [optional] 

## Methods

### NewScimBaseUser

`func NewScimBaseUser() *ScimBaseUser`

NewScimBaseUser instantiates a new ScimBaseUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimBaseUserWithDefaults

`func NewScimBaseUserWithDefaults() *ScimBaseUser`

NewScimBaseUserWithDefaults instantiates a new ScimBaseUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ScimBaseUser) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ScimBaseUser) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ScimBaseUser) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ScimBaseUser) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDisplayName

`func (o *ScimBaseUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ScimBaseUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ScimBaseUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ScimBaseUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetUserName

`func (o *ScimBaseUser) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ScimBaseUser) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ScimBaseUser) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ScimBaseUser) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetName

`func (o *ScimBaseUser) GetName() ScimBaseUserName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScimBaseUser) GetNameOk() (*ScimBaseUserName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScimBaseUser) SetName(v ScimBaseUserName)`

SetName sets Name field to given value.

### HasName

`func (o *ScimBaseUser) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


