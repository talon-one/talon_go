# ApplicationCifReferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationCartItemFilterId** | Pointer to **int32** | The ID of the Application Cart Item Filter that is referenced by a campaign. | [optional] 
**Campaigns** | Pointer to [**[]CampaignDetail**](CampaignDetail.md) | Campaigns that reference a speciifc Application Cart Item Filter. | [optional] 

## Methods

### GetApplicationCartItemFilterId

`func (o *ApplicationCifReferences) GetApplicationCartItemFilterId() int32`

GetApplicationCartItemFilterId returns the ApplicationCartItemFilterId field if non-nil, zero value otherwise.

### GetApplicationCartItemFilterIdOk

`func (o *ApplicationCifReferences) GetApplicationCartItemFilterIdOk() (int32, bool)`

GetApplicationCartItemFilterIdOk returns a tuple with the ApplicationCartItemFilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationCartItemFilterId

`func (o *ApplicationCifReferences) HasApplicationCartItemFilterId() bool`

HasApplicationCartItemFilterId returns a boolean if a field has been set.

### SetApplicationCartItemFilterId

`func (o *ApplicationCifReferences) SetApplicationCartItemFilterId(v int32)`

SetApplicationCartItemFilterId gets a reference to the given int32 and assigns it to the ApplicationCartItemFilterId field.

### GetCampaigns

`func (o *ApplicationCifReferences) GetCampaigns() []CampaignDetail`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *ApplicationCifReferences) GetCampaignsOk() ([]CampaignDetail, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaigns

`func (o *ApplicationCifReferences) HasCampaigns() bool`

HasCampaigns returns a boolean if a field has been set.

### SetCampaigns

`func (o *ApplicationCifReferences) SetCampaigns(v []CampaignDetail)`

SetCampaigns gets a reference to the given []CampaignDetail and assigns it to the Campaigns field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


