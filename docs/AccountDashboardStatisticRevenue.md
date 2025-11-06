# AccountDashboardStatisticRevenue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **float32** | All revenue that went through the client&#39;s shop (including purchases that didn’t trigger an effect). | 
**Influenced** | Pointer to **float32** | The revenue that was created by a purchase that triggered an effect (excluding web hooks, notifications). | 
**Datetime** | Pointer to [**time.Time**](time.Time.md) | Values aggregated for the specified date. | 

## Methods

### NewAccountDashboardStatisticRevenue

`func NewAccountDashboardStatisticRevenue(total float32, influenced float32, datetime time.Time, ) *AccountDashboardStatisticRevenue`

NewAccountDashboardStatisticRevenue instantiates a new AccountDashboardStatisticRevenue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDashboardStatisticRevenueWithDefaults

`func NewAccountDashboardStatisticRevenueWithDefaults() *AccountDashboardStatisticRevenue`

NewAccountDashboardStatisticRevenueWithDefaults instantiates a new AccountDashboardStatisticRevenue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *AccountDashboardStatisticRevenue) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AccountDashboardStatisticRevenue) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AccountDashboardStatisticRevenue) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetInfluenced

`func (o *AccountDashboardStatisticRevenue) GetInfluenced() float32`

GetInfluenced returns the Influenced field if non-nil, zero value otherwise.

### GetInfluencedOk

`func (o *AccountDashboardStatisticRevenue) GetInfluencedOk() (*float32, bool)`

GetInfluencedOk returns a tuple with the Influenced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfluenced

`func (o *AccountDashboardStatisticRevenue) SetInfluenced(v float32)`

SetInfluenced sets Influenced field to given value.


### GetDatetime

`func (o *AccountDashboardStatisticRevenue) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *AccountDashboardStatisticRevenue) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *AccountDashboardStatisticRevenue) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


