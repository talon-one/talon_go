# RevisionVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**AccountId** | Pointer to **int64** |  | 
**ApplicationId** | Pointer to **int64** |  | 
**CampaignId** | Pointer to **int64** |  | 
**Created** | Pointer to [**time.Time**](time.Time.md) |  | 
**CreatedBy** | Pointer to **int64** |  | 
**RevisionId** | Pointer to **int64** |  | 
**Version** | Pointer to **int64** |  | 
**Name** | Pointer to **string** | A user-facing name for this campaign. | [optional] 
**StartTime** | Pointer to [**NullableTime**](time.Time.md) | Timestamp when the campaign will become active. | [optional] 
**EndTime** | Pointer to [**NullableTime**](time.Time.md) | Timestamp when the campaign will become inactive. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this campaign. | [optional] 
**Description** | Pointer to **NullableString** | A detailed description of the campaign. | [optional] 
**ActiveRulesetId** | Pointer to **NullableInt32** | The ID of the ruleset this campaign will use. | [optional] 
**Tags** | Pointer to **[]string** | A list of tags for the campaign. | [optional] 
**CouponSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**ReferralSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**Limits** | Pointer to [**[]LimitConfig**](LimitConfig.md) | The set of limits that will operate for this campaign version. | [optional] 
**ReevaluateOnReturn** | Pointer to **bool** | Indicates whether this campaign should be reevaluated when a customer returns an item. | [optional] 
**Features** | Pointer to **[]string** | A list of features for the campaign. | [optional] 

## Methods

### NewRevisionVersion

`func NewRevisionVersion(id int64, accountId int64, applicationId int64, campaignId int64, created time.Time, createdBy int64, revisionId int64, version int64, ) *RevisionVersion`

NewRevisionVersion instantiates a new RevisionVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionVersionWithDefaults

`func NewRevisionVersionWithDefaults() *RevisionVersion`

NewRevisionVersionWithDefaults instantiates a new RevisionVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RevisionVersion) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RevisionVersion) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RevisionVersion) SetId(v int64)`

SetId sets Id field to given value.


### GetAccountId

`func (o *RevisionVersion) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RevisionVersion) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RevisionVersion) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetApplicationId

`func (o *RevisionVersion) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *RevisionVersion) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *RevisionVersion) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetCampaignId

`func (o *RevisionVersion) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *RevisionVersion) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *RevisionVersion) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetCreated

`func (o *RevisionVersion) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RevisionVersion) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RevisionVersion) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCreatedBy

`func (o *RevisionVersion) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RevisionVersion) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RevisionVersion) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.


### GetRevisionId

`func (o *RevisionVersion) GetRevisionId() int64`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *RevisionVersion) GetRevisionIdOk() (*int64, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *RevisionVersion) SetRevisionId(v int64)`

SetRevisionId sets RevisionId field to given value.


### GetVersion

`func (o *RevisionVersion) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RevisionVersion) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RevisionVersion) SetVersion(v int64)`

SetVersion sets Version field to given value.


### GetName

`func (o *RevisionVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RevisionVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RevisionVersion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RevisionVersion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTime

`func (o *RevisionVersion) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RevisionVersion) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RevisionVersion) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *RevisionVersion) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *RevisionVersion) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *RevisionVersion) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *RevisionVersion) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *RevisionVersion) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *RevisionVersion) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *RevisionVersion) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *RevisionVersion) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *RevisionVersion) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetAttributes

`func (o *RevisionVersion) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RevisionVersion) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RevisionVersion) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *RevisionVersion) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDescription

`func (o *RevisionVersion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RevisionVersion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RevisionVersion) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RevisionVersion) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RevisionVersion) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RevisionVersion) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetActiveRulesetId

`func (o *RevisionVersion) GetActiveRulesetId() int32`

GetActiveRulesetId returns the ActiveRulesetId field if non-nil, zero value otherwise.

### GetActiveRulesetIdOk

`func (o *RevisionVersion) GetActiveRulesetIdOk() (*int32, bool)`

GetActiveRulesetIdOk returns a tuple with the ActiveRulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesetId

`func (o *RevisionVersion) SetActiveRulesetId(v int32)`

SetActiveRulesetId sets ActiveRulesetId field to given value.

### HasActiveRulesetId

`func (o *RevisionVersion) HasActiveRulesetId() bool`

HasActiveRulesetId returns a boolean if a field has been set.

### SetActiveRulesetIdNil

`func (o *RevisionVersion) SetActiveRulesetIdNil(b bool)`

 SetActiveRulesetIdNil sets the value for ActiveRulesetId to be an explicit nil

### UnsetActiveRulesetId
`func (o *RevisionVersion) UnsetActiveRulesetId()`

UnsetActiveRulesetId ensures that no value is present for ActiveRulesetId, not even an explicit nil
### GetTags

`func (o *RevisionVersion) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RevisionVersion) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RevisionVersion) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RevisionVersion) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCouponSettings

`func (o *RevisionVersion) GetCouponSettings() CodeGeneratorSettings`

GetCouponSettings returns the CouponSettings field if non-nil, zero value otherwise.

### GetCouponSettingsOk

`func (o *RevisionVersion) GetCouponSettingsOk() (*CodeGeneratorSettings, bool)`

GetCouponSettingsOk returns a tuple with the CouponSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponSettings

`func (o *RevisionVersion) SetCouponSettings(v CodeGeneratorSettings)`

SetCouponSettings sets CouponSettings field to given value.

### HasCouponSettings

`func (o *RevisionVersion) HasCouponSettings() bool`

HasCouponSettings returns a boolean if a field has been set.

### GetReferralSettings

`func (o *RevisionVersion) GetReferralSettings() CodeGeneratorSettings`

GetReferralSettings returns the ReferralSettings field if non-nil, zero value otherwise.

### GetReferralSettingsOk

`func (o *RevisionVersion) GetReferralSettingsOk() (*CodeGeneratorSettings, bool)`

GetReferralSettingsOk returns a tuple with the ReferralSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralSettings

`func (o *RevisionVersion) SetReferralSettings(v CodeGeneratorSettings)`

SetReferralSettings sets ReferralSettings field to given value.

### HasReferralSettings

`func (o *RevisionVersion) HasReferralSettings() bool`

HasReferralSettings returns a boolean if a field has been set.

### GetLimits

`func (o *RevisionVersion) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *RevisionVersion) GetLimitsOk() (*[]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *RevisionVersion) SetLimits(v []LimitConfig)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *RevisionVersion) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetReevaluateOnReturn

`func (o *RevisionVersion) GetReevaluateOnReturn() bool`

GetReevaluateOnReturn returns the ReevaluateOnReturn field if non-nil, zero value otherwise.

### GetReevaluateOnReturnOk

`func (o *RevisionVersion) GetReevaluateOnReturnOk() (*bool, bool)`

GetReevaluateOnReturnOk returns a tuple with the ReevaluateOnReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReevaluateOnReturn

`func (o *RevisionVersion) SetReevaluateOnReturn(v bool)`

SetReevaluateOnReturn sets ReevaluateOnReturn field to given value.

### HasReevaluateOnReturn

`func (o *RevisionVersion) HasReevaluateOnReturn() bool`

HasReevaluateOnReturn returns a boolean if a field has been set.

### GetFeatures

`func (o *RevisionVersion) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *RevisionVersion) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *RevisionVersion) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *RevisionVersion) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


