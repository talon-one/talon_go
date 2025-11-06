# CreateTemplateCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A user-facing name for this campaign. | 
**Description** | Pointer to **string** | A detailed description of the campaign. | [optional] 
**TemplateId** | Pointer to **int64** | The ID of the Campaign Template which will be used in order to create the Campaign. | 
**CampaignAttributesOverrides** | Pointer to [**map[string]interface{}**](.md) | Custom Campaign Attributes. If the Campaign Template defines the same values, they will be overridden. | [optional] 
**TemplateParamValues** | Pointer to [**[]Binding**](Binding.md) | Actual values to replace the template placeholder values in the Ruleset bindings. Values for all Template Parameters must be provided. | [optional] 
**LimitOverrides** | Pointer to [**[]LimitConfig**](LimitConfig.md) | Limits for this Campaign. If the Campaign Template or Application define default values for the same limits, they will be overridden. | [optional] 
**CampaignGroups** | Pointer to **[]int64** | The IDs of the [campaign groups](https://docs.talon.one/docs/product/account/account-settings/managing-campaign-groups) this campaign belongs to.  | [optional] 
**Tags** | Pointer to **[]string** | A list of tags for the campaign. If the campaign template has tags, they will be overridden by this list. | [optional] 
**EvaluationGroupId** | Pointer to **int64** | The ID of the campaign evaluation group the campaign belongs to. | [optional] 
**LinkedStoreIds** | Pointer to **[]int64** | A list of store IDs that are linked to the campaign.  **Note:** Campaigns with linked store IDs will only be evaluated when there is a [customer session update](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/updateCustomerSessionV2) that references a linked store.  | [optional] 

## Methods

### NewCreateTemplateCampaign

`func NewCreateTemplateCampaign(name string, templateId int64, ) *CreateTemplateCampaign`

NewCreateTemplateCampaign instantiates a new CreateTemplateCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTemplateCampaignWithDefaults

`func NewCreateTemplateCampaignWithDefaults() *CreateTemplateCampaign`

NewCreateTemplateCampaignWithDefaults instantiates a new CreateTemplateCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTemplateCampaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTemplateCampaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTemplateCampaign) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateTemplateCampaign) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTemplateCampaign) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTemplateCampaign) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTemplateCampaign) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTemplateId

`func (o *CreateTemplateCampaign) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CreateTemplateCampaign) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *CreateTemplateCampaign) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.


### GetCampaignAttributesOverrides

`func (o *CreateTemplateCampaign) GetCampaignAttributesOverrides() map[string]interface{}`

GetCampaignAttributesOverrides returns the CampaignAttributesOverrides field if non-nil, zero value otherwise.

### GetCampaignAttributesOverridesOk

`func (o *CreateTemplateCampaign) GetCampaignAttributesOverridesOk() (*map[string]interface{}, bool)`

GetCampaignAttributesOverridesOk returns a tuple with the CampaignAttributesOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignAttributesOverrides

`func (o *CreateTemplateCampaign) SetCampaignAttributesOverrides(v map[string]interface{})`

SetCampaignAttributesOverrides sets CampaignAttributesOverrides field to given value.

### HasCampaignAttributesOverrides

`func (o *CreateTemplateCampaign) HasCampaignAttributesOverrides() bool`

HasCampaignAttributesOverrides returns a boolean if a field has been set.

### GetTemplateParamValues

`func (o *CreateTemplateCampaign) GetTemplateParamValues() []Binding`

GetTemplateParamValues returns the TemplateParamValues field if non-nil, zero value otherwise.

### GetTemplateParamValuesOk

`func (o *CreateTemplateCampaign) GetTemplateParamValuesOk() (*[]Binding, bool)`

GetTemplateParamValuesOk returns a tuple with the TemplateParamValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateParamValues

`func (o *CreateTemplateCampaign) SetTemplateParamValues(v []Binding)`

SetTemplateParamValues sets TemplateParamValues field to given value.

### HasTemplateParamValues

`func (o *CreateTemplateCampaign) HasTemplateParamValues() bool`

HasTemplateParamValues returns a boolean if a field has been set.

### GetLimitOverrides

`func (o *CreateTemplateCampaign) GetLimitOverrides() []LimitConfig`

GetLimitOverrides returns the LimitOverrides field if non-nil, zero value otherwise.

### GetLimitOverridesOk

`func (o *CreateTemplateCampaign) GetLimitOverridesOk() (*[]LimitConfig, bool)`

GetLimitOverridesOk returns a tuple with the LimitOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitOverrides

`func (o *CreateTemplateCampaign) SetLimitOverrides(v []LimitConfig)`

SetLimitOverrides sets LimitOverrides field to given value.

### HasLimitOverrides

`func (o *CreateTemplateCampaign) HasLimitOverrides() bool`

HasLimitOverrides returns a boolean if a field has been set.

### GetCampaignGroups

`func (o *CreateTemplateCampaign) GetCampaignGroups() []int64`

GetCampaignGroups returns the CampaignGroups field if non-nil, zero value otherwise.

### GetCampaignGroupsOk

`func (o *CreateTemplateCampaign) GetCampaignGroupsOk() (*[]int64, bool)`

GetCampaignGroupsOk returns a tuple with the CampaignGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignGroups

`func (o *CreateTemplateCampaign) SetCampaignGroups(v []int64)`

SetCampaignGroups sets CampaignGroups field to given value.

### HasCampaignGroups

`func (o *CreateTemplateCampaign) HasCampaignGroups() bool`

HasCampaignGroups returns a boolean if a field has been set.

### GetTags

`func (o *CreateTemplateCampaign) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateTemplateCampaign) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateTemplateCampaign) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateTemplateCampaign) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEvaluationGroupId

`func (o *CreateTemplateCampaign) GetEvaluationGroupId() int64`

GetEvaluationGroupId returns the EvaluationGroupId field if non-nil, zero value otherwise.

### GetEvaluationGroupIdOk

`func (o *CreateTemplateCampaign) GetEvaluationGroupIdOk() (*int64, bool)`

GetEvaluationGroupIdOk returns a tuple with the EvaluationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationGroupId

`func (o *CreateTemplateCampaign) SetEvaluationGroupId(v int64)`

SetEvaluationGroupId sets EvaluationGroupId field to given value.

### HasEvaluationGroupId

`func (o *CreateTemplateCampaign) HasEvaluationGroupId() bool`

HasEvaluationGroupId returns a boolean if a field has been set.

### GetLinkedStoreIds

`func (o *CreateTemplateCampaign) GetLinkedStoreIds() []int64`

GetLinkedStoreIds returns the LinkedStoreIds field if non-nil, zero value otherwise.

### GetLinkedStoreIdsOk

`func (o *CreateTemplateCampaign) GetLinkedStoreIdsOk() (*[]int64, bool)`

GetLinkedStoreIdsOk returns a tuple with the LinkedStoreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedStoreIds

`func (o *CreateTemplateCampaign) SetLinkedStoreIds(v []int64)`

SetLinkedStoreIds sets LinkedStoreIds field to given value.

### HasLinkedStoreIds

`func (o *CreateTemplateCampaign) HasLinkedStoreIds() bool`

HasLinkedStoreIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


