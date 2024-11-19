# RevisionVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**AccountId** | Pointer to **int32** |  | 
**ApplicationId** | Pointer to **int32** |  | 
**CampaignId** | Pointer to **int32** |  | 
**Created** | Pointer to [**time.Time**](time.Time.md) |  | 
**CreatedBy** | Pointer to **int32** |  | 
**RevisionId** | Pointer to **int32** |  | 
**Version** | Pointer to **int32** |  | 
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

### GetId

`func (o *RevisionVersion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RevisionVersion) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *RevisionVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *RevisionVersion) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetAccountId

`func (o *RevisionVersion) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RevisionVersion) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *RevisionVersion) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *RevisionVersion) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetApplicationId

`func (o *RevisionVersion) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *RevisionVersion) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *RevisionVersion) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *RevisionVersion) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetCampaignId

`func (o *RevisionVersion) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *RevisionVersion) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *RevisionVersion) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *RevisionVersion) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.

### GetCreated

`func (o *RevisionVersion) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RevisionVersion) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *RevisionVersion) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *RevisionVersion) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetCreatedBy

`func (o *RevisionVersion) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RevisionVersion) GetCreatedByOk() (int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *RevisionVersion) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *RevisionVersion) SetCreatedBy(v int32)`

SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.

### GetRevisionId

`func (o *RevisionVersion) GetRevisionId() int32`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *RevisionVersion) GetRevisionIdOk() (int32, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRevisionId

`func (o *RevisionVersion) HasRevisionId() bool`

HasRevisionId returns a boolean if a field has been set.

### SetRevisionId

`func (o *RevisionVersion) SetRevisionId(v int32)`

SetRevisionId gets a reference to the given int32 and assigns it to the RevisionId field.

### GetVersion

`func (o *RevisionVersion) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RevisionVersion) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *RevisionVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *RevisionVersion) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetName

`func (o *RevisionVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RevisionVersion) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *RevisionVersion) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *RevisionVersion) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetStartTime

`func (o *RevisionVersion) GetStartTime() NullableTime`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RevisionVersion) GetStartTimeOk() (NullableTime, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartTime

`func (o *RevisionVersion) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTime

`func (o *RevisionVersion) SetStartTime(v NullableTime)`

SetStartTime gets a reference to the given NullableTime and assigns it to the StartTime field.

### SetStartTimeExplicitNull

`func (o *RevisionVersion) SetStartTimeExplicitNull(b bool)`

SetStartTimeExplicitNull (un)sets StartTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The StartTime value is set to nil even if false is passed
### GetEndTime

`func (o *RevisionVersion) GetEndTime() NullableTime`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *RevisionVersion) GetEndTimeOk() (NullableTime, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndTime

`func (o *RevisionVersion) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTime

`func (o *RevisionVersion) SetEndTime(v NullableTime)`

SetEndTime gets a reference to the given NullableTime and assigns it to the EndTime field.

### SetEndTimeExplicitNull

`func (o *RevisionVersion) SetEndTimeExplicitNull(b bool)`

SetEndTimeExplicitNull (un)sets EndTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EndTime value is set to nil even if false is passed
### GetAttributes

`func (o *RevisionVersion) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RevisionVersion) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *RevisionVersion) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *RevisionVersion) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetDescription

`func (o *RevisionVersion) GetDescription() NullableString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RevisionVersion) GetDescriptionOk() (NullableString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *RevisionVersion) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *RevisionVersion) SetDescription(v NullableString)`

SetDescription gets a reference to the given NullableString and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *RevisionVersion) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetActiveRulesetId

`func (o *RevisionVersion) GetActiveRulesetId() NullableInt32`

GetActiveRulesetId returns the ActiveRulesetId field if non-nil, zero value otherwise.

### GetActiveRulesetIdOk

`func (o *RevisionVersion) GetActiveRulesetIdOk() (NullableInt32, bool)`

GetActiveRulesetIdOk returns a tuple with the ActiveRulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveRulesetId

`func (o *RevisionVersion) HasActiveRulesetId() bool`

HasActiveRulesetId returns a boolean if a field has been set.

### SetActiveRulesetId

`func (o *RevisionVersion) SetActiveRulesetId(v NullableInt32)`

SetActiveRulesetId gets a reference to the given NullableInt32 and assigns it to the ActiveRulesetId field.

### SetActiveRulesetIdExplicitNull

`func (o *RevisionVersion) SetActiveRulesetIdExplicitNull(b bool)`

SetActiveRulesetIdExplicitNull (un)sets ActiveRulesetId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ActiveRulesetId value is set to nil even if false is passed
### GetTags

`func (o *RevisionVersion) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RevisionVersion) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *RevisionVersion) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *RevisionVersion) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.

### GetCouponSettings

`func (o *RevisionVersion) GetCouponSettings() CodeGeneratorSettings`

GetCouponSettings returns the CouponSettings field if non-nil, zero value otherwise.

### GetCouponSettingsOk

`func (o *RevisionVersion) GetCouponSettingsOk() (CodeGeneratorSettings, bool)`

GetCouponSettingsOk returns a tuple with the CouponSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponSettings

`func (o *RevisionVersion) HasCouponSettings() bool`

HasCouponSettings returns a boolean if a field has been set.

### SetCouponSettings

`func (o *RevisionVersion) SetCouponSettings(v CodeGeneratorSettings)`

SetCouponSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CouponSettings field.

### GetReferralSettings

`func (o *RevisionVersion) GetReferralSettings() CodeGeneratorSettings`

GetReferralSettings returns the ReferralSettings field if non-nil, zero value otherwise.

### GetReferralSettingsOk

`func (o *RevisionVersion) GetReferralSettingsOk() (CodeGeneratorSettings, bool)`

GetReferralSettingsOk returns a tuple with the ReferralSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralSettings

`func (o *RevisionVersion) HasReferralSettings() bool`

HasReferralSettings returns a boolean if a field has been set.

### SetReferralSettings

`func (o *RevisionVersion) SetReferralSettings(v CodeGeneratorSettings)`

SetReferralSettings gets a reference to the given CodeGeneratorSettings and assigns it to the ReferralSettings field.

### GetLimits

`func (o *RevisionVersion) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *RevisionVersion) GetLimitsOk() ([]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimits

`func (o *RevisionVersion) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimits

`func (o *RevisionVersion) SetLimits(v []LimitConfig)`

SetLimits gets a reference to the given []LimitConfig and assigns it to the Limits field.

### GetFeatures

`func (o *RevisionVersion) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *RevisionVersion) GetFeaturesOk() ([]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatures

`func (o *RevisionVersion) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeatures

`func (o *RevisionVersion) SetFeatures(v []string)`

SetFeatures gets a reference to the given []string and assigns it to the Features field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


