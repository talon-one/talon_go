# RiskDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**NotificationId** | Pointer to **int64** | The ID of the risk notification rule that flagged this risk. | 
**FeatureDate** | Pointer to **string** | The date of the activity data in which this risk was detected. The anomaly detection pipeline scores complete 24-hour cycles, so this is always the day before the risk was reported, not the reporting date itself.  | 
**GroupKey** | Pointer to **string** | The Application group this risk was detected in. Contains the Application ID, or &#x60;__GLOBAL__&#x60; for metrics that are not grouped by Application.  | 
**ApplicationId** | Pointer to **int64** | The ID of the Application this risk belongs to. Absent for global metrics. | [optional] 
**Status** | Pointer to **string** | The triage lifecycle status of this risk. | 
**Criticality** | Pointer to **string** | The critical classification bucket of this risk. | 
**Entity** | Pointer to **string** | The entity type the risk was detected in. | 
**Activity** | Pointer to **string** | The activity metric the risk was detected in. | 
**TimeFrame** | Pointer to **string** | The rolling time window of the risk evaluation. | 
**ReportedDate** | Pointer to [**time.Time**](time.Time.md) | The time the ML service reported this risk. | 
**AffectedEntityCount** | Pointer to **int64** | The total number of entities affected by this risk. | 
**Description** | Pointer to **string** | Human-readable description of the detected anomaly. | [optional] 
**DiscardReason** | Pointer to **string** | The reason this risk was discarded. Only present on discarded risks. | [optional] 
**StatusComment** | Pointer to **string** | The free-text details of the latest reclassification action: the description for resolving confirmed risks, or the details for discarding risks.  | [optional] 
**StatusChangedBy** | Pointer to **int64** | The ID of the user who performed the latest reclassification action. | [optional] 
**StatusChangedAt** | Pointer to [**time.Time**](time.Time.md) | The time of the latest reclassification action. | [optional] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent update. | 
**AffectedEntities** | Pointer to [**[]RiskAffectedEntityItem**](RiskAffectedEntityItem.md) | The affected entities with the highest severity ratios, in descending order. | 

## Methods

### NewRiskDetail

`func NewRiskDetail(id int64, created time.Time, notificationId int64, featureDate string, groupKey string, status string, criticality string, entity string, activity string, timeFrame string, reportedDate time.Time, affectedEntityCount int64, modified time.Time, affectedEntities []RiskAffectedEntityItem, ) *RiskDetail`

NewRiskDetail instantiates a new RiskDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskDetailWithDefaults

`func NewRiskDetailWithDefaults() *RiskDetail`

NewRiskDetailWithDefaults instantiates a new RiskDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RiskDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskDetail) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *RiskDetail) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RiskDetail) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RiskDetail) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetNotificationId

`func (o *RiskDetail) GetNotificationId() int64`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *RiskDetail) GetNotificationIdOk() (*int64, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *RiskDetail) SetNotificationId(v int64)`

SetNotificationId sets NotificationId field to given value.


### GetFeatureDate

`func (o *RiskDetail) GetFeatureDate() string`

GetFeatureDate returns the FeatureDate field if non-nil, zero value otherwise.

### GetFeatureDateOk

`func (o *RiskDetail) GetFeatureDateOk() (*string, bool)`

GetFeatureDateOk returns a tuple with the FeatureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureDate

`func (o *RiskDetail) SetFeatureDate(v string)`

SetFeatureDate sets FeatureDate field to given value.


### GetGroupKey

`func (o *RiskDetail) GetGroupKey() string`

GetGroupKey returns the GroupKey field if non-nil, zero value otherwise.

### GetGroupKeyOk

`func (o *RiskDetail) GetGroupKeyOk() (*string, bool)`

GetGroupKeyOk returns a tuple with the GroupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupKey

`func (o *RiskDetail) SetGroupKey(v string)`

SetGroupKey sets GroupKey field to given value.


### GetApplicationId

`func (o *RiskDetail) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *RiskDetail) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *RiskDetail) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *RiskDetail) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetStatus

`func (o *RiskDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RiskDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RiskDetail) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCriticality

`func (o *RiskDetail) GetCriticality() string`

GetCriticality returns the Criticality field if non-nil, zero value otherwise.

### GetCriticalityOk

`func (o *RiskDetail) GetCriticalityOk() (*string, bool)`

GetCriticalityOk returns a tuple with the Criticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticality

`func (o *RiskDetail) SetCriticality(v string)`

SetCriticality sets Criticality field to given value.


### GetEntity

`func (o *RiskDetail) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *RiskDetail) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *RiskDetail) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetActivity

`func (o *RiskDetail) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *RiskDetail) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *RiskDetail) SetActivity(v string)`

