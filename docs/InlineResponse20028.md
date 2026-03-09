# InlineResponse20028

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | 
**Data** | Pointer to [**[]CustomerProfile**](CustomerProfile.md) |  | 

## Methods

### NewInlineResponse20028

`func NewInlineResponse20028(hasMore bool, data []CustomerProfile, ) *InlineResponse20028`

NewInlineResponse20028 instantiates a new InlineResponse20028 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20028WithDefaults

`func NewInlineResponse20028WithDefaults() *InlineResponse20028`

NewInlineResponse20028WithDefaults instantiates a new InlineResponse20028 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *InlineResponse20028) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InlineResponse20028) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InlineResponse20028) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *InlineResponse20028) GetData() []CustomerProfile`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20028) GetDataOk() (*[]CustomerProfile, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20028) SetData(v []CustomerProfile)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


