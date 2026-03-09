# InlineResponse20027

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**TotalResultSize** | Pointer to **int64** |  | [optional] 
**Data** | Pointer to [**[]CustomerProfile**](CustomerProfile.md) |  | 

## Methods

### NewInlineResponse20027

`func NewInlineResponse20027(data []CustomerProfile, ) *InlineResponse20027`

NewInlineResponse20027 instantiates a new InlineResponse20027 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20027WithDefaults

`func NewInlineResponse20027WithDefaults() *InlineResponse20027`

NewInlineResponse20027WithDefaults instantiates a new InlineResponse20027 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *InlineResponse20027) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InlineResponse20027) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InlineResponse20027) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *InlineResponse20027) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetTotalResultSize

`func (o *InlineResponse20027) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *InlineResponse20027) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *InlineResponse20027) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.

### HasTotalResultSize

`func (o *InlineResponse20027) HasTotalResultSize() bool`

HasTotalResultSize returns a boolean if a field has been set.

### GetData

`func (o *InlineResponse20027) GetData() []CustomerProfile`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20027) GetDataOk() (*[]CustomerProfile, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20027) SetData(v []CustomerProfile)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


