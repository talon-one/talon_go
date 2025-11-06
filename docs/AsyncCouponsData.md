# AsyncCouponsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchID** | Pointer to **string** | The ID of the batch to which the coupon belongs.  **Note:** The Batch ID is generated when coupons are created.  | 
**TypeOfChange** | Pointer to **string** |  | 
**Operation** | Pointer to **string** |  | 
**EmployeeName** | Pointer to **string** |  | 
**NotificationType** | Pointer to **string** | The type of the notification | 

## Methods

### NewAsyncCouponsData

`func NewAsyncCouponsData(batchID string, typeOfChange string, operation string, employeeName string, notificationType string, ) *AsyncCouponsData`

NewAsyncCouponsData instantiates a new AsyncCouponsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncCouponsDataWithDefaults

`func NewAsyncCouponsDataWithDefaults() *AsyncCouponsData`

NewAsyncCouponsDataWithDefaults instantiates a new AsyncCouponsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchID

`func (o *AsyncCouponsData) GetBatchID() string`

GetBatchID returns the BatchID field if non-nil, zero value otherwise.

### GetBatchIDOk

`func (o *AsyncCouponsData) GetBatchIDOk() (*string, bool)`

GetBatchIDOk returns a tuple with the BatchID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchID

`func (o *AsyncCouponsData) SetBatchID(v string)`

SetBatchID sets BatchID field to given value.


### GetTypeOfChange

`func (o *AsyncCouponsData) GetTypeOfChange() string`

GetTypeOfChange returns the TypeOfChange field if non-nil, zero value otherwise.

### GetTypeOfChangeOk

`func (o *AsyncCouponsData) GetTypeOfChangeOk() (*string, bool)`

GetTypeOfChangeOk returns a tuple with the TypeOfChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfChange

`func (o *AsyncCouponsData) SetTypeOfChange(v string)`

SetTypeOfChange sets TypeOfChange field to given value.


### GetOperation

`func (o *AsyncCouponsData) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *AsyncCouponsData) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *AsyncCouponsData) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetEmployeeName

`func (o *AsyncCouponsData) GetEmployeeName() string`

GetEmployeeName returns the EmployeeName field if non-nil, zero value otherwise.

### GetEmployeeNameOk

`func (o *AsyncCouponsData) GetEmployeeNameOk() (*string, bool)`

GetEmployeeNameOk returns a tuple with the EmployeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName

`func (o *AsyncCouponsData) SetEmployeeName(v string)`

SetEmployeeName sets EmployeeName field to given value.


### GetNotificationType

`func (o *AsyncCouponsData) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *AsyncCouponsData) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *AsyncCouponsData) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


