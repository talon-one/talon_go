# InlineResponse20030

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**[]ApplicationSession**](ApplicationSession.md) |  | 

## Methods

### NewInlineResponse20030

`func NewInlineResponse20030(data []ApplicationSession, ) *InlineResponse20030`

NewInlineResponse20030 instantiates a new InlineResponse20030 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20030WithDefaults

`func NewInlineResponse20030WithDefaults() *InlineResponse20030`

NewInlineResponse20030WithDefaults instantiates a new InlineResponse20030 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *InlineResponse20030) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InlineResponse20030) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InlineResponse20030) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *InlineResponse20030) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetData

`func (o *InlineResponse20030) GetData() []ApplicationSession`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20030) GetDataOk() (*[]ApplicationSession, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20030) SetData(v []ApplicationSession)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


