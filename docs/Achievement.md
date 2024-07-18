# Achievement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Name** | Pointer to **string** | The internal name of the achievement used in API requests.  **Note**: The name should start with a letter. This cannot be changed after the achievement has been created.  | 
**Title** | Pointer to **string** | The display name for the achievement in the Campaign Manager. | 
**Description** | Pointer to **string** | A description of the achievement. | 
**Target** | Pointer to **float32** | The required number of actions or the transactional milestone to complete the achievement. | 
**Period** | Pointer to **string** | The relative duration after which the achievement ends and resets for a particular customer profile.  **Note**: The &#x60;period&#x60; does not start when the achievement is created.  The period is a **positive real number** followed by one letter indicating the time unit.  Examples: &#x60;30s&#x60;, &#x60;40m&#x60;, &#x60;1h&#x60;, &#x60;5D&#x60;, &#x60;7W&#x60;, &#x60;10M&#x60;, &#x60;15Y&#x60;.  Available units:  - &#x60;s&#x60;: seconds - &#x60;m&#x60;: minutes - &#x60;h&#x60;: hours - &#x60;D&#x60;: days - &#x60;W&#x60;: weeks - &#x60;M&#x60;: months - &#x60;Y&#x60;: years  You can also round certain units down to the beginning of period and up to the end of period.: - &#x60;_D&#x60; for rounding down days only. Signifies the start of the day. Example: &#x60;30D_D&#x60; - &#x60;_U&#x60; for rounding up days, weeks, months and years. Signifies the end of the day, week, month or year. Example: &#x60;23W_U&#x60;  **Note**: You can either use the round down and round up option or set an absolute period.  | 
**PeriodEndOverride** | Pointer to [**TimePoint**](TimePoint.md) |  | [optional] 
**CampaignId** | Pointer to **int32** | ID of the campaign, to which the achievement belongs to | 
**UserId** | Pointer to **int32** | ID of the user that created this achievement. | 
**CreatedBy** | Pointer to **string** | Name of the user that created the achievement.  **Note**: This is not available if the user has been deleted.  | 
**HasProgress** | Pointer to **bool** | Indicates if a customer has made progress in the achievement. | [optional] 

## Methods

### GetId

`func (o *Achievement) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Achievement) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Achievement) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Achievement) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *Achievement) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Achievement) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Achievement) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Achievement) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetName

`func (o *Achievement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Achievement) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Achievement) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Achievement) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTitle

`func (o *Achievement) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Achievement) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *Achievement) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *Achievement) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetDescription

`func (o *Achievement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Achievement) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Achievement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Achievement) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetTarget

`func (o *Achievement) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Achievement) GetTargetOk() (float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTarget

`func (o *Achievement) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTarget

`func (o *Achievement) SetTarget(v float32)`

SetTarget gets a reference to the given float32 and assigns it to the Target field.

### GetPeriod

`func (o *Achievement) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *Achievement) GetPeriodOk() (string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriod

`func (o *Achievement) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriod

`func (o *Achievement) SetPeriod(v string)`

SetPeriod gets a reference to the given string and assigns it to the Period field.

### GetPeriodEndOverride

`func (o *Achievement) GetPeriodEndOverride() TimePoint`

GetPeriodEndOverride returns the PeriodEndOverride field if non-nil, zero value otherwise.

### GetPeriodEndOverrideOk

`func (o *Achievement) GetPeriodEndOverrideOk() (TimePoint, bool)`

GetPeriodEndOverrideOk returns a tuple with the PeriodEndOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodEndOverride

`func (o *Achievement) HasPeriodEndOverride() bool`

HasPeriodEndOverride returns a boolean if a field has been set.

### SetPeriodEndOverride

`func (o *Achievement) SetPeriodEndOverride(v TimePoint)`

SetPeriodEndOverride gets a reference to the given TimePoint and assigns it to the PeriodEndOverride field.

### GetCampaignId

`func (o *Achievement) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *Achievement) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *Achievement) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *Achievement) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.

### GetUserId

`func (o *Achievement) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Achievement) GetUserIdOk() (int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *Achievement) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *Achievement) SetUserId(v int32)`

SetUserId gets a reference to the given int32 and assigns it to the UserId field.

### GetCreatedBy

`func (o *Achievement) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Achievement) GetCreatedByOk() (string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *Achievement) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *Achievement) SetCreatedBy(v string)`

SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.

### GetHasProgress

`func (o *Achievement) GetHasProgress() bool`

GetHasProgress returns the HasProgress field if non-nil, zero value otherwise.

### GetHasProgressOk

`func (o *Achievement) GetHasProgressOk() (bool, bool)`

GetHasProgressOk returns a tuple with the HasProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasProgress

`func (o *Achievement) HasHasProgress() bool`

HasHasProgress returns a boolean if a field has been set.

### SetHasProgress

`func (o *Achievement) SetHasProgress(v bool)`

SetHasProgress gets a reference to the given bool and assigns it to the HasProgress field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


