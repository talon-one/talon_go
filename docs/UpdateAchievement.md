# UpdateAchievement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description of the achievement. | [optional] 
**Name** | Pointer to **string** | The internal name of the achievement used in API requests.  | [optional] 
**Period** | Pointer to **string** | The relative duration after which the achievement ends and resets for a particular customer profile.  | [optional] 
**PeriodEndOverride** | Pointer to [**TimePoint**](TimePoint.md) |  | [optional] 
**Target** | Pointer to **float32** | The required number of actions or the transactional milestone to complete the achievement. | [optional] 
**Title** | Pointer to **string** | The display name for the achievement in the Campaign Manager. | [optional] 

## Methods

### GetDescription

`func (o *UpdateAchievement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAchievement) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *UpdateAchievement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *UpdateAchievement) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetName

`func (o *UpdateAchievement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAchievement) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *UpdateAchievement) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *UpdateAchievement) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetPeriod

`func (o *UpdateAchievement) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *UpdateAchievement) GetPeriodOk() (string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriod

`func (o *UpdateAchievement) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriod

`func (o *UpdateAchievement) SetPeriod(v string)`

SetPeriod gets a reference to the given string and assigns it to the Period field.

### GetPeriodEndOverride

`func (o *UpdateAchievement) GetPeriodEndOverride() TimePoint`

GetPeriodEndOverride returns the PeriodEndOverride field if non-nil, zero value otherwise.

### GetPeriodEndOverrideOk

`func (o *UpdateAchievement) GetPeriodEndOverrideOk() (TimePoint, bool)`

GetPeriodEndOverrideOk returns a tuple with the PeriodEndOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodEndOverride

`func (o *UpdateAchievement) HasPeriodEndOverride() bool`

HasPeriodEndOverride returns a boolean if a field has been set.

### SetPeriodEndOverride

`func (o *UpdateAchievement) SetPeriodEndOverride(v TimePoint)`

SetPeriodEndOverride gets a reference to the given TimePoint and assigns it to the PeriodEndOverride field.

### GetTarget

`func (o *UpdateAchievement) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *UpdateAchievement) GetTargetOk() (float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTarget

`func (o *UpdateAchievement) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTarget

`func (o *UpdateAchievement) SetTarget(v float32)`

SetTarget gets a reference to the given float32 and assigns it to the Target field.

### GetTitle

`func (o *UpdateAchievement) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateAchievement) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *UpdateAchievement) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *UpdateAchievement) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


