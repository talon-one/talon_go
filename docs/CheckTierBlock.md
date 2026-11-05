# CheckTierBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**Operator** | Pointer to **string** | An indicator of how the block compares its elements. | 
**Subledger** | Pointer to **string** | The name of the subledger to check a customer&#39;s tier in. Can be empty if this block checks the loyalty program&#39;s main ledger instead of a subledger. A subledger can only be used when the loyalty program that the tier is in has subledgers enabled. To specify the main ledger, provide an empty string (\&quot;\&quot;). | 
**Tier** | Pointer to [**TierBlockReference**](TierBlockReference.md) |  | 
**OnFailure** | Pointer to **[]map[string]interface{}** | Promotion blocks evaluated when this block fails or returns false. | [optional] 

## Methods

### NewCheckTierBlock

`func NewCheckTierBlock(type_ string, operator string, subledger string, tier TierBlockReference, ) *CheckTierBlock`

NewCheckTierBlock instantiates a new CheckTierBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckTierBlockWithDefaults

`func NewCheckTierBlockWithDefaults() *CheckTierBlock`

NewCheckTierBlockWithDefaults instantiates a new CheckTierBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckTierBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckTierBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckTierBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CheckTierBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CheckTierBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckTierBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckTierBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *CheckTierBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CheckTierBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CheckTierBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CheckTierBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *CheckTierBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CheckTierBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CheckTierBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetSubledger

`func (o *CheckTierBlock) GetSubledger() string`

GetSubledger returns the Subledger field if non-nil, zero value otherwise.

### GetSubledgerOk

`func (o *CheckTierBlock) GetSubledgerOk() (*string, bool)`

GetSubledgerOk returns a tuple with the Subledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledger

`func (o *CheckTierBlock) SetSubledger(v string)`

SetSubledger sets Subledger field to given value.


### GetTier

`func (o *CheckTierBlock) GetTier() TierBlockReference`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *CheckTierBlock) GetTierOk() (*TierBlockReference, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *CheckTierBlock) SetTier(v TierBlockReference)`

SetTier sets Tier field to given value.


### GetOnFailure

`func (o *CheckTierBlock) GetOnFailure() []map[string]interface{}`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *CheckTierBlock) GetOnFailureOk() (*[]map[string]interface{}, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *CheckTierBlock) SetOnFailure(v []map[string]interface{})`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *CheckTierBlock) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


