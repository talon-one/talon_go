# CampaignSetV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**ApplicationId** | Pointer to **int32** | The ID of the application that owns this entity. | 
**Version** | Pointer to **int32** | Version of the campaign set. | 
**Set** | Pointer to [**CampaignPrioritiesV2**](CampaignPrioritiesV2.md) |  | 

## Methods

### GetId

`func (o *CampaignSetV2) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignSetV2) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *CampaignSetV2) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *CampaignSetV2) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *CampaignSetV2) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CampaignSetV2) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *CampaignSetV2) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *CampaignSetV2) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetApplicationId

`func (o *CampaignSetV2) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CampaignSetV2) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *CampaignSetV2) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *CampaignSetV2) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetVersion

`func (o *CampaignSetV2) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CampaignSetV2) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *CampaignSetV2) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *CampaignSetV2) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetSet

`func (o *CampaignSetV2) GetSet() CampaignPrioritiesV2`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *CampaignSetV2) GetSetOk() (CampaignPrioritiesV2, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSet

`func (o *CampaignSetV2) HasSet() bool`

HasSet returns a boolean if a field has been set.

### SetSet

`func (o *CampaignSetV2) SetSet(v CampaignPrioritiesV2)`

SetSet gets a reference to the given CampaignPrioritiesV2 and assigns it to the Set field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


