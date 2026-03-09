# IntegrationHubEventPayloadLoyaltyProfileBasedNotification

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
**CurrentTier** | Pointer to **string** |  | [optional] 
**OldTier** | Pointer to **string** |  | [optional] 
**TierExpirationDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**TimestampOfTierChange** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**PointsRequiredToTheNextTier** | Pointer to **float32** |  | [optional] 
**NextTier** | Pointer to **string** |  | [optional] 

## Methods

### NewIntegrationHubEventPayloadLoyaltyProfileBasedNotification

`func NewIntegrationHubEventPayloadLoyaltyProfileBasedNotification(profileIntegrationID string, loyaltyProgramID int64, subledgerID string, sourceOfEvent string, currentPoints float32, publishedAt time.Time, ) *IntegrationHubEventPayloadLoyaltyProfileBasedNotification`

NewIntegrationHubEventPayloadLoyaltyProfileBasedNotification instantiates a new IntegrationHubEventPayloadLoyaltyProfileBasedNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationHubEventPayloadLoyaltyProfileBasedNotificationWithDefaults

`func NewIntegrationHubEventPayloadLoyaltyProfileBasedNotificationWithDefaults() *IntegrationHubEventPayloadLoyaltyProfileBasedNotification`

NewIntegrationHubEventPayloadLoyaltyProfileBasedNotificationWithDefaults instantiates a new IntegrationHubEventPayloadLoyaltyProfileBasedNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileIntegrationID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetProfileIntegrationID() string`

GetProfileIntegrationID returns the ProfileIntegrationID field if non-nil, zero value otherwise.

### GetProfileIntegrationIDOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetProfileIntegrationIDOk() (*string, bool)`

GetProfileIntegrationIDOk returns a tuple with the ProfileIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) SetProfileIntegrationID(v string)`

SetProfileIntegrationID sets ProfileIntegrationID field to given value.


### GetLoyaltyProgramID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetSubledgerID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetSourceOfEvent

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetSourceOfEvent() string`

GetSourceOfEvent returns the SourceOfEvent field if non-nil, zero value otherwise.

### GetSourceOfEventOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetSourceOfEventOk() (*string, bool)`

GetSourceOfEventOk returns a tuple with the SourceOfEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfEvent

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) SetSourceOfEvent(v string)`

SetSourceOfEvent sets SourceOfEvent field to given value.


### GetEmployeeName

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetEmployeeName() string`

GetEmployeeName returns the EmployeeName field if non-nil, zero value otherwise.

### GetEmployeeNameOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetEmployeeNameOk() (*string, bool)`

GetEmployeeNameOk returns a tuple with the EmployeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) SetEmployeeName(v string)`

SetEmployeeName sets EmployeeName field to given value.

### HasEmployeeName

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) HasEmployeeName() bool`

HasEmployeeName returns a boolean if a field has been set.

### GetUserID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) SetUserID(v int64)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetCurrentPoints

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetCurrentPoints() float32`

GetCurrentPoints returns the CurrentPoints field if non-nil, zero value otherwise.

### GetCurrentPointsOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetCurrentPointsOk() (*float32, bool)`

GetCurrentPointsOk returns a tuple with the CurrentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPoints

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) SetCurrentPoints(v float32)`

SetCurrentPoints sets CurrentPoints field to given value.


### GetActions

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetActions() []IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetActionsOk() (*[]IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) SetActions(v []IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetPublishedAt

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.


### GetCurrentTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetCurrentTier() string`

GetCurrentTier returns the CurrentTier field if non-nil, zero value otherwise.

### GetCurrentTierOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetCurrentTierOk() (*string, bool)`

GetCurrentTierOk returns a tuple with the CurrentTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) SetCurrentTier(v string)`

SetCurrentTier sets CurrentTier field to given value.

### HasCurrentTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) HasCurrentTier() bool`

HasCurrentTier returns a boolean if a field has been set.

### GetOldTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetOldTier() string`

GetOldTier returns the OldTier field if non-nil, zero value otherwise.

### GetOldTierOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetOldTierOk() (*string, bool)`

GetOldTierOk returns a tuple with the OldTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) SetOldTier(v string)`

SetOldTier sets OldTier field to given value.

### HasOldTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) HasOldTier() bool`

HasOldTier returns a boolean if a field has been set.

### GetTierExpirationDate

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetTierExpirationDate() time.Time`

GetTierExpirationDate returns the TierExpirationDate field if non-nil, zero value otherwise.

### GetTierExpirationDateOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetTierExpirationDateOk() (*time.Time, bool)`

GetTierExpirationDateOk returns a tuple with the TierExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierExpirationDate

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) SetTierExpirationDate(v time.Time)`

SetTierExpirationDate sets TierExpirationDate field to given value.

### HasTierExpirationDate

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) HasTierExpirationDate() bool`

HasTierExpirationDate returns a boolean if a field has been set.

### GetTimestampOfTierChange

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetTimestampOfTierChange() time.Time`

GetTimestampOfTierChange returns the TimestampOfTierChange field if non-nil, zero value otherwise.

### GetTimestampOfTierChangeOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetTimestampOfTierChangeOk() (*time.Time, bool)`

GetTimestampOfTierChangeOk returns a tuple with the TimestampOfTierChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOfTierChange

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) SetTimestampOfTierChange(v time.Time)`

SetTimestampOfTierChange sets TimestampOfTierChange field to given value.

### HasTimestampOfTierChange

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) HasTimestampOfTierChange() bool`

HasTimestampOfTierChange returns a boolean if a field has been set.

### GetPointsRequiredToTheNextTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetPointsRequiredToTheNextTier() float32`

GetPointsRequiredToTheNextTier returns the PointsRequiredToTheNextTier field if non-nil, zero value otherwise.

### GetPointsRequiredToTheNextTierOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetPointsRequiredToTheNextTierOk() (*float32, bool)`

GetPointsRequiredToTheNextTierOk returns a tuple with the PointsRequiredToTheNextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsRequiredToTheNextTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) SetPointsRequiredToTheNextTier(v float32)`

SetPointsRequiredToTheNextTier sets PointsRequiredToTheNextTier field to given value.

### HasPointsRequiredToTheNextTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) HasPointsRequiredToTheNextTier() bool`

HasPointsRequiredToTheNextTier returns a boolean if a field has been set.

### GetNextTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetNextTier() string`

GetNextTier returns the NextTier field if non-nil, zero value otherwise.

### GetNextTierOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) GetNextTierOk() (*string, bool)`

GetNextTierOk returns a tuple with the NextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) SetNextTier(v string)`

SetNextTier sets NextTier field to given value.

### HasNextTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedNotification) HasNextTier() bool`

HasNextTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


