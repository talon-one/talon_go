# BaseNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to [**map[string]interface{}**](.md) | Indicates which notification properties to apply. | 
**Enabled** | Pointer to **bool** | Indicates whether the notification is activated. | [optional] [default to true]
**Webhook** | Pointer to [**BaseNotificationWebhook**](BaseNotificationWebhook.md) |  | 
**Id** | Pointer to **int32** | Unique ID for this entity. | 
**Type** | Pointer to **string** | The notification type. | 

## Methods

### GetPolicy

`func (o *BaseNotification) GetPolicy() map[string]interface{}`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *BaseNotification) GetPolicyOk() (map[string]interface{}, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPolicy

`func (o *BaseNotification) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicy

`func (o *BaseNotification) SetPolicy(v map[string]interface{})`

SetPolicy gets a reference to the given map[string]interface{} and assigns it to the Policy field.

### GetEnabled

`func (o *BaseNotification) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BaseNotification) GetEnabledOk() (bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabled

`func (o *BaseNotification) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabled

`func (o *BaseNotification) SetEnabled(v bool)`

SetEnabled gets a reference to the given bool and assigns it to the Enabled field.

### GetWebhook

`func (o *BaseNotification) GetWebhook() BaseNotificationWebhook`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *BaseNotification) GetWebhookOk() (BaseNotificationWebhook, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebhook

`func (o *BaseNotification) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### SetWebhook

`func (o *BaseNotification) SetWebhook(v BaseNotificationWebhook)`

SetWebhook gets a reference to the given BaseNotificationWebhook and assigns it to the Webhook field.

### GetId

`func (o *BaseNotification) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseNotification) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *BaseNotification) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *BaseNotification) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetType

`func (o *BaseNotification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseNotification) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *BaseNotification) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *BaseNotification) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


