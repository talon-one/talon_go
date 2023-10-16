# LoyaltyProgramTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the loyalty ledger transaction. | 
**ProgramId** | Pointer to **int32** | ID of the loyalty program. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | Date and time the loyalty transaction occurred. | 
**Type** | Pointer to **string** | Type of transaction. Possible values:   - &#x60;addition&#x60;: Signifies added points.   - &#x60;subtraction&#x60;: Signifies deducted points.  | 
**Amount** | Pointer to **float32** | Amount of loyalty points added or deducted in the transaction. | 
**Name** | Pointer to **string** | Name or reason for the loyalty ledger transaction. | 
**StartDate** | Pointer to **string** | When points become active. Possible values:   - &#x60;immediate&#x60;: Points are immediately active.   - a timestamp value: Points become active at a given date and time.  | 
**ExpiryDate** | Pointer to **string** | When points expire. Possible values:   - &#x60;unlimited&#x60;: Points have no expiration date.   - a timestamp value: Points expire at a given date and time.  | 
**CustomerProfileId** | Pointer to **string** | Customer profile integration ID used in the loyalty program. | [optional] 
**CardIdentifier** | Pointer to **string** | The alphanumeric identifier of the loyalty card.  | [optional] 
**SubledgerId** | Pointer to **string** | ID of the subledger. | 
**CustomerSessionId** | Pointer to **string** | ID of the customer session where the transaction occurred. | [optional] 
**ImportId** | Pointer to **int32** | ID of the import where the transaction occurred. | [optional] 
**UserId** | Pointer to **int32** | ID of the user who manually added or deducted points. Applies only to manual transactions. | [optional] 
**UserEmail** | Pointer to **string** | The email of the Campaign Manager account that manually added or deducted points. Applies only to manual transactions. | [optional] 
**RulesetId** | Pointer to **int32** | ID of the ruleset containing the rule that triggered the effect. Applies only for transactions that resulted from a customer session. | [optional] 
**RuleName** | Pointer to **string** | Name of the rule that triggered the effect. Applies only for transactions that resulted from a customer session. | [optional] 

## Methods

### GetId

`func (o *LoyaltyProgramTransaction) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoyaltyProgramTransaction) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *LoyaltyProgramTransaction) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *LoyaltyProgramTransaction) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetProgramId

`func (o *LoyaltyProgramTransaction) GetProgramId() int32`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *LoyaltyProgramTransaction) GetProgramIdOk() (int32, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgramId

`func (o *LoyaltyProgramTransaction) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### SetProgramId

`func (o *LoyaltyProgramTransaction) SetProgramId(v int32)`

SetProgramId gets a reference to the given int32 and assigns it to the ProgramId field.

### GetCreated

`func (o *LoyaltyProgramTransaction) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LoyaltyProgramTransaction) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *LoyaltyProgramTransaction) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *LoyaltyProgramTransaction) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetType

`func (o *LoyaltyProgramTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoyaltyProgramTransaction) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *LoyaltyProgramTransaction) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *LoyaltyProgramTransaction) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetAmount

