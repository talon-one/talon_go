# ProductUnitAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ProductUnitAnalyticsDataPoint**](ProductUnitAnalyticsDataPoint.md) |  | 
**Totals** | Pointer to [**ProductUnitAnalyticsTotals**](ProductUnitAnalytics_totals.md) |  | 

## Methods

### NewProductUnitAnalytics

`func NewProductUnitAnalytics(data []ProductUnitAnalyticsDataPoint, totals ProductUnitAnalyticsTotals, ) *ProductUnitAnalytics`

NewProductUnitAnalytics instantiates a new ProductUnitAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductUnitAnalyticsWithDefaults

`func NewProductUnitAnalyticsWithDefaults() *ProductUnitAnalytics`

NewProductUnitAnalyticsWithDefaults instantiates a new ProductUnitAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ProductUnitAnalytics) GetData() []ProductUnitAnalyticsDataPoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProductUnitAnalytics) GetDataOk() (*[]ProductUnitAnalyticsDataPoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProductUnitAnalytics) SetData(v []ProductUnitAnalyticsDataPoint)`

SetData sets Data field to given value.


### GetTotals

`func (o *ProductUnitAnalytics) GetTotals() ProductUnitAnalyticsTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *ProductUnitAnalytics) GetTotalsOk() (*ProductUnitAnalyticsTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *ProductUnitAnalytics) SetTotals(v ProductUnitAnalyticsTotals)`

SetTotals sets Totals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


