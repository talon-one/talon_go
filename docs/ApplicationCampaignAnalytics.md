# ApplicationCampaignAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The start of the aggregation time frame in UTC. | 
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The end of the aggregation time frame in UTC. | 
**CampaignId** | Pointer to **int32** | The ID of the campaign. | 
**CampaignName** | Pointer to **string** | The name of the campaign. | 
**CampaignTags** | Pointer to **[]string** | A list of tags for the campaign. | 
**CampaignState** | Pointer to **string** | The state of the campaign.  **Note:** A disabled or archived campaign is not evaluated for rules or coupons.  | 
**TotalRevenue** | Pointer to [**AnalyticsDataPointWithTrendAndInfluencedRate**](AnalyticsDataPointWithTrendAndInfluencedRate.md) |  | [optional] 
**SessionsCount** | Pointer to [**AnalyticsDataPointWithTrendAndInfluencedRate**](AnalyticsDataPointWithTrendAndInfluencedRate.md) |  | [optional] 
**AvgItemsPerSession** | Pointer to [**AnalyticsDataPointWithTrendAndUplift**](AnalyticsDataPointWithTrendAndUplift.md) |  | [optional] 
**AvgSessionValue** | Pointer to [**AnalyticsDataPointWithTrendAndUplift**](AnalyticsDataPointWithTrendAndUplift.md) |  | [optional] 
**TotalDiscounts** | Pointer to [**AnalyticsDataPointWithTrend**](AnalyticsDataPointWithTrend.md) |  | [optional] 
**CouponsCount** | Pointer to [**AnalyticsDataPointWithTrend**](AnalyticsDataPointWithTrend.md) |  | [optional] 

## Methods

### GetStartTime

`func (o *ApplicationCampaignAnalytics) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ApplicationCampaignAnalytics) GetStartTimeOk() (time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartTime

`func (o *ApplicationCampaignAnalytics) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTime

`func (o *ApplicationCampaignAnalytics) SetStartTime(v time.Time)`

SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.

### GetEndTime

`func (o *ApplicationCampaignAnalytics) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ApplicationCampaignAnalytics) GetEndTimeOk() (time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndTime

`func (o *ApplicationCampaignAnalytics) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTime

`func (o *ApplicationCampaignAnalytics) SetEndTime(v time.Time)`

SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.

### GetCampaignId

