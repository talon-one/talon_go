# CampaignRulesetChangedNotificationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** | The type of the event. Can be one of the following: [&#39;campaign_state_changed&#39;, &#39;campaign_ruleset_changed&#39;, &#39;campaign_edited&#39;, &#39;campaign_created&#39;, &#39;campaign_deleted&#39;]  | 
**Campaign** | Pointer to [**Campaign**](Campaign.md) |  | 
**OldRuleset** | Pointer to [**Ruleset**](Ruleset.md) |  | [optional] 
**Ruleset** | Pointer to [**Ruleset**](Ruleset.md) |  | [optional] 

## Methods

### NewCampaignRulesetChangedNotificationItem

`func NewCampaignRulesetChangedNotificationItem(event string, campaign Campaign, ) *CampaignRulesetChangedNotificationItem`

NewCampaignRulesetChangedNotificationItem instantiates a new CampaignRulesetChangedNotificationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignRulesetChangedNotificationItemWithDefaults

`func NewCampaignRulesetChangedNotificationItemWithDefaults() *CampaignRulesetChangedNotificationItem`

NewCampaignRulesetChangedNotificationItemWithDefaults instantiates a new CampaignRulesetChangedNotificationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *CampaignRulesetChangedNotificationItem) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CampaignRulesetChangedNotificationItem) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CampaignRulesetChangedNotificationItem) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCampaign

`func (o *CampaignRulesetChangedNotificationItem) GetCampaign() Campaign`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *CampaignRulesetChangedNotificationItem) GetCampaignOk() (*Campaign, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *CampaignRulesetChangedNotificationItem) SetCampaign(v Campaign)`

SetCampaign sets Campaign field to given value.


### GetOldRuleset

`func (o *CampaignRulesetChangedNotificationItem) GetOldRuleset() Ruleset`

GetOldRuleset returns the OldRuleset field if non-nil, zero value otherwise.

### GetOldRulesetOk

`func (o *CampaignRulesetChangedNotificationItem) GetOldRulesetOk() (*Ruleset, bool)`

GetOldRulesetOk returns a tuple with the OldRuleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldRuleset

`func (o *CampaignRulesetChangedNotificationItem) SetOldRuleset(v Ruleset)`

SetOldRuleset sets OldRuleset field to given value.

### HasOldRuleset

`func (o *CampaignRulesetChangedNotificationItem) HasOldRuleset() bool`

HasOldRuleset returns a boolean if a field has been set.

### GetRuleset

`func (o *CampaignRulesetChangedNotificationItem) GetRuleset() Ruleset`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *CampaignRulesetChangedNotificationItem) GetRulesetOk() (*Ruleset, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *CampaignRulesetChangedNotificationItem) SetRuleset(v Ruleset)`

SetRuleset sets Ruleset field to given value.

### HasRuleset

`func (o *CampaignRulesetChangedNotificationItem) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


