# BestPriorPriceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skus** | Pointer to **[]string** | List of product SKUs to check when determining the best prior price. | 
**TimeframeEndDate** | Pointer to [**time.Time**](time.Time.md) | The end date and time that defines the latest time for retrieving historical SKU prices. | 
**Timeframe** | Pointer to **string** | The number of days prior to the timeframeEndDate. Only prices within this look back period are considered for the best prior price evaluation. | 
**StrictEndDate** | Pointer to **bool** | This property is **deprecated**. Use &#x60;timeframeEndDateType&#x60; instead.  Indicates whether the timeframe includes the start of the current sale. - When &#x60;false&#x60;, the timeframe includes the start date of the current sale. - When &#x60;true&#x60;, the timeframe strictly uses the number of days specified in &#x60;timeframe&#x60;.  | 
**TimeframeEndDateType** | Pointer to **string** | Sets the timeframe for retrieving historical pricing data. Can be one of the following values: - &#x60;strict&#x60;: The timeframe ends at the &#x60;timeframeEndDate&#x60; value. - &#x60;price&#x60;: The timeframe ends at the start of the current &#x60;contextId&#x60; with the current price value. Identical price records are merged. If there is no &#x60;contextId&#x60; for the most recent price, the most recent timestamp for the price is used.  - &#x60;sale&#x60;:  The timeframe ends at the start of current &#x60;contextId&#x60; and takes the prices prior to the start of the &#x60;contextId&#x60; into account.  | [optional] 
**Target** | Pointer to [**BestPriorTarget**](BestPriorTarget.md) |  | [optional] 

## Methods

### NewBestPriorPriceRequest

`func NewBestPriorPriceRequest(skus []string, timeframeEndDate time.Time, timeframe string, strictEndDate bool, ) *BestPriorPriceRequest`

NewBestPriorPriceRequest instantiates a new BestPriorPriceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBestPriorPriceRequestWithDefaults

`func NewBestPriorPriceRequestWithDefaults() *BestPriorPriceRequest`

NewBestPriorPriceRequestWithDefaults instantiates a new BestPriorPriceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkus

`func (o *BestPriorPriceRequest) GetSkus() []string`

GetSkus returns the Skus field if non-nil, zero value otherwise.

### GetSkusOk

`func (o *BestPriorPriceRequest) GetSkusOk() (*[]string, bool)`

GetSkusOk returns a tuple with the Skus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkus

`func (o *BestPriorPriceRequest) SetSkus(v []string)`

SetSkus sets Skus field to given value.


### GetTimeframeEndDate

`func (o *BestPriorPriceRequest) GetTimeframeEndDate() time.Time`

GetTimeframeEndDate returns the TimeframeEndDate field if non-nil, zero value otherwise.

### GetTimeframeEndDateOk

`func (o *BestPriorPriceRequest) GetTimeframeEndDateOk() (*time.Time, bool)`

GetTimeframeEndDateOk returns a tuple with the TimeframeEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframeEndDate

`func (o *BestPriorPriceRequest) SetTimeframeEndDate(v time.Time)`

SetTimeframeEndDate sets TimeframeEndDate field to given value.


### GetTimeframe

`func (o *BestPriorPriceRequest) GetTimeframe() string`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *BestPriorPriceRequest) GetTimeframeOk() (*string, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *BestPriorPriceRequest) SetTimeframe(v string)`

SetTimeframe sets Timeframe field to given value.


### GetStrictEndDate

`func (o *BestPriorPriceRequest) GetStrictEndDate() bool`

GetStrictEndDate returns the StrictEndDate field if non-nil, zero value otherwise.

### GetStrictEndDateOk

`func (o *BestPriorPriceRequest) GetStrictEndDateOk() (*bool, bool)`

GetStrictEndDateOk returns a tuple with the StrictEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrictEndDate

`func (o *BestPriorPriceRequest) SetStrictEndDate(v bool)`

SetStrictEndDate sets StrictEndDate field to given value.


### GetTimeframeEndDateType

`func (o *BestPriorPriceRequest) GetTimeframeEndDateType() string`

GetTimeframeEndDateType returns the TimeframeEndDateType field if non-nil, zero value otherwise.

### GetTimeframeEndDateTypeOk

`func (o *BestPriorPriceRequest) GetTimeframeEndDateTypeOk() (*string, bool)`

GetTimeframeEndDateTypeOk returns a tuple with the TimeframeEndDateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframeEndDateType

`func (o *BestPriorPriceRequest) SetTimeframeEndDateType(v string)`

SetTimeframeEndDateType sets TimeframeEndDateType field to given value.

### HasTimeframeEndDateType

`func (o *BestPriorPriceRequest) HasTimeframeEndDateType() bool`

HasTimeframeEndDateType returns a boolean if a field has been set.

### GetTarget

`func (o *BestPriorPriceRequest) GetTarget() BestPriorTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *BestPriorPriceRequest) GetTargetOk() (*BestPriorTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *BestPriorPriceRequest) SetTarget(v BestPriorTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *BestPriorPriceRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


