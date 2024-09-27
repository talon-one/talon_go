# Revision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**ActivateAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**AccountId** | Pointer to **int32** |  | 
**ApplicationId** | Pointer to **int32** |  | 
**CampaignId** | Pointer to **int32** |  | 
**Created** | Pointer to [**time.Time**](time.Time.md) |  | 
**CreatedBy** | Pointer to **int32** |  | 
**ActivatedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ActivatedBy** | Pointer to **int32** |  | [optional] 
**CurrentVersion** | Pointer to [**RevisionVersion**](RevisionVersion.md) |  | [optional] 

## Methods

### GetId

`func (o *Revision) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Revision) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Revision) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Revision) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetActivateAt

`func (o *Revision) GetActivateAt() time.Time`

GetActivateAt returns the ActivateAt field if non-nil, zero value otherwise.

### GetActivateAtOk

`func (o *Revision) GetActivateAtOk() (time.Time, bool)`

GetActivateAtOk returns a tuple with the ActivateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivateAt

`func (o *Revision) HasActivateAt() bool`

HasActivateAt returns a boolean if a field has been set.

### SetActivateAt

`func (o *Revision) SetActivateAt(v time.Time)`

SetActivateAt gets a reference to the given time.Time and assigns it to the ActivateAt field.

### GetAccountId

`func (o *Revision) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Revision) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *Revision) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *Revision) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetApplicationId

`func (o *Revision) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Revision) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *Revision) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *Revision) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetCampaignId

`func (o *Revision) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *Revision) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *Revision) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *Revision) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.

### GetCreated

`func (o *Revision) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Revision) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Revision) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Revision) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetCreatedBy

`func (o *Revision) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Revision) GetCreatedByOk() (int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *Revision) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *Revision) SetCreatedBy(v int32)`

SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.

### GetActivatedAt

`func (o *Revision) GetActivatedAt() time.Time`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *Revision) GetActivatedAtOk() (time.Time, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivatedAt

`func (o *Revision) HasActivatedAt() bool`

HasActivatedAt returns a boolean if a field has been set.

### SetActivatedAt

`func (o *Revision) SetActivatedAt(v time.Time)`

SetActivatedAt gets a reference to the given time.Time and assigns it to the ActivatedAt field.

### GetActivatedBy

`func (o *Revision) GetActivatedBy() int32`

GetActivatedBy returns the ActivatedBy field if non-nil, zero value otherwise.

### GetActivatedByOk

`func (o *Revision) GetActivatedByOk() (int32, bool)`

GetActivatedByOk returns a tuple with the ActivatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivatedBy

`func (o *Revision) HasActivatedBy() bool`

HasActivatedBy returns a boolean if a field has been set.

### SetActivatedBy

`func (o *Revision) SetActivatedBy(v int32)`

SetActivatedBy gets a reference to the given int32 and assigns it to the ActivatedBy field.

### GetCurrentVersion

`func (o *Revision) GetCurrentVersion() RevisionVersion`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *Revision) GetCurrentVersionOk() (RevisionVersion, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrentVersion

`func (o *Revision) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### SetCurrentVersion

`func (o *Revision) SetCurrentVersion(v RevisionVersion)`

SetCurrentVersion gets a reference to the given RevisionVersion and assigns it to the CurrentVersion field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


