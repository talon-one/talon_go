# ApplicationAnalyticsDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The start of the aggregation time frame in UTC. | 
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The end of the aggregation time frame in UTC. | 
**TotalRevenue** | Pointer to [**AnalyticsDataPoint**](AnalyticsDataPoint.md) |  | [optional] 
**SessionsCount** | Pointer to [**AnalyticsDataPoint**](AnalyticsDataPoint.md) |  | [optional] 
**AvgItemsPerSession** | Pointer to [**AnalyticsDataPoint**](AnalyticsDataPoint.md) |  | [optional] 
**AvgSessionValue** | Pointer to [**AnalyticsDataPoint**](AnalyticsDataPoint.md) |  | [optional] 
**TotalDiscounts** | Pointer to **float32** | The total value of discounts given for cart items in influenced sessions. | [optional] 
**CouponsCount** | Pointer to **float32** | The number of times a coupon was successfully redeemed in influenced sessions. | [optional] 

## Methods

### GetStartTime

`func (o *ApplicationAnalyticsDataPoint) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ApplicationAnalyticsDataPoint) GetStartTimeOk() (time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartTime

`func (o *ApplicationAnalyticsDataPoint) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTime

`func (o *ApplicationAnalyticsDataPoint) SetStartTime(v time.Time)`

SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.

### GetEndTime

`func (o *ApplicationAnalyticsDataPoint) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ApplicationAnalyticsDataPoint) GetEndTimeOk() (time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndTime

`func (o *ApplicationAnalyticsDataPoint) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTime

`func (o *ApplicationAnalyticsDataPoint) SetEndTime(v time.Time)`

SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.

### GetTotalRevenue

`func (o *ApplicationAnalyticsDataPoint) GetTotalRevenue() AnalyticsDataPoint`

GetTotalRevenue returns the TotalRevenue field if non-nil, zero value otherwise.

### GetTotalRevenueOk

`func (o *ApplicationAnalyticsDataPoint) GetTotalRevenueOk() (AnalyticsDataPoint, bool)`

GetTotalRevenueOk returns a tuple with the TotalRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalRevenue

`func (o *ApplicationAnalyticsDataPoint) HasTotalRevenue() bool`

HasTotalRevenue returns a boolean if a field has been set.

### SetTotalRevenue

`func (o *ApplicationAnalyticsDataPoint) SetTotalRevenue(v AnalyticsDataPoint)`

SetTotalRevenue gets a reference to the given AnalyticsDataPoint and assigns it to the TotalRevenue field.

### GetSessionsCount

`func (o *ApplicationAnalyticsDataPoint) GetSessionsCount() AnalyticsDataPoint`

GetSessionsCount returns the SessionsCount field if non-nil, zero value otherwise.

### GetSessionsCountOk

`func (o *ApplicationAnalyticsDataPoint) GetSessionsCountOk() (AnalyticsDataPoint, bool)`

GetSessionsCountOk returns a tuple with the SessionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSessionsCount

`func (o *ApplicationAnalyticsDataPoint) HasSessionsCount() bool`

HasSessionsCount returns a boolean if a field has been set.

### SetSessionsCount

`func (o *ApplicationAnalyticsDataPoint) SetSessionsCount(v AnalyticsDataPoint)`

SetSessionsCount gets a reference to the given AnalyticsDataPoint and assigns it to the SessionsCount field.

### GetAvgItemsPerSession

`func (o *ApplicationAnalyticsDataPoint) GetAvgItemsPerSession() AnalyticsDataPoint`

GetAvgItemsPerSession returns the AvgItemsPerSession field if non-nil, zero value otherwise.

### GetAvgItemsPerSessionOk

`func (o *ApplicationAnalyticsDataPoint) GetAvgItemsPerSessionOk() (AnalyticsDataPoint, bool)`

GetAvgItemsPerSessionOk returns a tuple with the AvgItemsPerSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAvgItemsPerSession

`func (o *ApplicationAnalyticsDataPoint) HasAvgItemsPerSession() bool`

HasAvgItemsPerSession returns a boolean if a field has been set.

### SetAvgItemsPerSession

`func (o *ApplicationAnalyticsDataPoint) SetAvgItemsPerSession(v AnalyticsDataPoint)`

SetAvgItemsPerSession gets a reference to the given AnalyticsDataPoint and assigns it to the AvgItemsPerSession field.

### GetAvgSessionValue

`func (o *ApplicationAnalyticsDataPoint) GetAvgSessionValue() AnalyticsDataPoint`

GetAvgSessionValue returns the AvgSessionValue field if non-nil, zero value otherwise.

### GetAvgSessionValueOk

`func (o *ApplicationAnalyticsDataPoint) GetAvgSessionValueOk() (AnalyticsDataPoint, bool)`

GetAvgSessionValueOk returns a tuple with the AvgSessionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAvgSessionValue

`func (o *ApplicationAnalyticsDataPoint) HasAvgSessionValue() bool`

HasAvgSessionValue returns a boolean if a field has been set.

### SetAvgSessionValue

`func (o *ApplicationAnalyticsDataPoint) SetAvgSessionValue(v AnalyticsDataPoint)`

SetAvgSessionValue gets a reference to the given AnalyticsDataPoint and assigns it to the AvgSessionValue field.

### GetTotalDiscounts

`func (o *ApplicationAnalyticsDataPoint) GetTotalDiscounts() float32`

GetTotalDiscounts returns the TotalDiscounts field if non-nil, zero value otherwise.

### GetTotalDiscountsOk

`func (o *ApplicationAnalyticsDataPoint) GetTotalDiscountsOk() (float32, bool)`

GetTotalDiscountsOk returns a tuple with the TotalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalDiscounts

`func (o *ApplicationAnalyticsDataPoint) HasTotalDiscounts() bool`

HasTotalDiscounts returns a boolean if a field has been set.

### SetTotalDiscounts

`func (o *ApplicationAnalyticsDataPoint) SetTotalDiscounts(v float32)`

SetTotalDiscounts gets a reference to the given float32 and assigns it to the TotalDiscounts field.

### GetCouponsCount

`func (o *ApplicationAnalyticsDataPoint) GetCouponsCount() float32`

GetCouponsCount returns the CouponsCount field if non-nil, zero value otherwise.

### GetCouponsCountOk

`func (o *ApplicationAnalyticsDataPoint) GetCouponsCountOk() (float32, bool)`

GetCouponsCountOk returns a tuple with the CouponsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponsCount

`func (o *ApplicationAnalyticsDataPoint) HasCouponsCount() bool`

HasCouponsCount returns a boolean if a field has been set.

### SetCouponsCount

`func (o *ApplicationAnalyticsDataPoint) SetCouponsCount(v float32)`

SetCouponsCount gets a reference to the given float32 and assigns it to the CouponsCount field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


