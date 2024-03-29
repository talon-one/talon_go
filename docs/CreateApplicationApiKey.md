# CreateApplicationApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title for API Key. | 
**Expires** | Pointer to [**time.Time**](time.Time.md) | The date the API key expired. | 
**Platform** | Pointer to **string** | The third-party platform the API key is valid for. Use &#x60;none&#x60; for a generic API key to be used from your own integration layer.  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


