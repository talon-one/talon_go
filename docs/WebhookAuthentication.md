# WebhookAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** | The name of the user who created the webhook authentication. | 
**ModifiedBy** | Pointer to **string** | The name of the user who last modified the webhook authentication. | 
**Webhooks** | Pointer to [**[]WebhookAuthenticationWebhookRef**](WebhookAuthenticationWebhookRef.md) |  | 
**Name** | Pointer to **string** | The name of the webhook authentication. | 
**Type** | Pointer to **string** |  | 
**Data** | Pointer to [**map[string]interface{}**](.md) |  | 
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 

## Methods

### NewWebhookAuthentication

`func NewWebhookAuthentication(createdBy string, modifiedBy string, webhooks []WebhookAuthenticationWebhookRef, name string, type_ string, data map[string]interface{}, id int64, created time.Time, modified time.Time, ) *WebhookAuthentication`

NewWebhookAuthentication instantiates a new WebhookAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookAuthenticationWithDefaults

`func NewWebhookAuthenticationWithDefaults() *WebhookAuthentication`

NewWebhookAuthenticationWithDefaults instantiates a new WebhookAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *WebhookAuthentication) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WebhookAuthentication) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WebhookAuthentication) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedBy

`func (o *WebhookAuthentication) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *WebhookAuthentication) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *WebhookAuthentication) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetWebhooks

`func (o *WebhookAuthentication) GetWebhooks() []WebhookAuthenticationWebhookRef`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *WebhookAuthentication) GetWebhooksOk() (*[]WebhookAuthenticationWebhookRef, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *WebhookAuthentication) SetWebhooks(v []WebhookAuthenticationWebhookRef)`

SetWebhooks sets Webhooks field to given value.


### GetName

`func (o *WebhookAuthentication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookAuthentication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookAuthentication) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *WebhookAuthentication) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookAuthentication) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookAuthentication) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *WebhookAuthentication) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookAuthentication) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookAuthentication) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetId

`func (o *WebhookAuthentication) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookAuthentication) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookAuthentication) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *WebhookAuthentication) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WebhookAuthentication) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WebhookAuthentication) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *WebhookAuthentication) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *WebhookAuthentication) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *WebhookAuthentication) SetModified(v time.Time)`

SetModified sets Modified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


