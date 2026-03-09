# InlineResponse20046

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**TotalResultSize** | Pointer to **int64** |  | [optional] 
**Data** | Pointer to [**[]Store**](Store.md) |  | 

## Methods

### NewInlineResponse20046

`func NewInlineResponse20046(data []Store, ) *InlineResponse20046`

NewInlineResponse20046 instantiates a new InlineResponse20046 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20046WithDefaults

`func NewInlineResponse20046WithDefaults() *InlineResponse20046`

NewInlineResponse20046WithDefaults instantiates a new InlineResponse20046 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *InlineResponse20046) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InlineResponse20046) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InlineResponse20046) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *InlineResponse20046) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetTotalResultSize

`func (o *InlineResponse20046) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *InlineResponse20046) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *InlineResponse20046) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.

### HasTotalResultSize

`func (o *InlineResponse20046) HasTotalResultSize() bool`

HasTotalResultSize returns a boolean if a field has been set.

### GetData

`func (o *InlineResponse20046) GetData() []Store`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20046) GetDataOk() (*[]Store, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20046) SetData(v []Store)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


