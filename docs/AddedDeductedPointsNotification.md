# AddedDeductedPointsNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmployeeName** | Pointer to **string** | The name of the employee who added or deducted points. | 
**LoyaltyProgramID** | Pointer to **int64** | The ID of the loyalty program. | 
**NotificationType** | Pointer to **string** | The type of notification. | 
**ProfileIntegrationID** | Pointer to **string** | The integration ID of the customer profile to whom points were added or deducted. | 
**SessionIntegrationID** | Pointer to **string** | The integration ID of the session through which the points were earned or lost. | 
**SubledgerID** | Pointer to **string** | The ID of the subledger within the loyalty program where these points were added. | 
**TypeOfChange** | Pointer to **string** | The notification source, that is, it indicates whether the points were added or deducted via one of the following routes:  - [The Campaign Manager](/docs/product/getting-started)  - [Management API](/management-api#tag/Loyalty)  - [Rule Engine](/docs/product/applications/evaluation-order-for-rules-and-filters)  | 
**UserID** | Pointer to **int64** | The ID of the employee who added or deducted points. | 
**Amount** | Pointer to **float32** | The amount of added or deducted loyalty points. | 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | The expiration date for loyalty points. | [optional] 
**Operation** | Pointer to **string** | The action (addition or deduction) made with loyalty points. | 
**Reason** | Pointer to **string** | The reason for the points addition or deduction. | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | The start date for loyalty points. | [optional] 
**TransactionUUID** | Pointer to **string** | The identifier of the transaction in the loyalty ledger. | 

## Methods

### NewAddedDeductedPointsNotification

`func NewAddedDeductedPointsNotification(employeeName string, loyaltyProgramID int64, notificationType string, profileIntegrationID string, sessionIntegrationID string, subledgerID string, typeOfChange string, userID int64, amount float32, operation string, reason string, transactionUUID string, ) *AddedDeductedPointsNotification`

NewAddedDeductedPointsNotification instantiates a new AddedDeductedPointsNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddedDeductedPointsNotificationWithDefaults

`func NewAddedDeductedPointsNotificationWithDefaults() *AddedDeductedPointsNotification`

NewAddedDeductedPointsNotificationWithDefaults instantiates a new AddedDeductedPointsNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployeeName

`func (o *AddedDeductedPointsNotification) GetEmployeeName() string`

GetEmployeeName returns the EmployeeName field if non-nil, zero value otherwise.

### GetEmployeeNameOk

`func (o *AddedDeductedPointsNotification) GetEmployeeNameOk() (*string, bool)`

GetEmployeeNameOk returns a tuple with the EmployeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName

`func (o *AddedDeductedPointsNotification) SetEmployeeName(v string)`

SetEmployeeName sets EmployeeName field to given value.


### GetLoyaltyProgramID

`func (o *AddedDeductedPointsNotification) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *AddedDeductedPointsNotification) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *AddedDeductedPointsNotification) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetNotificationType

`func (o *AddedDeductedPointsNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *AddedDeductedPointsNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *AddedDeductedPointsNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetProfileIntegrationID

`func (o *AddedDeductedPointsNotification) GetProfileIntegrationID() string`

GetProfileIntegrationID returns the ProfileIntegrationID field if non-nil, zero value otherwise.

### GetProfileIntegrationIDOk

`func (o *AddedDeductedPointsNotification) GetProfileIntegrationIDOk() (*string, bool)`

GetProfileIntegrationIDOk returns a tuple with the ProfileIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationID

`func (o *AddedDeductedPointsNotification) SetProfileIntegrationID(v string)`

SetProfileIntegrationID sets ProfileIntegrationID field to given value.


### GetSessionIntegrationID

`func (o *AddedDeductedPointsNotification) GetSessionIntegrationID() string`

GetSessionIntegrationID returns the SessionIntegrationID field if non-nil, zero value otherwise.

### GetSessionIntegrationIDOk

`func (o *AddedDeductedPointsNotification) GetSessionIntegrationIDOk() (*string, bool)`

GetSessionIntegrationIDOk returns a tuple with the SessionIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIntegrationID

`func (o *AddedDeductedPointsNotification) SetSessionIntegrationID(v string)`

SetSessionIntegrationID sets SessionIntegrationID field to given value.


### GetSubledgerID

`func (o *AddedDeductedPointsNotification) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *AddedDeductedPointsNotification) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *AddedDeductedPointsNotification) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetTypeOfChange

`func (o *AddedDeductedPointsNotification) GetTypeOfChange() string`

GetTypeOfChange returns the TypeOfChange field if non-nil, zero value otherwise.

### GetTypeOfChangeOk

`func (o *AddedDeductedPointsNotification) GetTypeOfChangeOk() (*string, bool)`

GetTypeOfChangeOk returns a tuple with the TypeOfChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfChange

`func (o *AddedDeductedPointsNotification) SetTypeOfChange(v string)`

SetTypeOfChange sets TypeOfChange field to given value.


### GetUserID

`func (o *AddedDeductedPointsNotification) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *AddedDeductedPointsNotification) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *AddedDeductedPointsNotification) SetUserID(v int64)`

SetUserID sets UserID field to given value.


### GetAmount

`func (o *AddedDeductedPointsNotification) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AddedDeductedPointsNotification) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AddedDeductedPointsNotification) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetExpiryDate

`func (o *AddedDeductedPointsNotification) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *AddedDeductedPointsNotification) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *AddedDeductedPointsNotification) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *AddedDeductedPointsNotification) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetOperation

`func (o *AddedDeductedPointsNotification) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *AddedDeductedPointsNotification) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *AddedDeductedPointsNotification) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetReason

`func (o *AddedDeductedPointsNotification) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AddedDeductedPointsNotification) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AddedDeductedPointsNotification) SetReason(v string)`

SetReason sets Reason field to given value.


### GetStartDate

`func (o *AddedDeductedPointsNotification) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AddedDeductedPointsNotification) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AddedDeductedPointsNotification) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AddedDeductedPointsNotification) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTransactionUUID

`func (o *AddedDeductedPointsNotification) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *AddedDeductedPointsNotification) GetTransactionUUIDOk() (*string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionUUID

`func (o *AddedDeductedPointsNotification) SetTransactionUUID(v string)`

SetTransactionUUID sets TransactionUUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


