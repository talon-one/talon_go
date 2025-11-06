# ScimNewUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Status of the user. | [optional] 
**DisplayName** | Pointer to **string** | Display name of the user. | [optional] 
**UserName** | Pointer to **string** | Unique identifier of the user. This is usually an email address. | 
**Name** | Pointer to [**ScimBaseUserName**](ScimBaseUser_name.md) |  | [optional] 

## Methods

### NewScimNewUser

`func NewScimNewUser(userName string, ) *ScimNewUser`

NewScimNewUser instantiates a new ScimNewUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimNewUserWithDefaults

`func NewScimNewUserWithDefaults() *ScimNewUser`

NewScimNewUserWithDefaults instantiates a new ScimNewUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ScimNewUser) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ScimNewUser) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ScimNewUser) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ScimNewUser) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDisplayName

`func (o *ScimNewUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ScimNewUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ScimNewUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ScimNewUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetUserName

`func (o *ScimNewUser) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ScimNewUser) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ScimNewUser) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetName

`func (o *ScimNewUser) GetName() ScimBaseUserName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScimNewUser) GetNameOk() (*ScimBaseUserName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScimNewUser) SetName(v ScimBaseUserName)`

SetName sets Name field to given value.

### HasName

`func (o *ScimNewUser) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


