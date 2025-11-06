# Store

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Name** | Pointer to **string** | The name of the store. | 
**Description** | Pointer to **string** | The description of the store. | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | The attributes of the store. | [optional] 
**IntegrationId** | Pointer to **string** | The integration ID of the store. You choose this ID when you create a store.  **Note**: You cannot edit the &#x60;integrationId&#x60; after the store has been created.  | 
**ApplicationId** | Pointer to **int64** | The ID of the Application that owns this entity. | 
**Updated** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent update on this entity. | 
**LinkedCampaignIds** | Pointer to **[]int64** | A list of IDs of the campaigns that are linked with current store. | [optional] 

## Methods

### NewStore

`func NewStore(id int64, created time.Time, name string, description string, integrationId string, applicationId int64, updated time.Time, ) *Store`

NewStore instantiates a new Store object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreWithDefaults

`func NewStoreWithDefaults() *Store`

NewStoreWithDefaults instantiates a new Store object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Store) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Store) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Store) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *Store) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Store) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Store) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetName

`func (o *Store) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Store) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Store) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Store) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Store) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Store) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAttributes

`func (o *Store) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Store) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Store) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Store) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetIntegrationId

`func (o *Store) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *Store) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *Store) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetApplicationId

`func (o *Store) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Store) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *Store) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetUpdated

`func (o *Store) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Store) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Store) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### GetLinkedCampaignIds

`func (o *Store) GetLinkedCampaignIds() []int64`

GetLinkedCampaignIds returns the LinkedCampaignIds field if non-nil, zero value otherwise.

### GetLinkedCampaignIdsOk

`func (o *Store) GetLinkedCampaignIdsOk() (*[]int64, bool)`

GetLinkedCampaignIdsOk returns a tuple with the LinkedCampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedCampaignIds

`func (o *Store) SetLinkedCampaignIds(v []int64)`

SetLinkedCampaignIds sets LinkedCampaignIds field to given value.

### HasLinkedCampaignIds

`func (o *Store) HasLinkedCampaignIds() bool`

HasLinkedCampaignIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


