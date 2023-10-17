# NotificationWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 
**ApplicationId** | Pointer to **int32** | The ID of the application that owns this entity. | 
**Url** | Pointer to **string** | API URL for the given webhook-based notification. | 
**Headers** | Pointer to **[]string** | List of API HTTP headers for the given webhook-based notification. | 
**Enabled** | Pointer to **bool** | Indicates whether the notification is activated. | [optional] [default to true]

## Methods

### GetId

`func (o *NotificationWebhook) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationWebhook) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *NotificationWebhook) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *NotificationWebhook) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *NotificationWebhook) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NotificationWebhook) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *NotificationWebhook) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *NotificationWebhook) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetModified

`func (o *NotificationWebhook) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *NotificationWebhook) GetModifiedOk() (time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *NotificationWebhook) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *NotificationWebhook) SetModified(v time.Time)`

SetModified gets a reference to the given time.Time and assigns it to the Modified field.

### GetApplicationId

`func (o *NotificationWebhook) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *NotificationWebhook) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *NotificationWebhook) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *NotificationWebhook) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetUrl

`func (o *NotificationWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationWebhook) GetUrlOk() (string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUrl

`func (o *NotificationWebhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrl

`func (o *NotificationWebhook) SetUrl(v string)`

SetUrl gets a reference to the given string and assigns it to the Url field.

### GetHeaders

`func (o *NotificationWebhook) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *NotificationWebhook) GetHeadersOk() ([]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHeaders

`func (o *NotificationWebhook) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeaders

`func (o *NotificationWebhook) SetHeaders(v []string)`

SetHeaders gets a reference to the given []string and assigns it to the Headers field.

### GetEnabled

`func (o *NotificationWebhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NotificationWebhook) GetEnabledOk() (bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabled

`func (o *NotificationWebhook) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabled

`func (o *NotificationWebhook) SetEnabled(v bool)`

SetEnabled gets a reference to the given bool and assigns it to the Enabled field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


