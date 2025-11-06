# RollbackDeductedLoyaltyPointsEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramId** | Pointer to **int64** | The ID of the loyalty program where these points were reimbursed. | 
**SubLedgerId** | Pointer to **string** | The ID of the subledger within the loyalty program where these points were reimbursed. | 
**Value** | Pointer to **float32** | The amount of reimbursed points that were added. | 
**RecipientIntegrationId** | Pointer to **string** | The user for whom these points were reimbursed. | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Date after which the reimbursed points will be valid. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Date after which the reimbursed points will expire. | [optional] 
**TransactionUUID** | Pointer to **string** | The identifier of &#39;addition&#39; entries added to the ledger as the &#x60;deductLoyaltyPoints&#x60; effect is rolled back. | 
**CardIdentifier** | Pointer to **string** | The alphanumeric identifier of the loyalty card.  | [optional] 

## Methods

### NewRollbackDeductedLoyaltyPointsEffectProps

`func NewRollbackDeductedLoyaltyPointsEffectProps(programId int64, subLedgerId string, value float32, recipientIntegrationId string, transactionUUID string, ) *RollbackDeductedLoyaltyPointsEffectProps`

NewRollbackDeductedLoyaltyPointsEffectProps instantiates a new RollbackDeductedLoyaltyPointsEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollbackDeductedLoyaltyPointsEffectPropsWithDefaults

`func NewRollbackDeductedLoyaltyPointsEffectPropsWithDefaults() *RollbackDeductedLoyaltyPointsEffectProps`

NewRollbackDeductedLoyaltyPointsEffectPropsWithDefaults instantiates a new RollbackDeductedLoyaltyPointsEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgramId

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetProgramId() int64`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetProgramIdOk() (*int64, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *RollbackDeductedLoyaltyPointsEffectProps) SetProgramId(v int64)`

SetProgramId sets ProgramId field to given value.


### GetSubLedgerId

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetSubLedgerId() string`

GetSubLedgerId returns the SubLedgerId field if non-nil, zero value otherwise.

### GetSubLedgerIdOk

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetSubLedgerIdOk() (*string, bool)`

GetSubLedgerIdOk returns a tuple with the SubLedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubLedgerId

`func (o *RollbackDeductedLoyaltyPointsEffectProps) SetSubLedgerId(v string)`

SetSubLedgerId sets SubLedgerId field to given value.


### GetValue

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RollbackDeductedLoyaltyPointsEffectProps) SetValue(v float32)`

SetValue sets Value field to given value.


### GetRecipientIntegrationId

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetRecipientIntegrationIdOk() (*string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientIntegrationId

`func (o *RollbackDeductedLoyaltyPointsEffectProps) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId sets RecipientIntegrationId field to given value.


### GetStartDate

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *RollbackDeductedLoyaltyPointsEffectProps) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *RollbackDeductedLoyaltyPointsEffectProps) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *RollbackDeductedLoyaltyPointsEffectProps) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *RollbackDeductedLoyaltyPointsEffectProps) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetTransactionUUID

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetTransactionUUIDOk() (*string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionUUID

`func (o *RollbackDeductedLoyaltyPointsEffectProps) SetTransactionUUID(v string)`

SetTransactionUUID sets TransactionUUID field to given value.


### GetCardIdentifier

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *RollbackDeductedLoyaltyPointsEffectProps) GetCardIdentifierOk() (*string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIdentifier

`func (o *RollbackDeductedLoyaltyPointsEffectProps) SetCardIdentifier(v string)`

SetCardIdentifier sets CardIdentifier field to given value.

### HasCardIdentifier

`func (o *RollbackDeductedLoyaltyPointsEffectProps) HasCardIdentifier() bool`

HasCardIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


