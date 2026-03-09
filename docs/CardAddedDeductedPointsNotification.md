# CardAddedDeductedPointsNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardIdentifier** | Pointer to **string** | Loyalty card identification number. | 
**EmployeeName** | Pointer to **string** | The name of the employee who added or deducted points. | 
**LoyaltyProgramID** | Pointer to **int64** | The ID of the loyalty program. | 
**NotificationType** | Pointer to **string** | The type of notification. | 
**ProfileIntegrationIDs** | Pointer to **[]string** | The integration ID of the customer profile to whom points were added or deducted. | 
**SessionIntegrationID** | Pointer to **string** | The integration ID of the session through which the points were earned or lost. | 
**SubledgerID** | Pointer to **string** | The ID of the subledger within the loyalty program where these points were added or deducted. | 
**TypeOfChange** | Pointer to **string** | The notification source, that is, it indicates whether the points were added or deducted via one of the following routes:  - [The Campaign Manager](/docs/product/getting-started)  - [Management API](/management-api#tag/Loyalty)  - [Rule Engine](/docs/product/applications/evaluation-order-for-rules-and-filters)  | 
**UserID** | Pointer to **int64** | The ID of the employee who added or deducted points. | 
**UsersPerCardLimit** | Pointer to **int64** | The max amount of user profiles with whom a card can be shared. This can be set to &#x60;0&#x60; for no limit. | 
**Amount** | Pointer to **float32** | The amount of added or deducted loyalty points. | 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | The expiration date for loyalty points. | [optional] 
**Operation** | Pointer to **string** | The action (addition or deduction) made with loyalty points. | 
**Reason** | Pointer to **string** | The reason for the points addition or deduction. | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | The start date for loyalty points. | [optional] 
**TransactionUUID** | Pointer to **string** | The identifier of the transaction in the loyalty ledger. | 

## Methods

### NewCardAddedDeductedPointsNotification

`func NewCardAddedDeductedPointsNotification(cardIdentifier string, employeeName string, loyaltyProgramID int64, notificationType string, profileIntegrationIDs []string, sessionIntegrationID string, subledgerID string, typeOfChange string, userID int64, usersPerCardLimit int64, amount float32, operation string, reason string, transactionUUID string, ) *CardAddedDeductedPointsNotification`

NewCardAddedDeductedPointsNotification instantiates a new CardAddedDeductedPointsNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardAddedDeductedPointsNotificationWithDefaults

`func NewCardAddedDeductedPointsNotificationWithDefaults() *CardAddedDeductedPointsNotification`

NewCardAddedDeductedPointsNotificationWithDefaults instantiates a new CardAddedDeductedPointsNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardIdentifier

`func (o *CardAddedDeductedPointsNotification) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *CardAddedDeductedPointsNotification) GetCardIdentifierOk() (*string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIdentifier

`func (o *CardAddedDeductedPointsNotification) SetCardIdentifier(v string)`

SetCardIdentifier sets CardIdentifier field to given value.


### GetEmployeeName

`func (o *CardAddedDeductedPointsNotification) GetEmployeeName() string`

GetEmployeeName returns the EmployeeName field if non-nil, zero value otherwise.

### GetEmployeeNameOk

`func (o *CardAddedDeductedPointsNotification) GetEmployeeNameOk() (*string, bool)`

GetEmployeeNameOk returns a tuple with the EmployeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName

`func (o *CardAddedDeductedPointsNotification) SetEmployeeName(v string)`

SetEmployeeName sets EmployeeName field to given value.


### GetLoyaltyProgramID

`func (o *CardAddedDeductedPointsNotification) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *CardAddedDeductedPointsNotification) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *CardAddedDeductedPointsNotification) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetNotificationType

`func (o *CardAddedDeductedPointsNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *CardAddedDeductedPointsNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *CardAddedDeductedPointsNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetProfileIntegrationIDs

`func (o *CardAddedDeductedPointsNotification) GetProfileIntegrationIDs() []string`

GetProfileIntegrationIDs returns the ProfileIntegrationIDs field if non-nil, zero value otherwise.

### GetProfileIntegrationIDsOk

`func (o *CardAddedDeductedPointsNotification) GetProfileIntegrationIDsOk() (*[]string, bool)`

GetProfileIntegrationIDsOk returns a tuple with the ProfileIntegrationIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationIDs

`func (o *CardAddedDeductedPointsNotification) SetProfileIntegrationIDs(v []string)`

SetProfileIntegrationIDs sets ProfileIntegrationIDs field to given value.


### GetSessionIntegrationID

`func (o *CardAddedDeductedPointsNotification) GetSessionIntegrationID() string`

GetSessionIntegrationID returns the SessionIntegrationID field if non-nil, zero value otherwise.

### GetSessionIntegrationIDOk

`func (o *CardAddedDeductedPointsNotification) GetSessionIntegrationIDOk() (*string, bool)`

GetSessionIntegrationIDOk returns a tuple with the SessionIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIntegrationID

`func (o *CardAddedDeductedPointsNotification) SetSessionIntegrationID(v string)`

SetSessionIntegrationID sets SessionIntegrationID field to given value.


### GetSubledgerID

`func (o *CardAddedDeductedPointsNotification) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *CardAddedDeductedPointsNotification) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *CardAddedDeductedPointsNotification) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetTypeOfChange

`func (o *CardAddedDeductedPointsNotification) GetTypeOfChange() string`

GetTypeOfChange returns the TypeOfChange field if non-nil, zero value otherwise.

### GetTypeOfChangeOk

`func (o *CardAddedDeductedPointsNotification) GetTypeOfChangeOk() (*string, bool)`

GetTypeOfChangeOk returns a tuple with the TypeOfChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfChange

`func (o *CardAddedDeductedPointsNotification) SetTypeOfChange(v string)`

SetTypeOfChange sets TypeOfChange field to given value.


### GetUserID

`func (o *CardAddedDeductedPointsNotification) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *CardAddedDeductedPointsNotification) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *CardAddedDeductedPointsNotification) SetUserID(v int64)`

SetUserID sets UserID field to given value.


### GetUsersPerCardLimit

`func (o *CardAddedDeductedPointsNotification) GetUsersPerCardLimit() int64`

GetUsersPerCardLimit returns the UsersPerCardLimit field if non-nil, zero value otherwise.

### GetUsersPerCardLimitOk

`func (o *CardAddedDeductedPointsNotification) GetUsersPerCardLimitOk() (*int64, bool)`

GetUsersPerCardLimitOk returns a tuple with the UsersPerCardLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersPerCardLimit

`func (o *CardAddedDeductedPointsNotification) SetUsersPerCardLimit(v int64)`

SetUsersPerCardLimit sets UsersPerCardLimit field to given value.


### GetAmount

`func (o *CardAddedDeductedPointsNotification) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CardAddedDeductedPointsNotification) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CardAddedDeductedPointsNotification) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetExpiryDate

`func (o *CardAddedDeductedPointsNotification) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *CardAddedDeductedPointsNotification) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *CardAddedDeductedPointsNotification) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *CardAddedDeductedPointsNotification) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetOperation

`func (o *CardAddedDeductedPointsNotification) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *CardAddedDeductedPointsNotification) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *CardAddedDeductedPointsNotification) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetReason

`func (o *CardAddedDeductedPointsNotification) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CardAddedDeductedPointsNotification) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CardAddedDeductedPointsNotification) SetReason(v string)`

SetReason sets Reason field to given value.


### GetStartDate

`func (o *CardAddedDeductedPointsNotification) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CardAddedDeductedPointsNotification) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CardAddedDeductedPointsNotification) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CardAddedDeductedPointsNotification) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTransactionUUID

`func (o *CardAddedDeductedPointsNotification) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *CardAddedDeductedPointsNotification) GetTransactionUUIDOk() (*string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionUUID

`func (o *CardAddedDeductedPointsNotification) SetTransactionUUID(v string)`

SetTransactionUUID sets TransactionUUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


