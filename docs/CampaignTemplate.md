# CampaignTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**AccountId** | Pointer to **int32** | The ID of the account that owns this entity. | 
**UserId** | Pointer to **int32** | The ID of the user associated with this entity. | 
**Name** | Pointer to **string** | The campaign template name. | 
**Description** | Pointer to **string** | Customer-facing text that explains the objective of the template. | 
**Instructions** | Pointer to **string** | Customer-facing text that explains how to use the template. For example, you can use this property to explain the available attributes of this template, and how they can be modified when a user uses this template to create a new campaign. | 
**CampaignAttributes** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | The campaign attributes that campaigns created from this template will have by default. | [optional] 
**CouponAttributes** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | The campaign attributes that coupons created from this template will have by default. | [optional] 
**State** | Pointer to **string** | Only campaign templates in &#39;available&#39; state may be used to create campaigns. | 
**ActiveRulesetId** | Pointer to **int32** | The ID of the ruleset this campaign template will use. | [optional] 
**Tags** | Pointer to **[]string** | A list of tags for the campaign template. | [optional] 
**Features** | Pointer to **[]string** | A list of features for the campaign template. | [optional] 
**CouponSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**ReferralSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**Limits** | Pointer to [**[]TemplateLimitConfig**](TemplateLimitConfig.md) | The set of limits that operate for this campaign template. | [optional] 
**TemplateParams** | Pointer to [**[]CampaignTemplateParams**](CampaignTemplateParams.md) | Fields which can be used to replace values in a rule. | [optional] 
**ApplicationsIds** | Pointer to **[]int32** |  | 
**CampaignCollections** | Pointer to [**[]CampaignTemplateCollection**](CampaignTemplateCollection.md) | The campaign collections from the blueprint campaign for the template. | [optional] 
**DefaultCampaignGroupId** | Pointer to **int32** | The default campaign group ID. | [optional] 
**CampaignType** | Pointer to **string** | The campaign type. Possible type values:   - &#x60;cartItem&#x60;: Type of campaign that can apply effects only to cart items.   - &#x60;advanced&#x60;: Type of campaign that can apply effects to customer sessions and cart items.  | [default to CAMPAIGN_TYPE_ADVANCED]
**Updated** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent update to the campaign template or any of its elements. | [optional] 
**UpdatedBy** | Pointer to **string** | Name of the user who last updated this campaign template, if available. | [optional] 
**ValidApplicationIds** | Pointer to **[]int32** | The IDs of the Applications that are related to this entity. | 

## Methods

### GetId

`func (o *CampaignTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignTemplate) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *CampaignTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *CampaignTemplate) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *CampaignTemplate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CampaignTemplate) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *CampaignTemplate) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *CampaignTemplate) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetAccountId

`func (o *CampaignTemplate) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CampaignTemplate) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *CampaignTemplate) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *CampaignTemplate) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetUserId

`func (o *CampaignTemplate) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CampaignTemplate) GetUserIdOk() (int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *CampaignTemplate) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *CampaignTemplate) SetUserId(v int32)`

SetUserId gets a reference to the given int32 and assigns it to the UserId field.

### GetName

`func (o *CampaignTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignTemplate) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *CampaignTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *CampaignTemplate) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *CampaignTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignTemplate) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *CampaignTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *CampaignTemplate) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetInstructions

`func (o *CampaignTemplate) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *CampaignTemplate) GetInstructionsOk() (string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInstructions

`func (o *CampaignTemplate) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructions

`func (o *CampaignTemplate) SetInstructions(v string)`

SetInstructions gets a reference to the given string and assigns it to the Instructions field.

### GetCampaignAttributes

`func (o *CampaignTemplate) GetCampaignAttributes() map[string]map[string]interface{}`

GetCampaignAttributes returns the CampaignAttributes field if non-nil, zero value otherwise.

### GetCampaignAttributesOk

`func (o *CampaignTemplate) GetCampaignAttributesOk() (map[string]map[string]interface{}, bool)`

GetCampaignAttributesOk returns a tuple with the CampaignAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignAttributes

`func (o *CampaignTemplate) HasCampaignAttributes() bool`

HasCampaignAttributes returns a boolean if a field has been set.

### SetCampaignAttributes

`func (o *CampaignTemplate) SetCampaignAttributes(v map[string]map[string]interface{})`

SetCampaignAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the CampaignAttributes field.

### GetCouponAttributes

`func (o *CampaignTemplate) GetCouponAttributes() map[string]map[string]interface{}`

GetCouponAttributes returns the CouponAttributes field if non-nil, zero value otherwise.

### GetCouponAttributesOk

`func (o *CampaignTemplate) GetCouponAttributesOk() (map[string]map[string]interface{}, bool)`

GetCouponAttributesOk returns a tuple with the CouponAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponAttributes

`func (o *CampaignTemplate) HasCouponAttributes() bool`

HasCouponAttributes returns a boolean if a field has been set.

### SetCouponAttributes

`func (o *CampaignTemplate) SetCouponAttributes(v map[string]map[string]interface{})`

SetCouponAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the CouponAttributes field.

### GetState

`func (o *CampaignTemplate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CampaignTemplate) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *CampaignTemplate) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *CampaignTemplate) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetActiveRulesetId

