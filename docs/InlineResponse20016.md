# InlineResponse20016

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** | true means there is more data in the source collection to request.. | 
**Data** | Pointer to [**[]CardLedgerTransactionLogEntry**](CardLedgerTransactionLogEntry.md) | List of loyalty card transaction logs. | 

## Methods

### GetHasMore

`func (o *InlineResponse20016) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InlineResponse20016) GetHasMoreOk() (bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasMore

`func (o *InlineResponse20016) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### SetHasMore

`func (o *InlineResponse20016) SetHasMore(v bool)`

SetHasMore gets a reference to the given bool and assigns it to the HasMore field.

### GetData

`func (o *InlineResponse20016) GetData() []CardLedgerTransactionLogEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20016) GetDataOk() ([]CardLedgerTransactionLogEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasData

`func (o *InlineResponse20016) HasData() bool`

HasData returns a boolean if a field has been set.

### SetData

`func (o *InlineResponse20016) SetData(v []CardLedgerTransactionLogEntry)`

SetData gets a reference to the given []CardLedgerTransactionLogEntry and assigns it to the Data field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


