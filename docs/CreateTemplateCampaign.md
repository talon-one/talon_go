# CreateTemplateCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A user-facing name for this campaign. | 
**Description** | Pointer to **string** | A detailed description of the campaign. | [optional] 
**TemplateId** | Pointer to **int32** | The ID of the Campaign Template which will be used in order to create the Campaign. | 
**CampaignAttributesOverrides** | Pointer to [**map[string]interface{}**](.md) | Custom Campaign Attributes. If the Campaign Template defines the same values, they will be overridden. | [optional] 
**TemplateParamValues** | Pointer to [**[]Binding**](Binding.md) | Actual values to replace the template placeholder values in the Ruleset bindings. Values for all Template Parameters must be provided. | [optional] 
**LimitOverrides** | Pointer to [**[]LimitConfig**](LimitConfig.md) | Limits for this Campaign. If the Campaign Template or Application define default values for the same limits, they will be overridden. | [optional] 
**CampaignGroups** | Pointer to **[]int32** | The IDs of the [campaign groups](https://docs.talon.one/docs/product/account/managing-campaign-groups/) this campaign belongs to.  | [optional] 
**Tags** | Pointer to **[]string** | A list of tags for the campaign. If the campaign template has tags, they will be overridden by this list. | [optional] 

## Methods

### GetName

`func (o *CreateTemplateCampaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTemplateCampaign) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *CreateTemplateCampaign) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *CreateTemplateCampaign) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *CreateTemplateCampaign) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTemplateCampaign) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *CreateTemplateCampaign) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *CreateTemplateCampaign) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetTemplateId

`func (o *CreateTemplateCampaign) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CreateTemplateCampaign) GetTemplateIdOk() (int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTemplateId

`func (o *CreateTemplateCampaign) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateId

`func (o *CreateTemplateCampaign) SetTemplateId(v int32)`

SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.

### GetCampaignAttributesOverrides

`func (o *CreateTemplateCampaign) GetCampaignAttributesOverrides() map[string]interface{}`

GetCampaignAttributesOverrides returns the CampaignAttributesOverrides field if non-nil, zero value otherwise.

### GetCampaignAttributesOverridesOk

`func (o *CreateTemplateCampaign) GetCampaignAttributesOverridesOk() (map[string]interface{}, bool)`

GetCampaignAttributesOverridesOk returns a tuple with the CampaignAttributesOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignAttributesOverrides

`func (o *CreateTemplateCampaign) HasCampaignAttributesOverrides() bool`

HasCampaignAttributesOverrides returns a boolean if a field has been set.

### SetCampaignAttributesOverrides

`func (o *CreateTemplateCampaign) SetCampaignAttributesOverrides(v map[string]interface{})`

SetCampaignAttributesOverrides gets a reference to the given map[string]interface{} and assigns it to the CampaignAttributesOverrides field.

### GetTemplateParamValues

`func (o *CreateTemplateCampaign) GetTemplateParamValues() []Binding`

GetTemplateParamValues returns the TemplateParamValues field if non-nil, zero value otherwise.

### GetTemplateParamValuesOk

`func (o *CreateTemplateCampaign) GetTemplateParamValuesOk() ([]Binding, bool)`

GetTemplateParamValuesOk returns a tuple with the TemplateParamValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTemplateParamValues

`func (o *CreateTemplateCampaign) HasTemplateParamValues() bool`

HasTemplateParamValues returns a boolean if a field has been set.

### SetTemplateParamValues

`func (o *CreateTemplateCampaign) SetTemplateParamValues(v []Binding)`

SetTemplateParamValues gets a reference to the given []Binding and assigns it to the TemplateParamValues field.

### GetLimitOverrides

`func (o *CreateTemplateCampaign) GetLimitOverrides() []LimitConfig`

GetLimitOverrides returns the LimitOverrides field if non-nil, zero value otherwise.

### GetLimitOverridesOk

`func (o *CreateTemplateCampaign) GetLimitOverridesOk() ([]LimitConfig, bool)`

GetLimitOverridesOk returns a tuple with the LimitOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimitOverrides

`func (o *CreateTemplateCampaign) HasLimitOverrides() bool`

HasLimitOverrides returns a boolean if a field has been set.

### SetLimitOverrides

`func (o *CreateTemplateCampaign) SetLimitOverrides(v []LimitConfig)`

SetLimitOverrides gets a reference to the given []LimitConfig and assigns it to the LimitOverrides field.

### GetCampaignGroups

`func (o *CreateTemplateCampaign) GetCampaignGroups() []int32`

GetCampaignGroups returns the CampaignGroups field if non-nil, zero value otherwise.

### GetCampaignGroupsOk

`func (o *CreateTemplateCampaign) GetCampaignGroupsOk() ([]int32, bool)`

GetCampaignGroupsOk returns a tuple with the CampaignGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignGroups

`func (o *CreateTemplateCampaign) HasCampaignGroups() bool`

HasCampaignGroups returns a boolean if a field has been set.

### SetCampaignGroups

`func (o *CreateTemplateCampaign) SetCampaignGroups(v []int32)`

SetCampaignGroups gets a reference to the given []int32 and assigns it to the CampaignGroups field.

### GetTags

`func (o *CreateTemplateCampaign) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateTemplateCampaign) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *CreateTemplateCampaign) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *CreateTemplateCampaign) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