`func (o *ApplicationCampaignAnalytics) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *ApplicationCampaignAnalytics) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *ApplicationCampaignAnalytics) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *ApplicationCampaignAnalytics) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.

### GetCampaignName

`func (o *ApplicationCampaignAnalytics) GetCampaignName() string`

GetCampaignName returns the CampaignName field if non-nil, zero value otherwise.

### GetCampaignNameOk

`func (o *ApplicationCampaignAnalytics) GetCampaignNameOk() (string, bool)`

GetCampaignNameOk returns a tuple with the CampaignName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignName

`func (o *ApplicationCampaignAnalytics) HasCampaignName() bool`

HasCampaignName returns a boolean if a field has been set.

### SetCampaignName

`func (o *ApplicationCampaignAnalytics) SetCampaignName(v string)`

SetCampaignName gets a reference to the given string and assigns it to the CampaignName field.

### GetCampaignTags

`func (o *ApplicationCampaignAnalytics) GetCampaignTags() []string`

GetCampaignTags returns the CampaignTags field if non-nil, zero value otherwise.

### GetCampaignTagsOk

`func (o *ApplicationCampaignAnalytics) GetCampaignTagsOk() ([]string, bool)`

GetCampaignTagsOk returns a tuple with the CampaignTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignTags

`func (o *ApplicationCampaignAnalytics) HasCampaignTags() bool`

HasCampaignTags returns a boolean if a field has been set.

### SetCampaignTags

`func (o *ApplicationCampaignAnalytics) SetCampaignTags(v []string)`

SetCampaignTags gets a reference to the given []string and assigns it to the CampaignTags field.

### GetCampaignState

`func (o *ApplicationCampaignAnalytics) GetCampaignState() string`

GetCampaignState returns the CampaignState field if non-nil, zero value otherwise.

### GetCampaignStateOk

`func (o *ApplicationCampaignAnalytics) GetCampaignStateOk() (string, bool)`

GetCampaignStateOk returns a tuple with the CampaignState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignState

`func (o *ApplicationCampaignAnalytics) HasCampaignState() bool`

HasCampaignState returns a boolean if a field has been set.

### SetCampaignState

`func (o *ApplicationCampaignAnalytics) SetCampaignState(v string)`

SetCampaignState gets a reference to the given string and assigns it to the CampaignState field.

### GetTotalRevenue

`func (o *ApplicationCampaignAnalytics) GetTotalRevenue() AnalyticsDataPointWithTrendAndInfluencedRate`

GetTotalRevenue returns the TotalRevenue field if non-nil, zero value otherwise.

### GetTotalRevenueOk

`func (o *ApplicationCampaignAnalytics) GetTotalRevenueOk() (AnalyticsDataPointWithTrendAndInfluencedRate, bool)`

GetTotalRevenueOk returns a tuple with the TotalRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalRevenue

`func (o *ApplicationCampaignAnalytics) HasTotalRevenue() bool`

HasTotalRevenue returns a boolean if a field has been set.

### SetTotalRevenue

`func (o *ApplicationCampaignAnalytics) SetTotalRevenue(v AnalyticsDataPointWithTrendAndInfluencedRate)`

SetTotalRevenue gets a reference to the given AnalyticsDataPointWithTrendAndInfluencedRate and assigns it to the TotalRevenue field.

### GetSessionsCount

`func (o *ApplicationCampaignAnalytics) GetSessionsCount() AnalyticsDataPointWithTrendAndInfluencedRate`

GetSessionsCount returns the SessionsCount field if non-nil, zero value otherwise.

### GetSessionsCountOk

`func (o *ApplicationCampaignAnalytics) GetSessionsCountOk() (AnalyticsDataPointWithTrendAndInfluencedRate, bool)`

GetSessionsCountOk returns a tuple with the SessionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSessionsCount

`func (o *ApplicationCampaignAnalytics) HasSessionsCount() bool`

HasSessionsCount returns a boolean if a field has been set.

### SetSessionsCount

`func (o *ApplicationCampaignAnalytics) SetSessionsCount(v AnalyticsDataPointWithTrendAndInfluencedRate)`

SetSessionsCount gets a reference to the given AnalyticsDataPointWithTrendAndInfluencedRate and assigns it to the SessionsCount field.

### GetAvgItemsPerSession

`func (o *ApplicationCampaignAnalytics) GetAvgItemsPerSession() AnalyticsDataPointWithTrendAndUplift`

GetAvgItemsPerSession returns the AvgItemsPerSession field if non-nil, zero value otherwise.

### GetAvgItemsPerSessionOk

`func (o *ApplicationCampaignAnalytics) GetAvgItemsPerSessionOk() (AnalyticsDataPointWithTrendAndUplift, bool)`

GetAvgItemsPerSessionOk returns a tuple with the AvgItemsPerSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAvgItemsPerSession

`func (o *ApplicationCampaignAnalytics) HasAvgItemsPerSession() bool`

HasAvgItemsPerSession returns a boolean if a field has been set.

### SetAvgItemsPerSession

`func (o *ApplicationCampaignAnalytics) SetAvgItemsPerSession(v AnalyticsDataPointWithTrendAndUplift)`

SetAvgItemsPerSession gets a reference to the given AnalyticsDataPointWithTrendAndUplift and assigns it to the AvgItemsPerSession field.

### GetAvgSessionValue

`func (o *ApplicationCampaignAnalytics) GetAvgSessionValue() AnalyticsDataPointWithTrendAndUplift`

GetAvgSessionValue returns the AvgSessionValue field if non-nil, zero value otherwise.

### GetAvgSessionValueOk

`func (o *ApplicationCampaignAnalytics) GetAvgSessionValueOk() (AnalyticsDataPointWithTrendAndUplift, bool)`

GetAvgSessionValueOk returns a tuple with the AvgSessionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAvgSessionValue

`func (o *ApplicationCampaignAnalytics) HasAvgSessionValue() bool`

HasAvgSessionValue returns a boolean if a field has been set.

### SetAvgSessionValue

`func (o *ApplicationCampaignAnalytics) SetAvgSessionValue(v AnalyticsDataPointWithTrendAndUplift)`

SetAvgSessionValue gets a reference to the given AnalyticsDataPointWithTrendAndUplift and assigns it to the AvgSessionValue field.

### GetTotalDiscounts

`func (o *ApplicationCampaignAnalytics) GetTotalDiscounts() AnalyticsDataPointWithTrend`

GetTotalDiscounts returns the TotalDiscounts field if non-nil, zero value otherwise.

### GetTotalDiscountsOk

`func (o *ApplicationCampaignAnalytics) GetTotalDiscountsOk() (AnalyticsDataPointWithTrend, bool)`

GetTotalDiscountsOk returns a tuple with the TotalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalDiscounts

`func (o *ApplicationCampaignAnalytics) HasTotalDiscounts() bool`

HasTotalDiscounts returns a boolean if a field has been set.

### SetTotalDiscounts

`func (o *ApplicationCampaignAnalytics) SetTotalDiscounts(v AnalyticsDataPointWithTrend)`

SetTotalDiscounts gets a reference to the given AnalyticsDataPointWithTrend and assigns it to the TotalDiscounts field.

### GetCouponsCount

`func (o *ApplicationCampaignAnalytics) GetCouponsCount() AnalyticsDataPointWithTrend`

GetCouponsCount returns the CouponsCount field if non-nil, zero value otherwise.

### GetCouponsCountOk

`func (o *ApplicationCampaignAnalytics) GetCouponsCountOk() (AnalyticsDataPointWithTrend, bool)`

GetCouponsCountOk returns a tuple with the CouponsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponsCount

`func (o *ApplicationCampaignAnalytics) HasCouponsCount() bool`

HasCouponsCount returns a boolean if a field has been set.

### SetCouponsCount

`func (o *ApplicationCampaignAnalytics) SetCouponsCount(v AnalyticsDataPointWithTrend)`

SetCouponsCount gets a reference to the given AnalyticsDataPointWithTrend and assigns it to the CouponsCount field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


