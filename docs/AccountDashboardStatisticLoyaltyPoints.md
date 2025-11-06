# AccountDashboardStatisticLoyaltyPoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **float32** | Total loyalty points earned by users. | 
**Datetime** | Pointer to [**time.Time**](time.Time.md) | Values aggregated for the specified date. | 

## Methods

### NewAccountDashboardStatisticLoyaltyPoints

`func NewAccountDashboardStatisticLoyaltyPoints(total float32, datetime time.Time, ) *AccountDashboardStatisticLoyaltyPoints`

NewAccountDashboardStatisticLoyaltyPoints instantiates a new AccountDashboardStatisticLoyaltyPoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDashboardStatisticLoyaltyPointsWithDefaults

`func NewAccountDashboardStatisticLoyaltyPointsWithDefaults() *AccountDashboardStatisticLoyaltyPoints`

NewAccountDashboardStatisticLoyaltyPointsWithDefaults instantiates a new AccountDashboardStatisticLoyaltyPoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *AccountDashboardStatisticLoyaltyPoints) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AccountDashboardStatisticLoyaltyPoints) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AccountDashboardStatisticLoyaltyPoints) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetDatetime

`func (o *AccountDashboardStatisticLoyaltyPoints) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *AccountDashboardStatisticLoyaltyPoints) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *AccountDashboardStatisticLoyaltyPoints) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


