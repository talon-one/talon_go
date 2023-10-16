# RollbackAddedLoyaltyPointsEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramId** | Pointer to **int32** | The ID of the loyalty program where the points were originally added. | 
**SubLedgerId** | Pointer to **string** | The ID of the subledger within the loyalty program where these points were originally added. | 
**Value** | Pointer to **float32** | The amount of points that were rolled back. | 
**RecipientIntegrationId** | Pointer to **string** | The user for whom these points were originally added. | 
**TransactionUUID** | Pointer to **string** | The identifier of &#39;deduction&#39; entry added to the ledger as the &#x60;addLoyaltyPoints&#x60; effect is rolled back. | 
**CartItemPosition** | Pointer to **float32** | The index of the item in the cart items for which the loyalty points were rolled back. | [optional] 
**CartItemSubPosition** | Pointer to **float32** | For cart items with &#x60;quantity&#x60; &gt; 1, the sub-position indicates to which item the loyalty points were rolled back.  | [optional] 
**CardIdentifier** | Pointer to **string** | The alphanumeric identifier of the loyalty card.  | [optional] 

## Methods

### GetProgramId

`func (o *RollbackAddedLoyaltyPointsEffectProps) GetProgramId() int32`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *RollbackAddedLoyaltyPointsEffectProps) GetProgramIdOk() (int32, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgramId

`func (o *RollbackAddedLoyaltyPointsEffectProps) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### SetProgramId

`func (o *RollbackAddedLoyaltyPointsEffectProps) SetProgramId(v int32)`

SetProgramId gets a reference to the given int32 and assigns it to the ProgramId field.

### GetSubLedgerId

`func (o *RollbackAddedLoyaltyPointsEffectProps) GetSubLedgerId() string`

GetSubLedgerId returns the SubLedgerId field if non-nil, zero value otherwise.

### GetSubLedgerIdOk

`func (o *RollbackAddedLoyaltyPointsEffectProps) GetSubLedgerIdOk() (string, bool)`

GetSubLedgerIdOk returns a tuple with the SubLedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubLedgerId

`func (o *RollbackAddedLoyaltyPointsEffectProps) HasSubLedgerId() bool`

HasSubLedgerId returns a boolean if a field has been set.

### SetSubLedgerId

`func (o *RollbackAddedLoyaltyPointsEffectProps) SetSubLedgerId(v string)`

SetSubLedgerId gets a reference to the given string and assigns it to the SubLedgerId field.

### GetValue

`func (o *RollbackAddedLoyaltyPointsEffectProps) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RollbackAddedLoyaltyPointsEffectProps) GetValueOk() (float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *RollbackAddedLoyaltyPointsEffectProps) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *RollbackAddedLoyaltyPointsEffectProps) SetValue(v float32)`

SetValue gets a reference to the given float32 and assigns it to the Value field.

### GetRecipientIntegrationId

`func (o *RollbackAddedLoyaltyPointsEffectProps) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *RollbackAddedLoyaltyPointsEffectProps) GetRecipientIntegrationIdOk() (string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecipientIntegrationId

`func (o *RollbackAddedLoyaltyPointsEffectProps) HasRecipientIntegrationId() bool`

HasRecipientIntegrationId returns a boolean if a field has been set.

### SetRecipientIntegrationId

`func (o *RollbackAddedLoyaltyPointsEffectProps) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId gets a reference to the given string and assigns it to the RecipientIntegrationId field.

### GetTransactionUUID

`func (o *RollbackAddedLoyaltyPointsEffectProps) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *RollbackAddedLoyaltyPointsEffectProps) GetTransactionUUIDOk() (string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTransactionUUID

`func (o *RollbackAddedLoyaltyPointsEffectProps) HasTransactionUUID() bool`

HasTransactionUUID returns a boolean if a field has been set.

### SetTransactionUUID

`func (o *RollbackAddedLoyaltyPointsEffectProps) SetTransactionUUID(v string)`

SetTransactionUUID gets a reference to the given string and assigns it to the TransactionUUID field.

### GetCartItemPosition

`func (o *RollbackAddedLoyaltyPointsEffectProps) GetCartItemPosition() float32`

GetCartItemPosition returns the CartItemPosition field if non-nil, zero value otherwise.

### GetCartItemPositionOk

`func (o *RollbackAddedLoyaltyPointsEffectProps) GetCartItemPositionOk() (float32, bool)`

GetCartItemPositionOk returns a tuple with the CartItemPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCartItemPosition

`func (o *RollbackAddedLoyaltyPointsEffectProps) HasCartItemPosition() bool`

HasCartItemPosition returns a boolean if a field has been set.

### SetCartItemPosition

`func (o *RollbackAddedLoyaltyPointsEffectProps) SetCartItemPosition(v float32)`

SetCartItemPosition gets a reference to the given float32 and assigns it to the CartItemPosition field.

### GetCartItemSubPosition

`func (o *RollbackAddedLoyaltyPointsEffectProps) GetCartItemSubPosition() float32`

GetCartItemSubPosition returns the CartItemSubPosition field if non-nil, zero value otherwise.

### GetCartItemSubPositionOk

`func (o *RollbackAddedLoyaltyPointsEffectProps) GetCartItemSubPositionOk() (float32, bool)`

GetCartItemSubPositionOk returns a tuple with the CartItemSubPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCartItemSubPosition

`func (o *RollbackAddedLoyaltyPointsEffectProps) HasCartItemSubPosition() bool`

HasCartItemSubPosition returns a boolean if a field has been set.

### SetCartItemSubPosition

`func (o *RollbackAddedLoyaltyPointsEffectProps) SetCartItemSubPosition(v float32)`

SetCartItemSubPosition gets a reference to the given float32 and assigns it to the CartItemSubPosition field.

### GetCardIdentifier

`func (o *RollbackAddedLoyaltyPointsEffectProps) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *RollbackAddedLoyaltyPointsEffectProps) GetCardIdentifierOk() (string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCardIdentifier

`func (o *RollbackAddedLoyaltyPointsEffectProps) HasCardIdentifier() bool`

HasCardIdentifier returns a boolean if a field has been set.

### SetCardIdentifier

`func (o *RollbackAddedLoyaltyPointsEffectProps) SetCardIdentifier(v string)`

SetCardIdentifier gets a reference to the given string and assigns it to the CardIdentifier field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


