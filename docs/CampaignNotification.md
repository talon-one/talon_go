# CampaignNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** | The type of the event. Can be one of the following: [&#39;campaign_state_changed&#39;, &#39;campaign_ruleset_changed&#39;, &#39;campaign_edited&#39;, &#39;campaign_created&#39;, &#39;campaign_deleted&#39;]  | 
**Campaign** | Pointer to [**Campaign**](.md) |  | 
**OldState** | Pointer to **string** | The campaign&#39;s old state. Can be one of the following: [&#39;running&#39;, &#39;disabled&#39;, &#39;scheduled&#39;, &#39;expired&#39;, &#39;draft&#39;, &#39;archived&#39;]  | 
**NewState** | Pointer to **string** | The campaign&#39;s new state. Can be one of the following: [&#39;running&#39;, &#39;disabled&#39;, &#39;scheduled&#39;, &#39;expired&#39;, &#39;draft&#39;, &#39;archived&#39;]  | 
**Ruleset** | Pointer to [**Ruleset**](.md) |  | [optional] 
**OldRuleset** | Pointer to [**Ruleset**](Ruleset.md) |  | [optional] 
**OldCampaign** | Pointer to [**Campaign**](.md) |  | 
**EvaluationPosition** | Pointer to [**CampaignEvaluationPosition**](.md) |  | 
**DeletedAt** | Pointer to [**time.Time**](time.Time.md) | Time when the campaign was deleted. | 

## Methods

### GetEvent

`func (o *CampaignNotification) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CampaignNotification) GetEventOk() (string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvent

`func (o *CampaignNotification) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEvent

`func (o *CampaignNotification) SetEvent(v string)`

SetEvent gets a reference to the given string and assigns it to the Event field.

### GetCampaign

`func (o *CampaignNotification) GetCampaign() Campaign`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *CampaignNotification) GetCampaignOk() (Campaign, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaign

`func (o *CampaignNotification) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### SetCampaign

`func (o *CampaignNotification) SetCampaign(v Campaign)`

SetCampaign gets a reference to the given Campaign and assigns it to the Campaign field.

### GetOldState

`func (o *CampaignNotification) GetOldState() string`

GetOldState returns the OldState field if non-nil, zero value otherwise.

### GetOldStateOk

`func (o *CampaignNotification) GetOldStateOk() (string, bool)`

GetOldStateOk returns a tuple with the OldState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOldState

`func (o *CampaignNotification) HasOldState() bool`

HasOldState returns a boolean if a field has been set.

### SetOldState

`func (o *CampaignNotification) SetOldState(v string)`

SetOldState gets a reference to the given string and assigns it to the OldState field.

### GetNewState

`func (o *CampaignNotification) GetNewState() string`

GetNewState returns the NewState field if non-nil, zero value otherwise.

### GetNewStateOk

`func (o *CampaignNotification) GetNewStateOk() (string, bool)`

GetNewStateOk returns a tuple with the NewState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNewState

`func (o *CampaignNotification) HasNewState() bool`

HasNewState returns a boolean if a field has been set.

### SetNewState

`func (o *CampaignNotification) SetNewState(v string)`

SetNewState gets a reference to the given string and assigns it to the NewState field.

### GetRuleset

`func (o *CampaignNotification) GetRuleset() Ruleset`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *CampaignNotification) GetRulesetOk() (Ruleset, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRuleset

`func (o *CampaignNotification) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.

### SetRuleset

`func (o *CampaignNotification) SetRuleset(v Ruleset)`

SetRuleset gets a reference to the given Ruleset and assigns it to the Ruleset field.

### GetOldRuleset

`func (o *CampaignNotification) GetOldRuleset() Ruleset`

GetOldRuleset returns the OldRuleset field if non-nil, zero value otherwise.

### GetOldRulesetOk

`func (o *CampaignNotification) GetOldRulesetOk() (Ruleset, bool)`

GetOldRulesetOk returns a tuple with the OldRuleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOldRuleset

`func (o *CampaignNotification) HasOldRuleset() bool`

HasOldRuleset returns a boolean if a field has been set.

### SetOldRuleset

`func (o *CampaignNotification) SetOldRuleset(v Ruleset)`

SetOldRuleset gets a reference to the given Ruleset and assigns it to the OldRuleset field.

### GetOldCampaign

`func (o *CampaignNotification) GetOldCampaign() Campaign`

GetOldCampaign returns the OldCampaign field if non-nil, zero value otherwise.

### GetOldCampaignOk

`func (o *CampaignNotification) GetOldCampaignOk() (Campaign, bool)`

GetOldCampaignOk returns a tuple with the OldCampaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOldCampaign

`func (o *CampaignNotification) HasOldCampaign() bool`

HasOldCampaign returns a boolean if a field has been set.

### SetOldCampaign

`func (o *CampaignNotification) SetOldCampaign(v Campaign)`

SetOldCampaign gets a reference to the given Campaign and assigns it to the OldCampaign field.

### GetEvaluationPosition

`func (o *CampaignNotification) GetEvaluationPosition() CampaignEvaluationPosition`

GetEvaluationPosition returns the EvaluationPosition field if non-nil, zero value otherwise.

### GetEvaluationPositionOk

`func (o *CampaignNotification) GetEvaluationPositionOk() (CampaignEvaluationPosition, bool)`

GetEvaluationPositionOk returns a tuple with the EvaluationPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvaluationPosition

`func (o *CampaignNotification) HasEvaluationPosition() bool`

HasEvaluationPosition returns a boolean if a field has been set.

### SetEvaluationPosition

`func (o *CampaignNotification) SetEvaluationPosition(v CampaignEvaluationPosition)`

SetEvaluationPosition gets a reference to the given CampaignEvaluationPosition and assigns it to the EvaluationPosition field.

### GetDeletedAt

`func (o *CampaignNotification) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CampaignNotification) GetDeletedAtOk() (time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeletedAt

`func (o *CampaignNotification) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAt

`func (o *CampaignNotification) SetDeletedAt(v time.Time)`

SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


