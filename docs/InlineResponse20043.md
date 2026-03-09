# InlineResponse20043

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | Pointer to **int64** |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**[]Change**](Change.md) |  | 

## Methods

### NewInlineResponse20043

`func NewInlineResponse20043(data []Change, ) *InlineResponse20043`

NewInlineResponse20043 instantiates a new InlineResponse20043 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20043WithDefaults

`func NewInlineResponse20043WithDefaults() *InlineResponse20043`

NewInlineResponse20043WithDefaults instantiates a new InlineResponse20043 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *InlineResponse20043) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *InlineResponse20043) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *InlineResponse20043) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.

### HasTotalResultSize

`func (o *InlineResponse20043) HasTotalResultSize() bool`

HasTotalResultSize returns a boolean if a field has been set.

### GetHasMore

`func (o *InlineResponse20043) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InlineResponse20043) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InlineResponse20043) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *InlineResponse20043) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetData

`func (o *InlineResponse20043) GetData() []Change`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20043) GetDataOk() (*[]Change, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20043) SetData(v []Change)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


