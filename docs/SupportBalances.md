# SupportBalances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threshold** | Pointer to **float32** | The maximum number of loyalty points the support agent is allowed to award for this loyalty program. Not present if the agent has no configured limit.  | [optional] 
**AwardedPoints** | Pointer to **float32** | The total number of loyalty points already awarded to this customer profile by this support agent.  | 
**RemainingBalance** | Pointer to **float32** | The remaining number of loyalty points the support agent can still award to this customer profile. Not present if the agent has no configured limit.  | [optional] 

## Methods

### NewSupportBalances

`func NewSupportBalances(awardedPoints float32, ) *SupportBalances`

NewSupportBalances instantiates a new SupportBalances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportBalancesWithDefaults

`func NewSupportBalancesWithDefaults() *SupportBalances`

NewSupportBalancesWithDefaults instantiates a new SupportBalances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreshold

`func (o *SupportBalances) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *SupportBalances) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *SupportBalances) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *SupportBalances) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetAwardedPoints

`func (o *SupportBalances) GetAwardedPoints() float32`

GetAwardedPoints returns the AwardedPoints field if non-nil, zero value otherwise.

### GetAwardedPointsOk

`func (o *SupportBalances) GetAwardedPointsOk() (*float32, bool)`

GetAwardedPointsOk returns a tuple with the AwardedPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwardedPoints

`func (o *SupportBalances) SetAwardedPoints(v float32)`

SetAwardedPoints sets AwardedPoints field to given value.


### GetRemainingBalance

`func (o *SupportBalances) GetRemainingBalance() float32`

GetRemainingBalance returns the RemainingBalance field if non-nil, zero value otherwise.

### GetRemainingBalanceOk

`func (o *SupportBalances) GetRemainingBalanceOk() (*float32, bool)`

GetRemainingBalanceOk returns a tuple with the RemainingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingBalance

`func (o *SupportBalances) SetRemainingBalance(v float32)`

SetRemainingBalance sets RemainingBalance field to given value.

### HasRemainingBalance

`func (o *SupportBalances) HasRemainingBalance() bool`

HasRemainingBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


