# Store

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Name** | Pointer to **string** | The name of the store. | 
**Description** | Pointer to **string** | The description of the store. | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | The attributes of the store. | [optional] 
**IntegrationId** | Pointer to **string** | The integration ID of the store. You choose this ID when you create a store.  **Note**: You cannot edit the &#x60;integrationId&#x60; after the store has been created.  | 
**ApplicationId** | Pointer to **int32** | The ID of the Application that owns this entity. | 
**Updated** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent update on this entity. | 
**LinkedCampaignIds** | Pointer to **[]int32** | A list of IDs of the campaigns that are linked with current store. | [optional] 

## Methods

### GetId

`func (o *Store) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Store) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Store) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Store) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *Store) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Store) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Store) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Store) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetName

`func (o *Store) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Store) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Store) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Store) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Store) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Store) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Store) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Store) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetAttributes

`func (o *Store) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Store) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *Store) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *Store) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetIntegrationId

`func (o *Store) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *Store) GetIntegrationIdOk() (string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegrationId

`func (o *Store) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationId

`func (o *Store) SetIntegrationId(v string)`

SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.

### GetApplicationId

`func (o *Store) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Store) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *Store) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *Store) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetUpdated

`func (o *Store) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Store) GetUpdatedOk() (time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *Store) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *Store) SetUpdated(v time.Time)`

SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.

### GetLinkedCampaignIds

`func (o *Store) GetLinkedCampaignIds() []int32`

GetLinkedCampaignIds returns the LinkedCampaignIds field if non-nil, zero value otherwise.

### GetLinkedCampaignIdsOk

`func (o *Store) GetLinkedCampaignIdsOk() ([]int32, bool)`

GetLinkedCampaignIdsOk returns a tuple with the LinkedCampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkedCampaignIds

`func (o *Store) HasLinkedCampaignIds() bool`

HasLinkedCampaignIds returns a boolean if a field has been set.

### SetLinkedCampaignIds

`func (o *Store) SetLinkedCampaignIds(v []int32)`

SetLinkedCampaignIds gets a reference to the given []int32 and assigns it to the LinkedCampaignIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


