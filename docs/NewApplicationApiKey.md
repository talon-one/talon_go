# NewApplicationApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title of the API key. | 
**Expires** | Pointer to [**time.Time**](time.Time.md) | The date the API key expires. | 
**Platform** | Pointer to **string** | The third-party platform the API key is valid for. Use &#x60;none&#x60; for a generic API key to be used from your own integration layer.  | [optional] 
**Type** | Pointer to **string** | The API key type. Can be empty or &#x60;staging&#x60;.  Staging API keys can only be used for dry requests with the [Update customer session](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/updateCustomerSessionV2) endpoint, [Update customer profile](https://docs.talon.one/integration-api#tag/Customer-profiles/operation/updateCustomerProfileV2) endpoint, and [Track event](https://docs.talon.one/integration-api#tag/Events/operation/trackEventV2) endpoint.  When using the _Update customer profile_ endpoint with a staging API key, the query parameter &#x60;runRuleEngine&#x60; must be &#x60;true&#x60;.  | [optional] 
**TimeOffset** | Pointer to **int32** | A time offset in nanoseconds associated with the API key. When making a request using the API key, rule evaluation is based on a date that is calculated by adding the offset to the current date.  | [optional] 
**Id** | Pointer to **int32** | ID of the API Key. | 
**CreatedBy** | Pointer to **int32** | ID of user who created. | 
**AccountID** | Pointer to **int32** | ID of account the key is used for. | 
**ApplicationID** | Pointer to **int32** | ID of application the key is used for. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The date the API key was created. | 
**Key** | Pointer to **string** | The API key. | 

## Methods

### GetTitle

`func (o *NewApplicationApiKey) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewApplicationApiKey) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *NewApplicationApiKey) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *NewApplicationApiKey) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetExpires

`func (o *NewApplicationApiKey) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *NewApplicationApiKey) GetExpiresOk() (time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpires

`func (o *NewApplicationApiKey) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpires

`func (o *NewApplicationApiKey) SetExpires(v time.Time)`

SetExpires gets a reference to the given time.Time and assigns it to the Expires field.

### GetPlatform

`func (o *NewApplicationApiKey) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *NewApplicationApiKey) GetPlatformOk() (string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlatform

`func (o *NewApplicationApiKey) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatform

`func (o *NewApplicationApiKey) SetPlatform(v string)`

SetPlatform gets a reference to the given string and assigns it to the Platform field.

### GetType

`func (o *NewApplicationApiKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewApplicationApiKey) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *NewApplicationApiKey) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *NewApplicationApiKey) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetTimeOffset

`func (o *NewApplicationApiKey) GetTimeOffset() int32`

GetTimeOffset returns the TimeOffset field if non-nil, zero value otherwise.

### GetTimeOffsetOk

`func (o *NewApplicationApiKey) GetTimeOffsetOk() (int32, bool)`

GetTimeOffsetOk returns a tuple with the TimeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimeOffset

`func (o *NewApplicationApiKey) HasTimeOffset() bool`

HasTimeOffset returns a boolean if a field has been set.

### SetTimeOffset

`func (o *NewApplicationApiKey) SetTimeOffset(v int32)`

SetTimeOffset gets a reference to the given int32 and assigns it to the TimeOffset field.

### GetId

`func (o *NewApplicationApiKey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NewApplicationApiKey) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *NewApplicationApiKey) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *NewApplicationApiKey) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreatedBy

`func (o *NewApplicationApiKey) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *NewApplicationApiKey) GetCreatedByOk() (int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *NewApplicationApiKey) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *NewApplicationApiKey) SetCreatedBy(v int32)`

SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.

### GetAccountID

`func (o *NewApplicationApiKey) GetAccountID() int32`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *NewApplicationApiKey) GetAccountIDOk() (int32, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountID

`func (o *NewApplicationApiKey) HasAccountID() bool`

HasAccountID returns a boolean if a field has been set.

### SetAccountID

`func (o *NewApplicationApiKey) SetAccountID(v int32)`

SetAccountID gets a reference to the given int32 and assigns it to the AccountID field.

### GetApplicationID

`func (o *NewApplicationApiKey) GetApplicationID() int32`

GetApplicationID returns the ApplicationID field if non-nil, zero value otherwise.

### GetApplicationIDOk

`func (o *NewApplicationApiKey) GetApplicationIDOk() (int32, bool)`

GetApplicationIDOk returns a tuple with the ApplicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationID

`func (o *NewApplicationApiKey) HasApplicationID() bool`

HasApplicationID returns a boolean if a field has been set.

### SetApplicationID

`func (o *NewApplicationApiKey) SetApplicationID(v int32)`

SetApplicationID gets a reference to the given int32 and assigns it to the ApplicationID field.

### GetCreated

`func (o *NewApplicationApiKey) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NewApplicationApiKey) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *NewApplicationApiKey) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *NewApplicationApiKey) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetKey

`func (o *NewApplicationApiKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *NewApplicationApiKey) GetKeyOk() (string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKey

`func (o *NewApplicationApiKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKey

`func (o *NewApplicationApiKey) SetKey(v string)`

SetKey gets a reference to the given string and assigns it to the Key field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


