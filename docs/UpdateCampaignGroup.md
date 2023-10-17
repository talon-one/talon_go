# UpdateCampaignGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this campaign access group. | 
**Description** | Pointer to **string** | A longer description of the campaign access group. | [optional] 
**SubscribedApplicationsIds** | Pointer to **[]int32** | A list of IDs of the Applications that this campaign access group is enabled for. | [optional] 
**CampaignIds** | Pointer to **[]int32** | A list of IDs of the campaigns that are part of the campaign access group. | [optional] 

## Methods

### GetName

`func (o *UpdateCampaignGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCampaignGroup) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *UpdateCampaignGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *UpdateCampaignGroup) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *UpdateCampaignGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCampaignGroup) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *UpdateCampaignGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *UpdateCampaignGroup) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetSubscribedApplicationsIds

`func (o *UpdateCampaignGroup) GetSubscribedApplicationsIds() []int32`

GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field if non-nil, zero value otherwise.

### GetSubscribedApplicationsIdsOk

`func (o *UpdateCampaignGroup) GetSubscribedApplicationsIdsOk() ([]int32, bool)`

GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubscribedApplicationsIds

`func (o *UpdateCampaignGroup) HasSubscribedApplicationsIds() bool`

HasSubscribedApplicationsIds returns a boolean if a field has been set.

### SetSubscribedApplicationsIds

`func (o *UpdateCampaignGroup) SetSubscribedApplicationsIds(v []int32)`

SetSubscribedApplicationsIds gets a reference to the given []int32 and assigns it to the SubscribedApplicationsIds field.

### GetCampaignIds

`func (o *UpdateCampaignGroup) GetCampaignIds() []int32`

GetCampaignIds returns the CampaignIds field if non-nil, zero value otherwise.

### GetCampaignIdsOk

`func (o *UpdateCampaignGroup) GetCampaignIdsOk() ([]int32, bool)`

GetCampaignIdsOk returns a tuple with the CampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignIds

`func (o *UpdateCampaignGroup) HasCampaignIds() bool`

HasCampaignIds returns a boolean if a field has been set.

### SetCampaignIds

`func (o *UpdateCampaignGroup) SetCampaignIds(v []int32)`

SetCampaignIds gets a reference to the given []int32 and assigns it to the CampaignIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


