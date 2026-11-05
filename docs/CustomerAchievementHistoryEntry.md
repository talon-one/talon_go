# CustomerAchievementHistoryEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The status of the achievement. | 
**Progress** | Pointer to **float32** | The current progress of the customer in the achievement. | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which the customer started the achievement. | [optional] 
**CompletionDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the customer completed the achievement. | [optional] 
**EndDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the achievement ends and resets for the customer. | [optional] 
**ProgressTrackerId** | Pointer to **int64** | The ID of the customer&#39;s progress tracker for this achievement. | 

## Methods

### NewCustomerAchievementHistoryEntry

`func NewCustomerAchievementHistoryEntry(status string, progress float32, progressTrackerId int64, ) *CustomerAchievementHistoryEntry`

NewCustomerAchievementHistoryEntry instantiates a new CustomerAchievementHistoryEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAchievementHistoryEntryWithDefaults

`func NewCustomerAchievementHistoryEntryWithDefaults() *CustomerAchievementHistoryEntry`

NewCustomerAchievementHistoryEntryWithDefaults instantiates a new CustomerAchievementHistoryEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CustomerAchievementHistoryEntry) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerAchievementHistoryEntry) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerAchievementHistoryEntry) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetProgress

`func (o *CustomerAchievementHistoryEntry) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *CustomerAchievementHistoryEntry) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *CustomerAchievementHistoryEntry) SetProgress(v float32)`

SetProgress sets Progress field to given value.


### GetStartDate

`func (o *CustomerAchievementHistoryEntry) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CustomerAchievementHistoryEntry) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CustomerAchievementHistoryEntry) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CustomerAchievementHistoryEntry) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetCompletionDate

`func (o *CustomerAchievementHistoryEntry) GetCompletionDate() time.Time`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *CustomerAchievementHistoryEntry) GetCompletionDateOk() (*time.Time, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *CustomerAchievementHistoryEntry) SetCompletionDate(v time.Time)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *CustomerAchievementHistoryEntry) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### GetEndDate

`func (o *CustomerAchievementHistoryEntry) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CustomerAchievementHistoryEntry) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CustomerAchievementHistoryEntry) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CustomerAchievementHistoryEntry) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetProgressTrackerId

`func (o *CustomerAchievementHistoryEntry) GetProgressTrackerId() int64`

GetProgressTrackerId returns the ProgressTrackerId field if non-nil, zero value otherwise.

### GetProgressTrackerIdOk

`func (o *CustomerAchievementHistoryEntry) GetProgressTrackerIdOk() (*int64, bool)`

GetProgressTrackerIdOk returns a tuple with the ProgressTrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTrackerId

`func (o *CustomerAchievementHistoryEntry) SetProgressTrackerId(v int64)`

SetProgressTrackerId sets ProgressTrackerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


