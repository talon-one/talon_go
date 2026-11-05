# LedgerTransactionLogEntryManagementAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionUUID** | Pointer to **string** | Unique identifier of the transaction in the UUID format. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | Date and time the loyalty transaction occurred. | 
**ProgramId** | Pointer to **int64** | ID of the loyalty program. | 
**CustomerSessionId** | Pointer to **string** | ID of the customer session where the transaction occurred. | [optional] 
**StoreIntegrationId** | Pointer to **string** | The integration ID of the store where the transaction occurred. Only set for transactions created by a customer session or event that referenced a store. | [optional] 
**Type** | Pointer to **string** | Type of transaction. Possible values:   - &#x60;addition&#x60;: Signifies added points.   - &#x60;subtraction&#x60;: Signifies deducted points.  | 
**Name** | Pointer to **string** | Name or reason of the loyalty ledger transaction. | 
**StartDate** | Pointer to **string** | When points become active. Possible values:   - &#x60;immediate&#x60;: Points are immediately active.   - &#x60;on_action&#x60;: Points become active based on the customer&#39;s action.   - a timestamp value: Points become active at a given date and time.  | 
**ExpiryDate** | Pointer to **string** | Date when points expire. Possible values are:   - &#x60;unlimited&#x60;: Points have no expiration date.   - &#x60;timestamp value&#x60;: Points expire on the given date.  | 
**SubledgerId** | Pointer to **string** | ID of the subledger. | 
**Amount** | Pointer to **float32** | Amount of loyalty points added or deducted in the transaction. | 
**Id** | Pointer to **int64** | ID of the loyalty ledger transaction. | 
**RulesetId** | Pointer to **int64** | The ID of the ruleset containing the rule that triggered this effect. | [optional] 
**RuleName** | Pointer to **string** | The name of the rule that triggered this effect. | [optional] 
**Flags** | Pointer to [**LoyaltyLedgerEntryFlags**](LoyaltyLedgerEntryFlags.md) |  | [optional] 
**ValidityDuration** | Pointer to **string** | The duration for which the points remain active, relative to the activation date.  **Note**: This only applies to points for which &#x60;awaitsActivation&#x60; is &#x60;true&#x60; and &#x60;expiryDate&#x60; is not set.  | [optional] 
**ReferenceTransactionUUIDs** | Pointer to **[]string** | UUIDs of the addition transactions from which this subtraction deducted loyalty points. Only returned when &#x60;includeReferences&#x60; is &#x60;true&#x60; and the transaction has references.  | [optional] 

## Methods

### NewLedgerTransactionLogEntryManagementAPI

`func NewLedgerTransactionLogEntryManagementAPI(transactionUUID string, created time.Time, programId int64, type_ string, name string, startDate string, expiryDate string, subledgerId string, amount float32, id int64, ) *LedgerTransactionLogEntryManagementAPI`

NewLedgerTransactionLogEntryManagementAPI instantiates a new LedgerTransactionLogEntryManagementAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerTransactionLogEntryManagementAPIWithDefaults

`func NewLedgerTransactionLogEntryManagementAPIWithDefaults() *LedgerTransactionLogEntryManagementAPI`

NewLedgerTransactionLogEntryManagementAPIWithDefaults instantiates a new LedgerTransactionLogEntryManagementAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionUUID

`func (o *LedgerTransactionLogEntryManagementAPI) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *LedgerTransactionLogEntryManagementAPI) GetTransactionUUIDOk() (*string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionUUID

`func (o *LedgerTransactionLogEntryManagementAPI) SetTransactionUUID(v string)`

SetTransactionUUID sets TransactionUUID field to given value.


### GetCreated

`func (o *LedgerTransactionLogEntryManagementAPI) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LedgerTransactionLogEntryManagementAPI) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LedgerTransactionLogEntryManagementAPI) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetProgramId

`func (o *LedgerTransactionLogEntryManagementAPI) GetProgramId() int64`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *LedgerTransactionLogEntryManagementAPI) GetProgramIdOk() (*int64, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *LedgerTransactionLogEntryManagementAPI) SetProgramId(v int64)`

SetProgramId sets ProgramId field to given value.


### GetCustomerSessionId

`func (o *LedgerTransactionLogEntryManagementAPI) GetCustomerSessionId() string`

GetCustomerSessionId returns the CustomerSessionId field if non-nil, zero value otherwise.

### GetCustomerSessionIdOk

`func (o *LedgerTransactionLogEntryManagementAPI) GetCustomerSessionIdOk() (*string, bool)`

GetCustomerSessionIdOk returns a tuple with the CustomerSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSessionId

`func (o *LedgerTransactionLogEntryManagementAPI) SetCustomerSessionId(v string)`

SetCustomerSessionId sets CustomerSessionId field to given value.

### HasCustomerSessionId

`func (o *LedgerTransactionLogEntryManagementAPI) HasCustomerSessionId() bool`

HasCustomerSessionId returns a boolean if a field has been set.

### GetStoreIntegrationId

`func (o *LedgerTransactionLogEntryManagementAPI) GetStoreIntegrationId() string`

GetStoreIntegrationId returns the StoreIntegrationId field if non-nil, zero value otherwise.

### GetStoreIntegrationIdOk

`func (o *LedgerTransactionLogEntryManagementAPI) GetStoreIntegrationIdOk() (*string, bool)`

GetStoreIntegrationIdOk returns a tuple with the StoreIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreIntegrationId

`func (o *LedgerTransactionLogEntryManagementAPI) SetStoreIntegrationId(v string)`

SetStoreIntegrationId sets StoreIntegrationId field to given value.

### HasStoreIntegrationId

`func (o *LedgerTransactionLogEntryManagementAPI) HasStoreIntegrationId() bool`

HasStoreIntegrationId returns a boolean if a field has been set.

### GetType

`func (o *LedgerTransactionLogEntryManagementAPI) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LedgerTransactionLogEntryManagementAPI) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LedgerTransactionLogEntryManagementAPI) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *LedgerTransactionLogEntryManagementAPI) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LedgerTransactionLogEntryManagementAPI) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LedgerTransactionLogEntryManagementAPI) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *LedgerTransactionLogEntryManagementAPI) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *LedgerTransactionLogEntryManagementAPI) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *LedgerTransactionLogEntryManagementAPI) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetExpiryDate

