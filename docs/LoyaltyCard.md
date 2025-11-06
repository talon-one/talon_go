# LoyaltyCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**ProgramID** | Pointer to **int64** | The ID of the loyalty program that owns this entity. | 
**ProgramName** | Pointer to **string** | The integration name of the loyalty program that owns this entity. | [optional] 
**ProgramTitle** | Pointer to **string** | The Campaign Manager-displayed name of the loyalty program that owns this entity. | [optional] 
**Status** | Pointer to **string** | Status of the loyalty card. Can be &#x60;active&#x60; or &#x60;inactive&#x60;.  | 
**BlockReason** | Pointer to **string** | Reason for transferring and blocking the loyalty card.  | [optional] 
**Identifier** | Pointer to **string** | The alphanumeric identifier of the loyalty card.  | 
**UsersPerCardLimit** | Pointer to **int64** | The max amount of customer profiles that can be linked to the card. 0 means unlimited.  | 
**Profiles** | Pointer to [**[]LoyaltyCardProfileRegistration**](LoyaltyCardProfileRegistration.md) | Integration IDs of the customers profiles linked to the card. | [optional] 
**Ledger** | Pointer to [**LedgerInfo**](LedgerInfo.md) |  | [optional] 
**Subledgers** | Pointer to [**map[string]LedgerInfo**](LedgerInfo.md) | Displays point balances of the card in the subledgers of the loyalty program. | [optional] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent update of the loyalty card. | [optional] 
**OldCardIdentifier** | Pointer to **string** | The alphanumeric identifier of the loyalty card.  | [optional] 
**NewCardIdentifier** | Pointer to **string** | The alphanumeric identifier of the loyalty card.  | [optional] 
**BatchId** | Pointer to **string** | The ID of the batch in which the loyalty card was created. | [optional] 

## Methods

### NewLoyaltyCard

`func NewLoyaltyCard(id int64, created time.Time, programID int64, status string, identifier string, usersPerCardLimit int64, ) *LoyaltyCard`

NewLoyaltyCard instantiates a new LoyaltyCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyCardWithDefaults

`func NewLoyaltyCardWithDefaults() *LoyaltyCard`

NewLoyaltyCardWithDefaults instantiates a new LoyaltyCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoyaltyCard) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoyaltyCard) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoyaltyCard) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *LoyaltyCard) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LoyaltyCard) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LoyaltyCard) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetProgramID

`func (o *LoyaltyCard) GetProgramID() int64`

GetProgramID returns the ProgramID field if non-nil, zero value otherwise.

### GetProgramIDOk

`func (o *LoyaltyCard) GetProgramIDOk() (*int64, bool)`

GetProgramIDOk returns a tuple with the ProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramID

`func (o *LoyaltyCard) SetProgramID(v int64)`

SetProgramID sets ProgramID field to given value.


### GetProgramName

`func (o *LoyaltyCard) GetProgramName() string`

GetProgramName returns the ProgramName field if non-nil, zero value otherwise.

### GetProgramNameOk

`func (o *LoyaltyCard) GetProgramNameOk() (*string, bool)`

GetProgramNameOk returns a tuple with the ProgramName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramName

`func (o *LoyaltyCard) SetProgramName(v string)`

SetProgramName sets ProgramName field to given value.

### HasProgramName

`func (o *LoyaltyCard) HasProgramName() bool`

HasProgramName returns a boolean if a field has been set.

### GetProgramTitle

`func (o *LoyaltyCard) GetProgramTitle() string`

GetProgramTitle returns the ProgramTitle field if non-nil, zero value otherwise.

### GetProgramTitleOk

`func (o *LoyaltyCard) GetProgramTitleOk() (*string, bool)`

GetProgramTitleOk returns a tuple with the ProgramTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramTitle

`func (o *LoyaltyCard) SetProgramTitle(v string)`

SetProgramTitle sets ProgramTitle field to given value.

### HasProgramTitle

`func (o *LoyaltyCard) HasProgramTitle() bool`

HasProgramTitle returns a boolean if a field has been set.

### GetStatus

