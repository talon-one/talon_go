# NewNotificationWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | API URL for the given webhook-based notification. | 
**Headers** | Pointer to **[]string** | List of API HTTP headers for the given webhook-based notification. | 
**Enabled** | Pointer to **bool** | Indicates whether the notification is activated. | [optional] [default to true]

## Methods

### NewNewNotificationWebhook

`func NewNewNotificationWebhook(url string, headers []string, ) *NewNotificationWebhook`

NewNewNotificationWebhook instantiates a new NewNotificationWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewNotificationWebhookWithDefaults

`func NewNewNotificationWebhookWithDefaults() *NewNotificationWebhook`

NewNewNotificationWebhookWithDefaults instantiates a new NewNotificationWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *NewNotificationWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NewNotificationWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NewNotificationWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *NewNotificationWebhook) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *NewNotificationWebhook) GetHeadersOk() (*[]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *NewNotificationWebhook) SetHeaders(v []string)`

SetHeaders sets Headers field to given value.


### GetEnabled

`func (o *NewNotificationWebhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NewNotificationWebhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NewNotificationWebhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NewNotificationWebhook) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


