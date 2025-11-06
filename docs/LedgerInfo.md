# LedgerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentBalance** | Pointer to **float32** | Sum of currently active points. | 
**PendingBalance** | Pointer to **float32** | Sum of pending points. | 
**NegativeBalance** | Pointer to **float32** | Sum of negative points. This implies that &#x60;currentBalance&#x60; is &#x60;0&#x60;. | [optional] 
**ExpiredBalance** | Pointer to **float32** | **DEPRECATED** Value is shown as 0.  | 
**SpentBalance** | Pointer to **float32** | **DEPRECATED** Value is shown as 0.  | 
**TentativeCurrentBalance** | Pointer to **float32** | The tentative points balance, reflecting the &#x60;currentBalance&#x60; and all point additions and deductions within the current open customer session. When the session is closed, the effects are applied and the &#x60;currentBalance&#x60; is updated to this value.  **Note:** Tentative balances are specific to the current session and do not take into account other open sessions for the given customer.  | 
**TentativePendingBalance** | Pointer to **float32** | The tentative points balance, reflecting the &#x60;pendingBalance&#x60; and all point additions with a future activation date within the current open customer session. When the session is closed, the effects are applied and the &#x60;pendingBalance&#x60; is updated to this value.  **Note:** Tentative balances are specific to the current session and do not take into account other open sessions for the given customer.  | [optional] 
**TentativeNegativeBalance** | Pointer to **float32** | The tentative negative balance after all additions and deductions from the current customer session are applied to &#x60;negativeBalance&#x60;. When the session is closed, the tentative effects are applied and &#x60;negativeBalance&#x60; is updated to this value.  **Note:** Tentative balances are specific to the current session and do not take into account other open sessions for the given customer.  | [optional] 
**CurrentTier** | Pointer to [**Tier**](Tier.md) |  | [optional] 
**PointsToNextTier** | Pointer to **float32** | Points required to move up a tier. | [optional] 

## Methods

### NewLedgerInfo

`func NewLedgerInfo(currentBalance float32, pendingBalance float32, expiredBalance float32, spentBalance float32, tentativeCurrentBalance float32, ) *LedgerInfo`

NewLedgerInfo instantiates a new LedgerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerInfoWithDefaults

`func NewLedgerInfoWithDefaults() *LedgerInfo`

NewLedgerInfoWithDefaults instantiates a new LedgerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentBalance

`func (o *LedgerInfo) GetCurrentBalance() float32`

GetCurrentBalance returns the CurrentBalance field if non-nil, zero value otherwise.

### GetCurrentBalanceOk

`func (o *LedgerInfo) GetCurrentBalanceOk() (*float32, bool)`

GetCurrentBalanceOk returns a tuple with the CurrentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBalance

`func (o *LedgerInfo) SetCurrentBalance(v float32)`

SetCurrentBalance sets CurrentBalance field to given value.


### GetPendingBalance

`func (o *LedgerInfo) GetPendingBalance() float32`

GetPendingBalance returns the PendingBalance field if non-nil, zero value otherwise.

### GetPendingBalanceOk

`func (o *LedgerInfo) GetPendingBalanceOk() (*float32, bool)`

GetPendingBalanceOk returns a tuple with the PendingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingBalance

`func (o *LedgerInfo) SetPendingBalance(v float32)`

SetPendingBalance sets PendingBalance field to given value.


### GetNegativeBalance

`func (o *LedgerInfo) GetNegativeBalance() float32`

GetNegativeBalance returns the NegativeBalance field if non-nil, zero value otherwise.

### GetNegativeBalanceOk

`func (o *LedgerInfo) GetNegativeBalanceOk() (*float32, bool)`

GetNegativeBalanceOk returns a tuple with the NegativeBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeBalance

`func (o *LedgerInfo) SetNegativeBalance(v float32)`

SetNegativeBalance sets NegativeBalance field to given value.

### HasNegativeBalance

`func (o *LedgerInfo) HasNegativeBalance() bool`

HasNegativeBalance returns a boolean if a field has been set.

### GetExpiredBalance

`func (o *LedgerInfo) GetExpiredBalance() float32`

GetExpiredBalance returns the ExpiredBalance field if non-nil, zero value otherwise.

### GetExpiredBalanceOk

`func (o *LedgerInfo) GetExpiredBalanceOk() (*float32, bool)`

