# AddedDeductedPointsBalancesNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmployeeName** | Pointer to **string** | The name of the employee who added or deducted points. | 
**LoyaltyProgramID** | Pointer to **int64** | The ID of the loyalty program. | 
**NotificationType** | Pointer to **string** | The type of notification. | 
**ProfileIntegrationID** | Pointer to **string** | The integration ID of the customer profile to whom points were added or deducted. | 
**SessionIntegrationID** | Pointer to **string** | The integration ID of the session through which the points were earned or lost. | 
**SubledgerID** | Pointer to **string** | The ID of the subledger within the loyalty program where these points were added. | 
**TypeOfChange** | Pointer to **string** | The notification source, that is, it indicates whether the points were added or deducted via one of the following routes: - [The Campaign Manager](/docs/product/getting-started) - [Management API](/management-api#tag/Loyalty) - [Rule Engine](/docs/product/applications/evaluation-order-for-rules-and-filters)  | 
**UserID** | Pointer to **int64** | The ID of the employee who added or deducted points. | 
**Actions** | Pointer to [**[]AddedDeductedPointsBalancesAction**](AddedDeductedPointsBalancesAction.md) | The list of actions that have been triggered in the loyalty program. | 
**CurrentPoints** | Pointer to **float32** | The current points balance. | 

## Methods

### NewAddedDeductedPointsBalancesNotification

`func NewAddedDeductedPointsBalancesNotification(employeeName string, loyaltyProgramID int64, notificationType string, profileIntegrationID string, sessionIntegrationID string, subledgerID string, typeOfChange string, userID int64, actions []AddedDeductedPointsBalancesAction, currentPoints float32, ) *AddedDeductedPointsBalancesNotification`

NewAddedDeductedPointsBalancesNotification instantiates a new AddedDeductedPointsBalancesNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddedDeductedPointsBalancesNotificationWithDefaults

`func NewAddedDeductedPointsBalancesNotificationWithDefaults() *AddedDeductedPointsBalancesNotification`

NewAddedDeductedPointsBalancesNotificationWithDefaults instantiates a new AddedDeductedPointsBalancesNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployeeName

`func (o *AddedDeductedPointsBalancesNotification) GetEmployeeName() string`

GetEmployeeName returns the EmployeeName field if non-nil, zero value otherwise.

### GetEmployeeNameOk

`func (o *AddedDeductedPointsBalancesNotification) GetEmployeeNameOk() (*string, bool)`

GetEmployeeNameOk returns a tuple with the EmployeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName

`func (o *AddedDeductedPointsBalancesNotification) SetEmployeeName(v string)`

SetEmployeeName sets EmployeeName field to given value.


### GetLoyaltyProgramID

`func (o *AddedDeductedPointsBalancesNotification) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *AddedDeductedPointsBalancesNotification) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *AddedDeductedPointsBalancesNotification) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetNotificationType

`func (o *AddedDeductedPointsBalancesNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *AddedDeductedPointsBalancesNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *AddedDeductedPointsBalancesNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetProfileIntegrationID

`func (o *AddedDeductedPointsBalancesNotification) GetProfileIntegrationID() string`

GetProfileIntegrationID returns the ProfileIntegrationID field if non-nil, zero value otherwise.

### GetProfileIntegrationIDOk

`func (o *AddedDeductedPointsBalancesNotification) GetProfileIntegrationIDOk() (*string, bool)`

GetProfileIntegrationIDOk returns a tuple with the ProfileIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationID

`func (o *AddedDeductedPointsBalancesNotification) SetProfileIntegrationID(v string)`

SetProfileIntegrationID sets ProfileIntegrationID field to given value.


### GetSessionIntegrationID

`func (o *AddedDeductedPointsBalancesNotification) GetSessionIntegrationID() string`

GetSessionIntegrationID returns the SessionIntegrationID field if non-nil, zero value otherwise.

### GetSessionIntegrationIDOk

`func (o *AddedDeductedPointsBalancesNotification) GetSessionIntegrationIDOk() (*string, bool)`

GetSessionIntegrationIDOk returns a tuple with the SessionIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIntegrationID

`func (o *AddedDeductedPointsBalancesNotification) SetSessionIntegrationID(v string)`

SetSessionIntegrationID sets SessionIntegrationID field to given value.


### GetSubledgerID

`func (o *AddedDeductedPointsBalancesNotification) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *AddedDeductedPointsBalancesNotification) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *AddedDeductedPointsBalancesNotification) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetTypeOfChange

`func (o *AddedDeductedPointsBalancesNotification) GetTypeOfChange() string`

GetTypeOfChange returns the TypeOfChange field if non-nil, zero value otherwise.

### GetTypeOfChangeOk

`func (o *AddedDeductedPointsBalancesNotification) GetTypeOfChangeOk() (*string, bool)`

GetTypeOfChangeOk returns a tuple with the TypeOfChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfChange

`func (o *AddedDeductedPointsBalancesNotification) SetTypeOfChange(v string)`

SetTypeOfChange sets TypeOfChange field to given value.


### GetUserID

`func (o *AddedDeductedPointsBalancesNotification) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *AddedDeductedPointsBalancesNotification) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *AddedDeductedPointsBalancesNotification) SetUserID(v int64)`

SetUserID sets UserID field to given value.


### GetActions

`func (o *AddedDeductedPointsBalancesNotification) GetActions() []AddedDeductedPointsBalancesAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AddedDeductedPointsBalancesNotification) GetActionsOk() (*[]AddedDeductedPointsBalancesAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AddedDeductedPointsBalancesNotification) SetActions(v []AddedDeductedPointsBalancesAction)`

SetActions sets Actions field to given value.


### GetCurrentPoints

`func (o *AddedDeductedPointsBalancesNotification) GetCurrentPoints() float32`

GetCurrentPoints returns the CurrentPoints field if non-nil, zero value otherwise.

### GetCurrentPointsOk

`func (o *AddedDeductedPointsBalancesNotification) GetCurrentPointsOk() (*float32, bool)`

GetCurrentPointsOk returns a tuple with the CurrentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPoints

`func (o *AddedDeductedPointsBalancesNotification) SetCurrentPoints(v float32)`

SetCurrentPoints sets CurrentPoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


