# CampaignStoreBudget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**CampaignId** | Pointer to **int32** | The ID of the campaign that owns this entity. | 
**StoreId** | Pointer to **int32** | The ID of the store. | 
**Limits** | Pointer to [**[]LimitConfig**](LimitConfig.md) | The set of budget limits for stores linked to the campaign. | 

## Methods

### GetId

`func (o *CampaignStoreBudget) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignStoreBudget) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *CampaignStoreBudget) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *CampaignStoreBudget) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *CampaignStoreBudget) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CampaignStoreBudget) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *CampaignStoreBudget) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *CampaignStoreBudget) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetCampaignId

`func (o *CampaignStoreBudget) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CampaignStoreBudget) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *CampaignStoreBudget) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *CampaignStoreBudget) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.

### GetStoreId

`func (o *CampaignStoreBudget) GetStoreId() int32`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *CampaignStoreBudget) GetStoreIdOk() (int32, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStoreId

`func (o *CampaignStoreBudget) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### SetStoreId

`func (o *CampaignStoreBudget) SetStoreId(v int32)`

SetStoreId gets a reference to the given int32 and assigns it to the StoreId field.

### GetLimits

`func (o *CampaignStoreBudget) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *CampaignStoreBudget) GetLimitsOk() ([]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimits

`func (o *CampaignStoreBudget) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimits

`func (o *CampaignStoreBudget) SetLimits(v []LimitConfig)`

SetLimits gets a reference to the given []LimitConfig and assigns it to the Limits field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


