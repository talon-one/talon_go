# LoyaltyDashboardData

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

`func (o *LoyaltyDashboardData) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *LoyaltyDashboardData) GetDateOk() (time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDate

`func (o *LoyaltyDashboardData) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDate

`func (o *LoyaltyDashboardData) SetDate(v time.Time)`

SetDate gets a reference to the given time.Time and assigns it to the Date field.

### GetTotalActivePoints

`func (o *LoyaltyDashboardData) GetTotalActivePoints() float32`

GetTotalActivePoints returns the TotalActivePoints field if non-nil, zero value otherwise.

### GetTotalActivePointsOk

`func (o *LoyaltyDashboardData) GetTotalActivePointsOk() (float32, bool)`

GetTotalActivePointsOk returns a tuple with the TotalActivePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalActivePoints

`func (o *LoyaltyDashboardData) HasTotalActivePoints() bool`

HasTotalActivePoints returns a boolean if a field has been set.

### SetTotalActivePoints

`func (o *LoyaltyDashboardData) SetTotalActivePoints(v float32)`

SetTotalActivePoints gets a reference to the given float32 and assigns it to the TotalActivePoints field.

### GetTotalPendingPoints

`func (o *LoyaltyDashboardData) GetTotalPendingPoints() float32`

GetTotalPendingPoints returns the TotalPendingPoints field if non-nil, zero value otherwise.

### GetTotalPendingPointsOk

`func (o *LoyaltyDashboardData) GetTotalPendingPointsOk() (float32, bool)`

GetTotalPendingPointsOk returns a tuple with the TotalPendingPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalPendingPoints

`func (o *LoyaltyDashboardData) HasTotalPendingPoints() bool`

HasTotalPendingPoints returns a boolean if a field has been set.

### SetTotalPendingPoints

`func (o *LoyaltyDashboardData) SetTotalPendingPoints(v float32)`

SetTotalPendingPoints gets a reference to the given float32 and assigns it to the TotalPendingPoints field.

### GetTotalSpentPoints

`func (o *LoyaltyDashboardData) GetTotalSpentPoints() float32`

GetTotalSpentPoints returns the TotalSpentPoints field if non-nil, zero value otherwise.

### GetTotalSpentPointsOk

`func (o *LoyaltyDashboardData) GetTotalSpentPointsOk() (float32, bool)`

GetTotalSpentPointsOk returns a tuple with the TotalSpentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalSpentPoints

`func (o *LoyaltyDashboardData) HasTotalSpentPoints() bool`

HasTotalSpentPoints returns a boolean if a field has been set.

### SetTotalSpentPoints

`func (o *LoyaltyDashboardData) SetTotalSpentPoints(v float32)`

SetTotalSpentPoints gets a reference to the given float32 and assigns it to the TotalSpentPoints field.

### GetTotalExpiredPoints

`func (o *LoyaltyDashboardData) GetTotalExpiredPoints() float32`

GetTotalExpiredPoints returns the TotalExpiredPoints field if non-nil, zero value otherwise.

### GetTotalExpiredPointsOk

`func (o *LoyaltyDashboardData) GetTotalExpiredPointsOk() (float32, bool)`

GetTotalExpiredPointsOk returns a tuple with the TotalExpiredPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalExpiredPoints

`func (o *LoyaltyDashboardData) HasTotalExpiredPoints() bool`

HasTotalExpiredPoints returns a boolean if a field has been set.

### SetTotalExpiredPoints

`func (o *LoyaltyDashboardData) SetTotalExpiredPoints(v float32)`

SetTotalExpiredPoints gets a reference to the given float32 and assigns it to the TotalExpiredPoints field.

### GetTotalMembers

`func (o *LoyaltyDashboardData) GetTotalMembers() float32`

GetTotalMembers returns the TotalMembers field if non-nil, zero value otherwise.

### GetTotalMembersOk

`func (o *LoyaltyDashboardData) GetTotalMembersOk() (float32, bool)`

GetTotalMembersOk returns a tuple with the TotalMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalMembers

`func (o *LoyaltyDashboardData) HasTotalMembers() bool`

HasTotalMembers returns a boolean if a field has been set.

### SetTotalMembers

`func (o *LoyaltyDashboardData) SetTotalMembers(v float32)`

SetTotalMembers gets a reference to the given float32 and assigns it to the TotalMembers field.

### GetNewMembers

`func (o *LoyaltyDashboardData) GetNewMembers() float32`

GetNewMembers returns the NewMembers field if non-nil, zero value otherwise.

### GetNewMembersOk

`func (o *LoyaltyDashboardData) GetNewMembersOk() (float32, bool)`

GetNewMembersOk returns a tuple with the NewMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNewMembers

`func (o *LoyaltyDashboardData) HasNewMembers() bool`

HasNewMembers returns a boolean if a field has been set.

### SetNewMembers

`func (o *LoyaltyDashboardData) SetNewMembers(v float32)`

SetNewMembers gets a reference to the given float32 and assigns it to the NewMembers field.

### GetSpentPoints

`func (o *LoyaltyDashboardData) GetSpentPoints() LoyaltyDashboardPointsBreakdown`

GetSpentPoints returns the SpentPoints field if non-nil, zero value otherwise.

### GetSpentPointsOk

`func (o *LoyaltyDashboardData) GetSpentPointsOk() (LoyaltyDashboardPointsBreakdown, bool)`

GetSpentPointsOk returns a tuple with the SpentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpentPoints

`func (o *LoyaltyDashboardData) HasSpentPoints() bool`

HasSpentPoints returns a boolean if a field has been set.

### SetSpentPoints

`func (o *LoyaltyDashboardData) SetSpentPoints(v LoyaltyDashboardPointsBreakdown)`

SetSpentPoints gets a reference to the given LoyaltyDashboardPointsBreakdown and assigns it to the SpentPoints field.

### GetEarnedPoints

`func (o *LoyaltyDashboardData) GetEarnedPoints() LoyaltyDashboardPointsBreakdown`

GetEarnedPoints returns the EarnedPoints field if non-nil, zero value otherwise.

### GetEarnedPointsOk

`func (o *LoyaltyDashboardData) GetEarnedPointsOk() (LoyaltyDashboardPointsBreakdown, bool)`

GetEarnedPointsOk returns a tuple with the EarnedPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEarnedPoints

`func (o *LoyaltyDashboardData) HasEarnedPoints() bool`

HasEarnedPoints returns a boolean if a field has been set.

### SetEarnedPoints

`func (o *LoyaltyDashboardData) SetEarnedPoints(v LoyaltyDashboardPointsBreakdown)`

SetEarnedPoints gets a reference to the given LoyaltyDashboardPointsBreakdown and assigns it to the EarnedPoints field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