`func (o *LoyaltyCard) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LoyaltyCard) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LoyaltyCard) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetBlockReason

`func (o *LoyaltyCard) GetBlockReason() string`

GetBlockReason returns the BlockReason field if non-nil, zero value otherwise.

### GetBlockReasonOk

`func (o *LoyaltyCard) GetBlockReasonOk() (*string, bool)`

GetBlockReasonOk returns a tuple with the BlockReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockReason

`func (o *LoyaltyCard) SetBlockReason(v string)`

SetBlockReason sets BlockReason field to given value.

### HasBlockReason

`func (o *LoyaltyCard) HasBlockReason() bool`

HasBlockReason returns a boolean if a field has been set.

### GetIdentifier

`func (o *LoyaltyCard) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *LoyaltyCard) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *LoyaltyCard) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetUsersPerCardLimit

`func (o *LoyaltyCard) GetUsersPerCardLimit() int64`

GetUsersPerCardLimit returns the UsersPerCardLimit field if non-nil, zero value otherwise.

### GetUsersPerCardLimitOk

`func (o *LoyaltyCard) GetUsersPerCardLimitOk() (*int64, bool)`

GetUsersPerCardLimitOk returns a tuple with the UsersPerCardLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersPerCardLimit

`func (o *LoyaltyCard) SetUsersPerCardLimit(v int64)`

SetUsersPerCardLimit sets UsersPerCardLimit field to given value.


### GetProfiles

`func (o *LoyaltyCard) GetProfiles() []LoyaltyCardProfileRegistration`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *LoyaltyCard) GetProfilesOk() (*[]LoyaltyCardProfileRegistration, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *LoyaltyCard) SetProfiles(v []LoyaltyCardProfileRegistration)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *LoyaltyCard) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetLedger

`func (o *LoyaltyCard) GetLedger() LedgerInfo`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *LoyaltyCard) GetLedgerOk() (*LedgerInfo, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *LoyaltyCard) SetLedger(v LedgerInfo)`

SetLedger sets Ledger field to given value.

### HasLedger

`func (o *LoyaltyCard) HasLedger() bool`

HasLedger returns a boolean if a field has been set.

### GetSubledgers

`func (o *LoyaltyCard) GetSubledgers() map[string]LedgerInfo`

GetSubledgers returns the Subledgers field if non-nil, zero value otherwise.

### GetSubledgersOk

`func (o *LoyaltyCard) GetSubledgersOk() (*map[string]LedgerInfo, bool)`

GetSubledgersOk returns a tuple with the Subledgers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgers

`func (o *LoyaltyCard) SetSubledgers(v map[string]LedgerInfo)`

SetSubledgers sets Subledgers field to given value.

### HasSubledgers

`func (o *LoyaltyCard) HasSubledgers() bool`

HasSubledgers returns a boolean if a field has been set.

### GetModified

`func (o *LoyaltyCard) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *LoyaltyCard) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *LoyaltyCard) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *LoyaltyCard) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetOldCardIdentifier

`func (o *LoyaltyCard) GetOldCardIdentifier() string`

GetOldCardIdentifier returns the OldCardIdentifier field if non-nil, zero value otherwise.

### GetOldCardIdentifierOk

`func (o *LoyaltyCard) GetOldCardIdentifierOk() (*string, bool)`

GetOldCardIdentifierOk returns a tuple with the OldCardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldCardIdentifier

`func (o *LoyaltyCard) SetOldCardIdentifier(v string)`

SetOldCardIdentifier sets OldCardIdentifier field to given value.

### HasOldCardIdentifier

`func (o *LoyaltyCard) HasOldCardIdentifier() bool`

HasOldCardIdentifier returns a boolean if a field has been set.

### GetNewCardIdentifier

`func (o *LoyaltyCard) GetNewCardIdentifier() string`

GetNewCardIdentifier returns the NewCardIdentifier field if non-nil, zero value otherwise.

### GetNewCardIdentifierOk

`func (o *LoyaltyCard) GetNewCardIdentifierOk() (*string, bool)`

GetNewCardIdentifierOk returns a tuple with the NewCardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCardIdentifier

`func (o *LoyaltyCard) SetNewCardIdentifier(v string)`

SetNewCardIdentifier sets NewCardIdentifier field to given value.

### HasNewCardIdentifier

`func (o *LoyaltyCard) HasNewCardIdentifier() bool`

HasNewCardIdentifier returns a boolean if a field has been set.

### GetBatchId

`func (o *LoyaltyCard) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *LoyaltyCard) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *LoyaltyCard) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *LoyaltyCard) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


