# GroupBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**Operator** | Pointer to **string** | Logical operator applied across child blocks. &#x60;all&#x60; requires every child to pass, &#x60;atLeastOne&#x60; requires at least one, &#x60;none&#x60; requires all to fail. | 
**Blocks** | Pointer to **[]map[string]interface{}** | Child blocks evaluated according to the operator. | 
**OnFailure** | Pointer to **[]map[string]interface{}** | Blocks evaluated when this block fails or returns false. | [optional] 
**OnError** | Pointer to [**map[string][]map[string]interface{}**](array.md) | Named error handlers evaluated when a specific error occurs. | [optional] 

## Methods

### NewGroupBlock

`func NewGroupBlock(type_ string, operator string, blocks []map[string]interface{}, ) *GroupBlock`

NewGroupBlock instantiates a new GroupBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupBlockWithDefaults

`func NewGroupBlockWithDefaults() *GroupBlock`

NewGroupBlockWithDefaults instantiates a new GroupBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GroupBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *GroupBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GroupBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GroupBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GroupBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *GroupBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *GroupBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *GroupBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetBlocks

`func (o *GroupBlock) GetBlocks() []map[string]interface{}`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *GroupBlock) GetBlocksOk() (*[]map[string]interface{}, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *GroupBlock) SetBlocks(v []map[string]interface{})`

SetBlocks sets Blocks field to given value.


### GetOnFailure

`func (o *GroupBlock) GetOnFailure() []map[string]interface{}`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *GroupBlock) GetOnFailureOk() (*[]map[string]interface{}, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *GroupBlock) SetOnFailure(v []map[string]interface{})`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *GroupBlock) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.

### GetOnError

`func (o *GroupBlock) GetOnError() map[string][]map[string]interface{}`

GetOnError returns the OnError field if non-nil, zero value otherwise.

### GetOnErrorOk

`func (o *GroupBlock) GetOnErrorOk() (*map[string][]map[string]interface{}, bool)`

GetOnErrorOk returns a tuple with the OnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnError

`func (o *GroupBlock) SetOnError(v map[string][]map[string]interface{})`

SetOnError sets OnError field to given value.

### HasOnError

`func (o *GroupBlock) HasOnError() bool`

HasOnError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


