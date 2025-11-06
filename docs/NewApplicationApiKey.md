# NewApplicationAPIKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title of the API key. | 
**Expires** | Pointer to [**time.Time**](time.Time.md) | The date the API key expires. | 
**Platform** | Pointer to **string** | The third-party platform the API key is valid for. Use &#x60;none&#x60; for a generic API key to be used from your own integration layer.  | [optional] 
**Type** | Pointer to **string** | The API key type. Can be empty or &#x60;staging&#x60;.  Staging API keys can only be used for dry requests with the [Update customer session](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/updateCustomerSessionV2) endpoint, [Update customer profile](https://docs.talon.one/integration-api#tag/Customer-profiles/operation/updateCustomerProfileV2) endpoint, and [Track event](https://docs.talon.one/integration-api#tag/Events/operation/trackEventV2) endpoint.  When using the _Update customer profile_ endpoint with a staging API key, the query parameter &#x60;runRuleEngine&#x60; must be &#x60;true&#x60;.  | [optional] 
**TimeOffset** | Pointer to **int64** | A time offset in nanoseconds associated with the API key. When making a request using the API key, rule evaluation is based on a date that is calculated by adding the offset to the current date.  | [optional] 
**Id** | Pointer to **int64** | ID of the API Key. | 
**CreatedBy** | Pointer to **int64** | ID of user who created. | 
**AccountID** | Pointer to **int64** | ID of account the key is used for. | 
**ApplicationID** | Pointer to **int64** | ID of application the key is used for. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The date the API key was created. | 
**Key** | Pointer to **string** | The API key. | 

## Methods

### NewNewApplicationAPIKey

`func NewNewApplicationAPIKey(title string, expires time.Time, id int64, createdBy int64, accountID int64, applicationID int64, created time.Time, key string, ) *NewApplicationAPIKey`

NewNewApplicationAPIKey instantiates a new NewApplicationAPIKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewApplicationAPIKeyWithDefaults

`func NewNewApplicationAPIKeyWithDefaults() *NewApplicationAPIKey`

NewNewApplicationAPIKeyWithDefaults instantiates a new NewApplicationAPIKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *NewApplicationAPIKey) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewApplicationAPIKey) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewApplicationAPIKey) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetExpires

`func (o *NewApplicationAPIKey) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *NewApplicationAPIKey) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *NewApplicationAPIKey) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.


### GetPlatform

`func (o *NewApplicationAPIKey) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *NewApplicationAPIKey) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *NewApplicationAPIKey) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *NewApplicationAPIKey) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetType

`func (o *NewApplicationAPIKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewApplicationAPIKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NewApplicationAPIKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NewApplicationAPIKey) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTimeOffset

`func (o *NewApplicationAPIKey) GetTimeOffset() int64`

GetTimeOffset returns the TimeOffset field if non-nil, zero value otherwise.

### GetTimeOffsetOk

`func (o *NewApplicationAPIKey) GetTimeOffsetOk() (*int64, bool)`

GetTimeOffsetOk returns a tuple with the TimeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffset

`func (o *NewApplicationAPIKey) SetTimeOffset(v int64)`

SetTimeOffset sets TimeOffset field to given value.

### HasTimeOffset

`func (o *NewApplicationAPIKey) HasTimeOffset() bool`

HasTimeOffset returns a boolean if a field has been set.

### GetId

`func (o *NewApplicationAPIKey) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NewApplicationAPIKey) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NewApplicationAPIKey) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *NewApplicationAPIKey) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *NewApplicationAPIKey) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *NewApplicationAPIKey) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.


### GetAccountID

`func (o *NewApplicationAPIKey) GetAccountID() int64`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *NewApplicationAPIKey) GetAccountIDOk() (*int64, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *NewApplicationAPIKey) SetAccountID(v int64)`

SetAccountID sets AccountID field to given value.


### GetApplicationID

`func (o *NewApplicationAPIKey) GetApplicationID() int64`

GetApplicationID returns the ApplicationID field if non-nil, zero value otherwise.

### GetApplicationIDOk

`func (o *NewApplicationAPIKey) GetApplicationIDOk() (*int64, bool)`

GetApplicationIDOk returns a tuple with the ApplicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationID

`func (o *NewApplicationAPIKey) SetApplicationID(v int64)`

SetApplicationID sets ApplicationID field to given value.


### GetCreated

`func (o *NewApplicationAPIKey) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NewApplicationAPIKey) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NewApplicationAPIKey) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetKey

`func (o *NewApplicationAPIKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *NewApplicationAPIKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *NewApplicationAPIKey) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


