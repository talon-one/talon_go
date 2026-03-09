# InlineResponse20022

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | 
**Data** | Pointer to [**[]CollectionItem**](CollectionItem.md) |  | 

## Methods

### NewInlineResponse20022

`func NewInlineResponse20022(hasMore bool, data []CollectionItem, ) *InlineResponse20022`

NewInlineResponse20022 instantiates a new InlineResponse20022 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20022WithDefaults

`func NewInlineResponse20022WithDefaults() *InlineResponse20022`

NewInlineResponse20022WithDefaults instantiates a new InlineResponse20022 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *InlineResponse20022) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InlineResponse20022) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InlineResponse20022) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *InlineResponse20022) GetData() []CollectionItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20022) GetDataOk() (*[]CollectionItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20022) SetData(v []CollectionItem)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


