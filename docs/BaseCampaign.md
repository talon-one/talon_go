# BaseCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** | A list of tags for the campaign. | 
**ActiveRulesetId** | Pointer to **int32** | [ID of Ruleset](https://docs.talon.one/management-api#operation/getRulesets) this campaign applies on customer session evaluation.  | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this campaign. | [optional] 
**CampaignGroups** | Pointer to **[]int32** | The IDs of the [campaign groups](https://docs.talon.one/docs/product/account/managing-campaign-groups) this campaign belongs to.  | [optional] 
**CouponSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**Description** | Pointer to **string** | A detailed description of the campaign. | [optional] 
**EndTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the campaign will become inactive. | [optional] 
**Features** | Pointer to **[]string** | The features enabled in this campaign. | 
**Limits** | Pointer to [**[]LimitConfig**](LimitConfig.md) | The set of [budget limits](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-budgets) for this campaign.  | 
**LinkedStoreIds** | Pointer to **[]int32** | A list of store IDs that you want to link to the campaign.  **Note:** Campaigns with linked store IDs will only be evaluated when there is a [customer session update](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/updateCustomerSessionV2) that references a linked store.  | [optional] 
**Name** | Pointer to **string** | A user-facing name for this campaign. | 
**ReferralSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**StartTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the campaign will become active. | [optional] 
**State** | Pointer to **string** | A disabled or archived campaign is not evaluated for rules or coupons.  | [default to STATE_ENABLED]
**Type** | Pointer to **string** | The campaign type. Possible type values:   - &#x60;cartItem&#x60;: Type of campaign that can apply effects only to cart items.   - &#x60;advanced&#x60;: Type of campaign that can apply effects to customer sessions and cart items.  | [optional] [default to TYPE_ADVANCED]

## Methods

### GetTags

`func (o *BaseCampaign) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BaseCampaign) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *BaseCampaign) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *BaseCampaign) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.

### GetActiveRulesetId

`func (o *BaseCampaign) GetActiveRulesetId() int32`

GetActiveRulesetId returns the ActiveRulesetId field if non-nil, zero value otherwise.

### GetActiveRulesetIdOk

`func (o *BaseCampaign) GetActiveRulesetIdOk() (int32, bool)`

GetActiveRulesetIdOk returns a tuple with the ActiveRulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveRulesetId

`func (o *BaseCampaign) HasActiveRulesetId() bool`

HasActiveRulesetId returns a boolean if a field has been set.

### SetActiveRulesetId

`func (o *BaseCampaign) SetActiveRulesetId(v int32)`

SetActiveRulesetId gets a reference to the given int32 and assigns it to the ActiveRulesetId field.

### GetAttributes

`func (o *BaseCampaign) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BaseCampaign) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *BaseCampaign) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *BaseCampaign) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetCampaignGroups

`func (o *BaseCampaign) GetCampaignGroups() []int32`

GetCampaignGroups returns the CampaignGroups field if non-nil, zero value otherwise.

### GetCampaignGroupsOk

`func (o *BaseCampaign) GetCampaignGroupsOk() ([]int32, bool)`

GetCampaignGroupsOk returns a tuple with the CampaignGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignGroups

`func (o *BaseCampaign) HasCampaignGroups() bool`

HasCampaignGroups returns a boolean if a field has been set.

### SetCampaignGroups

`func (o *BaseCampaign) SetCampaignGroups(v []int32)`

SetCampaignGroups gets a reference to the given []int32 and assigns it to the CampaignGroups field.

### GetCouponSettings

`func (o *BaseCampaign) GetCouponSettings() CodeGeneratorSettings`

GetCouponSettings returns the CouponSettings field if non-nil, zero value otherwise.

### GetCouponSettingsOk

`func (o *BaseCampaign) GetCouponSettingsOk() (CodeGeneratorSettings, bool)`

GetCouponSettingsOk returns a tuple with the CouponSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponSettings

`func (o *BaseCampaign) HasCouponSettings() bool`

HasCouponSettings returns a boolean if a field has been set.

### SetCouponSettings

`func (o *BaseCampaign) SetCouponSettings(v CodeGeneratorSettings)`

SetCouponSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CouponSettings field.

### GetDescription

`func (o *BaseCampaign) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseCampaign) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *BaseCampaign) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *BaseCampaign) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetEndTime

`func (o *BaseCampaign) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *BaseCampaign) GetEndTimeOk() (time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndTime

`func (o *BaseCampaign) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTime

`func (o *BaseCampaign) SetEndTime(v time.Time)`

SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.

### GetFeatures

`func (o *BaseCampaign) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *BaseCampaign) GetFeaturesOk() ([]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatures

`func (o *BaseCampaign) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeatures

`func (o *BaseCampaign) SetFeatures(v []string)`

SetFeatures gets a reference to the given []string and assigns it to the Features field.

### GetLimits

`func (o *BaseCampaign) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *BaseCampaign) GetLimitsOk() ([]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimits

`func (o *BaseCampaign) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimits

`func (o *BaseCampaign) SetLimits(v []LimitConfig)`

SetLimits gets a reference to the given []LimitConfig and assigns it to the Limits field.

### GetLinkedStoreIds

`func (o *BaseCampaign) GetLinkedStoreIds() []int32`

GetLinkedStoreIds returns the LinkedStoreIds field if non-nil, zero value otherwise.

### GetLinkedStoreIdsOk

`func (o *BaseCampaign) GetLinkedStoreIdsOk() ([]int32, bool)`

GetLinkedStoreIdsOk returns a tuple with the LinkedStoreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkedStoreIds

`func (o *BaseCampaign) HasLinkedStoreIds() bool`

HasLinkedStoreIds returns a boolean if a field has been set.

### SetLinkedStoreIds

`func (o *BaseCampaign) SetLinkedStoreIds(v []int32)`

SetLinkedStoreIds gets a reference to the given []int32 and assigns it to the LinkedStoreIds field.

### GetName

`func (o *BaseCampaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseCampaign) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *BaseCampaign) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *BaseCampaign) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetReferralSettings

`func (o *BaseCampaign) GetReferralSettings() CodeGeneratorSettings`

GetReferralSettings returns the ReferralSettings field if non-nil, zero value otherwise.

### GetReferralSettingsOk

`func (o *BaseCampaign) GetReferralSettingsOk() (CodeGeneratorSettings, bool)`

GetReferralSettingsOk returns a tuple with the ReferralSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralSettings

`func (o *BaseCampaign) HasReferralSettings() bool`

HasReferralSettings returns a boolean if a field has been set.

### SetReferralSettings

`func (o *BaseCampaign) SetReferralSettings(v CodeGeneratorSettings)`

SetReferralSettings gets a reference to the given CodeGeneratorSettings and assigns it to the ReferralSettings field.

### GetStartTime

`func (o *BaseCampaign) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BaseCampaign) GetStartTimeOk() (time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartTime

`func (o *BaseCampaign) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTime

`func (o *BaseCampaign) SetStartTime(v time.Time)`

SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.

### GetState

`func (o *BaseCampaign) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BaseCampaign) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *BaseCampaign) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *BaseCampaign) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetType

`func (o *BaseCampaign) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseCampaign) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *BaseCampaign) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *BaseCampaign) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


