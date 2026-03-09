# LoyaltyLedgerEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to [**time.Time**](time.Time.md) |  | 
**ProgramID** | Pointer to **int64** |  | 
**CustomerProfileID** | Pointer to **string** |  | [optional] 
**CardID** | Pointer to **int64** |  | [optional] 
**CustomerSessionID** | Pointer to **string** |  | [optional] 
**EventID** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** | The type of the ledger transaction. Possible values are: - &#x60;addition&#x60; - &#x60;subtraction&#x60; - &#x60;expire&#x60; - &#x60;expiring&#x60; (for expiring points ledgers)  | 
**Amount** | Pointer to **float32** |  | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Name** | Pointer to **string** | A name referencing the condition or effect that added this entry, or the specific name provided in an API call. | 
**SubLedgerID** | Pointer to **string** | This specifies if we are adding loyalty points to the main ledger or a subledger. | 
**UserID** | Pointer to **int64** | This is the ID of the user who created this entry, if the addition or subtraction was done manually. | [optional] 
**Archived** | Pointer to **bool** | Indicates if the entry belongs to the archived session. | [optional] 
**Flags** | Pointer to [**LoyaltyLedgerEntryFlags**](LoyaltyLedgerEntryFlags.md) |  | [optional] 
**ValidityDuration** | Pointer to **string** | The duration for which the points remain active, relative to the  activation date.  **Note**: This only applies to points for which &#x60;awaitsActivation&#x60; is &#x60;true&#x60; and &#x60;expiryDate&#x60; is not set.  | [optional] 

## Methods

### NewLoyaltyLedgerEntry

`func NewLoyaltyLedgerEntry(created time.Time, programID int64, type_ string, amount float32, name string, subLedgerID string, ) *LoyaltyLedgerEntry`

NewLoyaltyLedgerEntry instantiates a new LoyaltyLedgerEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyLedgerEntryWithDefaults

`func NewLoyaltyLedgerEntryWithDefaults() *LoyaltyLedgerEntry`

NewLoyaltyLedgerEntryWithDefaults instantiates a new LoyaltyLedgerEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *LoyaltyLedgerEntry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LoyaltyLedgerEntry) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LoyaltyLedgerEntry) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetProgramID

`func (o *LoyaltyLedgerEntry) GetProgramID() int64`

GetProgramID returns the ProgramID field if non-nil, zero value otherwise.

### GetProgramIDOk

`func (o *LoyaltyLedgerEntry) GetProgramIDOk() (*int64, bool)`

GetProgramIDOk returns a tuple with the ProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramID

`func (o *LoyaltyLedgerEntry) SetProgramID(v int64)`

SetProgramID sets ProgramID field to given value.


### GetCustomerProfileID

`func (o *LoyaltyLedgerEntry) GetCustomerProfileID() string`

GetCustomerProfileID returns the CustomerProfileID field if non-nil, zero value otherwise.

### GetCustomerProfileIDOk

`func (o *LoyaltyLedgerEntry) GetCustomerProfileIDOk() (*string, bool)`

GetCustomerProfileIDOk returns a tuple with the CustomerProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProfileID

`func (o *LoyaltyLedgerEntry) SetCustomerProfileID(v string)`

SetCustomerProfileID sets CustomerProfileID field to given value.

### HasCustomerProfileID

`func (o *LoyaltyLedgerEntry) HasCustomerProfileID() bool`

HasCustomerProfileID returns a boolean if a field has been set.

### GetCardID

`func (o *LoyaltyLedgerEntry) GetCardID() int64`

GetCardID returns the CardID field if non-nil, zero value otherwise.

### GetCardIDOk

`func (o *LoyaltyLedgerEntry) GetCardIDOk() (*int64, bool)`

GetCardIDOk returns a tuple with the CardID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardID

`func (o *LoyaltyLedgerEntry) SetCardID(v int64)`

SetCardID sets CardID field to given value.

### HasCardID

`func (o *LoyaltyLedgerEntry) HasCardID() bool`

