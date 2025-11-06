# NewCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A user-facing name for this campaign. | 
**Description** | Pointer to **string** | A detailed description of the campaign. | [optional] 
**StartTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the campaign will become active. | [optional] 
**EndTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the campaign will become inactive. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this campaign. | [optional] 
**State** | Pointer to **string** | A disabled or archived campaign is not evaluated for rules or coupons.  | [default to "enabled"]
**ActiveRulesetId** | Pointer to **int64** | [ID of Ruleset](https://docs.talon.one/management-api#operation/getRulesets) this campaign applies on customer session evaluation.  | [optional] 
**Tags** | Pointer to **[]string** | A list of tags for the campaign. | 
**Features** | Pointer to **[]string** | The features enabled in this campaign. | 
**CouponSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**ReferralSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**Limits** | Pointer to [**[]LimitConfig**](LimitConfig.md) | The set of [budget limits](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-budgets) for this campaign.  | 
**CampaignGroups** | Pointer to **[]int64** | The IDs of the [campaign groups](https://docs.talon.one/docs/product/account/managing-campaign-groups) this campaign belongs to.  | [optional] 
**Type** | Pointer to **string** | The campaign type. Possible type values:   - &#x60;cartItem&#x60;: Type of campaign that can apply effects only to cart items.   - &#x60;advanced&#x60;: Type of campaign that can apply effects to customer sessions and cart items.  | [optional] [default to "advanced"]
**LinkedStoreIds** | Pointer to **[]int64** | A list of store IDs that you want to link to the campaign.  **Note:** Campaigns with linked store IDs will only be evaluated when there is a [customer session update](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/updateCustomerSessionV2) that references a linked store.  | [optional] 
**EvaluationGroupId** | Pointer to **int64** | The ID of the campaign evaluation group the campaign belongs to. | [optional] 

## Methods

### NewNewCampaign

`func NewNewCampaign(name string, state string, tags []string, features []string, limits []LimitConfig, ) *NewCampaign`

NewNewCampaign instantiates a new NewCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewCampaignWithDefaults

`func NewNewCampaignWithDefaults() *NewCampaign`

NewNewCampaignWithDefaults instantiates a new NewCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewCampaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewCampaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewCampaign) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NewCampaign) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewCampaign) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewCampaign) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewCampaign) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartTime

`func (o *NewCampaign) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *NewCampaign) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *NewCampaign) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *NewCampaign) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *NewCampaign) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *NewCampaign) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *NewCampaign) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *NewCampaign) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetAttributes

`func (o *NewCampaign) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NewCampaign) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NewCampaign) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *NewCampaign) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetState

`func (o *NewCampaign) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NewCampaign) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NewCampaign) SetState(v string)`

SetState sets State field to given value.


### GetActiveRulesetId

`func (o *NewCampaign) GetActiveRulesetId() int64`

GetActiveRulesetId returns the ActiveRulesetId field if non-nil, zero value otherwise.

### GetActiveRulesetIdOk

`func (o *NewCampaign) GetActiveRulesetIdOk() (*int64, bool)`

GetActiveRulesetIdOk returns a tuple with the ActiveRulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesetId

`func (o *NewCampaign) SetActiveRulesetId(v int64)`

SetActiveRulesetId sets ActiveRulesetId field to given value.

### HasActiveRulesetId

`func (o *NewCampaign) HasActiveRulesetId() bool`

HasActiveRulesetId returns a boolean if a field has been set.

### GetTags

`func (o *NewCampaign) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NewCampaign) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NewCampaign) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetFeatures

`func (o *NewCampaign) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *NewCampaign) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *NewCampaign) SetFeatures(v []string)`

SetFeatures sets Features field to given value.


### GetCouponSettings

`func (o *NewCampaign) GetCouponSettings() CodeGeneratorSettings`

GetCouponSettings returns the CouponSettings field if non-nil, zero value otherwise.

### GetCouponSettingsOk

`func (o *NewCampaign) GetCouponSettingsOk() (*CodeGeneratorSettings, bool)`

GetCouponSettingsOk returns a tuple with the CouponSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponSettings

`func (o *NewCampaign) SetCouponSettings(v CodeGeneratorSettings)`

SetCouponSettings sets CouponSettings field to given value.

### HasCouponSettings

`func (o *NewCampaign) HasCouponSettings() bool`

HasCouponSettings returns a boolean if a field has been set.

### GetReferralSettings

`func (o *NewCampaign) GetReferralSettings() CodeGeneratorSettings`

GetReferralSettings returns the ReferralSettings field if non-nil, zero value otherwise.

### GetReferralSettingsOk

`func (o *NewCampaign) GetReferralSettingsOk() (*CodeGeneratorSettings, bool)`

GetReferralSettingsOk returns a tuple with the ReferralSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralSettings

`func (o *NewCampaign) SetReferralSettings(v CodeGeneratorSettings)`

SetReferralSettings sets ReferralSettings field to given value.

### HasReferralSettings

`func (o *NewCampaign) HasReferralSettings() bool`

HasReferralSettings returns a boolean if a field has been set.

### GetLimits

`func (o *NewCampaign) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *NewCampaign) GetLimitsOk() (*[]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *NewCampaign) SetLimits(v []LimitConfig)`

SetLimits sets Limits field to given value.


### GetCampaignGroups

`func (o *NewCampaign) GetCampaignGroups() []int64`

GetCampaignGroups returns the CampaignGroups field if non-nil, zero value otherwise.

### GetCampaignGroupsOk

`func (o *NewCampaign) GetCampaignGroupsOk() (*[]int64, bool)`

GetCampaignGroupsOk returns a tuple with the CampaignGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignGroups

`func (o *NewCampaign) SetCampaignGroups(v []int64)`

SetCampaignGroups sets CampaignGroups field to given value.

### HasCampaignGroups

`func (o *NewCampaign) HasCampaignGroups() bool`

HasCampaignGroups returns a boolean if a field has been set.

### GetType

`func (o *NewCampaign) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewCampaign) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NewCampaign) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NewCampaign) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinkedStoreIds

`func (o *NewCampaign) GetLinkedStoreIds() []int64`

GetLinkedStoreIds returns the LinkedStoreIds field if non-nil, zero value otherwise.

### GetLinkedStoreIdsOk

`func (o *NewCampaign) GetLinkedStoreIdsOk() (*[]int64, bool)`

GetLinkedStoreIdsOk returns a tuple with the LinkedStoreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedStoreIds

`func (o *NewCampaign) SetLinkedStoreIds(v []int64)`

SetLinkedStoreIds sets LinkedStoreIds field to given value.

### HasLinkedStoreIds

`func (o *NewCampaign) HasLinkedStoreIds() bool`

HasLinkedStoreIds returns a boolean if a field has been set.

### GetEvaluationGroupId

`func (o *NewCampaign) GetEvaluationGroupId() int64`

GetEvaluationGroupId returns the EvaluationGroupId field if non-nil, zero value otherwise.

### GetEvaluationGroupIdOk

`func (o *NewCampaign) GetEvaluationGroupIdOk() (*int64, bool)`

GetEvaluationGroupIdOk returns a tuple with the EvaluationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationGroupId

`func (o *NewCampaign) SetEvaluationGroupId(v int64)`

SetEvaluationGroupId sets EvaluationGroupId field to given value.

### HasEvaluationGroupId

`func (o *NewCampaign) HasEvaluationGroupId() bool`

HasEvaluationGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


