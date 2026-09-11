# TriggerWebhookBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**Webhook** | Pointer to [**TriggerWebhookBlockWebhook**](TriggerWebhookBlock_webhook.md) |  | 
**Params** | Pointer to [**map[string]interface{}**](.md) | The webhook&#39;s parameters, in configured order. Each property name is the parameter&#39;s title, lowercased with spaces replaced by underscores (for example, &#x60;Order ID&#x60; becomes &#x60;order_id&#x60;); falls back to &#x60;param_0&#x60;, &#x60;param_1&#x60;, and so on if a title is blank or collides with another. | [optional] 
**OnError** | Pointer to [**map[string][]map[string]interface{}**](array.md) | Named error handlers evaluated when a specific error occurs. | [optional] 

## Methods

### NewTriggerWebhookBlock

`func NewTriggerWebhookBlock(type_ string, webhook TriggerWebhookBlockWebhook, ) *TriggerWebhookBlock`

NewTriggerWebhookBlock instantiates a new TriggerWebhookBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerWebhookBlockWithDefaults

`func NewTriggerWebhookBlockWithDefaults() *TriggerWebhookBlock`

NewTriggerWebhookBlockWithDefaults instantiates a new TriggerWebhookBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TriggerWebhookBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TriggerWebhookBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TriggerWebhookBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TriggerWebhookBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TriggerWebhookBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TriggerWebhookBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TriggerWebhookBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *TriggerWebhookBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TriggerWebhookBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TriggerWebhookBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TriggerWebhookBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetWebhook

`func (o *TriggerWebhookBlock) GetWebhook() TriggerWebhookBlockWebhook`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *TriggerWebhookBlock) GetWebhookOk() (*TriggerWebhookBlockWebhook, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *TriggerWebhookBlock) SetWebhook(v TriggerWebhookBlockWebhook)`

SetWebhook sets Webhook field to given value.


### GetParams

`func (o *TriggerWebhookBlock) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *TriggerWebhookBlock) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *TriggerWebhookBlock) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *TriggerWebhookBlock) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetOnError

`func (o *TriggerWebhookBlock) GetOnError() map[string][]map[string]interface{}`

GetOnError returns the OnError field if non-nil, zero value otherwise.

### GetOnErrorOk

`func (o *TriggerWebhookBlock) GetOnErrorOk() (*map[string][]map[string]interface{}, bool)`

GetOnErrorOk returns a tuple with the OnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnError

`func (o *TriggerWebhookBlock) SetOnError(v map[string][]map[string]interface{})`

SetOnError sets OnError field to given value.

### HasOnError

`func (o *TriggerWebhookBlock) HasOnError() bool`

HasOnError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


