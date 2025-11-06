# CreateCouponData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ExtendedCoupon**](ExtendedCoupon.md) | The array of coupons codes. If 1000 or fewer coupons are requested, all coupon data is sent. If 1001 or more coupons are requested, only &#x60;BatchID&#x60; is sent. | [optional] 
**TotalResultSize** | Pointer to **int64** |  | [optional] 
**BatchID** | Pointer to **string** | The ID of the batch to which the coupon belongs.  **Note:** The Batch ID is generated when coupons are created.  | [optional] 
**TypeOfChange** | Pointer to **string** |  | 
**Operation** | Pointer to **string** |  | 
**EmployeeName** | Pointer to **string** |  | 
**NotificationType** | Pointer to **string** | The type of the not | 

## Methods

### NewCreateCouponData

`func NewCreateCouponData(typeOfChange string, operation string, employeeName string, notificationType string, ) *CreateCouponData`

NewCreateCouponData instantiates a new CreateCouponData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCouponDataWithDefaults

`func NewCreateCouponDataWithDefaults() *CreateCouponData`

NewCreateCouponDataWithDefaults instantiates a new CreateCouponData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateCouponData) GetData() []ExtendedCoupon`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateCouponData) GetDataOk() (*[]ExtendedCoupon, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateCouponData) SetData(v []ExtendedCoupon)`

SetData sets Data field to given value.

### HasData

`func (o *CreateCouponData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalResultSize

`func (o *CreateCouponData) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *CreateCouponData) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *CreateCouponData) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.

### HasTotalResultSize

`func (o *CreateCouponData) HasTotalResultSize() bool`

HasTotalResultSize returns a boolean if a field has been set.

### GetBatchID

`func (o *CreateCouponData) GetBatchID() string`

GetBatchID returns the BatchID field if non-nil, zero value otherwise.

### GetBatchIDOk

`func (o *CreateCouponData) GetBatchIDOk() (*string, bool)`

GetBatchIDOk returns a tuple with the BatchID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchID

`func (o *CreateCouponData) SetBatchID(v string)`

SetBatchID sets BatchID field to given value.

### HasBatchID

`func (o *CreateCouponData) HasBatchID() bool`

HasBatchID returns a boolean if a field has been set.

### GetTypeOfChange

`func (o *CreateCouponData) GetTypeOfChange() string`

GetTypeOfChange returns the TypeOfChange field if non-nil, zero value otherwise.

### GetTypeOfChangeOk

`func (o *CreateCouponData) GetTypeOfChangeOk() (*string, bool)`

GetTypeOfChangeOk returns a tuple with the TypeOfChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfChange

`func (o *CreateCouponData) SetTypeOfChange(v string)`

SetTypeOfChange sets TypeOfChange field to given value.


### GetOperation

`func (o *CreateCouponData) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *CreateCouponData) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *CreateCouponData) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetEmployeeName

`func (o *CreateCouponData) GetEmployeeName() string`

GetEmployeeName returns the EmployeeName field if non-nil, zero value otherwise.

### GetEmployeeNameOk

`func (o *CreateCouponData) GetEmployeeNameOk() (*string, bool)`

GetEmployeeNameOk returns a tuple with the EmployeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName

`func (o *CreateCouponData) SetEmployeeName(v string)`

SetEmployeeName sets EmployeeName field to given value.


### GetNotificationType

`func (o *CreateCouponData) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *CreateCouponData) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *CreateCouponData) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


