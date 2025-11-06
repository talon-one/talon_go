# AccountDashboardStatisticCampaigns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Live** | Pointer to **int64** | Number of campaigns that are active and live (across all Applications). | 
**EndingSoon** | Pointer to **int64** | Campaigns scheduled to expire sometime in the next 7 days. | 
**LowOnBudget** | Pointer to **int64** | Campaigns with less than 10% of budget left. | 

## Methods

### NewAccountDashboardStatisticCampaigns

`func NewAccountDashboardStatisticCampaigns(live int64, endingSoon int64, lowOnBudget int64, ) *AccountDashboardStatisticCampaigns`

NewAccountDashboardStatisticCampaigns instantiates a new AccountDashboardStatisticCampaigns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDashboardStatisticCampaignsWithDefaults

`func NewAccountDashboardStatisticCampaignsWithDefaults() *AccountDashboardStatisticCampaigns`

NewAccountDashboardStatisticCampaignsWithDefaults instantiates a new AccountDashboardStatisticCampaigns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLive

`func (o *AccountDashboardStatisticCampaigns) GetLive() int64`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *AccountDashboardStatisticCampaigns) GetLiveOk() (*int64, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *AccountDashboardStatisticCampaigns) SetLive(v int64)`

SetLive sets Live field to given value.


### GetEndingSoon

`func (o *AccountDashboardStatisticCampaigns) GetEndingSoon() int64`

GetEndingSoon returns the EndingSoon field if non-nil, zero value otherwise.

### GetEndingSoonOk

`func (o *AccountDashboardStatisticCampaigns) GetEndingSoonOk() (*int64, bool)`

GetEndingSoonOk returns a tuple with the EndingSoon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingSoon

`func (o *AccountDashboardStatisticCampaigns) SetEndingSoon(v int64)`

SetEndingSoon sets EndingSoon field to given value.


### GetLowOnBudget

`func (o *AccountDashboardStatisticCampaigns) GetLowOnBudget() int64`

GetLowOnBudget returns the LowOnBudget field if non-nil, zero value otherwise.

### GetLowOnBudgetOk

`func (o *AccountDashboardStatisticCampaigns) GetLowOnBudgetOk() (*int64, bool)`

GetLowOnBudgetOk returns a tuple with the LowOnBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowOnBudget

`func (o *AccountDashboardStatisticCampaigns) SetLowOnBudget(v int64)`

SetLowOnBudget sets LowOnBudget field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


