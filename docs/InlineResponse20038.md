# InlineResponse20038

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**TotalResultSize** | Pointer to **int64** |  | [optional] 
**Data** | Pointer to [**[]CatalogItem**](CatalogItem.md) |  | 

## Methods

### NewInlineResponse20038

`func NewInlineResponse20038(data []CatalogItem, ) *InlineResponse20038`

NewInlineResponse20038 instantiates a new InlineResponse20038 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20038WithDefaults

`func NewInlineResponse20038WithDefaults() *InlineResponse20038`

NewInlineResponse20038WithDefaults instantiates a new InlineResponse20038 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *InlineResponse20038) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InlineResponse20038) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InlineResponse20038) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *InlineResponse20038) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetTotalResultSize

`func (o *InlineResponse20038) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *InlineResponse20038) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *InlineResponse20038) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.

### HasTotalResultSize

`func (o *InlineResponse20038) HasTotalResultSize() bool`

HasTotalResultSize returns a boolean if a field has been set.

### GetData

`func (o *InlineResponse20038) GetData() []CatalogItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20038) GetDataOk() (*[]CatalogItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20038) SetData(v []CatalogItem)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


