# ExperimentVariantResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VariantId** | Pointer to **int64** | The ID of the variant. | [optional] 
**VariantName** | Pointer to **string** | The name of the variant. | [optional] 
**VariantWeight** | Pointer to **int64** | The weight of the variant. | [optional] 
**TotalRevenue** | Pointer to **float32** | The total, pre-discount value of all items purchased in a customer session. | [optional] 
**SessionsCount** | Pointer to **float32** | The number of all closed sessions. | [optional] 
**AvgItemsPerSession** | Pointer to **float32** | The number of items from sessions divided by the number of sessions. | [optional] 
**AvgSessionValue** | Pointer to **float32** | The average customer session value, calculated by dividing the revenue value by the number of sessions. | [optional] 
**AvgDiscountedSessionValue** | Pointer to **float32** | The average customer session value, calculated by dividing the revenue value by the number of sessions. | [optional] 
**TotalDiscounts** | Pointer to **float32** | The total value of discounts given for cart items in sessions. | [optional] 
**CouponsCount** | Pointer to **float32** | The number of times a coupon was successfully redeemed in sessions. | [optional] 

## Methods

### NewExperimentVariantResult

`func NewExperimentVariantResult() *ExperimentVariantResult`

NewExperimentVariantResult instantiates a new ExperimentVariantResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentVariantResultWithDefaults

`func NewExperimentVariantResultWithDefaults() *ExperimentVariantResult`

NewExperimentVariantResultWithDefaults instantiates a new ExperimentVariantResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariantId

`func (o *ExperimentVariantResult) GetVariantId() int64`

GetVariantId returns the VariantId field if non-nil, zero value otherwise.

### GetVariantIdOk

`func (o *ExperimentVariantResult) GetVariantIdOk() (*int64, bool)`

GetVariantIdOk returns a tuple with the VariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantId

`func (o *ExperimentVariantResult) SetVariantId(v int64)`

SetVariantId sets VariantId field to given value.

### HasVariantId

`func (o *ExperimentVariantResult) HasVariantId() bool`

HasVariantId returns a boolean if a field has been set.

### GetVariantName

`func (o *ExperimentVariantResult) GetVariantName() string`

GetVariantName returns the VariantName field if non-nil, zero value otherwise.

### GetVariantNameOk

`func (o *ExperimentVariantResult) GetVariantNameOk() (*string, bool)`

GetVariantNameOk returns a tuple with the VariantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantName

`func (o *ExperimentVariantResult) SetVariantName(v string)`

SetVariantName sets VariantName field to given value.

### HasVariantName

`func (o *ExperimentVariantResult) HasVariantName() bool`

HasVariantName returns a boolean if a field has been set.

### GetVariantWeight

`func (o *ExperimentVariantResult) GetVariantWeight() int64`

GetVariantWeight returns the VariantWeight field if non-nil, zero value otherwise.

### GetVariantWeightOk

`func (o *ExperimentVariantResult) GetVariantWeightOk() (*int64, bool)`

GetVariantWeightOk returns a tuple with the VariantWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantWeight

`func (o *ExperimentVariantResult) SetVariantWeight(v int64)`

SetVariantWeight sets VariantWeight field to given value.

### HasVariantWeight

`func (o *ExperimentVariantResult) HasVariantWeight() bool`

HasVariantWeight returns a boolean if a field has been set.

### GetTotalRevenue

`func (o *ExperimentVariantResult) GetTotalRevenue() float32`

GetTotalRevenue returns the TotalRevenue field if non-nil, zero value otherwise.

### GetTotalRevenueOk

`func (o *ExperimentVariantResult) GetTotalRevenueOk() (*float32, bool)`

GetTotalRevenueOk returns a tuple with the TotalRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRevenue

`func (o *ExperimentVariantResult) SetTotalRevenue(v float32)`

SetTotalRevenue sets TotalRevenue field to given value.

### HasTotalRevenue

`func (o *ExperimentVariantResult) HasTotalRevenue() bool`

HasTotalRevenue returns a boolean if a field has been set.

### GetSessionsCount

`func (o *ExperimentVariantResult) GetSessionsCount() float32`

GetSessionsCount returns the SessionsCount field if non-nil, zero value otherwise.

### GetSessionsCountOk

