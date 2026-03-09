# InlineResponse20033

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**TotalResultSize** | Pointer to **int64** |  | [optional] 
**Data** | Pointer to [**[]Audience**](Audience.md) |  | 

## Methods

### NewInlineResponse20033

`func NewInlineResponse20033(data []Audience, ) *InlineResponse20033`

NewInlineResponse20033 instantiates a new InlineResponse20033 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20033WithDefaults

`func NewInlineResponse20033WithDefaults() *InlineResponse20033`

NewInlineResponse20033WithDefaults instantiates a new InlineResponse20033 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *InlineResponse20033) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InlineResponse20033) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InlineResponse20033) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *InlineResponse20033) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetTotalResultSize

`func (o *InlineResponse20033) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *InlineResponse20033) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *InlineResponse20033) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.

### HasTotalResultSize

`func (o *InlineResponse20033) HasTotalResultSize() bool`

HasTotalResultSize returns a boolean if a field has been set.

### GetData

`func (o *InlineResponse20033) GetData() []Audience`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20033) GetDataOk() (*[]Audience, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20033) SetData(v []Audience)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


