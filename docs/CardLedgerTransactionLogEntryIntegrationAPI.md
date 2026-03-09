# CardLedgerTransactionLogEntryIntegrationAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionUUID** | Pointer to **string** | Unique identifier of the transaction in the UUID format. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | Date and time the loyalty card transaction occurred. | 
**ProgramId** | Pointer to **int64** | ID of the loyalty program. | 
**CardIdentifier** | Pointer to **string** | The alphanumeric identifier of the loyalty card.  | 
**CustomerSessionId** | Pointer to **string** | ID of the customer session where the transaction occurred. | [optional] 
**Type** | Pointer to **string** | Type of transaction. Possible values:   - &#x60;addition&#x60;: Signifies added points.   - &#x60;subtraction&#x60;: Signifies deducted points.  | 
**Name** | Pointer to **string** | Name or reason of the loyalty ledger transaction. | 
**StartDate** | Pointer to **string** | When points become active. Possible values:   - &#x60;immediate&#x60;: Points are active immediately.   - &#x60;on_action&#x60;: Points become active based on the customer&#39;s action.   - a timestamp value: Points become active at a given date and time.  | 
**ExpiryDate** | Pointer to **string** | Date when points expire. Possible values are:   - &#x60;unlimited&#x60;: Points have no expiration date.   - &#x60;timestamp value&#x60;: Points expire on the given date.  | 
**SubledgerId** | Pointer to **string** | ID of the subledger. | 
**Amount** | Pointer to **float32** | Amount of loyalty points added or deducted in the transaction. | 
**Id** | Pointer to **int64** | ID of the loyalty ledger transaction. | 
**RulesetId** | Pointer to **int64** | The ID of the ruleset containing the rule that triggered this effect. | [optional] 
**RuleName** | Pointer to **string** | The name of the rule that triggered this effect. | [optional] 
**ValidityDuration** | Pointer to **string** | The duration for which the points remain active, relative to the  activation date.  **Note**: This only applies to points for which &#x60;awaitsActivation&#x60; is &#x60;true&#x60; and &#x60;expiryDate&#x60; is not set.  | [optional] 

## Methods

### NewCardLedgerTransactionLogEntryIntegrationAPI

`func NewCardLedgerTransactionLogEntryIntegrationAPI(transactionUUID string, created time.Time, programId int64, cardIdentifier string, type_ string, name string, startDate string, expiryDate string, subledgerId string, amount float32, id int64, ) *CardLedgerTransactionLogEntryIntegrationAPI`

NewCardLedgerTransactionLogEntryIntegrationAPI instantiates a new CardLedgerTransactionLogEntryIntegrationAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardLedgerTransactionLogEntryIntegrationAPIWithDefaults

`func NewCardLedgerTransactionLogEntryIntegrationAPIWithDefaults() *CardLedgerTransactionLogEntryIntegrationAPI`

NewCardLedgerTransactionLogEntryIntegrationAPIWithDefaults instantiates a new CardLedgerTransactionLogEntryIntegrationAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionUUID

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetTransactionUUIDOk() (*string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionUUID

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) SetTransactionUUID(v string)`

SetTransactionUUID sets TransactionUUID field to given value.


### GetCreated

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetProgramId

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetProgramId() int64`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetProgramIdOk() (*int64, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) SetProgramId(v int64)`

SetProgramId sets ProgramId field to given value.


### GetCardIdentifier

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetCardIdentifierOk() (*string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIdentifier

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) SetCardIdentifier(v string)`

SetCardIdentifier sets CardIdentifier field to given value.


### GetCustomerSessionId

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetCustomerSessionId() string`

GetCustomerSessionId returns the CustomerSessionId field if non-nil, zero value otherwise.

### GetCustomerSessionIdOk

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetCustomerSessionIdOk() (*string, bool)`

GetCustomerSessionIdOk returns a tuple with the CustomerSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSessionId

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) SetCustomerSessionId(v string)`

SetCustomerSessionId sets CustomerSessionId field to given value.

### HasCustomerSessionId

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) HasCustomerSessionId() bool`

HasCustomerSessionId returns a boolean if a field has been set.

### GetType

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetExpiryDate

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.


### GetSubledgerId

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetSubledgerId() string`

GetSubledgerId returns the SubledgerId field if non-nil, zero value otherwise.

### GetSubledgerIdOk

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetSubledgerIdOk() (*string, bool)`

GetSubledgerIdOk returns a tuple with the SubledgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerId

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) SetSubledgerId(v string)`

SetSubledgerId sets SubledgerId field to given value.


### GetAmount

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetId

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) SetId(v int64)`

SetId sets Id field to given value.


### GetRulesetId

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetRulesetId() int64`

GetRulesetId returns the RulesetId field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetRulesetIdOk() (*int64, bool)`

GetRulesetIdOk returns a tuple with the RulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetId

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) SetRulesetId(v int64)`

SetRulesetId sets RulesetId field to given value.

### HasRulesetId

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) HasRulesetId() bool`

HasRulesetId returns a boolean if a field has been set.

### GetRuleName

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetValidityDuration

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetValidityDuration() string`

GetValidityDuration returns the ValidityDuration field if non-nil, zero value otherwise.

### GetValidityDurationOk

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) GetValidityDurationOk() (*string, bool)`

GetValidityDurationOk returns a tuple with the ValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityDuration

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) SetValidityDuration(v string)`

SetValidityDuration sets ValidityDuration field to given value.

### HasValidityDuration

`func (o *CardLedgerTransactionLogEntryIntegrationAPI) HasValidityDuration() bool`

HasValidityDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


