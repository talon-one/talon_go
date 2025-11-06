# AccountDashboardStatisticReferrals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **float32** | Total number of referrals initiated by users. | 
**Datetime** | Pointer to [**time.Time**](time.Time.md) | Values aggregated for the specified date. | 

## Methods

### NewAccountDashboardStatisticReferrals

`func NewAccountDashboardStatisticReferrals(total float32, datetime time.Time, ) *AccountDashboardStatisticReferrals`

NewAccountDashboardStatisticReferrals instantiates a new AccountDashboardStatisticReferrals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDashboardStatisticReferralsWithDefaults

`func NewAccountDashboardStatisticReferralsWithDefaults() *AccountDashboardStatisticReferrals`

NewAccountDashboardStatisticReferralsWithDefaults instantiates a new AccountDashboardStatisticReferrals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *AccountDashboardStatisticReferrals) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AccountDashboardStatisticReferrals) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AccountDashboardStatisticReferrals) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetDatetime

`func (o *AccountDashboardStatisticReferrals) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *AccountDashboardStatisticReferrals) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *AccountDashboardStatisticReferrals) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


