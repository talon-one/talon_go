# WebhookActivationLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationRequestUuid** | Pointer to **string** | UUID reference of the integration request that triggered the effect with the webhook | 
**WebhookId** | Pointer to **int32** | ID of the webhook that triggered the request | 
**ApplicationId** | Pointer to **int32** | ID of the application that triggered the webhook | 
**CampaignId** | Pointer to **int32** | ID of the campaign that triggered the webhook | 
**Created** | Pointer to [**time.Time**](time.Time.md) | Timestamp of request | 

## Methods

### GetIntegrationRequestUuid

`func (o *WebhookActivationLogEntry) GetIntegrationRequestUuid() string`

GetIntegrationRequestUuid returns the IntegrationRequestUuid field if non-nil, zero value otherwise.

### GetIntegrationRequestUuidOk

`func (o *WebhookActivationLogEntry) GetIntegrationRequestUuidOk() (string, bool)`

GetIntegrationRequestUuidOk returns a tuple with the IntegrationRequestUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegrationRequestUuid

`func (o *WebhookActivationLogEntry) HasIntegrationRequestUuid() bool`

HasIntegrationRequestUuid returns a boolean if a field has been set.

### SetIntegrationRequestUuid

`func (o *WebhookActivationLogEntry) SetIntegrationRequestUuid(v string)`

SetIntegrationRequestUuid gets a reference to the given string and assigns it to the IntegrationRequestUuid field.

### GetWebhookId

`func (o *WebhookActivationLogEntry) GetWebhookId() int32`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookActivationLogEntry) GetWebhookIdOk() (int32, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebhookId

`func (o *WebhookActivationLogEntry) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### SetWebhookId

`func (o *WebhookActivationLogEntry) SetWebhookId(v int32)`

SetWebhookId gets a reference to the given int32 and assigns it to the WebhookId field.

### GetApplicationId

`func (o *WebhookActivationLogEntry) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *WebhookActivationLogEntry) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *WebhookActivationLogEntry) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *WebhookActivationLogEntry) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetCampaignId

`func (o *WebhookActivationLogEntry) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *WebhookActivationLogEntry) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *WebhookActivationLogEntry) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *WebhookActivationLogEntry) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.

### GetCreated

`func (o *WebhookActivationLogEntry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WebhookActivationLogEntry) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *WebhookActivationLogEntry) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *WebhookActivationLogEntry) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


