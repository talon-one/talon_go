# WebhookAuthenticationWebhookRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Title** | Pointer to **string** | The title of the webhook authentication. | 
**Description** | Pointer to **string** | A description of the webhook authentication. | [optional] 

## Methods

### NewWebhookAuthenticationWebhookRef

`func NewWebhookAuthenticationWebhookRef(id int64, title string, ) *WebhookAuthenticationWebhookRef`

NewWebhookAuthenticationWebhookRef instantiates a new WebhookAuthenticationWebhookRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookAuthenticationWebhookRefWithDefaults

`func NewWebhookAuthenticationWebhookRefWithDefaults() *WebhookAuthenticationWebhookRef`

NewWebhookAuthenticationWebhookRefWithDefaults instantiates a new WebhookAuthenticationWebhookRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookAuthenticationWebhookRef) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookAuthenticationWebhookRef) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookAuthenticationWebhookRef) SetId(v int64)`

SetId sets Id field to given value.


### GetTitle

`func (o *WebhookAuthenticationWebhookRef) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebhookAuthenticationWebhookRef) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebhookAuthenticationWebhookRef) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *WebhookAuthenticationWebhookRef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookAuthenticationWebhookRef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookAuthenticationWebhookRef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookAuthenticationWebhookRef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


