# CardLedgerTransactionLogEntryIntegrationApi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to [**time.Time**](time.Time.md) | Date and time the loyalty card transaction occurred. | 
**ProgramId** | Pointer to **int32** | ID of the loyalty program. | 
**CardIdentifier** | Pointer to **string** | The alphanumeric identifier of the loyalty card.  | 
**CustomerSessionId** | Pointer to **string** | ID of the customer session where the transaction occurred. | [optional] 
**Type** | Pointer to **string** | Type of transaction. Possible values:   - &#x60;addition&#x60;: Signifies added points.   - &#x60;subtraction&#x60;: Signifies deducted points.  | 
**Name** | Pointer to **string** | Name or reason of the loyalty ledger transaction. | 
**StartDate** | Pointer to **string** | When points become active. Possible values:   - &#x60;immediate&#x60;: Points are active immediately.   - a timestamp value: Points become active at a given date and time.  | 
**ExpiryDate** | Pointer to **string** | Date when points expire. Possible values are:   - &#x60;unlimited&#x60;: Points have no expiration date.   - &#x60;timestamp value&#x60;: Points expire on the given date.  | 
**SubledgerId** | Pointer to **string** | ID of the subledger. | 
**Amount** | Pointer to **float32** | Amount of loyalty points added or deducted in the transaction. | 
**Id** | Pointer to **int32** | ID of the loyalty ledger transaction. | 
**RulesetId** | Pointer to **int32** | The ID of the ruleset containing the rule that triggered this effect. | [optional] 
**RuleName** | Pointer to **string** | The name of the rule that triggered this effect. | [optional] 

## Methods

### GetCreated

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *CardLedgerTransactionLogEntryIntegrationApi) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *CardLedgerTransactionLogEntryIntegrationApi) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetProgramId

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetProgramId() int32`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetProgramIdOk() (int32, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgramId

`func (o *CardLedgerTransactionLogEntryIntegrationApi) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### SetProgramId

`func (o *CardLedgerTransactionLogEntryIntegrationApi) SetProgramId(v int32)`

SetProgramId gets a reference to the given int32 and assigns it to the ProgramId field.

### GetCardIdentifier

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetCardIdentifierOk() (string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCardIdentifier

`func (o *CardLedgerTransactionLogEntryIntegrationApi) HasCardIdentifier() bool`

HasCardIdentifier returns a boolean if a field has been set.

### SetCardIdentifier

`func (o *CardLedgerTransactionLogEntryIntegrationApi) SetCardIdentifier(v string)`

SetCardIdentifier gets a reference to the given string and assigns it to the CardIdentifier field.

### GetCustomerSessionId

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetCustomerSessionId() string`

GetCustomerSessionId returns the CustomerSessionId field if non-nil, zero value otherwise.

### GetCustomerSessionIdOk

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetCustomerSessionIdOk() (string, bool)`

GetCustomerSessionIdOk returns a tuple with the CustomerSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerSessionId

`func (o *CardLedgerTransactionLogEntryIntegrationApi) HasCustomerSessionId() bool`

HasCustomerSessionId returns a boolean if a field has been set.

### SetCustomerSessionId

`func (o *CardLedgerTransactionLogEntryIntegrationApi) SetCustomerSessionId(v string)`

SetCustomerSessionId gets a reference to the given string and assigns it to the CustomerSessionId field.

### GetType

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *CardLedgerTransactionLogEntryIntegrationApi) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *CardLedgerTransactionLogEntryIntegrationApi) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetName

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *CardLedgerTransactionLogEntryIntegrationApi) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *CardLedgerTransactionLogEntryIntegrationApi) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetStartDate

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetStartDateOk() (string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *CardLedgerTransactionLogEntryIntegrationApi) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *CardLedgerTransactionLogEntryIntegrationApi) SetStartDate(v string)`

SetStartDate gets a reference to the given string and assigns it to the StartDate field.

### GetExpiryDate

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetExpiryDateOk() (string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *CardLedgerTransactionLogEntryIntegrationApi) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *CardLedgerTransactionLogEntryIntegrationApi) SetExpiryDate(v string)`

SetExpiryDate gets a reference to the given string and assigns it to the ExpiryDate field.

### GetSubledgerId

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetSubledgerId() string`

GetSubledgerId returns the SubledgerId field if non-nil, zero value otherwise.

### GetSubledgerIdOk

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetSubledgerIdOk() (string, bool)`

GetSubledgerIdOk returns a tuple with the SubledgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubledgerId

`func (o *CardLedgerTransactionLogEntryIntegrationApi) HasSubledgerId() bool`

HasSubledgerId returns a boolean if a field has been set.

### SetSubledgerId

`func (o *CardLedgerTransactionLogEntryIntegrationApi) SetSubledgerId(v string)`

SetSubledgerId gets a reference to the given string and assigns it to the SubledgerId field.

### GetAmount

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetAmountOk() (float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAmount

`func (o *CardLedgerTransactionLogEntryIntegrationApi) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmount

`func (o *CardLedgerTransactionLogEntryIntegrationApi) SetAmount(v float32)`

SetAmount gets a reference to the given float32 and assigns it to the Amount field.

### GetId

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *CardLedgerTransactionLogEntryIntegrationApi) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *CardLedgerTransactionLogEntryIntegrationApi) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetRulesetId

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetRulesetId() int32`

GetRulesetId returns the RulesetId field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetRulesetIdOk() (int32, bool)`

GetRulesetIdOk returns a tuple with the RulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRulesetId

`func (o *CardLedgerTransactionLogEntryIntegrationApi) HasRulesetId() bool`

HasRulesetId returns a boolean if a field has been set.

### SetRulesetId

`func (o *CardLedgerTransactionLogEntryIntegrationApi) SetRulesetId(v int32)`

SetRulesetId gets a reference to the given int32 and assigns it to the RulesetId field.

### GetRuleName

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *CardLedgerTransactionLogEntryIntegrationApi) GetRuleNameOk() (string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRuleName

`func (o *CardLedgerTransactionLogEntryIntegrationApi) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### SetRuleName

`func (o *CardLedgerTransactionLogEntryIntegrationApi) SetRuleName(v string)`

SetRuleName gets a reference to the given string and assigns it to the RuleName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


