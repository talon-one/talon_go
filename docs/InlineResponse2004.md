# InlineResponse2004

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | 
**Data** | Pointer to [**[]LedgerTransactionLogEntryIntegrationAPI**](LedgerTransactionLogEntryIntegrationAPI.md) |  | 

## Methods

### NewInlineResponse2004

`func NewInlineResponse2004(hasMore bool, data []LedgerTransactionLogEntryIntegrationAPI, ) *InlineResponse2004`

NewInlineResponse2004 instantiates a new InlineResponse2004 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2004WithDefaults

`func NewInlineResponse2004WithDefaults() *InlineResponse2004`

NewInlineResponse2004WithDefaults instantiates a new InlineResponse2004 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *InlineResponse2004) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InlineResponse2004) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InlineResponse2004) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *InlineResponse2004) GetData() []LedgerTransactionLogEntryIntegrationAPI`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse2004) GetDataOk() (*[]LedgerTransactionLogEntryIntegrationAPI, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse2004) SetData(v []LedgerTransactionLogEntryIntegrationAPI)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


