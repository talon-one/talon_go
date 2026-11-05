# InlineResponse2008

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | 
**Data** | Pointer to [**[]CardLedgerPointsEntryIntegrationAPI**](CardLedgerPointsEntryIntegrationAPI.md) |  | 

## Methods

### NewInlineResponse2008

`func NewInlineResponse2008(hasMore bool, data []CardLedgerPointsEntryIntegrationAPI, ) *InlineResponse2008`

NewInlineResponse2008 instantiates a new InlineResponse2008 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2008WithDefaults

`func NewInlineResponse2008WithDefaults() *InlineResponse2008`

NewInlineResponse2008WithDefaults instantiates a new InlineResponse2008 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *InlineResponse2008) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InlineResponse2008) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InlineResponse2008) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *InlineResponse2008) GetData() []CardLedgerPointsEntryIntegrationAPI`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse2008) GetDataOk() (*[]CardLedgerPointsEntryIntegrationAPI, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse2008) SetData(v []CardLedgerPointsEntryIntegrationAPI)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


