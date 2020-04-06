# DeductLoyaltyPointsEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleTitle** | Pointer to **string** | The title of the rule that contained triggered this points deduction | 
**ProgramId** | Pointer to **int32** | The ID of the loyalty program where these points were added | 
**SubLedgerId** | Pointer to **string** | The ID of the subledger within the loyalty program where these points were added | 
**Value** | Pointer to **float32** | The amount of points that were deducted | 

## Methods

### GetRuleTitle

`func (o *DeductLoyaltyPointsEffectProps) GetRuleTitle() string`

GetRuleTitle returns the RuleTitle field if non-nil, zero value otherwise.

### GetRuleTitleOk

`func (o *DeductLoyaltyPointsEffectProps) GetRuleTitleOk() (string, bool)`

GetRuleTitleOk returns a tuple with the RuleTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRuleTitle

`func (o *DeductLoyaltyPointsEffectProps) HasRuleTitle() bool`

HasRuleTitle returns a boolean if a field has been set.

### SetRuleTitle

`func (o *DeductLoyaltyPointsEffectProps) SetRuleTitle(v string)`

SetRuleTitle gets a reference to the given string and assigns it to the RuleTitle field.

### GetProgramId

`func (o *DeductLoyaltyPointsEffectProps) GetProgramId() int32`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *DeductLoyaltyPointsEffectProps) GetProgramIdOk() (int32, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgramId

`func (o *DeductLoyaltyPointsEffectProps) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### SetProgramId

`func (o *DeductLoyaltyPointsEffectProps) SetProgramId(v int32)`

SetProgramId gets a reference to the given int32 and assigns it to the ProgramId field.

### GetSubLedgerId

`func (o *DeductLoyaltyPointsEffectProps) GetSubLedgerId() string`

GetSubLedgerId returns the SubLedgerId field if non-nil, zero value otherwise.

### GetSubLedgerIdOk

`func (o *DeductLoyaltyPointsEffectProps) GetSubLedgerIdOk() (string, bool)`

GetSubLedgerIdOk returns a tuple with the SubLedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubLedgerId

`func (o *DeductLoyaltyPointsEffectProps) HasSubLedgerId() bool`

HasSubLedgerId returns a boolean if a field has been set.

### SetSubLedgerId

`func (o *DeductLoyaltyPointsEffectProps) SetSubLedgerId(v string)`

SetSubLedgerId gets a reference to the given string and assigns it to the SubLedgerId field.

### GetValue

`func (o *DeductLoyaltyPointsEffectProps) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DeductLoyaltyPointsEffectProps) GetValueOk() (float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *DeductLoyaltyPointsEffectProps) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *DeductLoyaltyPointsEffectProps) SetValue(v float32)`

SetValue gets a reference to the given float32 and assigns it to the Value field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


