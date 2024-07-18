# RulesetAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | Pointer to **int32** | The ID of the campaign that owns this entity. | [optional] 
**TemplateId** | Pointer to **int32** | The ID of the campaign template that owns this entity. | [optional] 
**ActivatedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp indicating when this Ruleset was activated. | [optional] 

## Methods

### GetCampaignId

`func (o *RulesetAllOf) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *RulesetAllOf) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *RulesetAllOf) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *RulesetAllOf) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.

### GetTemplateId

`func (o *RulesetAllOf) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *RulesetAllOf) GetTemplateIdOk() (int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTemplateId

`func (o *RulesetAllOf) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateId

`func (o *RulesetAllOf) SetTemplateId(v int32)`

SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.

### GetActivatedAt

`func (o *RulesetAllOf) GetActivatedAt() time.Time`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *RulesetAllOf) GetActivatedAtOk() (time.Time, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivatedAt

`func (o *RulesetAllOf) HasActivatedAt() bool`

HasActivatedAt returns a boolean if a field has been set.

### SetActivatedAt

`func (o *RulesetAllOf) SetActivatedAt(v time.Time)`

SetActivatedAt gets a reference to the given time.Time and assigns it to the ActivatedAt field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


