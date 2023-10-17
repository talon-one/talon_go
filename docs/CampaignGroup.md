# CampaignGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 
**AccountId** | Pointer to **int32** | The ID of the account that owns this entity. | 
**Name** | Pointer to **string** | The name of this campaign access group. | 
**Description** | Pointer to **string** | A longer description of the campaign access group. | [optional] 
**SubscribedApplicationsIds** | Pointer to **[]int32** | A list of IDs of the Applications that this campaign access group is enabled for. | [optional] 
**CampaignIds** | Pointer to **[]int32** | A list of IDs of the campaigns that are part of the campaign access group. | [optional] 

## Methods

### GetId

`func (o *CampaignGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignGroup) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *CampaignGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *CampaignGroup) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *CampaignGroup) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CampaignGroup) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *CampaignGroup) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *CampaignGroup) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetModified

`func (o *CampaignGroup) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *CampaignGroup) GetModifiedOk() (time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *CampaignGroup) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *CampaignGroup) SetModified(v time.Time)`

SetModified gets a reference to the given time.Time and assigns it to the Modified field.

### GetAccountId

`func (o *CampaignGroup) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CampaignGroup) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *CampaignGroup) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *CampaignGroup) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetName

`func (o *CampaignGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignGroup) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *CampaignGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *CampaignGroup) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *CampaignGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignGroup) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *CampaignGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *CampaignGroup) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetSubscribedApplicationsIds

`func (o *CampaignGroup) GetSubscribedApplicationsIds() []int32`

GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field if non-nil, zero value otherwise.

### GetSubscribedApplicationsIdsOk

`func (o *CampaignGroup) GetSubscribedApplicationsIdsOk() ([]int32, bool)`

GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubscribedApplicationsIds

`func (o *CampaignGroup) HasSubscribedApplicationsIds() bool`

HasSubscribedApplicationsIds returns a boolean if a field has been set.

### SetSubscribedApplicationsIds

`func (o *CampaignGroup) SetSubscribedApplicationsIds(v []int32)`

SetSubscribedApplicationsIds gets a reference to the given []int32 and assigns it to the SubscribedApplicationsIds field.

### GetCampaignIds

`func (o *CampaignGroup) GetCampaignIds() []int32`

GetCampaignIds returns the CampaignIds field if non-nil, zero value otherwise.

### GetCampaignIdsOk

`func (o *CampaignGroup) GetCampaignIdsOk() ([]int32, bool)`

GetCampaignIdsOk returns a tuple with the CampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignIds

`func (o *CampaignGroup) HasCampaignIds() bool`

HasCampaignIds returns a boolean if a field has been set.

### SetCampaignIds

`func (o *CampaignGroup) SetCampaignIds(v []int32)`

SetCampaignIds gets a reference to the given []int32 and assigns it to the CampaignIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


