# NewRevisionVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A user-facing name for this campaign. | [optional] 
**StartTime** | Pointer to [**NullableTime**](time.Time.md) | Timestamp when the campaign will become active. | [optional] 
**EndTime** | Pointer to [**NullableTime**](time.Time.md) | Timestamp when the campaign will become inactive. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this campaign. | [optional] 
**Description** | Pointer to **NullableString** | A detailed description of the campaign. | [optional] 
**ActiveRulesetId** | Pointer to **NullableInt32** | The ID of the ruleset this campaign template will use. | [optional] 
**Tags** | Pointer to **[]string** | A list of tags for the campaign template. | [optional] 
**CouponSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**ReferralSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**Limits** | Pointer to [**[]LimitConfig**](LimitConfig.md) | The set of limits that will operate for this campaign version. | [optional] 
**Features** | Pointer to **[]string** | A list of features for the campaign template. | [optional] 

## Methods

### GetName

`func (o *NewRevisionVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewRevisionVersion) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *NewRevisionVersion) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *NewRevisionVersion) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetStartTime

`func (o *NewRevisionVersion) GetStartTime() NullableTime`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *NewRevisionVersion) GetStartTimeOk() (NullableTime, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartTime

`func (o *NewRevisionVersion) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTime

`func (o *NewRevisionVersion) SetStartTime(v NullableTime)`

SetStartTime gets a reference to the given NullableTime and assigns it to the StartTime field.

### SetStartTimeExplicitNull

`func (o *NewRevisionVersion) SetStartTimeExplicitNull(b bool)`

SetStartTimeExplicitNull (un)sets StartTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The StartTime value is set to nil even if false is passed
### GetEndTime

`func (o *NewRevisionVersion) GetEndTime() NullableTime`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *NewRevisionVersion) GetEndTimeOk() (NullableTime, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndTime

`func (o *NewRevisionVersion) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTime

`func (o *NewRevisionVersion) SetEndTime(v NullableTime)`

SetEndTime gets a reference to the given NullableTime and assigns it to the EndTime field.

### SetEndTimeExplicitNull

`func (o *NewRevisionVersion) SetEndTimeExplicitNull(b bool)`

SetEndTimeExplicitNull (un)sets EndTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EndTime value is set to nil even if false is passed
### GetAttributes

`func (o *NewRevisionVersion) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NewRevisionVersion) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *NewRevisionVersion) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *NewRevisionVersion) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetDescription

`func (o *NewRevisionVersion) GetDescription() NullableString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewRevisionVersion) GetDescriptionOk() (NullableString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *NewRevisionVersion) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *NewRevisionVersion) SetDescription(v NullableString)`

SetDescription gets a reference to the given NullableString and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *NewRevisionVersion) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetActiveRulesetId

`func (o *NewRevisionVersion) GetActiveRulesetId() NullableInt32`

GetActiveRulesetId returns the ActiveRulesetId field if non-nil, zero value otherwise.

### GetActiveRulesetIdOk

`func (o *NewRevisionVersion) GetActiveRulesetIdOk() (NullableInt32, bool)`

GetActiveRulesetIdOk returns a tuple with the ActiveRulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveRulesetId

`func (o *NewRevisionVersion) HasActiveRulesetId() bool`

HasActiveRulesetId returns a boolean if a field has been set.

### SetActiveRulesetId

`func (o *NewRevisionVersion) SetActiveRulesetId(v NullableInt32)`

SetActiveRulesetId gets a reference to the given NullableInt32 and assigns it to the ActiveRulesetId field.

### SetActiveRulesetIdExplicitNull

`func (o *NewRevisionVersion) SetActiveRulesetIdExplicitNull(b bool)`

SetActiveRulesetIdExplicitNull (un)sets ActiveRulesetId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ActiveRulesetId value is set to nil even if false is passed
### GetTags

`func (o *NewRevisionVersion) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NewRevisionVersion) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *NewRevisionVersion) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *NewRevisionVersion) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.

### GetCouponSettings

`func (o *NewRevisionVersion) GetCouponSettings() CodeGeneratorSettings`

GetCouponSettings returns the CouponSettings field if non-nil, zero value otherwise.

### GetCouponSettingsOk

`func (o *NewRevisionVersion) GetCouponSettingsOk() (CodeGeneratorSettings, bool)`

GetCouponSettingsOk returns a tuple with the CouponSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponSettings

`func (o *NewRevisionVersion) HasCouponSettings() bool`

HasCouponSettings returns a boolean if a field has been set.

### SetCouponSettings

`func (o *NewRevisionVersion) SetCouponSettings(v CodeGeneratorSettings)`

SetCouponSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CouponSettings field.

### GetReferralSettings

`func (o *NewRevisionVersion) GetReferralSettings() CodeGeneratorSettings`

GetReferralSettings returns the ReferralSettings field if non-nil, zero value otherwise.

### GetReferralSettingsOk

`func (o *NewRevisionVersion) GetReferralSettingsOk() (CodeGeneratorSettings, bool)`

GetReferralSettingsOk returns a tuple with the ReferralSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralSettings

`func (o *NewRevisionVersion) HasReferralSettings() bool`

HasReferralSettings returns a boolean if a field has been set.

### SetReferralSettings

`func (o *NewRevisionVersion) SetReferralSettings(v CodeGeneratorSettings)`

SetReferralSettings gets a reference to the given CodeGeneratorSettings and assigns it to the ReferralSettings field.

### GetLimits

`func (o *NewRevisionVersion) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *NewRevisionVersion) GetLimitsOk() ([]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimits

`func (o *NewRevisionVersion) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimits

`func (o *NewRevisionVersion) SetLimits(v []LimitConfig)`

SetLimits gets a reference to the given []LimitConfig and assigns it to the Limits field.

### GetFeatures

`func (o *NewRevisionVersion) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *NewRevisionVersion) GetFeaturesOk() ([]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatures

`func (o *NewRevisionVersion) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeatures

`func (o *NewRevisionVersion) SetFeatures(v []string)`

SetFeatures gets a reference to the given []string and assigns it to the Features field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


