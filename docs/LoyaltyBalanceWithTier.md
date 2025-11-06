# LoyaltyBalanceWithTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePoints** | Pointer to **float32** | Total amount of points awarded to this customer and available to spend. | [optional] 
**PendingPoints** | Pointer to **float32** | Total amount of points awarded to this customer but not available until their start date. | [optional] 
**SpentPoints** | Pointer to **float32** | Total amount of points already spent by this customer. | [optional] 
**ExpiredPoints** | Pointer to **float32** | Total amount of points awarded but never redeemed. They cannot be used anymore. | [optional] 
**NegativePoints** | Pointer to **float32** | Total amount of negative points. This implies that &#x60;activePoints&#x60; is &#x60;0&#x60;. | [optional] 
**CurrentTier** | Pointer to [**Tier**](Tier.md) |  | [optional] 
**ProjectedTier** | Pointer to [**ProjectedTier**](ProjectedTier.md) |  | [optional] 
**PointsToNextTier** | Pointer to **float32** | The number of points required to move up a tier. | [optional] 
**NextTierName** | Pointer to **string** | The name of the tier consecutive to the current tier. | [optional] 

## Methods

### NewLoyaltyBalanceWithTier

`func NewLoyaltyBalanceWithTier() *LoyaltyBalanceWithTier`

NewLoyaltyBalanceWithTier instantiates a new LoyaltyBalanceWithTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyBalanceWithTierWithDefaults

`func NewLoyaltyBalanceWithTierWithDefaults() *LoyaltyBalanceWithTier`

NewLoyaltyBalanceWithTierWithDefaults instantiates a new LoyaltyBalanceWithTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePoints

`func (o *LoyaltyBalanceWithTier) GetActivePoints() float32`

GetActivePoints returns the ActivePoints field if non-nil, zero value otherwise.

### GetActivePointsOk

`func (o *LoyaltyBalanceWithTier) GetActivePointsOk() (*float32, bool)`

GetActivePointsOk returns a tuple with the ActivePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePoints

`func (o *LoyaltyBalanceWithTier) SetActivePoints(v float32)`

SetActivePoints sets ActivePoints field to given value.

### HasActivePoints

`func (o *LoyaltyBalanceWithTier) HasActivePoints() bool`

HasActivePoints returns a boolean if a field has been set.

### GetPendingPoints

`func (o *LoyaltyBalanceWithTier) GetPendingPoints() float32`

GetPendingPoints returns the PendingPoints field if non-nil, zero value otherwise.

### GetPendingPointsOk

`func (o *LoyaltyBalanceWithTier) GetPendingPointsOk() (*float32, bool)`

GetPendingPointsOk returns a tuple with the PendingPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingPoints

`func (o *LoyaltyBalanceWithTier) SetPendingPoints(v float32)`

SetPendingPoints sets PendingPoints field to given value.

### HasPendingPoints

`func (o *LoyaltyBalanceWithTier) HasPendingPoints() bool`

HasPendingPoints returns a boolean if a field has been set.

### GetSpentPoints

`func (o *LoyaltyBalanceWithTier) GetSpentPoints() float32`

GetSpentPoints returns the SpentPoints field if non-nil, zero value otherwise.

### GetSpentPointsOk

`func (o *LoyaltyBalanceWithTier) GetSpentPointsOk() (*float32, bool)`

GetSpentPointsOk returns a tuple with the SpentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentPoints

`func (o *LoyaltyBalanceWithTier) SetSpentPoints(v float32)`

SetSpentPoints sets SpentPoints field to given value.

### HasSpentPoints

`func (o *LoyaltyBalanceWithTier) HasSpentPoints() bool`

HasSpentPoints returns a boolean if a field has been set.

### GetExpiredPoints

`func (o *LoyaltyBalanceWithTier) GetExpiredPoints() float32`

GetExpiredPoints returns the ExpiredPoints field if non-nil, zero value otherwise.

### GetExpiredPointsOk

`func (o *LoyaltyBalanceWithTier) GetExpiredPointsOk() (*float32, bool)`

