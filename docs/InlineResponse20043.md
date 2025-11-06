# InlineResponse20043

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | Pointer to **int64** |  | 
**Data** | Pointer to [**[]Export**](Export.md) |  | 

## Methods

### NewInlineResponse20043

`func NewInlineResponse20043(totalResultSize int64, data []Export, ) *InlineResponse20043`

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


### GetData

`func (o *InlineResponse20043) GetData() []Export`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20043) GetDataOk() (*[]Export, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20043) SetData(v []Export)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


