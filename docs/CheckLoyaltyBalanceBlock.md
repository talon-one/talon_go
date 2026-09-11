# CheckLoyaltyBalanceBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**Operator** | Pointer to **string** | An indicator of how the block compares the balance to the value. | 
**Program** | Pointer to [**CheckLoyaltyBalanceBlockProgram**](CheckLoyaltyBalanceBlock_program.md) |  | 
**Subledger** | Pointer to **string** | The name of the subledger to check the balance of. Can be empty if this block checks the loyalty program&#39;s main ledger balance instead of a subledger. | 
**Balance** | Pointer to **string** | The type of balance to check:  - &#x60;current&#x60; is the sum of currently active points  - &#x60;pending&#x60; is the sum of pending points.  - &#x60;negative&#x60; is the sum of negative points.  - &#x60;tentativeCurrent&#x60; is the tentative points balance within the current open customer session. | 
**Value** | Pointer to **float32** | The numeric value to compare the balance against. | 
**OnFailure** | Pointer to **[]map[string]interface{}** | Promotion blocks evaluated when this block fails or returns false. | [optional] 

## Methods

### NewCheckLoyaltyBalanceBlock

`func NewCheckLoyaltyBalanceBlock(type_ string, operator string, program CheckLoyaltyBalanceBlockProgram, subledger string, balance string, value float32, ) *CheckLoyaltyBalanceBlock`

NewCheckLoyaltyBalanceBlock instantiates a new CheckLoyaltyBalanceBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckLoyaltyBalanceBlockWithDefaults

`func NewCheckLoyaltyBalanceBlockWithDefaults() *CheckLoyaltyBalanceBlock`

NewCheckLoyaltyBalanceBlockWithDefaults instantiates a new CheckLoyaltyBalanceBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckLoyaltyBalanceBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckLoyaltyBalanceBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckLoyaltyBalanceBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CheckLoyaltyBalanceBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CheckLoyaltyBalanceBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckLoyaltyBalanceBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckLoyaltyBalanceBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *CheckLoyaltyBalanceBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CheckLoyaltyBalanceBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CheckLoyaltyBalanceBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CheckLoyaltyBalanceBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *CheckLoyaltyBalanceBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CheckLoyaltyBalanceBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CheckLoyaltyBalanceBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetProgram

`func (o *CheckLoyaltyBalanceBlock) GetProgram() CheckLoyaltyBalanceBlockProgram`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *CheckLoyaltyBalanceBlock) GetProgramOk() (*CheckLoyaltyBalanceBlockProgram, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *CheckLoyaltyBalanceBlock) SetProgram(v CheckLoyaltyBalanceBlockProgram)`

SetProgram sets Program field to given value.


### GetSubledger

`func (o *CheckLoyaltyBalanceBlock) GetSubledger() string`

GetSubledger returns the Subledger field if non-nil, zero value otherwise.

### GetSubledgerOk

`func (o *CheckLoyaltyBalanceBlock) GetSubledgerOk() (*string, bool)`

GetSubledgerOk returns a tuple with the Subledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledger

`func (o *CheckLoyaltyBalanceBlock) SetSubledger(v string)`

SetSubledger sets Subledger field to given value.


### GetBalance

`func (o *CheckLoyaltyBalanceBlock) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CheckLoyaltyBalanceBlock) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CheckLoyaltyBalanceBlock) SetBalance(v string)`

SetBalance sets Balance field to given value.


### GetValue

`func (o *CheckLoyaltyBalanceBlock) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CheckLoyaltyBalanceBlock) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CheckLoyaltyBalanceBlock) SetValue(v float32)`

SetValue sets Value field to given value.


### GetOnFailure

`func (o *CheckLoyaltyBalanceBlock) GetOnFailure() []map[string]interface{}`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *CheckLoyaltyBalanceBlock) GetOnFailureOk() (*[]map[string]interface{}, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *CheckLoyaltyBalanceBlock) SetOnFailure(v []map[string]interface{})`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *CheckLoyaltyBalanceBlock) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


