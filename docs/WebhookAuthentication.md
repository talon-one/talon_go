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
**Id** | Pointer to **int32** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 

## Methods

### GetCreatedBy

`func (o *WebhookAuthentication) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WebhookAuthentication) GetCreatedByOk() (string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *WebhookAuthentication) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *WebhookAuthentication) SetCreatedBy(v string)`

SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.

### GetModifiedBy

`func (o *WebhookAuthentication) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *WebhookAuthentication) GetModifiedByOk() (string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModifiedBy

`func (o *WebhookAuthentication) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### SetModifiedBy

`func (o *WebhookAuthentication) SetModifiedBy(v string)`

SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.

### GetWebhooks

`func (o *WebhookAuthentication) GetWebhooks() []WebhookAuthenticationWebhookRef`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *WebhookAuthentication) GetWebhooksOk() ([]WebhookAuthenticationWebhookRef, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebhooks

`func (o *WebhookAuthentication) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### SetWebhooks

`func (o *WebhookAuthentication) SetWebhooks(v []WebhookAuthenticationWebhookRef)`

SetWebhooks gets a reference to the given []WebhookAuthenticationWebhookRef and assigns it to the Webhooks field.

### GetName

`func (o *WebhookAuthentication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookAuthentication) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *WebhookAuthentication) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *WebhookAuthentication) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetType

`func (o *WebhookAuthentication) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookAuthentication) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *WebhookAuthentication) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *WebhookAuthentication) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetData

`func (o *WebhookAuthentication) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookAuthentication) GetDataOk() (map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasData

`func (o *WebhookAuthentication) HasData() bool`

HasData returns a boolean if a field has been set.

### SetData

`func (o *WebhookAuthentication) SetData(v map[string]interface{})`

SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.

### GetId

`func (o *WebhookAuthentication) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookAuthentication) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *WebhookAuthentication) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *WebhookAuthentication) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *WebhookAuthentication) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WebhookAuthentication) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *WebhookAuthentication) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *WebhookAuthentication) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetModified

`func (o *WebhookAuthentication) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *WebhookAuthentication) GetModifiedOk() (time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *WebhookAuthentication) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *WebhookAuthentication) SetModified(v time.Time)`

SetModified gets a reference to the given time.Time and assigns it to the Modified field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


