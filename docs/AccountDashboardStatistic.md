# AccountDashboardStatistic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revenue** | Pointer to [**[]AccountDashboardStatisticRevenue**](AccountDashboardStatisticRevenue.md) | Aggregated statistic for account revenue. | [optional] 
**Discounts** | Pointer to [**[]AccountDashboardStatisticDiscount**](AccountDashboardStatisticDiscount.md) | Aggregated statistic for account discount. | [optional] 
**LoyaltyPoints** | Pointer to [**[]AccountDashboardStatisticLoyaltyPoints**](AccountDashboardStatisticLoyaltyPoints.md) | Aggregated statistic for account loyalty points. | [optional] 
**Referrals** | Pointer to [**[]AccountDashboardStatisticReferrals**](AccountDashboardStatisticReferrals.md) | Aggregated statistic for account referrals. | [optional] 
**Campaigns** | Pointer to [**AccountDashboardStatisticCampaigns**](AccountDashboardStatisticCampaigns.md) |  | 

## Methods

### NewAccountDashboardStatistic

`func NewAccountDashboardStatistic(campaigns AccountDashboardStatisticCampaigns, ) *AccountDashboardStatistic`

NewAccountDashboardStatistic instantiates a new AccountDashboardStatistic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDashboardStatisticWithDefaults

`func NewAccountDashboardStatisticWithDefaults() *AccountDashboardStatistic`

NewAccountDashboardStatisticWithDefaults instantiates a new AccountDashboardStatistic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevenue

`func (o *AccountDashboardStatistic) GetRevenue() []AccountDashboardStatisticRevenue`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *AccountDashboardStatistic) GetRevenueOk() (*[]AccountDashboardStatisticRevenue, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *AccountDashboardStatistic) SetRevenue(v []AccountDashboardStatisticRevenue)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *AccountDashboardStatistic) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetDiscounts

`func (o *AccountDashboardStatistic) GetDiscounts() []AccountDashboardStatisticDiscount`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *AccountDashboardStatistic) GetDiscountsOk() (*[]AccountDashboardStatisticDiscount, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *AccountDashboardStatistic) SetDiscounts(v []AccountDashboardStatisticDiscount)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *AccountDashboardStatistic) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### GetLoyaltyPoints

`func (o *AccountDashboardStatistic) GetLoyaltyPoints() []AccountDashboardStatisticLoyaltyPoints`

GetLoyaltyPoints returns the LoyaltyPoints field if non-nil, zero value otherwise.

### GetLoyaltyPointsOk

`func (o *AccountDashboardStatistic) GetLoyaltyPointsOk() (*[]AccountDashboardStatisticLoyaltyPoints, bool)`

GetLoyaltyPointsOk returns a tuple with the LoyaltyPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyPoints

`func (o *AccountDashboardStatistic) SetLoyaltyPoints(v []AccountDashboardStatisticLoyaltyPoints)`

SetLoyaltyPoints sets LoyaltyPoints field to given value.

### HasLoyaltyPoints

`func (o *AccountDashboardStatistic) HasLoyaltyPoints() bool`

HasLoyaltyPoints returns a boolean if a field has been set.

### GetReferrals

`func (o *AccountDashboardStatistic) GetReferrals() []AccountDashboardStatisticReferrals`

GetReferrals returns the Referrals field if non-nil, zero value otherwise.

### GetReferralsOk

`func (o *AccountDashboardStatistic) GetReferralsOk() (*[]AccountDashboardStatisticReferrals, bool)`

GetReferralsOk returns a tuple with the Referrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrals

`func (o *AccountDashboardStatistic) SetReferrals(v []AccountDashboardStatisticReferrals)`

SetReferrals sets Referrals field to given value.

### HasReferrals

`func (o *AccountDashboardStatistic) HasReferrals() bool`

HasReferrals returns a boolean if a field has been set.

### GetCampaigns

`func (o *AccountDashboardStatistic) GetCampaigns() AccountDashboardStatisticCampaigns`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *AccountDashboardStatistic) GetCampaignsOk() (*AccountDashboardStatisticCampaigns, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaigns

`func (o *AccountDashboardStatistic) SetCampaigns(v AccountDashboardStatisticCampaigns)`

SetCampaigns sets Campaigns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


