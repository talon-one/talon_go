# NewCampaignTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** | A list of tags for the campaign template. | [optional] 
**CampaignAttributes** | Pointer to [**map[string]interface{}**](.md) | The campaign attributes that campaigns created from this template will have by default. | [optional] 
**CampaignCollections** | Pointer to [**[]CampaignTemplateCollection**](CampaignTemplateCollection.md) | The campaign collections from the blueprint campaign for the template. | [optional] 
**CampaignType** | Pointer to **string** | The campaign type. Possible type values:   - &#x60;cartItem&#x60;: Type of campaign that can apply effects only to cart items.   - &#x60;advanced&#x60;: Type of campaign that can apply effects to customer sessions and cart items.  | [default to CAMPAIGN_TYPE_ADVANCED]
**CouponAttributes** | Pointer to [**map[string]interface{}**](.md) | The campaign attributes that coupons created from this template will have by default. | [optional] 
**CouponSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**DefaultCampaignGroupId** | Pointer to **int32** | The default campaign group ID. | [optional] 
**Description** | Pointer to **string** | Customer-facing text that explains the objective of the template. | 
**Features** | Pointer to **[]string** | A list of features for the campaign template. | [optional] 
**Instructions** | Pointer to **string** | Customer-facing text that explains how to use the template. For example, you can use this property to explain the available attributes of this template, and how they can be modified when a user uses this template to create a new campaign. | 
**Limits** | Pointer to [**[]TemplateLimitConfig**](TemplateLimitConfig.md) | The set of limits that will operate for this campaign template. | [optional] 
**Name** | Pointer to **string** | The campaign template name. | 
**ReferralSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**State** | Pointer to **string** | Only Campaign Templates in &#39;available&#39; state may be used to create Campaigns. | 
**TemplateParams** | Pointer to [**[]CampaignTemplateParams**](CampaignTemplateParams.md) | Fields which can be used to replace values in a rule. | [optional] 

## Methods

### GetTags

`func (o *NewCampaignTemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NewCampaignTemplate) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *NewCampaignTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *NewCampaignTemplate) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.

### GetCampaignAttributes

`func (o *NewCampaignTemplate) GetCampaignAttributes() map[string]interface{}`

GetCampaignAttributes returns the CampaignAttributes field if non-nil, zero value otherwise.

### GetCampaignAttributesOk

`func (o *NewCampaignTemplate) GetCampaignAttributesOk() (map[string]interface{}, bool)`

GetCampaignAttributesOk returns a tuple with the CampaignAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignAttributes

`func (o *NewCampaignTemplate) HasCampaignAttributes() bool`

HasCampaignAttributes returns a boolean if a field has been set.

### SetCampaignAttributes

`func (o *NewCampaignTemplate) SetCampaignAttributes(v map[string]interface{})`

SetCampaignAttributes gets a reference to the given map[string]interface{} and assigns it to the CampaignAttributes field.

### GetCampaignCollections

`func (o *NewCampaignTemplate) GetCampaignCollections() []CampaignTemplateCollection`

GetCampaignCollections returns the CampaignCollections field if non-nil, zero value otherwise.

### GetCampaignCollectionsOk

`func (o *NewCampaignTemplate) GetCampaignCollectionsOk() ([]CampaignTemplateCollection, bool)`

GetCampaignCollectionsOk returns a tuple with the CampaignCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignCollections

`func (o *NewCampaignTemplate) HasCampaignCollections() bool`

HasCampaignCollections returns a boolean if a field has been set.

### SetCampaignCollections

`func (o *NewCampaignTemplate) SetCampaignCollections(v []CampaignTemplateCollection)`

SetCampaignCollections gets a reference to the given []CampaignTemplateCollection and assigns it to the CampaignCollections field.

### GetCampaignType

`func (o *NewCampaignTemplate) GetCampaignType() string`

GetCampaignType returns the CampaignType field if non-nil, zero value otherwise.

### GetCampaignTypeOk

`func (o *NewCampaignTemplate) GetCampaignTypeOk() (string, bool)`

GetCampaignTypeOk returns a tuple with the CampaignType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignType

`func (o *NewCampaignTemplate) HasCampaignType() bool`

HasCampaignType returns a boolean if a field has been set.

### SetCampaignType

`func (o *NewCampaignTemplate) SetCampaignType(v string)`

SetCampaignType gets a reference to the given string and assigns it to the CampaignType field.

### GetCouponAttributes

`func (o *NewCampaignTemplate) GetCouponAttributes() map[string]interface{}`

GetCouponAttributes returns the CouponAttributes field if non-nil, zero value otherwise.

### GetCouponAttributesOk

`func (o *NewCampaignTemplate) GetCouponAttributesOk() (map[string]interface{}, bool)`

GetCouponAttributesOk returns a tuple with the CouponAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponAttributes

`func (o *NewCampaignTemplate) HasCouponAttributes() bool`

HasCouponAttributes returns a boolean if a field has been set.

### SetCouponAttributes

`func (o *NewCampaignTemplate) SetCouponAttributes(v map[string]interface{})`

SetCouponAttributes gets a reference to the given map[string]interface{} and assigns it to the CouponAttributes field.

### GetCouponSettings

`func (o *NewCampaignTemplate) GetCouponSettings() CodeGeneratorSettings`

GetCouponSettings returns the CouponSettings field if non-nil, zero value otherwise.

### GetCouponSettingsOk

`func (o *NewCampaignTemplate) GetCouponSettingsOk() (CodeGeneratorSettings, bool)`

GetCouponSettingsOk returns a tuple with the CouponSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponSettings

`func (o *NewCampaignTemplate) HasCouponSettings() bool`

HasCouponSettings returns a boolean if a field has been set.

### SetCouponSettings

`func (o *NewCampaignTemplate) SetCouponSettings(v CodeGeneratorSettings)`

SetCouponSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CouponSettings field.

### GetDefaultCampaignGroupId

`func (o *NewCampaignTemplate) GetDefaultCampaignGroupId() int32`

GetDefaultCampaignGroupId returns the DefaultCampaignGroupId field if non-nil, zero value otherwise.

### GetDefaultCampaignGroupIdOk

`func (o *NewCampaignTemplate) GetDefaultCampaignGroupIdOk() (int32, bool)`

GetDefaultCampaignGroupIdOk returns a tuple with the DefaultCampaignGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefaultCampaignGroupId

`func (o *NewCampaignTemplate) HasDefaultCampaignGroupId() bool`

HasDefaultCampaignGroupId returns a boolean if a field has been set.

### SetDefaultCampaignGroupId

`func (o *NewCampaignTemplate) SetDefaultCampaignGroupId(v int32)`

SetDefaultCampaignGroupId gets a reference to the given int32 and assigns it to the DefaultCampaignGroupId field.

### GetDescription

`func (o *NewCampaignTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewCampaignTemplate) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *NewCampaignTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *NewCampaignTemplate) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetFeatures

