# ExpiringPointsNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | Pointer to **int64** |  | 
**Data** | Pointer to [**[]ExpiringPointsData**](ExpiringPointsData.md) | The array of expiring points. | 
**NotificationType** | Pointer to **string** | The type of notification. | 

## Methods

### NewExpiringPointsNotification

`func NewExpiringPointsNotification(totalResultSize int64, data []ExpiringPointsData, notificationType string, ) *ExpiringPointsNotification`

NewExpiringPointsNotification instantiates a new ExpiringPointsNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpiringPointsNotificationWithDefaults

`func NewExpiringPointsNotificationWithDefaults() *ExpiringPointsNotification`

NewExpiringPointsNotificationWithDefaults instantiates a new ExpiringPointsNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *ExpiringPointsNotification) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *ExpiringPointsNotification) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *ExpiringPointsNotification) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetData

`func (o *ExpiringPointsNotification) GetData() []ExpiringPointsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExpiringPointsNotification) GetDataOk() (*[]ExpiringPointsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExpiringPointsNotification) SetData(v []ExpiringPointsData)`

SetData sets Data field to given value.


### GetNotificationType

`func (o *ExpiringPointsNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *ExpiringPointsNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *ExpiringPointsNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


