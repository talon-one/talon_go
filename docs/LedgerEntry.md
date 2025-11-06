# LedgerEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**ProfileId** | Pointer to **string** | ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known &#x60;profileId&#x60;, we recommend you use a guest &#x60;profileId&#x60;.  | 
**AccountId** | Pointer to **int64** | The ID of the Talon.One account that owns this profile. | 
**LoyaltyProgramId** | Pointer to **int64** | ID of the ledger. | 
**EventId** | Pointer to **int64** | ID of the related event. | 
**Amount** | Pointer to **int64** | Amount of loyalty points. | 
**Reason** | Pointer to **string** | reason for awarding/deducting points. | 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiration date of the points. | 
**ReferenceId** | Pointer to **int64** | The ID of the balancing ledgerEntry. | [optional] 

## Methods

### NewLedgerEntry

`func NewLedgerEntry(id int64, created time.Time, profileId string, accountId int64, loyaltyProgramId int64, eventId int64, amount int64, reason string, expiryDate time.Time, ) *LedgerEntry`

NewLedgerEntry instantiates a new LedgerEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerEntryWithDefaults

`func NewLedgerEntryWithDefaults() *LedgerEntry`

NewLedgerEntryWithDefaults instantiates a new LedgerEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LedgerEntry) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LedgerEntry) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LedgerEntry) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *LedgerEntry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LedgerEntry) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LedgerEntry) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetProfileId

`func (o *LedgerEntry) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *LedgerEntry) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *LedgerEntry) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetAccountId

`func (o *LedgerEntry) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LedgerEntry) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LedgerEntry) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetLoyaltyProgramId

`func (o *LedgerEntry) GetLoyaltyProgramId() int64`

GetLoyaltyProgramId returns the LoyaltyProgramId field if non-nil, zero value otherwise.

### GetLoyaltyProgramIdOk

`func (o *LedgerEntry) GetLoyaltyProgramIdOk() (*int64, bool)`

GetLoyaltyProgramIdOk returns a tuple with the LoyaltyProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramId

`func (o *LedgerEntry) SetLoyaltyProgramId(v int64)`

SetLoyaltyProgramId sets LoyaltyProgramId field to given value.


### GetEventId

`func (o *LedgerEntry) GetEventId() int64`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *LedgerEntry) GetEventIdOk() (*int64, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *LedgerEntry) SetEventId(v int64)`

SetEventId sets EventId field to given value.


### GetAmount

`func (o *LedgerEntry) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LedgerEntry) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *LedgerEntry) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetReason

`func (o *LedgerEntry) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *LedgerEntry) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *LedgerEntry) SetReason(v string)`

SetReason sets Reason field to given value.


### GetExpiryDate

`func (o *LedgerEntry) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *LedgerEntry) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *LedgerEntry) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.


### GetReferenceId

`func (o *LedgerEntry) GetReferenceId() int64`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *LedgerEntry) GetReferenceIdOk() (*int64, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *LedgerEntry) SetReferenceId(v int64)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *LedgerEntry) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


