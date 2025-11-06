# BaseNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to [**map[string]interface{}**](.md) | Indicates which notification properties to apply. | 
**Enabled** | Pointer to **bool** | Indicates whether the notification is activated. | [optional] [default to true]
**Webhook** | Pointer to [**BaseNotificationWebhook**](BaseNotificationWebhook.md) |  | 
**Id** | Pointer to **int64** | Unique ID for this entity. | 
**Type** | Pointer to **string** | The notification type. | 

## Methods

### NewBaseNotification

`func NewBaseNotification(policy map[string]interface{}, webhook BaseNotificationWebhook, id int64, type_ string, ) *BaseNotification`

NewBaseNotification instantiates a new BaseNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseNotificationWithDefaults

`func NewBaseNotificationWithDefaults() *BaseNotification`

NewBaseNotificationWithDefaults instantiates a new BaseNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *BaseNotification) GetPolicy() map[string]interface{}`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *BaseNotification) GetPolicyOk() (*map[string]interface{}, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *BaseNotification) SetPolicy(v map[string]interface{})`

SetPolicy sets Policy field to given value.


### GetEnabled

`func (o *BaseNotification) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BaseNotification) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BaseNotification) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BaseNotification) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetWebhook

`func (o *BaseNotification) GetWebhook() BaseNotificationWebhook`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *BaseNotification) GetWebhookOk() (*BaseNotificationWebhook, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *BaseNotification) SetWebhook(v BaseNotificationWebhook)`

SetWebhook sets Webhook field to given value.


### GetId

`func (o *BaseNotification) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseNotification) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseNotification) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *BaseNotification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseNotification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseNotification) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


