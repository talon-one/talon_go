# InlineResponse20049

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**[]Achievement**](Achievement.md) |  | 

## Methods

### NewInlineResponse20049

`func NewInlineResponse20049(data []Achievement, ) *InlineResponse20049`

NewInlineResponse20049 instantiates a new InlineResponse20049 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20049WithDefaults

`func NewInlineResponse20049WithDefaults() *InlineResponse20049`

NewInlineResponse20049WithDefaults instantiates a new InlineResponse20049 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *InlineResponse20049) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InlineResponse20049) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InlineResponse20049) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *InlineResponse20049) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetData

`func (o *InlineResponse20049) GetData() []Achievement`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20049) GetDataOk() (*[]Achievement, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20049) SetData(v []Achievement)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


