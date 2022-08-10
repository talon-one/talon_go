# UpdateCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A user-facing name for this campaign. | 
**Description** | Pointer to **string** | A detailed description of the campaign. | [optional] 
**StartTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the campaign will become active. | [optional] 
**EndTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the campaign will become inactive. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this campaign. | [optional] 
**State** | Pointer to **string** | A disabled or archived campaign is not evaluated for rules or coupons.  | [optional] [default to STATE_ENABLED]
**ActiveRulesetId** | Pointer to **int32** | ID of Ruleset this campaign applies on customer session evaluation. | [optional] 
**Tags** | Pointer to **[]string** | A list of tags for the campaign. | 
**Features** | Pointer to **[]string** | A list of features for the campaign. | 
**CouponSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**ReferralSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**Limits** | Pointer to [**[]LimitConfig**](LimitConfig.md) | The set of limits that will operate for this campaign. | 
**CampaignGroups** | Pointer to **[]int32** | The IDs of the campaign groups that own this entity. | [optional] 

## Methods

### GetName

`func (o *UpdateCampaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCampaign) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *UpdateCampaign) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *UpdateCampaign) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *UpdateCampaign) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCampaign) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *UpdateCampaign) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *UpdateCampaign) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetStartTime

`func (o *UpdateCampaign) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UpdateCampaign) GetStartTimeOk() (time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartTime

`func (o *UpdateCampaign) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTime

`func (o *UpdateCampaign) SetStartTime(v time.Time)`

SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.

### GetEndTime

`func (o *UpdateCampaign) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UpdateCampaign) GetEndTimeOk() (time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndTime

`func (o *UpdateCampaign) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTime

`func (o *UpdateCampaign) SetEndTime(v time.Time)`

SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.

### GetAttributes

`func (o *UpdateCampaign) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateCampaign) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *UpdateCampaign) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *UpdateCampaign) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetState

`func (o *UpdateCampaign) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateCampaign) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *UpdateCampaign) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *UpdateCampaign) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetActiveRulesetId

`func (o *UpdateCampaign) GetActiveRulesetId() int32`

GetActiveRulesetId returns the ActiveRulesetId field if non-nil, zero value otherwise.

### GetActiveRulesetIdOk

`func (o *UpdateCampaign) GetActiveRulesetIdOk() (int32, bool)`

GetActiveRulesetIdOk returns a tuple with the ActiveRulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveRulesetId

`func (o *UpdateCampaign) HasActiveRulesetId() bool`

HasActiveRulesetId returns a boolean if a field has been set.

### SetActiveRulesetId

`func (o *UpdateCampaign) SetActiveRulesetId(v int32)`

SetActiveRulesetId gets a reference to the given int32 and assigns it to the ActiveRulesetId field.

### GetTags

`func (o *UpdateCampaign) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateCampaign) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *UpdateCampaign) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *UpdateCampaign) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.

### GetFeatures

`func (o *UpdateCampaign) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *UpdateCampaign) GetFeaturesOk() ([]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatures

`func (o *UpdateCampaign) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeatures

`func (o *UpdateCampaign) SetFeatures(v []string)`

SetFeatures gets a reference to the given []string and assigns it to the Features field.

### GetCouponSettings

`func (o *UpdateCampaign) GetCouponSettings() CodeGeneratorSettings`

GetCouponSettings returns the CouponSettings field if non-nil, zero value otherwise.

### GetCouponSettingsOk

`func (o *UpdateCampaign) GetCouponSettingsOk() (CodeGeneratorSettings, bool)`

GetCouponSettingsOk returns a tuple with the CouponSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponSettings

`func (o *UpdateCampaign) HasCouponSettings() bool`

HasCouponSettings returns a boolean if a field has been set.

### SetCouponSettings

`func (o *UpdateCampaign) SetCouponSettings(v CodeGeneratorSettings)`

SetCouponSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CouponSettings field.

### GetReferralSettings

`func (o *UpdateCampaign) GetReferralSettings() CodeGeneratorSettings`

GetReferralSettings returns the ReferralSettings field if non-nil, zero value otherwise.

### GetReferralSettingsOk

`func (o *UpdateCampaign) GetReferralSettingsOk() (CodeGeneratorSettings, bool)`

GetReferralSettingsOk returns a tuple with the ReferralSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralSettings

`func (o *UpdateCampaign) HasReferralSettings() bool`

HasReferralSettings returns a boolean if a field has been set.

### SetReferralSettings

`func (o *UpdateCampaign) SetReferralSettings(v CodeGeneratorSettings)`

SetReferralSettings gets a reference to the given CodeGeneratorSettings and assigns it to the ReferralSettings field.

### GetLimits

`func (o *UpdateCampaign) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *UpdateCampaign) GetLimitsOk() ([]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimits

`func (o *UpdateCampaign) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimits

`func (o *UpdateCampaign) SetLimits(v []LimitConfig)`

SetLimits gets a reference to the given []LimitConfig and assigns it to the Limits field.

### GetCampaignGroups

`func (o *UpdateCampaign) GetCampaignGroups() []int32`

GetCampaignGroups returns the CampaignGroups field if non-nil, zero value otherwise.

### GetCampaignGroupsOk

`func (o *UpdateCampaign) GetCampaignGroupsOk() ([]int32, bool)`

GetCampaignGroupsOk returns a tuple with the CampaignGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignGroups

`func (o *UpdateCampaign) HasCampaignGroups() bool`

HasCampaignGroups returns a boolean if a field has been set.

### SetCampaignGroups

`func (o *UpdateCampaign) SetCampaignGroups(v []int32)`

SetCampaignGroups gets a reference to the given []int32 and assigns it to the CampaignGroups field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