`func (o *LedgerTransactionLogEntryManagementAPI) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *LedgerTransactionLogEntryManagementAPI) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *LedgerTransactionLogEntryManagementAPI) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.


### GetSubledgerId

`func (o *LedgerTransactionLogEntryManagementAPI) GetSubledgerId() string`

GetSubledgerId returns the SubledgerId field if non-nil, zero value otherwise.

### GetSubledgerIdOk

`func (o *LedgerTransactionLogEntryManagementAPI) GetSubledgerIdOk() (*string, bool)`

GetSubledgerIdOk returns a tuple with the SubledgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerId

`func (o *LedgerTransactionLogEntryManagementAPI) SetSubledgerId(v string)`

SetSubledgerId sets SubledgerId field to given value.


### GetAmount

`func (o *LedgerTransactionLogEntryManagementAPI) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LedgerTransactionLogEntryManagementAPI) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *LedgerTransactionLogEntryManagementAPI) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetId

`func (o *LedgerTransactionLogEntryManagementAPI) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LedgerTransactionLogEntryManagementAPI) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LedgerTransactionLogEntryManagementAPI) SetId(v int64)`

SetId sets Id field to given value.


### GetRulesetId

`func (o *LedgerTransactionLogEntryManagementAPI) GetRulesetId() int64`

GetRulesetId returns the RulesetId field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *LedgerTransactionLogEntryManagementAPI) GetRulesetIdOk() (*int64, bool)`

GetRulesetIdOk returns a tuple with the RulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetId

`func (o *LedgerTransactionLogEntryManagementAPI) SetRulesetId(v int64)`

SetRulesetId sets RulesetId field to given value.

### HasRulesetId

`func (o *LedgerTransactionLogEntryManagementAPI) HasRulesetId() bool`

HasRulesetId returns a boolean if a field has been set.

### GetRuleName

`func (o *LedgerTransactionLogEntryManagementAPI) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *LedgerTransactionLogEntryManagementAPI) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *LedgerTransactionLogEntryManagementAPI) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *LedgerTransactionLogEntryManagementAPI) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetFlags

`func (o *LedgerTransactionLogEntryManagementAPI) GetFlags() LoyaltyLedgerEntryFlags`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *LedgerTransactionLogEntryManagementAPI) GetFlagsOk() (*LoyaltyLedgerEntryFlags, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *LedgerTransactionLogEntryManagementAPI) SetFlags(v LoyaltyLedgerEntryFlags)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *LedgerTransactionLogEntryManagementAPI) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetValidityDuration

`func (o *LedgerTransactionLogEntryManagementAPI) GetValidityDuration() string`

GetValidityDuration returns the ValidityDuration field if non-nil, zero value otherwise.

### GetValidityDurationOk

`func (o *LedgerTransactionLogEntryManagementAPI) GetValidityDurationOk() (*string, bool)`

GetValidityDurationOk returns a tuple with the ValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityDuration

`func (o *LedgerTransactionLogEntryManagementAPI) SetValidityDuration(v string)`

SetValidityDuration sets ValidityDuration field to given value.

### HasValidityDuration

`func (o *LedgerTransactionLogEntryManagementAPI) HasValidityDuration() bool`

HasValidityDuration returns a boolean if a field has been set.

### GetReferenceTransactionUUIDs

`func (o *LedgerTransactionLogEntryManagementAPI) GetReferenceTransactionUUIDs() []string`

GetReferenceTransactionUUIDs returns the ReferenceTransactionUUIDs field if non-nil, zero value otherwise.

### GetReferenceTransactionUUIDsOk

`func (o *LedgerTransactionLogEntryManagementAPI) GetReferenceTransactionUUIDsOk() (*[]string, bool)`

GetReferenceTransactionUUIDsOk returns a tuple with the ReferenceTransactionUUIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceTransactionUUIDs

`func (o *LedgerTransactionLogEntryManagementAPI) SetReferenceTransactionUUIDs(v []string)`

SetReferenceTransactionUUIDs sets ReferenceTransactionUUIDs field to given value.

### HasReferenceTransactionUUIDs

`func (o *LedgerTransactionLogEntryManagementAPI) HasReferenceTransactionUUIDs() bool`

HasReferenceTransactionUUIDs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


