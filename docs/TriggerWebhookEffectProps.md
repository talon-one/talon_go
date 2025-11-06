# TriggerWebhookEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookId** | Pointer to **float32** | The ID of the webhook that was triggered. | 
**WebhookName** | Pointer to **string** | The name of the webhook that was triggered. | 

## Methods

### NewTriggerWebhookEffectProps

`func NewTriggerWebhookEffectProps(webhookId float32, webhookName string, ) *TriggerWebhookEffectProps`

NewTriggerWebhookEffectProps instantiates a new TriggerWebhookEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerWebhookEffectPropsWithDefaults

`func NewTriggerWebhookEffectPropsWithDefaults() *TriggerWebhookEffectProps`

NewTriggerWebhookEffectPropsWithDefaults instantiates a new TriggerWebhookEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookId

`func (o *TriggerWebhookEffectProps) GetWebhookId() float32`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *TriggerWebhookEffectProps) GetWebhookIdOk() (*float32, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *TriggerWebhookEffectProps) SetWebhookId(v float32)`

SetWebhookId sets WebhookId field to given value.


### GetWebhookName

`func (o *TriggerWebhookEffectProps) GetWebhookName() string`

GetWebhookName returns the WebhookName field if non-nil, zero value otherwise.

### GetWebhookNameOk

`func (o *TriggerWebhookEffectProps) GetWebhookNameOk() (*string, bool)`

GetWebhookNameOk returns a tuple with the WebhookName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookName

`func (o *TriggerWebhookEffectProps) SetWebhookName(v string)`

SetWebhookName sets WebhookName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


