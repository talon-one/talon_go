# IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileIntegrationID** | Pointer to **string** |  | 
**LoyaltyProgramID** | Pointer to **int64** |  | 
**SubledgerID** | Pointer to **string** |  | 
**SourceOfEvent** | Pointer to **string** |  | 
**EmployeeName** | Pointer to **string** |  | [optional] 
**UserID** | Pointer to **int64** |  | [optional] 
**CurrentPoints** | Pointer to **float32** |  | 
**Actions** | Pointer to [**[]IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction**](IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction.md) |  | [optional] 
**PublishedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the event was published. | 

## Methods

### NewIntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification

`func NewIntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification(profileIntegrationID string, loyaltyProgramID int64, subledgerID string, sourceOfEvent string, currentPoints float32, publishedAt time.Time, ) *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification`

NewIntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification instantiates a new IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationWithDefaults

`func NewIntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationWithDefaults() *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification`

NewIntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationWithDefaults instantiates a new IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileIntegrationID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetProfileIntegrationID() string`

GetProfileIntegrationID returns the ProfileIntegrationID field if non-nil, zero value otherwise.

### GetProfileIntegrationIDOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetProfileIntegrationIDOk() (*string, bool)`

GetProfileIntegrationIDOk returns a tuple with the ProfileIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) SetProfileIntegrationID(v string)`

SetProfileIntegrationID sets ProfileIntegrationID field to given value.


### GetLoyaltyProgramID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetSubledgerID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetSourceOfEvent

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetSourceOfEvent() string`

GetSourceOfEvent returns the SourceOfEvent field if non-nil, zero value otherwise.

### GetSourceOfEventOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetSourceOfEventOk() (*string, bool)`

GetSourceOfEventOk returns a tuple with the SourceOfEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfEvent

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) SetSourceOfEvent(v string)`

SetSourceOfEvent sets SourceOfEvent field to given value.


### GetEmployeeName

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetEmployeeName() string`

GetEmployeeName returns the EmployeeName field if non-nil, zero value otherwise.

### GetEmployeeNameOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetEmployeeNameOk() (*string, bool)`

GetEmployeeNameOk returns a tuple with the EmployeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) SetEmployeeName(v string)`

SetEmployeeName sets EmployeeName field to given value.

### HasEmployeeName

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) HasEmployeeName() bool`

HasEmployeeName returns a boolean if a field has been set.

### GetUserID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) SetUserID(v int64)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetCurrentPoints

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetCurrentPoints() float32`

GetCurrentPoints returns the CurrentPoints field if non-nil, zero value otherwise.

### GetCurrentPointsOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetCurrentPointsOk() (*float32, bool)`

GetCurrentPointsOk returns a tuple with the CurrentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPoints

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) SetCurrentPoints(v float32)`

SetCurrentPoints sets CurrentPoints field to given value.


### GetActions

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetActions() []IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetActionsOk() (*[]IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) SetActions(v []IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetPublishedAt

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotification) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


