# BaseNotificationWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 
**Url** | Pointer to **string** | API URL for the given webhook-based notification. | 
**Headers** | Pointer to **[]string** | List of API HTTP headers for the given webhook-based notification. | 
**Enabled** | Pointer to **bool** | Indicates whether the notification is activated. | [optional] [default to true]

## Methods

### GetId

`func (o *BaseNotificationWebhook) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseNotificationWebhook) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *BaseNotificationWebhook) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *BaseNotificationWebhook) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *BaseNotificationWebhook) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BaseNotificationWebhook) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *BaseNotificationWebhook) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *BaseNotificationWebhook) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetModified

`func (o *BaseNotificationWebhook) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *BaseNotificationWebhook) GetModifiedOk() (time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *BaseNotificationWebhook) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *BaseNotificationWebhook) SetModified(v time.Time)`

SetModified gets a reference to the given time.Time and assigns it to the Modified field.

### GetUrl

`func (o *BaseNotificationWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BaseNotificationWebhook) GetUrlOk() (string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUrl

`func (o *BaseNotificationWebhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrl

`func (o *BaseNotificationWebhook) SetUrl(v string)`

SetUrl gets a reference to the given string and assigns it to the Url field.

### GetHeaders

`func (o *BaseNotificationWebhook) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *BaseNotificationWebhook) GetHeadersOk() ([]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHeaders

`func (o *BaseNotificationWebhook) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeaders

`func (o *BaseNotificationWebhook) SetHeaders(v []string)`

SetHeaders gets a reference to the given []string and assigns it to the Headers field.

### GetEnabled

`func (o *BaseNotificationWebhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BaseNotificationWebhook) GetEnabledOk() (bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabled

`func (o *BaseNotificationWebhook) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabled

`func (o *BaseNotificationWebhook) SetEnabled(v bool)`

SetEnabled gets a reference to the given bool and assigns it to the Enabled field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