GetExpiredBalanceOk returns a tuple with the ExpiredBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredBalance

`func (o *LedgerInfo) SetExpiredBalance(v float32)`

SetExpiredBalance sets ExpiredBalance field to given value.


### GetSpentBalance

`func (o *LedgerInfo) GetSpentBalance() float32`

GetSpentBalance returns the SpentBalance field if non-nil, zero value otherwise.

### GetSpentBalanceOk

`func (o *LedgerInfo) GetSpentBalanceOk() (*float32, bool)`

GetSpentBalanceOk returns a tuple with the SpentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentBalance

`func (o *LedgerInfo) SetSpentBalance(v float32)`

SetSpentBalance sets SpentBalance field to given value.


### GetTentativeCurrentBalance

`func (o *LedgerInfo) GetTentativeCurrentBalance() float32`

GetTentativeCurrentBalance returns the TentativeCurrentBalance field if non-nil, zero value otherwise.

### GetTentativeCurrentBalanceOk

`func (o *LedgerInfo) GetTentativeCurrentBalanceOk() (*float32, bool)`

GetTentativeCurrentBalanceOk returns a tuple with the TentativeCurrentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTentativeCurrentBalance

`func (o *LedgerInfo) SetTentativeCurrentBalance(v float32)`

SetTentativeCurrentBalance sets TentativeCurrentBalance field to given value.


### GetTentativePendingBalance

`func (o *LedgerInfo) GetTentativePendingBalance() float32`

GetTentativePendingBalance returns the TentativePendingBalance field if non-nil, zero value otherwise.

### GetTentativePendingBalanceOk

`func (o *LedgerInfo) GetTentativePendingBalanceOk() (*float32, bool)`

GetTentativePendingBalanceOk returns a tuple with the TentativePendingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTentativePendingBalance

`func (o *LedgerInfo) SetTentativePendingBalance(v float32)`

SetTentativePendingBalance sets TentativePendingBalance field to given value.

### HasTentativePendingBalance

`func (o *LedgerInfo) HasTentativePendingBalance() bool`

HasTentativePendingBalance returns a boolean if a field has been set.

### GetTentativeNegativeBalance

`func (o *LedgerInfo) GetTentativeNegativeBalance() float32`

GetTentativeNegativeBalance returns the TentativeNegativeBalance field if non-nil, zero value otherwise.

### GetTentativeNegativeBalanceOk

`func (o *LedgerInfo) GetTentativeNegativeBalanceOk() (*float32, bool)`

GetTentativeNegativeBalanceOk returns a tuple with the TentativeNegativeBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTentativeNegativeBalance

`func (o *LedgerInfo) SetTentativeNegativeBalance(v float32)`

SetTentativeNegativeBalance sets TentativeNegativeBalance field to given value.

### HasTentativeNegativeBalance

`func (o *LedgerInfo) HasTentativeNegativeBalance() bool`

HasTentativeNegativeBalance returns a boolean if a field has been set.

### GetCurrentTier

`func (o *LedgerInfo) GetCurrentTier() Tier`

GetCurrentTier returns the CurrentTier field if non-nil, zero value otherwise.

### GetCurrentTierOk

`func (o *LedgerInfo) GetCurrentTierOk() (*Tier, bool)`

GetCurrentTierOk returns a tuple with the CurrentTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTier

`func (o *LedgerInfo) SetCurrentTier(v Tier)`

SetCurrentTier sets CurrentTier field to given value.

### HasCurrentTier

`func (o *LedgerInfo) HasCurrentTier() bool`

HasCurrentTier returns a boolean if a field has been set.

### GetPointsToNextTier

`func (o *LedgerInfo) GetPointsToNextTier() float32`

GetPointsToNextTier returns the PointsToNextTier field if non-nil, zero value otherwise.

### GetPointsToNextTierOk

`func (o *LedgerInfo) GetPointsToNextTierOk() (*float32, bool)`

GetPointsToNextTierOk returns a tuple with the PointsToNextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsToNextTier

`func (o *LedgerInfo) SetPointsToNextTier(v float32)`

SetPointsToNextTier sets PointsToNextTier field to given value.

### HasPointsToNextTier

`func (o *LedgerInfo) HasPointsToNextTier() bool`

HasPointsToNextTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


