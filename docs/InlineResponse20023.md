# InlineResponse20023

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | 
**Data** | Pointer to [**[]AccessLogEntry**](AccessLogEntry.md) |  | 

## Methods

### NewInlineResponse20023

`func NewInlineResponse20023(hasMore bool, data []AccessLogEntry, ) *InlineResponse20023`

NewInlineResponse20023 instantiates a new InlineResponse20023 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20023WithDefaults

`func NewInlineResponse20023WithDefaults() *InlineResponse20023`

NewInlineResponse20023WithDefaults instantiates a new InlineResponse20023 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *InlineResponse20023) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InlineResponse20023) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InlineResponse20023) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *InlineResponse20023) GetData() []AccessLogEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20023) GetDataOk() (*[]AccessLogEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20023) SetData(v []AccessLogEntry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


