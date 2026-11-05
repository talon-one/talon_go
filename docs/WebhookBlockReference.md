# WebhookBlockReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The unique identifier of the webhook. | 
**Title** | Pointer to **string** | The display name of the webhook. | [optional] [readonly] 

## Methods

### NewWebhookBlockReference

`func NewWebhookBlockReference(id int64, ) *WebhookBlockReference`

NewWebhookBlockReference instantiates a new WebhookBlockReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookBlockReferenceWithDefaults

`func NewWebhookBlockReferenceWithDefaults() *WebhookBlockReference`

NewWebhookBlockReferenceWithDefaults instantiates a new WebhookBlockReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookBlockReference) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookBlockReference) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookBlockReference) SetId(v int64)`

SetId sets Id field to given value.


### GetTitle

`func (o *WebhookBlockReference) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebhookBlockReference) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebhookBlockReference) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WebhookBlockReference) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


