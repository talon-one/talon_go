# UpdateCampaignTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The campaign template name. | 
**Description** | Pointer to **string** | Customer-facing text that explains the objective of the template. | 
**Instructions** | Pointer to **string** | Customer-facing text that explains how to use the template. For example, you can use this property to explain the available attributes of this template, and how they can be modified when a user uses this template to create a new campaign. | 
**CampaignAttributes** | Pointer to [**map[string]interface{}**](.md) | The Campaign Attributes that Campaigns created from this template will have by default. | [optional] 
**CouponAttributes** | Pointer to [**map[string]interface{}**](.md) | The Campaign Attributes that Coupons created from this template will have by default. | [optional] 
**State** | Pointer to **string** | Only Campaign Templates in &#39;available&#39; state may be used to create Campaigns. | 
**ActiveRulesetId** | Pointer to **int32** | The ID of the Ruleset this Campaign Template will use. | [optional] 
**Tags** | Pointer to **[]string** | A list of tags for the campaign template. | [optional] 
**Features** | Pointer to **[]string** | A list of features for the campaign template. | [optional] 
**CouponSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**ReferralSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**Limits** | Pointer to [**[]TemplateLimitConfig**](TemplateLimitConfig.md) | The set of limits that will operate for this campaign template. | [optional] 
**TemplateParams** | Pointer to [**[]CampaignTemplateParams**](CampaignTemplateParams.md) | Template parameters are fields which can be used to replace values in a rule. | [optional] 
**ApplicationsIds** | Pointer to **[]int32** | A list of the IDs of the applications that are subscribed to this campaign template. | 
**CampaignCollections** | Pointer to [**[]CampaignTemplateCollection**](CampaignTemplateCollection.md) | The campaign collections from the blueprint campaign for the template. | [optional] 
**DefaultCampaignGroupId** | Pointer to **int32** | The default campaignGroupId. | [optional] 

## Methods

### GetName

`func (o *UpdateCampaignTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCampaignTemplate) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *UpdateCampaignTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *UpdateCampaignTemplate) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *UpdateCampaignTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCampaignTemplate) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *UpdateCampaignTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *UpdateCampaignTemplate) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetInstructions

`func (o *UpdateCampaignTemplate) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *UpdateCampaignTemplate) GetInstructionsOk() (string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInstructions

`func (o *UpdateCampaignTemplate) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructions

`func (o *UpdateCampaignTemplate) SetInstructions(v string)`

SetInstructions gets a reference to the given string and assigns it to the Instructions field.

### GetCampaignAttributes

`func (o *UpdateCampaignTemplate) GetCampaignAttributes() map[string]interface{}`

GetCampaignAttributes returns the CampaignAttributes field if non-nil, zero value otherwise.

### GetCampaignAttributesOk

`func (o *UpdateCampaignTemplate) GetCampaignAttributesOk() (map[string]interface{}, bool)`

GetCampaignAttributesOk returns a tuple with the CampaignAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignAttributes

`func (o *UpdateCampaignTemplate) HasCampaignAttributes() bool`

HasCampaignAttributes returns a boolean if a field has been set.

### SetCampaignAttributes

`func (o *UpdateCampaignTemplate) SetCampaignAttributes(v map[string]interface{})`

SetCampaignAttributes gets a reference to the given map[string]interface{} and assigns it to the CampaignAttributes field.

### GetCouponAttributes

`func (o *UpdateCampaignTemplate) GetCouponAttributes() map[string]interface{}`

GetCouponAttributes returns the CouponAttributes field if non-nil, zero value otherwise.

### GetCouponAttributesOk

`func (o *UpdateCampaignTemplate) GetCouponAttributesOk() (map[string]interface{}, bool)`

GetCouponAttributesOk returns a tuple with the CouponAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponAttributes

`func (o *UpdateCampaignTemplate) HasCouponAttributes() bool`

HasCouponAttributes returns a boolean if a field has been set.

### SetCouponAttributes

`func (o *UpdateCampaignTemplate) SetCouponAttributes(v map[string]interface{})`

SetCouponAttributes gets a reference to the given map[string]interface{} and assigns it to the CouponAttributes field.

### GetState

`func (o *UpdateCampaignTemplate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateCampaignTemplate) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *UpdateCampaignTemplate) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *UpdateCampaignTemplate) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetActiveRulesetId

`func (o *UpdateCampaignTemplate) GetActiveRulesetId() int32`

GetActiveRulesetId returns the ActiveRulesetId field if non-nil, zero value otherwise.

### GetActiveRulesetIdOk

`func (o *UpdateCampaignTemplate) GetActiveRulesetIdOk() (int32, bool)`

GetActiveRulesetIdOk returns a tuple with the ActiveRulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveRulesetId

`func (o *UpdateCampaignTemplate) HasActiveRulesetId() bool`

HasActiveRulesetId returns a boolean if a field has been set.

### SetActiveRulesetId

`func (o *UpdateCampaignTemplate) SetActiveRulesetId(v int32)`

SetActiveRulesetId gets a reference to the given int32 and assigns it to the ActiveRulesetId field.

### GetTags