`func (o *CampaignTemplate) GetActiveRulesetId() int32`

GetActiveRulesetId returns the ActiveRulesetId field if non-nil, zero value otherwise.

### GetActiveRulesetIdOk

`func (o *CampaignTemplate) GetActiveRulesetIdOk() (int32, bool)`

GetActiveRulesetIdOk returns a tuple with the ActiveRulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveRulesetId

`func (o *CampaignTemplate) HasActiveRulesetId() bool`

HasActiveRulesetId returns a boolean if a field has been set.

### SetActiveRulesetId

`func (o *CampaignTemplate) SetActiveRulesetId(v int32)`

SetActiveRulesetId gets a reference to the given int32 and assigns it to the ActiveRulesetId field.

### GetTags

`func (o *CampaignTemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CampaignTemplate) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *CampaignTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *CampaignTemplate) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.

### GetFeatures

`func (o *CampaignTemplate) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CampaignTemplate) GetFeaturesOk() ([]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatures

`func (o *CampaignTemplate) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeatures

`func (o *CampaignTemplate) SetFeatures(v []string)`

SetFeatures gets a reference to the given []string and assigns it to the Features field.

### GetCouponSettings

`func (o *CampaignTemplate) GetCouponSettings() CodeGeneratorSettings`

GetCouponSettings returns the CouponSettings field if non-nil, zero value otherwise.

### GetCouponSettingsOk

`func (o *CampaignTemplate) GetCouponSettingsOk() (CodeGeneratorSettings, bool)`

GetCouponSettingsOk returns a tuple with the CouponSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponSettings

`func (o *CampaignTemplate) HasCouponSettings() bool`

HasCouponSettings returns a boolean if a field has been set.

### SetCouponSettings

`func (o *CampaignTemplate) SetCouponSettings(v CodeGeneratorSettings)`

SetCouponSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CouponSettings field.

### GetReferralSettings

`func (o *CampaignTemplate) GetReferralSettings() CodeGeneratorSettings`

GetReferralSettings returns the ReferralSettings field if non-nil, zero value otherwise.

### GetReferralSettingsOk

`func (o *CampaignTemplate) GetReferralSettingsOk() (CodeGeneratorSettings, bool)`

GetReferralSettingsOk returns a tuple with the ReferralSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralSettings

`func (o *CampaignTemplate) HasReferralSettings() bool`

HasReferralSettings returns a boolean if a field has been set.

### SetReferralSettings

`func (o *CampaignTemplate) SetReferralSettings(v CodeGeneratorSettings)`

SetReferralSettings gets a reference to the given CodeGeneratorSettings and assigns it to the ReferralSettings field.

### GetLimits

`func (o *CampaignTemplate) GetLimits() []TemplateLimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *CampaignTemplate) GetLimitsOk() ([]TemplateLimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimits

`func (o *CampaignTemplate) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimits

`func (o *CampaignTemplate) SetLimits(v []TemplateLimitConfig)`

SetLimits gets a reference to the given []TemplateLimitConfig and assigns it to the Limits field.

### GetTemplateParams

`func (o *CampaignTemplate) GetTemplateParams() []CampaignTemplateParams`

GetTemplateParams returns the TemplateParams field if non-nil, zero value otherwise.

### GetTemplateParamsOk

`func (o *CampaignTemplate) GetTemplateParamsOk() ([]CampaignTemplateParams, bool)`

GetTemplateParamsOk returns a tuple with the TemplateParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTemplateParams

`func (o *CampaignTemplate) HasTemplateParams() bool`

HasTemplateParams returns a boolean if a field has been set.

### SetTemplateParams

`func (o *CampaignTemplate) SetTemplateParams(v []CampaignTemplateParams)`

SetTemplateParams gets a reference to the given []CampaignTemplateParams and assigns it to the TemplateParams field.

### GetApplicationsIds

`func (o *CampaignTemplate) GetApplicationsIds() []int32`

GetApplicationsIds returns the ApplicationsIds field if non-nil, zero value otherwise.

### GetApplicationsIdsOk

`func (o *CampaignTemplate) GetApplicationsIdsOk() ([]int32, bool)`

GetApplicationsIdsOk returns a tuple with the ApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationsIds

`func (o *CampaignTemplate) HasApplicationsIds() bool`

HasApplicationsIds returns a boolean if a field has been set.

### SetApplicationsIds

`func (o *CampaignTemplate) SetApplicationsIds(v []int32)`

SetApplicationsIds gets a reference to the given []int32 and assigns it to the ApplicationsIds field.

### GetCampaignCollections

`func (o *CampaignTemplate) GetCampaignCollections() []CampaignTemplateCollection`

GetCampaignCollections returns the CampaignCollections field if non-nil, zero value otherwise.

### GetCampaignCollectionsOk

`func (o *CampaignTemplate) GetCampaignCollectionsOk() ([]CampaignTemplateCollection, bool)`

GetCampaignCollectionsOk returns a tuple with the CampaignCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignCollections

`func (o *CampaignTemplate) HasCampaignCollections() bool`

HasCampaignCollections returns a boolean if a field has been set.

### SetCampaignCollections

`func (o *CampaignTemplate) SetCampaignCollections(v []CampaignTemplateCollection)`

SetCampaignCollections gets a reference to the given []CampaignTemplateCollection and assigns it to the CampaignCollections field.

### GetDefaultCampaignGroupId

`func (o *CampaignTemplate) GetDefaultCampaignGroupId() int32`

GetDefaultCampaignGroupId returns the DefaultCampaignGroupId field if non-nil, zero value otherwise.

### GetDefaultCampaignGroupIdOk

`func (o *CampaignTemplate) GetDefaultCampaignGroupIdOk() (int32, bool)`

GetDefaultCampaignGroupIdOk returns a tuple with the DefaultCampaignGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefaultCampaignGroupId

`func (o *CampaignTemplate) HasDefaultCampaignGroupId() bool`

HasDefaultCampaignGroupId returns a boolean if a field has been set.

### SetDefaultCampaignGroupId

`func (o *CampaignTemplate) SetDefaultCampaignGroupId(v int32)`

SetDefaultCampaignGroupId gets a reference to the given int32 and assigns it to the DefaultCampaignGroupId field.

### GetCampaignType

`func (o *CampaignTemplate) GetCampaignType() string`

GetCampaignType returns the CampaignType field if non-nil, zero value otherwise.

### GetCampaignTypeOk

`func (o *CampaignTemplate) GetCampaignTypeOk() (string, bool)`

GetCampaignTypeOk returns a tuple with the CampaignType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignType

`func (o *CampaignTemplate) HasCampaignType() bool`

HasCampaignType returns a boolean if a field has been set.

### SetCampaignType

`func (o *CampaignTemplate) SetCampaignType(v string)`

SetCampaignType gets a reference to the given string and assigns it to the CampaignType field.

### GetUpdated

`func (o *CampaignTemplate) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CampaignTemplate) GetUpdatedOk() (time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *CampaignTemplate) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *CampaignTemplate) SetUpdated(v time.Time)`

SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.

### GetUpdatedBy

`func (o *CampaignTemplate) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *CampaignTemplate) GetUpdatedByOk() (string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdatedBy

`func (o *CampaignTemplate) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedBy

`func (o *CampaignTemplate) SetUpdatedBy(v string)`

SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.

### GetValidApplicationIds

`func (o *CampaignTemplate) GetValidApplicationIds() []int32`

GetValidApplicationIds returns the ValidApplicationIds field if non-nil, zero value otherwise.

### GetValidApplicationIdsOk

`func (o *CampaignTemplate) GetValidApplicationIdsOk() ([]int32, bool)`

GetValidApplicationIdsOk returns a tuple with the ValidApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValidApplicationIds

`func (o *CampaignTemplate) HasValidApplicationIds() bool`

HasValidApplicationIds returns a boolean if a field has been set.

### SetValidApplicationIds

`func (o *CampaignTemplate) SetValidApplicationIds(v []int32)`

SetValidApplicationIds gets a reference to the given []int32 and assigns it to the ValidApplicationIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


