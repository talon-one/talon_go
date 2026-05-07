# NewCampaignTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The campaign template name. | 
**Description** | Pointer to **string** | Customer-facing text that explains the objective of the template. | 
**Instructions** | Pointer to **string** | Customer-facing text that explains how to use the template. For example, you can use this property to explain the available attributes of this template, and how they can be modified when a user uses this template to create a new campaign. | 
**CampaignAttributes** | Pointer to [**map[string]interface{}**](.md) | The campaign attributes that campaigns created from this template will have by default. | [optional] 
**CouponAttributes** | Pointer to [**map[string]interface{}**](.md) | The campaign attributes that coupons created from this template will have by default. | [optional] 
**State** | Pointer to **string** | Only Campaign Templates in &#39;available&#39; state may be used to create Campaigns. | 
**Tags** | Pointer to **[]string** | A list of tags for the campaign template. | [optional] 
**ReevaluateOnReturn** | Pointer to **bool** | Indicates whether campaigns created from this template should be reevaluated when a customer returns an item. | [optional] 
**Features** | Pointer to **[]string** | A list of features for the campaign template. | [optional] 
**CouponSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**CouponReservationSettings** | Pointer to [**CampaignTemplateCouponReservationSettings**](CampaignTemplateCouponReservationSettings.md) |  | [optional] 
**ReferralSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**Limits** | Pointer to [**[]TemplateLimitConfig**](TemplateLimitConfig.md) | The set of limits that will operate for this campaign template. | [optional] 
**TemplateParams** | Pointer to [**[]CampaignTemplateParams**](CampaignTemplateParams.md) | Fields which can be used to replace values in a rule. | [optional] 
**CampaignCollections** | Pointer to [**[]CampaignTemplateCollection**](CampaignTemplateCollection.md) | The campaign collections from the blueprint campaign for the template. | [optional] 
**DefaultCampaignGroupId** | Pointer to **int64** | The default campaign group ID. | [optional] 
**CampaignType** | Pointer to **string** | The campaign type. Possible type values:   - &#x60;cartItem&#x60;: Type of campaign that can apply effects only to cart items.   - &#x60;advanced&#x60;: Type of campaign that can apply effects to customer sessions and cart items.  | [default to "advanced"]

## Methods

### NewNewCampaignTemplate

`func NewNewCampaignTemplate(name string, description string, instructions string, state string, campaignType string, ) *NewCampaignTemplate`

NewNewCampaignTemplate instantiates a new NewCampaignTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewCampaignTemplateWithDefaults

`func NewNewCampaignTemplateWithDefaults() *NewCampaignTemplate`

NewNewCampaignTemplateWithDefaults instantiates a new NewCampaignTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewCampaignTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewCampaignTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewCampaignTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NewCampaignTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewCampaignTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewCampaignTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInstructions

`func (o *NewCampaignTemplate) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *NewCampaignTemplate) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *NewCampaignTemplate) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.


### GetCampaignAttributes

`func (o *NewCampaignTemplate) GetCampaignAttributes() map[string]interface{}`

GetCampaignAttributes returns the CampaignAttributes field if non-nil, zero value otherwise.

### GetCampaignAttributesOk

`func (o *NewCampaignTemplate) GetCampaignAttributesOk() (*map[string]interface{}, bool)`

GetCampaignAttributesOk returns a tuple with the CampaignAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignAttributes

`func (o *NewCampaignTemplate) SetCampaignAttributes(v map[string]interface{})`

SetCampaignAttributes sets CampaignAttributes field to given value.

### HasCampaignAttributes

`func (o *NewCampaignTemplate) HasCampaignAttributes() bool`

HasCampaignAttributes returns a boolean if a field has been set.

### GetCouponAttributes

`func (o *NewCampaignTemplate) GetCouponAttributes() map[string]interface{}`

GetCouponAttributes returns the CouponAttributes field if non-nil, zero value otherwise.

### GetCouponAttributesOk

`func (o *NewCampaignTemplate) GetCouponAttributesOk() (*map[string]interface{}, bool)`

GetCouponAttributesOk returns a tuple with the CouponAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponAttributes

`func (o *NewCampaignTemplate) SetCouponAttributes(v map[string]interface{})`

SetCouponAttributes sets CouponAttributes field to given value.

### HasCouponAttributes

`func (o *NewCampaignTemplate) HasCouponAttributes() bool`

HasCouponAttributes returns a boolean if a field has been set.

### GetState

`func (o *NewCampaignTemplate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NewCampaignTemplate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NewCampaignTemplate) SetState(v string)`

SetState sets State field to given value.


### GetTags

`func (o *NewCampaignTemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NewCampaignTemplate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NewCampaignTemplate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NewCampaignTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetReevaluateOnReturn

`func (o *NewCampaignTemplate) GetReevaluateOnReturn() bool`

GetReevaluateOnReturn returns the ReevaluateOnReturn field if non-nil, zero value otherwise.

### GetReevaluateOnReturnOk

`func (o *NewCampaignTemplate) GetReevaluateOnReturnOk() (*bool, bool)`

GetReevaluateOnReturnOk returns a tuple with the ReevaluateOnReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReevaluateOnReturn

`func (o *NewCampaignTemplate) SetReevaluateOnReturn(v bool)`

SetReevaluateOnReturn sets ReevaluateOnReturn field to given value.

### HasReevaluateOnReturn

`func (o *NewCampaignTemplate) HasReevaluateOnReturn() bool`

HasReevaluateOnReturn returns a boolean if a field has been set.

### GetFeatures

`func (o *NewCampaignTemplate) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *NewCampaignTemplate) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *NewCampaignTemplate) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *NewCampaignTemplate) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetCouponSettings

`func (o *NewCampaignTemplate) GetCouponSettings() CodeGeneratorSettings`

GetCouponSettings returns the CouponSettings field if non-nil, zero value otherwise.

### GetCouponSettingsOk

`func (o *NewCampaignTemplate) GetCouponSettingsOk() (*CodeGeneratorSettings, bool)`

GetCouponSettingsOk returns a tuple with the CouponSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponSettings

`func (o *NewCampaignTemplate) SetCouponSettings(v CodeGeneratorSettings)`

SetCouponSettings sets CouponSettings field to given value.

### HasCouponSettings

`func (o *NewCampaignTemplate) HasCouponSettings() bool`

HasCouponSettings returns a boolean if a field has been set.

### GetCouponReservationSettings

`func (o *NewCampaignTemplate) GetCouponReservationSettings() CampaignTemplateCouponReservationSettings`

GetCouponReservationSettings returns the CouponReservationSettings field if non-nil, zero value otherwise.

### GetCouponReservationSettingsOk

`func (o *NewCampaignTemplate) GetCouponReservationSettingsOk() (*CampaignTemplateCouponReservationSettings, bool)`

GetCouponReservationSettingsOk returns a tuple with the CouponReservationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponReservationSettings

`func (o *NewCampaignTemplate) SetCouponReservationSettings(v CampaignTemplateCouponReservationSettings)`

SetCouponReservationSettings sets CouponReservationSettings field to given value.

### HasCouponReservationSettings

`func (o *NewCampaignTemplate) HasCouponReservationSettings() bool`

HasCouponReservationSettings returns a boolean if a field has been set.

### GetReferralSettings

`func (o *NewCampaignTemplate) GetReferralSettings() CodeGeneratorSettings`

GetReferralSettings returns the ReferralSettings field if non-nil, zero value otherwise.

### GetReferralSettingsOk

`func (o *NewCampaignTemplate) GetReferralSettingsOk() (*CodeGeneratorSettings, bool)`

GetReferralSettingsOk returns a tuple with the ReferralSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralSettings

`func (o *NewCampaignTemplate) SetReferralSettings(v CodeGeneratorSettings)`

SetReferralSettings sets ReferralSettings field to given value.

### HasReferralSettings

`func (o *NewCampaignTemplate) HasReferralSettings() bool`

HasReferralSettings returns a boolean if a field has been set.

### GetLimits

`func (o *NewCampaignTemplate) GetLimits() []TemplateLimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *NewCampaignTemplate) GetLimitsOk() (*[]TemplateLimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *NewCampaignTemplate) SetLimits(v []TemplateLimitConfig)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *NewCampaignTemplate) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetTemplateParams

