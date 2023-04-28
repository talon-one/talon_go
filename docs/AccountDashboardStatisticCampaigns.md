# AccountDashboardStatisticCampaigns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Live** | Pointer to **int32** | Number of campaigns that are active and live (across all Applications). | 
**EndingSoon** | Pointer to **int32** | Campaigns with a schedule ending in 7 days or with only 10% of budget left. | 

## Methods

### GetLive

`func (o *AccountDashboardStatisticCampaigns) GetLive() int32`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *AccountDashboardStatisticCampaigns) GetLiveOk() (int32, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLive

`func (o *AccountDashboardStatisticCampaigns) HasLive() bool`

HasLive returns a boolean if a field has been set.

### SetLive

`func (o *AccountDashboardStatisticCampaigns) SetLive(v int32)`

SetLive gets a reference to the given int32 and assigns it to the Live field.

### GetEndingSoon

`func (o *AccountDashboardStatisticCampaigns) GetEndingSoon() int32`

GetEndingSoon returns the EndingSoon field if non-nil, zero value otherwise.

### GetEndingSoonOk

`func (o *AccountDashboardStatisticCampaigns) GetEndingSoonOk() (int32, bool)`

GetEndingSoonOk returns a tuple with the EndingSoon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndingSoon

`func (o *AccountDashboardStatisticCampaigns) HasEndingSoon() bool`

HasEndingSoon returns a boolean if a field has been set.

### SetEndingSoon

`func (o *AccountDashboardStatisticCampaigns) SetEndingSoon(v int32)`

SetEndingSoon gets a reference to the given int32 and assigns it to the EndingSoon field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


