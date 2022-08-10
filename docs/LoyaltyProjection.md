# LoyaltyProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projections** | Pointer to [**[]LoyaltyProjectionData**](LoyaltyProjectionData.md) |  | [optional] 
**TotalExpiringPoints** | Pointer to **float32** | Sum of points to be expired by the projection date set in the query parameter. | 
**TotalActivatingPoints** | Pointer to **float32** | Sum of points to be active by the projection date set in the query parameter. | 

## Methods

### GetProjections

`func (o *LoyaltyProjection) GetProjections() []LoyaltyProjectionData`

GetProjections returns the Projections field if non-nil, zero value otherwise.

### GetProjectionsOk

`func (o *LoyaltyProjection) GetProjectionsOk() ([]LoyaltyProjectionData, bool)`

GetProjectionsOk returns a tuple with the Projections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProjections

`func (o *LoyaltyProjection) HasProjections() bool`

HasProjections returns a boolean if a field has been set.

### SetProjections

`func (o *LoyaltyProjection) SetProjections(v []LoyaltyProjectionData)`

SetProjections gets a reference to the given []LoyaltyProjectionData and assigns it to the Projections field.

### GetTotalExpiringPoints

`func (o *LoyaltyProjection) GetTotalExpiringPoints() float32`

GetTotalExpiringPoints returns the TotalExpiringPoints field if non-nil, zero value otherwise.

### GetTotalExpiringPointsOk

`func (o *LoyaltyProjection) GetTotalExpiringPointsOk() (float32, bool)`

GetTotalExpiringPointsOk returns a tuple with the TotalExpiringPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalExpiringPoints

`func (o *LoyaltyProjection) HasTotalExpiringPoints() bool`

HasTotalExpiringPoints returns a boolean if a field has been set.

### SetTotalExpiringPoints

`func (o *LoyaltyProjection) SetTotalExpiringPoints(v float32)`

SetTotalExpiringPoints gets a reference to the given float32 and assigns it to the TotalExpiringPoints field.

### GetTotalActivatingPoints

`func (o *LoyaltyProjection) GetTotalActivatingPoints() float32`

GetTotalActivatingPoints returns the TotalActivatingPoints field if non-nil, zero value otherwise.

### GetTotalActivatingPointsOk

`func (o *LoyaltyProjection) GetTotalActivatingPointsOk() (float32, bool)`

GetTotalActivatingPointsOk returns a tuple with the TotalActivatingPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalActivatingPoints

`func (o *LoyaltyProjection) HasTotalActivatingPoints() bool`

HasTotalActivatingPoints returns a boolean if a field has been set.

### SetTotalActivatingPoints

`func (o *LoyaltyProjection) SetTotalActivatingPoints(v float32)`

SetTotalActivatingPoints gets a reference to the given float32 and assigns it to the TotalActivatingPoints field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


