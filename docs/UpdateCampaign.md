# UpdateCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A user-facing name for this campaign. | 
**Description** | Pointer to **string** | A detailed description of the campaign. | [optional] 
**StartTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the campaign will become active. | [optional] 
**EndTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the campaign will become inactive. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this campaign. | [optional] 
**State** | Pointer to **string** | A disabled or archived campaign is not evaluated for rules or coupons.  | [optional] [default to "enabled"]
**ActiveRulesetId** | Pointer to **int64** | [ID of Ruleset](https://docs.talon.one/management-api#operation/getRulesets) this campaign applies on customer session evaluation.  | [optional] 
**Tags** | Pointer to **[]string** | A list of tags for the campaign. | 
**ReevaluateOnReturn** | Pointer to **bool** | Indicates whether this campaign should be reevaluated when a customer returns an item. | [optional] 
**Features** | Pointer to **[]string** | A list of features for the campaign. | 
**CouponSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**ReferralSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**Limits** | Pointer to [**[]LimitConfig**](LimitConfig.md) | The set of limits that will operate for this campaign. | 
**CampaignGroups** | Pointer to **[]int64** | The IDs of the [campaign groups](https://docs.talon.one/docs/product/account/account-settings/managing-campaign-groups) this campaign belongs to.  | [optional] 
**EvaluationGroupId** | Pointer to **int64** | The ID of the campaign evaluation group the campaign belongs to. | [optional] 
**Type** | Pointer to **string** | The campaign type. Possible type values:   - &#x60;cartItem&#x60;: Type of campaign that can apply effects only to cart items.   - &#x60;advanced&#x60;: Type of campaign that can apply effects to customer sessions and cart items.  | [optional] [default to "advanced"]
**LinkedStoreIds** | Pointer to **[]int64** | A list of store IDs that you want to link to the campaign.  **Note:** - Campaigns with linked store IDs will only be evaluated when there is a [customer session update](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/updateCustomerSessionV2) that references a linked store. - If you linked stores to the campaign by uploading a CSV file, you cannot use this property and it should be empty. - Use of this property is limited to 50 stores. To link more than 50 stores, upload them via a CSV file.  | [optional] 

## Methods

### NewUpdateCampaign

`func NewUpdateCampaign(name string, tags []string, features []string, limits []LimitConfig, ) *UpdateCampaign`

NewUpdateCampaign instantiates a new UpdateCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCampaignWithDefaults

`func NewUpdateCampaignWithDefaults() *UpdateCampaign`

NewUpdateCampaignWithDefaults instantiates a new UpdateCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateCampaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCampaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCampaign) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateCampaign) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCampaign) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCampaign) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCampaign) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartTime

`func (o *UpdateCampaign) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UpdateCampaign) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UpdateCampaign) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UpdateCampaign) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *UpdateCampaign) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UpdateCampaign) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UpdateCampaign) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *UpdateCampaign) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetAttributes

`func (o *UpdateCampaign) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateCampaign) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateCampaign) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UpdateCampaign) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetState

`func (o *UpdateCampaign) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateCampaign) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateCampaign) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UpdateCampaign) HasState() bool`

HasState returns a boolean if a field has been set.

### GetActiveRulesetId

`func (o *UpdateCampaign) GetActiveRulesetId() int64`

GetActiveRulesetId returns the ActiveRulesetId field if non-nil, zero value otherwise.

### GetActiveRulesetIdOk

`func (o *UpdateCampaign) GetActiveRulesetIdOk() (*int64, bool)`

GetActiveRulesetIdOk returns a tuple with the ActiveRulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesetId

`func (o *UpdateCampaign) SetActiveRulesetId(v int64)`

SetActiveRulesetId sets ActiveRulesetId field to given value.

### HasActiveRulesetId

`func (o *UpdateCampaign) HasActiveRulesetId() bool`

HasActiveRulesetId returns a boolean if a field has been set.

### GetTags

