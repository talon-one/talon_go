# CreateApplicationAPIKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title of the API key. | 
**Expires** | Pointer to [**time.Time**](time.Time.md) | The date the API key expires. | 
**Platform** | Pointer to **string** | The third-party platform the API key is valid for. Use &#x60;none&#x60; for a generic API key to be used from your own integration layer.  | [optional] 
**Type** | Pointer to **string** | The API key type. Can be empty or &#x60;staging&#x60;.  Staging API keys can only be used for dry requests with the [Update customer session](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/updateCustomerSessionV2) endpoint, [Update customer profile](https://docs.talon.one/integration-api#tag/Customer-profiles/operation/updateCustomerProfileV2) endpoint, and [Track event](https://docs.talon.one/integration-api#tag/Events/operation/trackEventV2) endpoint.  When using the _Update customer profile_ endpoint with a staging API key, the query parameter &#x60;runRuleEngine&#x60; must be &#x60;true&#x60;.  | [optional] 
**TimeOffset** | Pointer to **int64** | A time offset in nanoseconds associated with the API key. When making a request using the API key, rule evaluation is based on a date that is calculated by adding the offset to the current date.  | [optional] 

## Methods

### NewCreateApplicationAPIKey

`func NewCreateApplicationAPIKey(title string, expires time.Time, ) *CreateApplicationAPIKey`

NewCreateApplicationAPIKey instantiates a new CreateApplicationAPIKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplicationAPIKeyWithDefaults

`func NewCreateApplicationAPIKeyWithDefaults() *CreateApplicationAPIKey`

NewCreateApplicationAPIKeyWithDefaults instantiates a new CreateApplicationAPIKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateApplicationAPIKey) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateApplicationAPIKey) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateApplicationAPIKey) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetExpires

`func (o *CreateApplicationAPIKey) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *CreateApplicationAPIKey) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *CreateApplicationAPIKey) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.


### GetPlatform

`func (o *CreateApplicationAPIKey) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreateApplicationAPIKey) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreateApplicationAPIKey) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CreateApplicationAPIKey) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetType

`func (o *CreateApplicationAPIKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateApplicationAPIKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateApplicationAPIKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateApplicationAPIKey) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTimeOffset

`func (o *CreateApplicationAPIKey) GetTimeOffset() int64`

GetTimeOffset returns the TimeOffset field if non-nil, zero value otherwise.

### GetTimeOffsetOk

`func (o *CreateApplicationAPIKey) GetTimeOffsetOk() (*int64, bool)`

GetTimeOffsetOk returns a tuple with the TimeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffset

`func (o *CreateApplicationAPIKey) SetTimeOffset(v int64)`

SetTimeOffset sets TimeOffset field to given value.

### HasTimeOffset

`func (o *CreateApplicationAPIKey) HasTimeOffset() bool`

HasTimeOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


