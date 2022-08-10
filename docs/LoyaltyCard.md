# LoyaltyCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was created. | 
**ProgramID** | Pointer to **int32** | The ID of the loyalty program that owns this entity. | 
**Status** | Pointer to **string** | Status of the loyalty card. Can be one of: [&#39;active&#39;, &#39;disabled&#39;]  | 
**Identifier** | Pointer to **string** | The alphanumeric identifier of the loyalty card. | 
**UsersPerCardLimit** | Pointer to **int32** | The max amount of user profiles a card can be shared with. 0 means unlimited.  | 
**Profiles** | Pointer to [**[]LoyaltyCardProfileRegistration**](LoyaltyCardProfileRegistration.md) | Integration IDs of the customers associated with the card. | [optional] 
**Ledger** | Pointer to [**LedgerInfo**](LedgerInfo.md) |  | [optional] 
**Subledgers** | Pointer to [**map[string]LedgerInfo**](LedgerInfo.md) | Displays point balances of the card in the subledgers of the loyalty program. | [optional] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent update of the loyalty card. | [optional] 

## Methods

### GetId

`func (o *LoyaltyCard) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoyaltyCard) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *LoyaltyCard) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *LoyaltyCard) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *LoyaltyCard) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LoyaltyCard) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *LoyaltyCard) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *LoyaltyCard) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetProgramID

`func (o *LoyaltyCard) GetProgramID() int32`

GetProgramID returns the ProgramID field if non-nil, zero value otherwise.

### GetProgramIDOk

`func (o *LoyaltyCard) GetProgramIDOk() (int32, bool)`

GetProgramIDOk returns a tuple with the ProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgramID

`func (o *LoyaltyCard) HasProgramID() bool`

HasProgramID returns a boolean if a field has been set.

### SetProgramID

`func (o *LoyaltyCard) SetProgramID(v int32)`

SetProgramID gets a reference to the given int32 and assigns it to the ProgramID field.

### GetStatus

`func (o *LoyaltyCard) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LoyaltyCard) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *LoyaltyCard) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *LoyaltyCard) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### GetIdentifier

`func (o *LoyaltyCard) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *LoyaltyCard) GetIdentifierOk() (string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIdentifier

`func (o *LoyaltyCard) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifier

`func (o *LoyaltyCard) SetIdentifier(v string)`

SetIdentifier gets a reference to the given string and assigns it to the Identifier field.

### GetUsersPerCardLimit

`func (o *LoyaltyCard) GetUsersPerCardLimit() int32`

GetUsersPerCardLimit returns the UsersPerCardLimit field if non-nil, zero value otherwise.

### GetUsersPerCardLimitOk

`func (o *LoyaltyCard) GetUsersPerCardLimitOk() (int32, bool)`

GetUsersPerCardLimitOk returns a tuple with the UsersPerCardLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsersPerCardLimit

`func (o *LoyaltyCard) HasUsersPerCardLimit() bool`

HasUsersPerCardLimit returns a boolean if a field has been set.

### SetUsersPerCardLimit

`func (o *LoyaltyCard) SetUsersPerCardLimit(v int32)`

SetUsersPerCardLimit gets a reference to the given int32 and assigns it to the UsersPerCardLimit field.

### GetProfiles

`func (o *LoyaltyCard) GetProfiles() []LoyaltyCardProfileRegistration`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *LoyaltyCard) GetProfilesOk() ([]LoyaltyCardProfileRegistration, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfiles

`func (o *LoyaltyCard) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfiles

`func (o *LoyaltyCard) SetProfiles(v []LoyaltyCardProfileRegistration)`

SetProfiles gets a reference to the given []LoyaltyCardProfileRegistration and assigns it to the Profiles field.

### GetLedger

`func (o *LoyaltyCard) GetLedger() LedgerInfo`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *LoyaltyCard) GetLedgerOk() (LedgerInfo, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLedger

`func (o *LoyaltyCard) HasLedger() bool`

HasLedger returns a boolean if a field has been set.

### SetLedger

`func (o *LoyaltyCard) SetLedger(v LedgerInfo)`

SetLedger gets a reference to the given LedgerInfo and assigns it to the Ledger field.

### GetSubledgers

`func (o *LoyaltyCard) GetSubledgers() map[string]LedgerInfo`

GetSubledgers returns the Subledgers field if non-nil, zero value otherwise.

### GetSubledgersOk

`func (o *LoyaltyCard) GetSubledgersOk() (map[string]LedgerInfo, bool)`

GetSubledgersOk returns a tuple with the Subledgers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubledgers

`func (o *LoyaltyCard) HasSubledgers() bool`

HasSubledgers returns a boolean if a field has been set.

### SetSubledgers

`func (o *LoyaltyCard) SetSubledgers(v map[string]LedgerInfo)`

SetSubledgers gets a reference to the given map[string]LedgerInfo and assigns it to the Subledgers field.

### GetModified

`func (o *LoyaltyCard) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *LoyaltyCard) GetModifiedOk() (time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *LoyaltyCard) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *LoyaltyCard) SetModified(v time.Time)`

SetModified gets a reference to the given time.Time and assigns it to the Modified field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


