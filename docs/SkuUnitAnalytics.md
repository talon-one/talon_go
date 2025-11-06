# SkuUnitAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SkuUnitAnalyticsDataPoint**](SkuUnitAnalyticsDataPoint.md) |  | 
**Totals** | Pointer to [**ProductUnitAnalyticsTotals**](ProductUnitAnalytics_totals.md) |  | 

## Methods

### NewSkuUnitAnalytics

`func NewSkuUnitAnalytics(data []SkuUnitAnalyticsDataPoint, totals ProductUnitAnalyticsTotals, ) *SkuUnitAnalytics`

NewSkuUnitAnalytics instantiates a new SkuUnitAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuUnitAnalyticsWithDefaults

`func NewSkuUnitAnalyticsWithDefaults() *SkuUnitAnalytics`

NewSkuUnitAnalyticsWithDefaults instantiates a new SkuUnitAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SkuUnitAnalytics) GetData() []SkuUnitAnalyticsDataPoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SkuUnitAnalytics) GetDataOk() (*[]SkuUnitAnalyticsDataPoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SkuUnitAnalytics) SetData(v []SkuUnitAnalyticsDataPoint)`

SetData sets Data field to given value.


### GetTotals

`func (o *SkuUnitAnalytics) GetTotals() ProductUnitAnalyticsTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *SkuUnitAnalytics) GetTotalsOk() (*ProductUnitAnalyticsTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *SkuUnitAnalytics) SetTotals(v ProductUnitAnalyticsTotals)`

SetTotals sets Totals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


