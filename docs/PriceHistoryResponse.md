# PriceHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to **string** | The SKU of the item for which historical prices should be retrieved. | 
**History** | Pointer to [**[]History**](History.md) |  | 

## Methods

### NewPriceHistoryResponse

`func NewPriceHistoryResponse(sku string, history []History, ) *PriceHistoryResponse`

NewPriceHistoryResponse instantiates a new PriceHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceHistoryResponseWithDefaults

`func NewPriceHistoryResponseWithDefaults() *PriceHistoryResponse`

NewPriceHistoryResponseWithDefaults instantiates a new PriceHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *PriceHistoryResponse) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *PriceHistoryResponse) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *PriceHistoryResponse) SetSku(v string)`

SetSku sets Sku field to given value.


### GetHistory

`func (o *PriceHistoryResponse) GetHistory() []History`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *PriceHistoryResponse) GetHistoryOk() (*[]History, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *PriceHistoryResponse) SetHistory(v []History)`

SetHistory sets History field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


