# LoyaltyBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePoints** | Pointer to **float32** | Total amount of points awarded to this customer and available to spend. | [optional] 
**PendingPoints** | Pointer to **float32** | Total amount of points awarded to this customer but not available until their start date. | [optional] 
**SpentPoints** | Pointer to **float32** | Total amount of points already spent by this customer. | [optional] 
**ExpiredPoints** | Pointer to **float32** | Total amount of points awarded but never redeemed. They cannot be used anymore. | [optional] 
**NegativePoints** | Pointer to **float32** | Total amount of negative points. This implies that &#x60;activePoints&#x60; is &#x60;0&#x60;. | [optional] 

## Methods

### NewLoyaltyBalance

`func NewLoyaltyBalance() *LoyaltyBalance`

NewLoyaltyBalance instantiates a new LoyaltyBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyBalanceWithDefaults

`func NewLoyaltyBalanceWithDefaults() *LoyaltyBalance`

NewLoyaltyBalanceWithDefaults instantiates a new LoyaltyBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePoints

`func (o *LoyaltyBalance) GetActivePoints() float32`

GetActivePoints returns the ActivePoints field if non-nil, zero value otherwise.

### GetActivePointsOk

`func (o *LoyaltyBalance) GetActivePointsOk() (*float32, bool)`

GetActivePointsOk returns a tuple with the ActivePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePoints

`func (o *LoyaltyBalance) SetActivePoints(v float32)`

SetActivePoints sets ActivePoints field to given value.

### HasActivePoints

`func (o *LoyaltyBalance) HasActivePoints() bool`

HasActivePoints returns a boolean if a field has been set.

### GetPendingPoints

`func (o *LoyaltyBalance) GetPendingPoints() float32`

GetPendingPoints returns the PendingPoints field if non-nil, zero value otherwise.

### GetPendingPointsOk

`func (o *LoyaltyBalance) GetPendingPointsOk() (*float32, bool)`

GetPendingPointsOk returns a tuple with the PendingPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingPoints

`func (o *LoyaltyBalance) SetPendingPoints(v float32)`

SetPendingPoints sets PendingPoints field to given value.

### HasPendingPoints

`func (o *LoyaltyBalance) HasPendingPoints() bool`

HasPendingPoints returns a boolean if a field has been set.

### GetSpentPoints

`func (o *LoyaltyBalance) GetSpentPoints() float32`

GetSpentPoints returns the SpentPoints field if non-nil, zero value otherwise.

### GetSpentPointsOk

`func (o *LoyaltyBalance) GetSpentPointsOk() (*float32, bool)`

GetSpentPointsOk returns a tuple with the SpentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentPoints

`func (o *LoyaltyBalance) SetSpentPoints(v float32)`

SetSpentPoints sets SpentPoints field to given value.

### HasSpentPoints

`func (o *LoyaltyBalance) HasSpentPoints() bool`

HasSpentPoints returns a boolean if a field has been set.

### GetExpiredPoints

`func (o *LoyaltyBalance) GetExpiredPoints() float32`

GetExpiredPoints returns the ExpiredPoints field if non-nil, zero value otherwise.

### GetExpiredPointsOk

`func (o *LoyaltyBalance) GetExpiredPointsOk() (*float32, bool)`

GetExpiredPointsOk returns a tuple with the ExpiredPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredPoints

`func (o *LoyaltyBalance) SetExpiredPoints(v float32)`

SetExpiredPoints sets ExpiredPoints field to given value.

### HasExpiredPoints

`func (o *LoyaltyBalance) HasExpiredPoints() bool`

HasExpiredPoints returns a boolean if a field has been set.

### GetNegativePoints

`func (o *LoyaltyBalance) GetNegativePoints() float32`

GetNegativePoints returns the NegativePoints field if non-nil, zero value otherwise.

### GetNegativePointsOk

`func (o *LoyaltyBalance) GetNegativePointsOk() (*float32, bool)`

GetNegativePointsOk returns a tuple with the NegativePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativePoints

`func (o *LoyaltyBalance) SetNegativePoints(v float32)`

SetNegativePoints sets NegativePoints field to given value.

### HasNegativePoints

`func (o *LoyaltyBalance) HasNegativePoints() bool`

HasNegativePoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


