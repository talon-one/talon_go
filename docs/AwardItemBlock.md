# AwardItemBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**Sku** | Pointer to **string** | The stock keeping unit of the item to award. | 
**Name** | Pointer to **string** | The display name of the item to award. | 
**Quantity** | Pointer to **string** | The number of items to award. Supports template placeholders (e.g. \&quot;{{$Session.Total / 2}}\&quot;) for dynamic quantities. | 
**Partial** | Pointer to **bool** | When set to &#x60;true&#x60;, applies a partial item reward if the remaining budget is insufficient to award the full reward. | [optional] 
**OnFailure** | Pointer to **[]map[string]interface{}** | Blocks evaluated when this block fails or returns false. | [optional] 
**OnError** | Pointer to [**map[string][]map[string]interface{}**](array.md) | Named error handlers evaluated when a specific error occurs. | [optional] 

## Methods

### NewAwardItemBlock

`func NewAwardItemBlock(type_ string, sku string, name string, quantity string, ) *AwardItemBlock`

NewAwardItemBlock instantiates a new AwardItemBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwardItemBlockWithDefaults

`func NewAwardItemBlockWithDefaults() *AwardItemBlock`

NewAwardItemBlockWithDefaults instantiates a new AwardItemBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AwardItemBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwardItemBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwardItemBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AwardItemBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AwardItemBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwardItemBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwardItemBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *AwardItemBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AwardItemBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AwardItemBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AwardItemBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSku

`func (o *AwardItemBlock) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *AwardItemBlock) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *AwardItemBlock) SetSku(v string)`

SetSku sets Sku field to given value.


### GetName

`func (o *AwardItemBlock) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwardItemBlock) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwardItemBlock) SetName(v string)`

SetName sets Name field to given value.


### GetQuantity

`func (o *AwardItemBlock) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AwardItemBlock) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AwardItemBlock) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetPartial

`func (o *AwardItemBlock) GetPartial() bool`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *AwardItemBlock) GetPartialOk() (*bool, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *AwardItemBlock) SetPartial(v bool)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *AwardItemBlock) HasPartial() bool`

HasPartial returns a boolean if a field has been set.

### GetOnFailure

`func (o *AwardItemBlock) GetOnFailure() []map[string]interface{}`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *AwardItemBlock) GetOnFailureOk() (*[]map[string]interface{}, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *AwardItemBlock) SetOnFailure(v []map[string]interface{})`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *AwardItemBlock) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.

### GetOnError

`func (o *AwardItemBlock) GetOnError() map[string][]map[string]interface{}`

GetOnError returns the OnError field if non-nil, zero value otherwise.

### GetOnErrorOk

`func (o *AwardItemBlock) GetOnErrorOk() (*map[string][]map[string]interface{}, bool)`

GetOnErrorOk returns a tuple with the OnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnError

`func (o *AwardItemBlock) SetOnError(v map[string][]map[string]interface{})`

SetOnError sets OnError field to given value.

### HasOnError

`func (o *AwardItemBlock) HasOnError() bool`

HasOnError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


