# LoyaltyLedgerEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to [**time.Time**](time.Time.md) |  | 
**ProgramID** | Pointer to **int32** |  | 
**CustomerProfileID** | Pointer to **string** |  | 
**CustomerSessionID** | Pointer to **string** |  | [optional] 
**EventID** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | 
**Amount** | Pointer to **float32** |  | 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Name** | Pointer to **string** | A name referencing the condition or effect that added this entry, or the specific name provided in an API call. | 
**SubLedgerID** | Pointer to **string** | This specifies if we are adding loyalty points to the main ledger or a subledger | 

## Methods

### GetCreated

`func (o *LoyaltyLedgerEntry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LoyaltyLedgerEntry) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *LoyaltyLedgerEntry) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *LoyaltyLedgerEntry) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetProgramID

`func (o *LoyaltyLedgerEntry) GetProgramID() int32`

GetProgramID returns the ProgramID field if non-nil, zero value otherwise.

### GetProgramIDOk

`func (o *LoyaltyLedgerEntry) GetProgramIDOk() (int32, bool)`

GetProgramIDOk returns a tuple with the ProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgramID

`func (o *LoyaltyLedgerEntry) HasProgramID() bool`

HasProgramID returns a boolean if a field has been set.

### SetProgramID

`func (o *LoyaltyLedgerEntry) SetProgramID(v int32)`

SetProgramID gets a reference to the given int32 and assigns it to the ProgramID field.

### GetCustomerProfileID

`func (o *LoyaltyLedgerEntry) GetCustomerProfileID() string`

GetCustomerProfileID returns the CustomerProfileID field if non-nil, zero value otherwise.

### GetCustomerProfileIDOk

`func (o *LoyaltyLedgerEntry) GetCustomerProfileIDOk() (string, bool)`

GetCustomerProfileIDOk returns a tuple with the CustomerProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerProfileID

`func (o *LoyaltyLedgerEntry) HasCustomerProfileID() bool`

HasCustomerProfileID returns a boolean if a field has been set.

### SetCustomerProfileID

`func (o *LoyaltyLedgerEntry) SetCustomerProfileID(v string)`

SetCustomerProfileID gets a reference to the given string and assigns it to the CustomerProfileID field.

### GetCustomerSessionID

`func (o *LoyaltyLedgerEntry) GetCustomerSessionID() string`

GetCustomerSessionID returns the CustomerSessionID field if non-nil, zero value otherwise.

### GetCustomerSessionIDOk

`func (o *LoyaltyLedgerEntry) GetCustomerSessionIDOk() (string, bool)`

GetCustomerSessionIDOk returns a tuple with the CustomerSessionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerSessionID

`func (o *LoyaltyLedgerEntry) HasCustomerSessionID() bool`

HasCustomerSessionID returns a boolean if a field has been set.

### SetCustomerSessionID

`func (o *LoyaltyLedgerEntry) SetCustomerSessionID(v string)`

SetCustomerSessionID gets a reference to the given string and assigns it to the CustomerSessionID field.

### GetEventID

`func (o *LoyaltyLedgerEntry) GetEventID() int32`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *LoyaltyLedgerEntry) GetEventIDOk() (int32, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventID

`func (o *LoyaltyLedgerEntry) HasEventID() bool`

HasEventID returns a boolean if a field has been set.

### SetEventID

`func (o *LoyaltyLedgerEntry) SetEventID(v int32)`

SetEventID gets a reference to the given int32 and assigns it to the EventID field.

### GetType

`func (o *LoyaltyLedgerEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoyaltyLedgerEntry) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *LoyaltyLedgerEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *LoyaltyLedgerEntry) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetAmount

`func (o *LoyaltyLedgerEntry) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LoyaltyLedgerEntry) GetAmountOk() (float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAmount

`func (o *LoyaltyLedgerEntry) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmount

`func (o *LoyaltyLedgerEntry) SetAmount(v float32)`

SetAmount gets a reference to the given float32 and assigns it to the Amount field.

### GetExpiryDate

`func (o *LoyaltyLedgerEntry) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *LoyaltyLedgerEntry) GetExpiryDateOk() (time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *LoyaltyLedgerEntry) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *LoyaltyLedgerEntry) SetExpiryDate(v time.Time)`

SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.

### GetName

`func (o *LoyaltyLedgerEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoyaltyLedgerEntry) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *LoyaltyLedgerEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *LoyaltyLedgerEntry) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetSubLedgerID

`func (o *LoyaltyLedgerEntry) GetSubLedgerID() string`

GetSubLedgerID returns the SubLedgerID field if non-nil, zero value otherwise.

### GetSubLedgerIDOk

`func (o *LoyaltyLedgerEntry) GetSubLedgerIDOk() (string, bool)`

GetSubLedgerIDOk returns a tuple with the SubLedgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubLedgerID

`func (o *LoyaltyLedgerEntry) HasSubLedgerID() bool`

HasSubLedgerID returns a boolean if a field has been set.

### SetSubLedgerID

`func (o *LoyaltyLedgerEntry) SetSubLedgerID(v string)`

SetSubLedgerID gets a reference to the given string and assigns it to the SubLedgerID field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


