# RuleV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the rule. | [optional] 
**ParentId** | Pointer to **string** | ID of the parent rule, if any. | [optional] 
**Title** | Pointer to **string** | A short description of the rule. | 
**Description** | Pointer to **string** | A longer description of the rule. | [optional] 
**Blocks** | Pointer to **[]map[string]interface{}** | The condition and effect blocks that make up this rule. | 

## Methods

### NewRuleV2

`func NewRuleV2(title string, blocks []map[string]interface{}, ) *RuleV2`

NewRuleV2 instantiates a new RuleV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleV2WithDefaults

`func NewRuleV2WithDefaults() *RuleV2`

NewRuleV2WithDefaults instantiates a new RuleV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RuleV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RuleV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentId

`func (o *RuleV2) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *RuleV2) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *RuleV2) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *RuleV2) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetTitle

`func (o *RuleV2) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RuleV2) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RuleV2) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *RuleV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuleV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuleV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RuleV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBlocks

`func (o *RuleV2) GetBlocks() []map[string]interface{}`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *RuleV2) GetBlocksOk() (*[]map[string]interface{}, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *RuleV2) SetBlocks(v []map[string]interface{})`

SetBlocks sets Blocks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


