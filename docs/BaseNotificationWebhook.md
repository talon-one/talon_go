# BaseNotificationWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 
**Url** | Pointer to **string** | API URL for the given webhook-based notification. | 
**Headers** | Pointer to **[]string** | List of API HTTP headers for the given webhook-based notification. | 
**Enabled** | Pointer to **bool** | Indicates whether the notification is activated. | [optional] [default to true]

## Methods

### NewBaseNotificationWebhook

`func NewBaseNotificationWebhook(id int64, created time.Time, modified time.Time, url string, headers []string, ) *BaseNotificationWebhook`

NewBaseNotificationWebhook instantiates a new BaseNotificationWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseNotificationWebhookWithDefaults

`func NewBaseNotificationWebhookWithDefaults() *BaseNotificationWebhook`

NewBaseNotificationWebhookWithDefaults instantiates a new BaseNotificationWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseNotificationWebhook) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseNotificationWebhook) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseNotificationWebhook) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *BaseNotificationWebhook) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BaseNotificationWebhook) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BaseNotificationWebhook) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *BaseNotificationWebhook) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *BaseNotificationWebhook) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *BaseNotificationWebhook) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetUrl

`func (o *BaseNotificationWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BaseNotificationWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BaseNotificationWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *BaseNotificationWebhook) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *BaseNotificationWebhook) GetHeadersOk() (*[]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *BaseNotificationWebhook) SetHeaders(v []string)`

SetHeaders sets Headers field to given value.


### GetEnabled

`func (o *BaseNotificationWebhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BaseNotificationWebhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BaseNotificationWebhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BaseNotificationWebhook) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


