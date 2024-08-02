# LoyaltyBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePoints** | Pointer to **float32** | Total amount of points awarded to this customer and available to spend. | [optional] 
**ExpiredPoints** | Pointer to **float32** | Total amount of points awarded but never redeemed. They cannot be used anymore. | [optional] 
**PendingPoints** | Pointer to **float32** | Total amount of points awarded to this customer but not available until their start date. | [optional] 
**SpentPoints** | Pointer to **float32** | Total amount of points already spent by this customer. | [optional] 

## Methods

### GetActivePoints

`func (o *LoyaltyBalance) GetActivePoints() float32`

GetActivePoints returns the ActivePoints field if non-nil, zero value otherwise.

### GetActivePointsOk

`func (o *LoyaltyBalance) GetActivePointsOk() (float32, bool)`

GetActivePointsOk returns a tuple with the ActivePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivePoints

`func (o *LoyaltyBalance) HasActivePoints() bool`

HasActivePoints returns a boolean if a field has been set.

### SetActivePoints

`func (o *LoyaltyBalance) SetActivePoints(v float32)`

SetActivePoints gets a reference to the given float32 and assigns it to the ActivePoints field.

### GetExpiredPoints

`func (o *LoyaltyBalance) GetExpiredPoints() float32`

GetExpiredPoints returns the ExpiredPoints field if non-nil, zero value otherwise.

### GetExpiredPointsOk

`func (o *LoyaltyBalance) GetExpiredPointsOk() (float32, bool)`

GetExpiredPointsOk returns a tuple with the ExpiredPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiredPoints

`func (o *LoyaltyBalance) HasExpiredPoints() bool`

HasExpiredPoints returns a boolean if a field has been set.

### SetExpiredPoints

`func (o *LoyaltyBalance) SetExpiredPoints(v float32)`

SetExpiredPoints gets a reference to the given float32 and assigns it to the ExpiredPoints field.

### GetPendingPoints

`func (o *LoyaltyBalance) GetPendingPoints() float32`

GetPendingPoints returns the PendingPoints field if non-nil, zero value otherwise.

### GetPendingPointsOk

`func (o *LoyaltyBalance) GetPendingPointsOk() (float32, bool)`

GetPendingPointsOk returns a tuple with the PendingPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPendingPoints

`func (o *LoyaltyBalance) HasPendingPoints() bool`

HasPendingPoints returns a boolean if a field has been set.

### SetPendingPoints

`func (o *LoyaltyBalance) SetPendingPoints(v float32)`

SetPendingPoints gets a reference to the given float32 and assigns it to the PendingPoints field.

### GetSpentPoints

`func (o *LoyaltyBalance) GetSpentPoints() float32`

GetSpentPoints returns the SpentPoints field if non-nil, zero value otherwise.

### GetSpentPointsOk

`func (o *LoyaltyBalance) GetSpentPointsOk() (float32, bool)`

GetSpentPointsOk returns a tuple with the SpentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpentPoints

`func (o *LoyaltyBalance) HasSpentPoints() bool`

HasSpentPoints returns a boolean if a field has been set.

### SetSpentPoints

`func (o *LoyaltyBalance) SetSpentPoints(v float32)`

SetSpentPoints gets a reference to the given float32 and assigns it to the SpentPoints field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