`func (o *LoyaltyProgramTransaction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LoyaltyProgramTransaction) GetAmountOk() (float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAmount

`func (o *LoyaltyProgramTransaction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmount

`func (o *LoyaltyProgramTransaction) SetAmount(v float32)`

SetAmount gets a reference to the given float32 and assigns it to the Amount field.

### GetName

`func (o *LoyaltyProgramTransaction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoyaltyProgramTransaction) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *LoyaltyProgramTransaction) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *LoyaltyProgramTransaction) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetStartDate

`func (o *LoyaltyProgramTransaction) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *LoyaltyProgramTransaction) GetStartDateOk() (string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *LoyaltyProgramTransaction) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *LoyaltyProgramTransaction) SetStartDate(v string)`

SetStartDate gets a reference to the given string and assigns it to the StartDate field.

### GetExpiryDate

`func (o *LoyaltyProgramTransaction) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *LoyaltyProgramTransaction) GetExpiryDateOk() (string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *LoyaltyProgramTransaction) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *LoyaltyProgramTransaction) SetExpiryDate(v string)`

SetExpiryDate gets a reference to the given string and assigns it to the ExpiryDate field.

### GetCustomerProfileId

`func (o *LoyaltyProgramTransaction) GetCustomerProfileId() string`

GetCustomerProfileId returns the CustomerProfileId field if non-nil, zero value otherwise.

### GetCustomerProfileIdOk

`func (o *LoyaltyProgramTransaction) GetCustomerProfileIdOk() (string, bool)`

GetCustomerProfileIdOk returns a tuple with the CustomerProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerProfileId

`func (o *LoyaltyProgramTransaction) HasCustomerProfileId() bool`

HasCustomerProfileId returns a boolean if a field has been set.

### SetCustomerProfileId

`func (o *LoyaltyProgramTransaction) SetCustomerProfileId(v string)`

SetCustomerProfileId gets a reference to the given string and assigns it to the CustomerProfileId field.

### GetCardIdentifier

`func (o *LoyaltyProgramTransaction) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *LoyaltyProgramTransaction) GetCardIdentifierOk() (string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCardIdentifier

`func (o *LoyaltyProgramTransaction) HasCardIdentifier() bool`

HasCardIdentifier returns a boolean if a field has been set.

### SetCardIdentifier

`func (o *LoyaltyProgramTransaction) SetCardIdentifier(v string)`

SetCardIdentifier gets a reference to the given string and assigns it to the CardIdentifier field.

### GetSubledgerId

`func (o *LoyaltyProgramTransaction) GetSubledgerId() string`

GetSubledgerId returns the SubledgerId field if non-nil, zero value otherwise.

### GetSubledgerIdOk

`func (o *LoyaltyProgramTransaction) GetSubledgerIdOk() (string, bool)`

GetSubledgerIdOk returns a tuple with the SubledgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubledgerId

`func (o *LoyaltyProgramTransaction) HasSubledgerId() bool`

HasSubledgerId returns a boolean if a field has been set.

### SetSubledgerId

`func (o *LoyaltyProgramTransaction) SetSubledgerId(v string)`

SetSubledgerId gets a reference to the given string and assigns it to the SubledgerId field.

### GetCustomerSessionId

`func (o *LoyaltyProgramTransaction) GetCustomerSessionId() string`

GetCustomerSessionId returns the CustomerSessionId field if non-nil, zero value otherwise.

### GetCustomerSessionIdOk

`func (o *LoyaltyProgramTransaction) GetCustomerSessionIdOk() (string, bool)`

GetCustomerSessionIdOk returns a tuple with the CustomerSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerSessionId

`func (o *LoyaltyProgramTransaction) HasCustomerSessionId() bool`

HasCustomerSessionId returns a boolean if a field has been set.

### SetCustomerSessionId

`func (o *LoyaltyProgramTransaction) SetCustomerSessionId(v string)`

SetCustomerSessionId gets a reference to the given string and assigns it to the CustomerSessionId field.

### GetImportId

`func (o *LoyaltyProgramTransaction) GetImportId() int32`

GetImportId returns the ImportId field if non-nil, zero value otherwise.

### GetImportIdOk

`func (o *LoyaltyProgramTransaction) GetImportIdOk() (int32, bool)`

GetImportIdOk returns a tuple with the ImportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImportId

`func (o *LoyaltyProgramTransaction) HasImportId() bool`

HasImportId returns a boolean if a field has been set.

### SetImportId

`func (o *LoyaltyProgramTransaction) SetImportId(v int32)`

SetImportId gets a reference to the given int32 and assigns it to the ImportId field.

### GetUserId

`func (o *LoyaltyProgramTransaction) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *LoyaltyProgramTransaction) GetUserIdOk() (int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *LoyaltyProgramTransaction) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *LoyaltyProgramTransaction) SetUserId(v int32)`

SetUserId gets a reference to the given int32 and assigns it to the UserId field.

### GetUserEmail

`func (o *LoyaltyProgramTransaction) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *LoyaltyProgramTransaction) GetUserEmailOk() (string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserEmail

`func (o *LoyaltyProgramTransaction) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### SetUserEmail

`func (o *LoyaltyProgramTransaction) SetUserEmail(v string)`

SetUserEmail gets a reference to the given string and assigns it to the UserEmail field.

### GetRulesetId

`func (o *LoyaltyProgramTransaction) GetRulesetId() int32`

GetRulesetId returns the RulesetId field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *LoyaltyProgramTransaction) GetRulesetIdOk() (int32, bool)`

GetRulesetIdOk returns a tuple with the RulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRulesetId

`func (o *LoyaltyProgramTransaction) HasRulesetId() bool`

HasRulesetId returns a boolean if a field has been set.

### SetRulesetId

`func (o *LoyaltyProgramTransaction) SetRulesetId(v int32)`

SetRulesetId gets a reference to the given int32 and assigns it to the RulesetId field.

### GetRuleName

`func (o *LoyaltyProgramTransaction) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *LoyaltyProgramTransaction) GetRuleNameOk() (string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRuleName

`func (o *LoyaltyProgramTransaction) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### SetRuleName

`func (o *LoyaltyProgramTransaction) SetRuleName(v string)`

SetRuleName gets a reference to the given string and assigns it to the RuleName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


