# CouponsNotificationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeOfChange** | Pointer to **string** | The type of change that occurred. | 
**Operation** | Pointer to **string** | The operation performed. | 
**EmployeeName** | Pointer to **string** | The name of the employee associated with the operation. | 
**Data** | Pointer to [**[]ExtendedCoupon**](ExtendedCoupon.md) | A list of extended coupon data. | [optional] 
**TotalResultSize** | Pointer to **int64** |  | [optional] 
**NotificationType** | Pointer to **string** | The type of the notification | 

## Methods

### NewCouponsNotificationData

`func NewCouponsNotificationData(typeOfChange string, operation string, employeeName string, notificationType string, ) *CouponsNotificationData`

NewCouponsNotificationData instantiates a new CouponsNotificationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponsNotificationDataWithDefaults

`func NewCouponsNotificationDataWithDefaults() *CouponsNotificationData`

NewCouponsNotificationDataWithDefaults instantiates a new CouponsNotificationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeOfChange

`func (o *CouponsNotificationData) GetTypeOfChange() string`

GetTypeOfChange returns the TypeOfChange field if non-nil, zero value otherwise.

### GetTypeOfChangeOk

`func (o *CouponsNotificationData) GetTypeOfChangeOk() (*string, bool)`

GetTypeOfChangeOk returns a tuple with the TypeOfChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfChange

`func (o *CouponsNotificationData) SetTypeOfChange(v string)`

SetTypeOfChange sets TypeOfChange field to given value.


### GetOperation

`func (o *CouponsNotificationData) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *CouponsNotificationData) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *CouponsNotificationData) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetEmployeeName

`func (o *CouponsNotificationData) GetEmployeeName() string`

GetEmployeeName returns the EmployeeName field if non-nil, zero value otherwise.

### GetEmployeeNameOk

`func (o *CouponsNotificationData) GetEmployeeNameOk() (*string, bool)`

GetEmployeeNameOk returns a tuple with the EmployeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName

`func (o *CouponsNotificationData) SetEmployeeName(v string)`

SetEmployeeName sets EmployeeName field to given value.


### GetData

`func (o *CouponsNotificationData) GetData() []ExtendedCoupon`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CouponsNotificationData) GetDataOk() (*[]ExtendedCoupon, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CouponsNotificationData) SetData(v []ExtendedCoupon)`

SetData sets Data field to given value.

### HasData

`func (o *CouponsNotificationData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalResultSize

`func (o *CouponsNotificationData) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *CouponsNotificationData) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *CouponsNotificationData) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.

### HasTotalResultSize

`func (o *CouponsNotificationData) HasTotalResultSize() bool`

HasTotalResultSize returns a boolean if a field has been set.

### GetNotificationType

`func (o *CouponsNotificationData) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *CouponsNotificationData) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *CouponsNotificationData) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


