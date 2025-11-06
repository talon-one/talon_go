# CustomerAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedCoupons** | Pointer to **int64** | Total accepted coupons for this customer. | 
**CreatedCoupons** | Pointer to **int64** | Total created coupons for this customer. | 
**FreeItems** | Pointer to **int64** | Total free items given to this customer. | 
**TotalOrders** | Pointer to **int64** | Total orders made by this customer. | 
**TotalDiscountedOrders** | Pointer to **int64** | Total orders made by this customer that had a discount. | 
**TotalRevenue** | Pointer to **float32** | Total Revenue across all closed sessions. | 
**TotalDiscounts** | Pointer to **float32** | The sum of discounts that were given across all closed sessions. | 

## Methods

### NewCustomerAnalytics

`func NewCustomerAnalytics(acceptedCoupons int64, createdCoupons int64, freeItems int64, totalOrders int64, totalDiscountedOrders int64, totalRevenue float32, totalDiscounts float32, ) *CustomerAnalytics`

NewCustomerAnalytics instantiates a new CustomerAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAnalyticsWithDefaults

`func NewCustomerAnalyticsWithDefaults() *CustomerAnalytics`

NewCustomerAnalyticsWithDefaults instantiates a new CustomerAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedCoupons

`func (o *CustomerAnalytics) GetAcceptedCoupons() int64`

GetAcceptedCoupons returns the AcceptedCoupons field if non-nil, zero value otherwise.

### GetAcceptedCouponsOk

`func (o *CustomerAnalytics) GetAcceptedCouponsOk() (*int64, bool)`

GetAcceptedCouponsOk returns a tuple with the AcceptedCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCoupons

`func (o *CustomerAnalytics) SetAcceptedCoupons(v int64)`

SetAcceptedCoupons sets AcceptedCoupons field to given value.


### GetCreatedCoupons

`func (o *CustomerAnalytics) GetCreatedCoupons() int64`

GetCreatedCoupons returns the CreatedCoupons field if non-nil, zero value otherwise.

### GetCreatedCouponsOk

`func (o *CustomerAnalytics) GetCreatedCouponsOk() (*int64, bool)`

GetCreatedCouponsOk returns a tuple with the CreatedCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedCoupons

`func (o *CustomerAnalytics) SetCreatedCoupons(v int64)`

SetCreatedCoupons sets CreatedCoupons field to given value.


### GetFreeItems

`func (o *CustomerAnalytics) GetFreeItems() int64`

GetFreeItems returns the FreeItems field if non-nil, zero value otherwise.

### GetFreeItemsOk

`func (o *CustomerAnalytics) GetFreeItemsOk() (*int64, bool)`

GetFreeItemsOk returns a tuple with the FreeItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeItems

`func (o *CustomerAnalytics) SetFreeItems(v int64)`

SetFreeItems sets FreeItems field to given value.


### GetTotalOrders

`func (o *CustomerAnalytics) GetTotalOrders() int64`

GetTotalOrders returns the TotalOrders field if non-nil, zero value otherwise.

### GetTotalOrdersOk

`func (o *CustomerAnalytics) GetTotalOrdersOk() (*int64, bool)`

GetTotalOrdersOk returns a tuple with the TotalOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOrders

`func (o *CustomerAnalytics) SetTotalOrders(v int64)`

SetTotalOrders sets TotalOrders field to given value.


### GetTotalDiscountedOrders

`func (o *CustomerAnalytics) GetTotalDiscountedOrders() int64`

GetTotalDiscountedOrders returns the TotalDiscountedOrders field if non-nil, zero value otherwise.

### GetTotalDiscountedOrdersOk

`func (o *CustomerAnalytics) GetTotalDiscountedOrdersOk() (*int64, bool)`

GetTotalDiscountedOrdersOk returns a tuple with the TotalDiscountedOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountedOrders

`func (o *CustomerAnalytics) SetTotalDiscountedOrders(v int64)`

SetTotalDiscountedOrders sets TotalDiscountedOrders field to given value.


### GetTotalRevenue

`func (o *CustomerAnalytics) GetTotalRevenue() float32`

GetTotalRevenue returns the TotalRevenue field if non-nil, zero value otherwise.

### GetTotalRevenueOk

`func (o *CustomerAnalytics) GetTotalRevenueOk() (*float32, bool)`

GetTotalRevenueOk returns a tuple with the TotalRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRevenue

`func (o *CustomerAnalytics) SetTotalRevenue(v float32)`

SetTotalRevenue sets TotalRevenue field to given value.


### GetTotalDiscounts

`func (o *CustomerAnalytics) GetTotalDiscounts() float32`

GetTotalDiscounts returns the TotalDiscounts field if non-nil, zero value otherwise.

### GetTotalDiscountsOk

`func (o *CustomerAnalytics) GetTotalDiscountsOk() (*float32, bool)`

GetTotalDiscountsOk returns a tuple with the TotalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscounts

`func (o *CustomerAnalytics) SetTotalDiscounts(v float32)`

SetTotalDiscounts sets TotalDiscounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