`func (o *ExperimentVariantResult) GetSessionsCountOk() (*float32, bool)`

GetSessionsCountOk returns a tuple with the SessionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsCount

`func (o *ExperimentVariantResult) SetSessionsCount(v float32)`

SetSessionsCount sets SessionsCount field to given value.

### HasSessionsCount

`func (o *ExperimentVariantResult) HasSessionsCount() bool`

HasSessionsCount returns a boolean if a field has been set.

### GetAvgItemsPerSession

`func (o *ExperimentVariantResult) GetAvgItemsPerSession() float32`

GetAvgItemsPerSession returns the AvgItemsPerSession field if non-nil, zero value otherwise.

### GetAvgItemsPerSessionOk

`func (o *ExperimentVariantResult) GetAvgItemsPerSessionOk() (*float32, bool)`

GetAvgItemsPerSessionOk returns a tuple with the AvgItemsPerSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgItemsPerSession

`func (o *ExperimentVariantResult) SetAvgItemsPerSession(v float32)`

SetAvgItemsPerSession sets AvgItemsPerSession field to given value.

### HasAvgItemsPerSession

`func (o *ExperimentVariantResult) HasAvgItemsPerSession() bool`

HasAvgItemsPerSession returns a boolean if a field has been set.

### GetAvgSessionValue

`func (o *ExperimentVariantResult) GetAvgSessionValue() float32`

GetAvgSessionValue returns the AvgSessionValue field if non-nil, zero value otherwise.

### GetAvgSessionValueOk

`func (o *ExperimentVariantResult) GetAvgSessionValueOk() (*float32, bool)`

GetAvgSessionValueOk returns a tuple with the AvgSessionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgSessionValue

`func (o *ExperimentVariantResult) SetAvgSessionValue(v float32)`

SetAvgSessionValue sets AvgSessionValue field to given value.

### HasAvgSessionValue

`func (o *ExperimentVariantResult) HasAvgSessionValue() bool`

HasAvgSessionValue returns a boolean if a field has been set.

### GetAvgDiscountedSessionValue

`func (o *ExperimentVariantResult) GetAvgDiscountedSessionValue() float32`

GetAvgDiscountedSessionValue returns the AvgDiscountedSessionValue field if non-nil, zero value otherwise.

### GetAvgDiscountedSessionValueOk

`func (o *ExperimentVariantResult) GetAvgDiscountedSessionValueOk() (*float32, bool)`

GetAvgDiscountedSessionValueOk returns a tuple with the AvgDiscountedSessionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDiscountedSessionValue

`func (o *ExperimentVariantResult) SetAvgDiscountedSessionValue(v float32)`

SetAvgDiscountedSessionValue sets AvgDiscountedSessionValue field to given value.

### HasAvgDiscountedSessionValue

`func (o *ExperimentVariantResult) HasAvgDiscountedSessionValue() bool`

HasAvgDiscountedSessionValue returns a boolean if a field has been set.

### GetTotalDiscounts

`func (o *ExperimentVariantResult) GetTotalDiscounts() float32`

GetTotalDiscounts returns the TotalDiscounts field if non-nil, zero value otherwise.

### GetTotalDiscountsOk

`func (o *ExperimentVariantResult) GetTotalDiscountsOk() (*float32, bool)`

GetTotalDiscountsOk returns a tuple with the TotalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscounts

`func (o *ExperimentVariantResult) SetTotalDiscounts(v float32)`

SetTotalDiscounts sets TotalDiscounts field to given value.

### HasTotalDiscounts

`func (o *ExperimentVariantResult) HasTotalDiscounts() bool`

HasTotalDiscounts returns a boolean if a field has been set.

### GetCouponsCount

`func (o *ExperimentVariantResult) GetCouponsCount() float32`

GetCouponsCount returns the CouponsCount field if non-nil, zero value otherwise.

### GetCouponsCountOk

`func (o *ExperimentVariantResult) GetCouponsCountOk() (*float32, bool)`

GetCouponsCountOk returns a tuple with the CouponsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponsCount

`func (o *ExperimentVariantResult) SetCouponsCount(v float32)`

SetCouponsCount sets CouponsCount field to given value.

### HasCouponsCount

`func (o *ExperimentVariantResult) HasCouponsCount() bool`

HasCouponsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