GetExpiredPointsOk returns a tuple with the ExpiredPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredPoints

`func (o *LoyaltyBalanceWithTier) SetExpiredPoints(v float32)`

SetExpiredPoints sets ExpiredPoints field to given value.

### HasExpiredPoints

`func (o *LoyaltyBalanceWithTier) HasExpiredPoints() bool`

HasExpiredPoints returns a boolean if a field has been set.

### GetNegativePoints

`func (o *LoyaltyBalanceWithTier) GetNegativePoints() float32`

GetNegativePoints returns the NegativePoints field if non-nil, zero value otherwise.

### GetNegativePointsOk

`func (o *LoyaltyBalanceWithTier) GetNegativePointsOk() (*float32, bool)`

GetNegativePointsOk returns a tuple with the NegativePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativePoints

`func (o *LoyaltyBalanceWithTier) SetNegativePoints(v float32)`

SetNegativePoints sets NegativePoints field to given value.

### HasNegativePoints

`func (o *LoyaltyBalanceWithTier) HasNegativePoints() bool`

HasNegativePoints returns a boolean if a field has been set.

### GetCurrentTier

`func (o *LoyaltyBalanceWithTier) GetCurrentTier() Tier`

GetCurrentTier returns the CurrentTier field if non-nil, zero value otherwise.

### GetCurrentTierOk

`func (o *LoyaltyBalanceWithTier) GetCurrentTierOk() (*Tier, bool)`

GetCurrentTierOk returns a tuple with the CurrentTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTier

`func (o *LoyaltyBalanceWithTier) SetCurrentTier(v Tier)`

SetCurrentTier sets CurrentTier field to given value.

### HasCurrentTier

`func (o *LoyaltyBalanceWithTier) HasCurrentTier() bool`

HasCurrentTier returns a boolean if a field has been set.

### GetProjectedTier

`func (o *LoyaltyBalanceWithTier) GetProjectedTier() ProjectedTier`

GetProjectedTier returns the ProjectedTier field if non-nil, zero value otherwise.

### GetProjectedTierOk

`func (o *LoyaltyBalanceWithTier) GetProjectedTierOk() (*ProjectedTier, bool)`

GetProjectedTierOk returns a tuple with the ProjectedTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedTier

`func (o *LoyaltyBalanceWithTier) SetProjectedTier(v ProjectedTier)`

SetProjectedTier sets ProjectedTier field to given value.

### HasProjectedTier

`func (o *LoyaltyBalanceWithTier) HasProjectedTier() bool`

HasProjectedTier returns a boolean if a field has been set.

### GetPointsToNextTier

`func (o *LoyaltyBalanceWithTier) GetPointsToNextTier() float32`

GetPointsToNextTier returns the PointsToNextTier field if non-nil, zero value otherwise.

### GetPointsToNextTierOk

`func (o *LoyaltyBalanceWithTier) GetPointsToNextTierOk() (*float32, bool)`

GetPointsToNextTierOk returns a tuple with the PointsToNextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsToNextTier

`func (o *LoyaltyBalanceWithTier) SetPointsToNextTier(v float32)`

SetPointsToNextTier sets PointsToNextTier field to given value.

### HasPointsToNextTier

`func (o *LoyaltyBalanceWithTier) HasPointsToNextTier() bool`

HasPointsToNextTier returns a boolean if a field has been set.

### GetNextTierName

`func (o *LoyaltyBalanceWithTier) GetNextTierName() string`

GetNextTierName returns the NextTierName field if non-nil, zero value otherwise.

### GetNextTierNameOk

`func (o *LoyaltyBalanceWithTier) GetNextTierNameOk() (*string, bool)`

GetNextTierNameOk returns a tuple with the NextTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTierName

`func (o *LoyaltyBalanceWithTier) SetNextTierName(v string)`

SetNextTierName sets NextTierName field to given value.

### HasNextTierName

`func (o *LoyaltyBalanceWithTier) HasNextTierName() bool`

HasNextTierName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


