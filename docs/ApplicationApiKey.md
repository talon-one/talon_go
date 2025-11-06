# ApplicationAPIKey

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

## Methods

### NewApplicationAPIKey

`func NewApplicationAPIKey(title string, expires time.Time, id int64, createdBy int64, accountID int64, applicationID int64, created time.Time, ) *ApplicationAPIKey`

NewApplicationAPIKey instantiates a new ApplicationAPIKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationAPIKeyWithDefaults

`func NewApplicationAPIKeyWithDefaults() *ApplicationAPIKey`

NewApplicationAPIKeyWithDefaults instantiates a new ApplicationAPIKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ApplicationAPIKey) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApplicationAPIKey) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApplicationAPIKey) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetExpires

`func (o *ApplicationAPIKey) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ApplicationAPIKey) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ApplicationAPIKey) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.


### GetPlatform

`func (o *ApplicationAPIKey) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ApplicationAPIKey) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ApplicationAPIKey) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ApplicationAPIKey) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetType

`func (o *ApplicationAPIKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationAPIKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationAPIKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApplicationAPIKey) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTimeOffset

`func (o *ApplicationAPIKey) GetTimeOffset() int64`

GetTimeOffset returns the TimeOffset field if non-nil, zero value otherwise.

### GetTimeOffsetOk

`func (o *ApplicationAPIKey) GetTimeOffsetOk() (*int64, bool)`

GetTimeOffsetOk returns a tuple with the TimeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffset

`func (o *ApplicationAPIKey) SetTimeOffset(v int64)`

SetTimeOffset sets TimeOffset field to given value.

### HasTimeOffset

`func (o *ApplicationAPIKey) HasTimeOffset() bool`

HasTimeOffset returns a boolean if a field has been set.

### GetId

`func (o *ApplicationAPIKey) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationAPIKey) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationAPIKey) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *ApplicationAPIKey) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ApplicationAPIKey) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ApplicationAPIKey) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.


### GetAccountID

`func (o *ApplicationAPIKey) GetAccountID() int64`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *ApplicationAPIKey) GetAccountIDOk() (*int64, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *ApplicationAPIKey) SetAccountID(v int64)`

SetAccountID sets AccountID field to given value.


### GetApplicationID

`func (o *ApplicationAPIKey) GetApplicationID() int64`

GetApplicationID returns the ApplicationID field if non-nil, zero value otherwise.

### GetApplicationIDOk

`func (o *ApplicationAPIKey) GetApplicationIDOk() (*int64, bool)`

GetApplicationIDOk returns a tuple with the ApplicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationID

`func (o *ApplicationAPIKey) SetApplicationID(v int64)`

SetApplicationID sets ApplicationID field to given value.


### GetCreated

`func (o *ApplicationAPIKey) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApplicationAPIKey) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApplicationAPIKey) SetCreated(v time.Time)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