`func (o *UpdateCampaign) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateCampaign) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateCampaign) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetReevaluateOnReturn

`func (o *UpdateCampaign) GetReevaluateOnReturn() bool`

GetReevaluateOnReturn returns the ReevaluateOnReturn field if non-nil, zero value otherwise.

### GetReevaluateOnReturnOk

`func (o *UpdateCampaign) GetReevaluateOnReturnOk() (*bool, bool)`

GetReevaluateOnReturnOk returns a tuple with the ReevaluateOnReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReevaluateOnReturn

`func (o *UpdateCampaign) SetReevaluateOnReturn(v bool)`

SetReevaluateOnReturn sets ReevaluateOnReturn field to given value.

### HasReevaluateOnReturn

`func (o *UpdateCampaign) HasReevaluateOnReturn() bool`

HasReevaluateOnReturn returns a boolean if a field has been set.

### GetFeatures

`func (o *UpdateCampaign) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *UpdateCampaign) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *UpdateCampaign) SetFeatures(v []string)`

SetFeatures sets Features field to given value.


### GetCouponSettings

`func (o *UpdateCampaign) GetCouponSettings() CodeGeneratorSettings`

GetCouponSettings returns the CouponSettings field if non-nil, zero value otherwise.

### GetCouponSettingsOk

`func (o *UpdateCampaign) GetCouponSettingsOk() (*CodeGeneratorSettings, bool)`

GetCouponSettingsOk returns a tuple with the CouponSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponSettings

`func (o *UpdateCampaign) SetCouponSettings(v CodeGeneratorSettings)`

SetCouponSettings sets CouponSettings field to given value.

### HasCouponSettings

`func (o *UpdateCampaign) HasCouponSettings() bool`

HasCouponSettings returns a boolean if a field has been set.

### GetReferralSettings

`func (o *UpdateCampaign) GetReferralSettings() CodeGeneratorSettings`

GetReferralSettings returns the ReferralSettings field if non-nil, zero value otherwise.

### GetReferralSettingsOk

`func (o *UpdateCampaign) GetReferralSettingsOk() (*CodeGeneratorSettings, bool)`

GetReferralSettingsOk returns a tuple with the ReferralSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralSettings

`func (o *UpdateCampaign) SetReferralSettings(v CodeGeneratorSettings)`

SetReferralSettings sets ReferralSettings field to given value.

### HasReferralSettings

`func (o *UpdateCampaign) HasReferralSettings() bool`

HasReferralSettings returns a boolean if a field has been set.

### GetLimits

`func (o *UpdateCampaign) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *UpdateCampaign) GetLimitsOk() (*[]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *UpdateCampaign) SetLimits(v []LimitConfig)`

SetLimits sets Limits field to given value.


### GetCampaignGroups

`func (o *UpdateCampaign) GetCampaignGroups() []int64`

GetCampaignGroups returns the CampaignGroups field if non-nil, zero value otherwise.

### GetCampaignGroupsOk

`func (o *UpdateCampaign) GetCampaignGroupsOk() (*[]int64, bool)`

GetCampaignGroupsOk returns a tuple with the CampaignGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignGroups

`func (o *UpdateCampaign) SetCampaignGroups(v []int64)`

SetCampaignGroups sets CampaignGroups field to given value.

### HasCampaignGroups

`func (o *UpdateCampaign) HasCampaignGroups() bool`

HasCampaignGroups returns a boolean if a field has been set.

### GetEvaluationGroupId

`func (o *UpdateCampaign) GetEvaluationGroupId() int64`

GetEvaluationGroupId returns the EvaluationGroupId field if non-nil, zero value otherwise.

### GetEvaluationGroupIdOk

`func (o *UpdateCampaign) GetEvaluationGroupIdOk() (*int64, bool)`

GetEvaluationGroupIdOk returns a tuple with the EvaluationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationGroupId

`func (o *UpdateCampaign) SetEvaluationGroupId(v int64)`

SetEvaluationGroupId sets EvaluationGroupId field to given value.

### HasEvaluationGroupId

`func (o *UpdateCampaign) HasEvaluationGroupId() bool`

HasEvaluationGroupId returns a boolean if a field has been set.

### GetType

`func (o *UpdateCampaign) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateCampaign) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateCampaign) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateCampaign) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinkedStoreIds

`func (o *UpdateCampaign) GetLinkedStoreIds() []int64`

GetLinkedStoreIds returns the LinkedStoreIds field if non-nil, zero value otherwise.

### GetLinkedStoreIdsOk

`func (o *UpdateCampaign) GetLinkedStoreIdsOk() (*[]int64, bool)`

GetLinkedStoreIdsOk returns a tuple with the LinkedStoreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedStoreIds

`func (o *UpdateCampaign) SetLinkedStoreIds(v []int64)`

SetLinkedStoreIds sets LinkedStoreIds field to given value.

### HasLinkedStoreIds

`func (o *UpdateCampaign) HasLinkedStoreIds() bool`

HasLinkedStoreIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


