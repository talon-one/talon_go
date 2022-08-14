# LoyaltyProjectionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to [**time.Time**](time.Time.md) | Specific date of projection. | 
**ExpiringPoints** | Pointer to **float32** | Points that will be expired by the specified date. | 
**ActivatingPoints** | Pointer to **float32** | Points that will be active by the specified date. | 
**ProjectedBalance** | Pointer to **float32** | Current balance plus projected active points, minus expiring points. | 

## Methods

### GetDate

`func (o *LoyaltyProjectionData) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *LoyaltyProjectionData) GetDateOk() (time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDate

`func (o *LoyaltyProjectionData) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDate

`func (o *LoyaltyProjectionData) SetDate(v time.Time)`

SetDate gets a reference to the given time.Time and assigns it to the Date field.

### GetExpiringPoints

`func (o *LoyaltyProjectionData) GetExpiringPoints() float32`

GetExpiringPoints returns the ExpiringPoints field if non-nil, zero value otherwise.

### GetExpiringPointsOk

`func (o *LoyaltyProjectionData) GetExpiringPointsOk() (float32, bool)`

GetExpiringPointsOk returns a tuple with the ExpiringPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiringPoints

`func (o *LoyaltyProjectionData) HasExpiringPoints() bool`

HasExpiringPoints returns a boolean if a field has been set.

### SetExpiringPoints

`func (o *LoyaltyProjectionData) SetExpiringPoints(v float32)`

SetExpiringPoints gets a reference to the given float32 and assigns it to the ExpiringPoints field.

### GetActivatingPoints

`func (o *LoyaltyProjectionData) GetActivatingPoints() float32`

GetActivatingPoints returns the ActivatingPoints field if non-nil, zero value otherwise.

### GetActivatingPointsOk

`func (o *LoyaltyProjectionData) GetActivatingPointsOk() (float32, bool)`

GetActivatingPointsOk returns a tuple with the ActivatingPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivatingPoints

`func (o *LoyaltyProjectionData) HasActivatingPoints() bool`

HasActivatingPoints returns a boolean if a field has been set.

### SetActivatingPoints

`func (o *LoyaltyProjectionData) SetActivatingPoints(v float32)`

SetActivatingPoints gets a reference to the given float32 and assigns it to the ActivatingPoints field.

### GetProjectedBalance

`func (o *LoyaltyProjectionData) GetProjectedBalance() float32`

GetProjectedBalance returns the ProjectedBalance field if non-nil, zero value otherwise.

### GetProjectedBalanceOk

`func (o *LoyaltyProjectionData) GetProjectedBalanceOk() (float32, bool)`

GetProjectedBalanceOk returns a tuple with the ProjectedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProjectedBalance

`func (o *LoyaltyProjectionData) HasProjectedBalance() bool`

HasProjectedBalance returns a boolean if a field has been set.

### SetProjectedBalance

`func (o *LoyaltyProjectionData) SetProjectedBalance(v float32)`

SetProjectedBalance gets a reference to the given float32 and assigns it to the ProjectedBalance field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