`func (o *NewCampaignTemplate) GetTemplateParams() []CampaignTemplateParams`

GetTemplateParams returns the TemplateParams field if non-nil, zero value otherwise.

### GetTemplateParamsOk

`func (o *NewCampaignTemplate) GetTemplateParamsOk() (*[]CampaignTemplateParams, bool)`

GetTemplateParamsOk returns a tuple with the TemplateParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateParams

`func (o *NewCampaignTemplate) SetTemplateParams(v []CampaignTemplateParams)`

SetTemplateParams sets TemplateParams field to given value.

### HasTemplateParams

`func (o *NewCampaignTemplate) HasTemplateParams() bool`

HasTemplateParams returns a boolean if a field has been set.

### GetCampaignCollections

`func (o *NewCampaignTemplate) GetCampaignCollections() []CampaignTemplateCollection`

GetCampaignCollections returns the CampaignCollections field if non-nil, zero value otherwise.

### GetCampaignCollectionsOk

`func (o *NewCampaignTemplate) GetCampaignCollectionsOk() (*[]CampaignTemplateCollection, bool)`

GetCampaignCollectionsOk returns a tuple with the CampaignCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignCollections

`func (o *NewCampaignTemplate) SetCampaignCollections(v []CampaignTemplateCollection)`

SetCampaignCollections sets CampaignCollections field to given value.

### HasCampaignCollections

`func (o *NewCampaignTemplate) HasCampaignCollections() bool`

HasCampaignCollections returns a boolean if a field has been set.

### GetDefaultCampaignGroupId

`func (o *NewCampaignTemplate) GetDefaultCampaignGroupId() int64`

GetDefaultCampaignGroupId returns the DefaultCampaignGroupId field if non-nil, zero value otherwise.

### GetDefaultCampaignGroupIdOk

`func (o *NewCampaignTemplate) GetDefaultCampaignGroupIdOk() (*int64, bool)`

GetDefaultCampaignGroupIdOk returns a tuple with the DefaultCampaignGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCampaignGroupId

`func (o *NewCampaignTemplate) SetDefaultCampaignGroupId(v int64)`

SetDefaultCampaignGroupId sets DefaultCampaignGroupId field to given value.

### HasDefaultCampaignGroupId

`func (o *NewCampaignTemplate) HasDefaultCampaignGroupId() bool`

HasDefaultCampaignGroupId returns a boolean if a field has been set.

### GetCampaignType

`func (o *NewCampaignTemplate) GetCampaignType() string`

GetCampaignType returns the CampaignType field if non-nil, zero value otherwise.

### GetCampaignTypeOk

`func (o *NewCampaignTemplate) GetCampaignTypeOk() (*string, bool)`

GetCampaignTypeOk returns a tuple with the CampaignType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignType

`func (o *NewCampaignTemplate) SetCampaignType(v string)`

SetCampaignType sets CampaignType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


