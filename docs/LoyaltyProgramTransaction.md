# LoyaltyProgramTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | ID of the loyalty ledger transaction. | 
**TransactionUUID** | Pointer to **string** | Unique identifier of the transaction in the UUID format. | 
**ProgramId** | Pointer to **int64** | ID of the loyalty program. | 
**CampaignId** | Pointer to **int64** | ID of the campaign. | [optional] 
**Created** | Pointer to [**time.Time**](time.Time.md) | Date and time the loyalty transaction occurred. | 
**Type** | Pointer to **string** | Type of transaction. Possible values:   - &#x60;addition&#x60;: Signifies added points.   - &#x60;subtraction&#x60;: Signifies deducted points.  | 
**Amount** | Pointer to **float32** | Amount of loyalty points added or deducted in the transaction. | 
**Name** | Pointer to **string** | Name or reason for the loyalty ledger transaction. | 
**StartDate** | Pointer to **string** | When points become active. Possible values:   - &#x60;immediate&#x60;: Points are immediately active.   - &#x60;on_action&#x60;: Points become active based on the customer&#39;s action.   - a timestamp value: Points become active at a given date and time.  | 
**ExpiryDate** | Pointer to **string** | When points expire. Possible values:   - &#x60;unlimited&#x60;: Points have no expiration date.   - a timestamp value: Points expire at a given date and time.  | 
**CustomerProfileId** | Pointer to **string** | Customer profile integration ID used in the loyalty program. | [optional] 
**CardIdentifier** | Pointer to **string** | The alphanumeric identifier of the loyalty card.  | [optional] 
**SubledgerId** | Pointer to **string** | ID of the subledger. | 
**CustomerSessionId** | Pointer to **string** | ID of the customer session where the transaction occurred. | [optional] 
**ImportId** | Pointer to **int64** | ID of the import where the transaction occurred. | [optional] 
**UserId** | Pointer to **int64** | ID of the user who manually added or deducted points. Applies only to manual transactions. | [optional] 
**UserEmail** | Pointer to **string** | The email of the Campaign Manager account that manually added or deducted points. Applies only to manual transactions. | [optional] 
**RulesetId** | Pointer to **int64** | ID of the ruleset containing the rule that triggered the effect. Applies only for transactions that resulted from a customer session. | [optional] 
**RuleName** | Pointer to **string** | Name of the rule that triggered the effect. Applies only for transactions that resulted from a customer session. | [optional] 
**Flags** | Pointer to [**LoyaltyLedgerEntryFlags**](LoyaltyLedgerEntryFlags.md) |  | [optional] 
**ValidityDuration** | Pointer to **string** | The duration for which the points remain active, relative to the  activation date.  **Note**: This only applies to points for which &#x60;awaitsActivation&#x60; is &#x60;true&#x60; and &#x60;expiryDate&#x60; is not set.  | [optional] 

## Methods

### NewLoyaltyProgramTransaction

`func NewLoyaltyProgramTransaction(id int64, transactionUUID string, programId int64, created time.Time, type_ string, amount float32, name string, startDate string, expiryDate string, subledgerId string, ) *LoyaltyProgramTransaction`

NewLoyaltyProgramTransaction instantiates a new LoyaltyProgramTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyProgramTransactionWithDefaults

`func NewLoyaltyProgramTransactionWithDefaults() *LoyaltyProgramTransaction`

NewLoyaltyProgramTransactionWithDefaults instantiates a new LoyaltyProgramTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoyaltyProgramTransaction) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoyaltyProgramTransaction) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoyaltyProgramTransaction) SetId(v int64)`

SetId sets Id field to given value.


### GetTransactionUUID

`func (o *LoyaltyProgramTransaction) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *LoyaltyProgramTransaction) GetTransactionUUIDOk() (*string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionUUID

`func (o *LoyaltyProgramTransaction) SetTransactionUUID(v string)`

SetTransactionUUID sets TransactionUUID field to given value.


### GetProgramId

`func (o *LoyaltyProgramTransaction) GetProgramId() int64`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *LoyaltyProgramTransaction) GetProgramIdOk() (*int64, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *LoyaltyProgramTransaction) SetProgramId(v int64)`

SetProgramId sets ProgramId field to given value.


### GetCampaignId

`func (o *LoyaltyProgramTransaction) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *LoyaltyProgramTransaction) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *LoyaltyProgramTransaction) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *LoyaltyProgramTransaction) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetCreated

`func (o *LoyaltyProgramTransaction) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LoyaltyProgramTransaction) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LoyaltyProgramTransaction) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetType

`func (o *LoyaltyProgramTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoyaltyProgramTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoyaltyProgramTransaction) SetType(v string)`

SetType sets Type field to given value.


### GetAmount

`func (o *LoyaltyProgramTransaction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LoyaltyProgramTransaction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *LoyaltyProgramTransaction) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetName

`func (o *LoyaltyProgramTransaction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoyaltyProgramTransaction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoyaltyProgramTransaction) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *LoyaltyProgramTransaction) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *LoyaltyProgramTransaction) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *LoyaltyProgramTransaction) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetExpiryDate

`func (o *LoyaltyProgramTransaction) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *LoyaltyProgramTransaction) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *LoyaltyProgramTransaction) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.


### GetCustomerProfileId

