# ExtendLoyaltyPointsExpiryDateEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramId** | Pointer to **int32** | ID of the loyalty program that contains these points. | 
**SubLedgerId** | Pointer to **string** | API name of the loyalty program subledger that contains these points. added. | 
**ExtensionDuration** | Pointer to **string** | Time frame by which the expiry date extends.  The time format is either: - immediate, or - an **integer** followed by a letter indicating the time unit.  Examples: &#x60;immediate&#x60;, &#x60;30s&#x60;, &#x60;40m&#x60;, &#x60;1h&#x60;, &#x60;5D&#x60;, &#x60;7W&#x60;, &#x60;10M&#x60;, &#x60;15Y&#x60;.  Available units:  - &#x60;s&#x60;: seconds - &#x60;m&#x60;: minutes - &#x60;h&#x60;: hours - &#x60;D&#x60;: days - &#x60;W&#x60;: weeks - &#x60;M&#x60;: months - &#x60;Y&#x60;: years  You can round certain units up or down: - &#x60;_D&#x60; for rounding down days only. Signifies the start of the day. - &#x60;_U&#x60; for rounding up days, weeks, months and years. Signifies the end of the day, week, month or year.  | 
**TransactionUUIDs** | Pointer to **[]string** | The list of identifiers of transactions affected affected by the extension. | [optional] 
**PreviousExpirationDate** | Pointer to [**time.Time**](time.Time.md) | Expiry date before applying the extension. | 

## Methods

### GetProgramId

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) GetProgramId() int32`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) GetProgramIdOk() (int32, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgramId

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### SetProgramId

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) SetProgramId(v int32)`

SetProgramId gets a reference to the given int32 and assigns it to the ProgramId field.

### GetSubLedgerId

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) GetSubLedgerId() string`

GetSubLedgerId returns the SubLedgerId field if non-nil, zero value otherwise.

### GetSubLedgerIdOk

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) GetSubLedgerIdOk() (string, bool)`

GetSubLedgerIdOk returns a tuple with the SubLedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubLedgerId

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) HasSubLedgerId() bool`

HasSubLedgerId returns a boolean if a field has been set.

### SetSubLedgerId

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) SetSubLedgerId(v string)`

SetSubLedgerId gets a reference to the given string and assigns it to the SubLedgerId field.

### GetExtensionDuration

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) GetExtensionDuration() string`

GetExtensionDuration returns the ExtensionDuration field if non-nil, zero value otherwise.

### GetExtensionDurationOk

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) GetExtensionDurationOk() (string, bool)`

GetExtensionDurationOk returns a tuple with the ExtensionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExtensionDuration

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) HasExtensionDuration() bool`

HasExtensionDuration returns a boolean if a field has been set.

### SetExtensionDuration

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) SetExtensionDuration(v string)`

SetExtensionDuration gets a reference to the given string and assigns it to the ExtensionDuration field.

### GetTransactionUUIDs

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) GetTransactionUUIDs() []string`

GetTransactionUUIDs returns the TransactionUUIDs field if non-nil, zero value otherwise.

### GetTransactionUUIDsOk

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) GetTransactionUUIDsOk() ([]string, bool)`

GetTransactionUUIDsOk returns a tuple with the TransactionUUIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTransactionUUIDs

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) HasTransactionUUIDs() bool`

HasTransactionUUIDs returns a boolean if a field has been set.

### SetTransactionUUIDs

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) SetTransactionUUIDs(v []string)`

SetTransactionUUIDs gets a reference to the given []string and assigns it to the TransactionUUIDs field.

### GetPreviousExpirationDate

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) GetPreviousExpirationDate() time.Time`

GetPreviousExpirationDate returns the PreviousExpirationDate field if non-nil, zero value otherwise.

### GetPreviousExpirationDateOk

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) GetPreviousExpirationDateOk() (time.Time, bool)`

GetPreviousExpirationDateOk returns a tuple with the PreviousExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreviousExpirationDate

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) HasPreviousExpirationDate() bool`

HasPreviousExpirationDate returns a boolean if a field has been set.

### SetPreviousExpirationDate

`func (o *ExtendLoyaltyPointsExpiryDateEffectProps) SetPreviousExpirationDate(v time.Time)`

SetPreviousExpirationDate gets a reference to the given time.Time and assigns it to the PreviousExpirationDate field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


