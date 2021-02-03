# RollbackDeductedLoyaltyPointsEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramId** | Pointer to **int32** | The ID of the loyalty program where these points were reimbursed | 
**SubLedgerId** | Pointer to **string** | The ID of the subledger within the loyalty program where these points were reimbursed | 
**Value** | Pointer to **float32** | The amount of reimbursed points that were added | 
**RecipientIntegrationId** | Pointer to **string** | The user for whom these points were reimbursed | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Date after which the reimbursed points will be valid | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Date after which the reimbursed points will expire | [optional] 
**TransactionUUID** | Pointer to **string** | The identifier of &#39;addition&#39; entries added to the ledger as the &#x60;deductLoyaltyPoints&#x60; effect is rolled back | 

## Methods

### GetProgramId

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetProgramId() int32`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetProgramIdOk() (int32, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgramId

`func (o *RollbackDeductedLoyaltyPointsEffectProps) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### SetProgramId

`func (o *RollbackDeductedLoyaltyPointsEffectProps) SetProgramId(v int32)`

SetProgramId gets a reference to the given int32 and assigns it to the ProgramId field.

### GetSubLedgerId

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetSubLedgerId() string`

GetSubLedgerId returns the SubLedgerId field if non-nil, zero value otherwise.

### GetSubLedgerIdOk

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetSubLedgerIdOk() (string, bool)`

GetSubLedgerIdOk returns a tuple with the SubLedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubLedgerId

`func (o *RollbackDeductedLoyaltyPointsEffectProps) HasSubLedgerId() bool`

HasSubLedgerId returns a boolean if a field has been set.

### SetSubLedgerId

`func (o *RollbackDeductedLoyaltyPointsEffectProps) SetSubLedgerId(v string)`

SetSubLedgerId gets a reference to the given string and assigns it to the SubLedgerId field.

### GetValue

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetValueOk() (float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *RollbackDeductedLoyaltyPointsEffectProps) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *RollbackDeductedLoyaltyPointsEffectProps) SetValue(v float32)`

SetValue gets a reference to the given float32 and assigns it to the Value field.

### GetRecipientIntegrationId

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetRecipientIntegrationIdOk() (string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecipientIntegrationId

`func (o *RollbackDeductedLoyaltyPointsEffectProps) HasRecipientIntegrationId() bool`

HasRecipientIntegrationId returns a boolean if a field has been set.

### SetRecipientIntegrationId

`func (o *RollbackDeductedLoyaltyPointsEffectProps) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId gets a reference to the given string and assigns it to the RecipientIntegrationId field.

### GetStartDate

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetStartDateOk() (time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *RollbackDeductedLoyaltyPointsEffectProps) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *RollbackDeductedLoyaltyPointsEffectProps) SetStartDate(v time.Time)`

SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.

### GetExpiryDate

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetExpiryDateOk() (time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *RollbackDeductedLoyaltyPointsEffectProps) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *RollbackDeductedLoyaltyPointsEffectProps) SetExpiryDate(v time.Time)`

SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.

### GetTransactionUUID

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetTransactionUUIDOk() (string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTransactionUUID

`func (o *RollbackDeductedLoyaltyPointsEffectProps) HasTransactionUUID() bool`

HasTransactionUUID returns a boolean if a field has been set.

### SetTransactionUUID

`func (o *RollbackDeductedLoyaltyPointsEffectProps) SetTransactionUUID(v string)`

SetTransactionUUID gets a reference to the given string and assigns it to the TransactionUUID field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


