# InlineResponse2003

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | 
**Data** | Pointer to [**[]CardLedgerTransactionLogEntryIntegrationAPI**](CardLedgerTransactionLogEntryIntegrationAPI.md) |  | 

## Methods

### NewInlineResponse2003

`func NewInlineResponse2003(hasMore bool, data []CardLedgerTransactionLogEntryIntegrationAPI, ) *InlineResponse2003`

NewInlineResponse2003 instantiates a new InlineResponse2003 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2003WithDefaults

`func NewInlineResponse2003WithDefaults() *InlineResponse2003`

NewInlineResponse2003WithDefaults instantiates a new InlineResponse2003 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *InlineResponse2003) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InlineResponse2003) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InlineResponse2003) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *InlineResponse2003) GetData() []CardLedgerTransactionLogEntryIntegrationAPI`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse2003) GetDataOk() (*[]CardLedgerTransactionLogEntryIntegrationAPI, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse2003) SetData(v []CardLedgerTransactionLogEntryIntegrationAPI)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


