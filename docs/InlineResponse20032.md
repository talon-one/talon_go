# InlineResponse20032

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | 
**Data** | Pointer to [**[]ApplicationEvent**](ApplicationEvent.md) |  | 

## Methods

### NewInlineResponse20032

`func NewInlineResponse20032(hasMore bool, data []ApplicationEvent, ) *InlineResponse20032`

NewInlineResponse20032 instantiates a new InlineResponse20032 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20032WithDefaults

`func NewInlineResponse20032WithDefaults() *InlineResponse20032`

NewInlineResponse20032WithDefaults instantiates a new InlineResponse20032 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *InlineResponse20032) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InlineResponse20032) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InlineResponse20032) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *InlineResponse20032) GetData() []ApplicationEvent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20032) GetDataOk() (*[]ApplicationEvent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20032) SetData(v []ApplicationEvent)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


