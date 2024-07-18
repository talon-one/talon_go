# LoyaltyCardAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | The alphanumeric identifier of the loyalty card.  | 
**UsersPerCardLimit** | Pointer to **int32** | The max amount of customer profiles that can be linked to the card. 0 means unlimited.  | 
**Profiles** | Pointer to [**[]LoyaltyCardProfileRegistration**](LoyaltyCardProfileRegistration.md) | Integration IDs of the customers profiles linked to the card. | [optional] 
**Ledger** | Pointer to [**LedgerInfo**](LedgerInfo.md) |  | [optional] 
**Subledgers** | Pointer to [**map[string]LedgerInfo**](LedgerInfo.md) | Displays point balances of the card in the subledgers of the loyalty program. | [optional] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent update of the loyalty card. | [optional] 
**OldCardIdentifier** | Pointer to **string** | The alphanumeric identifier of the loyalty card.  | [optional] 
**NewCardIdentifier** | Pointer to **string** | The alphanumeric identifier of the loyalty card.  | [optional] 

## Methods

### GetIdentifier

`func (o *LoyaltyCardAllOf) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *LoyaltyCardAllOf) GetIdentifierOk() (string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIdentifier

`func (o *LoyaltyCardAllOf) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifier

`func (o *LoyaltyCardAllOf) SetIdentifier(v string)`

SetIdentifier gets a reference to the given string and assigns it to the Identifier field.

### GetUsersPerCardLimit

`func (o *LoyaltyCardAllOf) GetUsersPerCardLimit() int32`

GetUsersPerCardLimit returns the UsersPerCardLimit field if non-nil, zero value otherwise.

### GetUsersPerCardLimitOk

`func (o *LoyaltyCardAllOf) GetUsersPerCardLimitOk() (int32, bool)`

GetUsersPerCardLimitOk returns a tuple with the UsersPerCardLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsersPerCardLimit

`func (o *LoyaltyCardAllOf) HasUsersPerCardLimit() bool`

HasUsersPerCardLimit returns a boolean if a field has been set.

### SetUsersPerCardLimit

`func (o *LoyaltyCardAllOf) SetUsersPerCardLimit(v int32)`

SetUsersPerCardLimit gets a reference to the given int32 and assigns it to the UsersPerCardLimit field.

### GetProfiles

`func (o *LoyaltyCardAllOf) GetProfiles() []LoyaltyCardProfileRegistration`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *LoyaltyCardAllOf) GetProfilesOk() ([]LoyaltyCardProfileRegistration, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfiles

`func (o *LoyaltyCardAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfiles

`func (o *LoyaltyCardAllOf) SetProfiles(v []LoyaltyCardProfileRegistration)`

SetProfiles gets a reference to the given []LoyaltyCardProfileRegistration and assigns it to the Profiles field.

### GetLedger

`func (o *LoyaltyCardAllOf) GetLedger() LedgerInfo`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *LoyaltyCardAllOf) GetLedgerOk() (LedgerInfo, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLedger

`func (o *LoyaltyCardAllOf) HasLedger() bool`

HasLedger returns a boolean if a field has been set.

### SetLedger

`func (o *LoyaltyCardAllOf) SetLedger(v LedgerInfo)`

SetLedger gets a reference to the given LedgerInfo and assigns it to the Ledger field.

### GetSubledgers

`func (o *LoyaltyCardAllOf) GetSubledgers() map[string]LedgerInfo`

GetSubledgers returns the Subledgers field if non-nil, zero value otherwise.

### GetSubledgersOk

`func (o *LoyaltyCardAllOf) GetSubledgersOk() (map[string]LedgerInfo, bool)`

GetSubledgersOk returns a tuple with the Subledgers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubledgers

`func (o *LoyaltyCardAllOf) HasSubledgers() bool`

HasSubledgers returns a boolean if a field has been set.

### SetSubledgers

`func (o *LoyaltyCardAllOf) SetSubledgers(v map[string]LedgerInfo)`

SetSubledgers gets a reference to the given map[string]LedgerInfo and assigns it to the Subledgers field.

### GetModified

`func (o *LoyaltyCardAllOf) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *LoyaltyCardAllOf) GetModifiedOk() (time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *LoyaltyCardAllOf) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *LoyaltyCardAllOf) SetModified(v time.Time)`

SetModified gets a reference to the given time.Time and assigns it to the Modified field.

### GetOldCardIdentifier

`func (o *LoyaltyCardAllOf) GetOldCardIdentifier() string`

GetOldCardIdentifier returns the OldCardIdentifier field if non-nil, zero value otherwise.

### GetOldCardIdentifierOk

`func (o *LoyaltyCardAllOf) GetOldCardIdentifierOk() (string, bool)`

GetOldCardIdentifierOk returns a tuple with the OldCardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOldCardIdentifier

`func (o *LoyaltyCardAllOf) HasOldCardIdentifier() bool`

HasOldCardIdentifier returns a boolean if a field has been set.

### SetOldCardIdentifier

`func (o *LoyaltyCardAllOf) SetOldCardIdentifier(v string)`

SetOldCardIdentifier gets a reference to the given string and assigns it to the OldCardIdentifier field.

### GetNewCardIdentifier

`func (o *LoyaltyCardAllOf) GetNewCardIdentifier() string`

GetNewCardIdentifier returns the NewCardIdentifier field if non-nil, zero value otherwise.

### GetNewCardIdentifierOk

`func (o *LoyaltyCardAllOf) GetNewCardIdentifierOk() (string, bool)`

GetNewCardIdentifierOk returns a tuple with the NewCardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNewCardIdentifier

`func (o *LoyaltyCardAllOf) HasNewCardIdentifier() bool`

HasNewCardIdentifier returns a boolean if a field has been set.

### SetNewCardIdentifier

`func (o *LoyaltyCardAllOf) SetNewCardIdentifier(v string)`

SetNewCardIdentifier gets a reference to the given string and assigns it to the NewCardIdentifier field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


