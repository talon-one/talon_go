# StrikethroughEffect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | Pointer to **int32** | The ID of the campaign that effect belongs to. | 
**RulesetId** | Pointer to **int32** | The ID of the ruleset containing the rule that triggered this effect. | 
**RuleIndex** | Pointer to **int32** | The position of the rule that triggered this effect within the ruleset. | 
**RuleName** | Pointer to **string** | The name of the rule that triggered this effect. | 
**Type** | Pointer to **string** | The type of this effect. | 
**Props** | Pointer to [**map[string]interface{}**](.md) |  | 
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The start of the time frame where the effect is active in UTC. | [optional] 
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The end of the time frame where the effect is active in UTC. | [optional] 

## Methods

### GetCampaignId

`func (o *StrikethroughEffect) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *StrikethroughEffect) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *StrikethroughEffect) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *StrikethroughEffect) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.

### GetRulesetId

`func (o *StrikethroughEffect) GetRulesetId() int32`

GetRulesetId returns the RulesetId field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *StrikethroughEffect) GetRulesetIdOk() (int32, bool)`

GetRulesetIdOk returns a tuple with the RulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRulesetId

`func (o *StrikethroughEffect) HasRulesetId() bool`

HasRulesetId returns a boolean if a field has been set.

### SetRulesetId

`func (o *StrikethroughEffect) SetRulesetId(v int32)`

SetRulesetId gets a reference to the given int32 and assigns it to the RulesetId field.

### GetRuleIndex

`func (o *StrikethroughEffect) GetRuleIndex() int32`

GetRuleIndex returns the RuleIndex field if non-nil, zero value otherwise.

### GetRuleIndexOk

`func (o *StrikethroughEffect) GetRuleIndexOk() (int32, bool)`

GetRuleIndexOk returns a tuple with the RuleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRuleIndex

`func (o *StrikethroughEffect) HasRuleIndex() bool`

HasRuleIndex returns a boolean if a field has been set.

### SetRuleIndex

`func (o *StrikethroughEffect) SetRuleIndex(v int32)`

SetRuleIndex gets a reference to the given int32 and assigns it to the RuleIndex field.

### GetRuleName

`func (o *StrikethroughEffect) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *StrikethroughEffect) GetRuleNameOk() (string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRuleName

`func (o *StrikethroughEffect) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### SetRuleName

`func (o *StrikethroughEffect) SetRuleName(v string)`

SetRuleName gets a reference to the given string and assigns it to the RuleName field.

### GetType

`func (o *StrikethroughEffect) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StrikethroughEffect) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *StrikethroughEffect) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *StrikethroughEffect) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetProps

`func (o *StrikethroughEffect) GetProps() map[string]interface{}`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *StrikethroughEffect) GetPropsOk() (map[string]interface{}, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProps

`func (o *StrikethroughEffect) HasProps() bool`

HasProps returns a boolean if a field has been set.

### SetProps

`func (o *StrikethroughEffect) SetProps(v map[string]interface{})`

SetProps gets a reference to the given map[string]interface{} and assigns it to the Props field.

### GetStartTime

`func (o *StrikethroughEffect) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *StrikethroughEffect) GetStartTimeOk() (time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartTime

`func (o *StrikethroughEffect) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTime

`func (o *StrikethroughEffect) SetStartTime(v time.Time)`

SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.

### GetEndTime

`func (o *StrikethroughEffect) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *StrikethroughEffect) GetEndTimeOk() (time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndTime

`func (o *StrikethroughEffect) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTime

`func (o *StrikethroughEffect) SetEndTime(v time.Time)`

SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


