# LoyaltyStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to [**time.Time**](time.Time.md) | Date at which data point was collected. | 
**TotalActivePoints** | Pointer to **float32** | Total of active points for this loyalty program. | 
**TotalPendingPoints** | Pointer to **float32** | Total of pending points for this loyalty program. | 
**TotalSpentPoints** | Pointer to **float32** | Total of spent points for this loyalty program. | 
**TotalExpiredPoints** | Pointer to **float32** | Total of expired points for this loyalty program. | 
**TotalMembers** | Pointer to **float32** | Number of loyalty program members. | 
**NewMembers** | Pointer to **float32** | Number of members who joined on this day. | 
**SpentPoints** | Pointer to [**LoyaltyDashboardPointsBreakdown**](LoyaltyDashboardPointsBreakdown.md) |  | 
**EarnedPoints** | Pointer to [**LoyaltyDashboardPointsBreakdown**](LoyaltyDashboardPointsBreakdown.md) |  | 

## Methods

### GetDate

`func (o *LoyaltyStatistics) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *LoyaltyStatistics) GetDateOk() (time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDate

`func (o *LoyaltyStatistics) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDate

`func (o *LoyaltyStatistics) SetDate(v time.Time)`

SetDate gets a reference to the given time.Time and assigns it to the Date field.

### GetTotalActivePoints

`func (o *LoyaltyStatistics) GetTotalActivePoints() float32`

GetTotalActivePoints returns the TotalActivePoints field if non-nil, zero value otherwise.

### GetTotalActivePointsOk

`func (o *LoyaltyStatistics) GetTotalActivePointsOk() (float32, bool)`

GetTotalActivePointsOk returns a tuple with the TotalActivePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalActivePoints

`func (o *LoyaltyStatistics) HasTotalActivePoints() bool`

HasTotalActivePoints returns a boolean if a field has been set.

### SetTotalActivePoints

`func (o *LoyaltyStatistics) SetTotalActivePoints(v float32)`

SetTotalActivePoints gets a reference to the given float32 and assigns it to the TotalActivePoints field.

### GetTotalPendingPoints

`func (o *LoyaltyStatistics) GetTotalPendingPoints() float32`

GetTotalPendingPoints returns the TotalPendingPoints field if non-nil, zero value otherwise.

### GetTotalPendingPointsOk

`func (o *LoyaltyStatistics) GetTotalPendingPointsOk() (float32, bool)`

GetTotalPendingPointsOk returns a tuple with the TotalPendingPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalPendingPoints

`func (o *LoyaltyStatistics) HasTotalPendingPoints() bool`

HasTotalPendingPoints returns a boolean if a field has been set.

### SetTotalPendingPoints

`func (o *LoyaltyStatistics) SetTotalPendingPoints(v float32)`

SetTotalPendingPoints gets a reference to the given float32 and assigns it to the TotalPendingPoints field.

### GetTotalSpentPoints

`func (o *LoyaltyStatistics) GetTotalSpentPoints() float32`

GetTotalSpentPoints returns the TotalSpentPoints field if non-nil, zero value otherwise.

### GetTotalSpentPointsOk

`func (o *LoyaltyStatistics) GetTotalSpentPointsOk() (float32, bool)`

GetTotalSpentPointsOk returns a tuple with the TotalSpentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalSpentPoints

`func (o *LoyaltyStatistics) HasTotalSpentPoints() bool`

HasTotalSpentPoints returns a boolean if a field has been set.

### SetTotalSpentPoints

`func (o *LoyaltyStatistics) SetTotalSpentPoints(v float32)`

SetTotalSpentPoints gets a reference to the given float32 and assigns it to the TotalSpentPoints field.

### GetTotalExpiredPoints

`func (o *LoyaltyStatistics) GetTotalExpiredPoints() float32`

GetTotalExpiredPoints returns the TotalExpiredPoints field if non-nil, zero value otherwise.

### GetTotalExpiredPointsOk

`func (o *LoyaltyStatistics) GetTotalExpiredPointsOk() (float32, bool)`

GetTotalExpiredPointsOk returns a tuple with the TotalExpiredPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalExpiredPoints

`func (o *LoyaltyStatistics) HasTotalExpiredPoints() bool`

HasTotalExpiredPoints returns a boolean if a field has been set.

### SetTotalExpiredPoints

`func (o *LoyaltyStatistics) SetTotalExpiredPoints(v float32)`

SetTotalExpiredPoints gets a reference to the given float32 and assigns it to the TotalExpiredPoints field.

### GetTotalMembers

`func (o *LoyaltyStatistics) GetTotalMembers() float32`

GetTotalMembers returns the TotalMembers field if non-nil, zero value otherwise.

### GetTotalMembersOk

`func (o *LoyaltyStatistics) GetTotalMembersOk() (float32, bool)`

GetTotalMembersOk returns a tuple with the TotalMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalMembers

`func (o *LoyaltyStatistics) HasTotalMembers() bool`

HasTotalMembers returns a boolean if a field has been set.

### SetTotalMembers

`func (o *LoyaltyStatistics) SetTotalMembers(v float32)`

SetTotalMembers gets a reference to the given float32 and assigns it to the TotalMembers field.

### GetNewMembers

`func (o *LoyaltyStatistics) GetNewMembers() float32`

GetNewMembers returns the NewMembers field if non-nil, zero value otherwise.

### GetNewMembersOk

`func (o *LoyaltyStatistics) GetNewMembersOk() (float32, bool)`

GetNewMembersOk returns a tuple with the NewMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNewMembers

`func (o *LoyaltyStatistics) HasNewMembers() bool`

HasNewMembers returns a boolean if a field has been set.

### SetNewMembers

`func (o *LoyaltyStatistics) SetNewMembers(v float32)`

SetNewMembers gets a reference to the given float32 and assigns it to the NewMembers field.

### GetSpentPoints

`func (o *LoyaltyStatistics) GetSpentPoints() LoyaltyDashboardPointsBreakdown`

GetSpentPoints returns the SpentPoints field if non-nil, zero value otherwise.

### GetSpentPointsOk

`func (o *LoyaltyStatistics) GetSpentPointsOk() (LoyaltyDashboardPointsBreakdown, bool)`

GetSpentPointsOk returns a tuple with the SpentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpentPoints

`func (o *LoyaltyStatistics) HasSpentPoints() bool`

HasSpentPoints returns a boolean if a field has been set.

### SetSpentPoints

`func (o *LoyaltyStatistics) SetSpentPoints(v LoyaltyDashboardPointsBreakdown)`

SetSpentPoints gets a reference to the given LoyaltyDashboardPointsBreakdown and assigns it to the SpentPoints field.

### GetEarnedPoints

`func (o *LoyaltyStatistics) GetEarnedPoints() LoyaltyDashboardPointsBreakdown`

GetEarnedPoints returns the EarnedPoints field if non-nil, zero value otherwise.

### GetEarnedPointsOk

`func (o *LoyaltyStatistics) GetEarnedPointsOk() (LoyaltyDashboardPointsBreakdown, bool)`

GetEarnedPointsOk returns a tuple with the EarnedPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEarnedPoints

`func (o *LoyaltyStatistics) HasEarnedPoints() bool`

HasEarnedPoints returns a boolean if a field has been set.

### SetEarnedPoints

`func (o *LoyaltyStatistics) SetEarnedPoints(v LoyaltyDashboardPointsBreakdown)`

SetEarnedPoints gets a reference to the given LoyaltyDashboardPointsBreakdown and assigns it to the EarnedPoints field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