`func (o *LoyaltyProgramTransaction) GetCustomerProfileId() string`

GetCustomerProfileId returns the CustomerProfileId field if non-nil, zero value otherwise.

### GetCustomerProfileIdOk

`func (o *LoyaltyProgramTransaction) GetCustomerProfileIdOk() (*string, bool)`

GetCustomerProfileIdOk returns a tuple with the CustomerProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProfileId

`func (o *LoyaltyProgramTransaction) SetCustomerProfileId(v string)`

SetCustomerProfileId sets CustomerProfileId field to given value.

### HasCustomerProfileId

`func (o *LoyaltyProgramTransaction) HasCustomerProfileId() bool`

HasCustomerProfileId returns a boolean if a field has been set.

### GetCardIdentifier

`func (o *LoyaltyProgramTransaction) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *LoyaltyProgramTransaction) GetCardIdentifierOk() (*string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIdentifier

`func (o *LoyaltyProgramTransaction) SetCardIdentifier(v string)`

SetCardIdentifier sets CardIdentifier field to given value.

### HasCardIdentifier

`func (o *LoyaltyProgramTransaction) HasCardIdentifier() bool`

HasCardIdentifier returns a boolean if a field has been set.

### GetSubledgerId

`func (o *LoyaltyProgramTransaction) GetSubledgerId() string`

GetSubledgerId returns the SubledgerId field if non-nil, zero value otherwise.

### GetSubledgerIdOk

`func (o *LoyaltyProgramTransaction) GetSubledgerIdOk() (*string, bool)`

GetSubledgerIdOk returns a tuple with the SubledgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerId

`func (o *LoyaltyProgramTransaction) SetSubledgerId(v string)`

SetSubledgerId sets SubledgerId field to given value.


### GetCustomerSessionId

`func (o *LoyaltyProgramTransaction) GetCustomerSessionId() string`

GetCustomerSessionId returns the CustomerSessionId field if non-nil, zero value otherwise.

### GetCustomerSessionIdOk

`func (o *LoyaltyProgramTransaction) GetCustomerSessionIdOk() (*string, bool)`

GetCustomerSessionIdOk returns a tuple with the CustomerSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSessionId

`func (o *LoyaltyProgramTransaction) SetCustomerSessionId(v string)`

SetCustomerSessionId sets CustomerSessionId field to given value.

### HasCustomerSessionId

`func (o *LoyaltyProgramTransaction) HasCustomerSessionId() bool`

HasCustomerSessionId returns a boolean if a field has been set.

### GetImportId

`func (o *LoyaltyProgramTransaction) GetImportId() int64`

GetImportId returns the ImportId field if non-nil, zero value otherwise.

### GetImportIdOk

`func (o *LoyaltyProgramTransaction) GetImportIdOk() (*int64, bool)`

GetImportIdOk returns a tuple with the ImportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportId

`func (o *LoyaltyProgramTransaction) SetImportId(v int64)`

SetImportId sets ImportId field to given value.

### HasImportId

`func (o *LoyaltyProgramTransaction) HasImportId() bool`

HasImportId returns a boolean if a field has been set.

### GetUserId

`func (o *LoyaltyProgramTransaction) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *LoyaltyProgramTransaction) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *LoyaltyProgramTransaction) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *LoyaltyProgramTransaction) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserEmail

`func (o *LoyaltyProgramTransaction) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *LoyaltyProgramTransaction) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *LoyaltyProgramTransaction) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *LoyaltyProgramTransaction) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetRulesetId

`func (o *LoyaltyProgramTransaction) GetRulesetId() int64`

GetRulesetId returns the RulesetId field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *LoyaltyProgramTransaction) GetRulesetIdOk() (*int64, bool)`

GetRulesetIdOk returns a tuple with the RulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetId

`func (o *LoyaltyProgramTransaction) SetRulesetId(v int64)`

SetRulesetId sets RulesetId field to given value.

### HasRulesetId

`func (o *LoyaltyProgramTransaction) HasRulesetId() bool`

HasRulesetId returns a boolean if a field has been set.

### GetRuleName

`func (o *LoyaltyProgramTransaction) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *LoyaltyProgramTransaction) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *LoyaltyProgramTransaction) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *LoyaltyProgramTransaction) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetFlags

`func (o *LoyaltyProgramTransaction) GetFlags() LoyaltyLedgerEntryFlags`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *LoyaltyProgramTransaction) GetFlagsOk() (*LoyaltyLedgerEntryFlags, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *LoyaltyProgramTransaction) SetFlags(v LoyaltyLedgerEntryFlags)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *LoyaltyProgramTransaction) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetValidityDuration

`func (o *LoyaltyProgramTransaction) GetValidityDuration() string`

GetValidityDuration returns the ValidityDuration field if non-nil, zero value otherwise.

### GetValidityDurationOk

`func (o *LoyaltyProgramTransaction) GetValidityDurationOk() (*string, bool)`

GetValidityDurationOk returns a tuple with the ValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityDuration

`func (o *LoyaltyProgramTransaction) SetValidityDuration(v string)`

SetValidityDuration sets ValidityDuration field to given value.

### HasValidityDuration

`func (o *LoyaltyProgramTransaction) HasValidityDuration() bool`

HasValidityDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


