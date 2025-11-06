# ProductUnitAnalyticsDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The start of the aggregation time frame in UTC. | 
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The end of the aggregation time frame in UTC. | 
**UnitsSold** | Pointer to [**AnalyticsDataPointWithTrend**](AnalyticsDataPointWithTrend.md) |  | 
**ProductId** | Pointer to **int64** | The ID of the product. | 
**ProductName** | Pointer to **string** | The name of the product. | 

## Methods

### NewProductUnitAnalyticsDataPoint

`func NewProductUnitAnalyticsDataPoint(startTime time.Time, endTime time.Time, unitsSold AnalyticsDataPointWithTrend, productId int64, productName string, ) *ProductUnitAnalyticsDataPoint`

NewProductUnitAnalyticsDataPoint instantiates a new ProductUnitAnalyticsDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductUnitAnalyticsDataPointWithDefaults

`func NewProductUnitAnalyticsDataPointWithDefaults() *ProductUnitAnalyticsDataPoint`

NewProductUnitAnalyticsDataPointWithDefaults instantiates a new ProductUnitAnalyticsDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *ProductUnitAnalyticsDataPoint) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ProductUnitAnalyticsDataPoint) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ProductUnitAnalyticsDataPoint) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *ProductUnitAnalyticsDataPoint) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ProductUnitAnalyticsDataPoint) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ProductUnitAnalyticsDataPoint) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetUnitsSold

`func (o *ProductUnitAnalyticsDataPoint) GetUnitsSold() AnalyticsDataPointWithTrend`

GetUnitsSold returns the UnitsSold field if non-nil, zero value otherwise.

### GetUnitsSoldOk

`func (o *ProductUnitAnalyticsDataPoint) GetUnitsSoldOk() (*AnalyticsDataPointWithTrend, bool)`

GetUnitsSoldOk returns a tuple with the UnitsSold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsSold

`func (o *ProductUnitAnalyticsDataPoint) SetUnitsSold(v AnalyticsDataPointWithTrend)`

SetUnitsSold sets UnitsSold field to given value.


### GetProductId

`func (o *ProductUnitAnalyticsDataPoint) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductUnitAnalyticsDataPoint) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductUnitAnalyticsDataPoint) SetProductId(v int64)`

SetProductId sets ProductId field to given value.


### GetProductName

`func (o *ProductUnitAnalyticsDataPoint) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *ProductUnitAnalyticsDataPoint) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *ProductUnitAnalyticsDataPoint) SetProductName(v string)`

SetProductName sets ProductName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


