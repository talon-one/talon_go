# ApplicationApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title for API Key | 
**Expires** | Pointer to [**time.Time**](time.Time.md) | The date the API key expired | 
**Platform** | Pointer to **string** | Platform the API key is valid for. | [optional] 
**Id** | Pointer to **int32** | ID of the API Key | 
**CreatedBy** | Pointer to **int32** | ID of user who created | 
**AccountID** | Pointer to **int32** | ID of account the key is used for | 
**ApplicationID** | Pointer to **int32** | ID of application the key is used for | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The date the API key was created | 

## Methods

### GetTitle

`func (o *ApplicationApiKey) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApplicationApiKey) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *ApplicationApiKey) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *ApplicationApiKey) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetExpires

`func (o *ApplicationApiKey) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ApplicationApiKey) GetExpiresOk() (time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpires

`func (o *ApplicationApiKey) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpires

`func (o *ApplicationApiKey) SetExpires(v time.Time)`

SetExpires gets a reference to the given time.Time and assigns it to the Expires field.

### GetPlatform

`func (o *ApplicationApiKey) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ApplicationApiKey) GetPlatformOk() (string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlatform

`func (o *ApplicationApiKey) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatform

`func (o *ApplicationApiKey) SetPlatform(v string)`

SetPlatform gets a reference to the given string and assigns it to the Platform field.

### GetId

`func (o *ApplicationApiKey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationApiKey) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *ApplicationApiKey) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *ApplicationApiKey) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreatedBy

`func (o *ApplicationApiKey) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ApplicationApiKey) GetCreatedByOk() (int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *ApplicationApiKey) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *ApplicationApiKey) SetCreatedBy(v int32)`

SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.

### GetAccountID

`func (o *ApplicationApiKey) GetAccountID() int32`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *ApplicationApiKey) GetAccountIDOk() (int32, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountID

`func (o *ApplicationApiKey) HasAccountID() bool`

HasAccountID returns a boolean if a field has been set.

### SetAccountID

`func (o *ApplicationApiKey) SetAccountID(v int32)`

SetAccountID gets a reference to the given int32 and assigns it to the AccountID field.

### GetApplicationID

`func (o *ApplicationApiKey) GetApplicationID() int32`

GetApplicationID returns the ApplicationID field if non-nil, zero value otherwise.

### GetApplicationIDOk

`func (o *ApplicationApiKey) GetApplicationIDOk() (int32, bool)`

GetApplicationIDOk returns a tuple with the ApplicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationID

`func (o *ApplicationApiKey) HasApplicationID() bool`

HasApplicationID returns a boolean if a field has been set.

### SetApplicationID

`func (o *ApplicationApiKey) SetApplicationID(v int32)`

SetApplicationID gets a reference to the given int32 and assigns it to the ApplicationID field.

### GetCreated

`func (o *ApplicationApiKey) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApplicationApiKey) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *ApplicationApiKey) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *ApplicationApiKey) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


