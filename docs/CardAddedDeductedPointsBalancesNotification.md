# CardAddedDeductedPointsBalancesNotification

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
**TypeOfChange** | Pointer to **string** | The notification source, that is, it indicates whether the points were added or deducted via one of the following routes: - [The Campaign Manager](/docs/product/getting-started) - [Management API](/management-api#tag/Loyalty) - [Rule Engine](/docs/product/applications/evaluation-order-for-rules-and-filters)  | 
**UserID** | Pointer to **int64** | The ID of the employee who added or deducted points. | 
**UsersPerCardLimit** | Pointer to **int64** | The max amount of user profiles with whom a card can be shared. This can be set to &#x60;0&#x60; for no limit. | 
**Actions** | Pointer to [**[]AddedDeductedPointsBalancesAction**](AddedDeductedPointsBalancesAction.md) | The list of actions that have been triggered in the loyalty program. | 
**CurrentPoints** | Pointer to **float32** | The current points balance. | 

## Methods

### NewCardAddedDeductedPointsBalancesNotification

`func NewCardAddedDeductedPointsBalancesNotification(cardIdentifier string, employeeName string, loyaltyProgramID int64, notificationType string, profileIntegrationIDs []string, sessionIntegrationID string, subledgerID string, typeOfChange string, userID int64, usersPerCardLimit int64, actions []AddedDeductedPointsBalancesAction, currentPoints float32, ) *CardAddedDeductedPointsBalancesNotification`

NewCardAddedDeductedPointsBalancesNotification instantiates a new CardAddedDeductedPointsBalancesNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardAddedDeductedPointsBalancesNotificationWithDefaults

`func NewCardAddedDeductedPointsBalancesNotificationWithDefaults() *CardAddedDeductedPointsBalancesNotification`

NewCardAddedDeductedPointsBalancesNotificationWithDefaults instantiates a new CardAddedDeductedPointsBalancesNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardIdentifier

`func (o *CardAddedDeductedPointsBalancesNotification) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *CardAddedDeductedPointsBalancesNotification) GetCardIdentifierOk() (*string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIdentifier

`func (o *CardAddedDeductedPointsBalancesNotification) SetCardIdentifier(v string)`

SetCardIdentifier sets CardIdentifier field to given value.


### GetEmployeeName

`func (o *CardAddedDeductedPointsBalancesNotification) GetEmployeeName() string`

GetEmployeeName returns the EmployeeName field if non-nil, zero value otherwise.

### GetEmployeeNameOk

`func (o *CardAddedDeductedPointsBalancesNotification) GetEmployeeNameOk() (*string, bool)`

GetEmployeeNameOk returns a tuple with the EmployeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName

`func (o *CardAddedDeductedPointsBalancesNotification) SetEmployeeName(v string)`

SetEmployeeName sets EmployeeName field to given value.


### GetLoyaltyProgramID

`func (o *CardAddedDeductedPointsBalancesNotification) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *CardAddedDeductedPointsBalancesNotification) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *CardAddedDeductedPointsBalancesNotification) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetNotificationType

`func (o *CardAddedDeductedPointsBalancesNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *CardAddedDeductedPointsBalancesNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *CardAddedDeductedPointsBalancesNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetProfileIntegrationIDs

`func (o *CardAddedDeductedPointsBalancesNotification) GetProfileIntegrationIDs() []string`

GetProfileIntegrationIDs returns the ProfileIntegrationIDs field if non-nil, zero value otherwise.

### GetProfileIntegrationIDsOk

`func (o *CardAddedDeductedPointsBalancesNotification) GetProfileIntegrationIDsOk() (*[]string, bool)`

GetProfileIntegrationIDsOk returns a tuple with the ProfileIntegrationIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationIDs

`func (o *CardAddedDeductedPointsBalancesNotification) SetProfileIntegrationIDs(v []string)`

SetProfileIntegrationIDs sets ProfileIntegrationIDs field to given value.


### GetSessionIntegrationID

`func (o *CardAddedDeductedPointsBalancesNotification) GetSessionIntegrationID() string`

GetSessionIntegrationID returns the SessionIntegrationID field if non-nil, zero value otherwise.

### GetSessionIntegrationIDOk

`func (o *CardAddedDeductedPointsBalancesNotification) GetSessionIntegrationIDOk() (*string, bool)`

GetSessionIntegrationIDOk returns a tuple with the SessionIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIntegrationID

`func (o *CardAddedDeductedPointsBalancesNotification) SetSessionIntegrationID(v string)`

SetSessionIntegrationID sets SessionIntegrationID field to given value.


### GetSubledgerID

`func (o *CardAddedDeductedPointsBalancesNotification) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *CardAddedDeductedPointsBalancesNotification) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *CardAddedDeductedPointsBalancesNotification) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetTypeOfChange

`func (o *CardAddedDeductedPointsBalancesNotification) GetTypeOfChange() string`

GetTypeOfChange returns the TypeOfChange field if non-nil, zero value otherwise.

### GetTypeOfChangeOk

`func (o *CardAddedDeductedPointsBalancesNotification) GetTypeOfChangeOk() (*string, bool)`

GetTypeOfChangeOk returns a tuple with the TypeOfChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfChange

`func (o *CardAddedDeductedPointsBalancesNotification) SetTypeOfChange(v string)`

SetTypeOfChange sets TypeOfChange field to given value.


### GetUserID

`func (o *CardAddedDeductedPointsBalancesNotification) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *CardAddedDeductedPointsBalancesNotification) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *CardAddedDeductedPointsBalancesNotification) SetUserID(v int64)`

SetUserID sets UserID field to given value.


### GetUsersPerCardLimit

`func (o *CardAddedDeductedPointsBalancesNotification) GetUsersPerCardLimit() int64`

GetUsersPerCardLimit returns the UsersPerCardLimit field if non-nil, zero value otherwise.

### GetUsersPerCardLimitOk

`func (o *CardAddedDeductedPointsBalancesNotification) GetUsersPerCardLimitOk() (*int64, bool)`

GetUsersPerCardLimitOk returns a tuple with the UsersPerCardLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersPerCardLimit

`func (o *CardAddedDeductedPointsBalancesNotification) SetUsersPerCardLimit(v int64)`

SetUsersPerCardLimit sets UsersPerCardLimit field to given value.


### GetActions

`func (o *CardAddedDeductedPointsBalancesNotification) GetActions() []AddedDeductedPointsBalancesAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CardAddedDeductedPointsBalancesNotification) GetActionsOk() (*[]AddedDeductedPointsBalancesAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CardAddedDeductedPointsBalancesNotification) SetActions(v []AddedDeductedPointsBalancesAction)`

SetActions sets Actions field to given value.


### GetCurrentPoints

`func (o *CardAddedDeductedPointsBalancesNotification) GetCurrentPoints() float32`

GetCurrentPoints returns the CurrentPoints field if non-nil, zero value otherwise.

### GetCurrentPointsOk

`func (o *CardAddedDeductedPointsBalancesNotification) GetCurrentPointsOk() (*float32, bool)`

GetCurrentPointsOk returns a tuple with the CurrentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPoints

`func (o *CardAddedDeductedPointsBalancesNotification) SetCurrentPoints(v float32)`

SetCurrentPoints sets CurrentPoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


