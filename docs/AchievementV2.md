# AchievementV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Name** | Pointer to **string** | The internal name of the achievement used in API requests.  **Note**: The name should start with a letter. This cannot be changed after the achievement has been created.  | 
**Title** | Pointer to **string** | The display name for the achievement in the Campaign Manager. | 
**Description** | Pointer to **string** | A description of the achievement. | 
**Target** | Pointer to **float32** | The required number of actions or the transactional milestone to complete the achievement. | 
**Period** | Pointer to **string** | The relative duration after which the achievement ends and resets for a particular customer profile.  **Note**: The &#x60;period&#x60; does not start when the achievement is created.  The period is a **positive real number** followed by one letter indicating the time unit.  Examples: &#x60;30s&#x60;, &#x60;40m&#x60;, &#x60;1h&#x60;, &#x60;5D&#x60;, &#x60;7W&#x60;, &#x60;10M&#x60;, &#x60;15Y&#x60;.  Available units:  - &#x60;s&#x60;: seconds - &#x60;m&#x60;: minutes - &#x60;h&#x60;: hours - &#x60;D&#x60;: days - &#x60;W&#x60;: weeks - &#x60;M&#x60;: months - &#x60;Y&#x60;: years  You can also round certain units down to the beginning of period and up to the end of period.: - &#x60;_D&#x60; for rounding down days only. Signifies the start of the day. Example: &#x60;30D_D&#x60; - &#x60;_U&#x60; for rounding up days, weeks, months and years. Signifies the end of the day, week, month or year. Example: &#x60;23W_U&#x60;  **Note**: You can either use the round down and round up option or set an absolute period.  | [optional] 
**RecurrencePolicy** | Pointer to **string** | The policy that determines if and how the achievement recurs. - &#x60;no_recurrence&#x60;: The achievement can be completed only once. - &#x60;on_expiration&#x60;: The achievement resets after it expires and becomes available again. - &#x60;on_completion&#x60;: When the customer progress status reaches &#x60;completed&#x60;, the achievement resets and becomes available again.  | 
**ActivationPolicy** | Pointer to **string** | The policy that determines how the achievement starts, ends, or resets. - &#x60;user_action&#x60;: The achievement ends or resets relative to when the customer started the achievement. - &#x60;fixed_schedule&#x60;: The achievement starts, ends, or resets for all customers following a fixed schedule.  | 
**FixedStartDate** | Pointer to [**time.Time**](time.Time.md) | The achievement&#39;s start date when &#x60;activationPolicy&#x60; is set to &#x60;fixed_schedule&#x60;.  **Note:** It must be an RFC3339 timestamp string.  | [optional] 
**EndDate** | Pointer to [**time.Time**](time.Time.md) | The achievement&#39;s end date. If defined, customers cannot participate in the achievement after this date.  **Note:** It must be an RFC3339 timestamp string.  | [optional] 
**AllowRollbackAfterCompletion** | Pointer to **bool** | When &#x60;true&#x60;, customer progress can be rolled back in completed achievements. | [optional] 
**Sandbox** | Pointer to **bool** | Indicates if this achievement is a live or sandbox achievement. Achievements of a given type can only be connected to Applications of the same type. | 
**SubscribedApplications** | Pointer to **[]int64** | A list containing the IDs of all applications that are subscribed to A list containing the IDs of all Applications that are connected to this achievement. | 
**Timezone** | Pointer to **string** | A string containing an IANA timezone descriptor. | 
**UserId** | Pointer to **int64** | The ID of the user that created this achievement. | 
**CreatedBy** | Pointer to **string** | Name of the user that created the achievement.  **Note**: This is not available if the user has been deleted.  | [optional] 
**HasProgress** | Pointer to **bool** | Indicates if a customer has made progress in the achievement. | [optional] 
**Status** | Pointer to **string** | The status of the achievement. | [optional] 

## Methods

### NewAchievementV2

`func NewAchievementV2(id int64, created time.Time, name string, title string, description string, target float32, recurrencePolicy string, activationPolicy string, sandbox bool, subscribedApplications []int64, timezone string, userId int64, ) *AchievementV2`

NewAchievementV2 instantiates a new AchievementV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchievementV2WithDefaults

`func NewAchievementV2WithDefaults() *AchievementV2`

NewAchievementV2WithDefaults instantiates a new AchievementV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AchievementV2) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AchievementV2) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AchievementV2) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *AchievementV2) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AchievementV2) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AchievementV2) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetName

