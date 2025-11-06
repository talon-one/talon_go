# InlineResponse20019

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** | true means there is more data in the source collection to request.. | 
**Data** | Pointer to [**[]CardLedgerTransactionLogEntry**](CardLedgerTransactionLogEntry.md) | List of loyalty card transaction logs. | 

## Methods

### NewInlineResponse20019

`func NewInlineResponse20019(hasMore bool, data []CardLedgerTransactionLogEntry, ) *InlineResponse20019`

NewInlineResponse20019 instantiates a new InlineResponse20019 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20019WithDefaults

`func NewInlineResponse20019WithDefaults() *InlineResponse20019`

NewInlineResponse20019WithDefaults instantiates a new InlineResponse20019 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *InlineResponse20019) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InlineResponse20019) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InlineResponse20019) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *InlineResponse20019) GetData() []CardLedgerTransactionLogEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20019) GetDataOk() (*[]CardLedgerTransactionLogEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20019) SetData(v []CardLedgerTransactionLogEntry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


