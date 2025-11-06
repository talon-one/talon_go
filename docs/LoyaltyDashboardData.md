# LoyaltyDashboardData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to [**time.Time**](time.Time.md) | Date at which data point was collected. | 
**TotalActivePoints** | Pointer to **float32** | Total of active points for this loyalty program. | 
**TotalPendingPoints** | Pointer to **float32** | Total of pending points for this loyalty program. | 
**TotalSpentPoints** | Pointer to **float32** | Total of spent points for this loyalty program. | 
**TotalExpiredPoints** | Pointer to **float32** | Total of expired points for this loyalty program. | 
**TotalNegativePoints** | Pointer to **float32** | Total of negative points for this loyalty program. | 
**TotalMembers** | Pointer to **float32** | Number of loyalty program members. | 
**NewMembers** | Pointer to **float32** | Number of members who joined on this day. | 
**SpentPoints** | Pointer to [**LoyaltyDashboardPointsBreakdown**](LoyaltyDashboardPointsBreakdown.md) |  | 
**EarnedPoints** | Pointer to [**LoyaltyDashboardPointsBreakdown**](LoyaltyDashboardPointsBreakdown.md) |  | 

## Methods

### NewLoyaltyDashboardData

`func NewLoyaltyDashboardData(date time.Time, totalActivePoints float32, totalPendingPoints float32, totalSpentPoints float32, totalExpiredPoints float32, totalNegativePoints float32, totalMembers float32, newMembers float32, spentPoints LoyaltyDashboardPointsBreakdown, earnedPoints LoyaltyDashboardPointsBreakdown, ) *LoyaltyDashboardData`

NewLoyaltyDashboardData instantiates a new LoyaltyDashboardData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyDashboardDataWithDefaults

`func NewLoyaltyDashboardDataWithDefaults() *LoyaltyDashboardData`

NewLoyaltyDashboardDataWithDefaults instantiates a new LoyaltyDashboardData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *LoyaltyDashboardData) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *LoyaltyDashboardData) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *LoyaltyDashboardData) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetTotalActivePoints

`func (o *LoyaltyDashboardData) GetTotalActivePoints() float32`

GetTotalActivePoints returns the TotalActivePoints field if non-nil, zero value otherwise.

### GetTotalActivePointsOk

`func (o *LoyaltyDashboardData) GetTotalActivePointsOk() (*float32, bool)`

GetTotalActivePointsOk returns a tuple with the TotalActivePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalActivePoints

`func (o *LoyaltyDashboardData) SetTotalActivePoints(v float32)`

SetTotalActivePoints sets TotalActivePoints field to given value.


### GetTotalPendingPoints

`func (o *LoyaltyDashboardData) GetTotalPendingPoints() float32`

GetTotalPendingPoints returns the TotalPendingPoints field if non-nil, zero value otherwise.

### GetTotalPendingPointsOk

`func (o *LoyaltyDashboardData) GetTotalPendingPointsOk() (*float32, bool)`

GetTotalPendingPointsOk returns a tuple with the TotalPendingPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPendingPoints

`func (o *LoyaltyDashboardData) SetTotalPendingPoints(v float32)`

SetTotalPendingPoints sets TotalPendingPoints field to given value.


### GetTotalSpentPoints

`func (o *LoyaltyDashboardData) GetTotalSpentPoints() float32`

GetTotalSpentPoints returns the TotalSpentPoints field if non-nil, zero value otherwise.

### GetTotalSpentPointsOk

`func (o *LoyaltyDashboardData) GetTotalSpentPointsOk() (*float32, bool)`

GetTotalSpentPointsOk returns a tuple with the TotalSpentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpentPoints

`func (o *LoyaltyDashboardData) SetTotalSpentPoints(v float32)`

SetTotalSpentPoints sets TotalSpentPoints field to given value.


### GetTotalExpiredPoints

`func (o *LoyaltyDashboardData) GetTotalExpiredPoints() float32`

GetTotalExpiredPoints returns the TotalExpiredPoints field if non-nil, zero value otherwise.

### GetTotalExpiredPointsOk

`func (o *LoyaltyDashboardData) GetTotalExpiredPointsOk() (*float32, bool)`

GetTotalExpiredPointsOk returns a tuple with the TotalExpiredPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExpiredPoints

`func (o *LoyaltyDashboardData) SetTotalExpiredPoints(v float32)`

SetTotalExpiredPoints sets TotalExpiredPoints field to given value.


### GetTotalNegativePoints

`func (o *LoyaltyDashboardData) GetTotalNegativePoints() float32`

GetTotalNegativePoints returns the TotalNegativePoints field if non-nil, zero value otherwise.

### GetTotalNegativePointsOk

`func (o *LoyaltyDashboardData) GetTotalNegativePointsOk() (*float32, bool)`

GetTotalNegativePointsOk returns a tuple with the TotalNegativePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNegativePoints

`func (o *LoyaltyDashboardData) SetTotalNegativePoints(v float32)`

SetTotalNegativePoints sets TotalNegativePoints field to given value.


### GetTotalMembers

`func (o *LoyaltyDashboardData) GetTotalMembers() float32`

GetTotalMembers returns the TotalMembers field if non-nil, zero value otherwise.

### GetTotalMembersOk

`func (o *LoyaltyDashboardData) GetTotalMembersOk() (*float32, bool)`

GetTotalMembersOk returns a tuple with the TotalMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMembers

`func (o *LoyaltyDashboardData) SetTotalMembers(v float32)`

SetTotalMembers sets TotalMembers field to given value.


### GetNewMembers

`func (o *LoyaltyDashboardData) GetNewMembers() float32`

GetNewMembers returns the NewMembers field if non-nil, zero value otherwise.

### GetNewMembersOk

`func (o *LoyaltyDashboardData) GetNewMembersOk() (*float32, bool)`

GetNewMembersOk returns a tuple with the NewMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMembers

`func (o *LoyaltyDashboardData) SetNewMembers(v float32)`

SetNewMembers sets NewMembers field to given value.


### GetSpentPoints

`func (o *LoyaltyDashboardData) GetSpentPoints() LoyaltyDashboardPointsBreakdown`

GetSpentPoints returns the SpentPoints field if non-nil, zero value otherwise.

### GetSpentPointsOk

`func (o *LoyaltyDashboardData) GetSpentPointsOk() (*LoyaltyDashboardPointsBreakdown, bool)`

GetSpentPointsOk returns a tuple with the SpentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentPoints

`func (o *LoyaltyDashboardData) SetSpentPoints(v LoyaltyDashboardPointsBreakdown)`

SetSpentPoints sets SpentPoints field to given value.


### GetEarnedPoints

`func (o *LoyaltyDashboardData) GetEarnedPoints() LoyaltyDashboardPointsBreakdown`

GetEarnedPoints returns the EarnedPoints field if non-nil, zero value otherwise.

### GetEarnedPointsOk

`func (o *LoyaltyDashboardData) GetEarnedPointsOk() (*LoyaltyDashboardPointsBreakdown, bool)`

GetEarnedPointsOk returns a tuple with the EarnedPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarnedPoints

`func (o *LoyaltyDashboardData) SetEarnedPoints(v LoyaltyDashboardPointsBreakdown)`

SetEarnedPoints sets EarnedPoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