`func (o *AchievementV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AchievementV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AchievementV2) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *AchievementV2) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AchievementV2) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AchievementV2) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *AchievementV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AchievementV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AchievementV2) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTarget

`func (o *AchievementV2) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AchievementV2) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AchievementV2) SetTarget(v float32)`

SetTarget sets Target field to given value.


### GetPeriod

`func (o *AchievementV2) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *AchievementV2) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *AchievementV2) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *AchievementV2) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetRecurrencePolicy

`func (o *AchievementV2) GetRecurrencePolicy() string`

GetRecurrencePolicy returns the RecurrencePolicy field if non-nil, zero value otherwise.

### GetRecurrencePolicyOk

`func (o *AchievementV2) GetRecurrencePolicyOk() (*string, bool)`

GetRecurrencePolicyOk returns a tuple with the RecurrencePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrencePolicy

`func (o *AchievementV2) SetRecurrencePolicy(v string)`

SetRecurrencePolicy sets RecurrencePolicy field to given value.


### GetActivationPolicy

`func (o *AchievementV2) GetActivationPolicy() string`

GetActivationPolicy returns the ActivationPolicy field if non-nil, zero value otherwise.

### GetActivationPolicyOk

`func (o *AchievementV2) GetActivationPolicyOk() (*string, bool)`

GetActivationPolicyOk returns a tuple with the ActivationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationPolicy

`func (o *AchievementV2) SetActivationPolicy(v string)`

SetActivationPolicy sets ActivationPolicy field to given value.


### GetFixedStartDate

`func (o *AchievementV2) GetFixedStartDate() time.Time`

GetFixedStartDate returns the FixedStartDate field if non-nil, zero value otherwise.

### GetFixedStartDateOk

`func (o *AchievementV2) GetFixedStartDateOk() (*time.Time, bool)`

GetFixedStartDateOk returns a tuple with the FixedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedStartDate

`func (o *AchievementV2) SetFixedStartDate(v time.Time)`

SetFixedStartDate sets FixedStartDate field to given value.

### HasFixedStartDate

`func (o *AchievementV2) HasFixedStartDate() bool`

HasFixedStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *AchievementV2) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AchievementV2) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AchievementV2) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AchievementV2) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetAllowRollbackAfterCompletion

`func (o *AchievementV2) GetAllowRollbackAfterCompletion() bool`

GetAllowRollbackAfterCompletion returns the AllowRollbackAfterCompletion field if non-nil, zero value otherwise.

### GetAllowRollbackAfterCompletionOk

`func (o *AchievementV2) GetAllowRollbackAfterCompletionOk() (*bool, bool)`

GetAllowRollbackAfterCompletionOk returns a tuple with the AllowRollbackAfterCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRollbackAfterCompletion

`func (o *AchievementV2) SetAllowRollbackAfterCompletion(v bool)`

SetAllowRollbackAfterCompletion sets AllowRollbackAfterCompletion field to given value.

### HasAllowRollbackAfterCompletion

`func (o *AchievementV2) HasAllowRollbackAfterCompletion() bool`

HasAllowRollbackAfterCompletion returns a boolean if a field has been set.

### GetSandbox

`func (o *AchievementV2) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *AchievementV2) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *AchievementV2) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.


### GetSubscribedApplications

`func (o *AchievementV2) GetSubscribedApplications() []int64`

GetSubscribedApplications returns the SubscribedApplications field if non-nil, zero value otherwise.

### GetSubscribedApplicationsOk

`func (o *AchievementV2) GetSubscribedApplicationsOk() (*[]int64, bool)`

GetSubscribedApplicationsOk returns a tuple with the SubscribedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedApplications

`func (o *AchievementV2) SetSubscribedApplications(v []int64)`

SetSubscribedApplications sets SubscribedApplications field to given value.


### GetTimezone

`func (o *AchievementV2) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AchievementV2) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AchievementV2) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetUserId

`func (o *AchievementV2) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AchievementV2) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AchievementV2) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetCreatedBy

`func (o *AchievementV2) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AchievementV2) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AchievementV2) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AchievementV2) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetHasProgress

`func (o *AchievementV2) GetHasProgress() bool`

GetHasProgress returns the HasProgress field if non-nil, zero value otherwise.

### GetHasProgressOk

`func (o *AchievementV2) GetHasProgressOk() (*bool, bool)`

GetHasProgressOk returns a tuple with the HasProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProgress

`func (o *AchievementV2) SetHasProgress(v bool)`

SetHasProgress sets HasProgress field to given value.

### HasHasProgress

`func (o *AchievementV2) HasHasProgress() bool`

HasHasProgress returns a boolean if a field has been set.

### GetStatus

`func (o *AchievementV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AchievementV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AchievementV2) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AchievementV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


