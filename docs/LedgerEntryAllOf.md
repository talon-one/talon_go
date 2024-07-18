# LedgerEntryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int32** | The ID of the Talon.One account that owns this profile. | 
**LoyaltyProgramId** | Pointer to **int32** | ID of the ledger. | 
**EventId** | Pointer to **int32** | ID of the related event. | 
**Amount** | Pointer to **int32** | Amount of loyalty points. | 
**Reason** | Pointer to **string** | reason for awarding/deducting points. | 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiration date of the points. | 
**ReferenceId** | Pointer to **int32** | The ID of the balancing ledgerEntry. | [optional] 

## Methods

### GetAccountId

`func (o *LedgerEntryAllOf) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LedgerEntryAllOf) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *LedgerEntryAllOf) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *LedgerEntryAllOf) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetLoyaltyProgramId

`func (o *LedgerEntryAllOf) GetLoyaltyProgramId() int32`

GetLoyaltyProgramId returns the LoyaltyProgramId field if non-nil, zero value otherwise.

### GetLoyaltyProgramIdOk

`func (o *LedgerEntryAllOf) GetLoyaltyProgramIdOk() (int32, bool)`

GetLoyaltyProgramIdOk returns a tuple with the LoyaltyProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoyaltyProgramId

`func (o *LedgerEntryAllOf) HasLoyaltyProgramId() bool`

HasLoyaltyProgramId returns a boolean if a field has been set.

### SetLoyaltyProgramId

`func (o *LedgerEntryAllOf) SetLoyaltyProgramId(v int32)`

SetLoyaltyProgramId gets a reference to the given int32 and assigns it to the LoyaltyProgramId field.

### GetEventId

`func (o *LedgerEntryAllOf) GetEventId() int32`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *LedgerEntryAllOf) GetEventIdOk() (int32, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventId

`func (o *LedgerEntryAllOf) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### SetEventId

`func (o *LedgerEntryAllOf) SetEventId(v int32)`

SetEventId gets a reference to the given int32 and assigns it to the EventId field.

### GetAmount

`func (o *LedgerEntryAllOf) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LedgerEntryAllOf) GetAmountOk() (int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAmount

`func (o *LedgerEntryAllOf) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmount

`func (o *LedgerEntryAllOf) SetAmount(v int32)`

SetAmount gets a reference to the given int32 and assigns it to the Amount field.

### GetReason

`func (o *LedgerEntryAllOf) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *LedgerEntryAllOf) GetReasonOk() (string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReason

`func (o *LedgerEntryAllOf) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReason

`func (o *LedgerEntryAllOf) SetReason(v string)`

SetReason gets a reference to the given string and assigns it to the Reason field.

### GetExpiryDate

`func (o *LedgerEntryAllOf) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *LedgerEntryAllOf) GetExpiryDateOk() (time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *LedgerEntryAllOf) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *LedgerEntryAllOf) SetExpiryDate(v time.Time)`

SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.

### GetReferenceId

`func (o *LedgerEntryAllOf) GetReferenceId() int32`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *LedgerEntryAllOf) GetReferenceIdOk() (int32, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferenceId

`func (o *LedgerEntryAllOf) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceId

`func (o *LedgerEntryAllOf) SetReferenceId(v int32)`

SetReferenceId gets a reference to the given int32 and assigns it to the ReferenceId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


