# AchievementBaseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The internal name of the achievement used in API requests.  **Note**: The name should start with a letter. This cannot be changed after the achievement has been created.  | [optional] 
**Title** | Pointer to **string** | The display name for the achievement in the Campaign Manager. | [optional] 
**Description** | Pointer to **string** | A description of the achievement. | [optional] 
**Target** | Pointer to **float32** | The required number of actions or the transactional milestone to complete the achievement. | [optional] 
**Period** | Pointer to **string** | The relative duration after which the achievement ends and resets for a particular customer profile.  **Note**: The &#x60;period&#x60; does not start when the achievement is created.  The period is a **positive real number** followed by one letter indicating the time unit.  Examples: &#x60;30s&#x60;, &#x60;40m&#x60;, &#x60;1h&#x60;, &#x60;5D&#x60;, &#x60;7W&#x60;, &#x60;10M&#x60;, &#x60;15Y&#x60;.  Available units:  - &#x60;s&#x60;: seconds - &#x60;m&#x60;: minutes - &#x60;h&#x60;: hours - &#x60;D&#x60;: days - &#x60;W&#x60;: weeks - &#x60;M&#x60;: months - &#x60;Y&#x60;: years  You can also round certain units down to the beginning of period and up to the end of period.: - &#x60;_D&#x60; for rounding down days only. Signifies the start of the day. Example: &#x60;30D_D&#x60; - &#x60;_U&#x60; for rounding up days, weeks, months and years. Signifies the end of the day, week, month or year. Example: &#x60;23W_U&#x60;  **Note**: You can either use the round down and round up option or set an absolute period.  | [optional] 
**RecurrencePolicy** | Pointer to **string** | The policy that determines if and how the achievement recurs. - &#x60;no_recurrence&#x60;: The achievement can be completed only once. - &#x60;on_expiration&#x60;: The achievement resets after it expires and becomes available again. - &#x60;on_completion&#x60;: When the customer progress status reaches &#x60;completed&#x60;, the achievement resets and becomes available again.  | [optional] 
**ActivationPolicy** | Pointer to **string** | The policy that determines how the achievement starts, ends, or resets. - &#x60;user_action&#x60;: The achievement ends or resets relative to when the customer started the achievement. - &#x60;fixed_schedule&#x60;: The achievement starts, ends, or resets for all customers following a fixed schedule.  | [optional] 
**FixedStartDate** | Pointer to [**time.Time**](time.Time.md) | The achievement&#39;s start date when &#x60;activationPolicy&#x60; is set to &#x60;fixed_schedule&#x60;.  **Note:** It must be an RFC3339 timestamp string.  | [optional] 
**EndDate** | Pointer to [**time.Time**](time.Time.md) | The achievement&#39;s end date. If defined, customers cannot participate in the achievement after this date.  **Note:** It must be an RFC3339 timestamp string.  | [optional] 
**AllowRollbackAfterCompletion** | Pointer to **bool** | When &#x60;true&#x60;, customer progress can be rolled back in completed achievements. | [optional] 
**Sandbox** | Pointer to **bool** | Indicates if this achievement is a live or sandbox achievement. Achievements of a given type can only be connected to Applications of the same type. | [optional] 
**SubscribedApplications** | Pointer to **[]int64** | A list containing the IDs of all applications that are subscribed to A list containing the IDs of all Applications that are connected to this achievement. | [optional] 
**Timezone** | Pointer to **string** | A string containing an IANA timezone descriptor. | [optional] 

## Methods

### NewAchievementBaseV2

`func NewAchievementBaseV2() *AchievementBaseV2`

NewAchievementBaseV2 instantiates a new AchievementBaseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchievementBaseV2WithDefaults

`func NewAchievementBaseV2WithDefaults() *AchievementBaseV2`

NewAchievementBaseV2WithDefaults instantiates a new AchievementBaseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AchievementBaseV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AchievementBaseV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AchievementBaseV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AchievementBaseV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *AchievementBaseV2) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AchievementBaseV2) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AchievementBaseV2) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AchievementBaseV2) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *AchievementBaseV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AchievementBaseV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AchievementBaseV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AchievementBaseV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTarget

`func (o *AchievementBaseV2) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AchievementBaseV2) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AchievementBaseV2) SetTarget(v float32)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AchievementBaseV2) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetPeriod

`func (o *AchievementBaseV2) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *AchievementBaseV2) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *AchievementBaseV2) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *AchievementBaseV2) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetRecurrencePolicy

`func (o *AchievementBaseV2) GetRecurrencePolicy() string`

GetRecurrencePolicy returns the RecurrencePolicy field if non-nil, zero value otherwise.

### GetRecurrencePolicyOk

`func (o *AchievementBaseV2) GetRecurrencePolicyOk() (*string, bool)`

GetRecurrencePolicyOk returns a tuple with the RecurrencePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrencePolicy

`func (o *AchievementBaseV2) SetRecurrencePolicy(v string)`

SetRecurrencePolicy sets RecurrencePolicy field to given value.

### HasRecurrencePolicy

`func (o *AchievementBaseV2) HasRecurrencePolicy() bool`

HasRecurrencePolicy returns a boolean if a field has been set.

### GetActivationPolicy

`func (o *AchievementBaseV2) GetActivationPolicy() string`

GetActivationPolicy returns the ActivationPolicy field if non-nil, zero value otherwise.

### GetActivationPolicyOk

`func (o *AchievementBaseV2) GetActivationPolicyOk() (*string, bool)`

GetActivationPolicyOk returns a tuple with the ActivationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationPolicy

`func (o *AchievementBaseV2) SetActivationPolicy(v string)`

SetActivationPolicy sets ActivationPolicy field to given value.

### HasActivationPolicy

`func (o *AchievementBaseV2) HasActivationPolicy() bool`

HasActivationPolicy returns a boolean if a field has been set.

### GetFixedStartDate

`func (o *AchievementBaseV2) GetFixedStartDate() time.Time`

GetFixedStartDate returns the FixedStartDate field if non-nil, zero value otherwise.

### GetFixedStartDateOk

`func (o *AchievementBaseV2) GetFixedStartDateOk() (*time.Time, bool)`

GetFixedStartDateOk returns a tuple with the FixedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedStartDate

`func (o *AchievementBaseV2) SetFixedStartDate(v time.Time)`

SetFixedStartDate sets FixedStartDate field to given value.

### HasFixedStartDate

`func (o *AchievementBaseV2) HasFixedStartDate() bool`

HasFixedStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *AchievementBaseV2) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AchievementBaseV2) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AchievementBaseV2) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AchievementBaseV2) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetAllowRollbackAfterCompletion

`func (o *AchievementBaseV2) GetAllowRollbackAfterCompletion() bool`

GetAllowRollbackAfterCompletion returns the AllowRollbackAfterCompletion field if non-nil, zero value otherwise.

### GetAllowRollbackAfterCompletionOk

`func (o *AchievementBaseV2) GetAllowRollbackAfterCompletionOk() (*bool, bool)`

GetAllowRollbackAfterCompletionOk returns a tuple with the AllowRollbackAfterCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRollbackAfterCompletion

`func (o *AchievementBaseV2) SetAllowRollbackAfterCompletion(v bool)`

SetAllowRollbackAfterCompletion sets AllowRollbackAfterCompletion field to given value.

### HasAllowRollbackAfterCompletion

`func (o *AchievementBaseV2) HasAllowRollbackAfterCompletion() bool`

HasAllowRollbackAfterCompletion returns a boolean if a field has been set.

### GetSandbox

`func (o *AchievementBaseV2) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *AchievementBaseV2) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *AchievementBaseV2) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *AchievementBaseV2) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### GetSubscribedApplications

`func (o *AchievementBaseV2) GetSubscribedApplications() []int64`

GetSubscribedApplications returns the SubscribedApplications field if non-nil, zero value otherwise.

### GetSubscribedApplicationsOk

`func (o *AchievementBaseV2) GetSubscribedApplicationsOk() (*[]int64, bool)`

GetSubscribedApplicationsOk returns a tuple with the SubscribedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedApplications

`func (o *AchievementBaseV2) SetSubscribedApplications(v []int64)`

SetSubscribedApplications sets SubscribedApplications field to given value.

### HasSubscribedApplications

`func (o *AchievementBaseV2) HasSubscribedApplications() bool`

HasSubscribedApplications returns a boolean if a field has been set.

### GetTimezone

`func (o *AchievementBaseV2) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AchievementBaseV2) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AchievementBaseV2) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *AchievementBaseV2) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


