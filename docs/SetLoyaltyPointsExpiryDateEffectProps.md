# SetLoyaltyPointsExpiryDateEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramId** | Pointer to **int64** | ID of the loyalty program that contains these points. | 
**SubLedgerId** | Pointer to **string** | API name of the loyalty program subledger that contains these points. | 
**NewExpiryDate** | Pointer to [**time.Time**](time.Time.md) | The specified expiry date and time for all active and pending point transactions in the loyalty program subledger. | 
**AffectedTransactions** | Pointer to [**[]LoyaltyLedgerEntryExpiryDateChange**](LoyaltyLedgerEntryExpiryDateChange.md) | List of transactions affected by the expiry date update. | [optional] 

## Methods

### NewSetLoyaltyPointsExpiryDateEffectProps

`func NewSetLoyaltyPointsExpiryDateEffectProps(programId int64, subLedgerId string, newExpiryDate time.Time, ) *SetLoyaltyPointsExpiryDateEffectProps`

NewSetLoyaltyPointsExpiryDateEffectProps instantiates a new SetLoyaltyPointsExpiryDateEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetLoyaltyPointsExpiryDateEffectPropsWithDefaults

`func NewSetLoyaltyPointsExpiryDateEffectPropsWithDefaults() *SetLoyaltyPointsExpiryDateEffectProps`

NewSetLoyaltyPointsExpiryDateEffectPropsWithDefaults instantiates a new SetLoyaltyPointsExpiryDateEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgramId

`func (o *SetLoyaltyPointsExpiryDateEffectProps) GetProgramId() int64`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *SetLoyaltyPointsExpiryDateEffectProps) GetProgramIdOk() (*int64, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *SetLoyaltyPointsExpiryDateEffectProps) SetProgramId(v int64)`

SetProgramId sets ProgramId field to given value.


### GetSubLedgerId

`func (o *SetLoyaltyPointsExpiryDateEffectProps) GetSubLedgerId() string`

GetSubLedgerId returns the SubLedgerId field if non-nil, zero value otherwise.

### GetSubLedgerIdOk

`func (o *SetLoyaltyPointsExpiryDateEffectProps) GetSubLedgerIdOk() (*string, bool)`

GetSubLedgerIdOk returns a tuple with the SubLedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubLedgerId

`func (o *SetLoyaltyPointsExpiryDateEffectProps) SetSubLedgerId(v string)`

SetSubLedgerId sets SubLedgerId field to given value.


### GetNewExpiryDate

`func (o *SetLoyaltyPointsExpiryDateEffectProps) GetNewExpiryDate() time.Time`

GetNewExpiryDate returns the NewExpiryDate field if non-nil, zero value otherwise.

### GetNewExpiryDateOk

`func (o *SetLoyaltyPointsExpiryDateEffectProps) GetNewExpiryDateOk() (*time.Time, bool)`

GetNewExpiryDateOk returns a tuple with the NewExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewExpiryDate

`func (o *SetLoyaltyPointsExpiryDateEffectProps) SetNewExpiryDate(v time.Time)`

SetNewExpiryDate sets NewExpiryDate field to given value.


### GetAffectedTransactions

`func (o *SetLoyaltyPointsExpiryDateEffectProps) GetAffectedTransactions() []LoyaltyLedgerEntryExpiryDateChange`

GetAffectedTransactions returns the AffectedTransactions field if non-nil, zero value otherwise.

### GetAffectedTransactionsOk

`func (o *SetLoyaltyPointsExpiryDateEffectProps) GetAffectedTransactionsOk() (*[]LoyaltyLedgerEntryExpiryDateChange, bool)`

GetAffectedTransactionsOk returns a tuple with the AffectedTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedTransactions

`func (o *SetLoyaltyPointsExpiryDateEffectProps) SetAffectedTransactions(v []LoyaltyLedgerEntryExpiryDateChange)`

SetAffectedTransactions sets AffectedTransactions field to given value.

### HasAffectedTransactions

`func (o *SetLoyaltyPointsExpiryDateEffectProps) HasAffectedTransactions() bool`

HasAffectedTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


