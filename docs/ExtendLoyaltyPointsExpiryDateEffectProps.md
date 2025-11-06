# ExtendLoyaltyPointsExpiryDateEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramId** | Pointer to **int64** | ID of the loyalty program that contains these points. | 
**SubLedgerId** | Pointer to **string** | API name of the loyalty program subledger that contains these points. added. | 
**ExtensionDuration** | Pointer to **string** | Time frame by which the expiry date extends.  The time format is either: - immediate, or - an **integer** followed by a letter indicating the time unit.  Examples: &#x60;immediate&#x60;, &#x60;30s&#x60;, &#x60;40m&#x60;, &#x60;1h&#x60;, &#x60;5D&#x60;, &#x60;7W&#x60;, &#x60;10M&#x60;, &#x60;15Y&#x60;.  Available units:  - &#x60;s&#x60;: seconds - &#x60;m&#x60;: minutes - &#x60;h&#x60;: hours - &#x60;D&#x60;: days - &#x60;W&#x60;: weeks - &#x60;M&#x60;: months - &#x60;Y&#x60;: years  You can round certain units up or down: - &#x60;_D&#x60; for rounding down days only. Signifies the start of the day. - &#x60;_U&#x60; for rounding up days, weeks, months and years. Signifies the end of the day, week, month or year.  | 
**AffectedTransactions** | Pointer to [**[]LoyaltyLedgerEntryExpiryDateChange**](LoyaltyLedgerEntryExpiryDateChange.md) | List of transactions affected by the expiry date update. | [optional] 

## Methods

### NewExtendLoyaltyPointsExpiryDateEffectProps

`func NewExtendLoyaltyPointsExpiryDateEffectProps(programId int64, subLedgerId string, extensionDuration string, ) *ExtendLoyaltyPointsExpiryDateEffectProps`

NewExtendLoyaltyPointsExpiryDateEffectProps instantiates a new ExtendLoyaltyPointsExpiryDateEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendLoyaltyPointsExpiryDateEffectPropsWithDefaults

`func NewExtendLoyaltyPointsExpiryDateEffectPropsWithDefaults() *ExtendLoyaltyPointsExpiryDateEffectProps`

NewExtendLoyaltyPointsExpiryDateEffectPropsWithDefaults instantiates a new ExtendLoyaltyPointsExpiryDateEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgramId

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) GetProgramId() int64`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) GetProgramIdOk() (*int64, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) SetProgramId(v int64)`

SetProgramId sets ProgramId field to given value.


### GetSubLedgerId

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) GetSubLedgerId() string`

GetSubLedgerId returns the SubLedgerId field if non-nil, zero value otherwise.

### GetSubLedgerIdOk

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) GetSubLedgerIdOk() (*string, bool)`

GetSubLedgerIdOk returns a tuple with the SubLedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubLedgerId

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) SetSubLedgerId(v string)`

SetSubLedgerId sets SubLedgerId field to given value.


### GetExtensionDuration

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) GetExtensionDuration() string`

GetExtensionDuration returns the ExtensionDuration field if non-nil, zero value otherwise.

### GetExtensionDurationOk

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) GetExtensionDurationOk() (*string, bool)`

GetExtensionDurationOk returns a tuple with the ExtensionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionDuration

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) SetExtensionDuration(v string)`

SetExtensionDuration sets ExtensionDuration field to given value.


### GetAffectedTransactions

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) GetAffectedTransactions() []LoyaltyLedgerEntryExpiryDateChange`

GetAffectedTransactions returns the AffectedTransactions field if non-nil, zero value otherwise.

### GetAffectedTransactionsOk

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) GetAffectedTransactionsOk() (*[]LoyaltyLedgerEntryExpiryDateChange, bool)`

GetAffectedTransactionsOk returns a tuple with the AffectedTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedTransactions

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) SetAffectedTransactions(v []LoyaltyLedgerEntryExpiryDateChange)`

SetAffectedTransactions sets AffectedTransactions field to given value.

### HasAffectedTransactions

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) HasAffectedTransactions() bool`

HasAffectedTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


