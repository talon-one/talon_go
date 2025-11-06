# ChangeLoyaltyTierLevelEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleTitle** | Pointer to **string** | The title of the rule that triggered the tier upgrade. | 
**ProgramId** | Pointer to **int64** | The ID of the loyalty program where these points were added. | 
**SubLedgerId** | Pointer to **string** | The ID of the subledger within the loyalty program where these points were added. | 
**PreviousTierName** | Pointer to **string** | The name of the tier from which the user was upgraded. | [optional] 
**NewTierName** | Pointer to **string** | The name of the tier to which the user has been upgraded. | 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | The expiration date of the new tier. | [optional] 

## Methods

### NewChangeLoyaltyTierLevelEffectProps

`func NewChangeLoyaltyTierLevelEffectProps(ruleTitle string, programId int64, subLedgerId string, newTierName string, ) *ChangeLoyaltyTierLevelEffectProps`

NewChangeLoyaltyTierLevelEffectProps instantiates a new ChangeLoyaltyTierLevelEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeLoyaltyTierLevelEffectPropsWithDefaults

`func NewChangeLoyaltyTierLevelEffectPropsWithDefaults() *ChangeLoyaltyTierLevelEffectProps`

NewChangeLoyaltyTierLevelEffectPropsWithDefaults instantiates a new ChangeLoyaltyTierLevelEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleTitle

`func (o *ChangeLoyaltyTierLevelEffectProps) GetRuleTitle() string`

GetRuleTitle returns the RuleTitle field if non-nil, zero value otherwise.

### GetRuleTitleOk

`func (o *ChangeLoyaltyTierLevelEffectProps) GetRuleTitleOk() (*string, bool)`

GetRuleTitleOk returns a tuple with the RuleTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleTitle

`func (o *ChangeLoyaltyTierLevelEffectProps) SetRuleTitle(v string)`

SetRuleTitle sets RuleTitle field to given value.


### GetProgramId

`func (o *ChangeLoyaltyTierLevelEffectProps) GetProgramId() int64`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *ChangeLoyaltyTierLevelEffectProps) GetProgramIdOk() (*int64, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *ChangeLoyaltyTierLevelEffectProps) SetProgramId(v int64)`

SetProgramId sets ProgramId field to given value.


### GetSubLedgerId

`func (o *ChangeLoyaltyTierLevelEffectProps) GetSubLedgerId() string`

GetSubLedgerId returns the SubLedgerId field if non-nil, zero value otherwise.

### GetSubLedgerIdOk

`func (o *ChangeLoyaltyTierLevelEffectProps) GetSubLedgerIdOk() (*string, bool)`

GetSubLedgerIdOk returns a tuple with the SubLedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubLedgerId

`func (o *ChangeLoyaltyTierLevelEffectProps) SetSubLedgerId(v string)`

SetSubLedgerId sets SubLedgerId field to given value.


### GetPreviousTierName

`func (o *ChangeLoyaltyTierLevelEffectProps) GetPreviousTierName() string`

GetPreviousTierName returns the PreviousTierName field if non-nil, zero value otherwise.

### GetPreviousTierNameOk

`func (o *ChangeLoyaltyTierLevelEffectProps) GetPreviousTierNameOk() (*string, bool)`

GetPreviousTierNameOk returns a tuple with the PreviousTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTierName

`func (o *ChangeLoyaltyTierLevelEffectProps) SetPreviousTierName(v string)`

SetPreviousTierName sets PreviousTierName field to given value.

### HasPreviousTierName

`func (o *ChangeLoyaltyTierLevelEffectProps) HasPreviousTierName() bool`

HasPreviousTierName returns a boolean if a field has been set.

### GetNewTierName

`func (o *ChangeLoyaltyTierLevelEffectProps) GetNewTierName() string`

GetNewTierName returns the NewTierName field if non-nil, zero value otherwise.

### GetNewTierNameOk

`func (o *ChangeLoyaltyTierLevelEffectProps) GetNewTierNameOk() (*string, bool)`

GetNewTierNameOk returns a tuple with the NewTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTierName

`func (o *ChangeLoyaltyTierLevelEffectProps) SetNewTierName(v string)`

SetNewTierName sets NewTierName field to given value.


### GetExpiryDate

`func (o *ChangeLoyaltyTierLevelEffectProps) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *ChangeLoyaltyTierLevelEffectProps) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *ChangeLoyaltyTierLevelEffectProps) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *ChangeLoyaltyTierLevelEffectProps) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


