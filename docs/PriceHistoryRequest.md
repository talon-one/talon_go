# PriceHistoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to **string** | The SKU of the item for which the historical prices are being retrieved. | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | The start date of the period for which historical prices should be retrieved. | 
**EndDate** | Pointer to [**time.Time**](time.Time.md) | The end date of the period for which historical prices should be retrieved. | 

## Methods

### NewPriceHistoryRequest

`func NewPriceHistoryRequest(sku string, startDate time.Time, endDate time.Time, ) *PriceHistoryRequest`

NewPriceHistoryRequest instantiates a new PriceHistoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceHistoryRequestWithDefaults

`func NewPriceHistoryRequestWithDefaults() *PriceHistoryRequest`

NewPriceHistoryRequestWithDefaults instantiates a new PriceHistoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *PriceHistoryRequest) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *PriceHistoryRequest) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *PriceHistoryRequest) SetSku(v string)`

SetSku sets Sku field to given value.


### GetStartDate

`func (o *PriceHistoryRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PriceHistoryRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PriceHistoryRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *PriceHistoryRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PriceHistoryRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PriceHistoryRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


