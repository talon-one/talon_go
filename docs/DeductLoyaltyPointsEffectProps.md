# DeductLoyaltyPointsEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleTitle** | Pointer to **string** | The title of the rule that contained triggered this points deduction. | 
**ProgramId** | Pointer to **int64** | The ID of the loyalty program where these points were added. | 
**SubLedgerId** | Pointer to **string** | The ID of the subledger within the loyalty program where these points were added. | 
**Value** | Pointer to **float32** | The amount of points that were deducted. | 
**TransactionUUID** | Pointer to **string** | The identifier of this deduction in the loyalty ledger. | 
**Name** | Pointer to **string** | The name property gets one of the following two values. It can be the loyalty program name or it can represent a reason for the respective deduction of loyalty points. The latter is an optional value defined in a deduction rule.  | 
**CardIdentifier** | Pointer to **string** | The alphanumeric identifier of the loyalty card.  | [optional] 

## Methods

### NewDeductLoyaltyPointsEffectProps

`func NewDeductLoyaltyPointsEffectProps(ruleTitle string, programId int64, subLedgerId string, value float32, transactionUUID string, name string, ) *DeductLoyaltyPointsEffectProps`

NewDeductLoyaltyPointsEffectProps instantiates a new DeductLoyaltyPointsEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeductLoyaltyPointsEffectPropsWithDefaults

`func NewDeductLoyaltyPointsEffectPropsWithDefaults() *DeductLoyaltyPointsEffectProps`

NewDeductLoyaltyPointsEffectPropsWithDefaults instantiates a new DeductLoyaltyPointsEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleTitle

`func (o *DeductLoyaltyPointsEffectProps) GetRuleTitle() string`

GetRuleTitle returns the RuleTitle field if non-nil, zero value otherwise.

### GetRuleTitleOk

`func (o *DeductLoyaltyPointsEffectProps) GetRuleTitleOk() (*string, bool)`

GetRuleTitleOk returns a tuple with the RuleTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleTitle

`func (o *DeductLoyaltyPointsEffectProps) SetRuleTitle(v string)`

SetRuleTitle sets RuleTitle field to given value.


### GetProgramId

`func (o *DeductLoyaltyPointsEffectProps) GetProgramId() int64`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *DeductLoyaltyPointsEffectProps) GetProgramIdOk() (*int64, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *DeductLoyaltyPointsEffectProps) SetProgramId(v int64)`

SetProgramId sets ProgramId field to given value.


### GetSubLedgerId

`func (o *DeductLoyaltyPointsEffectProps) GetSubLedgerId() string`

GetSubLedgerId returns the SubLedgerId field if non-nil, zero value otherwise.

### GetSubLedgerIdOk

`func (o *DeductLoyaltyPointsEffectProps) GetSubLedgerIdOk() (*string, bool)`

GetSubLedgerIdOk returns a tuple with the SubLedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubLedgerId

`func (o *DeductLoyaltyPointsEffectProps) SetSubLedgerId(v string)`

SetSubLedgerId sets SubLedgerId field to given value.


### GetValue

`func (o *DeductLoyaltyPointsEffectProps) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DeductLoyaltyPointsEffectProps) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DeductLoyaltyPointsEffectProps) SetValue(v float32)`

SetValue sets Value field to given value.


### GetTransactionUUID

`func (o *DeductLoyaltyPointsEffectProps) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *DeductLoyaltyPointsEffectProps) GetTransactionUUIDOk() (*string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionUUID

`func (o *DeductLoyaltyPointsEffectProps) SetTransactionUUID(v string)`

SetTransactionUUID sets TransactionUUID field to given value.


### GetName

`func (o *DeductLoyaltyPointsEffectProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeductLoyaltyPointsEffectProps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeductLoyaltyPointsEffectProps) SetName(v string)`

SetName sets Name field to given value.


### GetCardIdentifier

`func (o *DeductLoyaltyPointsEffectProps) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *DeductLoyaltyPointsEffectProps) GetCardIdentifierOk() (*string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIdentifier

`func (o *DeductLoyaltyPointsEffectProps) SetCardIdentifier(v string)`

SetCardIdentifier sets CardIdentifier field to given value.

### HasCardIdentifier

`func (o *DeductLoyaltyPointsEffectProps) HasCardIdentifier() bool`

HasCardIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


