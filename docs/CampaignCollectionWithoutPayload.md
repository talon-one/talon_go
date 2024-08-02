# CampaignCollectionWithoutPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int32** | The ID of the account that owns this entity. | 
**ApplicationId** | Pointer to **int32** | The ID of the Application that owns this entity. | [optional] 
**CampaignId** | Pointer to **int32** | The ID of the campaign that owns this entity. | [optional] 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**CreatedBy** | Pointer to **int32** | ID of the user who created this effect. | 
**Description** | Pointer to **string** | A short description of the purpose of this collection. | [optional] 
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 
**ModifiedBy** | Pointer to **int32** | ID of the user who last updated this effect if available. | [optional] 
**Name** | Pointer to **string** | The name of this collection. | 

## Methods

### GetAccountId

`func (o *CampaignCollectionWithoutPayload) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CampaignCollectionWithoutPayload) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *CampaignCollectionWithoutPayload) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *CampaignCollectionWithoutPayload) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetApplicationId

`func (o *CampaignCollectionWithoutPayload) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CampaignCollectionWithoutPayload) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *CampaignCollectionWithoutPayload) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *CampaignCollectionWithoutPayload) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetCampaignId

`func (o *CampaignCollectionWithoutPayload) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CampaignCollectionWithoutPayload) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *CampaignCollectionWithoutPayload) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *CampaignCollectionWithoutPayload) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.

### GetCreated

`func (o *CampaignCollectionWithoutPayload) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CampaignCollectionWithoutPayload) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *CampaignCollectionWithoutPayload) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *CampaignCollectionWithoutPayload) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetCreatedBy

`func (o *CampaignCollectionWithoutPayload) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CampaignCollectionWithoutPayload) GetCreatedByOk() (int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *CampaignCollectionWithoutPayload) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *CampaignCollectionWithoutPayload) SetCreatedBy(v int32)`

SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.

### GetDescription

`func (o *CampaignCollectionWithoutPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignCollectionWithoutPayload) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *CampaignCollectionWithoutPayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *CampaignCollectionWithoutPayload) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetId

`func (o *CampaignCollectionWithoutPayload) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignCollectionWithoutPayload) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *CampaignCollectionWithoutPayload) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *CampaignCollectionWithoutPayload) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetModified

`func (o *CampaignCollectionWithoutPayload) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *CampaignCollectionWithoutPayload) GetModifiedOk() (time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *CampaignCollectionWithoutPayload) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *CampaignCollectionWithoutPayload) SetModified(v time.Time)`

SetModified gets a reference to the given time.Time and assigns it to the Modified field.

### GetModifiedBy

`func (o *CampaignCollectionWithoutPayload) GetModifiedBy() int32`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *CampaignCollectionWithoutPayload) GetModifiedByOk() (int32, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModifiedBy

`func (o *CampaignCollectionWithoutPayload) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### SetModifiedBy

`func (o *CampaignCollectionWithoutPayload) SetModifiedBy(v int32)`

SetModifiedBy gets a reference to the given int32 and assigns it to the ModifiedBy field.

### GetName

`func (o *CampaignCollectionWithoutPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignCollectionWithoutPayload) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *CampaignCollectionWithoutPayload) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *CampaignCollectionWithoutPayload) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


