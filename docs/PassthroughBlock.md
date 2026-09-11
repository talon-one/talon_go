# PassthroughBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | The type discriminator for this block. | 
**Expression** | Pointer to **[]map[string]interface{}** | The raw Talang expression as an array. For a function call, the first element is the function name and subsequent elements are its arguments. For any other expression (for example a bare attribute path or a literal value), this is a single-element array containing that value. | 

## Methods

### NewPassthroughBlock

`func NewPassthroughBlock(type_ string, expression []map[string]interface{}, ) *PassthroughBlock`

NewPassthroughBlock instantiates a new PassthroughBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassthroughBlockWithDefaults

`func NewPassthroughBlockWithDefaults() *PassthroughBlock`

NewPassthroughBlockWithDefaults instantiates a new PassthroughBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PassthroughBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PassthroughBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PassthroughBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PassthroughBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PassthroughBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PassthroughBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PassthroughBlock) SetType(v string)`

SetType sets Type field to given value.


### GetExpression

`func (o *PassthroughBlock) GetExpression() []map[string]interface{}`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *PassthroughBlock) GetExpressionOk() (*[]map[string]interface{}, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *PassthroughBlock) SetExpression(v []map[string]interface{})`

SetExpression sets Expression field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


