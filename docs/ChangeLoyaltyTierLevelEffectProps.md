# ChangeLoyaltyTierLevelEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | The expiration date of the new tier. | [optional] 
**NewTierName** | Pointer to **string** | The name of the tier to which the user has been upgraded. | 
**PreviousTierName** | Pointer to **string** | The name of the tier from which the user was upgraded. | [optional] 
**ProgramId** | Pointer to **int32** | The ID of the loyalty program where these points were added. | 
**RuleTitle** | Pointer to **string** | The title of the rule that triggered the tier upgrade. | 
**SubLedgerId** | Pointer to **string** | The ID of the subledger within the loyalty program where these points were added. | 

## Methods

### GetExpiryDate

`func (o *ChangeLoyaltyTierLevelEffectProps) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *ChangeLoyaltyTierLevelEffectProps) GetExpiryDateOk() (time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *ChangeLoyaltyTierLevelEffectProps) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *ChangeLoyaltyTierLevelEffectProps) SetExpiryDate(v time.Time)`

SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.

### GetNewTierName

`func (o *ChangeLoyaltyTierLevelEffectProps) GetNewTierName() string`

GetNewTierName returns the NewTierName field if non-nil, zero value otherwise.

### GetNewTierNameOk

`func (o *ChangeLoyaltyTierLevelEffectProps) GetNewTierNameOk() (string, bool)`

GetNewTierNameOk returns a tuple with the NewTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNewTierName

`func (o *ChangeLoyaltyTierLevelEffectProps) HasNewTierName() bool`

HasNewTierName returns a boolean if a field has been set.

### SetNewTierName

`func (o *ChangeLoyaltyTierLevelEffectProps) SetNewTierName(v string)`

SetNewTierName gets a reference to the given string and assigns it to the NewTierName field.

### GetPreviousTierName

`func (o *ChangeLoyaltyTierLevelEffectProps) GetPreviousTierName() string`

GetPreviousTierName returns the PreviousTierName field if non-nil, zero value otherwise.

### GetPreviousTierNameOk

`func (o *ChangeLoyaltyTierLevelEffectProps) GetPreviousTierNameOk() (string, bool)`

GetPreviousTierNameOk returns a tuple with the PreviousTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreviousTierName

`func (o *ChangeLoyaltyTierLevelEffectProps) HasPreviousTierName() bool`

HasPreviousTierName returns a boolean if a field has been set.

### SetPreviousTierName

`func (o *ChangeLoyaltyTierLevelEffectProps) SetPreviousTierName(v string)`

SetPreviousTierName gets a reference to the given string and assigns it to the PreviousTierName field.

### GetProgramId

`func (o *ChangeLoyaltyTierLevelEffectProps) GetProgramId() int32`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *ChangeLoyaltyTierLevelEffectProps) GetProgramIdOk() (int32, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgramId

`func (o *ChangeLoyaltyTierLevelEffectProps) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### SetProgramId

`func (o *ChangeLoyaltyTierLevelEffectProps) SetProgramId(v int32)`

SetProgramId gets a reference to the given int32 and assigns it to the ProgramId field.

### GetRuleTitle

`func (o *ChangeLoyaltyTierLevelEffectProps) GetRuleTitle() string`

GetRuleTitle returns the RuleTitle field if non-nil, zero value otherwise.

### GetRuleTitleOk

`func (o *ChangeLoyaltyTierLevelEffectProps) GetRuleTitleOk() (string, bool)`

GetRuleTitleOk returns a tuple with the RuleTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRuleTitle

`func (o *ChangeLoyaltyTierLevelEffectProps) HasRuleTitle() bool`

HasRuleTitle returns a boolean if a field has been set.

### SetRuleTitle

`func (o *ChangeLoyaltyTierLevelEffectProps) SetRuleTitle(v string)`

SetRuleTitle gets a reference to the given string and assigns it to the RuleTitle field.

### GetSubLedgerId

`func (o *ChangeLoyaltyTierLevelEffectProps) GetSubLedgerId() string`

GetSubLedgerId returns the SubLedgerId field if non-nil, zero value otherwise.

### GetSubLedgerIdOk

`func (o *ChangeLoyaltyTierLevelEffectProps) GetSubLedgerIdOk() (string, bool)`

GetSubLedgerIdOk returns a tuple with the SubLedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubLedgerId

`func (o *ChangeLoyaltyTierLevelEffectProps) HasSubLedgerId() bool`

HasSubLedgerId returns a boolean if a field has been set.

### SetSubLedgerId

`func (o *ChangeLoyaltyTierLevelEffectProps) SetSubLedgerId(v string)`

SetSubLedgerId gets a reference to the given string and assigns it to the SubLedgerId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


