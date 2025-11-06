# InlineResponse2006

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | 
**Data** | Pointer to [**[]LedgerPointsEntryIntegrationAPI**](LedgerPointsEntryIntegrationAPI.md) |  | 

## Methods

### NewInlineResponse2006

`func NewInlineResponse2006(hasMore bool, data []LedgerPointsEntryIntegrationAPI, ) *InlineResponse2006`

NewInlineResponse2006 instantiates a new InlineResponse2006 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2006WithDefaults

`func NewInlineResponse2006WithDefaults() *InlineResponse2006`

NewInlineResponse2006WithDefaults instantiates a new InlineResponse2006 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *InlineResponse2006) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InlineResponse2006) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InlineResponse2006) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *InlineResponse2006) GetData() []LedgerPointsEntryIntegrationAPI`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse2006) GetDataOk() (*[]LedgerPointsEntryIntegrationAPI, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse2006) SetData(v []LedgerPointsEntryIntegrationAPI)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


