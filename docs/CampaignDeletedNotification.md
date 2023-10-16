# CampaignDeletedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Campaign** | Pointer to [**CampaignStateNotification**](CampaignStateNotification.md) |  | 
**DeletedAt** | Pointer to [**time.Time**](time.Time.md) | Time when the campaign was deleted. | 

## Methods

### GetCampaign

`func (o *CampaignDeletedNotification) GetCampaign() CampaignStateNotification`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *CampaignDeletedNotification) GetCampaignOk() (CampaignStateNotification, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaign

`func (o *CampaignDeletedNotification) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### SetCampaign

`func (o *CampaignDeletedNotification) SetCampaign(v CampaignStateNotification)`

SetCampaign gets a reference to the given CampaignStateNotification and assigns it to the Campaign field.

### GetDeletedAt

`func (o *CampaignDeletedNotification) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CampaignDeletedNotification) GetDeletedAtOk() (time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeletedAt

`func (o *CampaignDeletedNotification) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAt

`func (o *CampaignDeletedNotification) SetDeletedAt(v time.Time)`

SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