`func (o *NewCampaignTemplate) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *NewCampaignTemplate) GetFeaturesOk() ([]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatures

`func (o *NewCampaignTemplate) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeatures

`func (o *NewCampaignTemplate) SetFeatures(v []string)`

SetFeatures gets a reference to the given []string and assigns it to the Features field.

### GetInstructions

`func (o *NewCampaignTemplate) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *NewCampaignTemplate) GetInstructionsOk() (string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInstructions

`func (o *NewCampaignTemplate) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructions

`func (o *NewCampaignTemplate) SetInstructions(v string)`

SetInstructions gets a reference to the given string and assigns it to the Instructions field.

### GetLimits

`func (o *NewCampaignTemplate) GetLimits() []TemplateLimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *NewCampaignTemplate) GetLimitsOk() ([]TemplateLimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimits

`func (o *NewCampaignTemplate) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimits

`func (o *NewCampaignTemplate) SetLimits(v []TemplateLimitConfig)`

SetLimits gets a reference to the given []TemplateLimitConfig and assigns it to the Limits field.

### GetName

`func (o *NewCampaignTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewCampaignTemplate) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *NewCampaignTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *NewCampaignTemplate) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetReferralSettings

`func (o *NewCampaignTemplate) GetReferralSettings() CodeGeneratorSettings`

GetReferralSettings returns the ReferralSettings field if non-nil, zero value otherwise.

### GetReferralSettingsOk

`func (o *NewCampaignTemplate) GetReferralSettingsOk() (CodeGeneratorSettings, bool)`

GetReferralSettingsOk returns a tuple with the ReferralSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralSettings

`func (o *NewCampaignTemplate) HasReferralSettings() bool`

HasReferralSettings returns a boolean if a field has been set.

### SetReferralSettings

`func (o *NewCampaignTemplate) SetReferralSettings(v CodeGeneratorSettings)`

SetReferralSettings gets a reference to the given CodeGeneratorSettings and assigns it to the ReferralSettings field.

### GetState

`func (o *NewCampaignTemplate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NewCampaignTemplate) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *NewCampaignTemplate) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *NewCampaignTemplate) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetTemplateParams

`func (o *NewCampaignTemplate) GetTemplateParams() []CampaignTemplateParams`

GetTemplateParams returns the TemplateParams field if non-nil, zero value otherwise.

### GetTemplateParamsOk

`func (o *NewCampaignTemplate) GetTemplateParamsOk() ([]CampaignTemplateParams, bool)`

GetTemplateParamsOk returns a tuple with the TemplateParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTemplateParams

`func (o *NewCampaignTemplate) HasTemplateParams() bool`

HasTemplateParams returns a boolean if a field has been set.

### SetTemplateParams

`func (o *NewCampaignTemplate) SetTemplateParams(v []CampaignTemplateParams)`

SetTemplateParams gets a reference to the given []CampaignTemplateParams and assigns it to the TemplateParams field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