`func (o *UpdateCampaignTemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateCampaignTemplate) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *UpdateCampaignTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *UpdateCampaignTemplate) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.

### GetFeatures

`func (o *UpdateCampaignTemplate) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *UpdateCampaignTemplate) GetFeaturesOk() ([]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatures

`func (o *UpdateCampaignTemplate) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeatures

`func (o *UpdateCampaignTemplate) SetFeatures(v []string)`

SetFeatures gets a reference to the given []string and assigns it to the Features field.

### GetCouponSettings

`func (o *UpdateCampaignTemplate) GetCouponSettings() CodeGeneratorSettings`

GetCouponSettings returns the CouponSettings field if non-nil, zero value otherwise.

### GetCouponSettingsOk

`func (o *UpdateCampaignTemplate) GetCouponSettingsOk() (CodeGeneratorSettings, bool)`

GetCouponSettingsOk returns a tuple with the CouponSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponSettings

`func (o *UpdateCampaignTemplate) HasCouponSettings() bool`

HasCouponSettings returns a boolean if a field has been set.

### SetCouponSettings

`func (o *UpdateCampaignTemplate) SetCouponSettings(v CodeGeneratorSettings)`

SetCouponSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CouponSettings field.

### GetReferralSettings

`func (o *UpdateCampaignTemplate) GetReferralSettings() CodeGeneratorSettings`

GetReferralSettings returns the ReferralSettings field if non-nil, zero value otherwise.

### GetReferralSettingsOk

`func (o *UpdateCampaignTemplate) GetReferralSettingsOk() (CodeGeneratorSettings, bool)`

GetReferralSettingsOk returns a tuple with the ReferralSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralSettings

`func (o *UpdateCampaignTemplate) HasReferralSettings() bool`

HasReferralSettings returns a boolean if a field has been set.

### SetReferralSettings

`func (o *UpdateCampaignTemplate) SetReferralSettings(v CodeGeneratorSettings)`

SetReferralSettings gets a reference to the given CodeGeneratorSettings and assigns it to the ReferralSettings field.

### GetLimits

`func (o *UpdateCampaignTemplate) GetLimits() []TemplateLimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *UpdateCampaignTemplate) GetLimitsOk() ([]TemplateLimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimits

`func (o *UpdateCampaignTemplate) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimits

`func (o *UpdateCampaignTemplate) SetLimits(v []TemplateLimitConfig)`

SetLimits gets a reference to the given []TemplateLimitConfig and assigns it to the Limits field.

### GetTemplateParams

`func (o *UpdateCampaignTemplate) GetTemplateParams() []CampaignTemplateParams`

GetTemplateParams returns the TemplateParams field if non-nil, zero value otherwise.

### GetTemplateParamsOk

`func (o *UpdateCampaignTemplate) GetTemplateParamsOk() ([]CampaignTemplateParams, bool)`

GetTemplateParamsOk returns a tuple with the TemplateParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTemplateParams

`func (o *UpdateCampaignTemplate) HasTemplateParams() bool`

HasTemplateParams returns a boolean if a field has been set.

### SetTemplateParams

`func (o *UpdateCampaignTemplate) SetTemplateParams(v []CampaignTemplateParams)`

SetTemplateParams gets a reference to the given []CampaignTemplateParams and assigns it to the TemplateParams field.

### GetApplicationsIds

`func (o *UpdateCampaignTemplate) GetApplicationsIds() []int32`

GetApplicationsIds returns the ApplicationsIds field if non-nil, zero value otherwise.

### GetApplicationsIdsOk

`func (o *UpdateCampaignTemplate) GetApplicationsIdsOk() ([]int32, bool)`

GetApplicationsIdsOk returns a tuple with the ApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationsIds

`func (o *UpdateCampaignTemplate) HasApplicationsIds() bool`

HasApplicationsIds returns a boolean if a field has been set.

### SetApplicationsIds

`func (o *UpdateCampaignTemplate) SetApplicationsIds(v []int32)`

SetApplicationsIds gets a reference to the given []int32 and assigns it to the ApplicationsIds field.

### GetCampaignCollections

`func (o *UpdateCampaignTemplate) GetCampaignCollections() []CampaignTemplateCollection`

GetCampaignCollections returns the CampaignCollections field if non-nil, zero value otherwise.

### GetCampaignCollectionsOk

`func (o *UpdateCampaignTemplate) GetCampaignCollectionsOk() ([]CampaignTemplateCollection, bool)`

GetCampaignCollectionsOk returns a tuple with the CampaignCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignCollections

`func (o *UpdateCampaignTemplate) HasCampaignCollections() bool`

HasCampaignCollections returns a boolean if a field has been set.

### SetCampaignCollections

`func (o *UpdateCampaignTemplate) SetCampaignCollections(v []CampaignTemplateCollection)`

SetCampaignCollections gets a reference to the given []CampaignTemplateCollection and assigns it to the CampaignCollections field.

### GetDefaultCampaignGroupId

`func (o *UpdateCampaignTemplate) GetDefaultCampaignGroupId() int32`

GetDefaultCampaignGroupId returns the DefaultCampaignGroupId field if non-nil, zero value otherwise.

### GetDefaultCampaignGroupIdOk

`func (o *UpdateCampaignTemplate) GetDefaultCampaignGroupIdOk() (int32, bool)`

GetDefaultCampaignGroupIdOk returns a tuple with the DefaultCampaignGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefaultCampaignGroupId

`func (o *UpdateCampaignTemplate) HasDefaultCampaignGroupId() bool`

HasDefaultCampaignGroupId returns a boolean if a field has been set.

### SetDefaultCampaignGroupId

`func (o *UpdateCampaignTemplate) SetDefaultCampaignGroupId(v int32)`

SetDefaultCampaignGroupId gets a reference to the given int32 and assigns it to the DefaultCampaignGroupId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


