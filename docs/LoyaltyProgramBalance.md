# LoyaltyProgramBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentBalance** | Pointer to **float32** | Sum of currently active points. | 
**PendingBalance** | Pointer to **float32** | Sum of pending points. | 
**ExpiredBalance** | Pointer to **float32** | **DEPRECATED** Value is shown as 0.  | 
**SpentBalance** | Pointer to **float32** | **DEPRECATED** Value is shown as 0.  | 
**TentativeCurrentBalance** | Pointer to **float32** | The tentative points balance, reflecting the &#x60;currentBalance&#x60; and all point additions and deductions within the current open customer session. When the session is closed, the effects are applied and the &#x60;currentBalance&#x60; is updated to this value.  **Note:** Tentative balances are specific to the current session and do not take into account other open sessions for the given customer.  | 
**TentativePendingBalance** | Pointer to **float32** | The tentative points balance, reflecting the &#x60;pendingBalance&#x60; and all point additions with a future activation date within the current open customer session. When the session is closed, the effects are applied and the &#x60;pendingBalance&#x60; is updated to this value.  **Note:** Tentative balances are specific to the current session and do not take into account other open sessions for the given customer.  | [optional] 

## Methods

### GetCurrentBalance

`func (o *LoyaltyProgramBalance) GetCurrentBalance() float32`

GetCurrentBalance returns the CurrentBalance field if non-nil, zero value otherwise.

### GetCurrentBalanceOk

`func (o *LoyaltyProgramBalance) GetCurrentBalanceOk() (float32, bool)`

GetCurrentBalanceOk returns a tuple with the CurrentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrentBalance

`func (o *LoyaltyProgramBalance) HasCurrentBalance() bool`

HasCurrentBalance returns a boolean if a field has been set.

### SetCurrentBalance

`func (o *LoyaltyProgramBalance) SetCurrentBalance(v float32)`

SetCurrentBalance gets a reference to the given float32 and assigns it to the CurrentBalance field.

### GetPendingBalance

`func (o *LoyaltyProgramBalance) GetPendingBalance() float32`

GetPendingBalance returns the PendingBalance field if non-nil, zero value otherwise.

### GetPendingBalanceOk

`func (o *LoyaltyProgramBalance) GetPendingBalanceOk() (float32, bool)`

GetPendingBalanceOk returns a tuple with the PendingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPendingBalance

`func (o *LoyaltyProgramBalance) HasPendingBalance() bool`

HasPendingBalance returns a boolean if a field has been set.

### SetPendingBalance

`func (o *LoyaltyProgramBalance) SetPendingBalance(v float32)`

SetPendingBalance gets a reference to the given float32 and assigns it to the PendingBalance field.

### GetExpiredBalance

`func (o *LoyaltyProgramBalance) GetExpiredBalance() float32`

GetExpiredBalance returns the ExpiredBalance field if non-nil, zero value otherwise.

### GetExpiredBalanceOk

`func (o *LoyaltyProgramBalance) GetExpiredBalanceOk() (float32, bool)`

GetExpiredBalanceOk returns a tuple with the ExpiredBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiredBalance

`func (o *LoyaltyProgramBalance) HasExpiredBalance() bool`

HasExpiredBalance returns a boolean if a field has been set.

### SetExpiredBalance

`func (o *LoyaltyProgramBalance) SetExpiredBalance(v float32)`

SetExpiredBalance gets a reference to the given float32 and assigns it to the ExpiredBalance field.

### GetSpentBalance

`func (o *LoyaltyProgramBalance) GetSpentBalance() float32`

GetSpentBalance returns the SpentBalance field if non-nil, zero value otherwise.

### GetSpentBalanceOk

`func (o *LoyaltyProgramBalance) GetSpentBalanceOk() (float32, bool)`

GetSpentBalanceOk returns a tuple with the SpentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpentBalance

`func (o *LoyaltyProgramBalance) HasSpentBalance() bool`

HasSpentBalance returns a boolean if a field has been set.

### SetSpentBalance

`func (o *LoyaltyProgramBalance) SetSpentBalance(v float32)`

SetSpentBalance gets a reference to the given float32 and assigns it to the SpentBalance field.

### GetTentativeCurrentBalance

`func (o *LoyaltyProgramBalance) GetTentativeCurrentBalance() float32`

GetTentativeCurrentBalance returns the TentativeCurrentBalance field if non-nil, zero value otherwise.

### GetTentativeCurrentBalanceOk

`func (o *LoyaltyProgramBalance) GetTentativeCurrentBalanceOk() (float32, bool)`

GetTentativeCurrentBalanceOk returns a tuple with the TentativeCurrentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTentativeCurrentBalance

`func (o *LoyaltyProgramBalance) HasTentativeCurrentBalance() bool`

HasTentativeCurrentBalance returns a boolean if a field has been set.

### SetTentativeCurrentBalance

`func (o *LoyaltyProgramBalance) SetTentativeCurrentBalance(v float32)`

SetTentativeCurrentBalance gets a reference to the given float32 and assigns it to the TentativeCurrentBalance field.

### GetTentativePendingBalance

`func (o *LoyaltyProgramBalance) GetTentativePendingBalance() float32`

GetTentativePendingBalance returns the TentativePendingBalance field if non-nil, zero value otherwise.

### GetTentativePendingBalanceOk

`func (o *LoyaltyProgramBalance) GetTentativePendingBalanceOk() (float32, bool)`

GetTentativePendingBalanceOk returns a tuple with the TentativePendingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTentativePendingBalance

`func (o *LoyaltyProgramBalance) HasTentativePendingBalance() bool`

HasTentativePendingBalance returns a boolean if a field has been set.

### SetTentativePendingBalance

`func (o *LoyaltyProgramBalance) SetTentativePendingBalance(v float32)`

SetTentativePendingBalance gets a reference to the given float32 and assigns it to the TentativePendingBalance field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


