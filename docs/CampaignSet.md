# CampaignSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was created. | 
**ApplicationId** | Pointer to **int32** | The ID of the application that owns this entity. | 
**Version** | Pointer to **int32** | Version of the campaign set. | 
**Set** | Pointer to [**CampaignSetBranchNode**](CampaignSetBranchNode.md) |  | 

## Methods

### GetId

`func (o *CampaignSet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignSet) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *CampaignSet) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *CampaignSet) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *CampaignSet) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CampaignSet) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *CampaignSet) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *CampaignSet) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetApplicationId

`func (o *CampaignSet) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CampaignSet) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *CampaignSet) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *CampaignSet) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetVersion

`func (o *CampaignSet) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CampaignSet) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *CampaignSet) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *CampaignSet) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetSet

`func (o *CampaignSet) GetSet() CampaignSetBranchNode`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *CampaignSet) GetSetOk() (CampaignSetBranchNode, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSet

`func (o *CampaignSet) HasSet() bool`

HasSet returns a boolean if a field has been set.

### SetSet

`func (o *CampaignSet) SetSet(v CampaignSetBranchNode)`

SetSet gets a reference to the given CampaignSetBranchNode and assigns it to the Set field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


