# ExpiringCouponsNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | Pointer to **int64** |  | 
**Data** | Pointer to [**[]ExpiringCouponsData**](ExpiringCouponsData.md) | The array of expiring coupon notifications. | 
**NotificationType** | Pointer to **string** | The type of notification. | 

## Methods

### NewExpiringCouponsNotification

`func NewExpiringCouponsNotification(totalResultSize int64, data []ExpiringCouponsData, notificationType string, ) *ExpiringCouponsNotification`

NewExpiringCouponsNotification instantiates a new ExpiringCouponsNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpiringCouponsNotificationWithDefaults

`func NewExpiringCouponsNotificationWithDefaults() *ExpiringCouponsNotification`

NewExpiringCouponsNotificationWithDefaults instantiates a new ExpiringCouponsNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *ExpiringCouponsNotification) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *ExpiringCouponsNotification) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *ExpiringCouponsNotification) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetData

`func (o *ExpiringCouponsNotification) GetData() []ExpiringCouponsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExpiringCouponsNotification) GetDataOk() (*[]ExpiringCouponsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExpiringCouponsNotification) SetData(v []ExpiringCouponsData)`

SetData sets Data field to given value.


### GetNotificationType

`func (o *ExpiringCouponsNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *ExpiringCouponsNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *ExpiringCouponsNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


