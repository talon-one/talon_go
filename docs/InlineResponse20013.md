# InlineResponse20013

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | 
**Data** | Pointer to [**[]Referral**](Referral.md) |  | 

## Methods

### NewInlineResponse20013

`func NewInlineResponse20013(hasMore bool, data []Referral, ) *InlineResponse20013`

NewInlineResponse20013 instantiates a new InlineResponse20013 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20013WithDefaults

`func NewInlineResponse20013WithDefaults() *InlineResponse20013`

NewInlineResponse20013WithDefaults instantiates a new InlineResponse20013 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *InlineResponse20013) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InlineResponse20013) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InlineResponse20013) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *InlineResponse20013) GetData() []Referral`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20013) GetDataOk() (*[]Referral, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20013) SetData(v []Referral)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


