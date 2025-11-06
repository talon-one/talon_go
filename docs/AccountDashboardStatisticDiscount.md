# AccountDashboardStatisticDiscount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **float32** | Total discount value redeemed by users. | 
**Average** | Pointer to **float32** | Average discount percentage. | 
**Datetime** | Pointer to [**time.Time**](time.Time.md) | Values aggregated for the specified date. | 

## Methods

### NewAccountDashboardStatisticDiscount

`func NewAccountDashboardStatisticDiscount(total float32, average float32, datetime time.Time, ) *AccountDashboardStatisticDiscount`

NewAccountDashboardStatisticDiscount instantiates a new AccountDashboardStatisticDiscount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDashboardStatisticDiscountWithDefaults

`func NewAccountDashboardStatisticDiscountWithDefaults() *AccountDashboardStatisticDiscount`

NewAccountDashboardStatisticDiscountWithDefaults instantiates a new AccountDashboardStatisticDiscount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *AccountDashboardStatisticDiscount) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AccountDashboardStatisticDiscount) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AccountDashboardStatisticDiscount) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetAverage

`func (o *AccountDashboardStatisticDiscount) GetAverage() float32`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *AccountDashboardStatisticDiscount) GetAverageOk() (*float32, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *AccountDashboardStatisticDiscount) SetAverage(v float32)`

SetAverage sets Average field to given value.


### GetDatetime

`func (o *AccountDashboardStatisticDiscount) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *AccountDashboardStatisticDiscount) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *AccountDashboardStatisticDiscount) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


