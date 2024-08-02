# LedgerTransactionLogEntryIntegrationApi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | Amount of loyalty points added or deducted in the transaction. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | Date and time the loyalty transaction occurred. | 
**CustomerSessionId** | Pointer to **string** | ID of the customer session where the transaction occurred. | [optional] 
**ExpiryDate** | Pointer to **string** | Date when points expire. Possible values are:   - &#x60;unlimited&#x60;: Points have no expiration date.   - &#x60;timestamp value&#x60;: Points expire on the given date.  | 
**Id** | Pointer to **int32** | ID of the loyalty ledger transaction. | 
**Name** | Pointer to **string** | Name or reason of the loyalty ledger transaction. | 
**ProgramId** | Pointer to **int32** | ID of the loyalty program. | 
**RuleName** | Pointer to **string** | The name of the rule that triggered this effect. | [optional] 
**RulesetId** | Pointer to **int32** | The ID of the ruleset containing the rule that triggered this effect. | [optional] 
**StartDate** | Pointer to **string** | When points become active. Possible values:   - &#x60;immediate&#x60;: Points are immediately active.   - a timestamp value: Points become active at a given date and time.  | 
**SubledgerId** | Pointer to **string** | ID of the subledger. | 
**Type** | Pointer to **string** | Type of transaction. Possible values:   - &#x60;addition&#x60;: Signifies added points.   - &#x60;subtraction&#x60;: Signifies deducted points.  | 

## Methods

### GetAmount

`func (o *LedgerTransactionLogEntryIntegrationApi) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LedgerTransactionLogEntryIntegrationApi) GetAmountOk() (float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAmount

`func (o *LedgerTransactionLogEntryIntegrationApi) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmount

`func (o *LedgerTransactionLogEntryIntegrationApi) SetAmount(v float32)`

SetAmount gets a reference to the given float32 and assigns it to the Amount field.

### GetCreated

`func (o *LedgerTransactionLogEntryIntegrationApi) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LedgerTransactionLogEntryIntegrationApi) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *LedgerTransactionLogEntryIntegrationApi) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *LedgerTransactionLogEntryIntegrationApi) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetCustomerSessionId

`func (o *LedgerTransactionLogEntryIntegrationApi) GetCustomerSessionId() string`

GetCustomerSessionId returns the CustomerSessionId field if non-nil, zero value otherwise.

### GetCustomerSessionIdOk

`func (o *LedgerTransactionLogEntryIntegrationApi) GetCustomerSessionIdOk() (string, bool)`

GetCustomerSessionIdOk returns a tuple with the CustomerSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerSessionId

`func (o *LedgerTransactionLogEntryIntegrationApi) HasCustomerSessionId() bool`

HasCustomerSessionId returns a boolean if a field has been set.

### SetCustomerSessionId

`func (o *LedgerTransactionLogEntryIntegrationApi) SetCustomerSessionId(v string)`

SetCustomerSessionId gets a reference to the given string and assigns it to the CustomerSessionId field.

### GetExpiryDate

`func (o *LedgerTransactionLogEntryIntegrationApi) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *LedgerTransactionLogEntryIntegrationApi) GetExpiryDateOk() (string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *LedgerTransactionLogEntryIntegrationApi) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *LedgerTransactionLogEntryIntegrationApi) SetExpiryDate(v string)`

SetExpiryDate gets a reference to the given string and assigns it to the ExpiryDate field.

### GetId

`func (o *LedgerTransactionLogEntryIntegrationApi) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LedgerTransactionLogEntryIntegrationApi) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *LedgerTransactionLogEntryIntegrationApi) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *LedgerTransactionLogEntryIntegrationApi) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetName

`func (o *LedgerTransactionLogEntryIntegrationApi) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LedgerTransactionLogEntryIntegrationApi) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *LedgerTransactionLogEntryIntegrationApi) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *LedgerTransactionLogEntryIntegrationApi) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetProgramId

`func (o *LedgerTransactionLogEntryIntegrationApi) GetProgramId() int32`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *LedgerTransactionLogEntryIntegrationApi) GetProgramIdOk() (int32, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgramId

`func (o *LedgerTransactionLogEntryIntegrationApi) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### SetProgramId

`func (o *LedgerTransactionLogEntryIntegrationApi) SetProgramId(v int32)`

SetProgramId gets a reference to the given int32 and assigns it to the ProgramId field.

### GetRuleName

`func (o *LedgerTransactionLogEntryIntegrationApi) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *LedgerTransactionLogEntryIntegrationApi) GetRuleNameOk() (string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRuleName

`func (o *LedgerTransactionLogEntryIntegrationApi) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### SetRuleName

`func (o *LedgerTransactionLogEntryIntegrationApi) SetRuleName(v string)`

SetRuleName gets a reference to the given string and assigns it to the RuleName field.

### GetRulesetId

`func (o *LedgerTransactionLogEntryIntegrationApi) GetRulesetId() int32`

GetRulesetId returns the RulesetId field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *LedgerTransactionLogEntryIntegrationApi) GetRulesetIdOk() (int32, bool)`

GetRulesetIdOk returns a tuple with the RulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRulesetId

`func (o *LedgerTransactionLogEntryIntegrationApi) HasRulesetId() bool`

HasRulesetId returns a boolean if a field has been set.

### SetRulesetId

`func (o *LedgerTransactionLogEntryIntegrationApi) SetRulesetId(v int32)`

SetRulesetId gets a reference to the given int32 and assigns it to the RulesetId field.

### GetStartDate

`func (o *LedgerTransactionLogEntryIntegrationApi) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *LedgerTransactionLogEntryIntegrationApi) GetStartDateOk() (string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *LedgerTransactionLogEntryIntegrationApi) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *LedgerTransactionLogEntryIntegrationApi) SetStartDate(v string)`

SetStartDate gets a reference to the given string and assigns it to the StartDate field.

### GetSubledgerId

`func (o *LedgerTransactionLogEntryIntegrationApi) GetSubledgerId() string`

GetSubledgerId returns the SubledgerId field if non-nil, zero value otherwise.

### GetSubledgerIdOk

`func (o *LedgerTransactionLogEntryIntegrationApi) GetSubledgerIdOk() (string, bool)`

GetSubledgerIdOk returns a tuple with the SubledgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubledgerId

`func (o *LedgerTransactionLogEntryIntegrationApi) HasSubledgerId() bool`

HasSubledgerId returns a boolean if a field has been set.

### SetSubledgerId

`func (o *LedgerTransactionLogEntryIntegrationApi) SetSubledgerId(v string)`

SetSubledgerId gets a reference to the given string and assigns it to the SubledgerId field.

### GetType

`func (o *LedgerTransactionLogEntryIntegrationApi) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LedgerTransactionLogEntryIntegrationApi) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *LedgerTransactionLogEntryIntegrationApi) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *LedgerTransactionLogEntryIntegrationApi) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