HasCardID returns a boolean if a field has been set.

### GetCustomerSessionID

`func (o *LoyaltyLedgerEntry) GetCustomerSessionID() string`

GetCustomerSessionID returns the CustomerSessionID field if non-nil, zero value otherwise.

### GetCustomerSessionIDOk

`func (o *LoyaltyLedgerEntry) GetCustomerSessionIDOk() (*string, bool)`

GetCustomerSessionIDOk returns a tuple with the CustomerSessionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSessionID

`func (o *LoyaltyLedgerEntry) SetCustomerSessionID(v string)`

SetCustomerSessionID sets CustomerSessionID field to given value.

### HasCustomerSessionID

`func (o *LoyaltyLedgerEntry) HasCustomerSessionID() bool`

HasCustomerSessionID returns a boolean if a field has been set.

### GetEventID

`func (o *LoyaltyLedgerEntry) GetEventID() int64`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *LoyaltyLedgerEntry) GetEventIDOk() (*int64, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *LoyaltyLedgerEntry) SetEventID(v int64)`

SetEventID sets EventID field to given value.

### HasEventID

`func (o *LoyaltyLedgerEntry) HasEventID() bool`

HasEventID returns a boolean if a field has been set.

### GetType

`func (o *LoyaltyLedgerEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoyaltyLedgerEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoyaltyLedgerEntry) SetType(v string)`

SetType sets Type field to given value.


### GetAmount

`func (o *LoyaltyLedgerEntry) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LoyaltyLedgerEntry) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *LoyaltyLedgerEntry) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetStartDate

`func (o *LoyaltyLedgerEntry) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *LoyaltyLedgerEntry) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *LoyaltyLedgerEntry) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *LoyaltyLedgerEntry) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *LoyaltyLedgerEntry) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *LoyaltyLedgerEntry) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *LoyaltyLedgerEntry) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *LoyaltyLedgerEntry) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetName

`func (o *LoyaltyLedgerEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoyaltyLedgerEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoyaltyLedgerEntry) SetName(v string)`

SetName sets Name field to given value.


### GetSubLedgerID

`func (o *LoyaltyLedgerEntry) GetSubLedgerID() string`

GetSubLedgerID returns the SubLedgerID field if non-nil, zero value otherwise.

### GetSubLedgerIDOk

`func (o *LoyaltyLedgerEntry) GetSubLedgerIDOk() (*string, bool)`

GetSubLedgerIDOk returns a tuple with the SubLedgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubLedgerID

`func (o *LoyaltyLedgerEntry) SetSubLedgerID(v string)`

SetSubLedgerID sets SubLedgerID field to given value.


### GetUserID

`func (o *LoyaltyLedgerEntry) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *LoyaltyLedgerEntry) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *LoyaltyLedgerEntry) SetUserID(v int64)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *LoyaltyLedgerEntry) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetArchived

`func (o *LoyaltyLedgerEntry) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *LoyaltyLedgerEntry) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *LoyaltyLedgerEntry) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *LoyaltyLedgerEntry) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetFlags

`func (o *LoyaltyLedgerEntry) GetFlags() LoyaltyLedgerEntryFlags`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *LoyaltyLedgerEntry) GetFlagsOk() (*LoyaltyLedgerEntryFlags, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *LoyaltyLedgerEntry) SetFlags(v LoyaltyLedgerEntryFlags)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *LoyaltyLedgerEntry) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetValidityDuration

`func (o *LoyaltyLedgerEntry) GetValidityDuration() string`

GetValidityDuration returns the ValidityDuration field if non-nil, zero value otherwise.

### GetValidityDurationOk

`func (o *LoyaltyLedgerEntry) GetValidityDurationOk() (*string, bool)`

GetValidityDurationOk returns a tuple with the ValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityDuration

`func (o *LoyaltyLedgerEntry) SetValidityDuration(v string)`

SetValidityDuration sets ValidityDuration field to given value.

### HasValidityDuration

`func (o *LoyaltyLedgerEntry) HasValidityDuration() bool`

HasValidityDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