SetActivity sets Activity field to given value.


### GetTimeFrame

`func (o *RiskDetail) GetTimeFrame() string`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *RiskDetail) GetTimeFrameOk() (*string, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *RiskDetail) SetTimeFrame(v string)`

SetTimeFrame sets TimeFrame field to given value.


### GetReportedDate

`func (o *RiskDetail) GetReportedDate() time.Time`

GetReportedDate returns the ReportedDate field if non-nil, zero value otherwise.

### GetReportedDateOk

`func (o *RiskDetail) GetReportedDateOk() (*time.Time, bool)`

GetReportedDateOk returns a tuple with the ReportedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedDate

`func (o *RiskDetail) SetReportedDate(v time.Time)`

SetReportedDate sets ReportedDate field to given value.


### GetAffectedEntityCount

`func (o *RiskDetail) GetAffectedEntityCount() int64`

GetAffectedEntityCount returns the AffectedEntityCount field if non-nil, zero value otherwise.

### GetAffectedEntityCountOk

`func (o *RiskDetail) GetAffectedEntityCountOk() (*int64, bool)`

GetAffectedEntityCountOk returns a tuple with the AffectedEntityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedEntityCount

`func (o *RiskDetail) SetAffectedEntityCount(v int64)`

SetAffectedEntityCount sets AffectedEntityCount field to given value.


### GetDescription

`func (o *RiskDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RiskDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RiskDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RiskDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscardReason

`func (o *RiskDetail) GetDiscardReason() string`

GetDiscardReason returns the DiscardReason field if non-nil, zero value otherwise.

### GetDiscardReasonOk

`func (o *RiskDetail) GetDiscardReasonOk() (*string, bool)`

GetDiscardReasonOk returns a tuple with the DiscardReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscardReason

`func (o *RiskDetail) SetDiscardReason(v string)`

SetDiscardReason sets DiscardReason field to given value.

### HasDiscardReason

`func (o *RiskDetail) HasDiscardReason() bool`

HasDiscardReason returns a boolean if a field has been set.

### GetStatusComment

`func (o *RiskDetail) GetStatusComment() string`

GetStatusComment returns the StatusComment field if non-nil, zero value otherwise.

### GetStatusCommentOk

`func (o *RiskDetail) GetStatusCommentOk() (*string, bool)`

GetStatusCommentOk returns a tuple with the StatusComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusComment

`func (o *RiskDetail) SetStatusComment(v string)`

SetStatusComment sets StatusComment field to given value.

### HasStatusComment

`func (o *RiskDetail) HasStatusComment() bool`

HasStatusComment returns a boolean if a field has been set.

### GetStatusChangedBy

`func (o *RiskDetail) GetStatusChangedBy() int64`

GetStatusChangedBy returns the StatusChangedBy field if non-nil, zero value otherwise.

### GetStatusChangedByOk

`func (o *RiskDetail) GetStatusChangedByOk() (*int64, bool)`

GetStatusChangedByOk returns a tuple with the StatusChangedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChangedBy

`func (o *RiskDetail) SetStatusChangedBy(v int64)`

SetStatusChangedBy sets StatusChangedBy field to given value.

### HasStatusChangedBy

`func (o *RiskDetail) HasStatusChangedBy() bool`

HasStatusChangedBy returns a boolean if a field has been set.

### GetStatusChangedAt

`func (o *RiskDetail) GetStatusChangedAt() time.Time`

GetStatusChangedAt returns the StatusChangedAt field if non-nil, zero value otherwise.

### GetStatusChangedAtOk

`func (o *RiskDetail) GetStatusChangedAtOk() (*time.Time, bool)`

GetStatusChangedAtOk returns a tuple with the StatusChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChangedAt

`func (o *RiskDetail) SetStatusChangedAt(v time.Time)`

SetStatusChangedAt sets StatusChangedAt field to given value.

### HasStatusChangedAt

`func (o *RiskDetail) HasStatusChangedAt() bool`

HasStatusChangedAt returns a boolean if a field has been set.

### GetModified

`func (o *RiskDetail) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *RiskDetail) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *RiskDetail) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetAffectedEntities

`func (o *RiskDetail) GetAffectedEntities() []RiskAffectedEntityItem`

GetAffectedEntities returns the AffectedEntities field if non-nil, zero value otherwise.

### GetAffectedEntitiesOk

`func (o *RiskDetail) GetAffectedEntitiesOk() (*[]RiskAffectedEntityItem, bool)`

GetAffectedEntitiesOk returns a tuple with the AffectedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedEntities

`func (o *RiskDetail) SetAffectedEntities(v []RiskAffectedEntityItem)`

SetAffectedEntities sets AffectedEntities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


