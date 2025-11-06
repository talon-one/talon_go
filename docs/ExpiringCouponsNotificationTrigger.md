# ExpiringCouponsNotificationTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | The amount of period. | 
**Period** | Pointer to **string** | Notification period indicated by a letter; \&quot;w\&quot; means week, \&quot;d\&quot; means day. | 

## Methods

### NewExpiringCouponsNotificationTrigger

`func NewExpiringCouponsNotificationTrigger(amount int64, period string, ) *ExpiringCouponsNotificationTrigger`

NewExpiringCouponsNotificationTrigger instantiates a new ExpiringCouponsNotificationTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpiringCouponsNotificationTriggerWithDefaults

`func NewExpiringCouponsNotificationTriggerWithDefaults() *ExpiringCouponsNotificationTrigger`

NewExpiringCouponsNotificationTriggerWithDefaults instantiates a new ExpiringCouponsNotificationTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ExpiringCouponsNotificationTrigger) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ExpiringCouponsNotificationTrigger) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ExpiringCouponsNotificationTrigger) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetPeriod

`func (o *ExpiringCouponsNotificationTrigger) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *ExpiringCouponsNotificationTrigger) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *ExpiringCouponsNotificationTrigger) SetPeriod(v string)`

SetPeriod sets Period field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


