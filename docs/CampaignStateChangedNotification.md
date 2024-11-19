# CampaignStateChangedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Campaign** | Pointer to [**Campaign**](Campaign.md) |  | 
**OldState** | Pointer to **string** | The campaign&#39;s old state. Can be one of the following: [&#39;running&#39;, &#39;disabled&#39;, &#39;scheduled&#39;, &#39;expired&#39;, &#39;archived&#39;]  | 
**NewState** | Pointer to **string** | The campaign&#39;s new state. Can be one of the following: [&#39;running&#39;, &#39;disabled&#39;, &#39;scheduled&#39;, &#39;expired&#39;, &#39;archived&#39;]  | 
**Ruleset** | Pointer to [**Ruleset**](Ruleset.md) |  | [optional] 

## Methods

### GetCampaign

`func (o *CampaignStateChangedNotification) GetCampaign() Campaign`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *CampaignStateChangedNotification) GetCampaignOk() (Campaign, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaign

`func (o *CampaignStateChangedNotification) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### SetCampaign

`func (o *CampaignStateChangedNotification) SetCampaign(v Campaign)`

SetCampaign gets a reference to the given Campaign and assigns it to the Campaign field.

### GetOldState

`func (o *CampaignStateChangedNotification) GetOldState() string`

GetOldState returns the OldState field if non-nil, zero value otherwise.

### GetOldStateOk

`func (o *CampaignStateChangedNotification) GetOldStateOk() (string, bool)`

GetOldStateOk returns a tuple with the OldState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOldState

`func (o *CampaignStateChangedNotification) HasOldState() bool`

HasOldState returns a boolean if a field has been set.

### SetOldState

`func (o *CampaignStateChangedNotification) SetOldState(v string)`

SetOldState gets a reference to the given string and assigns it to the OldState field.

### GetNewState

`func (o *CampaignStateChangedNotification) GetNewState() string`

GetNewState returns the NewState field if non-nil, zero value otherwise.

### GetNewStateOk

`func (o *CampaignStateChangedNotification) GetNewStateOk() (string, bool)`

GetNewStateOk returns a tuple with the NewState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNewState

`func (o *CampaignStateChangedNotification) HasNewState() bool`

HasNewState returns a boolean if a field has been set.

### SetNewState

`func (o *CampaignStateChangedNotification) SetNewState(v string)`

SetNewState gets a reference to the given string and assigns it to the NewState field.

### GetRuleset

`func (o *CampaignStateChangedNotification) GetRuleset() Ruleset`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *CampaignStateChangedNotification) GetRulesetOk() (Ruleset, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRuleset

`func (o *CampaignStateChangedNotification) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.

### SetRuleset

`func (o *CampaignStateChangedNotification) SetRuleset(v Ruleset)`

SetRuleset gets a reference to the given Ruleset and assigns it to the Ruleset field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


