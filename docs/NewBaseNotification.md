# NewBaseNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to [**map[string]interface{}**](.md) | Indicates which notification properties to apply. | 
**Enabled** | Pointer to **bool** | Indicates whether the notification is activated. | [optional] [default to true]
**Webhook** | Pointer to [**NewNotificationWebhook**](NewNotificationWebhook.md) |  | 

## Methods

### NewNewBaseNotification

`func NewNewBaseNotification(policy map[string]interface{}, webhook NewNotificationWebhook, ) *NewBaseNotification`

NewNewBaseNotification instantiates a new NewBaseNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewBaseNotificationWithDefaults

`func NewNewBaseNotificationWithDefaults() *NewBaseNotification`

NewNewBaseNotificationWithDefaults instantiates a new NewBaseNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *NewBaseNotification) GetPolicy() map[string]interface{}`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *NewBaseNotification) GetPolicyOk() (*map[string]interface{}, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *NewBaseNotification) SetPolicy(v map[string]interface{})`

SetPolicy sets Policy field to given value.


### GetEnabled

`func (o *NewBaseNotification) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NewBaseNotification) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NewBaseNotification) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NewBaseNotification) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetWebhook

`func (o *NewBaseNotification) GetWebhook() NewNotificationWebhook`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *NewBaseNotification) GetWebhookOk() (*NewNotificationWebhook, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *NewBaseNotification) SetWebhook(v NewNotificationWebhook)`

SetWebhook sets Webhook field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


