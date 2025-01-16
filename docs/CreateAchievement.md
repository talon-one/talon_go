# CreateAchievement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The internal name of the achievement used in API requests.  **Note**: The name should start with a letter. This cannot be changed after the achievement has been created.  | 
**Title** | Pointer to **string** | The display name for the achievement in the Campaign Manager. | 
**Description** | Pointer to **string** | A description of the achievement. | 
**Target** | Pointer to **float32** | The required number of actions or the transactional milestone to complete the achievement. | 
**Period** | Pointer to **string** | The relative duration after which the achievement ends and resets for a particular customer profile.  **Note**: The &#x60;period&#x60; does not start when the achievement is created.  The period is a **positive real number** followed by one letter indicating the time unit.  Examples: &#x60;30s&#x60;, &#x60;40m&#x60;, &#x60;1h&#x60;, &#x60;5D&#x60;, &#x60;7W&#x60;, &#x60;10M&#x60;, &#x60;15Y&#x60;.  Available units:  - &#x60;s&#x60;: seconds - &#x60;m&#x60;: minutes - &#x60;h&#x60;: hours - &#x60;D&#x60;: days - &#x60;W&#x60;: weeks - &#x60;M&#x60;: months - &#x60;Y&#x60;: years  You can also round certain units down to the beginning of period and up to the end of period.: - &#x60;_D&#x60; for rounding down days only. Signifies the start of the day. Example: &#x60;30D_D&#x60; - &#x60;_U&#x60; for rounding up days, weeks, months and years. Signifies the end of the day, week, month or year. Example: &#x60;23W_U&#x60;  **Note**: You can either use the round down and round up option or set an absolute period.  | 
**PeriodEndOverride** | Pointer to [**TimePoint**](TimePoint.md) |  | [optional] 
**RecurrencePolicy** | Pointer to **string** | The policy that determines if and how the achievement recurs. - &#x60;no_recurrence&#x60;: The achievement can be completed only once. - &#x60;on_expiration&#x60;: The achievement resets after it expires and becomes available again.  | [optional] 
**ActivationPolicy** | Pointer to **string** | The policy that determines how the achievement starts, ends, or resets. - &#x60;user_action&#x60;: The achievement ends or resets relative to when the customer started the achievement. - &#x60;fixed_schedule&#x60;: The achievement starts, ends, or resets for all customers following a fixed schedule.  | [optional] 
**FixedStartDate** | Pointer to [**time.Time**](time.Time.md) | The achievement&#39;s start date when &#x60;activationPolicy&#x60; is set to &#x60;fixed_schedule&#x60;.  **Note:** It must be an RFC3339 timestamp string.  | [optional] 
**EndDate** | Pointer to [**time.Time**](time.Time.md) | The achievement&#39;s end date. If defined, customers cannot participate in the achievement after this date.  **Note:** It must be an RFC3339 timestamp string.  | [optional] 

## Methods

### GetName

`func (o *CreateAchievement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAchievement) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *CreateAchievement) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *CreateAchievement) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTitle

`func (o *CreateAchievement) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateAchievement) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *CreateAchievement) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *CreateAchievement) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetDescription

`func (o *CreateAchievement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAchievement) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *CreateAchievement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *CreateAchievement) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetTarget

`func (o *CreateAchievement) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CreateAchievement) GetTargetOk() (float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTarget

`func (o *CreateAchievement) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTarget

`func (o *CreateAchievement) SetTarget(v float32)`

SetTarget gets a reference to the given float32 and assigns it to the Target field.

### GetPeriod

`func (o *CreateAchievement) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CreateAchievement) GetPeriodOk() (string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriod

`func (o *CreateAchievement) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriod

`func (o *CreateAchievement) SetPeriod(v string)`

SetPeriod gets a reference to the given string and assigns it to the Period field.

### GetPeriodEndOverride

`func (o *CreateAchievement) GetPeriodEndOverride() TimePoint`

GetPeriodEndOverride returns the PeriodEndOverride field if non-nil, zero value otherwise.

### GetPeriodEndOverrideOk

`func (o *CreateAchievement) GetPeriodEndOverrideOk() (TimePoint, bool)`

GetPeriodEndOverrideOk returns a tuple with the PeriodEndOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodEndOverride

`func (o *CreateAchievement) HasPeriodEndOverride() bool`

HasPeriodEndOverride returns a boolean if a field has been set.

### SetPeriodEndOverride

`func (o *CreateAchievement) SetPeriodEndOverride(v TimePoint)`

SetPeriodEndOverride gets a reference to the given TimePoint and assigns it to the PeriodEndOverride field.

### GetRecurrencePolicy

`func (o *CreateAchievement) GetRecurrencePolicy() string`

GetRecurrencePolicy returns the RecurrencePolicy field if non-nil, zero value otherwise.

### GetRecurrencePolicyOk

`func (o *CreateAchievement) GetRecurrencePolicyOk() (string, bool)`

GetRecurrencePolicyOk returns a tuple with the RecurrencePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecurrencePolicy

`func (o *CreateAchievement) HasRecurrencePolicy() bool`

HasRecurrencePolicy returns a boolean if a field has been set.

### SetRecurrencePolicy

`func (o *CreateAchievement) SetRecurrencePolicy(v string)`

SetRecurrencePolicy gets a reference to the given string and assigns it to the RecurrencePolicy field.

### GetActivationPolicy

`func (o *CreateAchievement) GetActivationPolicy() string`

GetActivationPolicy returns the ActivationPolicy field if non-nil, zero value otherwise.

### GetActivationPolicyOk

`func (o *CreateAchievement) GetActivationPolicyOk() (string, bool)`

GetActivationPolicyOk returns a tuple with the ActivationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivationPolicy

`func (o *CreateAchievement) HasActivationPolicy() bool`

HasActivationPolicy returns a boolean if a field has been set.

### SetActivationPolicy

`func (o *CreateAchievement) SetActivationPolicy(v string)`

SetActivationPolicy gets a reference to the given string and assigns it to the ActivationPolicy field.

### GetFixedStartDate

`func (o *CreateAchievement) GetFixedStartDate() time.Time`

GetFixedStartDate returns the FixedStartDate field if non-nil, zero value otherwise.

### GetFixedStartDateOk

`func (o *CreateAchievement) GetFixedStartDateOk() (time.Time, bool)`

GetFixedStartDateOk returns a tuple with the FixedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFixedStartDate

`func (o *CreateAchievement) HasFixedStartDate() bool`

HasFixedStartDate returns a boolean if a field has been set.

### SetFixedStartDate

`func (o *CreateAchievement) SetFixedStartDate(v time.Time)`

SetFixedStartDate gets a reference to the given time.Time and assigns it to the FixedStartDate field.

### GetEndDate

`func (o *CreateAchievement) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CreateAchievement) GetEndDateOk() (time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndDate

`func (o *CreateAchievement) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDate

`func (o *CreateAchievement) SetEndDate(v time.Time)`

SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


