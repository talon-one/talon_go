# StoreAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**ApplicationId** | Pointer to **int32** | The ID of the application that owns this entity. | 
**Updated** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent update on this entity. | 
**LinkedCampaignIds** | Pointer to **[]int32** | A list of IDs of the campaigns that are linked with current store. | [optional] 

## Methods

### GetCreated

`func (o *StoreAllOf) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StoreAllOf) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *StoreAllOf) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *StoreAllOf) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetApplicationId

`func (o *StoreAllOf) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *StoreAllOf) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *StoreAllOf) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *StoreAllOf) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetUpdated

`func (o *StoreAllOf) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *StoreAllOf) GetUpdatedOk() (time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *StoreAllOf) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *StoreAllOf) SetUpdated(v time.Time)`

SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.

### GetLinkedCampaignIds

`func (o *StoreAllOf) GetLinkedCampaignIds() []int32`

GetLinkedCampaignIds returns the LinkedCampaignIds field if non-nil, zero value otherwise.

### GetLinkedCampaignIdsOk

`func (o *StoreAllOf) GetLinkedCampaignIdsOk() ([]int32, bool)`

GetLinkedCampaignIdsOk returns a tuple with the LinkedCampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkedCampaignIds

`func (o *StoreAllOf) HasLinkedCampaignIds() bool`

HasLinkedCampaignIds returns a boolean if a field has been set.

### SetLinkedCampaignIds

`func (o *StoreAllOf) SetLinkedCampaignIds(v []int32)`

SetLinkedCampaignIds gets a reference to the given []int32 and assigns it to the LinkedCampaignIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


