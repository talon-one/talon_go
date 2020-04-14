# LedgerEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was created. | 
**ProfileId** | Pointer to **string** | ID of the customers profile as used within this Talon.One account. May be omitted or set to the empty string if the customer does not yet have a known profile ID. | 
**AccountId** | Pointer to **int32** | The ID of the Talon.One account that owns this profile. | 
**LoyaltyProgramId** | Pointer to **int32** | ID of the ledger | 
**EventId** | Pointer to **int32** | ID of the related event | 
**Amount** | Pointer to **int32** | Amount of loyalty points | 
**Reason** | Pointer to **string** | reason for awarding/deducting points | 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiry date of the points | 
**ReferenceId** | Pointer to **int32** | The ID of the balancing ledgerEntry | [optional] 

## Methods

### GetId

`func (o *LedgerEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LedgerEntry) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *LedgerEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *LedgerEntry) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *LedgerEntry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LedgerEntry) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *LedgerEntry) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *LedgerEntry) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetProfileId

`func (o *LedgerEntry) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *LedgerEntry) GetProfileIdOk() (string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileId

`func (o *LedgerEntry) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileId

`func (o *LedgerEntry) SetProfileId(v string)`

SetProfileId gets a reference to the given string and assigns it to the ProfileId field.

### GetAccountId

`func (o *LedgerEntry) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LedgerEntry) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *LedgerEntry) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *LedgerEntry) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetLoyaltyProgramId

`func (o *LedgerEntry) GetLoyaltyProgramId() int32`

GetLoyaltyProgramId returns the LoyaltyProgramId field if non-nil, zero value otherwise.

### GetLoyaltyProgramIdOk

`func (o *LedgerEntry) GetLoyaltyProgramIdOk() (int32, bool)`

GetLoyaltyProgramIdOk returns a tuple with the LoyaltyProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoyaltyProgramId

`func (o *LedgerEntry) HasLoyaltyProgramId() bool`

HasLoyaltyProgramId returns a boolean if a field has been set.

### SetLoyaltyProgramId

`func (o *LedgerEntry) SetLoyaltyProgramId(v int32)`

SetLoyaltyProgramId gets a reference to the given int32 and assigns it to the LoyaltyProgramId field.

### GetEventId

`func (o *LedgerEntry) GetEventId() int32`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *LedgerEntry) GetEventIdOk() (int32, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventId

`func (o *LedgerEntry) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### SetEventId

`func (o *LedgerEntry) SetEventId(v int32)`

SetEventId gets a reference to the given int32 and assigns it to the EventId field.

### GetAmount

`func (o *LedgerEntry) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LedgerEntry) GetAmountOk() (int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAmount

`func (o *LedgerEntry) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmount

`func (o *LedgerEntry) SetAmount(v int32)`

SetAmount gets a reference to the given int32 and assigns it to the Amount field.

### GetReason

`func (o *LedgerEntry) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *LedgerEntry) GetReasonOk() (string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReason

`func (o *LedgerEntry) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReason

`func (o *LedgerEntry) SetReason(v string)`

SetReason gets a reference to the given string and assigns it to the Reason field.

### GetExpiryDate

`func (o *LedgerEntry) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *LedgerEntry) GetExpiryDateOk() (time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *LedgerEntry) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *LedgerEntry) SetExpiryDate(v time.Time)`

SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.

### GetReferenceId

`func (o *LedgerEntry) GetReferenceId() int32`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *LedgerEntry) GetReferenceIdOk() (int32, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferenceId

`func (o *LedgerEntry) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceId

`func (o *LedgerEntry) SetReferenceId(v int32)`

SetReferenceId gets a reference to the given int32 and assigns it to the ReferenceId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


