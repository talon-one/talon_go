# IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileIntegrationID** | Pointer to **string** |  | 
**LoyaltyProgramID** | Pointer to **int64** |  | 
**SubledgerID** | Pointer to **string** |  | 
**SourceOfEvent** | Pointer to **string** |  | 
**CurrentTier** | Pointer to **string** |  | [optional] 
**CurrentPoints** | Pointer to **float32** |  | 
**OldTier** | Pointer to **string** |  | [optional] 
**PointsRequiredToTheNextTier** | Pointer to **float32** |  | [optional] 
**NextTier** | Pointer to **string** |  | [optional] 
**TierExpirationDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**TimestampOfTierChange** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**PublishedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the event was published. | 

## Methods

### NewIntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification

`func NewIntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification(profileIntegrationID string, loyaltyProgramID int64, subledgerID string, sourceOfEvent string, currentPoints float32, publishedAt time.Time, ) *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification`

NewIntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification instantiates a new IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotificationWithDefaults

`func NewIntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotificationWithDefaults() *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification`

NewIntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotificationWithDefaults instantiates a new IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileIntegrationID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetProfileIntegrationID() string`

GetProfileIntegrationID returns the ProfileIntegrationID field if non-nil, zero value otherwise.

### GetProfileIntegrationIDOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetProfileIntegrationIDOk() (*string, bool)`

GetProfileIntegrationIDOk returns a tuple with the ProfileIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetProfileIntegrationID(v string)`

SetProfileIntegrationID sets ProfileIntegrationID field to given value.


### GetLoyaltyProgramID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetSubledgerID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetSourceOfEvent

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetSourceOfEvent() string`

GetSourceOfEvent returns the SourceOfEvent field if non-nil, zero value otherwise.

### GetSourceOfEventOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetSourceOfEventOk() (*string, bool)`

GetSourceOfEventOk returns a tuple with the SourceOfEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfEvent

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetSourceOfEvent(v string)`

SetSourceOfEvent sets SourceOfEvent field to given value.


### GetCurrentTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetCurrentTier() string`

GetCurrentTier returns the CurrentTier field if non-nil, zero value otherwise.

### GetCurrentTierOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetCurrentTierOk() (*string, bool)`

GetCurrentTierOk returns a tuple with the CurrentTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetCurrentTier(v string)`

SetCurrentTier sets CurrentTier field to given value.

### HasCurrentTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) HasCurrentTier() bool`

HasCurrentTier returns a boolean if a field has been set.

### GetCurrentPoints

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetCurrentPoints() float32`

GetCurrentPoints returns the CurrentPoints field if non-nil, zero value otherwise.

### GetCurrentPointsOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetCurrentPointsOk() (*float32, bool)`

GetCurrentPointsOk returns a tuple with the CurrentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPoints

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetCurrentPoints(v float32)`

SetCurrentPoints sets CurrentPoints field to given value.


### GetOldTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetOldTier() string`

GetOldTier returns the OldTier field if non-nil, zero value otherwise.

### GetOldTierOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetOldTierOk() (*string, bool)`

GetOldTierOk returns a tuple with the OldTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetOldTier(v string)`

SetOldTier sets OldTier field to given value.

### HasOldTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) HasOldTier() bool`

HasOldTier returns a boolean if a field has been set.

### GetPointsRequiredToTheNextTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetPointsRequiredToTheNextTier() float32`

GetPointsRequiredToTheNextTier returns the PointsRequiredToTheNextTier field if non-nil, zero value otherwise.

### GetPointsRequiredToTheNextTierOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetPointsRequiredToTheNextTierOk() (*float32, bool)`

GetPointsRequiredToTheNextTierOk returns a tuple with the PointsRequiredToTheNextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsRequiredToTheNextTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetPointsRequiredToTheNextTier(v float32)`

SetPointsRequiredToTheNextTier sets PointsRequiredToTheNextTier field to given value.

### HasPointsRequiredToTheNextTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) HasPointsRequiredToTheNextTier() bool`

HasPointsRequiredToTheNextTier returns a boolean if a field has been set.

### GetNextTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetNextTier() string`

GetNextTier returns the NextTier field if non-nil, zero value otherwise.

### GetNextTierOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetNextTierOk() (*string, bool)`

GetNextTierOk returns a tuple with the NextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetNextTier(v string)`

SetNextTier sets NextTier field to given value.

### HasNextTier

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) HasNextTier() bool`

HasNextTier returns a boolean if a field has been set.

### GetTierExpirationDate

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetTierExpirationDate() time.Time`

GetTierExpirationDate returns the TierExpirationDate field if non-nil, zero value otherwise.

### GetTierExpirationDateOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetTierExpirationDateOk() (*time.Time, bool)`

GetTierExpirationDateOk returns a tuple with the TierExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierExpirationDate

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetTierExpirationDate(v time.Time)`

SetTierExpirationDate sets TierExpirationDate field to given value.

### HasTierExpirationDate

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) HasTierExpirationDate() bool`

HasTierExpirationDate returns a boolean if a field has been set.

### GetTimestampOfTierChange

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetTimestampOfTierChange() time.Time`

GetTimestampOfTierChange returns the TimestampOfTierChange field if non-nil, zero value otherwise.

### GetTimestampOfTierChangeOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetTimestampOfTierChangeOk() (*time.Time, bool)`

GetTimestampOfTierChangeOk returns a tuple with the TimestampOfTierChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOfTierChange

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetTimestampOfTierChange(v time.Time)`

SetTimestampOfTierChange sets TimestampOfTierChange field to given value.

### HasTimestampOfTierChange

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) HasTimestampOfTierChange() bool`

HasTimestampOfTierChange returns a boolean if a field has been set.

### GetPublishedAt

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


