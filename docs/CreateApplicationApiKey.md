# CreateApplicationApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title of the API key. | 
**Expires** | Pointer to [**time.Time**](time.Time.md) | The date the API key expires. | 
**Platform** | Pointer to **string** | The third-party platform the API key is valid for. Use &#x60;none&#x60; for a generic API key to be used from your own integration layer.  | [optional] 
**Type** | Pointer to **string** | The API key type. Can be empty or &#x60;staging&#x60;.  Staging API keys can only be used for dry requests with the [Update customer session](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/updateCustomerSessionV2) endpoint, [Update customer profile](https://docs.talon.one/integration-api#tag/Customer-profiles/operation/updateCustomerProfileV2) endpoint, and [Track event](https://docs.talon.one/integration-api#tag/Events/operation/trackEventV2) endpoint.  When using the _Update customer profile_ endpoint with a staging API key, the query parameter &#x60;runRuleEngine&#x60; must be &#x60;true&#x60;.  | [optional] 
**TimeOffset** | Pointer to **int32** | A time offset in nanoseconds associated with the API key. When making a request using the API key, rule evaluation is based on a date that is calculated by adding the offset to the current date.  | [optional] 

## Methods

### GetTitle

`func (o *CreateApplicationApiKey) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateApplicationApiKey) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *CreateApplicationApiKey) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *CreateApplicationApiKey) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetExpires

`func (o *CreateApplicationApiKey) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *CreateApplicationApiKey) GetExpiresOk() (time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpires

`func (o *CreateApplicationApiKey) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpires

`func (o *CreateApplicationApiKey) SetExpires(v time.Time)`

SetExpires gets a reference to the given time.Time and assigns it to the Expires field.

### GetPlatform

`func (o *CreateApplicationApiKey) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreateApplicationApiKey) GetPlatformOk() (string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlatform

`func (o *CreateApplicationApiKey) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatform

`func (o *CreateApplicationApiKey) SetPlatform(v string)`

SetPlatform gets a reference to the given string and assigns it to the Platform field.

### GetType

`func (o *CreateApplicationApiKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateApplicationApiKey) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *CreateApplicationApiKey) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *CreateApplicationApiKey) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetTimeOffset

`func (o *CreateApplicationApiKey) GetTimeOffset() int32`

GetTimeOffset returns the TimeOffset field if non-nil, zero value otherwise.

### GetTimeOffsetOk

`func (o *CreateApplicationApiKey) GetTimeOffsetOk() (int32, bool)`

GetTimeOffsetOk returns a tuple with the TimeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimeOffset

`func (o *CreateApplicationApiKey) HasTimeOffset() bool`

HasTimeOffset returns a boolean if a field has been set.

### SetTimeOffset

`func (o *CreateApplicationApiKey) SetTimeOffset(v int32)`

SetTimeOffset gets a reference to the given int32 and assigns it to the TimeOffset field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


